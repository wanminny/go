package my

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestAll(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	})
	c := context.Background()

	t.Run("Test String Value", func(t *testing.T) {
		stringValue(c, client)
	})

	t.Run("Test List Value", func(t *testing.T) {
		listValue(c, client)
	})

	t.Run("Test Set Value", func(t *testing.T) {
		setValue(c, client)
	})

	t.Run("Test ZSet Value", func(t *testing.T) {
		zsetValue(c, client)
	})
	
	t.Run("Test Hashmap Value", func(t *testing.T) {
		hashtableValue(c, client)
	})
}
