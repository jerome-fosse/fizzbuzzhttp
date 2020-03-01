### Step 1 : Build the executable
FROM golang:1.13 AS builder

ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /go/src/fizzbuzzhttp
COPY . .

RUN go get -d -v ./...
RUN make install

### Step 2 : Build a smaller image
FROM alpine:latest

COPY  --from=builder  /go/bin/fizzbuzzhttp /go/bin/fizzbuzzhttp
RUN chmod +x /go/bin/fizzbuzzhttp

EXPOSE 8080

CMD ["/go/bin/fizzbuzzhttp"]

