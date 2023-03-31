package main

import "fmt"

func main() {
	slice2 := []string{"a", "b", "c", "d"}
	slice1 := []int{1, 2, 3, 4}

	newInts := reverse(slice1)
	newString := reverse(slice2)

	fmt.Println(newInts)
	fmt.Println(newString)

}

type constraintCustom interface {
	int | string
}

// para nao ter de criar duas funcoes para fazer do tipo int e string, o generic metodo e esse [T int | T string]
func reverse[T constraintCustom](slice []T) []T {
	newInts := make([]T, len(slice))

	newIntsLen := len(slice) - 1
	for i := 0; i < len(slice); i++ {
		newInts[newIntsLen] = slice[i]
		newIntsLen--
	}
	return newInts
}
