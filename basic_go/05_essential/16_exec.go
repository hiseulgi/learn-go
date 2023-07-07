// * Exec
// eksekusi perintah cli lewat kode program

package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	// exec.Command() -> menjalankan command
	// fungsi diatas di-chain dengan method Output() untuk mendapatkan outputnya
	// output yang dihasilkan berbentuk []byte
	output1, _ := exec.Command("ls").Output()
	fmt.Printf("-> ls\n%s\n", string(output1))

	output2, _ := exec.Command("pwd").Output()
	fmt.Printf("-> pwd\n%s\n", string(output2))

	output3, _ := exec.Command("git", "config", "user.name").Output()
	fmt.Printf("-> git config user.name\n%s\n", string(output3))

	/* -------------------------------------------------------------------------- */
	// rekomendasi penggunaan exec
	// misal command tsb tidak ada pada os tersebut, maka kita perlu kondisian dulu untuk deteksi osnya
	var output []byte
	var err any
	if runtime.GOOS == "windows" {
		output, err = exec.Command("cmd", "/C", "git config user.name").Output()
	} else {
		output, err = exec.Command("bash", "-c", "git config user.name").Output()
	}
	fmt.Println(string(output))
	_ = err
}
