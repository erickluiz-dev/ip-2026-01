package main

import f "fmt"

func main() {
    n, m, res := 0, 1000.0, 0.0
    f.Print("Digite o número de termos ")
    f.Scan(&n)

    if n <= 0 {
        return
    }

    for i := 1; i <= n; i++ {
        if (int(m) % 2) == 0 {
            res += m / float64(i)
        } else {
            res -= m / float64(i)
        }
        m -= 3
    }
    f.Printf("A soma é = %f ", res)
}