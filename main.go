package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
)

// Embed static assets into binary go:embed
//
//go:embed all:frontend/build
var assets embed.FS

func main() {
	app := NewApp()
	opts := NewOptions()

	opts.Bind = []interface{}{
		app,
	}

	opts.OnStartup = app.startup
	opts.OnShutdown = app.shutdown

	appErr := wails.Run(opts)
	if appErr != nil {
		panic(appErr)
	}
}
