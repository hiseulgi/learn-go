// // * program 2: process data
// // alur:
// // 1. baca data (read all json data)
// // 2. md5 id customer
// // 3. rename file ke md5 id customer
// // 3. perhitungan statistik
// // 5. pembuatan laporan
// //
// // berjalan selama:
// // 2023/07/02 17:20:03 done in 0.837925209 seconds

// package main

// import (
// 	"crypto/md5"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"time"
// )

// type Transaction struct {
// 	IdCustomer   string `json:"idCustomer"`
// 	ProductPrice int    `json:"productPrice"`
// 	Amount       int    `json:"amount"`
// 	TotalPrice   int    `json:"totalPrice"`
// 	RandomKey    string `json:"randomKey"`
// }

// type Report struct {
// 	SumOfSold               int
// 	SumOfTransactionPrice   int
// 	AverageTransactionPrice int
// }

// var ReportData Report

// var tempPath = filepath.Join(os.Getenv("TEMP"), "dummy-data")

// func main() {
// 	log.Println("start")
// 	start := time.Now()

// 	processData()

// 	duration := time.Since(start)
// 	log.Println("done in", duration.Seconds(), "seconds")
// }

// func processData() {
// 	// * 1. read all file
// 	files, err := ioutil.ReadDir(tempPath)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	for i, file := range files {
// 		if !file.IsDir() {
// 			filePath := filepath.Join(tempPath, file.Name())
// 			fileContent, err := ioutil.ReadFile(filePath)
// 			if err != nil {
// 				fmt.Println("Error:", err)
// 				continue
// 			}

// 			fileJson := Transaction{}
// 			err = json.Unmarshal(fileContent, &fileJson)
// 			if err != nil {
// 				fmt.Println("Error:", err)
// 				continue
// 			}

// 			// * 2. md5 id customer
// 			sum := fmt.Sprintf("%x", md5.Sum([]byte(fileJson.RandomKey)))

// 			// * 3. rename data
// 			newFileName := fmt.Sprintf("file-%s.json", sum)
// 			newFilePath := filepath.Join(tempPath, newFileName)

// 			err = os.Rename(filePath, newFilePath)
// 			if err != nil {
// 				fmt.Println("Error:", err)
// 				continue
// 			}

// 			// * 4. perhitungan statistik
// 			ReportData.SumOfSold += fileJson.Amount
// 			ReportData.SumOfTransactionPrice += fileJson.TotalPrice
// 		}

// 		if i%100 == 0 && i > 0 {
// 			log.Println(i, "files proceed")
// 		}
// 	}

// 	ReportData.AverageTransactionPrice = ReportData.SumOfTransactionPrice / ReportData.SumOfSold

// 	// * 5. pembuatan laporan
// 	filename := "transaction-report.json"

// 	content, err := json.Marshal(ReportData)
// 	if err != nil {
// 		log.Println("Error encoding json file", filename)
// 	}

// 	err = ioutil.WriteFile(filename, content, os.ModePerm)
// 	if err != nil {
// 		log.Println("Error writing file", filename)
// 	}
// }
