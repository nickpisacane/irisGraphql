package irisGraphql

import (
	"github.com/graphql-go/graphql"
	"github.com/imdario/mergo"
	"github.com/kataras/iris"
	"golang.org/x/net/context"
)

// ContextFunc a func that takes an *iris.Context, and returns a context.Context.
type ContextFunc func(*iris.Context) context.Context

// Config configuration for middleware.
type Config struct {
	Schema      graphql.Schema
	GraphiQL    bool
	ContextFunc ContextFunc
}

// DefaultConfig provides a Config instance with a ContextFunc that provides
// a context.Context with access to the iris.Context.
func DefaultConfig() Config {
	return Config{
		ContextFunc: func(ctx *iris.Context) context.Context {
			return context.WithValue(context.Background(), "ctx", ctx)
		},
	}
}

// Merge merges a config.
func (c Config) Merge(cfg []Config) (config Config) {
	if len(cfg) > 0 {
		config = cfg[0]
		mergo.Merge(&config, c)
	} else {
		config = c
	}
}
