/* -------------------------------------------------------------------------- */
/*                               // * interface                               */
/* -------------------------------------------------------------------------- */
// interface -> kumpulan definisi method yang tidak memiliki isi
// merupakan tipe data, zero value-nya adalah nil

// * embedded interface
// intinya interface di dalam interface, seperti pada struct

package main

import (
	"fmt"
	"math"
)

type calculate2d interface {
	area() float64
	circumference() float64
}

type calculate3d interface {
	volume() float64
}

type calculate interface {
	calculate2d
	calculate3d
}

type cube struct {
	side float64
}

func (c *cube) area() float64 {
	return math.Pow(c.side, 2) * 6
}

func (c *cube) circumference() float64 {
	return c.side * 12
}

func (c *cube) volume() float64 {
	return math.Pow(c.side, 3)
}

func main() {
	var solidGeometry calculate
	solidGeometry = &cube{6}

	fmt.Println(solidGeometry)
	fmt.Println(&solidGeometry)
	fmt.Println("circumference:", solidGeometry.circumference())
	fmt.Println("area:", solidGeometry.area())
	fmt.Println("volume:", solidGeometry.volume())
	fmt.Println()

	// ini nyoba aja, penasaran tentang method pointer
	// ternyata kalau pake ini sama aja antara objek biasa dan objek pointer
	// cuma kalau objek pointer ini perlu dideference terlebih dahulu ketika inisialisasi
	var kedua calculate
	kedua = &cube{4}
	fmt.Println(kedua)
	fmt.Println(&kedua)
}
