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
	// Create table Business
	tableName := "Business"

	svc := store.getSession()
	inputList := &dynamodb.ListTablesInput{}

	result, err := svc.ListTables(inputList)
	if err != nil {
		return err
	}

	for _, n := range result.TableNames {
		if *n == tableName {
			log.Info("table already exists")
			return nil
		}
	}

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("name"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("name"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	log.WithField("tableName", tableName).Info("Created the table")

	_, err = svc.CreateTable(input)
	if err != nil {
		return err
	}

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
