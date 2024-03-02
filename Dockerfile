FROM golang:1.18
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
VOLUME [ "/app/log.json" ]
CMD ["/app/main"]