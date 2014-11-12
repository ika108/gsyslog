package main

import {
    "fmt"
    "flag"
    "github.com/ika108/gsyslog/libs/input"
    "github.com/ika108/gsyslog/libs/codec"
    "github.com/ika108/gsyslog/libs/output"
    "github.com/ika108/gsyslog/libs/filter"
    "github.com/ika108/gsyslog/libs/queue"
    "github.com/ika108/gsyslog/libs/message"
}

type config {}

var configfile string
var configPtr config

func init () {
}

func main () {
    
}

func parseConfig (configFile string){
	flag.StringVar(&configfile, "configfile", `gsyslogd.conf`, "Define a configfile")
	flag.Parse()
	configPtr = parseConfig(configfile)
}