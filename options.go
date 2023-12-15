package main

import (
	wailsLogger "github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

func NewOptions() *options.App {
	opts := &options.App{
		Title:            "LAVT",
		BackgroundColour: options.NewRGBA(27, 27, 155, 0),
		Width:            1920,
		Height:           1080,
	}

	// Frontend should start as hidden
	opts.Frameless = true
	opts.StartHidden = false
	opts.HideWindowOnClose = true
	opts.WindowStartState = options.Maximised

	// Application renders on top of other windows
	opts.AlwaysOnTop = true

	// Logging
	fileLogger := wailsLogger.NewFileLogger("app.log")
	opts.Logger = fileLogger
	opts.LogLevel = wailsLogger.DEBUG

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

	return opts
}
