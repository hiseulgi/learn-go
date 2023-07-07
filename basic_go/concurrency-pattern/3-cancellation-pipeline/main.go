// * Context Cancellation Pipeline
// cancellation -> mekanisme untuk menggagalkan secara paksa proses konkuren yang sedang berjalan (timeout, error, faktor lain)

// penerapannya adalah dengan context.Context
// context berisi informasi deadlines, signal cancellation, dan data

// studi kasus: generate dummy files secara concurrently
// ada dua result nanti
// - proses sukses, karena execution time di bawah timeout
// - proses digagalkan secara paksa ditengah jalan, karena running time sudah melebihi batas timeout

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const totalFile = 3000
const contentLength = 5000
const timeoutDuration = 3 * time.Second

var tempPath = filepath.Join(os.Getenv("TEMP"), "worker-pool")

// meta data untuk pembuatan dummy file
type FileInfo struct {
	Index       int
	FileName    string
	WorkerIndex int
	Err         error
}

func main() {
	log.Println("start")
	start := time.Now()

	// untuk penggunaan method context ada banyak selain timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel() // wajib di cancel secara manual di akhir
	generateFilesWithContext(ctx)

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

/* -------------------------------------------------------------------------- */

func randomString(length int) string {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}

func generateFilesWithContext(ctx context.Context) {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	done := make(chan int)

	go func() {
		// pipeline 1: job distribution
		// dispatch goroutine untuk distribusi jobs
		chanFileIndex := generateFileIndexes(ctx)

		// pipeline  2: main logic (create file)
		// dispatch goroutine untuk start worker untuk melakukan pembuatan file (createFiles)
		createFilesWorker := 100
		chanFileResult := createFiles(ctx, chanFileIndex, createFilesWorker)

		// track and print output
		// tracking channel dari fan-in nilai return pipeline 2
		counterSuccess := 0
		for fileResult := range chanFileResult {
			if fileResult.Err != nil {
				log.Printf("error creating file %s. stack trace: %s", fileResult.FileName, fileResult.Err)
			} else {
				counterSuccess++
			}
		}

		// notify that the process is complete
		done <- counterSuccess
	}()

	select {
	case <-ctx.Done():
		log.Printf("generation process stopped. %s", ctx.Err())
	case counterSuccess := <-done:
		log.Printf("%d/%d of total files created", counterSuccess, totalFile)
	}
}

// media komunikasi antara proses distribusi jobs dengan proses fungsi createFiles()
// return chanFileIndex <-chan FileInfo
func generateFileIndexes(ctx context.Context) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	// make new metadata FileInfo then output it
	go func() {
		for i := 0; i < totalFile; i++ {
			select {
			case <-ctx.Done():
				break
			default:
				chanOut <- FileInfo{
					Index:    i,
					FileName: fmt.Sprintf("file-%d.txt", i),
				}
			}
		}
		close(chanOut)
	}()

	return chanOut
}

// fungsi gabungan antara fan-out fan-in
// return chanFileResult <-chan FileInfo
func createFiles(ctx context.Context, chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	// waitgroup to control the workers
	wg := new(sync.WaitGroup)

	// allocate N of workers
	wg.Add(numberOfWorkers)

	go func() {

		// dispatch N workers
		// (membagi pekerja untuk proses async)
		// kita akan membuat proses sebanyak numberOfWorkers agar dapat berjalan secara paralel. masing-masing proses tersebut akan listen dari channel chanIn. Dengan demikian, pembagian pekerjaan akan rata.
		for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
			go func(workerIndex int) {

				// listen to `chainIn` channel for incoming jobs (dari fungsi sebelumnya)
				for job := range chanIn {
					select {
					case <-ctx.Done():
						break
					default:
						// creating file
						filePath := filepath.Join(tempPath, job.FileName)
						content := randomString(contentLength)
						err := ioutil.WriteFile(filePath, []byte(content), os.ModePerm)

						log.Println("worker", workerIndex, "working on", job.FileName)

						// make new job result and send it to `chanOut`
						chanOut <- FileInfo{
							FileName:    job.FileName,
							WorkerIndex: workerIndex,
							Err:         err,
						}
					}

				}

				// mark down waitgroup if `chanIn is closed` and all jobs are finished
				wg.Done()
			}(workerIndex)
		}
	}()

	// wait until `chanIn` close and then all workers are done (notifier for close `chanOut`)
	go func() {
		wg.Wait()
		close(chanOut)
	}()
	return chanOut
}
