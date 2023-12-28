package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-chi/chi/v5"
	api "github.com/horrorsaur/LAVT/internal/backend/api/local"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type (
	App struct {
		ctx    context.Context
		client *api.ValorantClient
		router *chi.Mux
	}
)

func NewApp() *App {
	return &App{
		router: api.NewRouter(),
	}
}

// Called after the frontend has been created, but before index.html has been loaded.
func (a *App) startup(ctx context.Context) {
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

	if lockfile := <-watcher.Ch; lockfile != nil {
		a.client = api.NewClient(lockfile)

		entitlements, err := a.client.GetEntitlementsToken(ctx)
		if err != nil {
			runtime.LogInfof(ctx, "error getting entitlements: %+v", err)
			log.Print(err)
		}

		presences, err := a.client.GetPresences(ctx)
		if err != nil {
			runtime.LogInfof(ctx, "error getting presences: %+v", err)
			log.Print(err)
		}

		// cache request data
		ctx = context.WithValue(ctx, "entitlements", entitlements)
		ctx = context.WithValue(ctx, "presences", presences)
	}

	a.ctx = ctx

	log.Printf("entitlements: %+v\n\n\n", ctx.Value("entitlements"))
	log.Printf("presences: %+v\n\n\n", ctx.Value("presences"))
}

// Called after the frontend has been destroyed, just before the application terminates.
func (a *App) shutdown(ctx context.Context) {
	runtime.LogInfo(ctx, "Shutting down...")

	err := a.client.CloseWS(ctx)
	if err != nil {
		runtime.LogErrorf(ctx, "error on WS shutdown: %+v", err.Error())
	}
	runtime.LogInfo(ctx, "closed WS connection")
}
