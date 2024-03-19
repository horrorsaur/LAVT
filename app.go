package main

import (
	"context"
	"errors"
	"log"
	"strings"
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
// runtime may not work since wndw is initializing in a different thread
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
		log.Print("lockfile not found! watching for lockfile...")
		watcher.Watch(5000 * time.Millisecond) // every 5s
	}

	if lockfile := <-watcher.Ch; lockfile != nil {
		a.client = api.NewClient(lockfile)
	}
}

// This callback is called after the frontend has loaded index.html and its resources. It is given the application context
func (a *App) onDomReady(ctx context.Context) {
	runtime.LogInfo(ctx, "DOM is ready!")

	entitlements, err := a.client.GetEntitlementsToken(ctx)
	if err != nil {
		runtime.LogErrorf(ctx, "error getting entitlements: %+v", err)
	}

	session, err := a.client.GetLocalSession(ctx)
	if err != nil {
		runtime.LogErrorf(ctx, "error getting presences: %+v", err)
	}

	// save session data to context for now, i think this expires after a certain amount of time
	// or if the client is closed (user needs to reauth)
	ctx = context.WithValue(ctx, "session", session)

	runtime.LogInfo(ctx, "connecting to riot websocket")
	connErr := a.client.ConnectToSocket(ctx)
	if connErr != nil {
		runtime.LogError(ctx, connErr.Error())
		panic(connErr) // panic for now
	}
	runtime.LogInfo(ctx, "connected!")

	// todo: dig a lil more to see if there is a better way to grab this data
	// but this works 4 now
	subsErr := a.client.Subscribe(ctx, "OnJsonApiEvent")
	if subsErr != nil {
		runtime.LogError(ctx, subsErr.Error())
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				runtime.LogInfo(ctx, "context cancelled, shutting down")
				return
			case <-time.After(2 * time.Second):
				r := a.client.ReceiveSocketMsgs(ctx)
				if strings.Contains(r.Msg.Uri, "ares-pregame") {
					log.Printf("received pregame message: %+v\n", r)
					tmp := strings.Split(r.Msg.Uri, "/")
					matchId := tmp[len(tmp)-1]
					log.Printf("got match ID! %+v", matchId)
					if matchId != "" {
						resp, err := a.client.GetPreGameMatchDetails(ctx, matchId, entitlements)
						if err != nil {
							log.Printf("error getting pregame match details: %+v\n", err)
						}

						log.Printf("RESP: %+v\n", resp)

					}
				}
			}
		}
	}()
}

// Called after the frontend has been destroyed, just before the application terminates.
func (a *App) shutdown(ctx context.Context) {
	runtime.LogInfo(ctx, "Shutting down...")

	err := a.client.CloseSocket(ctx)
	if err != nil {
		runtime.LogErrorf(ctx, "error on WS shutdown: %+v", err.Error())
	}

	runtime.LogInfo(ctx, "closed WS connection")
}
