package main

import "fmt"
import "os"
import "io/ioutil"
import "encoding/json"

type config map[string]interface{}

type global_conf map[string]interface{}

type codecs_conf []interface{}

type filters_conf []interface{}

type inputs_conf []interface{}

type outputs_conf []interface{}

type queues_conf []interface{}

type settings struct {
	global global_conf
	codecs codecs_conf
	filters filters_conf
	inputs inputs_conf
	outputs outputs_conf
	queues queues_conf
}

func main() {

	myFile, err := readFile("gsyslogd.conf")
	if err != nil {
		fmt.Printf("Error with file : %v\n", err)
		panic("Error")
	}
	// fmt.Println(string(myFile))
	myStruct, err := importJSON(myFile)
	if err != nil {
		fmt.Printf("Error with file : %v\n", err)
		panic("Error")
	}

	config := parseJSON(myStruct)
	fmt.Println(config)

}

func readFile(fileName string) ([]byte, error) {
	fileReader, err := os.Open(fileName)
	defer fileReader.Close()
	if err != nil {
		return nil, fmt.Errorf("Error opening file %v : %v", fileName, err)
	}
	fileContent, err := ioutil.ReadAll(fileReader)

	if err != nil {
		return nil, fmt.Errorf("Error reading config file %v : %v", fileName, err)
	}

	return fileContent, nil
}

func importJSON(jsonBuf []byte) (config, error) {

	var myStruct config

	err := json.Unmarshal(jsonBuf, &myStruct)
	if err != nil {
		return myStruct, fmt.Errorf("Error parsing config file : %v", err)
	}
	return myStruct, nil
}

func parseJSON(config config) (settings){
	//fmt.Println("fullcontent:",*config)
	var settings settings
	for key, value := range map[string]interface{}(config) {
		// fmt.Printf("clef : %v value type : %T\n", key, value)
		switch {
		case key == "global":
			fmt.Printf("%s => %v (\033[31m%T\033[0m)\n",key, value,value)
		case key == "codecs":
			fmt.Printf("%s => %v (\033[31m%T\033[0m)\n",key, value,value)
		case key == "filters":
			fmt.Printf("%s => %v (\033[31m%T\033[0m)\n",key, value,value)
		case key == "inputs":
			fmt.Printf("%s => %v (\033[31m%T\033[0m)\n",key, value,value)
		case key == "outputs":
			fmt.Printf("%s => %v (\033[31m%T\033[0m)\n",key, value,value)
		case key == "queues":
			fmt.Printf("%s => %v (\033[31m%T\033[0m)\n",key, value,value)
		}
	}
	return settings
}
