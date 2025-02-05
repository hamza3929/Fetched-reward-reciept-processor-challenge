# Fetched-reward-reciept-processor-challenge
This project is a take home exercise for the company fetched where given a receipt calculate the amount of points the receipt earns and build a webservice that generates an ID for each receipt and stores the calculates points for the receipt

## Features
-Submit receipts via a POST /receipts/process endpoint.
-Retrieve points for a submitted receipt via a GET /receipts/{id}/points endpoint.
-In-memory storage (no database required).
-Docker support for easy deployment.

## Prerequisites
-[Go 1.19+](https://go.dev/doc/install)
-[Docker](https://docs.docker.com/desktop/)

## Installation & Setup
### Running Locally
1. Clone the repository:
   ```
   git clone <repository_url>
   cd <repository_directory>
   ```
2. Install dependencies and build the application:
   ```
   go mod tidy
   go build -o receipt-processor
   ```
3. Run the application:
   ```
   ./receipt-processor
   ```
4. The service will start on http://localhost:8080.

### Running with Docker
1. Build the Docker image:
   ```
   docker build -t receipt-processor .
   ```
2. Run the container:
   ```
   docker run -p 8080:8080 receipt-processor
   ```
3. The service will be accessible at http://localhost:8080.

## API Usage
# rest of readme to be finished











