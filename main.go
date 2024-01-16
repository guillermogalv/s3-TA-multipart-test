package main

func main() {
	cfg := Config{
		XApiKey: "<here-u-x-api-key>",
		BaseURL: "stage/enpoint",
	}
	init := Initialize(cfg)
	urls := GetSignedURLs(cfg, init.FileID)
	uploaded := UploadPNGFile(urls)
	FinalizeUpload(cfg, init.FileID, uploaded)
}
