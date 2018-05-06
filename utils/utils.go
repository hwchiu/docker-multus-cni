package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/containernetworking/cni/pkg/types"
)

type DelegatesObject []map[string]interface{}

type NormalCNIConfig map[string]interface{}

type PLuginCNIConfig struct {
	Name    string                   `json:"name"`
	Plugins []map[string]interface{} `json:"plugins"`
}

type MultusConfig struct {
	types.NetConf
	Delegates  []map[string]interface{} `json:"delegates"`
	Kubeconfig string                   `json:"kubeconfig"`
}

func LoadCNIConfig(path string) (*DelegatesObject, error) {
	//Try to plugin format
	//{
	//  name
	//  plugin [][]
	//}
	//Try to normal format
	//map[string]interface{}
	var data NormalCNIConfig

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	if _, ok := data["plugins"]; !ok {
		//normal plgin
		return &DelegatesObject{
			data,
		}, nil
	}
	//for plugin data, we se the first plugin and add the type into it.
	var pluginData PLuginCNIConfig
	err = json.Unmarshal(file, &pluginData)
	if err != nil {
		return nil, err
	}

	delegates := DelegatesObject{}
	for _, v := range pluginData.Plugins {
		if v["type"] == "portmap" {
			continue
		}
		v["name"] = pluginData.Name
		delegates = append(delegates, v)
	}
	return &delegates, nil
}

func NewMultusObject(config string) (*MultusConfig, error) {
	plugins := []map[string]interface{}{}
	data := MultusConfig{
		types.NetConf{
			Type:       "multus",
			CNIVersion: "0.3.1",
			Name:       "multus-cni",
		},
		plugins,
		config,
	}
	return &data, nil
}

func AddPluginsIntoMults(netConf *MultusConfig, cniObject *DelegatesObject) {
	netConf.Delegates = append(netConf.Delegates, *cniObject...)
}
