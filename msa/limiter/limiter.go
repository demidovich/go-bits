package limiter

import (
	"context"
	"errors"
	"time"
)

type Limiter struct {
	bucket chan struct{}
}

type Config struct {
	Interval time.Duration
	Tokens   int
}

func New(ctx context.Context, cfg Config) (*Limiter, error) {
	if cfg.Interval.Microseconds() < 10 {
		return nil, errors.New("the interval cannot be less than 10 microseconds")
	}

	if cfg.Tokens < 1 {
		return nil, errors.New("the tokens cannot be less than 1")
	}

	l := Limiter{
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
	for {
		select {
		case l.bucket <- struct{}{}:
		default:
			return
		}
	}
}
