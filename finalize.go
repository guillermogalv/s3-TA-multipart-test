package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Define the structure for your request payload
type CompleteMultipartUploadRequest struct {
	FileKey string `json:"fileKey"`
	FileId  string `json:"fileId"`
	Parts   []Part `json:"parts"`
}

// Part structure as expected by your Lambda function
type Part struct {
	ETag       string `json:"ETag"`
	PartNumber int    `json:"PartNumber"`
}

func FinalizeUpload(cfg Config, fileid string, ur []UploadResponse) {
	// Sample data for fileKey, fileId and parts
	fileKey := "testUpload/image.png"
	fileId := fileid
	parts := []Part{
		{ETag: ur[0].ETag, PartNumber: ur[0].PartNumber},
		{ETag: ur[1].ETag, PartNumber: ur[1].PartNumber},
		// Add more parts as necessary
	}

	// Create the request payload
	requestPayload := CompleteMultipartUploadRequest{
		FileKey: fileKey,
		FileId:  fileId,
		Parts:   parts,
	}

	// Marshal the request payload to JSON
	jsonData, err := json.Marshal(requestPayload)
	if err != nil {
		fmt.Println("Error marshalling request data:", err)
		return
	}

	// Define the API endpoint
	apiEndpoint := cfg.BaseURL + "/finalize" // Replace with your actual API endpoint

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", cfg.XApiKey)

	// Send the request using an HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Check and print the response status
	fmt.Println("Response status:", resp.Status)
}
