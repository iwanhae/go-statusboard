package monitor

import (
	"context"
	"fmt"
	"time"
)

func safeDoCheck(ctx context.Context, c Checker) (result *Check) {
	meta := c.GetMeta()
	now := time.Now()
	defer func() {
		if v := recover(); v != nil {
			result = &Check{
				Meta:      meta,
				IsSuccess: false,
				Result:    fmt.Sprintf("panic occuers:%v", v),
				CheckedAt: now,
				Duration:  time.Since(now),
			}
		}
	}()

	return c.DoCheck(ctx)
}
