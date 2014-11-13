package config

import (
    	"flag"
        "encoding/json"
)

type config struct {
	configFile string
	global map 
	codecs []codec
	filters []filter
	inputs []input
	outputs []output
	queues []queue
}

func parseConfig() *config, error {
    make myConfig config
    var configFilePtr = flag.StringVar("configfile", "gsyslogd.conf, "Path to the configuration file")
	flag.Parse()
	fileObj, fileErr := os.Open(*configFilePtr)
	if fileErr != nil {
	    return _,fmt.Errorf("Error opening file",*configFilePtr,fileErr)
	}
	
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