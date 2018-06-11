# GoTryGo

Practice implementation of the Go programming language

helloGo: contains exploritory implementation of Go while following The Complete Google's Go (golang) Programming Course
https://www.udemy.com/the-complete-google-golang-programming

goAPI: contains a go api example that can receive json requests and send json responses

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development.

### Prerequisites

Install Go: https://golang.org/doc/install

### How to run

Run the following commands
- helloGo
  - Build helloGo by executing in your terminal: go install helloGo
  - Run helloGo by executing in your terminal: "bin/helloGo" from the root directory
- goAPI
  - Build goAPI by executing in your terminal: go install goAPI
  - Run goAPI by executing in your terminal: "bin/goAPI" from the root directory

### Example JSON request
- GET http://localhost:3000/message 
- POST http://localhost:3000/message 
Body: {
	"Title": "Provide a title",
	"Body": "Provide a message"
}

### Example JSON response
{"message":{"Title":"Message title" "Body":"Message Body","Time":"2018-06-11T01:07:06.404700022-04:00"}}
