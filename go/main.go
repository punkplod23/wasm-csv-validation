package main

import (
	"bufio"
	"encoding/json"
	"strings"
	"syscall/js"
)

type CSVReport struct {
	Valid             bool          `json:"valid"`
	Errors            []string      `json:"errors"`
	Warning           []string      `json:"warning"`
	Info              []string      `json:"info"`
	BlankColumnsCount int           `json:"blank_columns_count"`
	BlankColumns      []BlankColumn `json:"blank_columns"`
	Headers           []string      `json:"headers,omitempty"`
	MappingRows       []CSVRow      `json:"mappingRows"`
}

type CSVRow struct {
	LineNumber int      `json:"line_number"`
	Columns    []string `json:"columns"`
}

type BlankColumn struct {
	ColumnName string `json:"column_name"`
	Line       int    `json:"line"`
}

func toJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}
	return string(b)
}

func main() {

	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}

// read https://pkg.go.dev/syscall/js#FuncOf for more information about js.FuncOf
func registerCallbacks() {
	js.Global().Set("getCSVReport", js.FuncOf(getCSVReport))
}

func getCSVReport(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return "No CSV content provided."
	}
	errors := []string{}
	blankColumns := []BlankColumn{}
	headers := []string{}
	mappingRows := []CSVRow{}
	blankColumnsCount := 0
	csvContent := args[0].String()
	if strings.TrimSpace(csvContent) == "" {
		errors = append(errors, "CSV is empty.")
	}
	if !strings.Contains(csvContent, ",") {
		errors = append(errors, "CSV does not appear to have any columns")
	}

	if len(errors) < 1 {
		scanner := bufio.NewScanner(strings.NewReader(csvContent))
		lineNumber := 0
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if lineNumber == 0 {
				// First line is the header
				headers = strings.Split(line, ",")
				lineNumber++
				continue
			}

			if line == "" {
				lineNumber++
				continue
			}

			columns := strings.Split(line, ",")
			hasBlank := false
			for i, column := range columns {
				if strings.TrimSpace(column) == "" {
					blankColumnsCount++
					hasBlank = true
					if i < len(headers) && len(blankColumns) < 10 {
						blankColumns = append(blankColumns, BlankColumn{
							ColumnName: headers[i],
							Line:       lineNumber + 1,
						})
					}
				}
			}

			if hasBlank && len(mappingRows) < 10 {
				mappingRows = append(mappingRows, CSVRow{
					LineNumber: lineNumber + 1,
					Columns:    columns,
				})
			}
			lineNumber++
		}
		if err := scanner.Err(); err != nil {
			errors = append(errors, err.Error())
		}

	}

	report := CSVReport{
		Valid:             true,
		Errors:            errors,
		Warning:           []string{},
		Info:              []string{},
		BlankColumnsCount: blankColumnsCount,
		BlankColumns:      blankColumns,
		Headers:           headers,
		MappingRows:       mappingRows,
	}
	return toJSON(report)
}
