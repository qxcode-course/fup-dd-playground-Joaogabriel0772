package main
import "fmt"
func main() {
  	var idade, quantidade int
	fmt.Scan(&idade)
	fmt.Scan(&quantidade)

	for i := 0; i < quantidade; i++ {
		fmt.Println(idade)
		idade += 2
	}
}
