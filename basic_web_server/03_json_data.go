// * JSON DATA
// format data sejuta umat

package main

import (
	"encoding/json"
	"fmt"
)

// akses struct harus bersifat publik
// tag json:"Name", menandakan dia bakal nyari properti Name pada Json yang kemudian akan masuk ke properti FullName Struct
type User struct {
	FullName string `json:"Name"`
	Age      int
}

type Users []User

func main() {
	// jsonString for input
	jsonString := `{"Name":"asep","Age":27, "Class":3}`
	jsonData := []byte(jsonString)

	data := User{}

	// decode json data to struct object
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("name:", data.FullName)
	fmt.Println("age:", data.Age)

	/* -------------------------------------------------------------------------- */
	// * Decode Json ke interface{}
	// kalau misal butuh cepat, pakai ini aja
	// var data1 map[string]interface{}
	var data1 interface{}

	if err := json.Unmarshal(jsonData, &data1); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(data1)

	/* -------------------------------------------------------------------------- */
	// * Decode Array Json ke Array Objek
	jsonString1 := `[
		{"Name":"bagus","Age":22},
		{"Name":"soleh","Age":20}
	]`
	jsonData1 := []byte(jsonString1)
	datas := Users{}

	if err := json.Unmarshal(jsonData1, &datas); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(datas[0])
	fmt.Println(datas[1])

	/* -------------------------------------------------------------------------- */
	// * Encode Object to Json String
	// mengubah objek ke json
	datasEncoded, err := json.Marshal(datas)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	datasEncodedString := string(datasEncoded)
	fmt.Println(datasEncoded)
	fmt.Println(datasEncodedString)
}

// * Hasil dan Kesimpulan
// - Jika data berlebih (lebih banyak dari properti struct), maka data yang berlebih tidak diapa-apain. Tetap bisa decode
// - Jika data kurang (ada properti struct yang kosong / tidak ada pada json), maka properti struct tersebut akan diisi dengan nilai default variabelnya. Tetap bisa decode

// * Cara kerja decode
// - ambil data string
// - ubah data string ke slice of byte (ngambil nilai ascii-nya)
// - ubah slice of byte menyesuaikan objek struct

// * Cara kerja encode
// - ambil objek
// - ubah data ke byte (udh json)
// - klo mau di string tinggal panggil var baru dan method string(...)
