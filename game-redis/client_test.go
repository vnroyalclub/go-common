package game_redis

import (
	"context"
	"testing"

	"github.com/vnroyalclub/go-common/game-redis/pb"
)

func TestNewGameRedisClient(t *testing.T) {
	/*	client, err := NewGameRedisClient("localhost:1981")
		if err != nil {
			t.Fatal("failed to new game redis client,err:", err)
		}
	*/

}

//测试字符串类型
func TestString(t *testing.T) {
	client, err := NewGameRedisClient("172.13.0.53:1981")
	if err != nil {
		t.Fatal("failed to new game redis client,err:", err)
	}

	var (
		pbKey = pb.Key{
			GameId: 1002,
			Key:    "test",
		}
		value = "hello world"
	)
	//set
	errMsg, err := client.Set(context.Background(), &pb.SetRequest{
		Key:   &pbKey,
		Value: value,
	})

	if err != nil || errMsg.ErrCode != pb.ErrCode_Sucess {
		t.Fatal("failed to set key value,err:", err)
	}

	getresp, err := client.Get(context.Background(), &pbKey)

	if err != nil || getresp.Value != value {
		t.Fatal("failed to get key value,err:", err)
	}

	errMsg, err = client.Expire(context.Background(), &pb.ExpireRequest{
		Key:    &pbKey,
		Type:   pb.DataType_String,
		Second: 300,
	})

	if err != nil || errMsg.ErrCode != pb.ErrCode_Sucess {
		t.Fatal("failed to expire err:", err)
	}

	errMsg, err = client.Del(context.Background(), &pb.DelRequest{
		Key:  &pbKey,
		Type: pb.DataType_Set,
	})

	if err != nil || errMsg.ErrCode != pb.ErrCode_Sucess {
		t.Fatal("failed to del err:", err)
	}
}

//测试hash类型
func TestHash(t *testing.T) {
	client, err := NewGameRedisClient("172.13.0.53:1981")
	if err != nil {
		t.Fatal("failed to new game redis client,err:", err)
	}

	var (
		pbKey = pb.Key{
			GameId: 1002,
			Key:    "test",
		}
		filedValue = map[string]string{
			"zhang": "san",
			"li":    "si",
			"wang":  "wu",
			"rank":  "2",
		}
	)
	//set
	errMsg, err := client.HmSet(context.Background(), &pb.HmSetRequest{
		Key:         &pbKey,
		FieldValues: filedValue,
	})

	if err != nil || errMsg.ErrCode != pb.ErrCode_Sucess {
		t.Fatal("failed to hmset,err:", err)
	}

	getresp, err := client.HmGet(context.Background(), &pb.HmGetRequest{
		Key:    &pbKey,
		Fields: []string{"zhang", "li"},
	})

	if err != nil {
		t.Fatal("failed to get key value,err:", err)
	}

	t.Log("hmget value:", getresp.FieldValues)
	for k, v := range getresp.FieldValues {
		if filedValue[k] != v {
			t.Fatal("failed to hmget")
		}
	}

	ret, err := client.HGetAll(context.Background(), &pbKey)
	if err != nil {
		t.Fatal("failed to hget all,err:", err)
	}

	t.Log("hgetall value:", ret.GetFieldValues())
	for k, v := range ret.FieldValues {
		if filedValue[k] != v {
			t.Fatal("failed to hget all")
		}
	}

	incrbyRet, err := client.Hincrby(context.Background(), &pb.HincrbyRequest{
		Key:   &pbKey,
		Field: "rank",
		Value: 2,
	})
	if err != nil || incrbyRet.ErrMsg.ErrCode != pb.ErrCode_Sucess {
		t.Fatal("failed to incrby,err:", err)
	}

	hexistRet, err := client.Hexists(context.Background(), &pb.HExistsRequest{
		Key:   &pbKey,
		Field: "rank",
	})

	t.Log("hexist ret:", hexistRet.String())

	if err != nil || !hexistRet.Exist {
		t.Fatal("failed to hexist")
	}

	hexistRet, err = client.Hexists(context.Background(), &pb.HExistsRequest{
		Key:   &pbKey,
		Field: "rank1",
	})

	t.Log("hexist ret:", hexistRet.String())
	if err != nil || hexistRet.Exist {
		t.Fatal("failed to hexist")
	}

	hdelRet, err := client.Hdel(context.Background(), &pb.HdelRequest{
		Key:    &pbKey,
		Fields: []string{"zhang", "rank", "zhuhu"},
	})

	t.Log("hdel ret:", hdelRet.String())
	if err != nil || hdelRet.ErrCode != pb.ErrCode_Sucess {
		t.Fatal("failed to hdel,err:", err)
	}

}

//测试list类型
func TestList(t *testing.T) {
	client, err := NewGameRedisClient("172.13.0.53:1981")
	if err != nil {
		t.Fatal("failed to new game redis client,err:", err)
	}

	var (
		pbKey = pb.Key{
			GameId: 1002,
			Key:    "test",
		}
	)

	client.Del(context.Background(), &pb.DelRequest{
		Key:  &pbKey,
		Type: pb.DataType_List,
	})

	lpushRet, err := client.Lpush(context.Background(), &pb.LPushRequest{
		Key:   &pbKey,
		Value: []string{"zhao", "qian", "sun", "li"},
	})

	if err != nil || lpushRet.ErrCode != pb.ErrCode_Sucess {
		t.Log("failed to lpush,err:", err)
	}

	lrangeRet, err := client.Lrange(context.Background(), &pb.LrangeRequest{
		Key:   &pbKey,
		Start: 0,
		End:   -1,
	})
	t.Log("lrangeRet:", lrangeRet.String())
	if err != nil || lrangeRet.ErrMsg.ErrCode != pb.ErrCode_Sucess {
		t.Fatal("failed to lrange")
	}

	lpopRet, err := client.Lpop(context.Background(), &pbKey)
	t.Log("lpop ret,", lpopRet.String())
	if err != nil || lpopRet.ErrMsg.ErrCode != pb.ErrCode_Sucess || lpopRet.Value != "li" {
		t.Fatal("failed to lpop")
	}

	RpopRet, err := client.Rpop(context.Background(), &pbKey)
	t.Log("rpop ret,", RpopRet.String())
	if err != nil || RpopRet.ErrMsg.ErrCode != pb.ErrCode_Sucess || RpopRet.Value != "zhao" {
		t.Fatal("failed to Rpop")
	}

	_, err = client.Rpush(context.Background(), &pb.RPushRequest{
		Key:   &pbKey,
		Value: []string{"zhou", "wu", "zheng", "wang"},
	})

	if err != nil || lpushRet.ErrCode != pb.ErrCode_Sucess {
		t.Log("failed to lpush,err:", err)
	}

	lrangeRet, err = client.Lrange(context.Background(), &pb.LrangeRequest{
		Key:   &pbKey,
		Start: 0,
		End:   -1,
	})
	t.Log("lrangeRet:", lrangeRet.String())
	if err != nil || lrangeRet.ErrMsg.ErrCode != pb.ErrCode_Sucess {
		t.Fatal("failed to lrange")
	}

	llenRet, err := client.Llen(context.Background(), &pbKey)
	if err != nil || llenRet.Len != 6 {
		t.Fatal("failed to llen")
	}
}

//测试set类型
func TestSet(t *testing.T) {
	client, err := NewGameRedisClient("172.13.0.53:1981")
	if err != nil {
		t.Fatal("failed to new game redis client,err:", err)
	}

	var (
		pbKey = pb.Key{
			GameId: 1002,
			Key:    "test",
		}
	)

	_, err = client.Sadd(context.Background(), &pb.SaddRequest{
		Key:     &pbKey,
		Members: []string{"a", "b", "c", "d", "a"},
	})

	if err != nil {
		t.Fatal("failed to sadd err:", err)
	}

	smembersRet, err := client.Smembers(context.Background(), &pbKey)

	t.Log("smember ret:", smembersRet.String())
	if err != nil || smembersRet.ErrMsg.ErrCode != pb.ErrCode_Sucess {
		t.Fatal("failed to smembers,err:", err)
	}

	SremRet, err := client.Srem(context.Background(), &pb.SremRequest{
		Key:     &pbKey,
		Members: []string{"a", "b"},
	})

	t.Log("SremRet ret:", SremRet.String())
	if err != nil || SremRet.ErrCode != pb.ErrCode_Sucess {
		t.Fatal("failed to Srem,err:", err)
	}

	smembersRet, err = client.Smembers(context.Background(), &pbKey)

	t.Log("smember ret:", smembersRet.String())
	if err != nil || smembersRet.ErrMsg.ErrCode != pb.ErrCode_Sucess {
		t.Fatal("failed to smembers,err:", err)
	}

	sisMember, err := client.SisMember(context.Background(), &pb.SisMemberRequest{
		Key:    &pbKey,
		Member: "c",
	})
	t.Log("sisMember ret:", sisMember.String())
	if err != nil || !sisMember.Exist {
		t.Fatal("failed to sisMember,err:", err)
	}

	sisMember, err = client.SisMember(context.Background(), &pb.SisMemberRequest{
		Key:    &pbKey,
		Member: "f",
	})
	t.Log("sisMember ret:", sisMember.String())
	if err != nil || sisMember.Exist {
		t.Fatal("failed to sisMember,err:", err)
	}
}

//测试Zset类型
func TestZSet(t *testing.T) {
	client, err := NewGameRedisClient("172.13.0.53:1981")
	if err != nil {
		t.Fatal("failed to new game redis client,err:", err)
	}

	var (
		pbKey = pb.Key{
			GameId: 1002,
			Key:    "test",
		}
	)

	_, err = client.Zadd(context.Background(), &pb.ZaddRequest{
		Key: &pbKey,
		MemberScores: []*pb.MemberScore{
			&pb.MemberScore{
				Member: "zhangsan",
				Score:  10086.7,
			},
			&pb.MemberScore{
				Member: "lisi",
				Score:  10087.2,
			},
			&pb.MemberScore{
				Member: "wangwu",
				Score:  10088.2,
			},
		},
	},
	)

	if err != nil {
		t.Fatal("failed to zadd err:", err)
	}

	ZrankRet, err := client.Zrank(context.Background(), &pb.ZrankRequest{
		Key:    &pbKey,
		Member: "zhangsan",
	})

	t.Log("ZrankRet:", ZrankRet.String())
	if err != nil || ZrankRet.ErrMsg.ErrCode != pb.ErrCode_Sucess || ZrankRet.Rank != 0 {
		t.Fatal("failed to Zrank,err:", err)
	}

	ZrevRankRet, err := client.ZrevRank(context.Background(), &pb.ZRevRankRequest{
		Key:    &pbKey,
		Member: "zhangsan",
	})

	t.Log("ZrevRankRet:", ZrevRankRet.String())
	if err != nil || ZrevRankRet.ErrMsg.ErrCode != pb.ErrCode_Sucess || ZrevRankRet.Rank != 2 {
		t.Fatal("failed to ZrevRankRet,err:", err)
	}

	//
	zrangeRet, err := client.Zrange(context.Background(), &pb.ZRangeRequest{
		Key:   &pbKey,
		Start: 0,
		End:   5,
	})

	t.Log("zrangeRet:", zrangeRet.String())

	//
	zRevRangeRet, err := client.ZrevRange(context.Background(), &pb.ZRevRangeRequest{
		Key:   &pbKey,
		Start: 0,
		End:   5,
	})

	t.Log("zRevRangeRet:", zRevRangeRet.String())

	zscoreRet, err := client.Zscore(context.Background(), &pb.ZscoreRequest{
		Key:    &pbKey,
		Member: "zhangsan",
	})

	if err != nil || zscoreRet.Score != 10086.7 {
		t.Fatal("failed to zscore")
	}
	_, err = client.Zrem(context.Background(), &pb.ZremRequest{
		Key:     &pbKey,
		Members: []string{"zhangsan", "lisi", "wji"},
	})

	if err != nil {
		t.Fatal("failed to zscore")
	}

	zrangeRet, err = client.Zrange(context.Background(), &pb.ZRangeRequest{
		Key:   &pbKey,
		Start: 0,
		End:   5,
	})

	t.Log("zrangeRet:", zrangeRet.String())

}
