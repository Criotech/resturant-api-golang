
# Imports a GO alpine image
FROM golang:1.18-alpine3.16

# Creates the application's directory
RUN mkdir -p /app

# Sets the work directory to application's folder
WORKDIR /app

# Copy files into application's folder
COPY .  .

# Install the dependencies
RUN go mod download

RUN go build -o main .

CMD ["/app/main"]
