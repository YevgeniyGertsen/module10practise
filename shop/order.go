package shop

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func OrderCost(reader *bufio.Reader) {
	fmt.Println("\n--- Расчёт стоимости заказа ---")

	fmt.Print("Введите цену товара: ")
	pStr, _ := reader.ReadString('\n')
	pStr = strings.TrimSpace(pStr)
	price, _ := strconv.Atoi(pStr)

	fmt.Print("Введите количество: ")
	qStr, _ := reader.ReadString('\n')
	qStr = strings.TrimSpace(qStr)
	qty, _ := strconv.Atoi(qStr)

	total := price * qty

	fmt.Printf("Стоимость заказа: %d₸\n", total)
}
