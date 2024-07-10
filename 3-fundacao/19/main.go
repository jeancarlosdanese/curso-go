// Interfaces vazias

package main

func main() {
	// A interface vazia é um tipo de interface que não possui métodos.
	// Isso significa que qualquer tipo de dado pode ser atribuído a uma variável do tipo interface{}.
	// Isso é útil quando você deseja escrever funções que aceitam qualquer tipo de dado.
	// Por exemplo, a função println aceita qualquer tipo de dado.
	// Isso é possível porque a função println aceita um número variável de argumentos do tipo interface{}.
	// Aqui está um exemplo de como você pode usar uma interface vazia:
	var a interface{}
	var b interface{}

	a = 1
	b = "Hello, World!"

	showType(a)
	showType(b)
}

func showType(t interface{}) {
	// Você pode usar uma declaração de switch para verificar o tipo de dado que foi atribuído a uma variável do tipo interface{}.
	// Aqui está um exemplo de como você pode fazer isso:
	switch t.(type) {
	case int:
		println("O tipo de t é int")
	case string:
		println("O tipo de t é string")
	default:
		println("O tipo de t é desconhecido")
	}
}
