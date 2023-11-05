## mdlinkstripper
:warning: chatGPTに書いてもらったコード

- markdownのリンク部分 `[example.com](http://example.com)` を `example.com http://emample.com` のようにする

- markdownで一気に書いた後にテキストファイルにペーストする際にmarkdownのリンクだと見づらいので、タイトルとリンクだけの形式にする際に使う

使い方
```
go run main.go your_markdown_file.md
```
