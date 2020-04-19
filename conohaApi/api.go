package ConohaApi

import (
	"../common"
	"bytes"
	"errors"
	"github.com/spf13/viper"
	"time"
)

func Login() (err error) {

	var pass = DataPasswordCredentials{Username: Common.Config.GetString("username"),
		Password: Common.Config.GetString("password")}
	var loginStruct = struct {
		Auth DataLogin `json:"auth"`
	}{
		Auth: DataLogin{PasswordCredentials: pass, TenantId: Common.Config.GetString("tenantId")},
	}

	body, err := Common.Post(ApiURLs.IdentityToken, Common.FastJsonMarshal(loginStruct))

	if nil == err {

		_, _, err := GetToken(body)
		if nil != err {
			return err
		}

		err = SaveJsonString(FileList.Token, string(body))
		if nil != err {
			return err
		}
	}

	return

}

func logout() {

}

func InitApi() (err error) {

	Common.SetRetryTimes(Common.Config.GetInt("retrytime"))
	Common.SetTimeOut(time.Duration(Common.Config.GetInt("timeout")) * time.Second)
	Common.SetHeader("Content-Type", "application/json")

	data, err := GetJsonString(FileList.Token)
	if nil != err {
		return err
	}
	token, _, err = GetToken([]byte(data))

	if nil != err {
		err = Login()
		if nil != err {
			return err
		}
		data, err = GetJsonString(FileList.Token)
		token, _, err = GetToken([]byte(data))
		if nil != err {

			return err
		}
	}

	if nil == err {
		Common.SetHeader("X-Auth-Token", token)
		tenantId = Common.Config.GetString("tenantId")
	}

	ServerIDList, err = ComputeServerList()
	if nil == err {
		SaveJsonString("serverIDList.json", Common.FastJsonMarshal(ServerIDList))
	}

	return err
}

func GetToken(jsonByte []byte) (token string, expires time.Time, err error) {
	bytes.NewReader(jsonByte)
	data := viper.New()

	data.SetConfigType("json")
	err = data.ReadConfig(bytes.NewReader(jsonByte))
	if nil != err {
		return
	}

	token = data.GetString("access.token.id")
	if len(token) < 3 {
		err = errors.New("token empty")
	}
	expires = data.GetTime("access.token.expires")

	if (expires.Local().Unix() - 600) < time.Now().Local().Unix() {
		err = errors.New("time expires")
	}

	return
}
