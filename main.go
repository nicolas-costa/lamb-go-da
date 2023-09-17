package main

import "github.com/aws/aws-lambda-go/lambda"

func pong() (string, error) {
	return "pong", nil
}

func main() {
	lambda.Start(pong)
}
