package ratelimit

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
)

// Limiter ..
type Limiter interface {
	Take(context.Context) time.Duration
}

// Take ..
func Take(ctx context.Context, l Limiter) {
	throttled := l.Take(ctx)

	for throttled.Milliseconds() < -10 {
		log.WithFields(
			log.Fields{
				"for": throttled.String(),
			},
		).Debug("throttled GitLab requests")

		throttled = l.Take(ctx)
	}
}
