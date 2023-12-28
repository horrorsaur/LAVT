package main

import (
	"embed"
	"net/http"

	"github.com/wailsapp/wails/v2"
	wailsLogger "github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

// Embed static assets into binary
//
//go:embed all:frontend/dist internal/templates
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	app := NewApp()
	opts := &options.App{
		Title:             "LAVT",
		BackgroundColour:  options.NewRGBA(27, 27, 155, 0),
		Width:             800, // 1920
		Height:            600, // 1080
		MinWidth:          1024,
		MinHeight:         768,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		WindowStartState:  options.Normal,
		AlwaysOnTop:       true,
		Logger:            wailsLogger.NewFileLogger("app.log"),
		LogLevel:          wailsLogger.DEBUG,
		AssetServer: &assetserver.Options{
			Assets: assets,
			Middleware: func(next http.Handler) http.Handler {
				app.router.NotFound(next.ServeHTTP)
				return app.router
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
		panic(err)
	}
}
