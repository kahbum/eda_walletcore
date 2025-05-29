package entity

import (
	"time"
)

type Account struct {
	ID        string
	Balance   float64
	UpdatedAt time.Time
}

func NewAccount(id string) *Account {
	if id == "" {
		return nil
	}
	account := &Account{
		ID:        id,
		Balance:   0,
		UpdatedAt: time.Now(),
	}
	return account
}

func (a *Account) UpdateBalance(amount float64) {
	a.Balance = amount
	a.UpdatedAt = time.Now()
}
