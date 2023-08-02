/*
   与大厅筹码道具相关的请求
*/

package hall

import (
	"fmt"

	gproto "git.huoys.com/vn/proto"
	"github.com/golang/protobuf/proto"
	"github.com/vnroyalclub/go-common/log"
)

//修改玩家道具信息(GameID 为大厅,在大厅其他子服务中使用)
func ModProps(playerId int64, sourceType gproto.TransferSource, hallUrl string, configIdCount map[int32]int64) (err error) {

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
		log.Error("failed to marshal web mod props,err:", err)
		return
	}

	res, err := httpPost(playerId, int32(gproto.Operation_OP_S_ReqWebModProps), body, hallUrl)
	if err != nil {
		log.Error("failed to mod web prop,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to mod props,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
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

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to load player props,errNum", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	err = proto.Unmarshal(res.Data, &props)
	if err != nil {
		log.Error("failed to unmarshal player props,err:", err)
		return
	}

	return
}

//修改大厅道具(在游戏中使用)
func ModHallProps(playerId int64, gameID, groupID int32, sourceType gproto.TransferSource, propsList map[int32]int64,
	isNotifyClient, isRaiseEvent bool, hallUrl string) (propsResps gproto.ModPropsResp, err error) {
	msgModProps := &gproto.ModProps{
		Source:         &sourceType,
		GameId:         proto.Int32(gameID),
		GroupId:        proto.Int32(groupID),
		NotifyClient:   proto.Bool(isNotifyClient),
		RaiseEvent:     proto.Bool(isRaiseEvent),
		RaiseTaskEvent: proto.Bool(true),
	}

	for key, value := range propsList {
		oneItem := &gproto.PropExchangeData{
			ConfigId: proto.Int32(key),
			Count:    proto.Int64(value),
		}
		msgModProps.Items = append(msgModProps.Items, oneItem)
	}

	body, err := proto.Marshal(msgModProps)
	if err != nil {
		log.Error("failed to marshal mod props, err:", err)
		return
	}

	res, err := httpPost(playerId, int32(gproto.Operation_OP_S_ReqModProps), body, hallUrl)
	if err != nil {
		log.Error("failed to mod hall props,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to mod props,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	err = proto.Unmarshal(res.GetData(), &propsResps)
	if err != nil {
		log.Error("failed to unmarshal data,err:", err)
		return
	}
	return
}

//查询玩家锁定筹码
func GetLockedChip(playerId int64, url string) (lockedChipRsp gproto.GetLockedChipInfoRsp, err error) {

	b := &gproto.GetLockedChipInfo{
		PlayerId: proto.Int64(playerId),
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

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqGetLockedChip), body, url)
	if err != nil {
		log.Error("failed to get lock chip,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to get lock chip,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	err = proto.Unmarshal(res.Data, &lockedChipRsp)
	if err != nil {
		log.Error("failed to unmarshal locked chip data,err:", err)
		return
	}
	return
}

//请求修改筹码
//玩家在游戏外的时候使用,只有在本对应的游戏中或者大厅的时候，才可以使用修改成功
func ReqModProps(playerId int64, gameID, groupID int32, sourceType gproto.TransferSource, propsList map[int32]int64,
	hallUrl string) (sucess bool, err error) {
	msgModProps := &gproto.ModProps{
		Source:         &sourceType,
		GameId:         proto.Int32(gameID),
		GroupId:        proto.Int32(groupID),
		NotifyClient:   proto.Bool(true),
		RaiseEvent:     proto.Bool(true),
		RaiseTaskEvent: proto.Bool(true),
	}

	for key, value := range propsList {
		oneItem := &gproto.PropExchangeData{
			ConfigId: proto.Int32(key),
			Count:    proto.Int64(value),
		}
		msgModProps.Items = append(msgModProps.Items, oneItem)
	}
	body, err := proto.Marshal(msgModProps)
	if err != nil {
		log.Error("failed to marshal mod props, err:", err)
		return
	}

	res, err := httpPost(playerId, int32(gproto.Operation_OP_S_ReqModPropsEx), body, hallUrl)
	if err != nil {
		log.Error("failed to mod hall props,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to mod props,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	sucess = true
	return
}

//查询商品订单状态
func GetWareOrderStatus(playerId int64, WareIds []string, hallUrl string) (orderStatus gproto.GetOrderStatusResp, err error) {

	if len(WareIds) == 0 {
		err = fmt.Errorf("len(WareIds) == 0")
		return
	}

	msgWareReqInfo := &gproto.GetOrderStatus{
		WareIds: WareIds,
	}
	body, err := proto.Marshal(msgWareReqInfo)
	if err != nil {
		log.Error("failed to marshal wareReqInfo,err:", err)
		return
	}

	res, err := httpPost(playerId, int32(gproto.Operation_OP_S_ReqGetOrderStatus), body, hallUrl)
	if err != nil {
		log.Error("failed to get order status,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("get order status invalid,retrun errNum:", res.GetEn())
		err = fmt.Errorf("errNum:%v", res.GetEn())
		return
	}

	err = proto.Unmarshal(res.GetData(), &orderStatus)
	if err != nil {
		log.Error("failed to unmarshal data,err:", err)
		return
	}

	return
}
