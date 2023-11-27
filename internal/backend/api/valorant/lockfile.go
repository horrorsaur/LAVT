package valorant

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	unknownLocalCacheDirectory = errors.New("could not finds users local cache directory")
	unknownValorantDirectory   = errors.New("could not find valorant directory")
	unexpectedLockfileData     = errors.New("lockfile data format is not expected")
	RiotClientLockfileNotFound = errors.New("the riot client lockfile was not found within the app dir")
)

type (
	// Riot LCU lockfile
	//
	// Process Name : Process ID : Port : Password : Protocol
	RiotClientLockfileInfo struct {
		Name     string
		PID      int
		Port     int
		Password string
		Protocol string
	}

	LockfileWatcher struct {
		Ch       chan (*RiotClientLockfileInfo)
		CacheDir string
	}

	LockfileReadDirectoryError struct{ Reason error }
)

func (e *LockfileReadDirectoryError) Error() string {
	return fmt.Sprintf("could not read lockfile directory. received %s", e.Reason)
}

func NewLockfileWatcher() (*LockfileWatcher, error) {
	// get the users local cache directory
	appDataDir, err := os.UserCacheDir()
	if err != nil {
		log.Printf("user cache dir err: %s", err)
		return nil, unknownLocalCacheDirectory
	}

	w := &LockfileWatcher{
		Ch:       make(chan *RiotClientLockfileInfo),
		CacheDir: appDataDir,
	}

	// check if the expected Riot Client directory exists
	if _, err := os.Stat(w.getLockfileDirectory()); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			log.Printf(err.Error())
			return nil, unknownValorantDirectory
		}
	}

	return w, nil
}

func (w LockfileWatcher) LockfilePath() string {
	return w.getLockfilePath()
}

// The path to the Riot Client lockfile
func (w LockfileWatcher) getLockfilePath() string {
	return filepath.Join(
		w.CacheDir,
		"Riot Games/Riot Client/Config/lockfile",
	)
}

// The path to the Riot Client config directory
func (w LockfileWatcher) getLockfileDirectory() string {
	return filepath.Join(
		w.CacheDir,
		"Riot Games/Riot Client/Config/",
	)
}

// Scan the given fs.DirEntry slice for the lockfile
func lockfileExists(entries []fs.DirEntry) bool {
	for _, file := range entries {
		if file.Name() == "lockfile" {
			return true
		}
	}
	return false
}

// Scans CacheDir for Riot Client lockfile.
func (w LockfileWatcher) Scan() (bool, error) {
	log.Println("scanning cache dir for lockfile")
	lockfileDir := w.getLockfileDirectory()

	entries, err := os.ReadDir(lockfileDir) // /Config
	if err != nil {
		log.Print(err.Error())
		return false, &LockfileReadDirectoryError{Reason: err}
	}

	lockfileExists := lockfileExists(entries)
	if lockfileExists {
		log.Print("lockfile found, reading data...")
		data, err := readLockfile(w.LockfilePath())
		if err != nil {
			panic(err)
		}

		log.Print("Sending lockfile data to channel")
		go func() {
			w.Ch <- data // send data to LockfileWatcher channel
			log.Print("Done!")
		}()

		return true, nil
	}

	log.Print("lockfile not found")
	return false, RiotClientLockfileNotFound
}

// Watches the Riot Client lockfile directory and send the lockfile data to LockfileWatcher.channel
func (w *LockfileWatcher) Watch(duration time.Duration) {
	ticker := time.NewTicker(duration)
	log.Println("starting watcher")

	go func() {
		for {
			select {
			case <-ticker.C:
				// scan for lockfile based on d duration
				ok, _ := w.Scan()
				if ok {
					defer ticker.Stop()

					data, err := readLockfile(w.LockfilePath())
					if err != nil {
						panic(err)
					}

					log.Print("sent lockfile file data to channel")
					w.Ch <- data // send data to LockfileWatcher channel
				}
			case <-w.Ch:
				return
			}
		}
	}()
}

// reads the Riot Client lockfile at the given path
func readLockfile(path string) (*RiotClientLockfileInfo, error) {
	res, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lockfileData := strings.Split(string(res), ":")
	if len(lockfileData) != int(5) {
		return nil, unexpectedLockfileData
	}

	name := lockfileData[0]
	PID, _ := strconv.Atoi(lockfileData[1])
	port, _ := strconv.Atoi(lockfileData[2])
	password := lockfileData[3]
	protocol := lockfileData[4]

	return &RiotClientLockfileInfo{
		Name:     name,
		PID:      PID,
		Port:     port,
		Password: password,
		Protocol: protocol,
	}, nil
}
