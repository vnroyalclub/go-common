package redis

import "testing"

func TestHash_Hset_Hget(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	var (
		hashName = "test"
		field    = "name"
		value    = "zhangsan"
	)
	h := Hash{

		HashName: hashName,
	}

	err = h.Hset(coon, field, value)
	if err != nil {
		t.Error("hash hset err:", err)
		return
	}

	v, err := h.Hget(coon, field)
	if err != nil || v != value {
		t.Error("hash hset_hget err:", err)
	}
	h.Del(coon)

}

func TestHash_Hmset_Hmget_Hdel(t *testing.T) {

	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	var (
		hashName   = "test"
		filedValue = map[interface{}]interface{}{
			"name": "zhangsan",
			"age":  17,
			"sex":  "man",
		}
	)

	h := Hash{
		HashName: hashName,
	}
	err = h.Hmset(coon, filedValue)
	if err != nil {
		t.Error("hash hmset err:", err)
		return
	}

	fv, err := h.Hmget(coon, "name", "age", "sex")
	if err != nil && len(fv) != 3 && fv["name"] != "zhangsan" && fv["age"] != "17" && fv["sex"] != "man" {
		t.Error("hash hget err:", err)
		return
	}

	err1 := h.Hdel(coon, "name")
	fv, err = h.Hmget(coon, "name", "age", "sex")
	if err1 != nil || err != nil ||
		len(fv) != 3 || fv["name"] != "" || fv["age"] != "17" || fv["sex"] != "man" {
		t.Error("hash hdel err:", err)
		return
	}

	h.Del(coon)

}

func TestHash_HsetNx_Del(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()
	var (
		hashName = "test"
		field    = "name"
		value    = "zhangsan"
	)
	h := Hash{
		HashName: hashName,
	}

	success, err := h.HsetNx(coon, field, value)
	if err != nil || !success {
		t.Error("hash hsetnx err:", err)
		return
	}

	success, err = h.HsetNx(coon, field, value)
	if err != nil || success {
		t.Error("hash hsetnx err:", err)
		return
	}

	err = h.Del(coon)
	if err != nil {
		t.Error("hash del err:", err)
		return
	}

	h.Del(coon)
}

func TestHash_Hincrby(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	var (
		hashName = "test"
		field    = "age"
		value    = "17"
	)
	h := Hash{
		HashName: hashName,
	}
	err = h.Hset(coon, field, value)
	if err != nil {
		t.Error("hash hmset err:", err)
		return
	}

	err = h.Hincrby(coon, field, 13)
	v, err1 := h.Hget(coon, field)
	if err != nil || err1 != nil || v != "30" {
		t.Error("hincrby err:", err)
	}

	h.Del(coon)
}

func TestHash_HincrbyFloat(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	var (
		hashName = "test"
		field    = "age"
		value    = 12.8
	)
	h := Hash{
		HashName: hashName,
	}
	err = h.Hset(coon, field, value)
	if err != nil {
		t.Error("hash hmset err:", err)
		return
	}

	err = h.HincrbyFloat(coon, field, 12.8)
	v, err1 := h.Hget(coon, field)
	if err != nil || err1 != nil || v != "25.6" {
		t.Error("hincrby err:", err)
	}

	h.Del(coon)
}

func TestHash_HgetAll(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	var (
		hashName   = "test"
		filedValue = map[interface{}]interface{}{
			"name": "zhangsan",
			"age":  17,
			"sex":  "man",
		}
	)

	h := Hash{
		HashName: hashName,
	}
	err = h.Hmset(coon, filedValue)
	if err != nil {
		t.Error("hash hmset err:", err)
		return
	}

	fv, err := h.HgetAll(coon)
	if err != nil || len(fv) != 3 || fv["name"] != "zhangsan" || fv["age"] != "17" || fv["sex"] != "man" {
		t.Error("hash hget err:", err)
		return
	}

	h.Del(coon)
}

func TestHash_Hexists(t *testing.T) {

	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	var (
		hashName = "test"
		field    = "name"
		value    = "zhangsan"
	)
	h := Hash{

		HashName: hashName,
	}

	err = h.Hset(coon, field, value)
	if err != nil {
		t.Error("hash hset err:", err)
		return
	}

	exist, err := h.Hexists(coon, field)
	if err != nil || !exist {
		t.Error("hexist err:", err, "exist:", exist)
		return
	}

	exist, err = h.Hexists(coon, "aeqiwh")
	if err != nil || exist {
		t.Error("hexist err:", err, "exist:", exist)
		return
	}

	h.Del(coon)
}
