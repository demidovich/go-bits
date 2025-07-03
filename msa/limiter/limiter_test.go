package limiter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_ConfigBadInterval(t *testing.T) {
	_, err := newTestingLimiter(t, 1*time.Microsecond, 10)

	assert.NotNil(t, err)
}

func Test_ConfigBadTokensCount(t *testing.T) {
	_, err := newTestingLimiter(t, 10*time.Millisecond, 0)

	assert.NotNil(t, err)
}

func Test_Allow(t *testing.T) {
	l, _ := newTestingLimiter(t, 1*time.Second, 2)

	ok1 := l.Allow()
	ok2 := l.Allow()
	ok3 := l.Allow()

	assert.True(t, ok1)
	assert.True(t, ok2)
	assert.False(t, ok3)
}

func Test_Refill(t *testing.T) {
	l, _ := newTestingLimiter(t, 10*time.Millisecond, 2)
	l.Allow()
	l.Allow()

	ok1 := l.Allow()
	assert.False(t, ok1)

	time.Sleep(15 * time.Millisecond)

	ok2 := l.Allow()
	assert.True(t, ok2)
}

func newTestingLimiter(t *testing.T, interval time.Duration, tokens int) (*Limiter, error) {
	cfg := Config{
		Interval: interval,
		Tokens:   tokens,
	}

	l, err := New(t.Context(), cfg)
	return l, err
}
