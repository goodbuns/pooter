FROM golang:1.14beta1-alpine3.11 AS builder

RUN apk update && apk add --no-cache git
WORKDIR $HOME/src
COPY . .

RUN go get -d -v ./...
RUN go build -o /go/bin/twooterserver github.com/liftM/twooter/cmd/twooterserver

FROM scratch

COPY --from=builder /go/bin/twooterserver /go/bin/twooterserver
ENTRYPOINT ["/go/bin/twooterserver"]
