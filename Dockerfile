FROM golang:1.22-alpine AS builder

WORKDIR /app
RUN apk add --no-cache make git

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build


FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/bin/simple-env-parser .
COPY .env ./.env

RUN chmod +x /app/simple-env-parser

ENTRYPOINT ["/app/simple-env-parser"]