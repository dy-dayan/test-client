package unitTest

import (
"github.com/golang/protobuf/proto"
"github.com/dy-dayan/testClient/client"
"github.com/dy-dayan/testClient/proto"
	"os"
)


func KeepAlive(c client.Client){
	dev := access.Device{
		Guid:                 "jerry",
	}
	account := access.Account{
		Email:                "",
		PhoneNum:             "",
		Password:             "",
		Token:                "",
	}

	reqHead := access.ReqHead{
		Dev:                  &dev,
		Ver:                  "0.0.0",
		Account:              &account,
	}
	content := []byte("hello")

	reqPkgHead := access.PkgReqHead{
		Seq:                  1,
	}

	reqBodyOne := access.ReqBody{
		Service:              "state",
		Method:               "keepalive",
		Content:              content,
	}


	reqBodys := []*access.ReqBody{&reqBodyOne}
	reqPkgBody := access.PkgReqBody{
		Head:                 &reqHead,
		Bodys:                reqBodys,
	}

	req := access.PkgReq{
		Head:                 &reqPkgHead,
		Body:                 &reqPkgBody,
	}

	reqMsg,_ := proto.Marshal(&req)
	_, err := c.PostMessage(reqMsg)
	if err != nil{
		os.Exit(-1)
	}
}