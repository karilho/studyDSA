package main

import (
    "fmt"
    "strings"
)

var array = []string{"a", "b", "c ", "d", "e"}

func logPairs(array []string) {
    for _, v := range array {
        for _, n := range array {
            if strings.Contains(v, n) {
                fmt.Println(v, n)
            }
        }
    }
}

func main() {
    logPairs(array)
}
