// * vendoring
// kapabilitas untuk mengunduh semua dependency / 3rd party, untuk disimpan di lokal dalam folder project (vendor)
// dengan adanya vendor, go dapat langsung menggunakan dependency tanpa harus download lagi di internet

package main

import (
	"fmt"

	"github.com/novalagung/gubrak/v2"
)

func main() {
	fmt.Println(gubrak.RandomInt(10, 20))
}

// * manfaat vendoring
// sisi kompatibilitas dan kestabilan 3rd party
// jadi dengan vendor, misal 3rd party yang kita gunakan ada update yang tidak backward compatible. maka aplikasi kita tetap aman karena yang digunakan ada di folder vendor
// apakah wajib? opsional sesuai kebutuhan
