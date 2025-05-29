package update_account_balance

import (
	"github.com/kahbum/eda_balance/internal/entity"
	"github.com/kahbum/eda_balance/internal/gateway"
)

type UpdateAccountBalanceInputDTO struct {
	ID      string  `json:"id"`
	Balance float64 `json:"balance"`
}

type UpdateAccountBalanceUseCase struct {
	AccountGateway gateway.AccountGateway
}

func NewUpdateAccountBalanceUseCase(a gateway.AccountGateway) *UpdateAccountBalanceUseCase {
	return &UpdateAccountBalanceUseCase{
		AccountGateway: a,
	}
}

func (uc *UpdateAccountBalanceUseCase) Execute(input UpdateAccountBalanceInputDTO) error {
	account := entity.NewAccount(input.ID)
	account.UpdateBalance(input.Balance)
	_, err := uc.AccountGateway.FindByID(account.ID)
	if err != nil {
		err := uc.AccountGateway.Save(account)
		if err != nil {
			return err
		}
		return nil
	}

	err = uc.AccountGateway.UpdateBalance(account)
	if err != nil {
		return err
	}
	return nil
}
