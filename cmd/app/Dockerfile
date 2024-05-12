FROM golang:1.22.2-alpine AS builder

WORKDIR /go/src/github.com/go-gin-intikom

COPY . ./

RUN go build -o bin/app cmd/app/app.go


FROM alpine
RUN     apk add --no-cache \
        ca-certificates

COPY ./configs/etc /etc
COPY --from=builder /go/src/github.com/go-gin-intikom/bin/app ./
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
ENV TZ=Asia/Jakarta
ENV ZONEINFO=/zoneinfo.zip
EXPOSE 8080
ENTRYPOINT ["./app"]