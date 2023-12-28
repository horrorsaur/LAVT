package local

import (
	"io"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO:
// Refactor LockfileWatcher to accept input to support custom Valorant installations

func TestLockfileWatcher_Watch(t *testing.T) {
	log.SetOutput(io.Discard)
	w := LockfileWatcher{

		Ch:       make(chan *RiotClientLockfileInfo),
		CacheDir: "../fixtures/",
	}

	ok, err := w.Scan()
	assert.True(t, ok)
	assert.Nil(t, err)
}
