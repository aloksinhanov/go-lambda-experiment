package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	awslambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type body struct {
	Message string `json:"message"`
}

// Handler is the Lambda function handler
func Mapper(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Lambda request", request.RequestContext.RequestID)

	b, _ := json.Marshal(body{Message: "hello mapper"})

	region := os.Getenv("AWS_REGION")

	sess, err := session.NewSession(&aws.Config{ // Use aws sdk to connect to dynamoDB
		Region: &region,
	})

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	client := lambda.New(sess)

	input := &lambda.InvokeInput{
		FunctionName:   aws.String("Transform"),
		InvocationType: aws.String("RequestResponse"),
		LogType:        aws.String("Tail"),
		Payload:        nil,
	}
	result, err := client.Invoke(input)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err.Error())
	} else {
		fmt.Println("Transformer called successfully!!")
	}

	fmt.Println(result)

	return events.APIGatewayProxyResponse{
		Body:       string(b),
		StatusCode: 200,
	}, nil
}

func main() {
	awslambda.Start(Mapper)
}
