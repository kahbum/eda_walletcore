package main

import (
	"database/sql"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/kahbum/eda_balance/internal/database"
	"github.com/kahbum/eda_balance/internal/event/handler"
	"github.com/kahbum/eda_balance/internal/usecase/get_account_balance"
	"github.com/kahbum/eda_balance/internal/web"
	"github.com/kahbum/eda_balance/internal/web/webserver"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "balance-mysql", "3306", "balance"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	accountDB := database.NewAccountDB(db)
	getAccountBalanceUseCase := get_account_balance.NewGetAccountBalanceUseCase(accountDB)

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	updateBalanceKafkaHandler := handler.NewUpdateBalanceKafkaHandler(configMap, accountDB)
	go updateBalanceKafkaHandler.Handle()

	webserver := webserver.NewWebServer(":3003")

	accountHandler := web.NewWebAccountHandler(*getAccountBalanceUseCase)
	webserver.AddHandler("/balances/{account_id}", accountHandler.GetAccountBalance)

	fmt.Println("Server is running")
	webserver.Start()
}
