package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) SetTransparentWindow() {
	runtime.WindowSetBackgroundColour(a.ctx, 0, 0, 0, 0)
}

// "github.com/lxn/win"
// hwnd := win.FindWindow(nil, syscall.StringToUTF16Ptr("Your App Title"))
// win.SetWindowLong(hwnd, win.GWL_EXSTYLE, win.GetWindowLong(hwnd, win.GWL_EXSTYLE)|win.WS_EX_LAYERED)
