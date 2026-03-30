package main
import "fmt"
func floor(num float64) int {
    inteiro := int(num)
    if num < 0 && float64(inteiro) != num {
        return inteiro - 1
    }
    return inteiro
}
func ceil(num float64) int {
    inteiro := int(num)
    if num > 0 && float64(inteiro) != num {
        return inteiro + 1
    }
    return inteiro
}
func round(num float64) int {
    inteiro := int(num)
    frac := num - float64(inteiro)
    if frac >= 0.5 {
        return inteiro + 1
    } else if frac <= -0.5 {
        return inteiro - 1
    }
    return inteiro
}
func main() {
    var num float64
    var op string

    fmt.Scan(&op)
    fmt.Scan(&num)

    switch op {
    case "f":
        fmt.Println(floor(num))
    case "c":
        fmt.Println(ceil(num))
    case "r":
        fmt.Println(round(num))
    default:
        fmt.Println("Operação inválida")
    }
}