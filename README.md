トレカ生成ツール。

## 背景生成方法

手動でやる。

- 背景テクスチャ画像を入手する。https://free-paper-texture.com/
- 250x400に近いサイズにリサイズする。https://www.iloveimg.com/ja/resize-image
- 250x400に切り抜く。https://www.iloveimg.com/ja/crop-image
- 角を丸くする。https://www.quickpicturetools.com/jp/rounded_corners/

## tree

```
$ tree
.
├── cmd.go
├── go.mod
├── go.sum
├── images
│   ├── bg
│   │   ├── normalize
│   │   │   └── bg.png
│   │   └── original
│   ├── card
│   │   └── jinno.svg
│   └── key
│       ├── normalize
│       │   └── jinno.png
│       └── original
│           └── jinno.png
├── kousei.drawio.svg
├── main.go
├── normalize.go
└── README.md
```
