/* -------------------------------------------------------------------------- */
/*                                // * reflect                                */
/* -------------------------------------------------------------------------- */
// teknik untuk inspeksi variabel, mengambil info dari variabel atau bahlan memanipulasinya
// bisa melihat struktur variabel, tipe, nilai pointer dll

// reflect yang biasa dipakai adalah ValueOf() dan TypeOf()
// ValueOf() mengambil nilainya
// TypeOf() mengambil tipe datanya

package main

import (
	"fmt"
	"reflect"
)

func main() {
	myNumber := 69

	reflectValue := reflect.ValueOf(myNumber) // return objek dalam tipe reflect.Value
	reflectType := reflect.TypeOf(myNumber)   // return string
	fmt.Println("tipe data myNumber:", reflectType)
	fmt.Println(reflectValue.Kind()) // ini sama mengambil tipe variabel juga, tapi dia lebih buat ke komparasi nilai

	// cek apakah myNumber adalah int
	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai myNumber:", reflectValue)
	}

	// * akses nilai dalam bentuk interface{}
	// karena interface secara default ketika dicetak adalah string. jadi cocok untuk segala jenis nilai
	fmt.Println(reflectValue.Interface())

	// * akses informasi property variabel objek
	// bisa melihat dan memanip informasi pada property
	s1 := &student{"asep", 3}
	s1.getPropertyInfo()

	// * akses informasi method variabel objek
	// sama kek diatas, bisa mengambil dengan cara mereflect methodnya ke variabel baru dan mencobanya dengan reflect juga
	s1ReflectValue := reflect.ValueOf(s1) // mengambil objek reflect dari s1

	method := s1ReflectValue.MethodByName("SetName") // mengambil method SetName dari objek reflect s1
	method.Call([]reflect.Value{
		reflect.ValueOf("daril"),
	}) // memanggil method reflect dengan parameter reflect juga
	s1.getPropertyInfo()
}

// * akses informasi property variabel objek
type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	// mengambil nilai reflect pada objek
	reflectValue := reflect.ValueOf(s)

	// mengambil tipe objek reflect apakah pointer
	if reflectValue.Kind() == reflect.Pointer {
		// mengambil objek reflect aslinya
		reflectValue = reflectValue.Elem()
	}

	// mengambil tipe data nya
	reflectType := reflectValue.Type()

	// loop setiap property pada objek
	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama:", reflectType.Field(i).Name)          // nama property
		fmt.Println("tipe data:", reflectType.Field(i).Type)     // tipe property
		fmt.Println("nilai:", reflectValue.Field(i).Interface()) // nilai property
		fmt.Println()
	}
}

// * akses informasi method variabel objek
func (s *student) SetName(name string) {
	s.Name = name
}
