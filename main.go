package main

import (
	"embed"
	"net/http"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
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
		Width:             1920,
		Height:            1080,
		MinWidth:          1024,
		MinHeight:         768,
		Frameless:         false,
		StartHidden:       true,
		HideWindowOnClose: false,
		WindowStartState:  options.Normal,
		AlwaysOnTop:       true,
		Logger:            logger.NewFileLogger(os.Getenv("DEFAULT_LOG_FILE_PATH")),
		LogLevel:          logger.DEBUG,
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
		OnDomReady: app.onDomReady,
	}

	err := wails.Run(opts)
	if err != nil {
		panic(err)
	}
}
