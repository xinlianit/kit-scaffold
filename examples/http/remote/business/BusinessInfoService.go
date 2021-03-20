package business

import (
	"context"
	"gitee.com/jirenyou/business.palm.proto/pb/go/service"
	"gitee.com/jirenyou/business.palm.proto/pb/go/transport/request"
	"gitee.com/jirenyou/business.palm.proto/pb/go/transport/response"
	"github.com/Shopify/sarama"
	"google.golang.org/grpc/status"
	"log"
)

type businessInfoService struct {
	client service.BusinessInfoServiceClient
}

func (s businessInfoService) GetBusinessInfo(ctx context.Context) (*response.GetBusinessInfoResponse, error) {
	req := &request.GetBusinessInfoRequest{
		BusinessId: 99,
	}
	rsp, err := s.client.GetBusinessInfo(ctx, req)

	// 发送kafka
	kafkaConfig := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	// 随机向partition发送消息
	kafkaConfig.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Producer.Return.Errors = true

	// 创建生产者
	kafkaAddrs := []string{"39.100.153.72:9092"}
	producer, err := sarama.NewSyncProducer(kafkaAddrs, kafkaConfig)
	if err != nil {
		log.Printf("kafka连接失败->error: %v", err)
	}else{
		defer producer.Close()

		// 发送消息(同步方式)
		msg := &sarama.ProducerMessage{
			Topic: "kit_examples",
			Key: sarama.StringEncoder("test_key"),
			Value: sarama.StringEncoder("测试kafka"),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("消息发送失败->error: %v", err)
		}
		log.Printf("消息发送成功->partition: %d, offset: %d", partition, offset)
	}

	if err != nil {
		if rsp, ok := status.FromError(err); ok {
			log.Printf("RPC 错误: code: %d, message: %s", rsp.Proto().GetCode(), rsp.Proto().GetMessage())
		}

		return nil, err
	}

	return rsp, nil
}
