package main

import "fmt"

func recursiveSum(list []int) int{
   if(len(list) == 0){
    return 0}
    return list[0]+ recursiveSum(list[1:])

}

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
    totalSumByRecursion := recursiveSum(list)
    fmt.Printf("Sum of the list of numbers: %d\n", totalSum)
    fmt.Printf("Sum of the list of numbers: %d\n", totalSumByRecursion)
}
