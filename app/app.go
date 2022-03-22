package app

import (
	"context"
	"github.com/gongxulei/go_kit/logger/origin"
	"github.com/gongxulei/go_kit/register"
	"net/url"
	"os"
)

type Application interface {
}

type App struct {
}

type OptionFun func()

type options struct {
	ctx       context.Context
	endpoints []*url.URL
	signals   []os.Signal
	log       origin.LoggerInterface
	register  register.Registry
	discovery register.Discovery
}
