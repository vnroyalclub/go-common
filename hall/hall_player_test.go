package hall

import (
	"testing"

	gproto "git.huoys.com/vn/proto"
	"github.com/golang/protobuf/proto"
)

func TestLoadPlayer(t *testing.T) {
	palyerId := int64(1032768)
	hallUrl := "http://172.13.0.53/hall/"

	ret, err := LoadPlayer(palyerId, hallUrl)
	if err != nil {
		t.Error("load player with props err:", err)
		return
	}

	t.Log("id:", ret.GetId(), "nick:", ret.GetNick(),
		"portrait:", ret.GetPortrait(), "vip:", ret.GetLevel())

}

func TestLoadPlayerWithProps(t *testing.T) {
	palyerId := int64(1032768)
	hallUrl := "http://172.13.0.53/hall/"
	propsId := map[string]int32{
		"VIP": int32(gproto.PropConfigID_EPC_VIPPoint),
	}

	ret, err := LoadPlayerWithProps(palyerId, hallUrl, propsId)
	if err != nil {
		t.Error("load player with props err:", err)
		return
	}

	t.Log("id:", ret.Data.Player.GetId(), "nick", ret.Data.Player.GetNick(),
		"portrait", ret.Data.Player.GetPortrait(), "vip", ret.Data.Player.GetLevel())

	for _, v := range ret.Data.Props {
		t.Log("configId:", v.GetConfigId(), "count:", v.GetCount())
	}

}

func TestFreezeAccount(t *testing.T) {
	url := "http://172.13.0.53/hall/"
	fa := gproto.FreezeAccount{
		PlayerId: []int64{1061456},
		Reason:   proto.String("测试冻结"),
		Remark:   proto.String("冻结备注"),
		Duration: proto.Int32(86400),
	}
	err := FreezeAccount(fa, url)
	if err != nil {
		t.Error("freeze account err:", err)
	}
}

func TestCLoseAccout(t *testing.T) {
	url := "http://172.13.0.53/hall/"
	fa := gproto.CloseAccount{
		PlayerId: []int64{10086},
		IP:       []string{},
		Imei:     []string{},
		Reason:   proto.String("reason"),
		Remark:   proto.String("remark"),
		Duration: proto.Int32(-1),
		OpBy:     proto.String("testor"),
	}
	err := CloseAccout(fa, url)
	if err != nil {
		t.Error("close account err:", err)
	}
}

func TestLoadPlayers(t *testing.T) {

	palyerId := []int64{1008232, 1000152}
	hallUrl := "http://172.13.0.53:9527/"

	ret, err := LoadPlayers(palyerId, hallUrl)
	if err != nil {
		t.Error("load player with props err:", err)
		return
	}

	for _, v := range ret.GetPlayers() {
		t.Log("id:", v.GetId(), "nick:", v.GetNick(),
			"portrait:", v.GetPortrait(), "vip:", v.GetLevel())
	}
}

func TestReqIfPlayerFrozen(t *testing.T) {

	palyerId := int64(1052412)
	hallUrl := "http://test.os.huoys.com/hall/"

	ret, err := ReqIfPlayerFrozen(palyerId, hallUrl)
	if err != nil {
		t.Error("req if player frozen err:", err)
		return
	}

	t.Log("req if player frozen,ret:", ret.String())
}

func TestLoadPlayerIpImei(t *testing.T) {

	palyerId := int64(1409766)
	hallUrl := "http://test.os.huoys.com/hall/"

	ret, err := LoadPlayerIpImei(palyerId, hallUrl)
	if err != nil {
		t.Error("Load Player IpImei err:", err)
		return
	}
	t.Log("Load Player IpImei,ret:", ret.String())
}
