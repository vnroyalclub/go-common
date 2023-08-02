package kafka

import (
	"fmt"
)

var (
	handler = func(msg []byte, args ...interface{}) {
		if len(args) != 1 {
			fmt.Println("arg is empty,")
			return
		}
		value, ok := args[0].(int)
		if !ok {
			fmt.Println("args 1 not type int")
			return
		}
		fmt.Println("value:", value)
		fmt.Println("msg:", string(msg))
	}
)

func ExampleConsumer() {
	fmt.Println("consumer...")
	consumer, err := NewConsumer([]string{"172.13.0.51:9092"}, "zhuma")
	if err != nil {
		fmt.Println("failed to new consumer,err:", err)
		return
	}
	defer consumer.Close()
	err= consumer.Consume(map[string]Handler{
		"test":Handler{
			Run:handler,
			Args:[]interface{}{1},
		},
	})
}
