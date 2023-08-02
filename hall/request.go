package hall

import (
	"fmt"
	"net/http"

	"git.huoys.com/vn/go-common/httputil"
	"git.huoys.com/vn/go-common/log"
	gproto "git.huoys.com/vn/proto"

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
		log.Error("failed http post,ops:", ops, "url:", url, "body:", string(body), "err:", err, "code:", code)
		err = fmt.Errorf("ops:%d,err:%v,code:%v", ops, err, code)
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
