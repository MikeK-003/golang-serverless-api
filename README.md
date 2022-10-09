# golang-serverless-api

This is my first project using golang and AWS, it aims to create a serverless api using those two.

This is just a demonstration of what I know in the 2 days I've been studying this.

Reproducing the project requires 
1. Visual Studio Code, with golang and the REST Client extension.
2. An AWS account to utilize AWS API Gateway and AWS Lambda.

I will write exactly how I ran this code:

1. Download all the contents of this repository as a ZIP. Extract wherever.
2. Open AWS Lambda and create a new function. Change its runtime to Go 1.x and upload its source code from the .zip file. Edit the runtime handler to main.
3. Open AWS API Gateway and create a HTTP API. Have it integrate the lambda function from before.
4. Copy the invoke URL, open tests.http in Visual Studio Code, and replace its existing link with yours. 
5. Using the REST Client extension in VSC, click the Send Request button for a successful output.
