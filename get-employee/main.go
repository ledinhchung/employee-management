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
	employeeName := "Dat"

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String(employeeName),
			},
		},
	})

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("%v", err.Error()),
			StatusCode: 500,
		}, nil
	}

	employee := Employee{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &employee)

	if err != nil {
		fmt.Println(err.Error())
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("%v", err.Error()),
			StatusCode: 500,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Welcome to GoLang and Lambda %v, %v", employee.Name, employee.Email),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
