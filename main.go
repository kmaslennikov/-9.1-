package main

import (
	"fmt"
	"math"
)

func main() {
	var counterUint8, counterUint16 int
	var tempUint8 uint8
	var tempuint16 uint16
	for i := 0; i < math.MaxUint32; i++ {
		if tempUint8 == math.MaxUint8 {
			counterUint8++
		}
		if tempuint16 == math.MaxUint16 {
			counterUint16++
		}
		tempUint8++
		tempuint16++
	}
	fmt.Printf("%d приходится переполнений чисел типа uint8 в диапазоне от 0 до uint32\n", counterUint8)

	fmt.Printf("%d приходится переполнений чисел типа uint16 в диапазоне от 0 до uint32\n", counterUint16)
}
