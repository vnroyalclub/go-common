package hall

import (
	"fmt"

	"git.huoys.com/vn/go-common/log"
	gproto "git.huoys.com/vn/proto"
	"github.com/golang/protobuf/proto"
)

func LoadRobotsID(lvFrom, lvTo int32, hallUrl string) (robotResp gproto.LoadRobotsResp, err error) {

	msg := &gproto.LoadRobots{
		LevelFrom: proto.Int32(lvFrom),
		LevelTo:   proto.Int32(lvTo),
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		log.Error("failed to marshal load robots body,err:", err)
		return
	}

	res, err := httpPost(0, int32(gproto.Operation_OP_S_ReqLoadRobots), data, hallUrl)

	if err != nil {
		log.Error("failed to request robot hall props,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to req handle game char msg:", *res.En)
		err = fmt.Errorf("errNum:%v", *res.En)
		return
	}

	err = proto.Unmarshal(res.GetData(), &robotResp)
	if err != nil {
		log.Error("LoadRobotsID  %v", err.Error())
		return
	}

	return
}

//修改机器人道具(在游戏中使用)
func ModRobotHallProps(gameId, groupId int32, playerId int64,
	source *gproto.TransferSource, propsList map[int32]int64, hallUrl string) (err error) {

	if len(propsList) == 0 {
		return
	}

	msgModProps := &gproto.ModProps{
		Source:         source,
		GameId:         proto.Int32(gameId),
		GroupId:        proto.Int32(groupId),
		NotifyClient:   proto.Bool(false),
		RaiseEvent:     proto.Bool(false),
		RaiseTaskEvent: proto.Bool(false),
	}

	for key, value := range propsList {
		oneItem := &gproto.PropExchangeData{
			ConfigId: proto.Int32(key),
			Count:    proto.Int64(value),
		}
		msgModProps.Items = append(msgModProps.Items, oneItem)
	}

	data, err := proto.Marshal(msgModProps)
	if err != nil {
		log.Error("failed to marshal data,err:", err)
		return
	}

	res, err := httpPost(playerId, int32(gproto.Operation_OP_S_ReqRobotsModProps), data, hallUrl)

	if err != nil {
		log.Error("failed to request robot hall props,err:", err)
		return
	}

	if res.GetEn() != int32(gproto.ErrorCode_Success) {
		log.Error("failed to req handle game char msg:", *res.En)
		err = fmt.Errorf("errNum:%v", *res.En)
		return
	}

	return
}
