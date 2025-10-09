package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest() (string, error) {
	return "Hello, World!", nil
}

func main() {
	lambda.Start(HandleRequest)
}