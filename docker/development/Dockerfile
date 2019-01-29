FROM golang:latest

ENV TZ=America/Sao_Paulo
RUN echo "America/Sao_Paulo" > /etc/timezone

RUN go get -u github.com/pilu/fresh
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/bwmarrin/discordgo

ENV GOCACHE=off

WORKDIR /go/src/github.com/palmirinha/app

EXPOSE 8080

ENTRYPOINT ["fresh"]