// * Arguments & Flag
// arguments : data opsional yang disisipkan ketika eksekusi program
// flag : ekstensi dari argument

package main

import (
	"flag"
	"fmt"
)

func main() {
	// * Arguments

	// // mengambil seluruh args, berisi path file exec juga
	// argsRaw := os.Args
	// fmt.Println(argsRaw)

	// // hanya mengambl slice args saja
	// args := argsRaw[1:]
	// fmt.Printf("%#v\n", args)

	// go run *.go aku sugab "donat enak"
	// untuk args yang memiliki space harus menggunakan tanda  petik ("")

	// * Flag
	// -> berguna seperti arguments, parameterize eksekusi program
	// penulisannya key-value

	// flag bertipe string, (key, default value, keterangan)
	// nilai balik fungsi ini adalah string pointer, jadi perlu di deference dulu klo mau dipake
	name := flag.String("name", "anonymous", "type your name")
	age := flag.Int64("age", 0, "type your age")

	flag.Parse()
	fmt.Println("name:", *name)
	fmt.Println("age:", *age)

	// go run *.go -name="bagus adi" -age=25
}
