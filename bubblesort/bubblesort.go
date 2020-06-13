package main

import (
	"fmt"
)

// BubbleSort sorts a slice using bubble sort algorithm
func BubbleSort(numbersSlice []int) {

	for i := 0; i < len(numbersSlice); i++ {
		for j := len(numbersSlice) - 1; j >= i+1; j-- {
			if numbersSlice[j] < numbersSlice[j-1] {
				Swap(numbersSlice, j)

			}
		}
	}

}

// Swap sorts values
func Swap(tosort []int, index int) {

	tosort[index], tosort[index-1] = tosort[index-1], tosort[index]

}

func main() {
	values := make([]int, 10)
	fmt.Println("Enter upto 10 numbers")
	for i := 0; i < len(values); i++ {

		_, err := fmt.Scanf("%d\n", &values[i])

		if err != nil {
			fmt.Println(err)
		}

	}

	// sort the slice
	BubbleSort(values)

	// print sorted slice
	for _, value := range values {
		fmt.Printf("%v ", value)
	}

}
