package smf

import (
	"fiveGCHub/common"
	"fmt"
	"io/ioutil"
	"net/http"
)

const SmfURL = "http://127.0.0.1:5000/#/"

type SmfCrawler struct {
	Name string
	Conf *SmfConfig
}

type SmfConfig struct {
	Interval int
	Version  string
	Enable   bool
}

func NewSmfCrawler(name string, cfg *SmfConfig) *SmfCrawler {
	return &SmfCrawler{name, cfg}
}

func (sc *SmfCrawler) Run() {
	if !sc.Conf.Enable {
		return
	}
	// Just an example, this is actually webui
	// TODO: Replace with real SMF
	resp, err := http.Get(SmfURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)
}

func (sc *SmfCrawler) GetInterval() int {
	return sc.Conf.Interval
}

func GetSmfConf() *SmfConfig {
	// TODO: use viper to get config
	cfg := &SmfConfig{5, "0.0.1", true}
	return cfg
}

func init() {
	name := "smf_retrieve"
	cfg := GetSmfConf()
	common.Register(name, NewSmfCrawler(name, cfg))
}
