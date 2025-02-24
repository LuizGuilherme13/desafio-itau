package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/LuizGuilherme13/desafio-itau/models"
)

type Server struct {
	Addr  string
	Store models.Storage
}

func NewServer(addr string) *Server {
	return &Server{Addr: addr, Store: models.Storage{}}
}

func (s *Server) Start() error {
	http.HandleFunc("POST /transacao", s.HandleNewTransaction)
	http.HandleFunc("DELETE /transacao", s.HandleDeleteTransactions)

	log.Println("Server running on port", s.Addr)
	return http.ListenAndServe(s.Addr, nil)
}

func (s *Server) HandleNewTransaction(w http.ResponseWriter, r *http.Request) {
	t := models.Transaction{}

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if t == (models.Transaction{}) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if t.DateTime.After(time.Now()) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if t.Value < 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	s.Store.Add(t)
	fmt.Println(s.Store.Transactions)
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) HandleDeleteTransactions(w http.ResponseWriter, r *http.Request) {
	s.Store.Transactions = []models.Transaction{}

	w.WriteHeader(http.StatusOK)
}
