FROM golang:1.18.3-alpine3.16 as builder

COPY ./controllers /go/src/github.com/vagnerpelais/napp-agenda/controllers/
COPY ./database /go/src/github.com/vagnerpelais/napp-agenda/database/
COPY ./models /go/src/github.com/vagnerpelais/napp-agenda/models/
COPY ./config /go/src/github.com/vagnerpelais/napp-agenda/config/
COPY ./repositories /go/src/github.com/vagnerpelais/napp-agenda/repositories/
COPY ./services /go/src/github.com/vagnerpelais/napp-agenda/services/
COPY ./server /go/src/github.com/vagnerpelais/napp-agenda/server/


COPY ./go.mod /go/src/github.com/vagnerpelais/napp-agenda/
COPY ./main.go /go/src/github.com/vagnerpelais/napp-agenda/


WORKDIR /go/src/github.com/vagnerpelais/napp-agenda

RUN go get

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/main github.com/vagnerpelais/napp-agenda

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/vagnerpelais/napp-agenda/build/main /usr/bin/main
EXPOSE 5050 5050

RUN apk add --update curl && apk add --update tar && apk add --no-cache tzdata && rm -rf /var/cache/apk/*

RUN rm -f /etc/localtime; ln -s /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime;

ENTRYPOINT ["/usr/bin/main"]