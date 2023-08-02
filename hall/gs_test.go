package hall

import "testing"

func TestGetGameAddress(t *testing.T) {

	ret,err:=GetGameAddress(2004, 11, "1103", "http://172.13.0.53:9000/2004/")
	if err!=nil{
		t.Error("get game address:",err)
		return
	}

	t.Log(ret.String())

}

func TestRegisterGame(t *testing.T) {

	gsUrl:="http://172.13.0.53:9000/2004/"
	groupId:=int32(11)
	serverId:="1103"
	port:=int32(1098)


	err:=RegisterGame(groupId,serverId,port,gsUrl)

	if err!=nil{
		t.Error("failed to redister game,err:",err)
		return
	}

}
