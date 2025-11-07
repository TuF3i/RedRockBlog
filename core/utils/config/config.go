package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func InitConfig(ConfigPath string) (Config, error) {
	//初始化配置文件
	root := ConfigRoot{}
	root.configPath = ConfigPath
	root.conf = Config{}
	err := root.readConfig()
	return root.conf, err
}

func (root *ConfigRoot) readConfig() error {
	//读取配置文件
	data, err := os.ReadFile(root.configPath)
	if err != nil {
		return fmt.Errorf("read Config File Error - %v", err.Error())
	}

	//解析配置到json对象
	err = json.Unmarshal(data, &root.conf)
	if err != nil {
		return fmt.Errorf("unmarshal JSON Obj Error - %v", err.Error())
	}

	root.conf.RedirectURL = fmt.Sprintf("http://%v:%v/callback", root.conf.Domain, root.conf.ApiListeningPort)
	return nil
}
