package monitor

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type HttpChecker struct {
	Meta                 *Meta
	ExpectedResponseCode int
	Request              *http.Request
	Client               *http.Client
}

func (c HttpChecker) GetMeta() *Meta {
	return c.Meta
}

func (c HttpChecker) DoCheck(ctx context.Context) *Check {
	now := time.Now()
	res, err := c.Client.Do(c.Request)
	duration := time.Since(now)

	if err != nil {
		return &Check{
			Meta:      c.Meta,
			IsSuccess: false,
			Result:    ErrorResult{err.Error()},
			CheckedAt: now,
			Duration:  duration,
		}
	} else if c.ExpectedResponseCode != 0 && res.StatusCode != c.ExpectedResponseCode {
		return &Check{
			Meta:      c.Meta,
			IsSuccess: false,
			Result: ErrorResult{
				fmt.Sprintf("expected %d, but got %d response", c.ExpectedResponseCode, res.StatusCode),
			},
			CheckedAt: now,
			Duration:  duration,
		}
	}
	return &Check{
		Meta:      c.Meta,
		IsSuccess: true,
		Result: map[string]interface{}{
			"status":         res.StatusCode,
			"content-length": res.ContentLength,
		},
		CheckedAt: now,
		Duration:  duration,
	}
}

func NewSimpleHttpChecker(name, description, method, endpoint string, interval, timeout time.Duration) Checker {
	u, _ := url.Parse(endpoint)
	return HttpChecker{
		Meta: &Meta{
			Name:        name,
			Description: description,
			Interval:    interval,
		},
		Request: &http.Request{
			Method: method,
			URL:    u,
		},
		Client: &http.Client{
			Timeout: timeout,
		},
	}
}
