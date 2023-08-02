module github.com/vnroyalclub/go-common

go 1.12

require (
	git.huoys.com/vn/proto v0.0.0-20201121110355-00205aba2644
	github.com/Shopify/sarama v1.23.1
	github.com/fsnotify/fsnotify v1.4.7
	github.com/golang/protobuf v1.4.0
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/spf13/viper v1.6.2
	github.com/yougg/log4go v0.0.0-20170306024712-5c831e7a9a40
	google.golang.org/grpc v1.26.0
	google.golang.org/protobuf v1.22.0 // indirect
	gopkg.in/jcmturner/goidentity.v3 v3.0.0 // indirect
)

replace git.huoys.com/vn/proto => github.com/vnroyalclub/go_proto v0.0.0-20230213101009-85b4ebd503db
