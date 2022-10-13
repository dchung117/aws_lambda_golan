package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// create event and response data types
type MyEvent struct {
	Name string `json:"What is your name?"`
	Age  int    `json:"How old are you?"`
}

type MyResponse struct {
	Message string `json:"Answer"`
}

// define lambda function
func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	// return the message w/ person's name and age
	return MyResponse{Message: fmt.Sprintf("%s is %d years old.", event.Name, event.Age)}, nil
}

func main() {
	// initialize internal lambda endpoint, listening for requests to pass to HandleLambdaEvent
	lambda.Start(HandleLambdaEvent)
}
