// * time duration
// mecari suatu durasi pada program (runtime), jadi kita bisa melihat delta/selisih dari dua objek time.Time
// contoh penerapannya untuk benchmarking durasi

package main

import (
	"fmt"
	"time"
)

func main() {
	// * time duration
	// bakal menghitung waktu yang berlalu sejak suatu objek time
	start := time.Now()
	time.Sleep(time.Second * 5)
	duration := time.Since(start) // menghitung waktu sejak objek start

	// melihat durasi-durasinya
	// ada banyak method untuk melihatnya seperti:
	// Nanosecond, Millisecond, Second, Minute, Hour
	fmt.Println("waktu yang berlalu dalam detik:", duration.Seconds())
	fmt.Println("waktu yang berlalu dalam menit:", duration.Minutes())
	fmt.Println("waktu yang berlalu dalam jam:", duration.Hours())
	fmt.Println("=================")

	// * perhitungan durasi antar dua objek time
	// menghitung selisih dari dua objek time dengan method .sub
	t1 := time.Now()
	time.Sleep(time.Second * 5)
	t2 := time.Now()
	delta := t2.Sub(t1)
	fmt.Println("waktu berlalu lagi dalam detik:", delta)

	// * konversi angka ke objek time.Duration
	n := 5
	duration1 := time.Duration(n) * time.Second
	fmt.Println(duration1)
	fmt.Println(duration1.Seconds())
}
