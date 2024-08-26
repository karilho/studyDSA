package main

import (
	"fmt"
)


func gradingStudents(grades []int32) []int32 {
	for k, v := range grades {
		if v % 5  >= 3 && v >= 38 {
			grades[k] = v + (5 - (v % 5))
		}

	}
	return grades
}

func gradingStudents2(grades []int32) []int32 {
    for i, grade := range grades {
        if grade >= 38 && grade % 5 >= 3 {
            grades[i] = grade + (5 - (grade % 5))
        }
    }
    return grades
}

func main3() {
	gradingStudents([]int32{4,73,67,38,33, 84})


	
	fmt.Println("Done")

}