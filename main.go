package main

import (
	"ev/function/sample/hdlr"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(hdlr.Handler)
}
