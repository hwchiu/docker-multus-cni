package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/containernetworking/cni/pkg/types"
)

type CNIObject map[string]interface{}

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

func NewMultusObject(config string) (*NetConf, error) {
	plugins := []map[string]interface{}{}
	data := NetConf{
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

func AddPluginsIntoMults(netConf *NetConf, cniObject *CNIObject) {
	netConf.Delegates = append(netConf.Delegates, *cniObject)
}
