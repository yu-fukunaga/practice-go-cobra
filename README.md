# practice-go-cobra
cobraを試すだけのリポジトリ

## 環境構築
TODO: write direnv install reference

```
$ make env
$ make setup
```

## おまけ
cobra-cliをインストールして、initコマンドを叩くと、cobra最小構成を生成できる
```
$ go install github.com/spf13/cobra-cli@latest
$ cobra-cli init
```