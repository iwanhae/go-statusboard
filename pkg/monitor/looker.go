package monitor

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rs/zerolog/log"
)

type Looker struct {
	CheckList []Checker
	Cache     []*Check
	ptr       int32

	ctx        context.Context
	cancel     context.CancelFunc
	subscriber map[chan *Check]string
	isRunning  bool
	mu         sync.Mutex   // Mutex for isRunning
	rwMu       sync.RWMutex // Mutext for subscriber
}

func (l *Looker) IsRunning() bool {
	return l.isRunning
}
func (l *Looker) Stop() {
	l.cancel()
	// make sure not to run mutiple times
	l.mu.Lock()
	l.isRunning = false
	l.mu.Unlock()
}

func (l *Looker) Start(ctx context.Context) error {
	l.ctx, l.cancel = context.WithCancel(ctx)

	if l.IsRunning() {
		return nil
	}
	// make sure not to run mutiple times
	l.mu.Lock()
	l.isRunning = true
	l.mu.Unlock()

	for _, checker := range l.CheckList {
		c := checker
		interval := checker.GetMeta().Interval
		go func() {
			for {
				if !l.isRunning {
					break
				}
				next := time.Now().Add(interval)
				l.notify(
					safeDoCheck(l.ctx, c),
				)
				rest := time.Until(next)
				select {
				case <-l.ctx.Done():
					log.Ctx(l.ctx).Info().Msg("looker loop exited")
					return
				case <-time.After(rest):
				}
			}
		}()
	}
	return nil
}

func (l *Looker) notify(check *Check) {
	l.rwMu.RLock()
	defer l.rwMu.RUnlock()
	log.Ctx(l.ctx).Info().Interface("check", check).Msg("notify")

	// Caching Logic
	if length := cap(l.Cache); length != 0 {
		i := atomic.AddInt32(&l.ptr, 1) - 1
		i = i % int32(length)
		l.Cache[i] = check
	}

	for k, v := range l.subscriber {
		select {
		case k <- check:
		default:
			log.Ctx(l.ctx).Error().Str("channel_id", v).Msg("fail to push Check to channel")
		}
	}
}

func (l *Looker) Subscribe(id string, c chan *Check) {
	l.rwMu.Lock()
	defer l.rwMu.Unlock()

	log.Ctx(l.ctx).Info().Str("id", id).Msg("new subscriber")
	if l.subscriber == nil {
		l.subscriber = make(map[chan *Check]string)
	}
	l.subscriber[c] = id
	go func() {
		for _, v := range l.Cache {
			if v == nil {
				break
			}
			c <- v
		}
	}()
}

func (l *Looker) UnSubscribe(c chan *Check) {
	id, ok := l.subscriber[c]
	if !ok {
		return
	}

	l.rwMu.Lock()
	defer l.rwMu.Unlock()
	log.Ctx(l.ctx).Info().Str("id", id).Msg("unsubscribe")
	delete(l.subscriber, c)
}
