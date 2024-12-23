FROM golang:1.23-alpine as build

WORKDIR /code

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /code/cmd/consumer

# 実行コマンドを指定
CMD ["go", "run", "main.go"]
