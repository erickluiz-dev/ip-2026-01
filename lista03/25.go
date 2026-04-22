package main

import (
    f "fmt"
    m "math"
)

func main() {
    S := 0.0
    x := 15.0

    for i := 0.0; i <= 14; i++ {
        if int(x)%2 == 0 {
            S -= m.Pow(2, i) / m.Pow(x, 2)
        } else {
            S += m.Pow(2, i) / m.Pow(x, 2)
        }
        x--
    }

    f.Printf("S = %f", S)
}