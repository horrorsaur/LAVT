package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

// Embed static assets into binary go:embed

//go:embed all:frontend/build
var assets embed.FS

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title:            "LAVT",
		Width:            1080,
		Height:           720,
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		WindowStartState: options.Fullscreen,
		Fullscreen:       true,
		Frameless:        true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			// Webview is transparent when alpha=0
			WebviewIsTransparent: true,
			WindowIsTranslucent:  false, // windows frosty thing
			BackdropType:         windows.Mica,
			Theme:                windows.SystemDefault,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
