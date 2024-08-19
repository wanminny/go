package my

import (
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"testing"
)

func TestPubSub(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	})
	c := context.Background()
	pubsub(c, client)
}
