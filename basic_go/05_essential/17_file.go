// * File
// teknik dasar untuk operasi file pada Go

package main

import (
	"fmt"
	"io"
	"os"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

// * new file cara 1
// func createFile(path string) {
// 	// detek apakah sudah ada file
// 	if _, err := os.Stat(path); err == nil {
// 		fmt.Println(err)
// 		fmt.Println("File exists")
// 	} else {
// 		// buat file baru jika belum ada
// 		file, err := os.Create(path)
// 		if isError(err) {
// 			return
// 		}
// 		defer file.Close()
// 		fmt.Println("New file has been created", path)
// 	}
// }

// * new file cara 2
// pake os.IsNotExist()
func createFile(path string) {
	// deteksi apakah file sudah ada
	// err bernilai nil jika file sudah ada
	var _, err = os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		// membuat file baru
		var file, err = os.Create(path)
		// jangan lupa setiap ada operasi pada file, ditutup juga
		defer file.Close()

		if isError(err) {
			return
		}

		fmt.Println("==> file berhasil dibuat", path)
		return
	}

	fmt.Println("failed! file sudah ada")
}

// * write isi file
func writeFile(path string) {
	// buka file dengan level akses RW
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// menulis data ke file
	_, err = file.WriteString("halo\n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("nama saya sugab\n")
	if isError(err) {
		return
	}

	// save file
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("file berhasil ditulis")
}

// * read file
func readFile(path string) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// read file
	var text = make([]byte, 1024)
	// proses pembacaan berurutan dari baris ke baris
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}
		// berhenti jika byte yang dibaca lebarnya = 0
		if n == 0 {
			break
		}
	}
	if isError(err) {
		return
	}

	fmt.Println("file berhasil dibaca")
	fmt.Println(string(text))
}

// * menghapus file
func deleteFile(path string) {
	err := os.Remove(path)
	if isError(err) {
		return
	}
	fmt.Println("file berhasil dihapus")
}

func main() {
	path := "/home/sugab/workspace/go/learn-go/test.txt"
	// createFile(path)
	// writeFile(path)
	// readFile(path)
	deleteFile(path)
}
