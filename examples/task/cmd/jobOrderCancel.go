/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
	"github.com/xinlianit/kit-scaffold/examples/task/consumer"
	"sync"
)

var (
	wg sync.WaitGroup
)

// jobOrderCancelCmd represents the jobOrderCancel command
var jobOrderCancelCmd = &cobra.Command{
	Use:   "jobOrderCancel",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("jobOrderCancel called")

		addrs := []string{"39.100.153.72:9092"}
		groupID := "scaffold"
		topic := "kit_examples"

		// 创建消费者组（消费者组）
		consumerGroup, err := sarama.NewConsumerGroup(addrs, groupID, consumer.NewKafkaConfig())
		if err != nil {
			fmt.Printf("消费者组启动失败 err:%v\n", err)
			return
		}

		for{
			consumerGroupHandler := new(consumer.ConsumerGroupHandler)

			if err := consumerGroup.Consume(context.Background(), []string{topic}, consumerGroupHandler); err != nil {
				fmt.Printf("消费者组消费失败 err:%v\n", err)
				return
			}

			fmt.Println("re balance")
		}

		// 创建消费者（消费者）
		//cons, err := sarama.NewConsumer(addrs, nil)
		//if err != nil {
		//	fmt.Printf("消费者启动失败 err:%v\n", err)
		//	return
		//}
		//
		//// 根据topic取到所有的分区
		//partitionList, err := cons.Partitions(topic)
		//if err != nil {
		//	fmt.Printf("分区获取失败 err:%v\n", err)
		//	return
		//}
		//
		//log.Println("分区列表：", partitionList)

		//// 遍历分区
		//for partition := range partitionList{
		//	// 针对每个分区创建一个对应的分区消费者
		//	pc, err := cons.ConsumePartition("kit_examples", int32(partition), sarama.OffsetNewest)
		//	if err != nil {
		//		fmt.Printf("分区消费者创建失败 err:%v\n", err)
		//	}
		//	defer pc.AsyncClose()
		//
		//	wg.Add(1)
		//
		//	// 异步从每个分区消费信息
		//	go func(sarama.PartitionConsumer) {
		//		defer wg.Done()
		//		for msg := range pc.Messages() {
		//			fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		//		}
		//	}(pc)
		//}

		//wg.Wait()
		//cons.Close()
	},
}

func init() {
	jobCmd.AddCommand(jobOrderCancelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jobOrderCancelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jobOrderCancelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
