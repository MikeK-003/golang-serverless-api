package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// the main aws url in tests.http is the "endpoint"
// it hits the lambda func which is the main func
// main func called in lambda, which calls handler func
// note that the url is how api gateway works. prior to it i had to test in lambda
func main() {
	lambda.Start(handler)
}

// this handler func takes in proxy requests and returns both a response and an error (is nil)
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// 1. (cont), the response has a body field and that is what we want to write into

	var person1 Person
	// we want to convert the JSON data and unmarshal the data to store it in this empty struct,
	// so that our future msg var uses that data, and we will then marshal the data back to JSON
	// using unmarshal func, pass in slice of bytes, equate it to the string, and store it into our new struct
	err := json.Unmarshal([]byte(request.Body), &person1)

	// if there is an error
	if err != nil {
		// pass an empty response, but also an error as the func is expecting one. norm error = empty reponse, code 500
		return events.APIGatewayProxyResponse{}, err
	}

	// 2. (cont), make a string out of the new go struct that we made from JSON data via unmarshalling
	msg := fmt.Sprintf("hello %v %v", *person1.FirstName, *person1.LastName)

	// now we will create the response struct. as its string field is a ptr, dereference it via ampersand
	responsebody1 := ResponseBody{
		Message: &msg,
	}

	// now we must convert the struct to JSON data.
	// jbytes is the name of our output, it is a function, it takes in any data the coder wants
	// (like how unmarshal outputted any data the coder wants)
	// and outputs both string of bytes and error status. hence the "jbytes, err =", they both equal the func's outputs.
	// here, we took in the responseBody struct with that specific string
	jbytes, err := json.Marshal(responsebody1)

	// if there is an error
	if err != nil {
		// pass an empty response, but also an error as the func is expecting one. norm error = empty reponse, code 500
		return events.APIGatewayProxyResponse{}, err
	}

	// 3. if we've encountered no errors during this whole process, create a response for the overall func to return
	// the response, again, is a struct. 2 of its fields include:
	// a body (the jbytes converted into a string), and a statuscode
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jbytes),
	}

	// return response and err value of nil
	return response, nil
}

// 1. this new program uses inputs/outputs of data.
// we must first use the json marshaler to convert incoming data into structs
// here is a struct. note the explicit json explanations
type Person struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}

// 2. at this point, with the creation of a struct and an unmarshal of json data (which is in the .http file) into a struct
// now we want to carry out the input into an output
// to this end, create a new struct
type ResponseBody struct {
	Message *string `json:"message"`
}
