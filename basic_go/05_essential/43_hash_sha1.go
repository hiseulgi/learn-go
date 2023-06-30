// * Hash SHA1
// Hash : algoritma enkripsi mengubah text ke deretean char acak
// jumlah char hasil hash selalu sama
// one-way encryption : hasil dari hash tidak bisa dikembalikan ke text asli
//
// SHA1 (Secure Hash Algorithm 1)
// algoritma hashing untuk enkripsi data
// hasil sha1 : data dengan lebar 20 byte (160 bit), biasa ditampilkan dalam bentuk bilangan heksadesimal 40 digit

package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func main() {
	// * penerapan hash sha1

	text := "sugab adalah rahasia"
	sha := sha1.New()                               // objek sha1 baru
	sha.Write([]byte(text))                         // set data yang akan di-hash (harus []byte)
	encrypted := sha.Sum(nil)                       // eksekusi proses hash, return hasil data ter-hash
	encryptedString := fmt.Sprintf("%x", encrypted) // format hash ke hex

	fmt.Println(text)
	fmt.Println(encrypted)
	fmt.Println(string(encrypted))
	fmt.Println(encryptedString)
	fmt.Println()

	/* -------------------------------------------------------------------------- */

	// * metode salting pada hash sha1
	// salt -> data acak yang digabungkan pada data asli sebelum proses hash dilakukan
	// berfungsi untuk mencegah serangan menggunakan metode pencocokan data-data yang hasil hash-nya adalah sama (dictionary attack)
	// hal tsb karena hash adalah enkripsi satu arah dengan lebar data yang sudah pasti, jadi sangat mungkin hasil hash pada data adalah sama
	text = "santorini adalah keren"
	fmt.Printf("Ori : %s\n\n", text)

	hashed1, _ := hashWithSalt(text)
	fmt.Println(hashed1)
	fmt.Println()

	hashed2, _ := hashWithSalt(text)
	fmt.Println(hashed2)
	fmt.Println()

	hashed3, _ := hashWithSalt(text)
	fmt.Println(hashed3)
	fmt.Println()

	// praktek penggunaan salt ini biasanya data hasil hash password akan disimpan pada db beserta saltnya pada kolom yang berbeda
	// perlu menyimpan dua nilai karena untuk pencocokan password setiap user melakukan login
	// user | password | salt
}

func hashWithSalt(text string) (string, string) {
	salt := fmt.Sprintf("%d", time.Now().UnixNano())
	saltedText := fmt.Sprintf("text: '%s', salt: %s", text, salt)
	fmt.Println(saltedText)
	sha := sha1.New()
	sha.Write([]byte(saltedText))
	encrypted := sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}
