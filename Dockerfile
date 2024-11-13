FROM golang:1.23-alpine3.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./bin/capital-gains ./cmd/cli/capital-gains/main.go

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/bin/capital-gains /app/bin/capital-gains
ENTRYPOINT ["/app/bin/capital-gains"]

