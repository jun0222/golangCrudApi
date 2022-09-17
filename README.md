# 目次

<!-- TOC -->

- [目次](#目次)
- [概要](#概要)
- [参考サイト](#参考サイト)
- [Go 言語のインストール](#go-言語のインストール)
- [コード実行例](#コード実行例)
- [テストの実行](#テストの実行)
  - [参考 URL](#参考-url)
- [エラーまとめ](#エラーまとめ)
- [API サンプル実行方法](#api-サンプル実行方法)
- [CRUD API 実行方法](#crud-api-実行方法)
  - [CREATE](#create)
  - [READ](#read)
  - [UPDATE](#update)
  - [DELETE](#delete)
- [ローカル環境 sequel pro 設定値](#ローカル環境-sequel-pro-設定値)
- [必要ライブラリインストールコマンド](#必要ライブラリインストールコマンド)

<!-- /TOC -->

# 概要

Go/Gin/Gorm/Docker/MySQL を使った CRUD API

# 参考サイト

- [フロントエンドエンジニアが Go 言語で ToDo リスト API を作ってみた](https://liginc.co.jp/584227)
- [テスト駆動開発で GO 言語を学びましょう](https://andmorefine.gitbook.io/learn-go-with-tests/)

# Go 言語のインストール

[Download and install](https://go.dev/doc/install) に従う  
↓ のようになればインストール OK

```go
$ go version
go version go1.15.7 darwin/amd64
```

# コード実行例

```go
$ go run helloWorld/hello.go
Hello, world
```

# テストの実行

```go
$ go test
package .: no Go files in 絶対パス
```

となる場合は

- `go mod init hello` しておく
- もしくは`go test`を実行したディレクトリに.go ファイルがない

## 参考 URL

https://andmorefine.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

# エラーまとめ

- `function main is undeclared in the main package`
  → main 関数が存在しない
- `main redeclared in this block previous declaration at ./hello.go:12:6`
  → main という関数名が被っている

# API サンプル実行方法

`apiSample/endpoint.go` に定義された処理を使う

```sh
# helloWorld用
COPY ./apiSample/endpoint.go .

# 〜略〜

# helloWorld用
CMD ["go", "run", "endpoint.go"]
```

を有効にして

```sh
docker-compose build # 関数やDockerfileなどに変更があった場合に必要
docker-compose up
```

すると http://localhost:8085 で hello というレスポンスが返ってくる

# CRUD API 実行方法

```sh
# todoアプリのcrud用（gin使用）
COPY ./crud/main.go .

# 〜略〜

# todoアプリのcrud用（gin使用）
CMD ["go", "run", "main.go"]
```

を有効にして

```
docker-compose build
docker-compose up
```

## CREATE

```sh
$ curl --location --request POST 'http://localhost:8085/todo'
{"id":8,"createdAt":"2022-09-17T16:01:05.686Z","updatedAt":"2022-09-17T16:01:05.686Z","deletedAt":null,"name":""}
```

という感じでレコードを create できる。
name とかは パラメータ から受け取っていないので、
受け取って更新できるようにしたい

## READ

```sql
INSERT INTO `todos` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`)
VALUES
	(2, '2022-01-01 00:00:00.000', '2022-01-01 00:00:00.000', NULL, 'aaaaaa');
```

のようなレコードがあれば、
http://localhost:8085/todo/2
で見られる

## UPDATE

```sh
$ curl --location --request PUT 'http://localhost:8085/todo/5'
{"id":5,"createdAt":"2022-09-17T15:56:24.767Z","updatedAt":"2022-09-17T16:07:21.169Z","deletedAt":null,"name":""}
```

で update 可能。
updated_at が更新される。

## DELETE

```sh
$ curl --location --request DELETE 'http://localhost:8085/todo/3'
{
    "id": 3,
    "createdAt": "2022-09-17T15:56:05.378Z",
    "updatedAt": "2022-09-17T15:56:05.378Z",
    "deletedAt": "2022-09-17T16:05:43.92Z",
    "name": ""
}
```

という感じで delete 可能
deleted_at にタイムスタンプが入る論理削除。

# ローカル環境 sequel pro 設定値

![picture 1](images/46ec2656b4aa081d7922b0be6e7ae7b9b4c0a3855ca5971b7a090e1490bc1ad5.png)
パスワードは xxx

# 必要ライブラリインストールコマンド

```sh
go get -u github.com/gorilla/mux # 参考：https://github.com/gorilla/mux

# 参考：https://liginc.co.jp/584227 でCRUD作成の項目でimportしているライブラリをinstall
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/gin-gonic/gin
```

→ .sh ファイルか docker 管理にする
