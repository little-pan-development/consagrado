#!/bin/bash
go get ./

if [ $APP_ENV == 'develop' ]; then
    go get github.com/pilu/fresh && fresh;
else
    go run main.go
fi
