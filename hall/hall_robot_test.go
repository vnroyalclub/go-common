package hall

import "testing"

func TestLoadRobotsID(t *testing.T) {
	hallUrl:="http://172.13.0.53:9527/"
	lvForm:=int32(1)
	lvTo:=int32(2)
	ret,err:=LoadRobotsID(lvForm,lvTo,hallUrl)
	if err!=nil{
		t.Error("failed to load robots Id,err:",err)
	}

	t.Log("ret:",ret.String())

}

func TestModRobotHallProps(t *testing.T) {

}