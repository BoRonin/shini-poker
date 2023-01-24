FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

CMD ["go", "run", "cmd/api/main.go"]



