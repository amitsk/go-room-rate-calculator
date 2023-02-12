package adapters

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DB struct {
	Client *dynamodb.Client
	Table  string
}

func NewDB(t string) DB {
	cfg, _ := config.LoadDefaultConfig(context.Background(),
		// config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:9000"}, nil
			},
		)),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "dummy", SecretAccessKey: "dummy", SessionToken: "dummy",
				Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}),
	)
	c := dynamodb.NewFromConfig(cfg)

	return DB{Client: c, Table: t}
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

func (db DB) TaxRate(zip string) (ZipcodeTaxRate, error) {
	zipcodeRate := ZipcodeTaxRate{Zipcode: zip}
	response, err := db.Client.GetItem(context.Background(), &dynamodb.GetItemInput{
		Key: zipcodeRate.GetKey(), TableName: aws.String(db.Table),
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
