package main

import (
	"log"
	"net/http"
)

func indexHandle(w http.ResponseWriter, req *http.Request) {
	str := []byte(`
<!DOCTYPE html>
<html lang="ja">
    <head>
        <meta charset="utf-8">
        <link  href="https://cdnjs.cloudflare.com/ajax/libs/viewerjs/1.11.7/viewer.css" rel="stylesheet">
        <style>
         ul { list-style-type: none; }
         li { display: inline-block; }
        </style>
    </head>
    <body>
        <ul id="images">
            <li><img src="/static/images/card/jinno.svg"></li>
            <li><img src="/static/images/card/nabeishi.svg"></li>
            <li><img src="/static/images/card/okawa.svg"></li>
            <li><img src="/static/images/card/r499.svg"></li>
        </ul>
    </body>
    <script src="https://code.jquery.com/jquery-2.2.0.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/viewerjs/1.11.7/viewer.min.js"></script>
    <script>
     var viewer = new Viewer(document.getElementById('images'));
    </script>
</html>
`)
	_, err := w.Write(str)
	if err != nil {
		log.Fatal(err)
	}
}
