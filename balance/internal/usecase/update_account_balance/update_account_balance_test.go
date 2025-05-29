package update_account_balance

import (
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/kahbum/eda_balance/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *AccountGatewayMock) UpdateBalance(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func TestUpdateAccountBalanceUseCase_NewAccount_Execute(t *testing.T) {
	inputDTO := UpdateAccountBalanceInputDTO{
		ID:      uuid.New().String(),
		Balance: 100,
	}
	accountMock := &AccountGatewayMock{}
	accountMock.On("FindByID", mock.Anything).Return(nil, sql.ErrNoRows)
	accountMock.On("Save", mock.MatchedBy(func(account *entity.Account) bool {
		return account.ID == inputDTO.ID && account.Balance == inputDTO.Balance
	})).Return(nil)

	uc := NewUpdateAccountBalanceUseCase(accountMock)
	err := uc.Execute(inputDTO)
	assert.Nil(t, err)
	accountMock.AssertExpectations(t)
	accountMock.AssertNumberOfCalls(t, "FindByID", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)
	accountMock.AssertNumberOfCalls(t, "UpdateBalance", 0)
}

func TestUpdateAccountBalanceUseCase_ExistingAccount_Execute(t *testing.T) {
	account := entity.NewAccount(uuid.New().String())
	inputDTO := UpdateAccountBalanceInputDTO{
		ID:      account.ID,
		Balance: 100,
	}
	accountMock := &AccountGatewayMock{}
	accountMock.On("FindByID", mock.Anything).Return(account, nil)
	accountMock.On("UpdateBalance", mock.MatchedBy(func(account *entity.Account) bool {
		return account.ID == inputDTO.ID && account.Balance == inputDTO.Balance
	})).Return(nil)

	uc := NewUpdateAccountBalanceUseCase(accountMock)
	err := uc.Execute(inputDTO)
	assert.Nil(t, err)
	accountMock.AssertExpectations(t)
	accountMock.AssertNumberOfCalls(t, "FindByID", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 0)
	accountMock.AssertNumberOfCalls(t, "UpdateBalance", 1)
}
