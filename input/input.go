package input

import (
	"bufio"
	"fmt"
	"os"
)

func GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение, где строки должны быть в кавычках: ")
	str, _ := reader.ReadString('\n')

	return str
}
