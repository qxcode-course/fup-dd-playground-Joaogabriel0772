package main
import "fmt"

func main() {
    var n int
    fmt.Scanln(&n)

    if n >= 0 {
        fmt.Println("SIM")
    } else {
        fmt.Println("NAO")
    }
}
