package roundrobin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Next(t *testing.T) {
	servers := []string{
		"server1",
		"server2",
	}

	rr := New(servers)

	assert.Equal(t, "server2", rr.Next())
	assert.Equal(t, "server1", rr.Next())
	assert.Equal(t, "server2", rr.Next())
	assert.Equal(t, "server1", rr.Next())
}
