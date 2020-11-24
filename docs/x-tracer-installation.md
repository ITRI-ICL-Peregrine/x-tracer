
## Pre-Requisite

1. Python/pip3
2. Kubernetes cluster with one master and one node, you can deploy it using [kubespray](https://sheenampathak.com/?p=598), for more detailed installation steps [click here](https://github.com/kubernetes-sigs/kubespray#quick-start)
3. go - [Install go](https://sheenampathak.com/?p=587) on the master node

## Software Requirements


| Packages | Version | 
| -------- | --------|
|Ubuntu OS | 18.04 LTS
|linux    | 4.15.0-124-generic
| Kubernetes     | v1.19    |
|docker| v19.03.13
| go    | v1.15.1
|kubespray | v2.14.2
|x-tracer-gocui| v1.0

x-tracer installation
---
x-tracer needs to be installed on Master node of the cluster.

You have to replace your $DOCKER_ID in Makefile
```
go get github.com/ITRI-ICL-Peregrine/x-tracer
cd $GOPATH/src/github.com/ITRI-ICL-Peregrine/x-tracer
make
make release
make publish
./bin/x-tracer
```
x-tracer CUI will be started, now you can choose the pod and probes.


