# Overview

This Go program is designed to facilitate the uploading of files to Amazon S3 using multipart upload and transfer acceleration features. The program initializes a configuration, retrieves signed URLs for the file to be uploaded, performs the file upload, and finalizes the upload process.


# Getting Started
## Prerequisites
- Go programming language installed
- AWS account with S3 access

## Configuration
Before running the program, you must configure it with your AWS API key and base URL. These can be obtained from the AWS Console.

## API Key and Base URL
```
cfg := Config{
    XApiKey: "<YOUR_API_KEY>",
    BaseURL: "<YOUR_BASE_URL>",
}

```

## Running the Program
Execute the program with the following command:

```
go run .

```

## Features
- Multipart Upload: Efficiently uploads large files in parts.
- Transfer Acceleration: Speeds up the transfer of files over long distances.

## More Information
- For detailed information on multipart upload and transfer acceleration, visit the [AWS blog post](https://aws.amazon.com/blogs/compute/uploading-large-objects-to-amazon-s3-using-multipart-upload-and-transfer-acceleration/).
- Understand the limits and capabilities of this process in the [Amazon S3 documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/qfacts.html).

** Note: ** Replace <YOUR_API_KEY> and <YOUR_BASE_URL> with your actual AWS credentials. Make sure to keep your API key secure and do not share it publicly.