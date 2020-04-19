package ConohaApi

import (
	"../common"
	"github.com/spf13/viper"
	"io/ioutil"
)

func GetJsonData(fileName string) (data *viper.Viper, err error) {
	data = viper.New()

	data.SetConfigType("json")
	data.SetConfigFile(trans2RealPath(fileName))
	err = data.ReadInConfig()

	return

	//if err != nil { // Handle errors reading the config file
	//	log.Println(fPath)
	//	panic(fmt.Errorf("Fatal error config file: %s \n", err))
	//}
}

func SaveJsonData(fileName string, data *viper.Viper) (err error) {
	err = data.WriteConfigAs(trans2RealPath(fileName))
	return
}

func GetJsonString(fileName string) (data string, err error) {
	bData, err := ioutil.ReadFile(trans2RealPath(fileName))
	return string(bData), err
}

func SaveJsonString(fileName string, data string) (err error) {

	err = ioutil.WriteFile(trans2RealPath(fileName), []byte(data), 755)
	return
}

func trans2RealPath(fileName string) (fPath string) {
	c, _ := Common.GetCurrentPath()
	fPath = c + "/data/" + fileName
	return
}
