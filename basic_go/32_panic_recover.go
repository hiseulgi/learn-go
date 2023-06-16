/* -------------------------------------------------------------------------- */
/*                         // * error, panic, recover                         */
/* -------------------------------------------------------------------------- */
// panic -> memunculkan panic error
// recover -> untuk menghandle-nya

// error adalah tipe data, isinya dapat nil
// error biasanya merupakan nilai return dari suatu fungsi / method pada go, jadi sangat penting untuk membuat handlingnya nanti

package main

import (
	"errors"
	"fmt"
	"strings"
)

// * panic
// -> menampilkan stack trace error sekaligus menghentikan flow goroutine (main() merupakan goroutine)
// semua proses akan berhenti kecuali

func validate(input string) (bool, error) {
	if strings.Trim(input, " ") == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

// * recover -> sebagai catch
// handling panic error, akan recover take-over goroutine yang sedang panic (pessan panic tidak akan muncul)

func catch() {
	if r := recover(); r != nil {
		fmt.Println("error occured:", r)
	} else {
		fmt.Println("app running perfectly")
	}
}

func main() {
	// penggunaan recover dengan memanggil fungsi eksternal
	// defer catch()

	// penggunaan recover dengan closure IIFE
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error occured:", r)
		} else {
			fmt.Println("app running perfectly")
		}
	}()

	name := ""
	fmt.Print("input nama: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("helo", name)
	} else {
		panic(err.Error())
		// fmt.Println("end") // unreachable code
	}
}

// kadangkala kita butuh recover panic pada setiap looping, jadi kita tidak akan menaruh defer pada scope main()
// tujuannya adalah agar seluruh program tidak berhenti, dan tetap melanjutkan loopingnya
// maka dari itu akan membuat IIFE itu di dalam loop tersendiri

// func main() {
// 	data := []string{"superman", "aquaman", "wonder woman"}
// 	for _, each := range data {
// 		func() {
// 			// recover untuk IIFE dalam perulangan
// 			defer func() {
// 				if r := recover(); r != nil {
// 					fmt.Println("Panic occured on looping", each, "| message:", r)
// 				} else {
// 					fmt.Println("Application running perfectly")
// 				}
// 			}()
// 			panic("some error happen")
// 		}()
// 	}
// }
