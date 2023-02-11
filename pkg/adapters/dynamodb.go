package adapters

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoTable struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

// TableExists determines whether a DynamoDB table exists.
func (table DynamoTable) TableExists() (bool, error) {
	exists := true
	_, err := table.DynamoDbClient.DescribeTable(
		context.TODO(), &dynamodb.DescribeTableInput{TableName: aws.String(table.TableName)},
	)
	if err != nil {
		var notFoundEx *types.ResourceNotFoundException
		if errors.As(err, &notFoundEx) {
			log.Printf("Table %v does not exist.\n", table.TableName)
			err = nil
		} else {
			log.Printf("Couldn't determine existence of table %v. Here's why: %v\n", table.TableName, err)
		}
		exists = false
	}
	return exists, err
}

type ZipcodeTaxRate struct {
	Zipcode string  `dynamodbav:"zipcode"`
	Rate    float32 `dynamodbav:"rate"`
}

// GetKey returns the composite primary key of the movie in a format that can be
// sent to DynamoDB.
func (zipcodeTax ZipcodeTaxRate) GetKey() map[string]types.AttributeValue {
	zip, err := attributevalue.Marshal(zipcodeTax.Zipcode)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"zipcode": zip}
}

// String returns the title, year, rating, and plot of a movie, formatted for the example.
func (zipcodeTax ZipcodeTaxRate) String() string {
	return fmt.Sprintf("Zipcode: %v\n\tRate: %v\n",
		zipcodeTax.Zipcode, zipcodeTax.Rate)
}

func (table DynamoTable) TaxRate(zip string) (ZipcodeTaxRate, error) {
	zipcodeRate := ZipcodeTaxRate{Zipcode: zip}
	response, err := table.DynamoDbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: zipcodeRate.GetKey(), TableName: aws.String(table.TableName),
	})
	if err != nil {
		log.Printf("Couldn't get info about %v. Here's why: %v\n", zip, err)
	} else {
		err = attributevalue.UnmarshalMap(response.Item, &zipcodeRate)
		if err != nil {
			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}
	}
	return zipcodeRate, err
}
