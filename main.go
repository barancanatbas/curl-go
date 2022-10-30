package main

import (
	"io"
	"log"
	"os"

	"github.com/barancanatbas/curl-go/curl"
	"github.com/urfave/cli/v2"
)

var Version = "v1.0.0"

func main() {
	app := &cli.App{
		Name:     "curl-go",
		Usage:    "A curl package written in golang.",
		Commands: Commands(os.Stdin),
		Version:  Version,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func Commands(reader io.Reader) []*cli.Command {
	return []*cli.Command{
		curl.Command(),
	}
}
