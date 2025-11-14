package shop

import (
	"bufio"
	"fmt"
	"strings"
)

func Greeting(reader *bufio.Reader) {
	fmt.Println("\n--- Приветствие пользователя ---")
	fmt.Print("Введите ваше имя: ")

	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Добро пожаловать в интернет-магазин, %s!\n", name)
}
