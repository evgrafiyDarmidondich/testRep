package utils

func CreateDivider(divider int, fl string) func(x int) any {
	dividerFunc := func(x int) any {
		if fl == "/" {
			return x / divider
		} else if fl == "*" {
			return x * divider
		} else if fl == "+" {
			return x + divider
		} else if fl == "-" {
			return x - divider
		} else {
			return "Неверный оператор"
		}
	}
	return dividerFunc
}

func IsChildern(num int) bool {
	return num < 18
}
