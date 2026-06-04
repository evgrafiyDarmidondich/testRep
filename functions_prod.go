package main

import (
	"fmt"
	"testrep/utils"
)

func Calculate(x, y int, act func(x, y int) int) int {
	return act(x, y)
}

// Точка входа
func main() {
	var mult func(x, y int) int // объявил переменную mult типа func
	fmt.Printf("Тип переменной mult - %T\n", mult)
	mult = func(x, y int) int { return x * y } // присвоил знвчение переменной mult
	fmt.Printf("mult type - %T. mult value = %v\n", mult, mult)
	fmt.Println(mult(3, 5))

	sum := func(x, y int) int { return x + y }
	fmt.Println(sum(2, 7))

	fmt.Println(Calculate(2, 3, sum))
	fmt.Println(Calculate(3, 7, func(x, y int) int { return x * y }))
	diveBy2 := utils.CreateDivider(2, "/") // Создали функцию при помощи функции
	diveBy10 := utils.CreateDivider(10, "d")

	fmt.Printf("100/2 = %v\n", diveBy2(100)) // в созданую функцию передали число 2 в качестве аргумента
	fmt.Printf("100*10 = %v\n", diveBy10(100))

	fmt.Println(utils.IsChildern(17))

	if isChild := utils.IsChildern(18); isChild == true {
		fmt.Println("Это ребенок")
	}

	utils.TestPakcage()
}
