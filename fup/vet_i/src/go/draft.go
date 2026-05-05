package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)
    vetor := make([]int, N)
    for i := 0; i < N; i++ {
        fmt.Scan(&vetor[i])
    }
    if N== 0{
        fmt.Print("\n")
    }
    for i := 0; i < N; i++ {
        fmt.Println(vetor[i])
    }
}