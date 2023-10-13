package main

import (
	"context"
	"flag"
	"log"
	"read-adviser-bot/storage/sql"

	tgClient "read-adviser-bot/clients/telegram"
	"read-adviser-bot/consumer/event-consumer"
	"read-adviser-bot/events/telegram"
)

const (
	tgBotHost         = "api.telegram.org"
	inFileStoragePath = "files_storage"
	mySqlStoragePath  = "root:password@tcp(docker.for.mac.localhost:3306)/mydbname?parseTime=true"
	batchSize         = 100
)

func main() {
	sqlStorage, err := sql.New(mySqlStoragePath)
	if err != nil {
		log.Fatal("could not create sql storage", err)
	}
	sqlStorage.Init(context.TODO())

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		sqlStorage,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
