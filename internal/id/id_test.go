package id_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stuckinforloop/dist/internal/id"
)

var (
	testIDOne = "c1n18717895732742165505"
)

func TestID(t *testing.T) {
	clientID := "c1"
	nodeID := "n1"

	source := rand.NewSource(0)
	idSource := id.New(&source)
	id := idSource.Generate(clientID, nodeID)
	assert.Equal(t, testIDOne, id)
}
