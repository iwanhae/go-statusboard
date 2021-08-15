package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/iwanhae/go-statusboard/pkg/monitor"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

type Meta struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Interval    string `json:"interval"`
}
type Stream struct {
	Name      string      `json:"name"`
	IsSuccess bool        `json:"is_success"`
	CheckedAt time.Time   `json:"checked_at"`
	Result    interface{} `json:"result"`
	Duration  int64       `json:"duration"`
}

func GetMetaHandler(looker *monitor.Looker) echo.HandlerFunc {
	var result []Meta
	for _, v := range looker.CheckList {
		meta := v.GetMeta()
		result = append(result, Meta{
			Name:        meta.Name,
			Description: meta.Description,
			Interval:    meta.Interval.String(),
		})
	}
	return func(c echo.Context) error {
		c.JSON(200, result)
		return nil
	}
}

func GetStreamHandler(looker *monitor.Looker) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		log.Ctx(ctx).Info().Msg("New Requests")
		ch := make(chan *monitor.Check, 10)
		looker.Subscribe(c.RealIP(), ch)

		r := c.Response()
		r.Header().Set("Content-Type", "text/event-stream")
		r.Header().Set("Cache-Control", "no-cache")
		r.Header().Set("Content-Encoding", "none")
		r.Header().Set("Access-Control-Allow-Origin", "*")
		r.WriteHeader(http.StatusOK)

		enc := json.NewEncoder(r)
	loop:
		for {
			select {
			case check := <-ch:
				fmt.Fprint(r, "data: ")
				enc.Encode(Stream{
					Name:      check.Meta.Name,
					IsSuccess: check.IsSuccess,
					CheckedAt: check.CheckedAt,
					Result:    check.Result,
					Duration:  check.Duration.Milliseconds(),
				})
				fmt.Fprint(r, "\n\n")
			case <-time.After(300 * time.Millisecond):
				r.Flush()
			case <-ctx.Done():
				log.Ctx(ctx).Info().Msg("Catch Context Closed")
				break loop
			case <-r.CloseNotify():
				log.Ctx(ctx).Info().Msg("Catch Response Closed")
				break loop
			}
		}
		looker.UnSubscribe(ch)
		log.Ctx(ctx).Info().Msg("Close Channel")
		close(ch)

		return nil
	}
}
