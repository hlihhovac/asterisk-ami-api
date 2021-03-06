FROM golang:1.8-alpine

WORKDIR /go/src/github.com/hlihhovac/asterisk-ami-api
COPY . .
RUN apk add --no-cache git && go get github.com/Masterminds/glide
RUN glide i
RUN go build -o asterisk-ami-api main.go
RUN rm -rf vendor

EXPOSE 3000

CMD ["./asterisk-ami-api", "-conf", "api.conf"]
