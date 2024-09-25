FROM golang:1.22.7-alpine3.20 AS builder
WORKDIR /app
COPY cmd cmd
COPY pkg pkg
COPY go.mod go.mod
COPY go.sum go.sum
RUN go build -o main cmd/lambda/main.go

FROM alpine:3.20
COPY --from=builder /app/main /main
ENTRYPOINT [ "/main" ]