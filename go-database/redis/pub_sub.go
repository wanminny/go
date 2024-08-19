package my

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func publish(c context.Context, client *redis.Client, chName string, msg any) {
	cmd := client.Publish(c, chName, msg)
	if cmd.Err() == nil {
		n := cmd.Val()
		fmt.Printf("%s sent msg %+v into channel %s, %d listeners\n", c.Value("publisher"), msg, chName, n)
	} else {
		fmt.Printf("%s sent msg %+v into channel %s failed\n", c.Value("publisher"), msg, chName)
	}
}

func subscribe(c context.Context, client *redis.Client, chNames []string) {
	ps := client.Subscribe(c, chNames...)
	defer ps.Close()

	for {
		if msg, err := ps.ReceiveMessage(c); err != nil {
			fmt.Println(msg)
			break
		} else {
			fmt.Printf("%s gets a msg %+v from %s\n", c.Value("subscriber"), msg.Payload, msg.Channel)
		}
	}
}

func pubsub(c context.Context, client *redis.Client) {
	// publishers
	c1 := context.WithValue(c, "publisher", "p1")
	c2 := context.WithValue(c, "publisher", "p2")
	chName1, chName2 := "ch1", "ch2"

	c3 := context.WithValue(c, "subscriber", "s1")
	c4 := context.WithValue(c, "subscriber", "s2")

	go subscribe(c3, client, []string{chName1})
	go subscribe(c4, client, []string{chName2})
	time.Sleep(time.Second)

	go publish(c1, client, chName1, "兄弟")
	go publish(c1, client, chName1, "香草泥")
	time.Sleep(time.Second)

	c5 := context.WithValue(c, "subscriber", "s3")
	go subscribe(c5, client, []string{chName1, chName2})
	time.Sleep(time.Second)

	go publish(c2, client, chName2, "Bro")
	go publish(c2, client, chName2, "You are so sweet")
	time.Sleep(time.Second)
}
