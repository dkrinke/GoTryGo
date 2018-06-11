package main

import (
	"encoding/json"               //For creating JSON responses
	"fmt"                         //STDOUT
	"github.com/gorilla/handlers" //Provides CORS which tells the browser that it's OK for JS to communicate with the server (go get github.com/gorilla/handlers)
	"github.com/gorilla/mux"      //Router that will take requests and decide what should be done (go get github.com/gorilla/mux)
	"log"                         //Logs when the server exits
	"net/http"                    //Provides the representation of HTTP requests, responses, and is Responsible for running the server
	"time"                        //Provides functionality for measuring and displaying time
)

type ResponseDTO struct {
	Title string
	Body  string
	Time  time.Time
}

type RequestDTO struct {
	Title string
	Body  string
}

var (
	response = ResponseDTO{Title: "default title", Body: "default body", Time: time.Now()}
  request RequestDTO //Used for incoming requests
)

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/healthcheck", healthCheck).Methods("GET")
	router.HandleFunc("/message", handleQryMessage).Methods("GET")
	router.HandleFunc("/message", handlePostMessage).Methods("POST")

	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":3000",
		handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func handleQryMessage(w http.ResponseWriter, r *http.Request) {
  //Send json response
	json.NewEncoder(w).Encode(map[string]ResponseDTO{"message": response})
}

func handlePostMessage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request) //Decode the incoming Request object
	if err != nil {
		panic(err)
	}

  //Assigning the incoming arguments to the response object
  response.Title = request.Title
  response.Body = request.Body
	response.Time = time.Now()
  
	defer r.Body.Close() //Defer the close of the body

  //Send json response
	json.NewEncoder(w).Encode(map[string]ResponseDTO{"message": response})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Still alive!") //Health check to confirm the server is running
}
