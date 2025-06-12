package main

import (
	"bufio"
	"strings"
	"syscall/js"
)

func main() {

	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}

// read https://pkg.go.dev/syscall/js#FuncOf for more information about js.FuncOf
func registerCallbacks() {
	js.Global().Set("validateCSV", js.FuncOf(validateCSV))
}

func validateCSV(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return "No CSV content provided."
	}
	csvContent := args[0].String()

	// Use a scanner to process the CSV line by line to save memory
	scanner := bufio.NewScanner(strings.NewReader(csvContent))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			println(line)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	if strings.TrimSpace(csvContent) == "" {
		return "CSV is empty."
	}
	if !strings.Contains(csvContent, ",") {
		return "CSV does not appear to have any columns."
	}
	// Further validation logic can be added here
	return csvContent
}
