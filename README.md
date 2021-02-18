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

デプロイ

```shell
docker build -f ./build/package/Dockerfile -t hello-world-lambda .
```

あとは tag を打って ECR に push する（ドキュメント参照）

lambda との紐づけ方もドキュメント参照
