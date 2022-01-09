FROM golang:1.17 AS builder

WORKDIR /app
ADD . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o mggers-app .

FROM alpine:latest AS production
COPY --from=builder /app .

CMD ["./mggers-app"]
