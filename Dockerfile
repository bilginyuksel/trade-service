# build phase
FROM golang:1.17

WORKDIR /app
COPY . /app

ENV CGO_ENABLED=0

RUN go build -o trade-service ./cmd

# execution phase
FROM alpine:latest

WORKDIR /

COPY --from=0 /app/trade-service ./
COPY --from=0 /app/.config ./.config

CMD ["./trade-service"]