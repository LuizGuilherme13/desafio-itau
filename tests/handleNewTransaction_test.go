package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/LuizGuilherme13/desafio-itau/api"
)

func TestHandleNewTransaction(t *testing.T) {
	tests := []struct {
		name         string
		body         string
		expectedCode int
	}{
		{"test1", fmt.Sprintf(`{"valor": 100.5, "dataHora": "%s"}`, time.Now().Format(time.RFC3339)), 201},
		{"test2", fmt.Sprintf(`{"valor": -999, "dataHora": "%s"}`, time.Now().Format(time.RFC3339)), 422},
		{"test3", fmt.Sprintf(`{"valor": 10.2345, "dataHora": "%s"}`, time.Now().Format(time.RFC3339)), 201},
		{"test4", fmt.Sprintf(`{"valor": 100.5, "dataHora": "%s"}`, time.Now().Add(1*time.Hour).Format(time.RFC3339)), 422},
		{"test5", ``, 400},
		{"test6", `{`, 400},
		{"test7", `{}`, 422},
	}

	for _, tt := range tests {
		r := httptest.NewRequest(http.MethodPost, "/transacao", bytes.NewBufferString(tt.body))
		w := httptest.NewRecorder()

		server := api.NewServer(":8080")

		server.HandleNewTransaction(w, r)

		res := w.Result()

		if tt.expectedCode != res.StatusCode {
			t.Errorf("%s FAIL - expected: %d, got: %d \n", tt.name, tt.expectedCode, res.StatusCode)
		}
	}
}
