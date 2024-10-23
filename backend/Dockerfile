# BUILD STAGE
FROM golang:1.22.2 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/bin/main /app

# RUN STAGE
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/main .

ENTRYPOINT [ "/main" ]