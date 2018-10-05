package main

//http://www.somkiat.cc/hello-golang-in-aws-lambda/
import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID    float64 `json:"id"`
	Value string  `json:"value"`
}

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func Handler(request Request) (Response, error) {
	return Response{
		Message: fmt.Sprintf("Request ID %f", request.ID),
		Ok:      true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
