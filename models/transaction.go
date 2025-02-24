package models

import "time"

type Transaction struct {
	Value    float64   `json:"valor"`
	DateTime time.Time `json:"dataHora"`
}
