package get_account_balance

import (
	"github.com/kahbum/eda_balance/internal/gateway"
)

type GetAccountBalanceInputDTO struct {
	ID string `json:"id"`
}

type GetAccountBalanceOutputDTO struct {
	ID      string  `json:"id"`
	Balance float64 `json:"balance"`
}

type GetAccountBalanceUseCase struct {
	AccountGateway gateway.AccountGateway
}

func NewGetAccountBalanceUseCase(a gateway.AccountGateway) *GetAccountBalanceUseCase {
	return &GetAccountBalanceUseCase{
		AccountGateway: a,
	}
}

func (uc *GetAccountBalanceUseCase) Execute(input GetAccountBalanceInputDTO) (*GetAccountBalanceOutputDTO, error) {
	account, err := uc.AccountGateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetAccountBalanceOutputDTO{
		ID:      account.ID,
		Balance: account.Balance,
	}, nil
}
