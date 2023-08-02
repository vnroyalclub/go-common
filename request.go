package hall

import (
	"encoding/json"
	"fmt"
	"net/http"

	gproto "github.com/vnroyalclub/go_proto"
	"github.com/vnroyalclub/go-common/httputil"
	"github.com/vnroyalclub/go-common/log"

	"github.com/golang/protobuf/proto"
)

//与大厅交互
func httpPost(playerId int64, ops int32, body []byte, url string) (hr gproto.HttpResult, err error) {
	//组URL
	if playerId > 0 {
		url = fmt.Sprintf("%s?ops=%d&playerid=%d", url, ops, playerId)
	} else {
		url = fmt.Sprintf("%s?ops=%d", url, ops)
	}

	//发送请求
	code, res, err := httputil.Request(http.MethodPost, url, string(body), nil)
	if err != nil || code != http.StatusOK {
		log.Error("failed http post,ops:", ops, "url:", url, "body:", string(body), "err:", err)
		err = fmt.Errorf("ops:%d,err:%v", ops, err)
		return
	}

	//解析结果
	err = proto.Unmarshal(res, &hr)
	if err != nil {
		log.Error("failed to unmarshal http post result ,ops:", ops, "url:", url, "res:", string(res))
		return
	}

	return
}

//加载玩家数据(基本数据)
func LoadPlayer(playerId int64, hallUrl string) (playerData gproto.PlayerData, err error) {

	res, err := httpPost(playerId, int32(gproto.Operation_OP_S_ReqLoadPlayer), nil, hallUrl)

	if err != nil {
		log.Error("failed to load player data,err:", err)
		return
	}

	if *res.En != int32(gproto.ErrorCode_Success) {
		log.Error("failed to load player data,errNum", *res.En)
		err = fmt.Errorf("errNum:%v", *res.En)
		return
	}

	err = proto.Unmarshal(res.Data, &playerData)
	if err != nil {
		log.Error("failed to unmarshal player data,err:", err)
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

	if *res.En != int32(gproto.ErrorCode_Success) {
		log.Error("load player with props errNum:", *res.En)
		err = fmt.Errorf("errNum:%v", *res.En)
		return
	}

	err = proto.Unmarshal(res.Data, &lpwpr)
	if err != nil {
		log.Error("failed to load player with props,err:", err)
		return
	}

	return
}

// //身份验证
// func VerifyToken(playerId int64, token string, url string) (valid bool, err error) {

// 	tokenInfo := &gproto.TokenInfo{
// 		PlayerId: proto.Int64(playerId),
// 		Token:    proto.String(token),
// 	}
// 	body, err := proto.Marshal(tokenInfo)
// 	if err != nil {
// 		log.Error("failed to marshal token info,err:", err)
// 		return
// 	}

// 	res, err := httpPost(playerId, int32(gproto.Operation_OP_S_ReqVerifyToken), body, url)
// 	if err != nil {
// 		log.Error("failed to verify token,err:", err)
// 		return
// 	}

// 	if *res.En != int32(gproto.ErrorCode_Success) {
// 		log.Info("token invalid,retrun errNum:", *res.En)
// 		return
// 	}

// 	valid = true
// 	return
// }

//修改玩家道具信息(GameID 为大厅)
func ModProps(playerId int64, sourceType gproto.TransferSource, url string, configIdCount map[int32]int64) (err error) {

	if len(configIdCount) == 0 {
		return fmt.Errorf("configid count is empty")
	}

	props := &gproto.WebModProps{
		GameId:    proto.Int32(1),
		Source:    &sourceType,
		PlayerIds: []int64{playerId},
	}

	for k, v := range configIdCount {
		configId := k
		count := v
		props.Items = append(props.Items, &gproto.PropExchangeData{
			ConfigId: &configId,
			Count:    &count,
		})
	}

	body, err := proto.Marshal(props)
	if err != nil {
		log.Error("failed to marshal token info,err:", err)
		return
	}

	res, err := httpPost(playerId, int32(gproto.Operation_OP_S_ReqWebModProps), body, url)
	if err != nil {
		log.Error("failed to verify token,err:", err)
		return
	}

	if *res.En != int32(gproto.ErrorCode_Success) {
		log.Error("failed to mod props,retrun errNum:", *res.En)
		err = fmt.Errorf("errNum:%v", *res.En)
		return
	}

	return
}

//拉取玩家道具信息
func LoadProps(playerId int64, hallUrl string) (props gproto.LoadPropsResp, err error) {

	res, err := httpPost(playerId, int32(gproto.Operation_OP_S_ReqLoadProps), nil, hallUrl)

	if err != nil {
		log.Error("failed to load player props,err:", err)
		return
	}

	if *res.En != int32(gproto.ErrorCode_Success) {
		log.Error("failed to load player props,errNum", *res.En)
		err = fmt.Errorf("errNum:%v", *res.En)
		return
	}

	err = proto.Unmarshal(res.Data, &props)
	if err != nil {
		log.Error("failed to unmarshal player props,err:", err)
		return
	}

	return
}

//发送日志
func SendLogs(logMsgs gproto.LogMsgs, url string) (err error) {
	body, err := json.Marshal(logMsgs)
	if err != nil {
		log.Error("failed to marshal logmsgs,err:", err)
		return
	}
	header := map[string]string{}
	code, _, err := httputil.Request(http.MethodPost, url, string(body), header)
	if code != http.StatusOK || err != nil {
		log.Error("failed to send log msg,code:", code, "err:", err)
		err = fmt.Errorf("code:%d,err:%v", code, err)
		return
	}

	return
}

//冻结玩家账号
func FreezeAccount(fa gproto.FreezeAccount, url string) (err error) {
	body, err := proto.Marshal(&fa)
	if err != nil {
		log.Error("failed to marshal freeze account,err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqFreezeAccount), body, url)
	if err != nil {
		log.Error("failed to freeze account,err:", err)
		return
	}

	if *res.En != int32(gproto.ErrorCode_Success) {
		log.Error("failed to freeze account,retrun errNum:", *res.En)
		err = fmt.Errorf("errNum:%v", *res.En)
		return
	}
	return
}

//跑马灯
// func StartMarquee(marquee gproto.Marquee, url string) (err error) {

// 	body, err := proto.Marshal(&marquee)
// 	if err != nil {
// 		log.Error("failed to marshal marquee,err:", err)
// 		return
// 	}

// 	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqStartMarquee), body, url)
// 	if err != nil {
// 		log.Error("failed to start marquee,err:", err)
// 		return
// 	}

// 	if *res.En != int32(gproto.ErrorCode_Success) {
// 		log.Error("failed to start marquee,retrun errNum:", *res.En)
// 		err = fmt.Errorf("errNum:%v", *res.En)
// 		return
// 	}
// 	return
// }

func CloseAccout(closeAccount gproto.CloseAccount, url string) (err error) {

	body, err := proto.Marshal(&closeAccount)
	if err != nil {
		log.Error("failed to marshal close account,err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqCloseAccount), body, url)
	if err != nil {
		log.Error("failed to close account,err:", err)
		return
	}

	if *res.En != int32(gproto.ErrorCode_Success) {
		log.Error("failed to start close account,retrun errNum:", *res.En)
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
