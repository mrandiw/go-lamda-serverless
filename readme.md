# Serverless Lambda API

A Go-based REST API built with the Echo framework that can run both locally and as an AWS Lambda function using Apex Gateway.

## Features

- Dual-mode execution: runs as a standard HTTP server locally or as a Lambda function in AWS
- Easy deployment to AWS Lambda with provided Makefile
- Simple API with demonstration endpoints
- Environment-based configuration
- Docker support for local development and Lambda builds

## Prerequisites

- Go 1.19 or newer
- Make (for using the Makefile)
- AWS CLI (for deploying to Lambda)

## Installation

1. Clone this repository:

```bash
git clone <repository-url>
cd echo-lambda-api
```

## Project Structure

```
.
├── main.go            # Main application code
├── Makefile           # Build and deployment automation
├── go.mod             # Go module definition
├── go.sum             # Go module checksums
└── README.md          # This file
```

## Running Locally

To run the application in local mode:

```bash
make run
```

This will start the server on port 8080 (or the port specified in the PORT environment variable).

You can also build the binary and run it directly:

```bash
make build
```

## API Endpoints

The application provides the following endpoints:

- `GET /` - Returns a "Hello, World!" message
- `GET /get-student` - Returns student API test message
- `GET /get-grade` - Returns grade API test message
- `GET /get-teacher` - Returns teacher API test message