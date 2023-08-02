/*
  涉及logserver相关的请求
  发送日志
*/

package hall

import (
	"encoding/json"
	"fmt"
	"net/http"

	gproto "git.huoys.com/vn/proto"
	"github.com/vnroyalclub/go-common/httputil"
	"github.com/vnroyalclub/go-common/log"
)

//发送日志
func SendLogs(logMsgs gproto.LogMsgs, logServerUrl string) (err error) {
	body, err := json.Marshal(logMsgs)
	if err != nil {
		log.Error("failed to marshal logmsgs,err:", err)
		return
	}
	header := map[string]string{}
	code, _, err := httputil.Request(http.MethodPost, logServerUrl, string(body), header)
	if code != http.StatusOK || err != nil {
		log.Error("failed to send log msg,code:", code, "err:", err)
		err = fmt.Errorf("code:%d,err:%v", code, err)
		return
	}

	return
}
