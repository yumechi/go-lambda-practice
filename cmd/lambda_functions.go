// refer: https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/golang-handler.html
package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

type MyResponse struct {
	Message string `json:"Answer"`
}

func HandleLambdaEvent(ctx context.Context, event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
