package monitor

import (
	"context"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHttpChecker(t *testing.T) {
	t.Parallel()

	t.Run("Success", func(t *testing.T) {
		u, _ := url.Parse("https://google.com")
		h := HttpChecker{
			Meta: &Meta{
				Name:        "Test",
				Description: "Test",
				Interval:    time.Second,
			},
			Request: &http.Request{
				Method: "GET",
				URL:    u,
			},
			Client: &http.Client{
				Timeout: 5 * time.Second,
			},
		}

		now := time.Now()
		c := h.DoCheck(context.Background())
		assert.Equal(t, "Test", c.Meta.Name)
		assert.True(t, c.CheckedAt.After(now))
		assert.Greater(t, c.Duration, 0*time.Second)
		assert.True(t, c.IsSuccess)
	})

	t.Run("Fail", func(t *testing.T) {
		u, _ := url.Parse("https://non-exists-url.com")
		h := HttpChecker{
			Meta: &Meta{
				Name:        "Test",
				Description: "Test",
				Interval:    time.Second,
			},
			Request: &http.Request{
				Method: "GET",
				URL:    u,
			},
			Client: &http.Client{
				Timeout: 5 * time.Second,
			},
		}

		now := time.Now()
		c := h.DoCheck(context.Background())
		assert.Equal(t, "Test", c.Meta.Name)
		assert.True(t, c.CheckedAt.After(now))
		assert.Greater(t, c.Duration, 0*time.Second)
		assert.False(t, c.IsSuccess)
	})

}
