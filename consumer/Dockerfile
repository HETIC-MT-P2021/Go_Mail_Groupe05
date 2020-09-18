FROM golang:alpine

RUN apk add --no-cache git
RUN mkdir /consumer

ADD . /consumer
WORKDIR /consumer

RUN ls

RUN go mod tidy
RUN go build -o main .

CMD ["./main"]