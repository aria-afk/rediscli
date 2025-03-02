package redis

import (
	"context"
	"fmt"

	goredis "github.com/redis/go-redis/v9"
)

// TODO: Replicants

// Client struct stores context from cobra and exposes all clients
type Redis struct {
	Ctx    context.Context
	Client *goredis.Client
}

// NOTE: This is a first pass and assumes we will be building proper URI's
// priority will be standard cli->env->defaults

// Options for connecting to redis; these are provided via cli-args
// and are constructed in the cmd/root.go Run function
type RedisOpts struct {
	URI string
}

// Return a new instance of a RedisStruct or an error if *any* connection failed.
func NewRedis(ctx context.Context, opts RedisOpts) (*Redis, error) {
	r := &Redis{Ctx: ctx}

	// Primary conn
	parsedopts, err := goredis.ParseURL(opts.URI)
	if err != nil {
		return r, fmt.Errorf("Error connecting to client\nMessage: %s", err)
	}
	r.Client = goredis.NewClient(parsedopts)
	return r, nil
}
