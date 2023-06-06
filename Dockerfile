# Stage 1: Build the Go binary
FROM golang:1.17 AS builder

WORKDIR /app

COPY . .

RUN go mod download

WORKDIR /app/cmd/there-yet/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o thereyet .

# Stage 2: Create the final minimal image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/cmd/there-yet/thereyet .

EXPOSE 8080

CMD ["./thereyet"]
