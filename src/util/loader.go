package util

import (
	"log"
	"os"
	"plugin"
)

// LoadFunction loads the shared object containing the function handler.
func LoadFunction(moduleName string, functionHandler string) func(string) string {
	// Load module

	// 1. Open the so file to load the symbols
	//    defined by the module $MODULE_NAME.
	module, err := plugin.Open("/openbrisk/" + moduleName + ".so")
	if err != nil {
		log.Fatalf(err.Error())
		os.Exit(1)
	}
	log.Printf("Loaded module: %s.so", moduleName)

	// 2. Look up a symbol (the exported function)
	//    in this case, function defined by $FUNCTION_HANDLER
	symbol, err := module.Lookup(functionHandler)
	if err != nil {
		log.Fatalf(err.Error())
		os.Exit(1)
	}
	log.Printf("Loaded function symbol: %s", functionHandler)

	// 3. Assert that loaded symbol is of a desired type
	//    in this case func() string.
	/*function, ok := symbol.(func() string)
	if !ok {
		log.Fatalf("Function module is missing function: %s() string", functionHandler)
		os.Exit(1)
	}
	log.Printf("Created function delegate for handler: %s", functionHandler)*/

	//return function

	switch function := symbol.(type) {
	case func() string:
		return func(data string) string {
			return function()
		}
	case func(string) string:
		return function
	default:
		log.Fatalf("Function entry point with correct signature not found.")
		os.Exit(1)
	}

	return nil
}
