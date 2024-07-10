// Condicionais: if, else if, else

package main

func main() {
	a := 10
	b := 20

	compare(a, b)

	b = 10
	compare(a, b)

	a = 20
	compare(a, b)

	c := 30
	if a > b || a > c {
		println("a > b ou a > c")
	} else {
		println("a <= b e a <= c")
	}

	condicao := 1
	switch condicao {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("default")
	}
}

func compare(a, b int) {
	if a > b {
		println("a > b")
	} else if a < b {
		println("a < b")
	} else {
		println("a == b")
	}
}
