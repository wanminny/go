package my

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

// value是简单的string
func stringValue(ctx context.Context, client *redis.Client) {
	key := "name"
	value := "Sam"
	defer client.Del(ctx, key)

	err := client.Set(ctx, key, value, 1*time.Second).Err()
	checkError(err)

	client.Expire(ctx, key, 3*time.Second)
	time.Sleep(2 * time.Second)

	v2, err := client.Get(ctx, key).Result()
	checkError(err)
	fmt.Println(v2)
}

// value是List
func listValue(ctx context.Context, client *redis.Client) {
	key := "ids"
	defer client.Del(ctx, key)

	values := []interface{}{"S", "A", "M", 8, 8, 4, 8}
	err := client.RPush(ctx, key, values...).Err()
	checkError(err)

	v2, err := client.LRange(ctx, key, 1, 4).Result()
	checkError(err)
	fmt.Println(v2)
}

// value是Set
func setValue(ctx context.Context, client *redis.Client) {
	key := "ids"
	defer client.Del(ctx, key)

	values := []interface{}{"S", "A", "M", 8, 8, 4, 8}
	err := client.SAdd(ctx, key, values...).Err()
	checkError(err)

	var query any
	query = 8
	if client.SIsMember(ctx, key, query).Val() {
		fmt.Println("Set has element:", query)
	} else {
		fmt.Println("Set doesn't have element:", query)
	}
	query = "8"
	if client.SIsMember(ctx, key, query).Val() {
		fmt.Println("Set has element:", query)
	} else {
		fmt.Println("Set doesn't have element:", query)
	}
	query = 5
	if client.SIsMember(ctx, key, query).Val() {
		fmt.Println("Set has element:", query)
	} else {
		fmt.Println("Set doesn't have element:", query)
	}

	for _, e := range client.SMembers(ctx, key).Val() {
		fmt.Println(e)
	}

	key2 := "ids2"
	defer client.Del(ctx, key2)
	values = []interface{}{"S", "U", "N", "8", "9", "6", "6"}
	err = client.SAdd(ctx, key2, values...).Err()
	checkError(err)

	fmt.Println("k1-k2")
	for _, e := range client.SDiff(ctx, key, key2).Val() {
		fmt.Println(e)
	}
	fmt.Println("k2-k1")
	for _, e := range client.SDiff(ctx, key2, key).Val() {
		fmt.Println(e)
	}
	fmt.Println("k1+k2")
	for _, e := range client.SInter(ctx, key, key2).Val() {
		fmt.Println(e)
	}
}

// value是ZSet
func zsetValue(ctx context.Context, client *redis.Client) {
	key := "ids"
	defer client.Del(ctx, key)

	values := []redis.Z{
		{Member: "Sam", Score: 100},
		{Member: "Greenery", Score: 99},
		{Member: "YKY", Score: 101},
	}
	err := client.ZAdd(ctx, key, values...).Err()
	checkError(err)

	for _, e := range client.ZRange(ctx, key, 0, -1).Val() {
		fmt.Println(e)
	}
}

// value是HashTable
func hashtableValue(ctx context.Context, client *redis.Client) {
	student1 := map[string]any{"Name": "Sam", "Gender": 1, "Age": 19}
	student2 := map[string]any{"Name": "Greenery", "Gender": 0.5, "Age": 23}
	checkError(client.HSet(ctx, "s1", student1).Err())
	checkError(client.HSet(ctx, "s2", student2).Err())

	age, err := client.HGet(ctx, "s2", "Name").Result()
	checkError(err)
	fmt.Println(age)

	for field, value := range client.HGetAll(ctx, "s1").Val() {
		fmt.Println(field, value)
	}

	client.Del(ctx, "s1")
	client.Del(ctx, "s2")
}

func checkError(err error) {
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exist")
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
