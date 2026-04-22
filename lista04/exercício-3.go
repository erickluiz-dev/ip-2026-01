package main

import f "fmt"

func main() {
    array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    f.Print(inverter(array))
}
func inverter(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    primeiro := 0
    ultimo := len(arr) - 1
    arr[primeiro], arr[ultimo] = arr[ultimo], arr[primeiro]

    inverter(arr[1:ultimo])

    return arr
}
