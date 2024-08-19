package my

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"sync"
	"time"
)

func TryLock(rc *redis.Client, key string, expire time.Duration) bool {
	cmd := rc.SetNX(context.Background(), key, "anything is ok", expire)
	if err := cmd.Err(); err == nil {
		return cmd.Val()
	} else {
		return false
	}
}

func ReleaseLock(rc *redis.Client, key string) {
	rc.Del(context.Background(), key)
}

func LockRace(client *redis.Client) {
	key := "iPhone 2/3"
	defer ReleaseLock(client, key)
	const P = 100
	wg := sync.WaitGroup{}
	wg.Add(P)
	for i := 0; i < P; i++ {
		go func(i int) {
			defer wg.Done()
			if TryLock(client, key, time.Hour) {
				fmt.Printf("%d got the iPhone!!\n", i)
			} else {
				fmt.Println(i, "is locked")
			}
		}(i)
	}
	wg.Wait()
}

func LockRace2(client *redis.Client, storage int) {
	keyLock := "lock"
	keyStorage := "store"
	client.Set(context.Background(), keyStorage, storage, 0)
	defer ReleaseLock(client, keyLock)
	defer client.Del(context.Background(), keyStorage)

	const P = 1000
	wg := sync.WaitGroup{}
	wg.Add(P)
	start := time.Now()
	for i := 0; i < P; i++ {
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			if TryLock(client, keyLock, 0) {
				if v := client.IncrBy(context.Background(), keyStorage, -1).Val(); v >= 0 {
					fmt.Printf("%d gets the No.%d iPhone! Use %v!\n", i, v+1, time.Now().Sub(start))
				}
				ReleaseLock(client, keyLock)
			}
		}(i)
	}
	wg.Wait()
}
