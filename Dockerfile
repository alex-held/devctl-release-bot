FROM golang as builder

WORKDIR /go/src/github.com/alex-held/devctl-release-bot
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go test -mod vendor ./... -cover
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor --ldflags "-s -w" -o devctl-release-bot cmd/action/*

FROM alpine

RUN mkdir -p /home/app

# Add non root user
RUN addgroup -S app && adduser app -S -G app
RUN chown app /home/app

WORKDIR /home/app

USER app

COPY --from=builder /go/src/github.com/alex-held/devctl-release-bot/devctl-release-bot /usr/local/bin/

CMD ["devctl-release-bot", "action"]
