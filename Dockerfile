FROM golang:1.17 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o mggers-app cmd/main.go

FROM alpine:latest AS launcher
COPY --from=builder /app .
CMD ["./mggers-app"]