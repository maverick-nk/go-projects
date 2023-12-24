/**
A simple server to serve
1. static index html
2. get health
*/

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
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
	portConfig := ":" + strconv.Itoa(*portPtr)
	server := &http.Server{
		Addr: portConfig,
	}
	// Views
	// 1. Static index.html
	http.Handle("/", http.FileServer(http.Dir("./static")))
	// 2. Get health
	http.HandleFunc("/health", healthHandler)

	go func() {
		// Start HTTP server at port
		logger.Printf("Starting server at port: %d\n", *portPtr)
		logger.Println("Visit - http://localhost", portConfig)
		if err := http.ListenAndServe(portConfig, nil); err != nil {
			logger.Fatal(err)
		}
	}()

	// Handle graceful termination
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigs
	logger.Println("Signal received: ", sig)

	shutdownContext, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()
	if err := server.Shutdown(shutdownContext); err != nil {
		logger.Fatal(err)
	}

	logger.Println("Gracefully shutting down")
	logger.Println("Server shutdown! Exiting")
}
