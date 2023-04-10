package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"mediafiledownload/config"
)

var configFile = flag.String("f", "config.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

}
