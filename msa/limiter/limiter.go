package limiter

import (
	"context"
	"errors"
	"time"
)

var (
	defaultTokens   = 100
	defaultInterval = 1 * time.Second
)

type Limiter struct {
	tokens int
	bucket chan struct{}
}

type Config struct {
	Interval time.Duration
	Tokens   int
}

func SetDefaultTokens(val int) {
	defaultTokens = val
}

func SetDefaultInterval(val time.Duration) {
	defaultInterval = val
}

func NewDefault(ctx context.Context) (*Limiter, error) {
	l, err := New(ctx, Config{
		Interval: defaultInterval,
		Tokens:   defaultTokens,
	})

	return l, err
}

func New(ctx context.Context, cfg Config) (*Limiter, error) {
	if cfg.Interval.Microseconds() < 10 {
		return nil, errors.New("the interval cannot be less than 10 microseconds")
	}

	if cfg.Tokens < 1 {
		return nil, errors.New("the tokens cannot be less than 1")
	}

	l := Limiter{
		tokens: cfg.Tokens,
		bucket: make(chan struct{}, cfg.Tokens),
	}

	l.initRefiller(ctx, cfg.Interval)
	l.refill()

	return &l, nil
}

func (l *Limiter) Allow() bool {
	select {
	case <-l.bucket:
		return true
	default:
		return false
	}
}

func (l *Limiter) initRefiller(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		defer ticker.Stop()
		for range ticker.C {
			select {
			case <-ctx.Done():
				return
			default:
				l.refill()
			}
		}
	}()
}

func (l *Limiter) refill() {
	for range l.tokens {
		select {
		case l.bucket <- struct{}{}:
		default:
			return
		}
	}
}
