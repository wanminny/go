package my

import (
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestLockRace(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	})
	t.Run("LockRace", func(t *testing.T) {
		LockRace(client)
	})
	t.Run("LockRace2", func(t *testing.T) {
		LockRace2(client, 10)
		/*
			244 gets the No.10 iPhone! Use 21.982042ms!
			767 gets the No.9 iPhone! Use 23.574209ms!
			399 gets the No.8 iPhone! Use 24.568542ms!
			390 gets the No.7 iPhone! Use 25.760834ms!
			537 gets the No.6 iPhone! Use 26.533334ms!
			314 gets the No.5 iPhone! Use 27.388375ms!
			649 gets the No.4 iPhone! Use 28.493209ms!
			354 gets the No.3 iPhone! Use 29.462459ms!
			353 gets the No.2 iPhone! Use 30.193125ms!
			196 gets the No.1 iPhone! Use 30.868209ms!
		*/
	})

}
