package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/openbrisk/brisk-runtime-go/src/util"
)

var (
	moduleName         string // $MODULE_NAME
	moduleDependencies string // $MODULE_NAME
	functionHandler    string // $FUNCTION_HANDLER
	functionTimeout    = 10   // $FUNCTION_TIMEOUT
	envError           error
	function           func(string) string /* func(any) any */
)

func main() {
	readEnvironment()

	// Load the function.
	function = util.LoadFunction(moduleName, functionHandler)

	http.HandleFunc("/healthz", func(response http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			response.WriteHeader(http.StatusOK)
		} else {
			response.WriteHeader(http.StatusNotFound)
		}
	})

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case
			"GET", "POST":
			executeFunction(response, request)
			break
		default:
			response.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	log.Println("Listening on port 8080 ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// readEnvironment reads the environment variables of the function and logs the results.
func readEnvironment() {
	moduleName = os.Getenv("MODULE_NAME")
	log.Printf("Using module: %s.so", moduleName)

	moduleDependencies = os.Getenv("MODULE_NAME")
	log.Printf("Using module deps: %s", moduleDependencies)

	functionHandler = os.Getenv("FUNCTION_HANDLER")
	log.Printf("Using function handler: %s", functionHandler)

	if os.Getenv("FUNCTION_TIMEOUT") != "" {
		functionTimeout, envError = strconv.Atoi(os.Getenv("FUNCTION_TIMEOUT"))
		if envError != nil {
			log.Fatal(envError)
		}
	}
	log.Printf("Using function timeout: %d seconds", functionTimeout)
}

// executeFunction executes the function.
func executeFunction(response http.ResponseWriter, request *http.Request) {
	start := time.Now()

	var isHeaderWritten = false
	var functionInput []byte
	var functionOutput []byte

	functionInput = getRequestBody(request)

	log.Println("Running function")
	functionOutput = []byte(function(string(functionInput)))

	// Check if the result contains the forward structure.
	/*functionResult, marshalError := util.UnmarshalFunctionResult(functionOutput)
	if marshalError == nil {
		log.Println("Found function result structure: writing forward header")
		decodedResult, decodingError := base64.StdEncoding.DecodeString(functionResult.Result)
		if decodingError == nil {
			log.Println("Decoding base64 encoded function result")
			functionOutput = decodedResult
		}
		if !isHeaderWritten {
			response.Header().Set("X-OpenBrisk-Forward", functionResult.Forward.To)
		}
	}*/

	// Write the duration header in nanoseconds.
	duration := time.Since(start)
	if !isHeaderWritten {
		response.Header().Set("X-OpenBrisk-Duration", fmt.Sprint(duration.Nanoseconds()))
		response.WriteHeader(http.StatusOK)
		response.Write(functionOutput)
		isHeaderWritten = true
	}
	log.Printf("Function execution duration: %s nanoseconds", fmt.Sprint(duration.Nanoseconds()))
}

// getRequestBody returns the request body data (the function input) as byte array.
func getRequestBody(request *http.Request) []byte {
	var body []byte
	var err error

	// No request body: return the empty byte array.
	if request.Body == nil {
		return body
	}

	// Close the body stream at the end of the function.
	defer request.Body.Close()

	// Read request body.
	body, err = ioutil.ReadAll(request.Body)
	// TODO: Better error handling: Send back per http response?
	if err != nil {
		return body
	}

	return body
}
