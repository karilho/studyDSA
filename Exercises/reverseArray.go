package main

import (
	"fmt"
)

func reverseArray(arr []int) {
    for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }

    fmt.Println(arr)
}

func reverseArraySameArray(arr []int) []int {
    // Aponta prp final do array para o início
    final := len(arr) - 1

	fmt.Println(arr)

    for inicio, _ := range arr[:len(arr)/2] { // Itera sobre a metade do array
		fmt.Println("valor do inicio: ", inicio, " valor do final: ", final)
        arr[inicio], arr[final] = arr[final], arr[inicio]
        final-- // Decrementa j para a próxima troca
		fmt.Println(arr)
    }

    fmt.Println(arr)
	return arr
}

func reverseArrayUsingMap(arr []int) {
	reversedMap := make(map[int]int, len(arr))

	for i, j := 0, len(arr)-1; i <= j; i, j = i+1, j-1 {
		reversedMap[i] = arr[j]
		reversedMap[j] = arr[i]
	}

	fmt.Println(reversedMap)
}

func reverseArrayWithForRange(arr []int) []int {
    reversed := make([]int, len(arr))
    
    for index, value := range arr {
		// Calcula o índice correspondente no array revertido
		// como funciona? 
		// -> o len(arr) é o tamanho do array
		// -> o -1 é para apontar para o último elemento inicialmente
		// -> o -index é para cada iteração, ele descer -1 para cada elemento do array, assim, trocando o índice
		reversed[len(arr)-1-index] = value
    }
    
    return reversed
}

func main() {
	var arr []int
	arr = []int{4, 2, 6, 3, 1, 7}
	reversed := reverseArraySameArray(arr)
	fmt.Println("Reversed array: ", reversed)
}
