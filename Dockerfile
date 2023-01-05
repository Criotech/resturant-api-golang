FROM golang:alpine

COPY . /app
WORKDIR /app

EXPOSE 9000

RUN CGO_ENABLED=0 GOOS=linux go build -o main

CMD ["./main"]