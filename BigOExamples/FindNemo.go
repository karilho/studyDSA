package main

import (
    "fmt"
    "time"
)

func main() {
    fish := []string{"dory", "bruce", "marlin", "nemo", "gill", "bloat", "nigel", "squirt", "darla", "hank"}

    for i := 0; i < 100000; i++ {
        fish = append(fish, "nemo")
    }

    findNemo(fish)
}

func findNemo(array []string) {
    timeStart := time.Now()
    for i := 0; i < len(array); i++ {
        if array[i] == "nemo" {
            fmt.Println("Found Nemo!")
        }
    }
    timeEnd := time.Now()

    fmt.Println("Call to find Nemo took", timeEnd.Sub(timeStart).Milliseconds(), "miliseconds")
}
