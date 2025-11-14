package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/YevgeniyGertsen/module10practise/shop"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== Практическая работа №1: Интернет-магазин ===")
		fmt.Println("1) Приветствие пользователя")
		fmt.Println("2) Добавление товара в корзину")
		fmt.Println("3) Расчёт стоимости заказа")
		fmt.Println("4) Определение статуса клиента")
		fmt.Println("5) Меню магазина (switch)")
		fmt.Println("6) Подсчёт количества товаров (for)")
		fmt.Println("7) Стоимость доставки (функция)")
		fmt.Println("0) Выход")
		fmt.Print("Выберите пункт меню: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, _ := strconv.Atoi(input)

		switch choice {
		case 0:
			fmt.Println("Выход.")
			return
		case 1:
			shop.Greeting(reader)
		case 2:
			shop.AddToCart(reader)
		case 3:
			shop.OrderCost(reader)
		case 4:
			shop.CustomerStatus(reader)
		case 5:
			shop.ShopMenu(reader)
		case 6:
			countItems(reader) // внутренняя функция
		case 7:
			shop.DeliveryTask(reader)
		default:
			fmt.Println("Такого пункта нет!")
		}
	}
}

// Задание 6 — работает прямо здесь, чтобы показать работу циклов
func countItems(reader *bufio.Reader) {
	fmt.Println("\n--- Подсчёт товаров в корзине ---")

	fmt.Print("Введите количество разных товаров: ")
	tStr, _ := reader.ReadString('\n')
	tStr = strings.TrimSpace(tStr)
	types, _ := strconv.Atoi(tStr)

	total := 0

	for i := 1; i <= types; i++ {
		fmt.Printf("Введите количество товара №%d: ", i)
		qStr, _ := reader.ReadString('\n')
		qStr = strings.TrimSpace(qStr)
		qty, _ := strconv.Atoi(qStr)

		total += qty
	}

	fmt.Printf("Всего товаров в корзине: %d шт.\n", total)
}
