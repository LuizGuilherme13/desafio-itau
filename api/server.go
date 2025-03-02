package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"slices"

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
	http.HandleFunc("GET /estatistica", s.HandleGetStatistic)

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
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) HandleDeleteTransactions(w http.ResponseWriter, r *http.Request) {
	s.Store.Transactions = []models.Transaction{}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) HandleGetStatistic(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	statistic := models.Statistic{}
	values := []float64{}

	for _, t := range s.Store.Transactions {
		diff := now.Sub(t.DateTime)

		if diff.Seconds() <= 60 {
			statistic.Count++
			statistic.Sum += t.Value
			statistic.Avg = statistic.Sum / float64(statistic.Count)

			values = append(values, t.Value)
		}
	}

	slices.Sort(values)

	statistic.Min = values[0]
	statistic.Max = values[len(values)-1]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(statistic); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
