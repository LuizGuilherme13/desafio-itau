package models

type Storage struct {
	Transactions []Transaction
}

func (s *Storage) Add(t Transaction) {
	s.Transactions = append(s.Transactions, t)
}
