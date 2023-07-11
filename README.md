# Rickenbacker

リポジトリをクローン後に、

```docker-compose build```

してもらうとDockerイメージのビルドが行われます。

### Dockerのよく使うコマンド

- 起動
```docker-compose up -d &```

- 廃棄
```docker-compose down```

- 停止
```docker-compose stop```

- 起動
```docker-compose start```

- バックエンドにSSH
```docker-compose exec backend bash```

- フロントエンドにSSH
```docker-compose exec frontend bash```

基本はupとdownを使えばいいと思います。

### Docker起動後のURL

- フロントエンド
```http://localhost:5173```
- バックエンド
```http://localhost:9090```
- PHPMyAdmin
```http://localhost:8888```
