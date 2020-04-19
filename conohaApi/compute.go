package ConohaApi

import (
	"../common"
	"fmt"
	"github.com/gobs/simplejson"
	"github.com/pkg/errors"
	"log"
	"time"
)

func ComputeFlavorList() {
	apiURL := ApiURLs.Compute + "/" + tenantId + "/flavors"
	body, err := Common.Get(apiURL)
	log.Println(err)
	fmt.Println(string(body))






	apiURL = ApiURLs.Compute + "/" + tenantId + "/flavors/ab7b9b6d-108c-4487-90a4-2da604ad6a92"
	body, err = Common.Get(apiURL)
	log.Println(err)
	fmt.Println(string(body))
	panic("xxxxxxxxxxxxxxxxx")

}
func ComputeImageList() {
	apiURL := ApiURLs.Compute + "/" + tenantId + "/images"
	body, err := Common.Get(apiURL)
	log.Println(err)
	fmt.Println(string(body))






	apiURL = ApiURLs.Compute + "/" + tenantId + "/images/37975142-2199-4ee9-9aa9-bc669e43f139"
	body, err = Common.Get(apiURL)
	log.Println(err)
	fmt.Println(string(body))
	panic("xxxxxxxxxxxxxxxxx")

}

func ComputeAdd() {
	apiURL := ""
	//1gb flover
	//apiURL = ApiURLs.Compute + "/" + tenantId + "/flavors/ab7b9b6d-108c-4487-90a4-2da604ad6a92"
	//apiURL = ApiURLs.Compute + "/" + tenantId + "/flavors"

	//centos7.6 image
	//apiURL = ApiURLs.Compute + "/" + tenantId + "/images/34c551c5-5494-4e32-9f56-4de023708eb4"
	//apiURL = ApiURLs.Compute + "/" + tenantId + "/images"
	//apiURL =  "https://account.tyo1.conoha.io/v1/" + tenantId + "/order-items/4c0e8735-fc2d-409b-937f-6399f85f350c"

	//gbody ,err := Common.Get(apiURL)
	//log.Println(err)
	//log.Println(string(gbody))
	//panic("xxxxxxxxxxxxxxxxx")

	apiURL = ApiURLs.Compute + "/" + tenantId + "/servers"
	data := "{\"server\":{\"imageRef\":\"37975142-2199-4ee9-9aa9-bc669e43f139\",\"flavorRef\":\"ab7b9b6d-108c-4487-90a4-2da604ad6a92\",\"adminPass\":\"adminPassa!@#@#@dminPass\"}}"
	//data := "{\"server\":{\"adminPass\":\"7AwxbUP6M4,R\",\"imageRef\":\"e2b62c96-abbc-41ae-a5f2-b0fe514b755c\",\"flavorRef\":\"294639c7-72ba-43a5-8ff2-513c8995b869\"}}"
	body, err := Common.Post(apiURL, data)
	//
	//apiURL = ApiURLs.Compute + "/" + tenantId + "/servers/83947220-4aa5-46b3-a774-04b875aa83b4"
	//body, err := Common.Delete(apiURL)

	log.Println(err)
	log.Println(string(body))
	panic("xxxxxxxxxxxxxxxxx")
}

func ComputeServerList() (idList []string, err error) {
	apiURL := ApiURLs.Compute + "/" + tenantId + "/servers"
	body, err := Common.Get(apiURL)

	var _serverList RespServerList
	err = Common.Json.Unmarshal(body, &_serverList)
	if nil != err {
		return nil, err
	}

	if len(_serverList.Servers) == 0 {
		var _errResponding RespError
		err = Common.Json.Unmarshal(body, &_errResponding)
		if nil != err {
			return nil, err
		}

		if _errResponding.BadRequest.Code > 0 {
			return nil, errors.New(string(body))
		}

	}

	if len(_serverList.Servers) > 0 {
		for _, v := range _serverList.Servers {
			idList = append(idList, v.ID)
		}
	}
	return

}

func ComputeServerInfo(id string) (serverInfo ServerInfo, err error) {
	apiURL := ApiURLs.Compute + "/" + tenantId + "/servers/" + id
	body, err := Common.Get(apiURL)
	if nil != err {
		return ServerInfo{}, err
	}
	var _errResponding RespError
	err = Common.Json.Unmarshal(body, &_errResponding)
	if nil != err {
		return ServerInfo{}, err
	}

	if _errResponding.BadRequest.Code > 0 {
		return ServerInfo{}, errors.New(string(body))
	}

	respJson, err := simplejson.LoadBytes(body)

	_status, ok := respJson.Get("server").CheckGet("status")
	if false == ok {
		return ServerInfo{}, errors.New(string(body))
	}
	serverInfo.Status = _status.MustString()
	_addressList, ok := respJson.Get("server").CheckGet("addresses")
	for _, v := range _addressList.MustMap() {

		var _address []RespServerInfoAddress
		Common.Json.UnmarshalFromString(Common.FastJsonMarshal(v), &_address)

		if len(_address) > 0 {
			serverInfo.IPV6Address = map[string]string{}
			serverInfo.IPV4Address = map[string]string{}
			for _, _v1 := range _address {
				if 6 == _v1.Version {
					serverInfo.IPV6Address[_v1.OS_EXT_IPS_MAC_macAddr] = _v1.Addr
				}
				if 4 == _v1.Version {
					serverInfo.IPV4Address[_v1.OS_EXT_IPS_MAC_macAddr] = _v1.Addr
				}
			}
		}

	}
	serverInfo.Name = respJson.Get("server").Get("metadata").Get("instance_name_tag").MustString()
	serverInfo.Created, _ = time.Parse(time.RFC3339, respJson.Get("server").Get("created").MustString())
	serverInfo.Updated, _ = time.ParseInLocation(time.RFC3339, respJson.Get("server").Get("updated").MustString(), time.Local)
	//SaveJsonString("serverInfo-"+id+".json", string(body))

	return serverInfo, nil

}

func ComputeServerReboot(id string) (err error) {
	apiURL := ApiURLs.Compute + "/" + tenantId + "/servers/" + id + "/action"

	body, err := Common.Post(apiURL, "{\"reboot\": {\"type\": \"SOFT\"}}")

	if len(body) > 3 {
		return errors.New(string(body))
	}
	return
}
func ComputeServerDelete(id string) (err error) {
	apiURL := ApiURLs.Compute + "/" + tenantId + "/servers/" + id

	body, err := Common.Delete(apiURL)
	log.Println("errrrrrrrrrrr", err)
	if len(body) > 3 {
		return errors.New(string(body))
	}
	return
}

func ComputeServerForceShutDown(id string) (err error) {
	apiURL := ApiURLs.Compute + "/" + tenantId + "/servers/" + id + "/action"

	body, err := Common.Post(apiURL, "{\"os-stop\": {\"force_shutdown\": true}}")

	if len(body) > 3 {
		return errors.New(string(body))
	}
	return
}

func ComputeServerStart(id string) (err error) {
	apiURL := ApiURLs.Compute + "/" + tenantId + "/servers/" + id + "/action"

	body, err := Common.Post(apiURL, "{\"os-start\": null}")

	log.Println(string(body))
	log.Println(err)
	if len(body) > 3 {
		return errors.New(string(body))
	}
	return
}
