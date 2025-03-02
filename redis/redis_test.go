package redis_test

import (
	"context"
	"testing"
	"time"

	"github.com/aria-afk/rediscli/redis"
)

func TestNewRedisPass(t *testing.T) {
	primaryUri := "redis://localhost:6379/0"

	ctx := context.Background()
	opts := redis.RedisOpts{
		URI: primaryUri,
	}

	r, err := redis.NewRedis(ctx, opts)
	if err != nil {
		t.Errorf("Error spawning redis connection:\n%s", err)
	}
	expires := time.Millisecond * 1
	err = r.Client.Set(r.Ctx, "test", "1", expires).Err()
	if err != nil {
		t.Errorf("Error running Set on primary client:\n%s", err)
	}
}
