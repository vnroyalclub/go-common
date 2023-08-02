package hall

import (
	"fmt"

	gproto "git.huoys.com/vn/proto"
	"github.com/golang/protobuf/proto"
	"github.com/vnroyalclub/go-common/log"
)

//请求关服
func CloserServer(closeServer gproto.CloseServer, url string) (err error) {

	body, err := proto.Marshal(&closeServer)
	if err != nil {
		log.Error("failed to marshal close server,err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqCloseServer), body, url)
	if err != nil {
		log.Error("failed to close server,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to close server,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}
	return
}

//请求开服
func OpenServer(openServer gproto.OpenServer, url string) (err error) {
	body, err := proto.Marshal(&openServer)
	if err != nil {
		log.Error("failed to marshal open server,err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqOpenServer), body, url)
	if err != nil {
		log.Error("failed to open server,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to open server,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}
	return
}

//玩家进入游戏成功,需要通知下大厅
func EnterGameNotifyHall(playerId int64, gameId, groupId int32, serverId string, tableId int32,
	tokenRequired bool, hallUrl string) (err error) {

	msg := &gproto.PlayerGameInfo{
		GameId:        &gameId,
		GroupId:       &groupId,
		ServerId:      &serverId,
		TableId:       &tableId,
		TokenRequired: &tokenRequired,
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		log.Error("enter game notify Mashal data error:", err)
		return
	}

	res, err := httpPost(playerId, int32(gproto.Operation_OP_PushUpdatePlayerGameInfo), data, hallUrl)
	if err != nil {
		log.Error("failed to notify player enter game,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to notify player enter game,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return
}

//玩家离开游戏成功,需要通知下大厅
func LeaveGameNotifyHall(playerID int64, hallUrl string) (err error) {

	res, err := httpPost(playerID, int32(gproto.Operation_OP_PushRemovePlayerGameInfo), nil, hallUrl)
	if err != nil {
		log.Error("failed to notify player leave game,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to notify player leave game,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}
	return
}

//通知玩家进入指定场的游戏(常用于比赛开始)
func NotifyPlayerIntoGame(playerId int64, gameId int32, groupId int32, serverId, extraData string,
	hallUrl string) (err error) {

	msg := &gproto.NotifyPlayerIntoGame{
		GameId:   &gameId,
		GroupId:  &groupId,
		ServerId: &serverId,
		ExtraMsg: &extraData,
		PlayerId: &playerId,
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		log.Error("notify player into game marshal data error:", err)
		return
	}

	res, err := httpPost(playerId, int32(gproto.Operation_OP_S_ReqNotifyPlayerIntoGame), data, hallUrl)
	if err != nil {
		log.Error("failed to notify player into enter game,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to notify player enter game,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return
}

//侧边框通知
func ReqGameArena(gameId, groupId, matchId int32, arenaInfo string, extraMsg string, hallUrl string) (err error) {

	reqData := &gproto.GameArena{
		GameId:    proto.Int32(gameId),
		GroupId:   proto.Int32(groupId),
		MatchId:   proto.Int32(matchId),
		ArenaInfo: proto.String(arenaInfo),
		ExtraData: proto.String(extraMsg),
	}

	data, err := proto.Marshal(reqData)
	if err != nil {
		log.Error("failed to marshal game areba,err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqGameArena), data, hallUrl)
	if err != nil {
		log.Error("failed to request game arena,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("ailed to request game arena:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return
}

//侧边框通知扩展，通知同类游戏
func ReqGameArenaEx(gameId, groupId, matchId int32, arenaInfo string, extraMsg string, hallUrl string) (err error) {

	reqData := &gproto.GameArena{
		GameId:    proto.Int32(gameId),
		GroupId:   proto.Int32(groupId),
		MatchId:   proto.Int32(matchId),
		ArenaInfo: proto.String(arenaInfo),
		ExtraData: proto.String(extraMsg),
	}

	data, err := proto.Marshal(reqData)
	if err != nil {
		log.Error("failed to marshal game areba,err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqGameArenaEx), data, hallUrl)
	if err != nil {
		log.Error("failed to request game arena,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("ailed to request game arena:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return
}

//删除侧边框通知
func DelReqGameArena(gameId, groupId, matchId int32, hallUrl string) (err error) {
	reqData := &gproto.GameArena{
		GameId:  proto.Int32(gameId),
		GroupId: proto.Int32(groupId),
		MatchId: proto.Int32(matchId),
	}

	data, err := proto.Marshal(reqData)
	if err != nil {
		log.Error("ReqDelMatchMessage Mashal data error:", err)
		return err
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqDeleteGameArena), data, hallUrl)
	if err != nil {
		log.Error("failed to request game arena,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to request game arena:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return
}

func Chat(playerID, level int64, msgType gproto.ChatMsgType, portrait, nick, message string,
	sex int32, url string) (err error) {

	var resSex gproto.Sex
	resSex = gproto.Sex(sex)
	base := &gproto.PlayerBase{
		PlayerId: &playerID,
		Portrait: &portrait,
		Nick:     &nick,
		Level:    &level,
		Online:   proto.Bool(true),
		Sex:      &resSex,
	}
	msg := &gproto.Chat{
		Who:         base,
		MessageType: &msgType,
		Message:     &message,
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		log.Error("marshal chat data err:", err)
		return
	}

	res, err := httpPost(playerID, int32(gproto.Operation_OP_S_ReqGameChat), data, url)
	if err != nil {
		log.Error("request game chat err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to request game chat:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return
}

//FilterChatMessage 聊天请求大厅过滤敏感词
func FilterChatMessage(msg string, hallUrl string) (content string, err error) {

	data, err := proto.Marshal(&gproto.GameChatMsg{
		Content: proto.String(msg),
	})

	if err != nil {
		log.Error("FilterChatMessage Mashal data error:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqHandleGameChatMsg), data, hallUrl)
	if err != nil {
		log.Error("failed to req handle game char msg,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to req handle game char msg:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	rsp := &gproto.GameChatMsgResp{}
	if err = proto.Unmarshal(res.GetData(), rsp); err != nil {
		log.Error("failed to unmarshal respond body,err:", err)
		return
	}

	content = rsp.GetContent()
	return
}

//校验游戏入场限制
func IfCanEnterGame(playerId int64, gameId int32, groupId int32, serverId, hallUrl string) (pass bool, err error) {

	reqData := &gproto.IfCanEnterGame{
		PlayerId: proto.Int64(playerId),
		GameId:   proto.Int32(gameId),
		GroupId:  proto.Int32(groupId),
		ServerId: proto.String(serverId),
	}

	data, err := proto.Marshal(reqData)
	if err != nil {
		log.Error("failed to marshal can enter game,err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqIfCanEnterGame), data, hallUrl)
	if err != nil {
		log.Error("failed to req can enter game ,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to req handle game char msg:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	rsp := &gproto.IfCanEnterGameResp{}
	if err = proto.Unmarshal(res.GetData(), rsp); err != nil {
		log.Error("failed to unmarshal respond body,err:", err)
		return
	}

	pass = rsp.GetCanEnter()
	return
}

//请求大厅转发游戏消息
func TransferGameMessage(transferGameMessage gproto.TransferGameMessage, hallUrl string) (err error) {

	reqData := &transferGameMessage
	data, err := proto.Marshal(reqData)
	if err != nil {
		log.Error("transfer game message marshal data err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqTransferGameMessage), data, hallUrl)
	if err != nil {
		log.Error("failed to req can enter game ,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to req handle game char msg:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return
}

func UpdateGameArena(gameId, groupId, matchId int32, arenaInfo string, extraData string, hallUrl string) (err error) {

	reqData := &gproto.GameArena{
		GameId:    proto.Int32(gameId),
		GroupId:   proto.Int32(groupId),
		MatchId:   proto.Int32(matchId),
		ArenaInfo: proto.String(arenaInfo),
		ExtraData: proto.String(extraData),
	}

	data, err := proto.Marshal(reqData)
	if err != nil {
		log.Error("update game arena marshal data err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqUpdateGameArena), data, hallUrl)
	if err != nil {
		log.Error("failed to req update game arena ,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to req update game arena err:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return
}

// TaskNotifyReq 任务通知，任务系统通过这个协议号通知到大厅，大厅给玩家转发
func TaskNotifyReq(playerID int64, taskType int32, taskInfo, hallUrl string) (err error) {
	msg := &gproto.TaskNotify{
		PlayerId: proto.Int64(playerID),
		TaskType: proto.Int32(taskType),
		TaskInfo: proto.String(taskInfo),
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Error("failed to Marshal data error:", err)
		return
	}

	res, err := httpPost(playerID, int32(gproto.Operation_OP_S_ReqTaskNotify), data, hallUrl)
	if err != nil {
		log.Error("failed to load player with props,err:", err)
		return
	}
	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to request game chat:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}
	return
}

// PublishActivityMsgReq 请求大厅处理 推送活动信息
func PublishActivityMsgReq(activityID int32, playerID int64, msg string, hallUrl string) (err error) {
	base := &gproto.PublishActivityMsg{
		ActivityId: proto.Int32(activityID),
		PlayerId:   proto.Int64(playerID),
		Msg:        proto.String(msg),
	}

	data, err := proto.Marshal(base)
	if err != nil {
		log.Error("marshal data err:", err)
		return
	}

	res, err := httpPost(playerID, int32(gproto.Operation_OP_S_ReqPublishActivityMsg), data, hallUrl)
	if err != nil {
		log.Error("request Publish Activity Msg err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to request Publish Activity Msg:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return
}

// GetSubLoseWin 请求大厅处理 批量获取玩家输分情况 day = 0 获取当天的输分情况
func GetSubLoseWin(playerIDs []int64, day int64, hallUrl string) (datas []*gproto.GetLoseSubWinData, err error) {
	base := &gproto.GetLoseSubWinReq{
		PlayerId: playerIDs,
		Day:      proto.Int64(day),
	}

	data, err := proto.Marshal(base)
	if err != nil {
		log.Error("marshal data err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqGetLoseSubWin), data, hallUrl)
	if err != nil {
		log.Error("request lose win Msg err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to req handle lose win msg:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	resp := &gproto.GetLoseSubWinRsp{}
	if err = proto.Unmarshal(res.GetData(), resp); err != nil {
		log.Error("failed to unmarshal respond body,err:", err)
		return
	}
	datas = resp.GetDatas()
	return
}

// GetSuperTreasureTicket 请求大厅处理 获取玩家奖票数量
func GetSuperTreasureTicket(playerID int64, hallUrl string) (smallTickets, bigTickets int32, err error) {
	base := &gproto.GetSuperTreasureTicketReq{
		PlayerId: proto.Int64(playerID),
	}

	data, err := proto.Marshal(base)
	if err != nil {
		log.Error("marshal data err:", err)
		return
	}

	res, err := httpPost(playerID, int32(gproto.Operation_OP_S_ReqGetSuperTreasureTicket), data, hallUrl)
	if err != nil {
		log.Error("request Get Super Treasure Ticket err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to request Get Super Treasure Ticket:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	resp := &gproto.SuperTreasureTicketRsp{}
	if err = proto.Unmarshal(res.GetData(), resp); err != nil {
		log.Error("failed to unmarshal respond body,err:", err)
		return
	}

	smallTickets = resp.GetSmallTickets()
	bigTickets = resp.GetBigTickets()
	return
}

// SetSuperTreasureTicket 请求大厅处理 更新玩家的奖票数量
func SetSuperTreasureTicket(playerID int64, smallTicketsReq, bigTicketsReq int32, hallUrl string) (smallTickets, bigTickets int32, err error) {
	base := &gproto.SetSuperTreasureTicketReq{
		PlayerId:     proto.Int64(playerID),
		SmallTickets: proto.Int32(smallTicketsReq),
		BigTickets:   proto.Int32(bigTicketsReq),
	}

	data, err := proto.Marshal(base)
	if err != nil {
		log.Error("marshal data err:", err)
		return
	}

	res, err := httpPost(playerID, int32(gproto.Operation_OP_S_ReqSetSuperTreasureTicketReq), data, hallUrl)
	if err != nil {
		log.Error("request Set Super Treasure Ticket err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to request Set Super Treasure Ticket:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	resp := &gproto.SuperTreasureTicketRsp{}
	if err = proto.Unmarshal(res.GetData(), resp); err != nil {
		log.Error("failed to unmarshal respond body,err:", err)
		return
	}

	smallTickets = resp.GetSmallTickets()
	bigTickets = resp.GetBigTickets()

	return
}

//身份验证
func VerifyToken(playerId int64, token string, url string) (valid bool, err error) {

	tokenInfo := &gproto.TokenInfo{
		PlayerId: proto.Int64(playerId),
		Token:    proto.String(token),
	}
	body, err := proto.Marshal(tokenInfo)
	if err != nil {
		log.Error("failed to marshal token info,err:", err)
		return
	}

	res, err := httpPost(playerId, int32(gproto.Operation_OP_S_ReqVerifyToken), body, url)
	if err != nil {
		log.Error("failed to verify token,err:", err)
		return
	}

	if *res.En != int32(gproto.ErrorCode_Success) {
		log.Info("token invalid,retrun errNum:", *res.En)
		return
	}

	valid = true
	return
}

//跑马灯
func StartMarquee(marquee gproto.Marquee, url string) (err error) {

	body, err := proto.Marshal(&marquee)
	if err != nil {
		log.Error("failed to marshal marquee,err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqStartMarquee), body, url)
	if err != nil {
		log.Error("failed to start marquee,err:", err)
		return
	}

	if *res.En != int32(gproto.ErrorCode_Success) {
		log.Error("failed to start marquee,retrun errNum:", *res.En)
		err = fmt.Errorf("errNum:%v", *res.En)
		return
	}
	return
}

// func SendEmail(mailReq gproto.WebMailReq, url string) (err error) {

// 	body, err := proto.Marshal(&mailReq)
// 	if err != nil {
// 		log.Error("failed to marshal mail req,err:", err)
// 		return
// 	}

// 	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqMailAdd), body, url)
// 	if err != nil {
// 		log.Error("failed to close account,err:", err)
// 		return
// 	}

// 	if *res.En != int32(gproto.ErrorCode_Success) {
// 		log.Error("failed to send mail,retrun errNum:", *res.En)
// 		err = fmt.Errorf("errNum:%v", *res.En)
// 		return
// 	}
// 	return
// }
