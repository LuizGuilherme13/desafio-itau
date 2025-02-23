package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Transaction struct {
	Value    float64   `json:"valor"`
	DateTime time.Time `json:"dataHora"`
}

func newTransaction(w http.ResponseWriter, r *http.Request) {
	t := Transaction{}

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if t == (Transaction{}) {
		http.Error(w, "", http.StatusUnprocessableEntity)
		return
	}

	if t.DateTime.After(time.Now()) {
		http.Error(w, "", http.StatusUnprocessableEntity)
		return
	}

	if t.Value < 0 {
		http.Error(w, "", http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func main() {
	http.HandleFunc("/transacao", newTransaction)

	log.Println("Server running on port:", 8080)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
