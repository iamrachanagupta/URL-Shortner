package utilities

import (
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"

	"gopkg.in/yaml.v2"
)

//The App_config structure is based on the config.yaml file.
//Any update in the config file needs an update to this structure
type App_config struct {
	AppURL   string `yaml:"appURL"`
	AppPort  string `yaml:"appPort"`
	Loglevel uint32 `yaml:"loglevel"`
}

//Placeholder for all the configs
var Data = make(map[string]App_config)

//ReadConfig reads the config from the project location config/config.yaml and
//updates the global variable Data with all the configurations in key value pair
func ReadConfig() {

	pwd, _ := os.Getwd()
	log.Debug("Current working Directory for ReadConfig:" + pwd)
	configfile := pwd + "/config/config.yaml"
	yfile, err := ioutil.ReadFile(configfile)

	//If config.yaml file is not found then panic occurs
	if err != nil {
		panic("Invalid config path:" + err.Error())
	}
	log.Debug("Reading config file: " + configfile)
	err = yaml.Unmarshal(yfile, &Data)

	//If unmarshaling fails then panic occurs
	if err != nil {
		log.Error("Unmarshal: %s", err)
		panic("Unmarshaling config failed with error: " + err.Error())
	}

}

//ReadTestConfig reads the config from the project location config/config.yaml and
//updates the global variable Data with all the configurations in key value pair.
//
//This is only used in running the unit test cases
func ReadTestConfig() {

	os.Chdir("../config")
	pwd, _ := os.Getwd()
	log.Debug("Current working Directory:" + pwd)
	configfile := pwd + "/config.yaml"
	yfile, _ := ioutil.ReadFile(configfile)
	fmt.Println("Reading config file :" + configfile)
	err := yaml.Unmarshal(yfile, &Data)
	if err != nil {
		log.Error("Unmarshal: %s", err)
	}
}
