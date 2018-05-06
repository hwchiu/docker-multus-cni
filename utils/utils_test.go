package utils

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func writeNoramCNI(t *testing.T) string {
	//Generate a sample file
	tmpfile, err := ioutil.TempFile(".", "calico-cni")
	assert.NoError(t, err)

	content := []byte(`
	{
  "name": "cni0",
  "cniVersion":"0.3.1",
  "type": "calico",
  "ipam": {
	  "type": "calico-ipam"
  }
}
	`)

	_, err = tmpfile.Write(content)
	assert.NoError(t, err)
	err = tmpfile.Close()
	assert.NoError(t, err)

	return tmpfile.Name()
}

func writePluginCNI(t *testing.T) string {
	//Generate a sample file
	tmpfile, err := ioutil.TempFile(".", "calico-cni")
	assert.NoError(t, err)

	content := []byte(`
	{
  "name": "cni0",
  "cniVersion":"0.3.1",
  "plugins":[
    {
          "nodename": "k8s-01",
          "type": "calico",
      "etcd_endpoints": "https://172.17.8.101:2379,https://172.17.8.102:2379,https://172.17.8.103:2379",
      "etcd_cert_file": "/etc/ssl/etcd/ssl/node-k8s-01.pem",
      "etcd_key_file": "/etc/ssl/etcd/ssl/node-k8s-01-key.pem",
      "etcd_ca_cert_file": "/etc/ssl/etcd/ssl/ca.pem",
      "log_level": "info",
      "ipam": {
        "type": "calico-ipam",
        "assign_ipv4": "true",
        "ipv4_pools": ["10.233.64.0/18"]
      },
              "kubernetes": {
        "kubeconfig": "/etc/kubernetes/node-kubeconfig.yaml"
      }
    },
    {
      "type":"portmap",
      "capabilities":{
        "portMappings":true
      }
    }
  ]
}
	`)

	_, err = tmpfile.Write(content)
	assert.NoError(t, err)
	err = tmpfile.Close()
	assert.NoError(t, err)

	return tmpfile.Name()
}

func TestLoadCNIConfig(t *testing.T) {
	//Generate a sample file

	pluginFile := writePluginCNI(t)
	defer os.Remove(pluginFile) // clean up

	obj, err := LoadCNIConfig(pluginFile)
	assert.NoError(t, err)
	assert.NotNil(t, obj)

	normalFile := writeNoramCNI(t)
	defer os.Remove(normalFile) // clean up

	obj2, err := LoadCNIConfig(normalFile)
	assert.NoError(t, err)
	assert.NotNil(t, obj2)

	netObj, err := NewMultusObject("/etc/config")
	assert.NoError(t, err)
	assert.NotNil(t, netObj)

	assert.Equal(t, "/etc/config", netObj.Kubeconfig)
	AddPluginsIntoMults(netObj, obj)
	AddPluginsIntoMults(netObj, obj2)
	assert.Equal(t, 2, len(netObj.Delegates))
	for _, v := range netObj.Delegates {
		assert.NotEqual(t, "", v["name"])
	}
}
