package get_account_balance

import (
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
	account := entity.NewAccount(uuid.New().String())
	account.UpdateBalance(100)
	accountMock := &AccountGatewayMock{}
	accountMock.On("FindByID", mock.Anything).Return(account, nil)

	uc := NewGetAccountBalanceUseCase(accountMock)
	inputDTO := GetAccountBalanceInputDTO{
		ID: account.ID,
	}
	accountResult, err := uc.Execute(inputDTO)
	assert.Nil(t, err)
	accountMock.AssertExpectations(t)
	accountMock.AssertNumberOfCalls(t, "FindByID", 1)

	assert.Equal(t, account.ID, accountResult.ID)
	assert.Equal(t, account.Balance, accountResult.Balance)
}
