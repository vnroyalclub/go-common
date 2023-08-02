package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

type Zset struct {
	ZsetName string
}

type MemberScore struct {
	Member string
	Score  int64
}

func (z Zset) ZRangeWithScroe(conn redis.Conn, start, stop int64) (memberScores []MemberScore, err error) {
	memberScores = make([]MemberScore, 0)
	values, err := redis.Strings(conn.Do(RcZrange, z.ZsetName, start, stop, "WITHSCORES"))
	if err != nil {
		return
	}

	if len(values)%2 != 0 {
		err = fmt.Errorf("expects even number of values result")
		return
	}

	for i := 0; i < len(values); i += 2 {
		score, err1 := strconv.ParseInt(values[i+1], 10, 64)
		if err1 != nil {
			err = err1
			return
		}
		memberScores = append(memberScores, MemberScore{
			Member: values[i],
			Score:  score,
		})
	}
	return
}

func (z Zset) ZIncrby(conn redis.Conn, increment int64, member interface{}) (err error) {
	_, err = conn.Do(RcZincrby, z.ZsetName, increment, member)
	return
}

func (z Zset) Zadd(conn redis.Conn, scoreMember map[int64]interface{}) (err error) {
	if len(scoreMember) == 0 {
		return
	}

	args := []interface{}{z.ZsetName}
	for score, member := range scoreMember {
		args = append(args, score, member)
	}

	_, err = conn.Do(RcZadd, args...)
	return
}

func (z Zset) Zscore(conn redis.Conn, member interface{}) (score int64, err error) {
	score, err = redis.Int64(conn.Do(RcZscore, z.ZsetName, member))
	return
}

func (z Zset) ZRevRangeWithScore(conn redis.Conn, start, stop int64) (memberScores []MemberScore, err error) {

	memberScores = make([]MemberScore, 0)
	values, err := redis.Strings(conn.Do(RcZrevRange, z.ZsetName, start, stop, "WITHSCORES"))
	if err != nil {
		return
	}

	if len(values)%2 != 0 {
		err = fmt.Errorf("expects even number of values result")
		return
	}

	for i := 0; i < len(values); i += 2 {
		score, err1 := strconv.ParseInt(values[i+1], 10, 64)
		if err1 != nil {
			err = err1
			return
		}
		memberScores = append(memberScores, MemberScore{
			Member: values[i],
			Score:  score,
		})
	}

	return
}

func (z Zset) Zrem(conn redis.Conn, members ...interface{}) (err error) {
	if len(members) == 0 {
		return
	}

	args := []interface{}{z.ZsetName}
	for _, member := range members {
		args = append(args, member)
	}
	_, err = conn.Do(RcZrem, args...)
	return
}

//倒叙排行rank,分数从低到高
func (z Zset) Zrank(conn redis.Conn, member interface{}) (rank int64, err error) {
	rank, err = redis.Int64(conn.Do(RcZrank, z.ZsetName, member))
	return
}

//获取排行，分数从高到低排序
func (z Zset) ZrevRank(conn redis.Conn, member interface{}) (rank int64, err error) {
	rank, err = redis.Int64(conn.Do(RcZrevRank, z.ZsetName, member))
	return
}

//删除集合
func (z Zset) Del(conn redis.Conn) (err error) {
	_, err = conn.Do(RcDel, z.ZsetName)
	return
}

//设置过期时间
func (z Zset) Expire(conn redis.Conn, second int32) (err error) {
	_, err = conn.Do(RcExpire, z.ZsetName, second)
	return
}
