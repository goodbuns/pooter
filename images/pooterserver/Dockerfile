FROM golang:1.14beta1-alpine3.11 AS builder

RUN apk update && apk add --no-cache git
WORKDIR $HOME/src
COPY . .

RUN go get -d -v ./...
RUN go build -o /go/bin/pooterserver github.com/liftM/pooter/cmd/pooterserver

FROM scratch

COPY --from=builder /go/bin/pooterserver /go/bin/pooterserver
ENTRYPOINT ["/go/bin/pooterserver"]
