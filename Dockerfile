FROM golang:1.20-alpine

# install git
RUN apk update && apk add --no-cache git

# where our file will be in the docker container
WORKDIR /app

# install air  which is used for hot reload each time a file is changed
RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod tidy

ENTRYPOINT  ["air", "-c", ".air.conf"]
