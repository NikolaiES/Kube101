# Stage 1: Build the Go application
FROM golang:1.22.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/

RUN go build -o goApp cmd/main.go

# Stage 2: Create a smaller image with only the built binary
FROM golang:1.22.3

WORKDIR /root/

COPY --from=builder /app/goApp .
COPY /files/templates/main.gohtml ./files/templates/main.gohtml

EXPOSE 8080

CMD ["./goApp"]

