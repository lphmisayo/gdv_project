package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvd_server/config"
	"gvd_server/global"
	"io/ioutil"
	"log"
)

const (
	ConfigFIle = "settings.yaml"
)

// InitConf 读取yaml文件配置
func InitConf() {
	//const ConfigFIle = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFIle)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error : %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal : %v", err)
	}
	log.Println("config yamlFile load Init success")
	global.Config = c
}

func SetYaml() error {
	//const ConfigFIle = "settings.yaml"
	data, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(ConfigFIle, data, 0777)
	if err != nil {
		return err
	}
	global.Log.Infoln("update settings success!")
	return nil
}
