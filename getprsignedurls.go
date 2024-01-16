package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetSignedURLs(cfg Config, fileId string) ResponsePayload {
	var responsePayload ResponsePayload

	// Set up the request data
	requestData := RequestPayload{
		FileKey: "testUpload/image.png",
		FileId:  fileId,
		Parts:   2,
	}

	// Marshal the data into a JSON payload
	payload, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error marshalling request data:", err)
		return responsePayload
	}

	// Define the URL of your Lambda function (replace with your actual endpoint)
	url := "https://j6sfpn2hdj.execute-api.us-east-1.amazonaws.com/v1/pre-signed-urls"

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return responsePayload
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", "p7OA8YZTw29D6o9DU5btJ6Dpyj0md0uE6njLq4lf") // Replace with your API key

	// Initialize an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return responsePayload
	}
	defer resp.Body.Close()

	// Read and parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return responsePayload
	}

	if err := json.Unmarshal(body, &responsePayload); err != nil {
		fmt.Println("Error parsing response JSON:", err)
		return responsePayload
	}

	// Print the response
	fmt.Printf("Received response: %+v\n", responsePayload)
	return responsePayload
}
