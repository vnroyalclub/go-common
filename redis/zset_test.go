package redis

import (
	"testing"
)

func TestZset_Zadd_Zsorce(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	z := Zset{
		ZsetName: "test",
	}
	err = z.Del(coon)
	if err != nil {
		t.Error("zset del err:", err)
		return
	}

	members := map[int64]interface{}{
		60: "zhangsan",
		70: "lisi",
	}
	err = z.Zadd(coon, members)
	if err != nil {
		t.Error("zset add err:", err)
		return
	}

	score, err := z.Zscore(coon, "zhangsan")
	if err != nil || score != 60 {
		t.Error("zset score err:", err)
		return
	}

	score, err = z.Zscore(coon, "lisi")
	if err != nil || score != 70 {
		t.Error("zset score err:", err)
		return
	}

	z.Del(coon)
}

func TestZset_ZIncrby(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	z := Zset{
		ZsetName: "test",
	}
	err = z.Del(coon)
	if err != nil {
		t.Error("zset del err:", err)
		return
	}

	members := map[int64]interface{}{
		60: "zhangsan",
	}
	err = z.Zadd(coon, members)
	if err != nil {
		t.Error("zset add err:", err)
		return
	}

	err = z.ZIncrby(coon, 18, "zhangsan")
	if err != nil {
		t.Error("zincrby err:", err)
		return
	}

	score, err := z.Zscore(coon, "zhangsan")
	if err != nil || score != 78 {
		t.Error("zset score err:", err)
		return
	}

	z.Del(coon)
}

func TestZset_Zrem(t *testing.T) {
	z := Zset{
		ZsetName: "test",
	}

	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	err = z.Del(coon)
	if err != nil {
		t.Error("zset del err:", err)
		return
	}

	members := map[int64]interface{}{
		60: "zhangsan",
		70: "lisi",
	}
	err = z.Zadd(coon, members)
	if err != nil {
		t.Error("zset add err:", err)
		return
	}

	err = z.Zrem(coon, "zhangsan")
	score, _ := z.Zscore(coon, "zhangsan")
	t.Log("score:", score)
	if err != nil || score != 0 {
		t.Error("zset zrem err:", err)
	}

	score, _ = z.Zscore(coon, "lisi")
	t.Log("score:", score)
	if err != nil || score != 70 {
		t.Error("zset zrem err:", err)
		return
	}

	z.Del(coon)
}

func TestZset_Zrank(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	z := Zset{
		ZsetName: "test",
	}
	err = z.Del(coon)
	if err != nil {
		t.Error("zset del err:", err)
		return
	}

	members := map[int64]interface{}{
		60: "zhangsan",
		70: "lisi",
	}
	err = z.Zadd(coon, members)
	if err != nil {
		t.Error("zset add err:", err)
		return
	}

	rank, err := z.Zrank(coon, "zhangsan")
	if err != nil || rank != 0 {
		t.Error("zrank err:", err)
		return
	}

	rank, err = z.Zrank(coon, "lisi")
	if err != nil || rank != 1 {
		t.Error("zrank err:", err)
		return
	}
}

func TestZset_ZrevRank(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	z := Zset{
		ZsetName: "test",
	}
	err = z.Del(coon)
	if err != nil {
		t.Error("zset del err:", err)
		return
	}

	members := map[int64]interface{}{
		60: "zhangsan",
		70: "lisi",
	}
	err = z.Zadd(coon, members)
	if err != nil {
		t.Error("zset add err:", err)
		return
	}

	rank, err := z.ZrevRank(coon, "zhangsan")
	if err != nil || rank != 1 {
		t.Error("zrank err:", err)
		return
	}

	rank, err = z.ZrevRank(coon, "lisi")
	if err != nil || rank != 0 {
		t.Error("zrank err:", err)
		return
	}
}

func TestZset_ZRangeWithScroe(t *testing.T) {

	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	z := Zset{
		ZsetName: "test",
	}
	err = z.Del(coon)
	if err != nil {
		t.Error("zset del err:", err)
		return
	}

	members := map[int64]interface{}{
		60: "zhangsan",
		70: "lisi",
		80: "wangwu",
	}
	err = z.Zadd(coon, members)
	if err != nil {
		t.Error("zset add err:", err)
		return
	}

	memberScore, err := z.ZRangeWithScroe(coon, 0, 0)
	if err != nil || len(memberScore) != 1 || memberScore[0].Score != 60 || memberScore[0].Member != "zhangsan" {
		t.Error("ZRangeWithScroe err:", err)
		return
	}

	z.Del(coon)
}

func TestZset_ZRevRangeWithScores(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()
	z := Zset{
		ZsetName: "test",
	}
	err = z.Del(coon)
	if err != nil {
		t.Error("zset del err:", err)
		return
	}

	members := map[int64]interface{}{
		60: "zhangsan",
		70: "lisi",
		80: "wangwu",
	}
	err = z.Zadd(coon, members)
	if err != nil {
		t.Error("zset add err:", err)
		return
	}

	memberScore, err := z.ZRevRangeWithScore(coon, 0, 0)
	if err != nil || len(memberScore) != 1 || memberScore[0].Score != 80 || memberScore[0].Member != "wangwu" {
		t.Error("ZRangeWithScroe err:", err)
		return
	}

	z.Del(coon)
}
