# Stage 1: Build the application
FROM golang:1.20.12-alpine AS builder

WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Build the Go application
RUN go build -o xdev

# Stage 2: Create a minimal runtime image
FROM alpine:3.19.0

ENV PORT=8000

WORKDIR /app

# Copy the built application from the previous stage
COPY --from=builder /app/xdev .
COPY entrypoint.sh .
# Expose the port on which the application will run
EXPOSE ${PORT}
# Start the application
ENTRYPOINT ["./entrypoint.sh"]