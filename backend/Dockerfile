# syntax=docker/dockerfile:1

# ベースイメージとして公式のGoイメージを使用
FROM golang:1.20.4-alpine as builder

# 作業ディレクトリを設定
WORKDIR /backend

# 必要なパッケージをインストール
RUN apk add --no-cache git

# ソースコードをコピー
COPY . .

# Goの依存関係をダウンロード
RUN go mod download
RUN go mod tidy
RUN go mod verify

# ビルド
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 新しいステージを作成し、実行可能ファイルをコピー
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /backend/main .

# Application
ENV APP_PORT="80"
# CDN
ENV STATIC_FILE_PATH="https://pub-f4af92d85ea3486a87712cb9821633f6.r2.dev/views"
ENV BACKEND_URL="https://novelchain.azurewebsites.net"

# アプリケーションを実行
CMD ["./main"]
