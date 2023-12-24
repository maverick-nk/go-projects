/**
A simple server to serve
1. static index html
2. get health
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func convertToJson(input map[string]interface{}) string {
	jsonByte, err := json.Marshal(input)
	if err != nil {
		fmt.Println("Error when converting response to json")
		return ""
	}
	return string(jsonByte)
}

func healthHandler(responseWriter http.ResponseWriter, request *http.Request) {
	// validations
	if request.URL.Path != "/health" {
		http.Error(responseWriter, "404 Not Found", http.StatusNotFound)
		return
	} else if request.Method != "GET" {
		http.Error(responseWriter, "Method Not Allowed", http.StatusMethodNotAllowed)
	} else {
		responseBody := make(map[string]interface{})
		responseBody["host"] = request.Host
		responseBody["servedAt"] = time.Now()

		// add headers
		for name, headers := range request.Header {
			for _, h := range headers {
				responseBody[name] = h
			}
		}
		fmt.Fprint(responseWriter, convertToJson(responseBody))
	}
}

func main() {
	// Configure logger
	logger := log.New(os.Stdout, "", log.LstdFlags)
	// Command line arguments
	portPtr := flag.Int("p", 8080, "Port number")
	flag.Parse()

	// HTTP Server
	// Views
	// 1. Static index.html
	http.Handle("/", http.FileServer(http.Dir("./static")))
	// 2. Get health
	http.HandleFunc("/health", healthHandler)

	// Start HTTP server at port
	logger.Printf("Starting server at port: %d\n", *portPtr)
	portConfig := ":" + strconv.Itoa(*portPtr)
	logger.Println("Visit - http://localhost", portConfig)
	log.Fatal(http.ListenAndServe(portConfig, nil))
}
