package main

import "fmt"

/*
Многострочный комментарий
*/
//Однострочный коментарий
func Greet() {
	fmt.Println("Привет, Друг!!!")
}

func PersonGreet(name string) {
	fmt.Println("Привет " + name)
}

func FioGreet(name, surname string) {
	fmt.Printf("Привет, %s %s\n", name, surname)
}

func Sum(a, b int) int {
	sum := a + b
	return sum
}

func SumAndMultiplay(a, b int) (int, int) {
	sum := a + b
	return sum, a * b
}

func NameSumAndMulty(a, b int) (sum int64, mult int64) {
	sum = int64(a + b)
	mult = int64(a * b)
	return
}

// Точка входа
func main() {
	Greet()
	PersonGreet("Вячеслав")
	FioGreet("Вячеслав", "Ерёмин")
	a1 := 1
	b1 := 4
	a2, b2 := 6, 3
	sum := Sum(a1, b1)
	fmt.Printf("Результат функции Sum = %v\n", sum)
	fmt.Printf("Результат функции Sum = %v\n", Sum(a2, b2))
	summa, mult := SumAndMultiplay(3, 4)
	_ = summa

	fmt.Println(summa)
	fmt.Println(mult)
	_, mutyplay64 := NameSumAndMulty(2, 6)
	fmt.Println(mutyplay64)
	sum64, _ := NameSumAndMulty(4, 9)
	fmt.Println(sum64)

}
