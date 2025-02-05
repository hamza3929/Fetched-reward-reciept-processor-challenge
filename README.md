# Fetched-reward-reciept-processor-challenge
This project is a take home exercise for the company fetched where given a receipt calculate the amount of points the receipt earns and build a webservice that generates an ID for each receipt and stores the calculated points for the receipt

## Features
- Submit receipts via a POST /receipts/process endpoint.
- Retrieve points for a submitted receipt via a GET /receipts/{id}/points endpoint.
- In-memory storage (no database required).
- Docker support for easy deployment.

## Prerequisites
- [Go 1.19+](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/desktop/)

## Installation & Setup
### Installing Linux
Before running any commands or going through any files make sure you have linux setup or a linux terminal as all the commands being run are linux based if you want to setup linux on your terminal there are these commands you follow to setup linux
```
wsl --install
```
after installing linux you can run this next command to start using linux on your terminal
```
bash
```
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
### Submit a Receipt
#### Endpoint:
```
POST /receipts/process
```
### Request Body:
```
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    { "shortDescription": "Mountain Dew 12PK", "price": "6.49" }
  ],
  "total": "6.49"
}
```

### Response:
```
{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
```

## Get Points for a Receipt
### Endpoint:
```
GET /receipts/{id}/points
```
### Response:
```
{ "points": 32 }
```
## Dockerfile Explanation
```
# Use the official Golang image as a build stage
FROM golang:1.19 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o receipt-processor .

# Use a compatible base image instead of Alpine (Alpine lacks some necessary libraries)
FROM debian:stable-slim

# Set working directory inside the container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/receipt-processor .

# Ensure the binary is executable
RUN chmod +x /root/receipt-processor

# Expose the port
EXPOSE 8080

# Start the application
CMD ["./receipt-processor"]
```
### Explanation:
- First stage: Uses a Golang image to build the application.
- Second stage: Uses a Debian-based image for runtime compatibility.
- WORKDIR: Sets the working directory inside the container.
- COPY go.mod go.sum & RUN go mod download: Ensures dependencies are downloaded before copying the source files.
- COPY --from=builder: Copies the compiled binary from the build stage.
- RUN chmod +x: Ensures the binary is executable.
- EXPOSE 8080: Exposes port 8080 for API access.
- CMD: Runs the compiled Go binary.

## Troubleshooting

If you encounter issues running the service:
1. Ensure Go and Docker are installed correctly.
2. If using Docker, try rebuilding the image with:
   ```
   docker build --no-cache -t receipt-processor .
   ```
3. Check if the binary exists inside the container:
   ```
   docker run --rm -it receipt-processor ls -lh /
   ```
4. Run the container in interactive mode for debugging:
   ```
   docker run --rm -it receipt-processor /bin/sh
   ```
## Testing 
### Submitting Receipt
1. Submitting JSON content, which then returns an ID for the submitted receipt.
   ```
   curl -X POST "http://localhost:8080/receipts/process" \
     -H "Content-Type: application/json" \
     -d '{
    "retailer": "Target",
    "purchaseDate": "2022-01-01",
    "purchaseTime": "13:01",
    "items": [
      {
        "shortDescription": "Mountain Dew 12PK",
        "price": "6.49"
      }
    ],
    "total": "6.49"
     }'
   ```
2. Submitting A JSON file, which then returns an ID for the submitted receipt.
   ```
   curl -X POST "http://localhost:8080/receipts/process" \
     -H "Content-Type: application/json" \
     -d @receipt.json
   ```
### Retreive Points
1. After Submitting the Receipt an ID returned for that receipt, with the given ID change the text id in the brackets to the given ID to get that receipt's points earned .
   ```
   curl -X GET "http://localhost:8080/receipts/{id}/points"

   ```





