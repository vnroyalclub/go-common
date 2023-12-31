package hall

import (
	gproto "github.com/vnroyalclub/go_proto"

	"testing"
)

func TestSendLogs(t *testing.T) {
	url := "http://127.0.0.1:1234/_internal/record"
	logs := gproto.LogMsgs{}

	err := SendLogs(logs, url)
	if err != nil {
		t.Error("send logs message err:", err)
	}
}
