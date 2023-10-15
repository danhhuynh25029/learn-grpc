FROM golang:1.21

COPY go.mod go.sum /app/

WORKDIR /app

RUN go mod download

ADD . .
RUN go build -o /usr/local/bin/app

CMD ["app"]