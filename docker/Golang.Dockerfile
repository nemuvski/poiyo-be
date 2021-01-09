FROM golang:1.15-alpine

# 必要なソフトウェアをインストール
RUN apk add --update --no-cache git

# コンテナの作業ディレクトリを変更
RUN mkdir /app
WORKDIR /app
ADD src ./

# Golangのパッケージのインストール
RUN go get github.com/lib/pq
