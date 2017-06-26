package main

import (
	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
)

type Config struct {
	BuyHost  string
	TimeHost string
	Mysql    Database
}

type Database struct {
	Host string
	User string
	Pswd string
	Name string
}

func ConfigInit(filename string) (Config, error) {
	var conf Config
	_, err := toml.DecodeFile(filename, &conf)
	if err != nil {
		glog.Errorln(err)
		return conf, err
	}
	return conf, nil

}
