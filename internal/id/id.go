package id

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// Source represents a source of unique identifiers. It contains a random number
// generator and a counter to generate monotonically increasing values.
type Source struct {
	rnd     *rand.Rand
	counter int64
}

// New creates a new Source instance with a random number generator.
// If source is not nil, it uses the provided rand.Source. Otherwise, it creates a new
// rand.Source using the current time.
// The returned Source has a counter initialized to 0.
func New(source *rand.Source) *Source {
	var rnd *rand.Rand

	if source != nil {
		rnd = rand.New(*source)
	} else {
		rnd = rand.New(rand.NewSource(time.Now().Unix()))
	}

	id := &Source{
		rnd:     rnd,
		counter: 0,
	}

	return id
}

// Generate generates a unique identifier string using the provided clientID and nodeID.
// It combines the clientID and nodeID with a monotonically increasing counter value
// to produce a unique identifier.
func (id *Source) Generate(clientID, nodeID string) string {
	delta := id.rnd.Int63()
	val := atomic.AddInt64(&id.counter, delta)

	return fmt.Sprintf("%s%s%d", clientID, nodeID, val)
}
