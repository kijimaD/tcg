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
	"log"
	"net/http"
	"os"

	svg "github.com/ajstarks/svgo"
)

func main() {
	exec()

	fileServer := http.FileServer(http.Dir("."))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	// http.Handle("/circle", http.HandlerFunc(circle))
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func circle(w http.ResponseWriter, req *http.Request) {
	const width = 500
	const height = 800
	const padding = 10

	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(width, height)

	// 全体枠
	s.Rect(0, 0, width/2+padding*2, height/2+padding*2, "fill:none;stroke:black;")
	// タイトル背景
	s.Rect(0, 0, width/2+padding*2, 36, "fill:#000;fill-opacity:0.5;stroke:black;")
	s.Text(width/4, 30, "橋跡", "text-anchor:middle;font-size:16px;fill:black")
	s.Text(padding, 30+30+height/4+padding*2+30, "旧陣之尾橋跡", "font-size:16px;fill:black")
	s.Image(padding, 30+30, width/2, height/4+padding*2, "/static/image.png")
	s.Rect(padding, 30+30, width/2, height/4+padding*2, "fill:none;stroke:black;")
	s.End()
}

func exec() {
	const width = 500
	const height = 800
	const padding = 10

	f, err := os.Create("test.svg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := svg.New(f)
	s.Start(width/2, height/2)

	// 全体枠
	s.Rect(0, 0, width/2+padding*2, height/2+padding*2, "fill:none;stroke:black;")
	// タイトル背景
	s.Rect(0, 0, width/2+padding*2, 36, "fill:#000;fill-opacity:0.5;stroke:black;")
	s.Text(width/4, 30, "橋跡", "text-anchor:middle;font-size:16px;fill:black")
	s.Text(padding, 30+30+height/4+padding*2+30, "旧陣之尾橋跡", "font-size:16px;fill:black")
	s.Image(padding, 30+30, width/2, height/4+padding*2, "/static/image.png")
	s.Rect(padding, 30+30, width/2, height/4+padding*2, "fill:none;stroke:black;")
	s.End()
}
