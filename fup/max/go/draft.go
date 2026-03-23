package main

import "fmt"
func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}
func main() {
    var a, b int
    fmt.Scanln(&a)
    fmt.Scanln(&b)

    resultado := max(a, b)
    fmt.Println(resultado)
}