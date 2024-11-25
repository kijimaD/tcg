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

type Place struct {
	// プログラム上で識別する名前。アルファベット
	Name string
	// タイトル
	Title string
	// カテゴリ
	PlaceCategory string
	// カード全体の背景
	BgPath string
	// キービジュアル
	KeyPath string
	// 説明
	Descs []string
}

func (p Place) build(w io.Writer) {
	s := svg.New(w)
	s.Start(cardWidth, cardHeight)

	var curY = 0

	// 背景
	bg := func() {
		s.Image(0, 0, cardWidth, cardHeight, fmt.Sprintf("data:image/png;base64,%s", base64nize(p.BgPath)))
	}

	// タイトル
	title := func() {
		curY += padding
		h := lineHeight * 2
		s.Rect(0, curY, cardWidth, h, "fill:black;fill-opacity:0.6;")
		s.Text(cardWidth/4, h, p.Title, fmt.Sprintf("text-anchor:middle;font-size:%dpx;fill:white;", lineHeight))
		s.Text(cardWidth-padding*2, h+6, p.PlaceCategory, fmt.Sprintf("text-anchor:middle;font-size:%dpx;fill:white;", lineHeight*2))
		curY += h
	}

	// キービジュアル
	keyVisual := func() {
		h := keyVisualHeight
		s.Rect(padding, curY, keyVisualWidth, h, "fill:none;")
		s.Image(padding, curY, keyVisualWidth, h, fmt.Sprintf("data:image/png;base64,%s", base64nize(p.KeyPath)))
		curY += h
	}

	// 説明文
	desc := func() {
		h := lineHeight * 7
		s.Rect(padding, curY, cardWidth-padding*2, h, "fill:white;fill-opacity:0.6;rx:8;ry:8;stroke:black;stroke-width:2px;")
		curY += padding * 2
		for _, desc := range p.Descs {
			s.Text(padding*2, curY, desc, fmt.Sprintf("font-size:%dpx;fill:black;", descFontSize))
			curY += lineHeight
		}
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
