package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/hwchiu/docker-multus-cni/utils"
)

func main() {
	var srcCNIDir string
	var outputCNI string
	var kubeconfig string

	flag.StringVar(&srcCNIDir, "srcDir", "/etc/cni/net.d/", "The souruce directory contains CNI configs")
	flag.StringVar(&outputCNI, "output", "/etc/cni/net.d/00-multus.conf", "The output location of multus CNI config")
	flag.StringVar(&kubeconfig, "kubeconfig", "/etc/kubernetes/admin.conf", "The kubernetes config location for Multus CNI")
	flag.Parse()

	if srcCNIDir == "" {
		log.Fatal("You need to specify the input CNI path")
	}

	netObj, err := utils.NewMultusObject(kubeconfig)
	if err != nil {
		log.Fatal("Fail to new the multus object")
	}

	err = filepath.Walk(srcCNIDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		cniObject, err := utils.LoadCNIConfig(path)
		if err != nil {
			log.Printf("Decoding the CNI  %s fail %v\n", path, err)
			return nil
		}
		utils.AddPluginsIntoMults(netObj, cniObject)
		return nil
	})

	//Add the masterplugin
	if len(netObj.Delegates) > 0 {
		netObj.Delegates[0]["masterplugin"] = true
	}
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
