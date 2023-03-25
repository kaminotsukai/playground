package goroutine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parallel(t *testing.T) {
	assert.Equal(t, SyncAtmicFunc(), int32(50005000))
	assert.Equal(t, SyncMutexFunc(), int32(50005000))
}
