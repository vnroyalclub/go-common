package hall

import (
	"fmt"

	gproto "github.com/vnroyalclub/go_proto"
	"github.com/golang/protobuf/proto"
	"github.com/vnroyalclub/go-common/log"
)

//加载玩家数据(基本数据)
func LoadPlayer(playerId int64, hallUrl string) (playerData gproto.PlayerData, err error) {

	res, err := httpPost(playerId, int32(gproto.Operation_OP_S_ReqLoadPlayer), nil, hallUrl)

	if err != nil {
		log.Error("failed to load player data,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to load player data,errNum", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	err = proto.Unmarshal(res.Data, &playerData)
	if err != nil {
		log.Error("failed to unmarshal player data,err:", err)
		return
	}

	return
}

//加载玩家数据(基本数据)
func LoadPlayers(playerIds []int64, hallUrl string) (playerData gproto.LoadPlayersResp, err error) {

	if len(playerIds) == 0 {
		err = fmt.Errorf("players len==0")
		log.Error("players len==0")
		return
	}

	b := gproto.LoadPlayers{
		PlayerIds: playerIds,
	}

	body, err := proto.Marshal(&b)
	if err != nil {
		log.Error("failed to marshal load players,err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqLoadPlayers), body, hallUrl)

	if err != nil {
		log.Error("failed to load players data,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to load players data,errNum", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	err = proto.Unmarshal(res.Data, &playerData)
	if err != nil {
		log.Error("failed to unmarshal players data,err:", err)
		return
	}

	return
}

//加载玩家数据和道具
func LoadPlayerWithProps(playerID int64, hallUrl string, propIDs map[string]int32) (
	lpwpr gproto.LoadPlayerWithPropsResp, err error) {

	msg := &gproto.LoadPlayerWithProps{
		PlayerId: proto.Int64(playerID),
	}

	for _, value := range propIDs {
		msg.PropIds = append(msg.PropIds, value)
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Error("failed to Marshal data error:", err)
		return
	}

	res, err := httpPost(playerID, int32(gproto.Operation_OP_S_ReqLoadPlayerWithProps), data, hallUrl)
	if err != nil {
		log.Error("failed to load player with props,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("load player with props errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	err = proto.Unmarshal(res.Data, &lpwpr)
	if err != nil {
		log.Error("failed to load player with props,err:", err)
		return
	}

	return
}

//拉取玩家ip,imei
func LoadPlayerIpImei(playerID int64, hallUrl string) (ret gproto.LoadPlayerIpImeiResp, err error) {
	msg := &gproto.LoadPlayerIpImei{
		PlayerId: proto.Int64(playerID),
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Error("failed to marshal load player Ipimei data error:", err)
		return
	}

	res, err := httpPost(playerID, int32(gproto.Operation_OP_S_ReqLoadPlayerIpImei), data, hallUrl)
	if err != nil {
		log.Error("failed to load player Ipimei data,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("load player Ipimei data errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	err = proto.Unmarshal(res.GetData(), &ret)
	if err != nil {
		log.Error("failed to load player Ipimei data,err:", err)
		return
	}

	return
}

//封号处理
func CloseAccout(closeAccount gproto.CloseAccount, hallUrl string) (err error) {

	body, err := proto.Marshal(&closeAccount)
	if err != nil {
		log.Error("failed to marshal close account,err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqCloseAccount), body, hallUrl)
	if err != nil {
		log.Error("failed to close account,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to start close account,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}
	return
}

//冻结玩家账号
func FreezeAccount(fa gproto.FreezeAccount, hallUrl string) (err error) {
	body, err := proto.Marshal(&fa)
	if err != nil {
		log.Error("failed to marshal freeze account,err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqFreezeAccount), body, hallUrl)
	if err != nil {
		log.Error("failed to freeze account,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to freeze account,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}
	return
}

//查询玩家在游戏中的信息
func LoadPlayerGameInfo(playerId int64, hallUrl string) (playerGameInfo gproto.PlayerGameInfo, err error) {

	b := &gproto.LoadPlayerGameInfo{
		PlayerId: proto.Int64(playerId),
	}

	body, err := proto.Marshal(b)
	if err != nil {
		log.Error("failed to marshal data,err:", err)
		return
	}
	res, err := httpPost(playerId, int32(gproto.Operation_OP_S_ReqLoadPlayerGameInfo), body, hallUrl)
	if err != nil {
		log.Error("failed to load player game info,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to load player data,errNum", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	err = proto.Unmarshal(res.Data, &playerGameInfo)
	if err != nil {
		log.Error("failed to unmarshal player game info data,err:", err)
		return
	}

	return
}

//查询冻结玩家账号
func LoadFreezeAccount(qeueryType int32, queryString string, halUrl string) (freezeAccInfo gproto.FreezeAccInfo, err error) {

	freezeAccInfo.Items = make([]*gproto.FreezeAccDetial, 0)

	b := &gproto.GetFreezeAcc{
		QueryType:   proto.Int32(qeueryType),
		QueryString: proto.String(queryString),
	}

	body, err := proto.Marshal(b)
	if err != nil {
		log.Error("failed to marshal data,err:", err)
		return
	}
	if err != nil {
		log.Error("failed to marshal get freeze,err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqLoadFreezeAccount), body, halUrl)
	if err != nil {
		log.Error("failed to get freeze account,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to freeze account,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	err = proto.Unmarshal(res.Data, &freezeAccInfo)
	if err != nil {
		log.Error("failed to unmarshal freeze acc info data,err:", err)
		return
	}
	return
}

// 判断玩家是否被东诶
func ReqIfPlayerFrozen(playerId int64, hallUrl string) (ifFrozenResp gproto.IfFrozenResp, err error) {

	b := &gproto.IfFrozen{
		PlayerId: proto.Int64(playerId),
	}

	body, err := proto.Marshal(b)
	if err != nil {
		log.Error("failed to marshal data,err:", err)
		return
	}

	res, err := httpPost(playerId, int32(gproto.Operation_OP_S_ReqIfPlayerFrozen), body, hallUrl)
	if err != nil {
		log.Error("failed to req if player frozen,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to load player data,errNum", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	err = proto.Unmarshal(res.Data, &ifFrozenResp)
	if err != nil {
		log.Error("failed to unmarshal ifFrozenResp,err:", err)
		return
	}

	return
}
