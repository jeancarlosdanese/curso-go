package matematica

func Soma[T int | float64](a, b T) T {
	return a + b
}

var A = 45

type Carro struct {
	Marca string
}

func (c Carro) Andar() string {
	return "O carro da marca " + c.Marca + " está andando"
}
