トレカ生成ツール。

<table>
<tr>
<td><img src="./images/card/jinno.svg" /></td>
<td><img src="./images/card/nabeishi.svg" /></td>
<td><img src="./images/card/okawa.svg" /></td>
<td><img src="./images/card/r499.svg" /></td>
</tr>
</table>

## command

```shell
go run . -h
───────────────────────────────────────────────────────
 _______ _______  ______
    |    |       |  ____
    |    |_____  |_____|
Trading Card Generator by kijimaD
───────────────────────────────────────────────────────

NAME:
   tcg - tcg [subcommand] [args]

USAGE:
   tcg [global options] command [command options]

VERSION:
   v0.0.1

DESCRIPTION:
   Trading Card Generation tool

COMMANDS:
   build         build
   server        server
   normalizeKey  normalizeKey
   normalizeBg   normalizeBg
   help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## tree

```
$ tree
.
├── cmd.go
├── go.mod
├── go.sum
├── images
│   ├── bg
│   │   ├── normalize
│   │   │   ├── patternA.png
│   │   │   └── patternB.png
│   │   └── original
│   │       ├── patternA.png
│   │       └── patternB.png
│   ├── card
│   │   └── jinno.svg
│   └── key
│       ├── normalize
│       │   └── jinno.png
│       └── original
│           └── jinno.png
├── kousei.drawio.svg
├── main.go
├── normalize.go
└── README.md

8 directories, 14 files
```
