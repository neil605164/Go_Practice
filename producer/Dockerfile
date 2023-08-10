FROM golang:1.20-alpine

# 安裝 git
# go install air
RUN apk add git \
    && apk add build-base \
    && go install github.com/cosmtrek/air@latest
