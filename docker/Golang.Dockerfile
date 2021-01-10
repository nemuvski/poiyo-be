FROM golang:1.15-alpine

# 必要なソフトウェアをインストール
RUN apk add --update --no-cache git

# コンテナの作業ディレクトリを変更
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD ./ /go/src/app
