### テストカバレッジの測定

```zsh
$ go test -cover ./... -coverprofile=cover.out
$ go tool cover -html=cover.out -o cover.html
```

「./...」は、Go言語のパッケージマネージャである「go」コマンドにおいて、カレントディレクトリ以下のすべてのサブディレクトリを含むパッケージをビルド、テスト、インストールするために使用される特殊なフラグ

