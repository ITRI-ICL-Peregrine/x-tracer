package xmgr

import (

"github.com/ITRI-ICL-Peregrine/x-tracer/kube"
"fmt"
)



func GetPods(namespace string) ([]string){

	kube.NAMESPACE = namespace
	pods,err := kube.GetPods()
	if err != nil {
                fmt.Println(err)
               // return "empty"
        }
	var n []string
	if len(pods.Items) > 0 {
		for _, pod := range pods.Items {
			pod_name := pod.GetName()
			n = append(n, pod_name)

		}
	}

	return n
}


func GetNamespaces() ([]string) {

	namespaces, err := kube.GetNamespaces()
	if err != nil {
		fmt.Println(err)
	}
	var ns []string
	if len(namespaces.Items) > 0 {
			for _, namespace := range namespaces.Items {
				//fmt.Println(namespace.GetName())
				ns = append(ns, namespace.GetName())
			}
	}

	return ns


}


func GetProbes() []string{

	pn := []string{"tcptracer", "tcpconnect", "tcpaccept", "tcplife", "execsnoop", "biosnoop", "cachestat", "All TCP Probes", "All Probes"}
        return pn
}



