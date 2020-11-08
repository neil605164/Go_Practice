# 第一層基底
FROM golang:1.14.0-alpine

# 安裝 git
# go get fresh
RUN apk add git \
    && go get github.com/pilu/fresh