package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	id := uuid.New().String()
	account := NewAccount(id)
	assert.NotNil(t, account)
	assert.Equal(t, id, account.ID)
	assert.Equal(t, float64(0), account.Balance)
}

func TestCreateAccountWithEmptyID(t *testing.T) {
	account := NewAccount("")
	assert.Nil(t, account)
}

func TestUpdateBalance(t *testing.T) {
	account := NewAccount(uuid.New().String())
	account.UpdateBalance(100)
	assert.Equal(t, float64(100), account.Balance)
}
