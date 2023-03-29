package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvd_project/config"
	"gvd_project/global"
	"io/ioutil"
	"log"
)

// 读取yaml文件配置
func CoreConf() {
	const ConfigFIle = "settings.yaml"
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
