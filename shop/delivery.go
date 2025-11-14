package shop

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// функция расчёта стоимости доставки
func DeliveryCost(distance float64) int {
	if distance <= 5 {
		return 500
	}
	if distance <= 15 {
		return 1200
	}
	return 2000
}

func DeliveryTask(reader *bufio.Reader) {
	fmt.Println("\n--- Стоимость доставки ---")

	fmt.Print("Введите расстояние до пункта доставки (км): ")
	dStr, _ := reader.ReadString('\n')
	dStr = strings.TrimSpace(dStr)
	distance, _ := strconv.ParseFloat(dStr, 64)

	cost := DeliveryCost(distance)

	fmt.Printf("Расстояние: %.1f км → Стоимость доставки: %d₸\n", distance, cost)
}
