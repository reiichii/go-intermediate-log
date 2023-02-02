## Examples

```
# docker-composeでmysqlを立てる
$ docker-compose up -d

# goのapiサーバを起動
$ DB_USER=docker DB_PASSWORD=docker DB_NAME=sampledb go run main.go

# リクエスト例
curl http://localhost:8080/article -X POST \
-H 'Authorization: Bearer xxx.xxx.xxx'
-d '{"title":"a","contents":"b","user_name":"invalid_user"}' -w '%{http_code}\n'
```

GCPコンソールからidトークンを取得し、認証client_idをソースの中で定義し直し、curlリクエストのヘッダにIDトークンのjwtを入れれば動く

[OAuth 2.0 を使用した Google API へのアクセス  |  Authorization  |  Google Developers](https://developers.google.com/identity/protocols/oauth2?hl=ja)
