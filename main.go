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
	err := wails.Run(&options.App{
		Title:             "LAVT",
		Width:             800,
		Height:            600,
		BackgroundColour:  options.NewRGBA(27, 27, 155, 0),
		WindowStartState:  options.Maximised,
		Frameless:         true,
		StartHidden:       true,
		HideWindowOnClose: true,
		AlwaysOnTop:       true,
		LogLevel:          logger.INFO,
		Logger:            logger.NewFileLogger("./app.log"),
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
