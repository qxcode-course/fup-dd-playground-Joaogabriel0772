package main
import "fmt"
func main() {
    var P, D1, D2 int
    fmt.Scan(&P)
    fmt.Scan(&D1)
    fmt.Scan(&D2)
    soma := D1 + D2
    if soma%2 == 0 {
        fmt.Println(P)
    } else {
        fmt.Println(1 - P)
    }
}