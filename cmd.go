package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

const splash = `───────────────────────────────────────────────────────
 _______ _______  ______
    |    |       |  ____
    |    |_____  |_____|
Trading Card Generator by kijimaD
───────────────────────────────────────────────────────
`

func NewMainApp() *cli.App {
	app := cli.NewApp()
	app.Name = "tcg"
	app.Usage = "tcg [subcommand] [args]"
	app.Description = "Trading Card Generation tool"
	app.DefaultCommand = CmdBuild.Name
	app.Version = "v0.0.1"
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		CmdBuild,
		CmdServer,
		CmdNormalizeKey,
		CmdNormalizeBg,
	}
	cli.AppHelpTemplate = fmt.Sprintf(`%s
%s
`, splash, cli.AppHelpTemplate)

	return app
}

func RunMainApp(app *cli.App, args ...string) error {
	err := app.Run(args)
	if err != nil {
		return fmt.Errorf("コマンド実行が失敗した: %w", err)
	}

	return nil
}

// ================

var CmdBuild = &cli.Command{
	Name:        "build",
	Usage:       "build",
	Description: "build",
	Action:      runBuild,
	Flags:       []cli.Flag{},
}

func runBuild(_ *cli.Context) error {
	{
		p := Place{
			Name:          "jinno_a",
			Title:         "旧陣之尾橋跡",
			PlaceCategory: "歴",
			BgPath:        "./images/bg/normalize/pattern_a.png",
			KeyPath:       "./images/key/normalize/jinno.png",
			Descs:         []string{"折口川に架かっていた橋の跡。", "橋台が残っている。"},
		}
		f, err := os.Create(fmt.Sprintf("./images/card/%s.svg", p.Name))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		p.build(f)
	}
	{
		p := Place{
			Name:          "jinno_b",
			Title:         "旧陣之尾橋跡",
			PlaceCategory: "歴",
			BgPath:        "./images/bg/normalize/pattern_b.png",
			KeyPath:       "./images/key/normalize/jinno.png",
			Descs:         []string{"折口川に架かっていた橋の跡。", "橋台が残っている。"},
		}
		f, err := os.Create(fmt.Sprintf("./images/card/%s.svg", p.Name))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		p.build(f)
	}
	{
		p := Place{
			Name:          "okawa_a",
			Title:         "旧大川トンネル",
			PlaceCategory: "歴",
			BgPath:        "./images/bg/normalize/pattern_b.png",
			KeyPath:       "./images/key/normalize/okawa.png",
			Descs:         []string{"大川の鉄道トンネル跡。", "両側から閉鎖してあり立ち入りできない。"},
		}
		f, err := os.Create(fmt.Sprintf("./images/card/%s.svg", p.Name))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		p.build(f)
	}
	{
		p := Place{
			Name:          "okawa_c",
			Title:         "旧大川トンネル",
			PlaceCategory: "歴",
			BgPath:        "./images/bg/normalize/pattern_c.png",
			KeyPath:       "./images/key/normalize/okawa.png",
			Descs:         []string{"大川の鉄道トンネル跡。", "両側から閉鎖してあり立ち入りできない。"},
		}
		f, err := os.Create(fmt.Sprintf("./images/card/%s.svg", p.Name))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		p.build(f)
	}

	return nil
}

// ================

var CmdNormalizeKey = &cli.Command{
	Name:        "normalizeKey",
	Usage:       "normalizeKey",
	Description: "normalizeKey",
	Action:      runNormalizeKey,
	Flags:       []cli.Flag{},
}

func runNormalizeKey(_ *cli.Context) error {
	{
		err := normalizeKey("./images/key/original/jinno.png", "./images/key/normalize/jinno.png")
		if err != nil {
			return err
		}
	}
	{
		err := normalizeKey("./images/key/original/okawa.png", "./images/key/normalize/okawa.png")
		if err != nil {
			return err
		}
	}

	return nil
}

// ================

var CmdNormalizeBg = &cli.Command{
	Name:        "normalizeBg",
	Usage:       "normalizeBg",
	Description: "normalizeBg",
	Action:      runNormalizeBg,
	Flags:       []cli.Flag{},
}

func runNormalizeBg(_ *cli.Context) error {
	{
		err := normalizeBg("./images/bg/original/pattern_a.png", "./images/bg/normalize/pattern_a.png")
		if err != nil {
			return err
		}
	}
	{
		err := normalizeBg("./images/bg/original/pattern_b.png", "./images/bg/normalize/pattern_b.png")
		if err != nil {
			return err
		}
	}
	{
		err := normalizeBg("./images/bg/original/pattern_c.png", "./images/bg/normalize/pattern_c.png")
		if err != nil {
			return err
		}
	}

	return nil
}

// ================

var CmdServer = &cli.Command{
	Name:        "server",
	Usage:       "server",
	Description: "server",
	Action:      runServer,
	Flags:       []cli.Flag{},
}

func runServer(_ *cli.Context) error {
	fileServer := http.FileServer(http.Dir("."))
	http.Handle("/", http.HandlerFunc(indexHandle))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	return nil
}

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
            <li><img src="/static/images/card/jinno_a.svg"></li>
            <li><img src="/static/images/card/jinno_b.svg"></li>
            <li><img src="/static/images/card/okawa_a.svg"></li>
            <li><img src="/static/images/card/okawa_c.svg"></li>
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
