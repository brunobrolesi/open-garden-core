FROM golang:1.21-alpine as builder

WORKDIR /usr/src/app

COPY . .

RUN go build -C ./cmd -buildvcs=false -o bin/open_garden_core

FROM alpine:3.17

RUN apk update

ENV WAIT_FOR_VERSION v2.2.4

RUN wget -q -O /usr/bin/wait-for https://raw.githubusercontent.com/eficode/wait-for/$WAIT_FOR_VERSION/wait-for && \
    chmod +x /usr/bin/wait-for  && \
    apk add --update --no-cache netcat-openbsd

EXPOSE 8080

COPY --from=builder /usr/src/app/cmd/bin .
COPY --from=builder /usr/src/app/.env .
