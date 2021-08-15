package main

import (
	"context"
	"os"

	"github.com/iwanhae/go-statusboard/pkg/config"
	"github.com/iwanhae/go-statusboard/pkg/monitor"
	"github.com/iwanhae/go-statusboard/pkg/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = zerolog.
		New(os.Stderr).With().
		Timestamp().
		Caller().
		Logger().
		Output(zerolog.ConsoleWriter{Out: os.Stderr})
	ctx := log.Logger.WithContext(context.Background())

	checkList, err := config.LoadConfig("./config.json")
	if err != nil {
		panic(err)
	}

	l := monitor.Looker{
		CheckList: checkList,
		Cache:     make([]*monitor.Check, 1000),
	}
	log.Ctx(ctx).Info().Msg("start monitoring")
	l.Start(ctx)

	e := server.NewServer(&l)
	log.Ctx(ctx).Info().Msg("go-statusboard started")
	e.Start(":80")

}
