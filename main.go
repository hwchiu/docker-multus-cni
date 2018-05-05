package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/hwchiu/docker-multus-cni/utils"
)

func main() {
	var inputCNI string
	var outputCNI string
	var kubeconfig string

	flag.StringVar(&inputCNI, "input", "", "The cni config we want to embed into multus CNI config")
	flag.StringVar(&outputCNI, "output", "/etc/cni/net.d/00-multus.conf", "The output location of multus CNI config")
	flag.StringVar(&kubeconfig, "kubeconfig", "/etc/kubernetes/admin.conf", "The kubernetes config location for Multus CNI")
	flag.Parse()

	if inputCNI == "" {
		log.Fatal("You need to specify the input CNI path")
	}

	cniObject, err := utils.LoadCNIConfig(inputCNI)
	if err != nil {
		log.Fatal("You need to specify the input CNI path")
	}

	netObj, err := utils.GenerateMultusObject(kubeconfig, cniObject)

	cniFile, err := os.Create(outputCNI)
	if err != nil {
		panic(err)
	}
	defer cniFile.Close()

	jsonData, err := json.MarshalIndent(netObj, "", " ")
	if err != nil {
		panic(err)
	}

	cniFile.Write(jsonData)
}
