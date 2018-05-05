package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/containernetworking/cni/pkg/types"
)

type CNIObject struct {
	CNIVersion string                   `json:"cniVersion,omitempty"`
	Name       string                   `json:"name,omitempty"`
	Plugins    []map[string]interface{} `json:"plugins,omitempty"`
}

type NetConf struct {
	types.NetConf
	Delegates  []map[string]interface{} `json:"delegates"`
	Kubeconfig string                   `json:"kubeconfig"`
}

func LoadCNIConfig(path string) (*CNIObject, error) {
	var data CNIObject

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &data)
	return &data, err
}

func GenerateMultusObject(config string, cniObject *CNIObject) (*NetConf, error) {
	data := NetConf{
		types.NetConf{
			Type:       "multus",
			CNIVersion: "0.3.1",
			Name:       "multus-cni",
		},
		cniObject.Plugins,
		config,
	}
	return &data, nil
}
