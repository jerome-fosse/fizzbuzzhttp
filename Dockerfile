FROM golang:latest

WORKDIR /go/src/fizzbuzzhttp
COPY . .

RUN go get -d -v ./...
RUN make install

EXPOSE 8080

CMD ["/go/bin/fizzbuzzhttp"]

