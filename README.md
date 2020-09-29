## 概要
<p>
2020年9月19日 - 2020年9月21日<br>
High Light Road で利用するAPIベース実装<br>
API仕様書は SwaggerEditor に定義ファイルの内容を入力して参照してください。
</p>

SwaggerEditor: <https://editor.swagger.io> <br>
定義ファイル: `./api-tomoshika.yaml`<br>

### API用のデータベースの接続情報を設定する
環境変数にデータベースの接続情報を設定します。<br>
ターミナルのセッション毎に設定したり、.bash_profileで設定を行います。

Macの場合
```
$ export MYSQL_USER=mori \
    MYSQL_PASSWORD=mori \
    MYSQL_HOST=127.0.0.1 \
    MYSQL_PORT=3306 \
    MYSQL_DATABASE=hlr_db
```

Windowsの場合
```
$ SET MYSQL_USER=mori
$ SET MYSQL_PASSWORD=mori
$ SET MYSQL_HOST=127.0.0.1
$ SET MYSQL_PORT=3306
$ SET MYSQL_DATABASE=hlr_db
```

## APIローカル起動方法
```
$ go run ./cmd/main.go
```
