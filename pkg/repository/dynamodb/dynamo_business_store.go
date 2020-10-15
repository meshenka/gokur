package dynamodb

import (
	"os"

	"github.com/meshenka/gokur/pkg/model"
	"github.com/meshenka/gokur/pkg/repository"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	log "github.com/sirupsen/logrus"
)

type dynamodbBusinessStore struct {
}

// NewDynamoBusinessStore constructor
func NewDynamoBusinessStore() repository.BusinessStore {
	return &dynamodbBusinessStore{}
}

func (store *dynamodbBusinessStore) Init() error {
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

	svc := store.getSession()

	_, err := svc.CreateTable(input)
	if err != nil {
		return err
	}

	log.WithField("tableName", tableName).Info("Created the table")
	return nil
}

func (store *dynamodbBusinessStore) GetByID() *model.Business {
	return nil
}

func (store *dynamodbBusinessStore) getSession() *dynamodb.DynamoDB {
	cfg := aws.Config{
		Endpoint: aws.String(os.Getenv("DYNAMO_ENDPOINT")),
		Region:   aws.String(os.Getenv("DYNAMO_REGION")),
	}
	sess := session.Must(session.NewSession())
	return dynamodb.New(sess, &cfg)

}
