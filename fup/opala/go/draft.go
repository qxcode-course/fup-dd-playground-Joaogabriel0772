package main
import "fmt"
func main() {
    var velocidade, tempoMin, consumo float64
    fmt.Scan(&velocidade)
    fmt.Scan(&tempoMin)
    fmt.Scan(&consumo)
    tempoHoras := tempoMin / 60.0
    distancia := velocidade * tempoHoras
    desempenho := distancia / consumo
    fmt.Printf("%.2f\n", desempenho)
}