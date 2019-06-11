FROM golang:1.12.4-alpine3.9
EXPOSE 80

RUN apk update && apk add gcc g++

WORKDIR /go/src/server

COPY . .

RUN go build

CMD ["./server"]