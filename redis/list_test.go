package redis

import "testing"

func TestList_Lpush_lpop(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	var (
		l = List{
			ListName: "test",
		}
		values = []string{"zhao", "qian", "sun", "li"}
	)

	l.Del(coon)

	err = l.Lpush(coon, values...)
	if err != nil {
		t.Error("list lpush err:", err)
		return
	}

	value, err := l.Lpop(coon)
	if err != nil || value != "li" {
		t.Error("list lpop err:", err, "value:", value)
		return
	}

	value, err = l.Lpop(coon)
	if err != nil || value != "sun" {
		t.Error("list lpop err:", err, "value:", value)
		return
	}

	err = l.Del(coon)
	if err != nil {
		t.Error("list del,err:", err)
		return
	}
	return
}

func TestList_Rpush_Rpop(t *testing.T) {

	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	var (
		l = List{
			ListName: "test",
		}
		values = []string{"zhao", "qian", "sun", "li"}
	)

	l.Del(coon)

	err = l.Rpush(coon, values...)
	if err != nil {
		t.Error("list lpush err:", err)
		return
	}

	value, err := l.Rpop(coon)
	if err != nil || value != "li" {
		t.Error("list lpop err:", err, "value:", value)
		return
	}

	value, err = l.Rpop(coon)
	if err != nil || value != "sun" {
		t.Error("list lpop err:", err, "value:", value)
		return
	}

	err = l.Del(coon)
	if err != nil {
		t.Error("list del,err:", err)
		return
	}
	return
}

func TestList_Lrem_Lrange(t *testing.T) {

	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	l := List{
		ListName: "test",
	}
	l.Del(coon)
	err = l.Rpush(coon, "zhao", "qian", "li", "qian", "sun")
	if err != nil {
		t.Error("rpush err:", err)
		return
	}

	err = l.Lrem(coon, 1, "qian")
	if err != nil {
		t.Error("lrem err:", err)
	}

	values, err := l.Lrange(coon, 0, -1)
	t.Log("values:", values)
	if err != nil || len(values) != 4 || values[1] != "li" {
		t.Error("list lrange err:", err)
		return
	}

	l.Del(coon)

}

func TestList_Ltrim(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	l := List{
		ListName: "test",
	}

	l.Del(coon)

	err = l.Rpush(coon, "zhao", "qian", "li", "qian", "sun")
	if err != nil {
		t.Error("rpush err:", err)
		return
	}

	err = l.Ltrim(coon, 1, 2)
	if err != nil {
		t.Error("lrem err:", err)
	}

	values, err := l.Lrange(coon, 0, -1)
	t.Log("values:", values)
	if err != nil || len(values) != 2 || values[0] != "qian" || values[1] != "li" {
		t.Error("list lrange err:", err)
		return
	}

	l.Del(coon)

}

func TestList_Llen(t *testing.T) {
	p, err := newPool()
	if err != nil {
		t.Error("new pool err:", err)
		return
	}
	coon := p.Get()
	defer coon.Close()

	var (
		l = List{
			ListName: "test",
		}
		values = []string{"zhao", "qian", "sun", "li"}
	)

	len, err := l.Llen(coon)
	if err != nil || len != 0 {
		t.Error("list Llen,err:", err)
		return
	}

	err = l.Lpush(coon, values...)
	if err != nil {
		t.Error("list lpush err:", err)
		return
	}

	len, err = l.Llen(coon)
	if err != nil || len != 4 {
		t.Error("list Llen,err:", err)
		return
	}
	l.Del(coon)
}
