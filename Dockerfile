FROM golang:latest

ENV SERVICE_PORT 8080

EXPOSE $SERVICE_PORT

RUN mkdir app/

ADD . /app

WORKDIR /app

RUN go build ./cmd/main/main.go

CMD ["./main"]