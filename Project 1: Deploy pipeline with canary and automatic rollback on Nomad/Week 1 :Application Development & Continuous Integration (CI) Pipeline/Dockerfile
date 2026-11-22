
# Stage 1: Build
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copy project files
COPY . .

# Compile Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o exe_binary .

# Stage 2: Runtime
FROM alpine:3.22

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/exe_binary .

EXPOSE 8080

CMD ["./exe_binary"]
