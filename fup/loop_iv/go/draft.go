package main
import "fmt"
func main() {
    var A, B int
    fmt.Scan(&A, &B)
    fmt.Print("[ ")
    step := 1
    if A > B {
        step = -1
    }
    for i := A; i != B; i += step {
        if i+step == B {
            fmt.Print(i)
        } else {
            fmt.Print(i, " ")
        }
    }
    fmt.Println(" ]")
}