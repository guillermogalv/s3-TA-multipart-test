package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"os"
	"sync"
)

func splitImage(img image.Image, numParts int) []*image.RGBA {
	// Get the dimensions of the original image
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	// Calculate the height of each part
	partHeight := height / numParts

	// Initialize an array to store the image parts
	imageParts := make([]*image.RGBA, numParts)

	// Split the image into parts
	for i := 0; i < numParts; i++ {
		// Calculate the Y coordinates for this part
		startY := i * partHeight
		endY := startY + partHeight

		// Create a new image for this part
		partImg := image.NewRGBA(image.Rect(0, 0, width, partHeight))

		// Copy the pixels from the original image to the part image
		for y := startY; y < endY; y++ {
			for x := 0; x < width; x++ {
				// Get the pixel color from the original image
				originalColor := img.At(x, y)

				// Set the pixel color in the part image
				partImg.Set(x, y-startY, originalColor)
			}
		}

		// Store the part image in the array
		imageParts[i] = partImg
	}

	return imageParts
}

// Function to upload a part to a pre-signed URL and return the ETag
func uploadPart(part *image.RGBA, partURL PartUploadURL, wg *sync.WaitGroup, ch chan<- UploadResponse) {
	defer wg.Done()

	// Encode the image part to a buffer
	buf := new(bytes.Buffer)
	err := png.Encode(buf, part)
	if err != nil {
		panic("Failed to encode image part: " + err.Error())
	}

	// Create a new request to upload this part
	req, err := http.NewRequest("PUT", partURL.SignedUrl, buf)
	if err != nil {
		panic("Failed to create request: " + err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("Failed to do request: " + err.Error())
	}
	defer resp.Body.Close()

	// Check response status code for success
	if resp.StatusCode != http.StatusOK {
		panic("Failed to upload part, status code: " + resp.Status)
	}

	// Extract the ETag from the response header
	etag := resp.Header.Get("ETag")
	fmt.Printf("Part %d uploaded successfully, ETag: %s\n", partURL.PartNumber, etag)

	// Send the ETag and PartNumber back through the channel
	ch <- UploadResponse{ETag: etag, PartNumber: partURL.PartNumber}
}

func UploadPNGFile(urls ResponsePayload) []UploadResponse {
	// Open the PNG image file
	imgFile, err := os.Open("./testimage.png") // Ensure this path is correct
	if err != nil {
		panic("Failed to open image file: " + err.Error())
	}
	defer imgFile.Close()

	// Decode the PNG image
	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic("Failed to decode PNG image: " + err.Error())
	}

	// Number of parts to split the image into
	numParts := 2

	// Split the image into parts
	imageParts := splitImage(img, numParts)

	if len(imageParts) != numParts {
		panic("Failed to split image into parts")
	}

	// Pre-signed URLs for uploading the parts
	preSignedURLs := []PartUploadURL{
		{SignedUrl: urls.Parts[0].SignedUrl, PartNumber: urls.Parts[0].PartNumber},
		{SignedUrl: urls.Parts[1].SignedUrl, PartNumber: urls.Parts[1].PartNumber},
	}

	var wg sync.WaitGroup

	// Create a channel to receive upload responses
	responseChannel := make(chan UploadResponse, numParts)
	defer close(responseChannel)

	// Upload each part concurrently using Goroutines
	for i, part := range imageParts {
		wg.Add(1)
		go uploadPart(part, preSignedURLs[i], &wg, responseChannel)
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	// Collect responses
	var responses []UploadResponse
	for i := 0; i < numParts; i++ {
		responses = append(responses, <-responseChannel)
	}

	return responses
}
