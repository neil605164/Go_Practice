package util

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var configFile []byte

type ChannelConfig struct {
	Channel Channel `yaml:"channel"`
}

type Channel struct {
	EmayReminderConfig EmayReminder `yaml:"emayReminder"`
	GuoduConfig        Guodu        `yaml:"guodu"`
}

type EmayReminder struct {
	UserId    string `yaml:"userId"`
	UserPws   string `yaml:"userPws"`
	Url       string `yaml:"url"`
	Threshold string `yaml:"threshold"`
}

type Guodu struct {
	UserId    string `yaml:"userId"`
	UserPws   string `yaml:"userPws"`
	Url       string `yaml:"url"`
	KeyStr    string `yaml:"keyStr"`
	Threshold string `yaml:"threshold"`
}

func GetChannelConfig() (e *ChannelConfig, err error) {
	err = yaml.Unmarshal(configFile, &e)
	return e, err
}

func init() {
	var err error
	configFile, err = ioutil.ReadFile("resource/config.yaml")
	if err != nil {
		log.Fatalf("yamlFile.Get err %v ", err)
	}
}
