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
ENV BUCKET_URL="https://pub-d3572231ffe04037aa4408917259b33d.r2.dev"

# アプリケーションを実行
CMD ["./main"]
