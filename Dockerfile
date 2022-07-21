FROM golang:1.18.3-alpine3.16 as builder

COPY ./controllers /go/src/napp-agenda/controllers/
COPY ./server /go/src/napp-agenda/server/
COPY ./database /go/src/napp-agenda/database/
COPY ./config /go/src/napp-agenda/config/
COPY ./models /go/src/napp-agenda/models/

COPY ../go.mod /go/src/napp-agenda/
COPY ../main.go /go/src/napp-agenda/


WORKDIR /go/src/napp-agenda/

RUN go get

RUN CGO_ENABLED=0 GOOS=linux go build -o build/main .

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/napp-agenda/build/main /usr/bin/main
EXPOSE 5050 5050

RUN apk add --update curl && apk add --update tar && apk add --no-cache tzdata && rm -rf /var/cache/apk/*

RUN rm -f /etc/localtime; ln -s /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime;

ENTRYPOINT ["/usr/bin/main"]