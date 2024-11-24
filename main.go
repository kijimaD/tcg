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
	normSize := 250

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
	const width = 500
	const height = 800
	const padding = 10

	s := svg.New(w)
	s.Start(width/2+padding*2, height/2+padding*2)

	// 全体枠
	s.Rect(0, 0, width/2+padding*2, height/2+padding*2, "fill:royalblue;rx:10;ry:10;")

	// 画像
	s.Image(padding, padding+16*2, 250, 250, fmt.Sprintf("data:image/png;base64,%s", imageBase("./normalize.png")))
	// 画像枠
	s.Rect(padding, padding+16*2, width/2, 250, "fill:none;stroke:gold;")

	// 本文
	s.Rect(padding, height/4+padding*6, width/2, 16*8, "fill:white;fill-opacity:1.0;rx:8;ry:8")
	s.Text(padding*2, height/4+padding*10, "橋台が残っている", "font-size:16px;fill:black")

	// タイトル
	s.Rect(0, padding, width/2+padding*2, 16*2, "fill:white;fill-opacity:1.0;stroke:black;")
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
