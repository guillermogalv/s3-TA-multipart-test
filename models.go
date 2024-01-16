package main

type Config struct {
	XApiKey string
	BaseURL string
}

type InitializePayload struct {
	Name string `json:"name"`
}

type ResponseInitialize struct {
	FileID  string `json:"fileId"`
	FileKey string `json:"fileKey"`
}

// Define a struct that matches the expected request structure
type RequestPayload struct {
	FileKey string `json:"fileKey"`
	FileId  string `json:"fileId"`
	Parts   int    `json:"parts"`
}

// Define a struct for the response
type ResponsePayload struct {
	Parts []struct {
		SignedUrl  string `json:"signedUrl"`
		PartNumber int    `json:"PartNumber"`
	} `json:"parts"`
}

// PartUploadURL represents a pre-signed URL for a part upload
type PartUploadURL struct {
	SignedUrl  string
	PartNumber int
}

// UploadResponse represents the response after uploading a part
type UploadResponse struct {
	ETag       string
	PartNumber int
}
