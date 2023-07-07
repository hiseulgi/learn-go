/// * Go Generics
// generic programming -> metode penulisan kode di mana tipe data dalam kode didefinisikan menggunakan tipe data yang tipe pastinya adalah dituliskan belakangan saat kode tersebut di-call
// (tipe datanya dipanggil ketika call pada params-nya)

// berbeda dengan interface{} karena ia akan membungkus data aslinya dan untuk mengakses datanya perlu type assertion (data.(int))
// sedangkan pada generic kita perlu mendefinisikan tipe data yang kompatibel untuk dipakai saat pemanggilan kode

package main

import "fmt"

// * fungsi non generics (only int)
// func Sum(numbers []int) int {
// 	var total int
// 	for _, e := range numbers {
// 		total += e
// 	}
// 	return total
// }

// * fungsi generics
// func FuncName[dataType <ComparableType>](params)
// fungsi ini dapat menerima int, float32, float64
func Sum[V int | float32 | float64](numbers []V) V {
	var total V
	for _, e := range numbers {
		total += e
	}
	return total
}

// * keyword comparable
// -> tipe data yang kompatibel dengan semua tipe yang ada
// ada dua params K dan V
// - K -> tipe map key yang kompatibel dengan semua tipe data
// - V -> tipe map value yang kompatibel dengan int dan float64
func SumNumber[K comparable, V int | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// * generic type constraint
// meringkas tipe data pada params generic function
// hanya bisa digunakan pada scope generic
type Number interface {
	int | int64 | float64
}

func SumNumbers2[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// * struct generic
type UserModel[T int | float64] struct {
	Name   string
	Scores []T
}

func main() {
	total1 := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(total1)

	total2 := Sum([]float32{1.4, 2.1, 3, 4.2, 5})
	fmt.Println(total2)

	// total3 := Sum([]float64{1, 2, 3.3, 4.3, 5.4})
	// cara pemanggilan juga bisa gini secara eksplisit
	total3 := Sum[float64]([]float64{1, 2, 3.3, 4.3, 5.4})
	fmt.Println(total3)

	/* -------------------------------------------------------------------------- */

	ints := map[string]int{"first": 23, "second": 54}
	floats := map[string]float64{"first": 23.11, "second": 3.56}

	fmt.Printf("generics sums with constrains: %v and %v\n", SumNumber(ints), SumNumber(floats))
	fmt.Printf("generics sums 2 with constrains: %v and %v\n", SumNumbers2(ints), SumNumbers2(floats))

	/* -------------------------------------------------------------------------- */

	var m1 UserModel[int]
	m1.Name = "asep"
	m1.Scores = []int{1, 2, 3}
	fmt.Println("scores:", m1.Scores)

	var m2 UserModel[float64]
	m2.Name = "udin"
	m2.Scores = []float64{0.5, 2.8, 3.1}
	fmt.Println("scores:", m2.Scores)
}
