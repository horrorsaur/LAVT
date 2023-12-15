package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	api "github.com/horrorsaur/LAVT/internal/backend/api/valorant"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type (
	App struct {
		ctx    context.Context
		client *api.ValorantClient
	}
)

func NewApp() *App {
	return &App{}
}

// Called after the frontend has been created, but before index.html has been loaded.
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// setup lockfile watcher
	watcher, watcherErr := api.NewLockfileWatcher()
	if watcherErr != nil {
		panic(watcherErr) // panic for now, valo isnt installed or custom directory
	}

	// scan cache directory for lockfile
	ok, scanErr := watcher.Scan()

	// lockfile not found, setup watcher
	if !ok && errors.Is(scanErr, api.RiotClientLockfileNotFound) {
		runtime.LogInfo(ctx, "setting up watcher to scan every 5 seconds")
		watcher.Watch(5000 * time.Millisecond) // every 5s
	}

	lockfile := <-watcher.Ch
	client := api.NewClient(lockfile, fmt.Sprintf("https://127.0.0.1:%v", lockfile.Port))
	a.client = client
}

// Called after the frontend has been destroyed, just before the application terminates.
func (a *App) shutdown(ctx context.Context) {
	runtime.LogInfo(a.ctx, "Shutting down!")
}

func (a App) GetSessionInfo() (*api.SessionResponse, error) {
	return a.client.GetSession(a.ctx)
}
