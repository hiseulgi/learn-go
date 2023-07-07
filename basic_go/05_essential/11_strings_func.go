// * fungsi string
// mengenal fungsi bawaan ada strings

package main

import (
	"fmt"
	"strings"
)

func main() {
	// * strings.Contains()
	// cek apakah suatu kata terdapat pada kata lain
	// kata/kalimat input; kata yang dicari
	isExists := strings.Contains("sugab ida", "ida")
	fmt.Println(isExists)

	// * strings.HasPrefix()
	// cek string diawali oleh string tertentu (awalan kata)
	isPrefix1 := strings.HasPrefix("sugab", "gab")
	fmt.Println(isPrefix1)

	// * strings.HasSuffix()
	// cek string diakhiri oleh string tertentu (akhiran kata)
	isSuffix1 := strings.HasSuffix("sugab", "gab")
	fmt.Println(isSuffix1)

	// * strings.Count()
	// menghitung jumlah karakter tertentu (parameter kedua)
	howMany := strings.Count("kobo kanaeru", "o")
	fmt.Println(howMany)

	// * strings.Index()
	// mencari posisi indeks sebuah string (parameter kedua)
	index1 := strings.Index("sugab ida", "a")
	fmt.Println(index1) // mengembalikan indeks yang paling awal

	// * strings.Replace()
	// mereplace string dengan string lain
	text := "bokobokobo"
	find := "o"
	replaceWith := "e"

	newText1 := strings.Replace(text, find, replaceWith, 1) // angka terakhir adalah berapa banyak yang diganti, dimula dari index paling awal
	fmt.Println(newText1)
	newText2 := strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2)
	newText3 := strings.Replace(text, find, replaceWith, -1) // semua substring
	fmt.Println(newText3)

	// * strings.Repeat()
	// mengulang string sebanyak data yang ditentukan
	str := strings.Repeat("lo", 4)
	fmt.Println(str)

	// * strings.Split()
	// memisahkan string dengan tanda pemisah yang bisa ditentukan; hasilnya slice of string
	string1 := strings.Split("Kobo Kanaeru Wibu", " ")
	fmt.Println(string1)
	string2 := strings.Split("sugab", "")
	fmt.Println(string2)

	// * strings.Join()
	// menggabungkan slice string menjadi string dengan pemisah tertentu
	data := []string{"banana", "papaya", "tomato"}
	dataJoin := strings.Join(data, "-")
	fmt.Println(dataJoin)

	// * strings.ToLower()
	// mengubah huruf string menjadi huruf kecil
	str1 := strings.ToLower("bAgUS")
	fmt.Println(str1)

	// * strings.ToUpper()
	// mengubah huruf string menjadi huruf besar
	str2 := strings.ToUpper("madang!")
	fmt.Println(str2)
}
