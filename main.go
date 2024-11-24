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
	const width = 250
	// カードの縦
	var height = int(width * 1.6)
	const padding = 10
	// キービジュアル
	const keyVisualWidth = 230
	const keyVisualHeight = 230

	s := svg.New(w)
	s.Start(width, height)

	// 全体枠
	s.Rect(0, 0, width+padding*2, height+padding*2, "fill:royalblue;rx:10;ry:10;")

	// キービジュアル
	s.Image(padding, padding+16*2, keyVisualWidth, keyVisualHeight, fmt.Sprintf("data:image/png;base64,%s", imageBase("./normalize.png")))
	// キービジュアル枠
	s.Rect(padding, padding+16*2, keyVisualWidth, keyVisualHeight, "fill:none;stroke:gold;")

	// 本文
	s.Rect(padding, 16*2+keyVisualWidth+padding, width-padding*2, 16*7, "fill:white;fill-opacity:1.0;rx:8;ry:8")
	s.Text(padding*2, 16*2+keyVisualWidth+padding*4, "橋台が残っている", "font-size:16px;fill:black")

	// タイトル
	s.Rect(0, padding, width+padding*2, 16*2, "fill:white;fill-opacity:1.0;stroke:black;")
	s.Text(width/4, 16*2, "旧陣之尾橋跡", "text-anchor:middle;font-size:16px;fill:black;")

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

func imageBase(src string) string {
	filePath := src // PNG画像のパス
	imageData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	base64Image := base64.StdEncoding.EncodeToString(imageData)

	return base64Image
}
