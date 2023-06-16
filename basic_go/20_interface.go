/* -------------------------------------------------------------------------- */
/*                               // * interface                               */
/* -------------------------------------------------------------------------- */
// interface -> kumpulan definisi method yang tidak memiliki isi
// merupakan tipe data, zero value-nya adalah nil

package main

import (
	"fmt"
	"math"
)

// * interface calculate method for plane geometry
type calculate interface {
	area() float64
	circumference() float64
}

// * square struct
type square struct {
	side float64
}

func (s square) area() float64 {
	return math.Pow(s.side, 2)
}

func (s square) circumference() float64 {
	return s.side * 4
}

// * circle struct
type circle struct {
	diameter float64
}

func (c circle) radius() float64 {
	return 0.5 * c.diameter
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius(), 2)
}

func (c circle) circumference() float64 {
	return 2 * math.Pi * c.radius()
}

func main() {
	// deklarasi var dengan interface calculate
	// var ini menampung blueprint dari method area dan circumference, masih bernilai nil
	var planeGeometry calculate

	// var ini menjadi objek square
	planeGeometry = square{8.0}
	fmt.Println("~ square")
	fmt.Println("area :", planeGeometry.area())
	fmt.Println("circumference :", planeGeometry.circumference())

	// var ini menjadi objek circle
	planeGeometry = circle{14.0}
	fmt.Println("~ circle")
	fmt.Println("area :", planeGeometry.area())
	fmt.Println("circumference :", planeGeometry.circumference())
	fmt.Println("radius :", planeGeometry.(circle).radius()) // karena radius tidak ada dalam interface, kita perlu memanggil struct terlebih dahulu agar dapat dipanggil method di dalamnya
	// atau istilahnya di-casting terelbih dahulu menjadi objek circle
}
