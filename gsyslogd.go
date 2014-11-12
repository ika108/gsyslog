package main

import (
	"flag"
	"fmt"
	"github.com/ika108/gsyslog/libs/codec"
	"github.com/ika108/gsyslog/libs/filter"
	"github.com/ika108/gsyslog/libs/input"
	"github.com/ika108/gsyslog/libs/message"
	"github.com/ika108/gsyslog/libs/output"
	"github.com/ika108/gsyslog/libs/queue"
        "encoding/json"
	"os"
)

type config struct {
	configFile string
	codecs []codec
	filters []filter
	inputs []input
	outputs []output
	queues []queue
}

var configFile string
var currentConfig config

func init() {
}

func main() {

}

func parseConfig(configFile string) {
	flag.StringVar(&configFile, "configfile", `gsyslogd.conf`, "Define a configfile")
	flag.Parse()
	fileObj, _ := os.Open(configFile)
	jsonDecoder := json.NewDecoder(fileObj)
	
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
	  fmt.Println("error:", err)
	}
	fmt.Println(configuration.Users) // output: [UserA, UserB]
}

func (myConfig *config) addCodec(codec *newCodec) {
	config.codecs = append(config.codecs,newcodec)
}

func (myConfig *config) addFilter(filter *newFilter) {
	config.filters = append(config.filters,newFilter)
}

func (myConfig *config) addInput(input *newInput) {
	config.inputs = append(config.inputs,newInput)
}

func (myConfig *config) addOutput(output *newOutput) {
	config.outputs = append(config.outputs,newOutput)
}

func (myConfig *config) addQueue(queue *newQueue) {
	config.queues = append(config.queues,newQueue)
}


