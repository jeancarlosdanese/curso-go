// Type assertion

package main

import "fmt"

func main() {
	var minhaVar interface{} = "Jean Carlos"

	println(minhaVar.(string))

	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)

	// panic: interface conversion: interface {} is string, not int
	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 é %v e o resultado de ok2 é %v\n", res2)
}
