package unitTest

import (
	"github.com/golang/protobuf/proto"
	"github.com/dy-dayan/test-client/client"
	"github.com/dy-dayan/test-client/proto"
)


func Hello(c client.Client){
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
		Service:              "greeter",
		Method:               "Greeter.Hello",
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
	c.PostMessage(reqMsg)
}