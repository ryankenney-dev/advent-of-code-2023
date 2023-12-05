# Start from the official Golang image to build the binary.
FROM golang:1.18 as builder

# Set the Current Working Directory inside the container.
WORKDIR /app

# Initialize Go module and download dependencies.
RUN go mod init github.com/ryankenney-dev/advent-of-code-2023
RUN go mod tidy

# Copy the source from the current directory to the Working Directory inside the container.
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch for the final image.
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]


