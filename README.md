## Docker-Multus-CNI[![Build Status](https://travis-ci.org/hwchiu/docker-multus-cni.svg?branch=master)](https://travis-ci.org/hwchiu/docker-multus-cni) [![codecov](https://codecov.io/gh/hwchiu/docker-multus-cni/branch/master/graph/badge.svg)](https://codecov.io/gh/hwchiu/docker-multus-cni) [![Docker Build Status](https://img.shields.io/docker/build/hwchiu/docker-multus-cni.svg)](https://hub.docker.com/r/hwchiu/docker-multus-cni/)
===================

This repo is used to deply the multus-CNI into your kubernetes cluster.
You can use the `kubespray` to start this docker image to copy the multus-cni's binary and other configs.

## Introduction

This docker image is used to deploy the multus-cni in kubernetes cluster.
In the docker image, it contains the following files.
1. CNI binary (multus)
2. CNI config, this file should be placed under **/etc/cni/net.d/**and it use the **ovs** as its default network.
3. Kubernetes CustomResourceDefinition yaml file, it will create a CRD object called Network and you can add any CNI as a Network object
and use it on your POD yaml files.
4. Our custome binary, it can embed the CNI config to multus config.

For the ovs network, you can refer to the **https://github.com/John-Lin/ovs-cni** to learn more about ovs network.

## How to use.
In order to copy the pre-build files into your system, you need to mount the vloume into the docker container
and tell the docker where you want to copy to.
You have three options to indicate what kind of the data you want to copy.
- b: the binary location
- c: the config location
- k: the CRD yaml location.
- g: generate the multus config from the arguments and pus the config into env `$DEST_CNI`

For example, i want to copy the binary to /tmp/test/bin, config to /tmp/test/conf and the yaml to /tmp/test/yaml. you can start the docker contaienr as below.
(Make sure you have already create the directory /tmp/test/bin, /tmp/test/conf, /tmp/test/yaml)
```
docker run --rm -v /tmp/test/bin:/cni/bin -v /tmp/test/conf:/conf -v /tmp/test/yaml:/yaml hwchiu/docker-multus-cni -b /cni/bin -c /conf -k /yaml
```
And you can see the files under the `/tmp/test/*` after the execution of docker container

Besides, you can copy the `test/calico.yaml` to any exist directory, take `/tmp/multus/` as example.
Use the following command to convert the calico CNI config into a Multus CNI config
```
docker run --rm -v /tmp/multus:/conf -e DEST_CNI=/conf/00-multus.conf hwchiu/docker-multus-cni -g /conf/calico.yaml
```
You can see the Multus CNI config in `/tmp/multus/00-multus.conf` and it looks like
```
{
 "cniVersion": "0.3.1",
 "name": "multus-cni",
 "type": "multus",
 "ipam": {},
 "dns": {},
 "delegates": [
  {
   "etcd_ca_cert_file": "/etc/ssl/etcd/ssl/ca.pem",
   "etcd_cert_file": "/etc/ssl/etcd/ssl/node-k8s-01.pem",
   "etcd_endpoints": "https://172.17.8.101:2379,https://172.17.8.102:2379,https://172.17.8.103:2379",
   "etcd_key_file": "/etc/ssl/etcd/ssl/node-k8s-01-key.pem",
   "ipam": {
    "assign_ipv4": "true",
    "ipv4_pools": [
     "10.233.64.0/18"
    ],
    "type": "calico-ipam"
   },
   "kubernetes": {
    "kubeconfig": "/etc/kubernetes/node-kubeconfig.yaml"
   },
   "log_level": "info",
   "nodename": "k8s-01",
   "type": "calico"
  },
  {
   "capabilities": {
    "portMappings": true
   },
   "type": "portmap"
  }
 ],
 "kubeconfig": "/etc/kubernetes/admin.conf"
}

```
