package awslambda

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	path := request.RequestContext.HTTP.Path

	response := events.LambdaFunctionURLResponse{}
	var err error = nil

	switch path {
	case "/greetings":
		response, err = greet()
	case "/hello":
		response, err = hello()
	default:
		response = events.LambdaFunctionURLResponse{
			StatusCode: 404,
			Body:       "invalid request",
		}
	}
	return response, err
}

func greet() (events.LambdaFunctionURLResponse, error) {
	return events.LambdaFunctionURLResponse{
		Body:       "Greetings",
		StatusCode: 200,
	}, nil
}

func hello() (events.LambdaFunctionURLResponse, error) {
	return events.LambdaFunctionURLResponse{
		Body:       "hello",
		StatusCode: 200,
	}, nil
}
