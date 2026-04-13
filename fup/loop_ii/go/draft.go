package main
import "fmt"
func main() {
 var A,B int
 fmt.Scan(&A, &B)
 fmt.Print("[ ")
    for i := A; i < B; i++ {
        if i == B-1 {
            fmt.Print(i)
        } else {
            fmt.Print(i, " ")
        }
    }
    fmt.Println(" ]")
}
