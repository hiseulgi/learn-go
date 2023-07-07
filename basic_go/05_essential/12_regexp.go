// * Regex
// regular expression -> teknik yang digunakan untuk pencocokan string dengan pola tertentu
// go menggunakan sintaks regex dari
// https://github.com/google/re2/wiki/Syntax

package main

import (
	"fmt"
	"regexp"
)

func main() {
	// * 0. penerapan regexp
	text := "banana burger soup"
	regex, err := regexp.Compile(`[a-z]+`)
	if err != nil {
		fmt.Println(err.Error())
	}

	res1 := regex.FindAllString(text, 2)
	fmt.Println(res1)
	res2 := regex.FindAllString(text, -1)
	fmt.Println(res2)

	// ekspresi [a-z]+ -> semua string yang hurufnya kecil
	// method FindAllString() -> mencari semua string yang sesuai dengan ekpresi regex, dengan nilai return berupa slice string
	// jumlah pencariannya bisa ditentukan (pada param kedua)

	// * 1. MatchString()
	// mendeteksi string sesuai pola regexp
	isMatch := regex.MatchString(text)
	fmt.Println(isMatch)

	// * 2. FindString()
	// mencari string yang sesuai pola regexp
	str := regex.FindString(text)
	fmt.Println(str)

	// return hanya 1 hasil saja (hasil pertama)

	// *  3. FindStringIndex()
	// mencari index string hasil return dari operasi regexp
	idx := regex.FindStringIndex(text)
	fmt.Println(idx)
	fmt.Println(str[:6])

	// * 4. ReplaceAllString()
	// replace semau string yang memenuhi kriteria regexp, dengan string lain
	str1 := regex.ReplaceAllString(text, "nabati")
	fmt.Println(str1)

	// * 5. ReplaceAllStringFunc()
	// replace semua string yang memenuhi kriteria regexp, dengan kondisi tertentu untuk setiap substring (dengan func)
	str2 := regex.ReplaceAllStringFunc(text, func(s string) string {
		if s == "burger" {
			return "hotdog"
		}
		return s
	})
	fmt.Println(str2)

	// * 5. Split()
	// memisah string dengan pemisah substring yang memenuhi kriteria regexp
	regex1, _ := regexp.Compile(`[a-b]+`) // separator
	str3 := regex1.Split(text, -1)        // -1 : semua karakter yang memenuhi kriteria regex akan menjadi separator
	// misal 2 : hanya 2 karakter pertama saja
	fmt.Println(str3)

}
