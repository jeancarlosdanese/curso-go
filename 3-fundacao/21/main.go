// Generics

package main

type MyNumber int

type Number interface {
	~int | float64
}

func SomaInteiros(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloats(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Comparar[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	m := map[string]int{"Jean": 1500, "Liane": 2000, "Bernardo": 3000}
	m2 := map[string]float64{"Jean": 1500.50, "Liane": 2000.60, "Bernardo": 3000.90}

	println(SomaInteiros(m))
	println(SomaFloats(m2))

	m3 := map[string]int{"Jean": 1500, "Liane": 2000, "Bernardo": 3000}
	m4 := map[string]float64{"Jean": 1500.50, "Liane": 2000.60, "Bernardo": 3000.90}

	m5 := map[string]MyNumber{"Jean": 1500, "Liane": 2000, "Bernardo": 3000}

	println(Soma(m3))
	println(Soma(m4))

	println(Soma(m5))
}
