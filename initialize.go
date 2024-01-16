package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Initialize(cfg Config) ResponseInitialize {
	// Define the URL
	url := cfg.BaseURL + "/initialize"
	var response ResponseInitialize

	data := InitializePayload{
		Name: "testUpload/image.png", // replace with the actual file name
	}

	// Create a POST request
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
		return response
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", cfg.XApiKey)

	// Create a client and execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request to API endpoint. %s\n", err)
		return response
	}

	// Close the response body when the function returns
	defer resp.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body. %s\n", err)
		return response
	}

	fmt.Printf("Response Body: %s\n", string(body)) // this will print: {"fileId":"z392I9jNrWfipV5nd4aR_1lohsAe6q963.Jx.HpM0s3qPfRSfaidm7iyHeqXHq1nQg_e8JfoOvk0MZle_b4bUwUhExcp98C8wpPp_Z2dslEXB4qK2UHEtPXrtpxZqfVJtWbSza1mi16n0fm3KwJFzA--","fileKey":"cotiviti.png"}

	err = json.Unmarshal([]byte(string(body)), &response)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %s\n", err)
		return response
	}
	return response
}
