# go-lambda-practice

これをもとにやってみる。

https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/go-image.html

## はまったところ

* ディレクトリ構成がよくわからない
    * 想定ではルートフォルダに lambda handler としたいプログラムが置かれそうな気がする
* alpineでのビルド
    * 一応ドキュメントの後の方には書いてある
    * 先に書いてある方法でやると lambda のエミュレーターとか含まれたイメージがビルドされる
    * これが結構ファイルサイズある(100MBくらい)ので、ローカルデバッグ用にしておき本格デプロイ時のみ alpine かなという感じ
        * ECRには利用制限があるため、それを考えても image サイズは小さいほうが良い
    * 手元ではどうすべきか悩んだ結果、Dockerfileごと分けてみることにした
* WORKDIRが空になる問題
    * lambda上で設定するか、Dockerfile上で設定する必要がある
    * Dockerfile上に書いておいて、必要なら lambda 上で上書きする運用が無難かなという感想
* alpine 利用時に rie [Runtime Interface Emulatorのこと](https://docs.aws.amazon.com/lambda/latest/dg/images-test.html)使ってデバッグするところ
    * サンプルに書いてある entry.sh の指定パスが Dockerfile 上に書くパスと異なる
    * `/usr/local/bin/aws-lambda-rie` に rie を入れるよう Dockerfile 側を書き換えるとよさそう
* デバッグ用の curl コマンド
    * 最後のコンマはいらないと思う（誤植？）

## 使い方とか

ローカルで動かす

```shell
docker build -f ./build/development/Dockerfile -t hello-world-lambda .
docker run -p 9000:8080 hello-world-lambda:latest /main
```

この状態で例の curl コマンド（ドキュメント参照）を打つと動くはず

デプロイ用の Dockerfile は分けることにした。

```shell
docker build -f ./build/production/Dockerfile -t hello-world-lambda .
```

あとは tag を打って ECR に push する（ドキュメント参照）

lambda との紐づけ方もドキュメント参照

## 参考

* ディレクトリ構成
  * [golang-standards/project-layout： Standard Go Project Layout](https://github.com/golang-standards/project-layout)
  * [Goにはディレクトリ構成のスタンダードがあるらしい。 - Qiita](https://qiita.com/sueken/items/87093e5941bfbc09bea8)
* gofmtの設定
  * [i++](http://increment.hatenablog.com/?page=1461757090)
  * [Goで開発を始める前に絶対に読んでほしいGolandの設定3選 - Qiita](https://qiita.com/tez/items/417c72a275fd1399645e#pre-commit%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB)
  * [misc/git/pre-commit - go - Git at Google](https://go.googlesource.com/go/+/dev.tls/misc/git/pre-commit)
* エラー回り
  * [standard_init_linux.go：211： exec user process caused "no such file or directory" の直し方 - Qiita](https://qiita.com/kabik/items/5591f62c0ef6ddef5db2)
    * 記載通り entry.sh が CRLF で保存されていた（がっくり）
* 環境変数
  * [Docker で環境変数をホストからコンテナに渡す方法（ホスト OS 側からゲスト OS に渡す方法各種） - Qiita](https://qiita.com/KEINOS/items/518610bc2fdf5999acf2)
    * 一応見ただけ
    * rie なしで run すると AWS のもろもろの環境変数がないといわれてエラーになるので調べていた
* lambdaに関して
  * [AWS Lambda の新機能 – コンテナイメージのサポート | Amazon Web Services ブログ](https://aws.amazon.com/jp/blogs/news/new-for-aws-lambda-container-image-support/)
