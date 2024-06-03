package main

import (
	"github.com/alecthomas/kong"
	"bytes"
	"fmt"
	"net/http"
)

type Context struct {
	Debug bool
}

type PublishCmd struct {
	Category string `arg:"" name:"category" help:"Event category." env:"CATEGORY,PLUGIN_CATEGORY" enum:"alerts,broadcasts,chronosphere,deploys,feature_flags,infrastructure,third_party"`
	ChronosphereEventsAPI string `arg:"" name:"chronosphere_events_api" help:"URL for Chronosphere Events API like https://ADDRESS.chronosphere.io/api/v1/data/events" env:"CHRONOSPHERE_EVENTS_API,PLUGIN_CHRONOSPHERE_EVENTS_API"`
	ChronosphereApiToken string `arg:"" name:"chronosphere_api_token" help:"API token for Chronosphere" env:"CHRONOSPHERE_API_TOKEN,PLUGIN_CHRONOSPHERE_API_TOKEN"`
}
func (p *PublishCmd) Run(ctx *Context) error {
	fmt.Println("Publishing event with category", p.Category)

	var jsonStr = []byte(`{"event": {"title": "Dustin test", "category": "deploys", "type": "deploy_start", "payload_object": {}, "labels": {}, "source": "local", "happened_at": "2024-06-03T12:42:00Z"}}`)

	if ctx.Debug {
		fmt.Println("API:", p.ChronosphereEventsAPI)
	}

	req, err := http.NewRequest("POST", p.ChronosphereEventsAPI, bytes.NewBuffer(jsonStr))
	req.Header.Set("API-Token", p.ChronosphereApiToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		if ctx.Debug {
			fmt.Println(resp)
		}
		panic(err)
	}
	return nil
}
func (p *PublishCmd) Validate() error {
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
