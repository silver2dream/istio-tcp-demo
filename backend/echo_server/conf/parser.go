package conf

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"

	"gopkg.in/yaml.v2"
)

var protocolConf Protocol

// ConfigParser  parser yaml config file into config struct
func ConfParser(file string, in interface{}, config interface{}) error {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		msg := fmt.Sprintf("failed to read '%s' with err: %s", file, err.Error())
		return errors.New(msg)
	}
	err = yaml.UnmarshalStrict(yamlFile, in)
	if err != nil {
		msg := fmt.Sprintf("failed to unmarshal '%s' with err: %s", file, err.Error())
		return errors.New(msg)
	}

	configArray := reflect.ValueOf(protocolConf)
	values := make([]interface{}, configArray.NumField())
	for i := 0; i < configArray.NumField(); i++ {
		values[i] = configArray.Field(i).Interface()
		proto := values[i].(Conf)
		if proto.Enable == true {
			config = proto
		}
	}
	return nil
}
