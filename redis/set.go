package redis

import (
	"github.com/gomodule/redigo/redis"
)

type Set struct {
	SetName string
}

//添加成员
func (s Set) Sadd(conn redis.Conn, members ...interface{}) (err error) {
	if len(members) == 0 {
		return
	}
	args := []interface{}{s.SetName}
	for _, arg := range members {
		args = append(args, arg)
	}

	_, err = conn.Do(RcSadd, args...)
	return
}

//删除集合
func (s Set) Del(conn redis.Conn) (err error) {
	_, err = conn.Do(RcDel, s.SetName)
	return
}

//获取到期时间
func (s Set) Ttl(conn redis.Conn) (ttl int64, err error) {
	ttl, err = redis.Int64(conn.Do(RcTtl, s.SetName))
	return
}

//设置过期时间
func (s Set) Expire(conn redis.Conn, second int32) (err error) {
	_, err = conn.Do(RcExpire, s.SetName, second)
	return
}

func (s Set) Srem(conn redis.Conn, members ...interface{}) (err error) {
	if len(members) == 0 {
		return
	}
	args := []interface{}{s.SetName}
	for _, arg := range members {
		args = append(args, arg)
	}
	_, err = conn.Do(RcSrem, args...)
	return
}

//获取集合所有成员
func (s Set) Smembers(conn redis.Conn) (values []string, err error) {
	values = make([]string, 0)
	values, err = redis.Strings(conn.Do(RcSmembers, s.SetName))
	return
}

//随机获取集合成员
func (s Set) SrandMembers(conn redis.Conn, count int) (values []string, err error) {
	if count == 0 {
		count = 1
	}

	values = make([]string, 0)
	values, err = redis.Strings(conn.Do(RcSRandMember, s.SetName, count))

	return
}

//随机获取集合成员
func (s Set) SisMember(conn redis.Conn, member interface{}) (is bool, err error) {
	v, err := redis.Int64(conn.Do(RcSisMember, s.SetName, member))
	if err != nil {
		return
	}

	//如果返回值为1，说明是成员，否则不是
	if v == 1 {
		is = true
	}
	return
}
