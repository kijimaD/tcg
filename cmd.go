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
    |    |_____  |_____| @kijimaD
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

var CmdBuild = &cli.Command{
	Name:        "build",
	Usage:       "build",
	Description: "build",
	Action:      runBuild,
	Flags:       []cli.Flag{},
}

func runBuild(_ *cli.Context) error {
	f, err := os.Create("test.svg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	build(f)

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
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.Handle("/check", http.HandlerFunc(checkHandle))
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	return nil
}

func checkHandle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	build(w)
}
