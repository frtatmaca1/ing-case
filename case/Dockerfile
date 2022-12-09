FROM docker.io/golang:alpine as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV GOSUMDB=off

WORKDIR /app

COPY . .

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -a -v -o server ./cmd/server

WORKDIR /app

EXPOSE 8080

CMD ["/app/server"]

