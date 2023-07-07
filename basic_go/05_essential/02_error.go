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
	"strconv"
	"strings"
)

func main() {
	// * simple error handling from built-in lib
	input := ""
	fmt.Print("masukkan angka: ")
	fmt.Scanln(&input)

	number := 0
	var err error

	// check is number or not
	// if there any error, then go on scope
	number, err = strconv.Atoi(input)
	if err == nil {
		fmt.Println(number, "adalah angka")
	} else {
		fmt.Println(input, "adalah bukan angka")
		fmt.Println(err.Error())
	}
	fmt.Println()

	// * custom error
	// ini pake method errors.New("customMessage")
	name := ""
	fmt.Print("masukkan nama: ")
	fmt.Scanln(&name)

	// untuk error handling biasanya pake struktur if else gini langsung
	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		fmt.Println(err.Error())
	}
}

// * func with custom error
func validate(input string) (bool, error) {
	if strings.Trim(input, " ") == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}
