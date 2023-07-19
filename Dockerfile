FROM golang:1.20-alpine

# install git
RUN apk update && apk add --no-cache git

# where our file will be in the docker container
WORKDIR /app

# copy the source from the current directory to the working directory inside the container
COPY . .

RUN go mod tidy

# install CompileDaemon which is used for hot reload each time a file is changed
RUN go build -o binary

ENTRYPOINT [ "/app/binary" ]
