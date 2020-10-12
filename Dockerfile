FROM golang:latest AS builder
WORKDIR /go/src/github.com/tonyskapunk/heyo/
COPY main.go .
RUN GOOS=linux CGO_ENABLED=0 go build -a -o heyo .

FROM alpine:latest
RUN addgroup -S app \
    && adduser -S -g app app \
    && apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /go/src/github.com/tonyskapunk/heyo/heyo .
RUN chown -R app:app ./
USER app
CMD ["./heyo"]