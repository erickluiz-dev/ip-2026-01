package main

import f "fmt"

func main() {
    s, n1 := 0.0, 38.0

    for i := 1; i <= 37; i++ {
        s += (n1 * (n1 - 1)) / float64(i)
        n1--
    }
    f.Printf("A soma é = %.0f ", s)
}