## サービス

https://poiyo.herokuapp.com

## 設定用意

**Config vars** に `GO_EXEC_ENV` (内容は `production` ) を用意する。

そのほかに環境変数を設定すること。 (`.env`を確認)

Firebaseのキーファイルは環境変数で持っておく。

```
heroku config:set FB_KEYFILE_JSON="$(< ./docker/app/poiyo-web-app.json)"
```

## :warning: デプロイ

開発後にデプロイする流れは次のようになる。

1. `master` ブランチに開発のコミットがされる。
2. `heroku` リポジトリの `master` に `production` の内容をプッシュする。
    - `git push heroku master`

※もしIDとパスワードが聞かれたら、 `heroku login` で事前にログインをしておくこと

## DB

管理画面のAdd-onをチェックすること。

CLIでも確認できる。

```
heroku pg:info
```

https://devcenter.heroku.com/ja/articles/heroku-postgresql を参照するとよい。

```
heroku pg:psql
heroku pg:psql postgresql-transparent-65826 --app poiyo
```

必要なエクステンションを入れる。

```
CREATE EXTENSION pgcrypto
    SCHEMA public
    VERSION "1.3";
```

その他、テーブル等は別ページに記述する。

## その他

### 変数確認

管理画面からも確認できる。

```
heroku config
```

### ログ確認

```
heroku logs --tail

heroku logs -p postgres -t
```

### DBの初期設定

```
alter database {database_name} set timezone = 'Asia/Tokyo';
```

一度終了して、再度接続する。 現在時刻を確認する。

```
select current_timestamp;
```

## トラブルシューティング

### Goのバージョン指定

デフォルトだと `1.12` となっている。 (2021/2/13の段階では)

`go.mod` に次のように指定してあげるとよい。

```
// +heroku goVersion go1.15
```


