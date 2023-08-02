package game_redis

import (
	"context"
	"fmt"
	"time"

	"github.com/vnroyalclub/go-common/game-redis/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func NewGameRedisClient(target string) (client pb.GameRedisClient, err error) {
	coon, err := grpc.DialContext(context.Background(), target,
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                time.Second * 30,
			Timeout:             time.Second * 30,
			PermitWithoutStream: true,
		}),
		//grpc.WithBlock(),
		//grpc.WithTimeout(time.Second * 30),
	)

	if err != nil {
		fmt.Println("[EROR] failed to new game redis client,err:", err)
		return
	}

	client = pb.NewGameRedisClient(coon)
	return
}
