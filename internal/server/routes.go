package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"slices"
	"time"

	"github.com/LuizGuilherme13/desafio-itau/internal/models"
	"github.com/LuizGuilherme13/desafio-itau/internal/utils/clog"
)

func (s *Server) MountRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /transacao", s.HandleNewTransaction)
	mux.HandleFunc("DELETE /transacao", s.HandleDeleteTransactions)
	mux.HandleFunc("GET /estatistica", s.HandleGetStatistic)

	return logMiddleware(mux)
}

func logMiddleware(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		clog.Info(r.Method, r.URL.Path)

		f.ServeHTTP(w, r)

	})
}

func (s *Server) HandleNewTransaction(w http.ResponseWriter, r *http.Request) {
	t := models.Transaction{}

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		clog.Error("(HandleNewTransaction)", errors.New("invalid json"))
		return
	}

	if t == (models.Transaction{}) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		clog.Error("(HandleNewTransaction)", errors.New("empty body"))
		return
	}

	if t.DateTime.After(time.Now()) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		clog.Error("(HandleNewTransaction)", errors.New("dateTime cannot be in the future"))
		return
	}

	if t.Value < 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		clog.Error("(HandleNewTransaction)", errors.New("value cannot be less than 0"))
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

	statistic := models.Statistic{}

	if len(s.Store.Transactions) > 0 {
		now := time.Now()
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
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(statistic); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
