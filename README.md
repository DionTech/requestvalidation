# about
This package is a wrapper for the github.com/go-playground/validator package.

You can validate a request (by struct) and get a status ("success" or "error") and when error, a map[string][]string of errors as a response.
This might be useful for a json-formatted api response.

You can also define your own error messages, using the struct field tag `message:"validation message here"`.

## usage, for example at aws lambda

``` 
go get go get github.com/DionTech/requestvalidation
```

```go 
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DionTech/requestvalidation"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)


type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest
type ExampleRequest struct {
    Name string `json:"name" validate:"required" message:"name is required"`
    Age int `json:"age" validate:"required,min=18" message:"age is required and must be 18 or older"`
}

func Handler(request Request) (Response, error) {
    var buf bytes.Buffer
    
    exampleRequest := ExampleRequest{}
    err := json.Unmarshal([]byte(request.Body), &exampleRequest)
    
    if err != nil {
		return handleError(err, 500), nil
	}

	requestValidator := requestvalidation.NewRequestValidator()
	validation, err := requestValidator.Validate(exampleRequest)

	if err != nil {
		body, err := json.Marshal(validation)
		if err != nil {
			return handleError(err, 500), nil
		}

		json.HTMLEscape(&buf, body)
		resp := Response{
			StatusCode:      422,
			IsBase64Encoded: false,
			Body:            buf.String(),
			Headers:         headers,
		}

		return resp, nil
	}
	
	//do some stuff here
    
}

func main() {
	lambda.Start(Handler)
}

func handleError(err error, statusCode int) Response {
	fmt.Printf("ERROR PDFSERVICE: %s", err.Error())

	return Response{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       err.Error(),
	}
}
```