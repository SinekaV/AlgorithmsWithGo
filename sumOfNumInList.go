package main

import "fmt"
 
func sum(list []int) int {
    total := 0
    for _, val := range list {
        total += val
    }
    return total
}

func main() {
    list := []int{1, 1, 1, 1, 1, 1}
    totalSum := sum(list)
    fmt.Printf("Sum of the list of numbers: %d\n", totalSum)
}
