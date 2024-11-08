# goのバージョン
FROM golang:1.23.2-alpine

# OSのインストール済みのパッケージをバージョンアップ&必要なパッケージをインストール
RUN apk upgrade --update && \
    apk --no-cache add git gcc musl-dev \
    chromium \
    fontconfig \
    font-noto-cjk \
    font-noto-cjk-extra 

# フォントの設定
RUN fc-cache -fv

# ワーキングディレクトリの設定
WORKDIR /go/src/app/

# ホストPCの[./src]フォルダーを[/go/src/app/]フォルダーにコピー
ADD ./mr-platform-contents-server/src .

# アプリケーションをビルド
RUN go build -o /go/bin/app .

# コンテナ起動時に実行するコマンド
CMD ["/go/bin/app"]
