/* -------------------------------------------------------------------------- */
/*                            // * channel - select                           */
/* -------------------------------------------------------------------------- */
// dengan adanya channel, goroutine dapat di-manage dengan mudah
// namun fungsi utama channel bukan untuk kontrol, melainkan sharing data antar goroutine

// ada kalanya ktia tidak hanya butuh satu channel saja untuk komunikasi data
// maka dari itu ada sintaks `select` yang akan mempermudah kontrol komunikasi data lewat banyak channel
// penggunaan `select` ini mirip seperti `switch` untuk seleksi kondisi

package main

import "fmt"

// function for find max number in numbers slice
func getMax(numbers []int, channel chan int) {
	maxOfNumber := 0 // var for max number

	for i, number := range numbers {
		if i == 0 {
			maxOfNumber = number
		}

		if maxOfNumber < number {
			maxOfNumber = number
		}
	}

	channel <- maxOfNumber // send data to channel
}

// function for calculate average number in numbers slice
func getAverage(numbers []int, channel chan float64) {
	sumOfNumber := 0 // temp var for sum of number

	for _, number := range numbers {
		sumOfNumber += number
	}

	averageOfNumber := float64(sumOfNumber) / float64(len(numbers))

	channel <- averageOfNumber // send data to channel
}

func main() {
	// initiate numbers slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// make channel and execute getMax async
	ch1 := make(chan int)
	go getMax(numbers, ch1)

	// make channel and execute getAverage async
	ch2 := make(chan float64)
	go getAverage(numbers, ch2)

	fmt.Println("numbers:", numbers)

	// async get data from two channel async using select
	for i := 0; i < 2; i++ {
		select {
		case maxOfNumber := <-ch1:
			fmt.Println("max of number:", maxOfNumber)
		case averageOfNumber := <-ch2:
			fmt.Println("average of number:", averageOfNumber)
		}
	}
}
