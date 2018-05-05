package utils

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadCNIConfig(t *testing.T) {
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

	defer os.Remove(tmpfile.Name()) // clean up

	_, err = tmpfile.Write(content)
	assert.NoError(t, err)
	err = tmpfile.Close()
	assert.NoError(t, err)

	obj, err := LoadCNIConfig(tmpfile.Name())
	assert.NoError(t, err)
	assert.NotNil(t, obj)

	netObj, err := GenerateMultusObject("/etc/config", obj)
	assert.NoError(t, err)
	assert.NotNil(t, netObj)

	assert.Equal(t, "/etc/config", netObj.Kubeconfig)
}
