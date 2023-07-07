// * program 2: process data secara konkurensi
// alur:
// 1. baca data (read all json data)
// 2. md5 id customer
// 3. rename file ke md5 id customer
// 3. perhitungan statistik
// 5. pembuatan laporan
//
// berjalan selama:
// 2023/07/02 18:06:47 done in 0.580669207 seconds
// 30,73% lebih cepat

package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Transaction struct {
	IdCustomer   string `json:"idCustomer"`
	ProductPrice int    `json:"productPrice"`
	Amount       int    `json:"amount"`
	TotalPrice   int    `json:"totalPrice"`
	RandomKey    string `json:"randomKey"`
}

type Report struct {
	Mu                      sync.Mutex
	SumOfSold               int
	SumOfTransactionPrice   int
	AverageTransactionPrice int
}

type FileInfo struct {
	FilePath     string // file location
	ContentData  Transaction
	MD5Sum       string // md5 sum of content
	IsRenamed    bool   // indicator when file is renamed already
	IsCalculated bool   // indicator when file is calculated already
}

var ReportData Report
var tempPath = filepath.Join(os.Getenv("TEMP"), "dummy-data")

func main() {
	log.Println("start")
	start := time.Now()

	// 1. baca data (read all json data)
	chanFileContent := readFiles()

	// 2. md5 id customer
	chanFileSum1 := getSum(chanFileContent)
	chanFileSum2 := getSum(chanFileContent)
	chanFileSum3 := getSum(chanFileContent)
	chanFileSum := mergeChanFileInfo(chanFileSum1, chanFileSum2, chanFileSum3)

	// 3. rename file ke md5 id customer
	chanRename1 := renameFile(chanFileSum)
	chanRename2 := renameFile(chanFileSum)
	chanRename3 := renameFile(chanFileSum)
	chanRename := mergeChanFileInfo(chanRename1, chanRename2, chanRename3)

	// 4. perhitungan statistik
	chanCalculate1 := calculateReport(chanRename)
	chanCalculate2 := calculateReport(chanRename)
	chanCalculate3 := calculateReport(chanRename)
	chanCalculate := mergeChanFileInfo(chanCalculate1, chanCalculate2, chanCalculate3)

	// print output
	counterCalculated := 0
	counterTotal := 0
	for dataFileInfo := range chanCalculate {
		if dataFileInfo.IsCalculated {
			counterCalculated++
		}
		counterTotal++
	}
	log.Printf("%d/%d files calculated", counterCalculated, counterTotal)

	// 5. pembuatan laporan
	go func() {
		<-chanCalculate
		makeReport(chanCalculate)
	}()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

/* -------------------------------------------------------------------------- */
// * 1. baca data (read all json data)
func readFiles() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		files, err := ioutil.ReadDir(tempPath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for _, file := range files {
			if !file.IsDir() {
				filePath := filepath.Join(tempPath, file.Name())
				fileContent, err := ioutil.ReadFile(filePath)
				if err != nil {
					fmt.Println("Error:", err)
					continue
				}

				fileJson := Transaction{}
				err = json.Unmarshal(fileContent, &fileJson)
				if err != nil {
					fmt.Println("Error:", err)
					continue
				}

				chanOut <- FileInfo{
					FilePath:    filePath,
					ContentData: fileJson,
				}
			}
		}

		close(chanOut)
	}()

	return chanOut
}

// * 2. md5 id customer
// fan-out
func getSum(chanIn <-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for fileInfo := range chanIn {

			fileInfo.MD5Sum = fmt.Sprintf("%x", md5.Sum([]byte(fileInfo.ContentData.RandomKey)))
			chanOut <- fileInfo
		}
		close(chanOut)
	}()

	return chanOut
}

// * 3. rename file ke md5 id customer
// fan-out
func renameFile(chanIn <-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for dataFileInfo := range chanIn {
			newFileName := fmt.Sprintf("file-%s.json", dataFileInfo.MD5Sum)
			newFilePath := filepath.Join(tempPath, newFileName)

			err := os.Rename(dataFileInfo.FilePath, newFilePath)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			dataFileInfo.IsRenamed = err == nil
			chanOut <- dataFileInfo
		}

		close(chanOut)
	}()

	return chanOut
}

// * 4. perhitungan statistik
func calculateReport(chanIn <-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for dataFileInfo := range chanIn {
			ReportData.Mu.Lock()

			ReportData.SumOfSold += dataFileInfo.ContentData.Amount
			ReportData.SumOfTransactionPrice += dataFileInfo.ContentData.TotalPrice

			ReportData.Mu.Unlock()

			dataFileInfo.IsCalculated = true
			chanOut <- dataFileInfo
		}
		close(chanOut)
	}()

	return chanOut
}

// * 5 pembuatan laporan
func makeReport(chanIn <-chan FileInfo) {
	// notifier menunggu semua proses selesai
	<-chanIn

	ReportData.AverageTransactionPrice = ReportData.SumOfTransactionPrice / ReportData.SumOfSold

	filename := "transaction-report.json"

	content, err := json.Marshal(Report{
		SumOfSold:               ReportData.SumOfSold,
		SumOfTransactionPrice:   ReportData.SumOfTransactionPrice,
		AverageTransactionPrice: ReportData.AverageTransactionPrice,
	})

	if err != nil {
		log.Println("Error encoding json file", filename)
	}

	err = ioutil.WriteFile(filename, content, os.ModePerm)
	if err != nil {
		log.Println("Error writing file", filename)
	}
}

// * fungsi pembantu: merge file info from many channel worker into one channel
// fan-in
func mergeChanFileInfo(chainInMany ...<-chan FileInfo) <-chan FileInfo {
	wg := new(sync.WaitGroup)
	chanOut := make(chan FileInfo)

	wg.Add(len(chainInMany))
	for _, eachChan := range chainInMany {
		go func(eachChan <-chan FileInfo) {
			for eachChanData := range eachChan {
				chanOut <- eachChanData
			}
			wg.Done()
		}(eachChan)
	}

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}
