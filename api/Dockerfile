
FROM golang:1.18-alpine

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . ./

ENV ADDR=0.0.0.0:4000

RUN go build -v -o server

ENTRYPOINT [ "/app/server" ]
