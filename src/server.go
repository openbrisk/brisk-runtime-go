package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	DefaultFunctionDirectory string = "/openbrisk/"
	DevFunctionDirectory     string = "../examples/"
)

var (
	moduleName         string      // 
	moduleDependencies string      // 
	functionHandler    string      // 
	functionTimeout    int    = 10 // NOTE: Define default value.
)

func main() {
	http.HandleFunc("/healthz", func(response http.ResponseWriter, request *http.Request) {
		if(request.Method == "GET") {
			response.Header().Set("Content-Type", "text/plain")
			response.Header().Set("Connection", "close")
			response.WriteHeader(http.StatusOK)
		} else {
			response.WriteHeader(http.StatusNotFound)
		}
	})

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		if(request.Method == "GET") {
			// TODO: Handle function invocation without parameters. 
		} else if(request.Method == "POST") {
			// TODO: Handle function invocation with parameters.
		} else {
			response.WriteHeader(http.StatusNotFound)
		}
	})

	fmt.Println("Listening on port 8080 ...")
	var error = http.ListenAndServe(":8080", nil)
	if error != nil {
		panic(error)
	}
}
