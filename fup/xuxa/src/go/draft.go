package main
import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if len(text) > 0 && text[len(text)-1] == '\n' {
		text = text[:len(text)-1]
	}
	runes := []rune(text)
	result := ""
	for i := len(runes) - 1; i >= 0; i-- {
		result += string(runes[i])
	}
	fmt.Println(result) 
}