package main

import (
	"fmt"
	"io"
	"os"

	"github.com/joho/godotenv"
	store "github.com/meshenka/gokur/pkg/repository/dynamodb"
	log "github.com/sirupsen/logrus"
)

const (
	// exitFail is the exit code if the program fails.
	exitFail = 1
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdout io.Writer) error {
	initLogger()
	initEnv()
	err := initDynamoDb()

	return err
}

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.Info("welcome to gokur")
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// @see https://medium.com/@mcleanjnathan/testing-with-dynamo-local-and-go-7b7000ef9602
func initDynamoDb() error {
	businessStore := store.NewDynamoBusinessStore()
	return businessStore.Init()
}
