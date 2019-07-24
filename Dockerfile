FROM golang:latest

WORKDIR $GOPATH/src/github.com/alonaprylypa/Project
RUN curl https://glide.sh/get | sh
ENV SERVICE_PORT 8080

EXPOSE $SERVICE_PORT

ADD . .
RUN glide install
RUN go build ./cmd/main/main.go

CMD ["./main"]