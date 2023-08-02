package hall

import (
	"testing"

	gproto "github.com/vnroyalclub/go_proto"
)

func TestModProps(t *testing.T) {
	palyerId := int64(1020640)
	hallUrl := "http://172.13.0.53/hall/"
	props := map[int32]int64{
		int32(gproto.PropConfigID_EPC_ChouMa):   1000000,
		int32(gproto.PropConfigID_EPC_ChengTuo): 2000,
	}

	err := ModProps(palyerId, gproto.TransferSource_TS_Agent, hallUrl, props)
	if err != nil {
		t.Error("mod props err:", err)
	}
}

func TestLoadProps(t *testing.T) {
	palyerId := int64(1027776)
	hallUrl := "http://172.13.0.53/hall/"

	props, err := LoadProps(palyerId, hallUrl)
	if err != nil {
		t.Error("load props err:", err)
		return
	}
	for k, _ := range props.Items {
		if props.Items[k].GetConfigId() == 2 {
			t.Log("chouma:", props.Items[k].GetCount())
		}
	}
}

func TestResComsumeProp(t *testing.T) {
	palyerId := int64(1471276)
	gameId := int32(2004)
	groupId := int32(10)
	source := gproto.TransferSource_TS_Game
	hallUrl := "http://172.13.0.53:9527/"
	props := map[int32]int64{
		int32(gproto.PropConfigID_EPC_ChouMa): -100000000,
	}

	sucess, err := ReqModProps(palyerId, gameId, groupId, source, props, hallUrl)
	if err != nil {
		t.Error("mod props err:", err)
	}

	t.Log("sucess:", sucess)
}
