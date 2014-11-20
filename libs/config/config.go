package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type config struct {
	Array []string `json:"array"`
	Hash  struct {
		String1 string `json:"string1"`
		String2 string `json:"string2"`
	} `json:"hash"`
	Maxthread  float64 `json:"maxthread"`
	Stringtest string  `json:"stringtest"`
}

func NewConfig() (*config, error) {
	var myConfigPtr *config
	var configFilePtr = flag.String("configfile", "gsyslogd.conf", "Path to the configuration file")
	flag.Parse()

	fileReader, err := os.Open(*configFilePtr)
	if err != nil {
		return myConfigPtr, fmt.Errorf("Error opening file", *configFilePtr, err)
	}
	defer fileReader.Close()

	fileContent, err := ioutil.ReadAll(fileReader)
	
	if err != nil {
		return myConfigPtr, fmt.Errorf("Error reading config file", err)
	}

	err = json.Unmarshal(fileContent, &myConfigPtr)

	fmt.Println(myConfigPtr.Maxthread,myConfigPtr.Stringtest,myConfigPtr.Hash.String1,myConfigPtr.Array)


	if err != nil {
		return myConfigPtr, fmt.Errorf("Error parsing config file", err)
	}
	return myConfigPtr, nil
}
