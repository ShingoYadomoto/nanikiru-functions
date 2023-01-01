package main

import (
	"github.com/ShingoYadomoto/nanikiru-functions/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	h := handler.Handler{}
	lambda.Start(h.GetAnswerLambdaHandler)
}
