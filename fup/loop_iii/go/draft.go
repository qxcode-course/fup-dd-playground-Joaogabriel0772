package main
import "fmt"
func main() {
  var a,b int
  fmt.Scan(&a,&b)
  fmt.Print("[ ")
    for i := a; i > b; i -- {
        if i == b-1 {
            fmt.Print(i)
        } else {
            fmt.Print(i, " ")
        }
    }
    fmt.Println("]")
}
