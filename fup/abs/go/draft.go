package main
import "fmt"
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    var a, b int

    fmt.Scanln(&a)
    fmt.Scanln(&b)

    diferenca := a - b
    resultado := abs(diferenca)

    fmt.Println(resultado)
}