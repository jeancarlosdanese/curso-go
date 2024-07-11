package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// ApiName define os tipos possíveis de ações em uma entidade.
type ApiName string

const (
	ViaCepName    ApiName = "ViaCEP"
	BrasilApiName ApiName = "BrasilAPI"
)

// Address - struct para armazenar os dados de endereço padronizados
type Address struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Cidade     string `json:"cidade"`
	Uf         string `json:"uf"`
}

// ViaCEP - struct para receber os dados de ViaCEP
type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

// BrasilAPI - struct para receber os dados de BrasilAPI
type BrasilAPI struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func main() {
	// Pega CEP da linha do comando (argumentos)
	for _, cep := range os.Args[1:] {

		// Timeout de 1 segundo
		timeout := 1 * time.Second

		// Contexto com timeout
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// // Criando um canal para cada API
		// ch1 := make(chan string)
		// ch2 := make(chan string)

		// Cria um canal para comunicação entre goroutines
		ch := make(chan string, 2) // Buffer de 2 posições

		go fetchAddress(ctx, "http://viacep.com.br/ws/"+cep+"/json/", ViaCepName, ch)
		go fetchAddress(ctx, "https://brasilapi.com.br/api/cep/v1/"+cep, BrasilApiName, ch)

		select {
		// case result := <-ch1:
		// 	fmt.Println(result)
		// case result := <-ch2:
		// 	fmt.Println(result)
		case result := <-ch:
			fmt.Println(result)
		case <-ctx.Done():
			fmt.Println("Timeout! No response within 1 second.")
		}
	}
}

func fetchAddress(ctx context.Context, url string, apiName ApiName, ch chan string) {
	log.Printf("Buscando endereço de %s: %s\n", apiName, url)
	starProgress := time.Now()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error fetching address:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ch <- fmt.Sprintf("Error from %s: CEP not found or other error", apiName)
		return
	}

	// Decodificando a resposta JSON e armazenando em uma struct
	var addr Address
	// Verifica qual API foi usada para buscar o endereço
	if apiName == ViaCepName {
		// Se for ViaCEP
		var viaCep ViaCEP
		if err := json.NewDecoder(resp.Body).Decode(&viaCep); err != nil {
			log.Printf("Error decoding ViaCEP response: %v", err)
			return
		}
		if viaCep.Cep == "" {
			ch <- fmt.Sprintf("Error from %s: CEP not found or other error", apiName)
			return
		}
		addr = Address{
			Cep:        viaCep.Cep,
			Logradouro: viaCep.Logradouro,
			Bairro:     viaCep.Bairro,
			Cidade:     viaCep.Localidade,
			Uf:         viaCep.Uf,
		}
	} else {
		// Se for BrasilAPI
		var brasilApi BrasilAPI
		if err := json.NewDecoder(resp.Body).Decode(&brasilApi); err != nil {
			log.Printf("Error decoding BrasilAPI response: %v", err)
			return
		}
		addr = Address{
			Cep:        brasilApi.Cep,
			Logradouro: brasilApi.Street,
			Bairro:     brasilApi.Neighborhood,
			Cidade:     brasilApi.City,
			Uf:         brasilApi.State,
		}
	}

	// Criando uma representação JSON com indentação de 4 espaços para melhor visualização
	jsonAddr, _ := json.MarshalIndent(addr, "", "    ")
	result := fmt.Sprintf("Response from %s: \n%s", apiName, jsonAddr)
	ch <- result
	log.Printf("Busca de %s foi concluída em: %v\n", apiName, time.Since(starProgress))
}
