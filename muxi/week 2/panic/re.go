package main

import (
	"fmt"
	"time"
)

type message struct {
	Topic     string
	Partition int32
	Offset    int64
}

type FeedEventDM struct {
	Type    string
	UserID  int
	Title   string
	Content string
}

type MSG struct {
	ms        message
	feedEvent FeedEventDM
}

const ConsumeNum = 5

func main() {
	var consumeMSG []MSG
	var lastConsumeTime time.Time
	msgs := make(chan MSG)

	go func() {
		for i := 0; ; i++ {
			msgs <- MSG{
				ms: message{Topic: "消费主题", Partition: 0, Offset: 0},
				feedEvent: FeedEventDM{
					Type: "grade", UserID: i, Title: "成绩提醒", Content: "您的成绩是xxx",
				},
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for msg := range msgs {
		consumeMSG = append(consumeMSG, msg)

		if len(consumeMSG) >= ConsumeNum {
			// 改动 1：先复制，再异步消费
			batch := make([]MSG, ConsumeNum)
			copy(batch, consumeMSG[:ConsumeNum])
			go fn(batch)

			lastConsumeTime = time.Now()
			consumeMSG = consumeMSG[ConsumeNum:]
		} else if !lastConsumeTime.IsZero() && time.Since(lastConsumeTime) > 5*time.Minute {
			if len(consumeMSG) > 0 {
				// 改动 2：批大小用实际长度 n，先复制再消费
				n := len(consumeMSG)
				batch := make([]MSG, n)
				copy(batch, consumeMSG[:n])
				go fn(batch)

				lastConsumeTime = time.Now()
				consumeMSG = consumeMSG[n:] // 原来是 [ConsumeNum:]，会越界
			}
		}
	}
}

func fn(m []MSG) {
	fmt.Printf("本次消费了%d条消息\n", len(m))
}
