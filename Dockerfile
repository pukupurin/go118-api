FROM golang:1.19.2-alpine3.16 AS debug

# タイムゾーン設定
RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

WORKDIR /root/api

COPY ./go.* /root/api/
RUN go mod download
RUN go install github.com/cosmtrek/air@latest

COPY . /root/api

CMD ["tail", "-f", "/dev/null"]

# build continer
FROM golang:1.19.2-alpine3.16 AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /tmp/app

COPY go.* ./
RUN go mod download

COPY . ./
RUN go build -o ./build/go-ent -ldflags="-w -s"


# runtime continer
FROM alpine:3.16.2

RUN apk --update --no-cache add tzdata curl && \
    rm -rf /var/cache/apk/*

ENV TZ Asia/Tokyo

WORKDIR /root

COPY --from=builder /tmp/app/build/go-ent /root/go-ent

EXPOSE 80
CMD ["/root/go-ent"]
