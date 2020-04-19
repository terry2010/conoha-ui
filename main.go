// Demo code for the Modal primitive.
package main

import (
	"./common"
	"./conohaApi"
	"./gui"
)

var err error

func main() {
	err = Common.InitConfig()

	if nil != err {
		Gui.ErrorAlert(err.Error())
		return
	}

	//ConohaApi.Login()
	err = ConohaApi.InitApi()
	if nil != err {
		Gui.ErrorAlert(err.Error())
		return
	}

	//ConohaApi.ComputeAdd()

	//_info,_ := ConohaApi.ComputeServerInfo("4c0e8735-fc2d-409b-937f-6399f85f350c")
	//log.Println(Common.FastJsonMarshal(_info))

	//ConohaApi.ComputeOrderItems()

	Gui.Run()

	//_info, err := ConohaApi.ComputeServerInfo("4c0e8735-fc2d-409b-937f-6399f85f350c")

	//err := ConohaApi.ComputeServerReboot("fc1054f1-0fa0-4b43-b3b0-d8a3b95a742b")

}
