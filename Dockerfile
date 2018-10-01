FROM golang:alpine AS builder

COPY . /go/src/github.com/zhangzhoujian/echo

RUN go install ./...

FROM alpine

ENV GOPATH /go

COPY --from=builder /go/bin/echo /go/bin/echo

CMD [ "/go/bin/echo" ]

EXPOSE 8080
