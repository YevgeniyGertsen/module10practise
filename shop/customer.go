package shop

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func CustomerStatus(reader *bufio.Reader) {
	fmt.Println("\n--- Определение статуса клиента ---")

	fmt.Print("Сколько покупок вы сделали за год? ")
	bStr, _ := reader.ReadString('\n')
	bStr = strings.TrimSpace(bStr)
	buys, _ := strconv.Atoi(bStr)

	if buys <= 5 {
		fmt.Println("Ваш статус: Обычный покупатель")
	} else if buys <= 20 {
		fmt.Println("Ваш статус: Постоянный покупатель")
	} else {
		fmt.Println("Ваш статус: VIP клиент")
	}
}
