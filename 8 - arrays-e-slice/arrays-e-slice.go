package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("arrays e slice")

	var array [5]string
	array[0] = "Posição 1"

	fmt.Println(array)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	slice := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array))
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:4]
	fmt.Println(slice2)

	// ARRAYS INTERNOS

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
