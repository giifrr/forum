FROM golang:1.20-alpine as builder

# install git
RUN apk update && apk add --no-cache git

# where our file will be in the docker container
WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o binary

FROM alpine:latest as baseImage

WORKDIR /

COPY --from=builder /app .

ENTRYPOINT ["./binary"]
