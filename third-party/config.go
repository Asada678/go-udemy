package main

import (
	"fmt"
	"gopkg.in/go-ini/ini.v1"
)

// ini

type ConfigList struct {
	Port      int
	Dbname    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustInt(8080),
		Dbname:    cfg.Section("db").Key("name").MustString("example.com"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Printf("Port = %v\n", Config.Port)
	fmt.Printf("Dbname = %v\n", Config.Dbname)
	fmt.Printf("SQLDriver = %v\n", Config.SQLDriver)

}
