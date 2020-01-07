package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Employee struct {
	Name    string
	Email   string
	Year    string
	Salary  string
	IsLeave string
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Init aws session
	// This is for us-east-1 Region
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	tableName := "Employees"
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := svc.Scan(params)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("%v", err.Error()),
			StatusCode: 500,
		}, nil
	}

	employees := []Employee{}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &employees)

	if err != nil {
		fmt.Println(err.Error())
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("%v", err.Error()),
			StatusCode: 500,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("%+v", employees),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
