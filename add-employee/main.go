package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type RequestEmployee struct {
	Name    string `json:"Name"`
	Email   string `json:"Email"`
	Year    string `json:"Year"`
	Salary  string `json:"Salary"`
	IsLeave string `json:"IsLeave"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Mapping request data
	inputData := RequestEmployee{}
	err := json.Unmarshal([]byte(request.Body), &inputData)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Request error %v", err.Error()),
			StatusCode: 500,
		}, nil
	}

	// Init aws session
	// This is for us-east-1 Region
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	item, err := dynamodbattribute.MarshalMap(inputData)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("DynamoDB error %v", err.Error()),
			StatusCode: 500,
		}, nil
	}

	tableName := "Employees"
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Save data error %v", err.Error()),
			StatusCode: 500,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Successfully"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
