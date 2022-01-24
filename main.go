package main

import (
	"fmt"
	"math"
)

func main() {
	var counterUint8, counterUint16 int

	for i := math.MaxUint32; i > 0; i -= math.MaxUint8 {
		counterUint8++
	}
	for i := math.MaxUint32; i > 0; i -= math.MaxUint16 {
		counterUint16++
	}
	fmt.Printf("%d приходится переполнений чисел типа uint8 в диапазоне от 0 до uint32\n", counterUint8)
	fmt.Printf("%d приходится переполнений чисел типа uint8 в диапазоне от 0 до uint32\n", counterUint16)

}
