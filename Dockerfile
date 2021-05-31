# Multi-Stage Dockerfile

FROM golang:1.16.4 AS builder

# We create an /app directory in which we'll put all of our project code
RUN mkdir /app
ADD . /app
WORKDIR /app

# Build applications to executable binary files
RUN CGO_ENABLED=0 GOOS=linux go build -o go-gin-currency

# Run Unit Testing
RUN go test ./...

# Use alpine to run the application
FROM alpine:latest AS production

# Copy output from builder stage to production stage
COPY --from=builder /app .

# Binary application files ready to go
CMD ["./go-gin-currency"]