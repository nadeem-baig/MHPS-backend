package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Lambda handler function
func HandleLambdaRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := fmt.Sprintf("Hello %s from Lambda!", request.QueryStringParameters["name"])
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       response,
	}, nil
}

// HTTP handler function for local server
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s from local server! v2", name)
}

func main() {
	// Check if the application is running in the AWS Lambda environment
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" {
		// If not in Lambda, start a local HTTP server
		http.HandleFunc("/hello", helloHandler)
		fmt.Println("Running locally on port 8080...")
		http.ListenAndServe(":8080", nil)
	} else {
		// If in Lambda, start the Lambda handler
		lambda.Start(HandleLambdaRequest)
	}
}
