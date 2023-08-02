package kafka

import (
	"fmt"
)

func ExampleProducer() {
	fmt.Println("producer...")

	producer, err := NewProducer([]string{"172.13.0.51:9092"})
	if err != nil {
		fmt.Println("failed to new producer,err:", err)
		return
	}
	defer producer.Close()

	err=producer.Producer(Message{
		Topic:"test",
		Value:[]byte("wo shi yi tiao yu"),
	})

	if err!=nil{
		fmt.Println("failed to send message,err:",err)
		return
	}
}
