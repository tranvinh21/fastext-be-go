FROM golang:1.24.2-alpine as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

FROM alpine:latest as run

WORKDIR /app

COPY --from=build /app/main /app/main

COPY .env .

EXPOSE 3000

ENTRYPOINT ["/app/main"]