/*
   涉及gs相关的请求
*/

package hall

import (
	"fmt"

	gproto "git.huoys.com/vn/proto"
	"github.com/golang/protobuf/proto"
	"github.com/vnroyalclub/go-common/log"
)

//获取游戏的地址
func GetGameAddress(gameId int32, groupId int32, serverId string, gsUrl string) (addressResp gproto.AllocServerAddressResp, err error) {

	addressResp.Servers = make([]*gproto.ServerAddr, 0)

	b := gproto.GetServerAddress{
		GameId:   proto.Int32(gameId),
		GroupId:  proto.Int32(groupId),
		ServerId: proto.String(serverId),
	}

	body, err := proto.Marshal(&b)
	if err != nil {
		log.Error("failed to marshal get server address,err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.ServiceOps_GS_Req_GetServerAddress), body, gsUrl)
	if err != nil {
		log.Error("failed to get server address,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to get server address,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	err = proto.Unmarshal(res.Data, &addressResp)
	if err != nil {
		log.Error("failed to unmarshal addressResp,err:", err)
		return
	}

	return
}

//游戏注册，想gs注册游戏
func RegisterGame(groupId int32, serverId string, port int32, gsUrl string) (err error) {

	msg := &gproto.RegisterGame{
		GroupId:  proto.Int32(groupId),
		ServerId: proto.String(serverId),
		Port:     proto.Int32(port),
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		log.Error("register game marshal data error:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.ServiceOps_GS_Req_GameRegister), data, gsUrl)
	if err != nil {
		log.Error("failed to register game,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to register game,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}
	return
}

//流水（自上次记录以来）
func RecordCapitalFlow(playerID int64, gameID, groupId int32, serverID string,
	profit, loss int64, gsUrl string) (err error) {

	msg := &gproto.RecordCapitalFlow{
		GameId:   proto.Int32(gameID),
		GroupId:  proto.Int32(groupId),
		ServerId: proto.String(serverID),
		Profit:   proto.Int64(profit),
		Loss:     proto.Int64(loss),
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		log.Error("failed to marshal record captial flow data err:", err)
		return
	}

	res, err := httpPost(playerID, int32(gproto.ServiceOps_GS_Req_RecordCapitalFlow), data, gsUrl)
	if err != nil {
		log.Error("failed to record capital flow,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Info("failed to record capital flow,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return
}

//游戏桌子数据更新
func GameTableUpdate(msg *gproto.GameTableUpdate, gsUrl string) (err error) {

	data, err := proto.Marshal(msg)
	if err != nil {
		log.Error("failed to marshal game table data err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.ServiceOps_GS_Req_GameTableUpdate), data, gsUrl)
	if err != nil {
		log.Error("failed to update game table data,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Info("failed to update game table data,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return
}

func LoadShareData(key string, gsUrl string) (rsp gproto.LoadValuesResp, err error) {

	msg := &gproto.LoadValues{}
	msg.Keys = append(msg.Keys, key)
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Error("failed to marshal request data,err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.ServiceOps_GS_Req_LoadShareData), data, gsUrl)
	if err != nil {
		log.Error("failed to load share data,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Info("failed to load share data,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	if err = proto.Unmarshal(res.GetData(), &rsp); err != nil {
		log.Error("failed to unmarshal respond body,err:", err)
		return
	}

	return
}

func UpdateShareData(playerId int64, mapValues map[string]string, gsUrl string) (err error) {

	msg := &gproto.UpdateKeyValues{}
	for key, value := range mapValues {
		item := &gproto.KeyValue{
			Key:   proto.String(key),
			Value: proto.String(value),
		}
		msg.Items = append(msg.Items, item)
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		log.Error("failed to marshal request data,err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.ServiceOps_GS_Req_UpdateShareData), data, gsUrl)
	if err != nil {
		log.Error("failed to update share data,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Info("failed to update share data,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return
}

//
func LoadKeyValues(gsUrl string, key ...string) (kv map[string]string, err error) {

	kv = make(map[string]string)

	msg := &gproto.LoadValues{}
	msg.Keys = append(msg.Keys, key...)

	data, err := proto.Marshal(msg)
	if err != nil {
		log.Error("load key values marshal data err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.ServiceOps_GS_Req_LoadValues), data, gsUrl)
	if err != nil {
		log.Error("failed to load key values ,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Info("failed to load key values,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	values := &gproto.LoadValuesResp{}
	err = proto.Unmarshal(res.Data, values)
	if err != nil {
		log.Error(" load key values data unmarshal err:", err)
		return
	}

	for _, item := range values.Items {
		kv[item.GetKey()] = item.GetValue()
	}

	return
}

//更新键值对
func UpdateKeyValues(playerId int64, mapValues map[string]string, gsUrl string) (err error) {

	msg := &gproto.UpdateKeyValues{}
	for key, value := range mapValues {
		item := &gproto.KeyValue{
			Key:   proto.String(key),
			Value: proto.String(value),
		}
		msg.Items = append(msg.Items, item)
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		log.Error("failed to marshal update key value,err:", err)
		return
	}

	res, err := httpPost(playerId, int32(gproto.ServiceOps_GS_Req_UpdateKeyValues), data, gsUrl)
	if err != nil {
		log.Error("failed to update key value ,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Info("failed to update key value,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return
}

//游戏服务器的更新
func GameUpdate(msg gproto.GameUpdate, gsUrl string) (err error) {
	data, err := proto.Marshal(&msg)
	if err != nil {
		log.Error("game update Mashal data error:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.ServiceOps_GS_Req_GameUpdate), data, gsUrl)
	if err != nil {
		log.Error("failed to game update ,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Info("failed to game update,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	return
}
