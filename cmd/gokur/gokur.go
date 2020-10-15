package main

import (
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/joho/godotenv"
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

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	cfg := aws.Config{
		Endpoint: aws.String(os.Getenv("DYNAMO_ENDPOINT")),
		Region:   aws.String(os.Getenv("DYNAMO_REGION")),
	}
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess, &cfg)

	// Create table Movies
	tableName := "Business"

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("lat"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("long"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("name"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("address"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("name"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("RANGE"),
			},
			{
				AttributeName: aws.String("lat"),
				KeyType:       aws.String("RANGE"),
			},
			{
				AttributeName: aws.String("long"),
				KeyType:       aws.String("RANGE"),
			},
			{
				AttributeName: aws.String("address"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		return err
	}

	log.WithField("tableName", tableName).Info("Created the table")
	return nil
}
