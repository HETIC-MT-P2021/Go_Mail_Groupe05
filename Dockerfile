FROM golang:alpine

RUN apk add --no-cache git
RUN mkdir /app

ADD . /app
WORKDIR /app

RUN go get
RUN go mod vendor
RUN go build -o main .

CMD ["./main"]
