FROM golang:1.24 as builder
WORKDIR /usr/src/app
COPY . .
RUN go mod download
RUN go build ./cmd/genshin-artifact-db

FROM golang:1.24
COPY --from=builder /usr/src/app/genshin-artifact-db /usr/local/bin

# データ永続化用ディレクトリを作成
RUN mkdir -p /var/lib/genshin-artifact-db
RUN mkdir -p /etc/config/genshin-artifact-db

EXPOSE 8080
CMD ["genshin-artifact-db"]
