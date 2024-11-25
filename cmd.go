package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

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
			Name:          "jinno",
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
			Name:          "nabeishi",
			Title:         "鍋石",
			PlaceCategory: "歴",
			BgPath:        "./images/bg/normalize/pattern_b.png",
			KeyPath:       "./images/key/normalize/nabeishi.png",
			Descs:         []string{"阿久根の七不思議の1つ。", "鍋の形をした岩。"},
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
			Name:          "r499",
			Title:         "国道499号線",
			PlaceCategory: "景",
			BgPath:        "./images/bg/normalize/pattern_c.png",
			KeyPath:       "./images/key/normalize/r499.png",
			Descs: []string{
				"阿久根市内の陸上区間はわずか62m",
				"しかない国道。",
				"市内で唯一の2車線道路区間。",
			},
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
			Name:          "okawa",
			Title:         "旧大川トンネル",
			PlaceCategory: "歴",
			BgPath:        "./images/bg/normalize/pattern_d.png",
			KeyPath:       "./images/key/normalize/okawa.png",
			Descs:         []string{"大川の鉄道トンネル跡。", "両側を閉鎖してあり侵入はできない。"},
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
	const srcDir = "./images/key/original"
	const destDir = "./images/key/normalize"

	c, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}
	for _, entry := range c {
		if !strings.HasSuffix(entry.Name(), ".png") {
			return fmt.Errorf("%sに.png以外のファイルが含まれている: %s", srcDir, entry.Name())
		}
		src := path.Join(srcDir, entry.Name())
		dest := path.Join(destDir, entry.Name())
		err := normalizeKey(src, dest)
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
	const srcDir = "./images/bg/original"
	const destDir = "./images/bg/normalize"

	c, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}
	for _, entry := range c {
		if !strings.HasSuffix(entry.Name(), ".png") {
			return fmt.Errorf("%sに.png以外のファイルが含まれている: %s", srcDir, entry.Name())
		}
		src := path.Join(srcDir, entry.Name())
		dest := path.Join(destDir, entry.Name())
		err := normalizeBg(src, dest)
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
