FROM golang:1.7

WORKDIR /go/src/github.com/HETIC-MT-P2021/Go_Mail_Groupe05

COPY main.go main.go

RUN go get github.com/streadway/amqp/...
RUN go build -o main .
RUN ls

CMD ./main