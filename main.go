package main

import (
	"github.com/alecthomas/kong"
	"fmt"
)

type Context struct {
	Debug bool
}

type PublishCmd struct {
	Category string `arg:"" name:"category" help:"Event category." env:"PLUGIN_CATEGORY"`
}
func (p *PublishCmd) Run(ctx *Context) error {
	fmt.Println("Publishing event with category", p.Category)
	return nil
}


var CLI struct {
	Debug bool `help:"Enable debug mode."`

	Publish PublishCmd `cmd:"" help:"Publish change event."`
}

func main() {
	ctx := kong.Parse(&CLI)
	err := ctx.Run(&Context{Debug: CLI.Debug})
	ctx.FatalIfErrorf(err)
}
