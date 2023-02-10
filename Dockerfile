FROM golang:1.19 as builder

ENV GO111MODULE=on
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .
EXPOSE 8080
CMD ["/app/main"]
