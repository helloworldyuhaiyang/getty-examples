/******************************************************
# DESC       : env var & configure
# MAINTAINER : Alex Stocks
# LICENCE    : Apache License 2.0
# EMAIL      : alexstocks@foxmail.com
# MOD        : 2016-09-06 16:53
# FILE       : config.go
******************************************************/

package main

import (
	"fmt"
	"os"
	"path"
	"time"
)

import (
	"github.com/AlexStocks/gocolor"
	log "github.com/AlexStocks/log4go"
	config "github.com/koding/multiconfig"
)

const (
	APP_CONF_FILE     string = "APP_CONF_FILE"
	APP_LOG_CONF_FILE string = "APP_LOG_CONF_FILE"
)

var (
	conf *Config
)

type (
	// Config holds supported types by the multiconfig package
	Config struct {
		Host        string   `default:"127.0.0.1"`
		Ports       []string `default:["10000"]`
		ProfilePort int      `default:"10086"`

		SessionTimeout string `default:"60s"`
		sessionTimeout time.Duration
	}
)

func initConf() {
	var (
		err      error
		confFile string
	)

	// configure
	confFile = os.Getenv(APP_CONF_FILE)
	if confFile == "" {
		panic(fmt.Sprintf("application configure file name is nil"))
		return // I know it is of no usage. Just Err Protection.
	}
	if path.Ext(confFile) != ".toml" {
		panic(fmt.Sprintf("application configure file name{%v} suffix must be .toml", confFile))
		return
	}
	conf = new(Config)
	config.MustLoadWithPath(confFile, conf)
	conf.sessionTimeout, err = time.ParseDuration(conf.SessionTimeout)
	if err != nil {
		panic(fmt.Sprintf("time.ParseDuration(SessionTimeout{%#v}) = error{%v}", conf.SessionTimeout, err))
		return
	}
	gocolor.Info("config{%#v}\n", conf)

	// log
	confFile = os.Getenv(APP_LOG_CONF_FILE)
	if confFile == "" {
		panic(fmt.Sprintf("log configure file name is nil"))
		return
	}
	if path.Ext(confFile) != ".xml" {
		panic(fmt.Sprintf("log configure file name{%v} suffix must be .xml", confFile))
		return
	}
	log.LoadConfiguration(confFile)

	return
}