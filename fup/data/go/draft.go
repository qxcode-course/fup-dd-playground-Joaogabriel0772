package main
import "fmt"
func main() {
   var h,m,d,mes,ano int
   fmt.Scan(&h)
   fmt.Scan(&m)
   fmt.Scan(&d)
   fmt.Scan(&mes)
   fmt.Scan(&ano)
    ano=ano % 100
    fmt.Printf("%02d:%02d %02d/%02d/%02d\n", h, m, d, mes, ano)
   }
