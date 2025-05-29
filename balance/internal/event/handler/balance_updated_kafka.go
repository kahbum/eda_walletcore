package handler

import (
	"encoding/json"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/kahbum/eda_balance/internal/gateway"
	"github.com/kahbum/eda_balance/internal/usecase/update_account_balance"
	"github.com/kahbum/eda_balance/pkg/kafka"
)

type UpdateBalanceKafkaHandlerInputDTO struct {
	Name    string `json:"Name"`
	Payload struct {
		AccountIDFrom        string  `json:"account_id_from"`
		AccountIDTo          string  `json:"account_id_to"`
		BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
		BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
	} `json:"Payload"`
}

type UpdateBalanceKafkaHandler struct {
	Kafka          *kafka.Consumer
	AccountGateway gateway.AccountGateway
}

func NewUpdateBalanceKafkaHandler(configMap ckafka.ConfigMap, accountGateway gateway.AccountGateway) *UpdateBalanceKafkaHandler {
	kafkaConsumer := kafka.NewConsumer(&configMap, []string{"balances"})
	return &UpdateBalanceKafkaHandler{
		Kafka:          kafkaConsumer,
		AccountGateway: accountGateway,
	}
}

func (h *UpdateBalanceKafkaHandler) Handle() {
	updateAccountBalanceUseCase := update_account_balance.NewUpdateAccountBalanceUseCase(h.AccountGateway)

	messageChan := make(chan *ckafka.Message)
	go h.Kafka.Consume(messageChan)

	for message := range messageChan {
		var inputDTO UpdateBalanceKafkaHandlerInputDTO
		if err := json.Unmarshal([]byte(message.Value), &inputDTO); err != nil {
			fmt.Printf("Error decoding update balance message from kafka: %s\n", err)
		} else {
			err := updateAccountBalanceUseCase.Execute(update_account_balance.UpdateAccountBalanceInputDTO{
				ID:      inputDTO.Payload.AccountIDFrom,
				Balance: inputDTO.Payload.BalanceAccountIDFrom,
			})
			if err != nil {
				fmt.Println(err)
			}

			err = updateAccountBalanceUseCase.Execute(update_account_balance.UpdateAccountBalanceInputDTO{
				ID:      inputDTO.Payload.AccountIDTo,
				Balance: inputDTO.Payload.BalanceAccountIDTo,
			})
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
