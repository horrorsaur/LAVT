package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

// TODO:
// Add custom logger
// Add lockfile validator for when Riot Client is closed

// Embed static assets into binary go:embed

//go:embed all:frontend/build
var assets embed.FS

func main() {
	app := NewApp()

	opts := &options.App{
		Title:            "LAVT",
		BackgroundColour: options.NewRGBA(27, 27, 155, 0),
	}

	// Frontend should start as hidden
	opts.Frameless = false
	opts.StartHidden = false
	opts.HideWindowOnClose = true
	opts.WindowStartState = options.Maximised

	// Application renders on top of other windows
	opts.AlwaysOnTop = true

	// Logging
	opts.LogLevel = logger.INFO
	opts.Logger = logger.NewFileLogger("./app.log")

	// Embedded assets
	opts.AssetServer = &assetserver.Options{
		Assets: assets,
	}

	// Windows platform options
	opts.Windows = &windows.Options{
		WebviewIsTransparent: true,
		WindowIsTranslucent:  true,
	}

	// Linux platform options
	opts.Linux = &linux.Options{
		WindowIsTranslucent: true,
	}

	opts.Bind = []interface{}{
		app,
	}

	opts.OnStartup = app.startup
	opts.OnShutdown = app.shutdown

	// Run the app
	err := wails.Run(opts)
	if err != nil {
		println("Error:", err.Error())
	}
}
