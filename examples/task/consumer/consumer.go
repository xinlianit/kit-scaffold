package consumer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

// 消费者组处理器
type ConsumerGroupHandler struct {

}

func (h ConsumerGroupHandler) Setup(s sarama.ConsumerGroupSession) error {
	fmt.Println("set up ....") // 当连接完毕的时候会通知这个，start
	return nil
}

func (h ConsumerGroupHandler) Cleanup(s sarama.ConsumerGroupSession) error {
	fmt.Println("Cleanup") // end，当这一次消费完毕，会通知，这里最好commit
	return nil
}
func (h ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error { // consume
	for msg := range claim.Messages() { // 接受topic消息
		fmt.Printf("[Consumer] Message topic:%q partition:%d offset:%d add:%d key:%s value:%s\n",
			msg.Topic, msg.Partition, msg.Offset, claim.HighWaterMarkOffset()-msg.Offset,
			string(msg.Key), string(msg.Value))
		sess.MarkMessage(msg, "") // 必须设置这个，不然你的偏移量无法提交。
	}
	return nil
}

// 创建kafka配置
func NewKafkaConfig() *sarama.Config {
	config := sarama.NewConfig()
	//config.ClientID = "sarama_demo" //
	//config.Version = sarama.V0_11_0_1 // kafka server的版本号
	config.Producer.Return.Successes = true // sync必须设置这个
	config.Producer.RequiredAcks = sarama.WaitForAll // 也就是等待foolower同步，才会返回
	config.Producer.Return.Errors = true
	config.Consumer.Return.Errors = true
	config.Metadata.Full = false // 不用拉取全部的信息
	config.Consumer.Offsets.AutoCommit.Enable = true // 自动提交偏移量，默认开启，说时候，我没找到手动提交。
	config.Consumer.Offsets.AutoCommit.Interval = time.Second // 这个看业务需求，commit提交频率，不然容易down机后造成重复消费。
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // 从最开始的地方消费，业务中看有没有需求，新业务重跑topic。
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange // rb策略，默认就是range
	return config
}
