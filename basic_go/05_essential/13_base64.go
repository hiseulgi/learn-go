// * Encode - Decode Base64
// fungsi-fungsi Base64 pada Golang
// Data yang akan di-encode harus bertipe []byte

package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// * EncodeToString() dan DecodeString()

	// data := "bagus sugab"

	// // encoding data
	// encodedString := base64.StdEncoding.EncodeToString([]byte(data))
	// fmt.Println(encodedString)

	// // decoding data
	// decodedByte, _ := base64.StdEncoding.DecodeString(encodedString)
	// decodedString := string(decodedByte)
	// fmt.Println(decodedString)

	/* -------------------------------------------------------------------------- */

	// * Encode() dan Decode()
	// beda dengan cara di atas
	// dimana pada cara ini kita perlu membuat penampung bertipe []byte dengan lebarnya juga

	// data := "sasugab"

	// // encode
	// encoded := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	// base64.StdEncoding.Encode(encoded, []byte(data))
	// encodedString := string(encoded)
	// fmt.Println(encodedString)

	// // decode
	// decoded := make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	// _, err := base64.StdEncoding.Decode(decoded, encoded)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// decodedString := string(decoded)
	// fmt.Println(decodedString)

	/* -------------------------------------------------------------------------- */

	// * URL Encode & Decode

	data := "https://google.com/"

	// encode
	encodedString := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedString)

	// decode
	decodedByte, _ := base64.URLEncoding.DecodeString(encodedString)
	decodedString := string(decodedByte)
	fmt.Println(decodedString)
}
