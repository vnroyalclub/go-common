package hall

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"testing"

	gproto "git.huoys.com/vn/proto"
	"github.com/golang/protobuf/proto"
)

func TestSendEmail(t *testing.T) {
	url := "http://172.13.0.53:9527/"
	source := gproto.TransferSource_TS_Msign_Sign
	fa := gproto.WebMailReq{
		PlayerIds:     []int64{1049152},
		From:          proto.String("system"),
		Title:         proto.String("恭喜你中奖"),
		Content:       proto.String("你中大奖了"),
		EffectiveDays: proto.Int32(30),
		Attachments: []*gproto.PropExchangeData{
			&gproto.PropExchangeData{
				ConfigId: proto.Int32(int32(gproto.PropConfigID_EPC_SlipGold)),
				Count:    proto.Int64(1),
			},
		},
		AttachmentsSource: &source,
	}
	err := SendEmail(fa, url)
	if err != nil {
		t.Error("close account err:", err)
	}
}

func TestStartMarquee(t *testing.T) {
	url := "http://172.13.0.53/hall/"
	fa := gproto.Marquee{
		Guid:    proto.String("fewuqwrei"),
		Start:   proto.String(time.Now().Format("01/02/2006 15:04:05")),
		End:     proto.String(time.Now().Add(60 * 60).Format("01/02/2006 15:04:05")),
		Msg:     proto.String("跑马灯测试"),
		Type:    proto.Int32(0),
		SubType: proto.Int32(int32(gproto.ChatMsgType_CMT_Activity_Normal)),
		IsLoop:  proto.Bool(false),
	}
	err := StartMarquee(fa, url)
	if err != nil {
		t.Error("freeze account err:", err)
	}
}

func TestVerifyToken(t *testing.T) {
	palyerId := int64(1016160)
	hallUrl := "http://172.13.0.53/hall/"
	token := "c71f0edc405c40a599500c4e8c0f14a2"

	valid, err := VerifyToken(palyerId, token, hallUrl)
	if err != nil {
		t.Error("verify err:", err)
		return
	}

	t.Log("valid:", valid)
}

func TestEnterGameNotifyHall(t *testing.T) {
	palyerId := int64(1016160)
	gameId := int32(2004)
	groupId := int32(11)
	tableId := int32(0)
	serverId := "1103"
	tokenRequire := true
	hallUrl := "http://172.13.0.53:9527/"

	err := EnterGameNotifyHall(palyerId, gameId, groupId, serverId, tableId, tokenRequire, hallUrl)
	if err != nil {
		t.Error("enter game:", err)
		return
	}

}

func TestLeaveGameNotifyHall(t *testing.T) {
	palyerId := int64(1016160)
	hallUrl := "http://172.13.0.53:9527/"

	err := LeaveGameNotifyHall(palyerId, hallUrl)
	if err != nil {
		t.Error("enter game:", err)
		return
	}
}

func TestNotifyPlayerIntoGame(t *testing.T) {
	palyerId := int64(1055720)
	gameId := int32(2004)
	groupId := int32(10)
	serverId := "1103"
	extraData := `{"ServerType":1,"tip":"老马，进入比赛了","matchId":1003}`
	hallUrl := "http://172.13.0.53:9527/"

	err := NotifyPlayerIntoGame(palyerId, gameId, groupId, serverId, extraData, hallUrl)
	if err != nil {
		t.Error("NotifyPlayerIntoGame:", err)
		return
	}
}

func TestReqGameArena(t *testing.T) {
	gameId := int32(2004)
	groupId := int32(10)
	matchId := int32(1001)
	//arenaInfo := "1103"
	extraData := ""
	hallUrl := "http://172.13.0.53:9527/"

	a := map[string]interface{}{
		"Name":       "德州比赛开始了",
		"GameID":     gameId,
		"BeginHour":  9,
		"BeginMin":   0,
		"EndHour":    22,
		"EndMin":     0,
		"ShowTime":   60,      // 展示时间
		"RemainTime": 0,       // 倒计时
		"Money":      1000000, // 筹码奖励
		//"Material":   map[int32]int64{},
		//"Prop":       map[int32]int64{},
	}

	body, err := json.Marshal(a)
	if err != nil {
		t.Error("failed to marshal,err:", err)
		return
	}

	err = ReqGameArena(gameId, groupId, matchId, string(body), extraData, hallUrl)
	if err != nil {
		t.Error("ReqGameArena:", err)
		return
	}
}

func TestDelReqGameArena(t *testing.T) {
	gameId := int32(2004)
	groupId := int32(10)
	matchId := int32(1001)
	//arenaInfo := "1103"
	hallUrl := "http://172.13.0.53:9527/"

	err := DelReqGameArena(gameId, groupId, matchId, hallUrl)
	if err != nil {
		t.Error("DelReqGameArena:", err)
		return
	}
}

func TestGetSubLoseWin(t *testing.T) {
	playerIDs := []int64{1027076}
	day := int64(0)
	hallUrl := "http://172.13.0.53:9527/"

	datas, err := GetSubLoseWin(playerIDs, day, hallUrl)
	if err != nil {
		t.Error("TestGetSubLoseWin:", err)
		return
	}
	fmt.Printf("datas:%+v", datas)
}

func TestGetSuperTreasureTicket(t *testing.T) {
	playerID := int64(1027076)

	hallUrl := "http://172.13.0.53:9527/"

	smallTickets, bigTickets, err := GetSuperTreasureTicket(playerID, hallUrl)
	if err != nil {
		t.Error("TestGetSuperTreasureTicket:", err)
		return
	}
	fmt.Printf("smallTickets:%v,bigTickets:%v", smallTickets, bigTickets)
}

func TestSetSuperTreasureTicket(t *testing.T) {
	playerID := int64(1027076)
	smallTicketsReq := int32(2)
	bigTicketsReq := int32(0)
	hallUrl := "http://172.13.0.53:9527/"

	smallTickets, bigTickets, err := SetSuperTreasureTicket(playerID, smallTicketsReq, bigTicketsReq, hallUrl)
	if err != nil {
		t.Error("TestGetSuperTreasureTicket:", err)
		return
	}
	fmt.Printf("smallTickets:%v,bigTickets:%v", smallTickets, bigTickets)
}

func TestIfCanEnterGame(t *testing.T) {
	playerID := int64(1027076)
	gameID := int32(3005)
	groupID := int32(1)
	serverID := "123456789"
	hallURL := "http://172.13.0.119/hall/"
	pass, err := IfCanEnterGame(playerID, gameID, groupID, serverID, hallURL)
	if err != nil {
		log.Println("error:", err)
	}

	log.Println("pass:", pass)

}

func TestGetGameTaxConfig(t *testing.T) {
	gameID := int32(3005)
	hallURL := "http://172.13.0.119/hall/"
	pass, err := GetGameTaxConfig(gproto.GetGameTaxConfig{GameId: proto.Int32(gameID)}, hallURL)
	if err != nil {
		log.Println("error:", err)
	}

	log.Println("pass:", pass)

}

func TestShipping(t *testing.T) {
	orderID := "1234567"
	hallURL := "http://localhost:9527"
	pass, err := ReqShipping(orderID, hallURL)
	if err != nil {
		log.Println("error:", err)
	}

	log.Println("pass:", pass)
}
