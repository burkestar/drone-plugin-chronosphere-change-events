package main

import (
	"github.com/alecthomas/kong"
	"bytes"
	"fmt"
	"encoding/json"
	"net/http"
)

type Context struct {
	Debug bool
}

type PublishCmd struct {
	ChronosphereEventsAPI string `required:"" name:"chronosphere_events_api" help:"URL for Chronosphere Events API like https://ADDRESS.chronosphere.io/api/v1/data/events" env:"CHRONOSPHERE_EVENTS_API,PLUGIN_CHRONOSPHERE_EVENTS_API"`
	ChronosphereApiToken string `required:"" name:"chronosphere_api_token" help:"API token for Chronosphere" env:"CHRONOSPHERE_API_TOKEN,PLUGIN_CHRONOSPHERE_API_TOKEN"`
	Category string `required:"" name:"category" help:"Event category." env:"CATEGORY,PLUGIN_CATEGORY" enum:"alerts,broadcasts,chronosphere,deploys,feature_flags,infrastructure,third_party"`
	Type string `required:"" name:"type" help:"Event type, which can be any custom value." env:"TYPE,PLUGIN_TYPE"`
	Title string `name:"title" help:"Title for this event. If not provided, one will be constructed dynamically from other fields" env:"TITLE,PLUGIN_TITLE"`
	Source string `name:"source" default:"unknown" help:"Source where this event comes from." env:"SOURCE,PLUGIN_SOURCE"`
	// HappenedAt string `name:"happened_at" default:"" help:"Timestamp when event happened, e.g. 2024-06-03T12:42:00Z" env:"HAPPENED_AT,PLUGIN_HAPPENED_AT"`
	// Labels as key=value, comma separated
}

type PublishEvent struct {
	Title string `json:"title"`
	Category string `json:"category"`
	Type string `json:"type"`
	// HappenedAt string `json:"happened_at"`
	Labels map[string]string `json:"labels"`
	PayloadJson string `json:"payload_json"`
	Source string `json:"source"`
}

type PublishEventPayload struct {
	Event PublishEvent `json:"event"`
}

func (p *PublishCmd) Run(ctx *Context) error {
	fmt.Println("Publishing event with category", p.Category)

	if ctx.Debug {
		fmt.Println("API:", p.ChronosphereEventsAPI)
	}

	if p.Title == "" {
		p.Title = fmt.Sprintf("%s (%s) from %s", p.Type, p.Category, p.Source)
	}

	payload := PublishEventPayload{
		PublishEvent{
			Title: p.Title,
			Category: p.Category,
			Type: p.Type,
			// HappenedAt: p.HappenedAt,
			Labels: map[string]string{},
			PayloadJson: "",
			Source: p.Source,
		},
	}

	var jsonStr, err = json.Marshal(payload)

	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", p.ChronosphereEventsAPI, bytes.NewBuffer(jsonStr))
	req.Header.Set("API-Token", p.ChronosphereApiToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		if ctx.Debug {
			fmt.Println("JSON", string(jsonStr))
			fmt.Println("Response status", resp.Status)
			fmt.Println("Response body", resp.Body)
		}
		panic(err)
	}
	defer resp.Body.Close()
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
