// * konversi antar tipe data

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// * fungsi strconv.Atoi()
	// konversi string ke int
	str1 := "157"
	num1, err1 := strconv.Atoi(str1)

	if err1 == nil {
		fmt.Println(num1)
	}

	// * fungsi strconv.Itoa()
	// konversi int ke string
	num2 := 169
	str2 := strconv.Itoa(num2)
	fmt.Println(str2)

	// * fungsi strconv.Parseint()
	// parsing string ke basis data tipe numerik non-desimal dengan lebar data yang bisa ditentukan
	str3 := "420169"
	// parameter; data string, basis bilangan pada string, lebar tipe data
	num3, err3 := strconv.ParseInt(str3, 10, 64)
	if err3 == nil {
		fmt.Println(num3)
	}

	// contoh lain ke basis 2 (biner)
	str3bin := "1010"
	num3bin, err3bin := strconv.ParseInt(str3bin, 2, 16)
	if err3bin == nil {
		fmt.Println(num3bin)
	}

	// * fungsi strconv.FormatInt()
	// konversi data numerik int64 ke string dengan basis numerik yang bisa ditentukan
	num4 := int64(24)
	str4 := strconv.FormatInt(num4, 10)
	fmt.Println(str4)

	// * fungsi strconv.ParseFloat()
	// konversi string ke float dengan lebar data yang bisa ditentukan
	// float32 atau float64
	str5 := "14.17"
	num5, err5 := strconv.ParseFloat(str5, 32)
	if err5 == nil {
		fmt.Println(num5)
	}

	// * fungsi strconv.FormatFloat()
	// konversi data float64 ke string dengan format eksponen, lebar digit desimal, dan lebar tipe data bisa ditentukan
	// parameter; num dalam float64, format eksponen (bisa liat di dokumentasi), lebar floating point, tipe data
	num6 := float64(51.57)
	str6 := strconv.FormatFloat(num6, 'f', 4, 64)
	fmt.Println(str6)

	// * fungsi strconv.ParseBool()
	// konversi string ke boolean
	str7 := "true"
	boolValue, err7 := strconv.ParseBool(str7)
	if err7 == nil {
		fmt.Println(boolValue)
	}

	// * fungsi strconv.FormatBool()
	boolValue2 := false
	str8 := strconv.FormatBool(boolValue2)
	fmt.Println(str8)

	/* -------------------------------------------------------------------------- */
	// * konversi data dengan teknik casting
	// keyword data bisa digunakan untuk casting
	// cara penggunaanya tinggal panggil tipe data tujuan casting seabgai fungsi, dan parameternya adalah input datanya
	a := float64(24)
	b := int32(29.00)
	fmt.Println(a, " ", b)

	// * casting string <-> byte
	// string sebenarnya adalah slice of byte
	// di Go sebuah char direpresentasikan oleh sebuah elemen slice byte
	// Tiap elemen slice berisi data int yang merupakan kode ASCII dari char dalam string
	// untuk mendapatkan slice byte dari data string -> meng-casting-nya ke tipe []byte
	text1 := "sugab"
	byte1 := []byte(text1)
	fmt.Println(byte1) // berisi int kode ASCII

	byte2 := []byte{104, 97, 108, 111}
	text2 := string(byte2)
	fmt.Println(text2)

	// selain itu char dalam string dapat kita lihat bentuk byte nya per karakter
	c := int64('h')
	fmt.Println(c)

	d := string(104)
	fmt.Println(d)

	// * type assertions pada interface kosong (interface{})
	// -> teknik untuk mengambil tipe data konkret dari data yang terbungkus interface{}
	// jadi hanya bisa dilakukan pada data bertipe interface{}
	// tujuannya adalah untuk mengambil tipe data yang sesuai pada suatu interface{}, jika tidak sesuai maka akan terjadi panic error
	data := map[string]any{
		"name":    "sugab",
		"grade":   2,
		"height":  169.5,
		"isMale":  true,
		"hobbies": []string{"eating", "anime"},
	}

	fmt.Println(data["name"].(string))
	// fmt.Println(data["name"].(int)) // ! bakal error, karena tipe datanya tidak sesuai
	fmt.Println(data["grade"].(int))
	fmt.Println(data["height"].(float64))
	fmt.Println(data["isMale"].(bool))
	fmt.Println(data["hobbies"].([]string))

	// melihat tipe data asli pada interface{}, jika kita tidak tahu tipe data aslinya
	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		case float64:
			fmt.Println(val.(float64))
		case bool:
			fmt.Println(val.(bool))
		case []string:
			fmt.Println(val.([]string))
		default:
			fmt.Println(val.(int))
		}
	}
}
