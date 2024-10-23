package main

import "fmt"

func numInList(list []int, num int) bool {
    for _, val := range list {
        if val == num {
            return true
        }
    }
    return false
}

func main() {
    list := []int{1, 2, 3, 4, 5, 6}
    num := 7
    found := numInList(list, num) 
    fmt.Println(found)
}
