## Introduction
This docker image is used to deploy the multus-cni in kubernetes cluster.
In the docker image, it contains the following files.
1. CNI binary (multus)
2. CNI config, this file should be placed under **/etc/cni/net.d/**and it use the **ovs** as its default network.
3. Kubernetes CustomResourceDefinition yaml file, it will create a CRD object called Network and you can add any CNI as a Network object
and use it on your POD yaml files.

For the ovs network, you can refer to the **https://github.com/John-Lin/ovs-cni** to learn more about ovs network.

## How to use.
In order to copy the pre-build files into your system, you need to mount the vloume into the docker container
and tell the docker where you want to copy to.
You have three options to indicate what kind of the data you want to copy.
- b: the binary location
- c: the config location
- k: the CRD yaml location.

For example, i want to copy the binary to /tmp/test/bin, config to /tmp/test/conf and the yaml to /tmp/test/yaml. you can start the docker contaienr as below.
(Make sure you have already create the directory /tmp/test/bin, /tmp/test/conf, /tmp/test/yaml)
```
sudo docker run --rm -v /tmp/test/bin:/cni/bin -v /tmp/test/conf:/conf -v /tmp/test/yaml:/yaml hwchiu/docker-multus-cni -b /cni/bin -c /conf -k /yaml
```
And you can see the files under the /tmp/test/* after the execution of docker container
