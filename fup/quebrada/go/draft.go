package main
import "fmt"
func main() {
    var n1, n2 int
    fmt.Scanln(&n1)
    fmt.Scanln(&n2)
    fmt.Println(n1 / n2)
    fmt.Println(n1 % n2)
    fmt.Printf("%.2f\n", float64(n1)/float64(n2))
}