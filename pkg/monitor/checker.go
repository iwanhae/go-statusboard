package monitor

import (
	"context"
	"time"
)

type Meta struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Interval    time.Duration `json:"interval"`
}

type Check struct {
	//Meta Metainfo of this Check
	Meta *Meta
	//IsSuccess true if success
	IsSuccess bool
	//Result contains detail of success or failiure
	Result interface{}

	CheckedAt time.Time
	Duration  time.Duration
}

type Checker interface {
	DoCheck(ctx context.Context) *Check
	GetMeta() *Meta
}

type ErrorResult struct {
	Error string `json:"error"`
}
