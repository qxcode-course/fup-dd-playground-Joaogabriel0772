package main
import "fmt"
func main() {
    var A, B int
    fmt.Scan(&A)
    fmt.Scan(&B)
    for i := A; i < B; i++ {
        fmt.Println(i)
    }
}