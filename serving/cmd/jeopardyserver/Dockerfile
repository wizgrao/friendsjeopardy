FROM golang:1.7.3
WORKDIR /go/src/github.com/wizgrao/friendsjeopardy/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app github.com/wizgrao/friendsjeopardy/serving/cmd/jeopardyserver

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /go/src/github.com/wizgrao/friendsjeopardy/app .
CMD ["./app"]
