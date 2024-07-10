// Maps entendendo maps e utilizando

package main

import "fmt"

func main() {
	salarios := map[string]float64{"Jean": 1000.0, "Carlos": 2000.0}
	fmt.Println(salarios)

	delete(salarios, "Jean")
	fmt.Println(salarios)

	salarios["Carlos"] = 3000.0

	salarios["Pedro"] = 1200.0

	fmt.Println(salarios["Carlos"])
	fmt.Println(salarios["Pedro"])

	sal := make(map[string]float64)
	fmt.Println(sal)

	sal["João"] = 5000.0
	sal["Maria"] = 6000.0
	fmt.Println(sal)

	sal1 := map[string]float64{}
	fmt.Println(sal1)

	sal1["João"] = 5000.0
	sal1["Maria"] = 6000.0
	fmt.Println(sal1)

	sal1["Jean"] = 5500.0

	for nome, salario := range sal1 {
		fmt.Printf("O salário de %s é %.2f\n", nome, salario)
	}

	_, ok := sal1["Jean"]
	if ok {
		fmt.Println("Jean está no map")
	} else {
		fmt.Println("Jean não está no map")
	}
}
