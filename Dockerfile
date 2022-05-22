# syntax=docker/dockerfile:1

# BUILD STAGE
FROM golang:1.18.2-alpine3.15 AS builder
ENV GOOS linux
ENV CGO_ENABLED 0
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/main.go

# RUN STAGE
FROM alpine:3.15 AS prod
WORKDIR /app
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/main .

EXPOSE 8080
CMD [ "/app/main"]
