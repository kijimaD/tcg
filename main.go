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
	"net/http"
	"os"

	svg "github.com/ajstarks/svgo"
)

func main() {
	inputPath := "image.png"
	outputPath := "normalize.png"
	normSize := 230

	img, err := loadImage(inputPath)
	if err != nil {
		fmt.Println("Error loading image:", err)
		return
	}
	croppedImg := SquareTrimImage(img, normSize)
	saveImage(croppedImg, outputPath)

	fileServer := http.FileServer(http.Dir("."))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.Handle("/check", http.HandlerFunc(check))
	err = http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func build(w io.Writer) {
	// カードの幅
	const cardWidth = 250
	// カードの縦
	var cardHeight = int(cardWidth * 1.6)
	const padding = 10
	// キービジュアル
	const keyVisualWidth = 230
	const keyVisualHeight = 230
	// 文字の高さ
	const lineHeight = 16

	s := svg.New(w)
	s.Start(cardWidth, cardHeight)

	// 全体枠
	body := func() {
		s.Rect(0, 0, cardWidth, cardHeight, "fill:royalblue;rx:10;ry:10;")
	}

	// キービジュアル
	keyVisualBG := func() {
		s.Rect(padding, padding+lineHeight*2, keyVisualWidth, keyVisualHeight, "fill:none;stroke:gold;")
	}
	keyVisual := func() {
		s.Image(padding, padding+lineHeight*2, keyVisualWidth, keyVisualHeight, fmt.Sprintf("data:image/png;base64,%s", base64nize("./normalize.png")))
	}

	// 説明文
	descBG := func() {
		s.Rect(padding, lineHeight*2+keyVisualWidth+padding, cardWidth-padding*2, lineHeight*7, "fill:white;fill-opacity:1.0;rx:8;ry:8")
	}
	desc := func() {
		s.Text(padding*2, lineHeight*2+keyVisualWidth+padding*4, "橋台が残っている", fmt.Sprintf("font-size:%dpx;fill:black", lineHeight))
	}

	// タイトル
	titleBG := func() {
		s.Rect(0, padding, cardWidth, lineHeight*2, "fill:white;fill-opacity:1.0;stroke:black;")
	}
	title := func() {
		s.Text(cardWidth/4, lineHeight*2, "旧陣之尾橋跡", fmt.Sprintf("text-anchor:middle;font-size:%dpx;fill:black;", lineHeight))
	}

	body()
	keyVisualBG()
	keyVisual()
	descBG()
	desc()
	titleBG()
	title()

	s.End()
}

func check(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	f, err := os.Create("test.svg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	mw := io.MultiWriter(f, w)
	build(mw)
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
