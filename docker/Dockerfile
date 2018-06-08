FROM golang

COPY entrypoint.sh /
RUN chmod +x /entrypoint.sh

RUN mkdir -p /go/src/bots/palmirinha/app
WORKDIR /go/src/bots/palmirinha/app

ENTRYPOINT [ "/entrypoint.sh" ]

EXPOSE 8080
