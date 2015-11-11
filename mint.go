package main

import (
	"os"

	"gopkg.in/leyra/cli.v1"
    //"github.com/tdewolff/minify/css"
    //"github.com/tdewolff/minify/html"
    //"github.com/tdewolff/minify/js"
)

func main() {
	app := cli.NewApp()

	app.Name = "mint"
	app.Usage = "a utility for minifying files"
	app.Commands = []cli.Command{
		{
			Name:  "javascript",
			Usage: "minify javascript files",
			Action: func(c *cli.Context) {
				
			},
			Aliases: []string{"js"},
		},
	}
	app.Run(os.Args)
}
