package main

import (
	"fmt"

	opt "github.com/romnn/configo"
)

type myApp struct {
	Config appConfig
}

type appConfig struct {
	Sanitize           *opt.Flag
	ShowErrors         *opt.Flag
	EnableExperimental *opt.Flag
}

func (c appConfig) defaultConfig() appConfig {
	return appConfig{
		EnableExperimental: opt.SetFlag(false),
		Sanitize:           opt.SetFlag(true),
	}
}

func (app *myApp) Start() {
	// Merge the default config
	opt.MergeConfig(&app.Config, app.Config.defaultConfig())

	// Will be true for all
	fmt.Printf("Sanitize=%t\n", opt.Enabled(app.Config.Sanitize))
	fmt.Printf("ShowErrors=%t\n", opt.Enabled(app.Config.ShowErrors))
	fmt.Printf("EnableExperimental=%t\n", opt.Enabled(app.Config.EnableExperimental))
}

func main() {
	app := &myApp{
		Config: appConfig{
			EnableExperimental: opt.SetFlag(true),
			ShowErrors:         opt.SetFlag(true),
		},
	}
	app.Start()
}
