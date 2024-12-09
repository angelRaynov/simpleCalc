# Use a Golang base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

ENV API_PORT=8080

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN go build -o calc .

# Expose the port that your server listens on
EXPOSE 8080

# Set the entry point to run the Lime API server
CMD ["./calc"]

#docker run calc