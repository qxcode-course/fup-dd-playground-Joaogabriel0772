package main
import "fmt"
func main() {
   var n1,n2 int
   var op string
   fmt.Scan(&n1,&n2,&op)
   if op == "+"{
    fmt.Println(n1+n2)
     }else if op == "-"{
    fmt.Println(n1-n2)
   } else if op == "*"{
    fmt.Println(n1*n2)
   } else if op == "/"{
    fmt.Println(n1/n2)
   }
   

}
