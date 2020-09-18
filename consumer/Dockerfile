FROM golang:alpine

RUN apk add --no-cache git
RUN mkdir /mail-consumer

ADD . /mail-consumer
WORKDIR /mail-consumer

RUN go get github.com/streadway/amqp
RUN go get github.com/joho/godotenv
RUN go get gopkg.in/gomail.v2
RUN go get github.com/HETIC-MT-P2021/Go_Mail_Groupe05/consumer/mailing

RUN go build -o main .

CMD ["./main"]