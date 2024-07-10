package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// DollarQuote represents the JSON received from the API (https://economia.awesomeapi.com.br/json/last/USD-BRL)
type DollarQuote struct {
	ID         string `json:"id,omitempty"`
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

// DollarQuoteReceived represents the JSON received from the API (https://economia.awesomeapi.com.br/json/last/USD-BRL)
type DollarQuoteReceived struct {
	USDBRL DollarQuote `json:"USDBRL"`
}

// Global variable to store the database connection
var db *sql.DB

func main() {
	// Create the quotes table if it doesn't exist, using SQLite and import.sql in this directory
	// Open the database connection
	var err error
	db, err = sql.Open("sqlite3", "./quotes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Register the handler for the /cotacao endpoint
	http.HandleFunc("/cotacao", handleCotacao)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCotacao(w http.ResponseWriter, r *http.Request) {
	// Definindo o timeout para 200ms conforme requisito
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	var quote DollarQuoteReceived
	err := fetchDollarQuote(ctx, &quote)
	if err != nil {
		// Verificação específica para erros de contexto, como timeout
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Timeout fetching dollar quote:", err)
			http.Error(w, "Timeout fetching dollar quote. Try again later", http.StatusRequestTimeout)
			return
		}

		log.Println("Error fetching dollar quote:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctxDB, cancelDB := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancelDB()

	err = saveToDB(ctxDB, db, quote)
	if err != nil {
		// Verificação específica para erros de contexto, como timeout
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Timeout save dollar quote:", err)
		} else {
			log.Println("Error fetching dollar quote:", err)
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote.USDBRL.Bid)
}

func fetchDollarQuote(ctx context.Context, quote *DollarQuoteReceived) error {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(quote)
	if err != nil {
		log.Println("Error decoding response:", err)
		return err
	}

	return nil
}

func saveToDB(ctxDB context.Context, db *sql.DB, quote DollarQuoteReceived) error {
	_, err := db.ExecContext(ctxDB, `
		INSERT INTO quotes (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, createDate) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		quote.USDBRL.Code, quote.USDBRL.Codein, quote.USDBRL.Name,
		quote.USDBRL.High, quote.USDBRL.Low, quote.USDBRL.VarBid,
		quote.USDBRL.PctChange, quote.USDBRL.Bid, quote.USDBRL.Ask,
		quote.USDBRL.Timestamp, quote.USDBRL.CreateDate)
	return err
}
