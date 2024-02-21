FROM golang:1.18
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
VOLUME [ "/app/numbers.txt", "/app/log.json" ]
CMD ["/app/main"]