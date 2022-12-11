FROM golang:1.18-alpine3.15

WORKDIR /usr/local/src/go01
COPY main.go /usr/local/src/go01/

EXPOSE 8888

CMD [ "go", "run", "main.go" ]