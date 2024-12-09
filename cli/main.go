package main

import (
	"bytes"
	"calc/internal/model"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Define command-line flags
	helpFlag := flag.Bool("help", false, "Show help information")
	typeFlag := flag.String("type", "", "The type of operation (errors, evaluate, validate). If errors is selected [expression] should be omitted.")
	expressionFlag := flag.String("expression", "", "The math expression to evaluate or validate")
	flag.Parse()

	if *helpFlag {
		printHelp()
		return
	}

	if *typeFlag == "" {
		fmt.Println("use -help to see the manual")
		return
	}

	if *expressionFlag == "" && (*typeFlag == "evaluate" || *typeFlag == "validate") {
		fmt.Println("Please provide a value for the 'expression' flag")
		return
	}

	var endpoint string
	var method string
	var body []byte

	switch *typeFlag {
	case "errors":
		endpoint = "http://localhost:8080/errors"
		method= http.MethodGet
	case "evaluate":
		endpoint = "http://localhost:8080/evaluate"
		body, _ = json.Marshal(model.EvaluateRequest{Expression: *expressionFlag})
		method= http.MethodPost
	case "validate":
		endpoint = "http://localhost:8080/validate"
		body, _ = json.Marshal(model.ValidateRequest{Expression: *expressionFlag})
		method= http.MethodPost

	default:
		fmt.Println("Invalid type argument. Supported types: errors, evaluate, validate")
		return
	}

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Failed to prepare request:", err)
		return
	}

	req.Header.Add("content-type","application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	if resp.StatusCode >= 400 {
		var errResp model.ErrorCLI
		if err := json.Unmarshal(respBody, &errResp); err != nil {
			fmt.Println("Failed to parse error response:", err)
			return
		}
		fmt.Println("Error:", errResp.Error)
	} else {
		if *typeFlag == "evaluate" {
			var resultResp model.EvaluateResponse
			if err := json.Unmarshal(respBody, &resultResp); err != nil {
				fmt.Println("Failed to parse result response:", err)
				return
			}
			fmt.Println("Result:", resultResp.Result)
			return
		}

		var validationResp model.ValidateResponse
		if err := json.Unmarshal(respBody, &validationResp); err != nil {
			fmt.Println("Failed to parse result response:", err)
			return
		}
		fmt.Println("Validation:", validationResp.Valid)
		return

	}
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  cliapp [type] [expression]")
	fmt.Println("Arguments:")
	fmt.Println("  type        : The type of operation (errors, evaluate, validate). If errors is selected [expression] should be omitted.")
	fmt.Println("  expression  : The math expression to evaluate or validate")
	fmt.Println("Flags:")
	flag.PrintDefaults()
}