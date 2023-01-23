FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

CMD ["go", "run", "cmd/api/main.go"]
# RUN CGO_ENABLED=0 go build -o backendApp ./cmd/api

# RUN chmod +x /app/backendApp

#small docker
# FROM alpine:latest
#
# RUN mkdir /app
#
# COPY --from=builder /app/backendApp /app
#
# CMD ["/app/backendApp"]



