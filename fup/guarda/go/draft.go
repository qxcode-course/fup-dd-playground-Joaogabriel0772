package main
import "fmt"
func main() {
    var w,l,a int
    fmt.Scan(&w, &l, &a)
    wifi := w == 1
    login := l == 1
    admin := a == 1

    if !wifi {
        fmt.Println("you must connect to wifi")
    } else if !login {
        fmt.Println("you need to login first")
    } else if !admin {
        fmt.Println("you must to login as admin")
    } else {
        fmt.Println("done")
    }
}