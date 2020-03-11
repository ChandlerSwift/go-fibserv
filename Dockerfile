FROM golang:1.14

WORKDIR /usr/src/app

COPY fibserv.go .

CMD ["go", "run", "fibserv.go"]
