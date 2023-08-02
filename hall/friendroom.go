package hall

import (
	"fmt"

	gproto "git.huoys.com/vn/proto"
	"github.com/golang/protobuf/proto"
	"github.com/vnroyalclub/go-common/log"
)

//请求存储房间
func ReqSaveRoom(hallUrl string, room gproto.SaveRoom) (err error) {
	reqData := &gproto.SaveRoom{
		GameId:     proto.Int32(room.GetGameId()),
		RoomNo:     proto.Int32(room.GetRoomNo()),
		Creator:    proto.Int64(room.GetCreator()),
		CreateTime: proto.Int64(room.GetCreateTime()),
		RoomId:     proto.Int64(room.GetRoomId()),
	}

	data, err := proto.Marshal(reqData)
	if err != nil {
		log.Error("ReqSaveRoom Mashal data error:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqSaveRoom), data, hallUrl)
	if err != nil {
		log.Error("failed to req save room,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Info("failed to req save room,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return
}

//请求删除房间
func ReqDelRoom(hallUrl string, room gproto.DelRoom) (err error) {

	reqData := &gproto.SaveRoom{
		GameId:     proto.Int32(room.GetGameId()),
		RoomNo:     proto.Int32(room.GetRoomNo()),
		Creator:    proto.Int64(room.GetCreator()),
		CreateTime: proto.Int64(room.GetCreateTime()),
	}

	data, err := proto.Marshal(reqData)
	if err != nil {
		log.Error("ReqDelRoom Mashal data error:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqDeleteRoom), data, hallUrl)
	if err != nil {
		log.Error("failed to req del room,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Info("failed to req del room,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return
}

//请求更新房间状态
func ReqUpdateRoomStatus(hallUrl string, roomStatus gproto.UpdateRoomStatus) (err error) {

	reqData := &gproto.UpdateRoomStatus{
		RoomNo:     proto.Int32(roomStatus.GetRoomNo()),
		Status:     proto.Int32(roomStatus.GetStatus()),
		CreateTime: proto.Int64(roomStatus.GetCreateTime()),
		GameId:     proto.Int32(roomStatus.GetGameId()),
	}

	data, err := proto.Marshal(reqData)
	if err != nil {
		log.Error("ReqUpdateRoomStatus Mashal data error:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqUpdateRoomStatus), data, hallUrl)
	if err != nil {
		log.Error("failed to req update room,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Info("failed to req update room,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return
}

//请求加入玩家
func ReqAddRoomPlayer(hallUrl string, player gproto.AddRoomPlayer) (err error) {

	reqData := &gproto.AddRoomPlayer{
		RoomNo:     proto.Int32(player.GetRoomNo()),
		CreateTime: proto.Int64(player.GetCreateTime()),
		Player:     player.Player,
		GameId:     player.GameId,
	}

	data, err := proto.Marshal(reqData)
	if err != nil {
		log.Error("ReqAddRoomPLayer Mashal data error:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqAddRoomPlayer), data, hallUrl)
	if err != nil {
		log.Error("failed to req add room player,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Info("failed to req add room player,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return err
}

//请求删除玩家
func ReqDelRoomPLayer(hallUrl string, player gproto.DelRoomPlayer) (err error) {

	reqData := &gproto.DelRoomPlayer{
		RoomNo:     proto.Int32(player.GetRoomNo()),
		CreateTime: proto.Int64(player.GetCreateTime()),
		Player:     player.Player,
		GameId:     player.GameId,
	}

	data, err := proto.Marshal(reqData)
	if err != nil {
		log.Error("ReqDelRoomPLayer Mashal data error:%v", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqDelRoomPlayer), data, hallUrl)
	if err != nil {
		log.Error("failed to req del room player,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Info("failed to req del room player,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return
}

//请求加入旁观玩家
func ReqAddRoomLooker(hallUrl string, player gproto.AddRoomLooker) (err error) {

	reqData := &gproto.AddRoomLooker{
		RoomNo:     proto.Int32(player.GetRoomNo()),
		CreateTime: proto.Int64(player.GetCreateTime()),
		Looker:     player.Looker,
		GameId:     player.GameId,
	}

	data, err := proto.Marshal(reqData)
	if err != nil {
		log.Error("ReqAddRoomLooker Mashal data error:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqAddRoomLooker), data, hallUrl)
	if err != nil {
		log.Error("failed to req add room looker ,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Info("failed to req add room looker,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return
}

//请求删除旁观玩家
func ReqDelRoomLooker(hallUrl string, player gproto.DelRoomLooker) (err error) {

	reqData := &gproto.DelRoomLooker{
		RoomNo:     proto.Int32(player.GetRoomNo()),
		CreateTime: proto.Int64(player.GetCreateTime()),
		Looker:     player.Looker,
		GameId:     player.GameId,
	}

	data, err := proto.Marshal(reqData)
	if err != nil {
		log.Error("ReqDelRoomLooker Mashal data error:%v", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqDelRoomLooker), data, hallUrl)
	if err != nil {
		log.Error("failed to req add room looker ,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Info("failed to req add room looker,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return
}

// 邀请进入房间
func ReqInviteIntoGame(hallUrl string, inviteInfo gproto.InvitePlayerIntoGame) (gproto.ErrorCode, error) {
	msg := &gproto.InvitePlayerIntoGame{
		GameId:   proto.Int32(inviteInfo.GetGameId()),
		GroupId:  proto.Int32(inviteInfo.GetGroupId()),
		ServerId: proto.String(inviteInfo.GetServerId()),
		ExtraMsg: proto.String(inviteInfo.GetExtraMsg()),
		RoomNo:   proto.Int64(inviteInfo.GetRoomNo()),
		Invitor:  inviteInfo.Invitor,
		Invitee:  inviteInfo.Invitee,
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		log.Error("ReqInviteIntoGame Mashal data error:", err.Error())
		return gproto.ErrorCode_SerializeFailed, err
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqInviteIntoGame), data, hallUrl)
	if err != nil {
		log.Error("failed to req add room looker ,err:", err)
		return gproto.ErrorCode_Failed, err
	}

	return gproto.ErrorCode(res.GetEn()), nil
}

//请求大厅保存游戏记录
func ReqSaveGameRecord(hallUrl string, gameRecord gproto.GameRecordData) (err error) {

	reqData := &gproto.GameRecordData{
		Key:        gameRecord.Key,
		Value:      gameRecord.Value,
		GameId:     gameRecord.GameId,
		Creator:    gameRecord.Creator,
		RoomNo:     gameRecord.RoomNo,
		CreateTime: gameRecord.CreateTime,
	}

	for _, player := range gameRecord.Players {
		reqData.Players = append(reqData.Players, player)
	}

	data, err := proto.Marshal(reqData)
	if err != nil {
		log.Error("ReqSaveGameRecord Mashal data error:%v", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqSaveGameRecord), data, hallUrl)
	if err != nil {
		log.Error("failed to req save game record,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Info("failed to req save game record,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}
	return
}
