package main

import (
    "fmt"
    "sort"
)

// Given 2 arrys,create a function to user knowtrue / false  - >whetherthese two arrys containsa any common items
// For Example:
// const array1 = ['a', 'b', 'c', 'x']
// const array2 = ['z', 'y', 'i']
// should return false

// const array1 = ['a', 'b', 'c', 'x']
// const array2 = ['z', 'y', 'x']
// should return true

// 2 parameters - arrays - no size limit
// return true or false

//my initial solution O(N²)
func containsCommons(arr1, arr2 []string) bool {
    for _, va := range arr1 {
        for _, v := range arr2 {
            if va == v {
                return true
            }
        }
    }
    return false
}

//O(M+N)
func containsCommons2(arr1, arr2 []string) bool {
    mapArr1 := make(map[string]bool)
    // Puts all value of keys of map = true
    for _, item := range arr1 {
        mapArr1[item] = true
    }

    // Check if the value of the key is true, value by value
    for _, item := range arr2 {
        if mapArr1[item] {
            return true
        }
    }

    return false
}

//s3 - O(n*log(n))
func containsCommons3(arr1, arr2 []string) bool {
    sort.Strings(arr1)
    sort.Strings(arr2)

    //double condition. If the first array is bigger than the second, the loop will stop, because they are sorted, so, the first array will never be equal to the second array
    for i, j := 0, 0; i < len(arr1) && j < len(arr2); {
        if arr1[i] == arr2[j] {
            return true
        } else if arr1[i] < arr2[j] {
            i++
        } else {
            j++
        }
    }

    return false
}

func main1() {
    var arr1 = []string{"xab", "base", "ca", "ty", "bebeb"}
    var arr2 = []string{"cac", "tega", "manteg", "f"}

    // fmt.Println(containsCommons(arr1, arr2))
    //fmt.Println(containsCommons2(arr1, arr2))
    fmt.Println(containsCommons3(arr1, arr2))

}
