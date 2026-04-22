package main

import f "fmt"

func main() {
    array := []float64{4.05, 98.4314, 3.14753, 2.25, 56.1237}
    f.Print(soma(array))
}
func soma(x []float64) float64 {
    if len(x) == 0 {
        return 0
    }
    return x[0] + soma(x[1:])
}
