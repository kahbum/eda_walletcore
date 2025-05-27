package gateway

import "github.com/kahbum/eda_walletcore/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
