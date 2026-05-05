package main
import "fmt"
func main() {
  var a,b int
  fmt.Scan(&a,&b)
  fmt.Print("[ ")
   for {
    if a<=b{
        if a%2 != 0 {
            fmt.Printf("%d ", a)
        } 
        a++
    } else{
        break
    }
   } 
   
   fmt.Println("]")
}
