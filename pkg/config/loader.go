package config

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/iwanhae/go-statusboard/pkg/monitor"
)

type ConfigSchema struct {
	Checks struct {
		HTTP []HttpConfig `json:"http"`
	} `json:"checks"`
}

type HttpConfig struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Method      string `json:"method"`
	Endpoint    string `json:"endpoint"`
	Timeout     string `json:"timeout"`
	Interval    string `json:"interval"`
}

func LoadConfig(path string) ([]monitor.Checker, error) {
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	c := &ConfigSchema{}
	err = json.Unmarshal(b, c)
	if err != nil {
		return nil, err
	}

	res := []monitor.Checker{}

	for _, v := range c.Checks.HTTP {
		interval, err := time.ParseDuration(v.Interval)
		if err != nil {
			return nil, err
		}
		timeout, err := time.ParseDuration(v.Timeout)
		if err != nil {
			return nil, err
		}
		res = append(res, monitor.NewSimpleHttpChecker(v.Name, v.Description, v.Method, v.Endpoint, interval, timeout))
	}
	return res, nil
}
