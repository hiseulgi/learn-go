// // * program 1: generate dummy data
// // berjalan selama:
// // 2023/07/02 17:19:52 done in 1.18737238 seconds

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"math/rand"
// 	"os"
// 	"path/filepath"
// 	"time"
// )

// const totalFile = 5000
// const keyLength = 2500

// type Transaction struct {
// 	IdCustomer   string `json:"idCustomer"`
// 	ProductPrice int    `json:"productPrice"`
// 	Amount       int    `json:"amount"`
// 	TotalPrice   int    `json:"totalPrice"`
// 	RandomKey    string `json:"randomKey"`
// }

// var tempPath = filepath.Join(os.Getenv("TEMP"), "dummy-data")

// func main() {
// 	log.Println("start")
// 	start := time.Now()

// 	generateFiles()

// 	duration := time.Since(start)
// 	log.Println("done in", duration.Seconds(), "seconds")
// }

// /* -------------------------------------------------------------------------- */

// var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

// func randomString(length int) string {
// 	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

// 	b := make([]rune, length)
// 	for i := range b {
// 		b[i] = letters[randomizer.Intn(len(letters))]
// 	}

// 	return string(b)
// }

// func randomNumber(limitMin, limitMax, multiplier int) int {
// 	number := randomizer.Intn(limitMax-limitMin) + limitMin

// 	return number * multiplier
// }

// func randomTransaction() Transaction {
// 	newTransaction := Transaction{
// 		IdCustomer:   randomString(10),
// 		ProductPrice: randomNumber(10, 1000, 1000),
// 		Amount:       randomNumber(1, 20, 1),
// 		RandomKey:    randomString(keyLength),
// 	}
// 	newTransaction.TotalPrice = newTransaction.ProductPrice * newTransaction.Amount

// 	return newTransaction
// }

// func generateFiles() {
// 	os.RemoveAll(tempPath)
// 	os.MkdirAll(tempPath, os.ModePerm)

// 	for i := 0; i < totalFile; i++ {
// 		filename := filepath.Join(tempPath, fmt.Sprintf("file-%d.json", i))
// 		transactionData := randomTransaction()

// 		content, err := json.Marshal(transactionData)
// 		if err != nil {
// 			log.Println("Error encoding json file", filename)
// 		}

// 		err = ioutil.WriteFile(filename, content, os.ModePerm)
// 		if err != nil {
// 			log.Println("Error writing file", filename)
// 		}

// 		if i%100 == 0 && i > 0 {
// 			log.Println(i, "files created")
// 		}
// 	}

// 	log.Printf("%d of total files created", totalFile)
// }
