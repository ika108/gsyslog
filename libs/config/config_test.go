package config

import "testing"

func TestLoadconfig(test *testing.T){
    myConfig,err := NewConfig()
    if err != nil {
        test.Fatalf("Loading conf has failed : %v\n",err)
    }
    test.Log(*myConfig)
}