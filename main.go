package main

import (
	"github.com/funfuck/lambda-sample-function/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.Handler)
}
