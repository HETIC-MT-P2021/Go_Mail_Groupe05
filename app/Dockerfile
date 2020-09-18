FROM golang:alpine

RUN apk add --no-cache git
RUN mkdir /app

ADD . /app
WORKDIR /app

RUN ls

RUN go mod tidy
RUN go build -o main .

CMD ["./main"]
