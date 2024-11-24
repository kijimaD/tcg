// -*- mode:go;mode:go-playground -*-
// snippet of code @ 2024-11-23 20:25:15

// === Go Playground ===
// Execute the snippet with:                 Ctl-Return
// Provide custom arguments to compile with: Alt-Return
// Other useful commands:
// - remove the snippet completely with its dir and all files: (go-playground-rm)
// - upload the current buffer to playground.golang.org:       (go-playground-upload)

package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"

	svg "github.com/ajstarks/svgo"
)

// カードの幅
const cardWidth = 250

// カードの縦
var cardHeight = int(cardWidth * 1.6)

// 余白
const padding = 10

// キービジュアル
const keyVisualWidth = 230
const keyVisualHeight = 230

// 文字の高さ
const lineHeight = 16
const descFontSize = 12

func main() {
	app := NewMainApp()
	err := RunMainApp(app, os.Args...)
	if err != nil {
		log.Fatal(err)
	}
}

func normalize() {
	inputPath := "./images/key/original/jinno.png"
	outputPath := "./images/key/normalize/jinno.png"

	img, err := loadImage(inputPath)
	if err != nil {
		fmt.Println("Error loading image:", err)
		return
	}
	croppedImg := SquareTrimImage(img, keyVisualWidth)
	saveImage(croppedImg, outputPath)
}

func build(w io.Writer) {
	s := svg.New(w)
	s.Start(cardWidth, cardHeight)

	var curY = 0

	// 背景
	bg := func() {
		s.Image(0, 0, cardWidth, cardHeight, fmt.Sprintf("data:image/png;base64,%s", base64nize("./images/bg/normalize/bg.png")))
	}

	// タイトル
	title := func() {
		curY += padding
		h := lineHeight * 2
		s.Rect(0, curY, cardWidth, h, "fill:white;fill-opacity:0.6;rx:8;ry:8;")
		s.Text(cardWidth/4, h, "旧陣之尾橋跡", fmt.Sprintf("text-anchor:middle;font-size:%dpx;fill:black;", lineHeight))
		s.Text(cardWidth-padding*2, h+6, "遺", fmt.Sprintf("text-anchor:middle;font-size:%dpx;fill:black;", lineHeight*2))
		curY += h
	}

	// キービジュアル
	keyVisual := func() {
		h := keyVisualHeight
		s.Rect(padding, curY, keyVisualWidth, h, "fill:none;")
		s.Image(padding, curY, keyVisualWidth, h, fmt.Sprintf("data:image/png;base64,%s", base64nize("./images/key/normalize/jinno.png")))
		curY += h
	}

	// 説明文
	desc := func() {
		h := lineHeight * 7
		s.Rect(padding, curY, cardWidth-padding*2, h, "fill:white;fill-opacity:0.6;rx:8;ry:8;stroke:black;stroke-width:2px;")
		curY += padding * 2
		s.Text(padding*2, curY, "橋台が残っている。", fmt.Sprintf("font-size:%dpx;fill:black;", descFontSize))
		curY += h
	}

	bg()
	title()
	keyVisual()
	desc()

	s.End()
}

func base64nize(src string) string {
	filePath := src // PNG画像のパス
	imageData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	base64Image := base64.StdEncoding.EncodeToString(imageData)

	return base64Image
}
