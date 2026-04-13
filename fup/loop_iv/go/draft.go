package main
import "fmt"
func main() {
    var A,B int
 fmt.Scan(&A, &B)
 fmt.Print("[ ")
    for i := A,B i++ {
        if i == B-1 {
            fmt.Print(i)
        } else {
            fmt.Print(i, " ")
        }
    }
    fmt.Println(" ]")
}
