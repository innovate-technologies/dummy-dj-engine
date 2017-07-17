package dummyengine

import (
	"fmt"
	"io"

	"github.com/chuckpreslar/emission"
	"github.com/innovate-technologies/DJ/data"
)

// Engine is the dummy engine
type Engine struct {
}

// New  returns a new engine
func New() *Engine {
	return &Engine{}
}

// Start starts the engine
func (e *Engine) Start(events *emission.Emitter) {
	fmt.Println("Started the engine")
}

// PutQueue replaced the queue
func (e *Engine) PutQueue(songs []data.Song) {
	for _, song := range songs {
		fmt.Println("queueing", song.ID)
	}
}

// Skip skips the current next song
func (e *Engine) Skip() {
	fmt.Println("Skipped song")
}

// PipeLive piped the reader to the stream
func (e *Engine) PipeLive(r io.Reader) {
	fmt.Println("Piping a live stream")
}

// SetMetadata sets the stream metadata
func (e *Engine) SetMetadata(song string) {
	fmt.Println("Setting metadata", song)
}
