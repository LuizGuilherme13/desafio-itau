package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/LuizGuilherme13/desafio-itau/internal/models"
	"github.com/LuizGuilherme13/desafio-itau/internal/server"
)

func TestHandleDeleteTransaction(t *testing.T) {
	server := server.New(":8080")
	server.Store = models.Storage{
		Transactions: []models.Transaction{
			{Value: 100.5, DateTime: time.Now()},
			{Value: 999, DateTime: time.Now()},
			{Value: 100.5555, DateTime: time.Now()},
			{Value: 1.006, DateTime: time.Now()},
			{Value: 77.1, DateTime: time.Now()},
		},
	}

	r := httptest.NewRequest(http.MethodDelete, "/transacao", nil)
	w := httptest.NewRecorder()

	server.HandleDeleteTransactions(w, r)

	if w.Result().StatusCode != 200 {
		t.Errorf("FAIL - expected: %d, got: %d", 200, w.Result().StatusCode)
	}

	if len(server.Store.Transactions) != 0 {
		t.Errorf("FAIL - expected: %v, got: %v", []models.Transaction{}, server.Store.Transactions)
	}
}
