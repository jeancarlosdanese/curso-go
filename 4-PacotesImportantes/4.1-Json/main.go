// Trabalhando com JSON

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int     `json:"numero"`
	Saldo  float64 `json:"saldo"`
}

func main() {
	conta := Conta{123, 100.50}

	res, err := json.Marshal(conta)
	if err != nil {
		fmt.Printf("Erro ao converter a conta em JSON: %v\n", err)
	}

	fmt.Printf("A conta em JSON: %s\n", res)

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		fmt.Printf("Erro ao converter a conta em JSON: %v\n", err)
	}

	jsonPuro := []byte(`{"numero": 124, "saldo": 101.50}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		fmt.Printf("Erro ao converter o JSON em conta: %v\n", err)
	}

	fmt.Printf("A conta convertida: numero %d, saldo %.2f\n", contaX.Numero, contaX.Saldo)
}
