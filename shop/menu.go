package shop

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ShopMenu(reader *bufio.Reader) {
	fmt.Println("\n--- Меню магазина ---")
	fmt.Println("1 — Посмотреть каталог")
	fmt.Println("2 — Посмотреть корзину")
	fmt.Println("3 — Оформить заказ")
	fmt.Print("Выберите действие: ")

	cStr, _ := reader.ReadString('\n')
	cStr = strings.TrimSpace(cStr)
	choice, _ := strconv.Atoi(cStr)

	switch choice {
	case 1:
		fmt.Println("Каталог: товары временно недоступны.")
	case 2:
		fmt.Println("Корзина пуста.")
	case 3:
		fmt.Println("Ваш заказ принят и обрабатывается.")
	default:
		fmt.Println("Неизвестная команда.")
	}
}
