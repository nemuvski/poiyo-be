# Poiyo Backend

## 実行環境構築

事前にDocker実行環境を用意する。用意した後に以下の手順に進む。

### `.env` に情報を記入

PostgreSQLといったサービスの設定情報を記述する。

```bash
cp docker/postgresql.env.example docker/postgresql.env
vi docker/postgresql.env

cp docker/pgadmin.env.example docker/pgadmin.env
vi docker/pgadmin.env
```

### イメージのビルドとコンテナ起動

Goを実行する環境とPostgreSQLを実行する環境を提供する。

```bash
docker-compose build
docker-compose up -d
```

時間がかかる場合があるので待つ。

### pgAdmin

http://localhost:5433

ユーザー名(Eメール)とパスワードは適宜 `.env` に記述した内容を入れる。

接続先のサーバーは、ホスト名 `postgresql` で登録する。

## 開発

### Go

Go言語実行用のコンテナで処理する。

```bash
docker exec -it app /bin/sh
```

その後、利用しているパッケージのダウンロード、またはプログラムを実行する。

```bash
go mod download
go run src/main.go
```
