import f "fmt"

func main() {
    x, n := 0, 0
    f.Print("Digite X ")
    f.Scan(&x)
    f.Print("Digite N ")
    f.Scan(&n)
    f.Println(pow(x, n))
}
func pow(base, expoente int) int {
    if expoente == 0 {
        return 1
    }
    return base * pow(base, expoente-1)
}
