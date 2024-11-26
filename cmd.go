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
	bytes, err := os.ReadFile("./raw.toml")
	if err != nil {
		return (err)
	}
	rawMaster, err := Load(string(bytes))
	if err != nil {
		return err
	}
	for _, place := range rawMaster.Raws.Places {
		f, err := os.Create(fmt.Sprintf("./images/card/%s.svg", place.Name))
		if err != nil {
			return err
		}
		defer f.Close()
		place.build(f)
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
