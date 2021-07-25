module github.com/xinlianit/kit-scaffold

go 1.13

require (
	gitee.com/jirenyou/business.palm.proto v1.0.0-2021021301
	gitee.com/plam-bfa/bfa-gateway-api v0.0.0-20210718080622-626de8ecde17
	github.com/Shopify/sarama v1.28.0
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/fsnotify/fsnotify v1.4.9
	github.com/go-kit/kit v0.10.0
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.7.3
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/hashicorp/consul/api v1.8.1
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/jinzhu/copier v0.2.3
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/miekg/dns v1.1.27 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/nacos-group/nacos-sdk-go v1.0.8
	github.com/prometheus/client_golang v1.3.0
	github.com/robfig/cron/v3 v3.0.0 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.8.1
	github.com/xinlianit/go-util v1.0.0-2021072301
	go.uber.org/zap v1.18.1
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/guregu/null.v4 v4.0.0 // indirect
	gorm.io/driver/mysql v1.1.1
	gorm.io/gorm v1.21.12
)

// 工具包
replace github.com/xinlianit/go-util => github.com/xinlianit/go-util v1.0.0-2021072301
