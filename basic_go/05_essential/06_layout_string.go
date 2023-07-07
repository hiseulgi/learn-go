// * Layout Format String

package main

import "fmt"

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

var (
	data = student{
		name:        "wick",
		height:      170.5,
		age:         21,
		isGraduated: true,
		hobbies:     []string{"eating", "anime"},
	}
)

func main() {
	// biner
	fmt.Printf("%b\n", data.age)

	// numerik basis 10
	fmt.Printf("%d\n", data.age)

	// numerik basis 8
	fmt.Printf("%o\n", data.age)

	// numerik basis 16
	fmt.Printf("%x\n", data.age)

	// numerik desimal dalam bentuk notasi numerik (scientific notation)
	fmt.Printf("%e\n", data.height)

	// float / numerik desimal
	fmt.Printf("%f\n", data.height)
	fmt.Printf("%.2f\n", data.height)
	fmt.Printf("%.5f\n", data.height)
	fmt.Printf("%.f\n", data.height)

	// numerik decimal (%g)
	// bedanya, lebar digit desimal adalah sesuai dengan datanya
	fmt.Printf("%g\n", 0.12)
	fmt.Printf("%.5g\n", 0.12)

	// alamat pointer
	fmt.Printf("%p\n", &data.name)

	// unicode
	fmt.Printf("%c\n", 1400)

	// escape string
	fmt.Printf("%q\n", `" name \ height "`)

	// string
	fmt.Printf("%s\n", data.name)

	// boolean
	fmt.Printf("%t\n", data.isGraduated)

	// mengambil tipe variabel
	fmt.Printf("%T\n", data.isGraduated)
	fmt.Printf("%T\n", data.height)
	fmt.Printf("%T\n", data.hobbies)

	/* -------------------------------------------------------------------------- */
	// format data apa saja (interface{})
	fmt.Printf("%v\n", data)

	// format struct dengan nama property dan valuenya
	fmt.Printf("%+v\n", data)

	// format struct dan objeknya
	fmt.Printf("%#v\n", data)

	// menulis persen (%) pada print
	fmt.Printf("%%\n")
}
