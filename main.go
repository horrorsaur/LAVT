package main

import (
	"embed"
	"net/http"

	"github.com/horrorsaur/LAVT/internal/backend/api/valorant"
	"github.com/wailsapp/wails/v2"
	wailsLogger "github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

// Embed static assets into binary go:embed
//
//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	app := NewApp()
	router := valorant.NewRouter()

	// Logging
	fileLogger := wailsLogger.NewFileLogger("app.log")

	opts := &options.App{
		Title:             "LAVT",
		BackgroundColour:  options.NewRGBA(27, 27, 155, 0),
		Width:             1920,
		Height:            1080,
		MinWidth:          1024,
		MinHeight:         768,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		WindowStartState:  options.Maximised,
		AlwaysOnTop:       true,
		Logger:            fileLogger,
		LogLevel:          wailsLogger.DEBUG,
		AssetServer: &assetserver.Options{
			Assets: assets,
			Middleware: func(next http.Handler) http.Handler {
				router.NotFound(next.ServeHTTP)
				return router
			},
		},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
		},
		Bind: []interface{}{
			app,
		},
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
	}

	err := wails.Run(opts)
	if err != nil {
		fileLogger.Error(err.Error())
		panic(err)
	}
}
