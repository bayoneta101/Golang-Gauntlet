FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ARG SERVICE
COPY ${SERVICE}/ ./${SERVICE}/

RUN go build -o /bin/service ./${SERVICE}/main/

FROM alpine:latest

COPY --from=builder /bin/service /service

ENTRYPOINT ["/service"]
