package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
	"time"

	"github.com/LuizGuilherme13/desafio-itau/api"
	"github.com/LuizGuilherme13/desafio-itau/models"
)

func TestHandleGetStatistic(t *testing.T) {
	expectedStatistic := models.Statistic{}

	server := api.NewServer(":8080")
	server.Store.Transactions = []models.Transaction{
		{Value: 100.00, DateTime: time.Now().Add(-1 * time.Minute)},
		{Value: 99.11, DateTime: time.Now().Add(-1 * time.Minute)},
		{Value: 88.22, DateTime: time.Now().Add(-1 * time.Minute)},
		{Value: 77.33, DateTime: time.Now().Add(-1 * time.Minute)},
		{Value: 66.44, DateTime: time.Now().Add(-1 * time.Minute)},
		{Value: 55.55, DateTime: time.Now().Add(-20 * time.Second)},
		{Value: 44.66, DateTime: time.Now().Add(-50 * time.Second)},
		{Value: 22.77, DateTime: time.Now().Add(-59 * time.Second)},
		{Value: 11.88, DateTime: time.Now()},
	}

	now := time.Now()
	values := []float64{}

	for _, t := range server.Store.Transactions {
		diff := now.Sub(t.DateTime)

		if diff.Seconds() <= 60 {
			expectedStatistic.Count++
			expectedStatistic.Sum += t.Value
			expectedStatistic.Avg = expectedStatistic.Sum / float64(expectedStatistic.Count)

			values = append(values, t.Value)
		}

	}

	slices.Sort(values)

	expectedStatistic.Min = values[0]
	expectedStatistic.Max = values[len(values)-1]

	r := httptest.NewRequest(http.MethodGet, "/estatistica", nil)
	w := httptest.NewRecorder()

	server.HandleGetStatistic(w, r)

	res := w.Result()

	if res.StatusCode != 200 {
		t.Errorf("FAIL - expected: %d, got: %d", 200, res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	statistic := models.Statistic{}

	if err := json.Unmarshal(body, &statistic); err != nil {
		t.Error(err)
	}

	if expectedStatistic.Count != statistic.Count {
		t.Errorf("FAIL - Count expected: %d, got: %d", expectedStatistic.Count, statistic.Count)
	}
	if expectedStatistic.Sum != statistic.Sum {
		t.Errorf("FAIL - Sum expected: %f, got: %f", expectedStatistic.Sum, statistic.Sum)
	}
	if expectedStatistic.Avg != statistic.Avg {
		t.Errorf("FAIL - Avg expected: %f, got: %f", expectedStatistic.Avg, statistic.Avg)
	}
	if expectedStatistic.Min != statistic.Min {
		t.Errorf("FAIL - Min expected: %f, got: %f", expectedStatistic.Min, statistic.Min)
	}
	if expectedStatistic.Max != statistic.Max {
		t.Errorf("FAIL - Max expected: %f, got: %f", expectedStatistic.Max, statistic.Max)
	}

}
