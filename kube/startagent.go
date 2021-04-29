package kube

import(
	"github.com/ITRI-ICL-Peregrine/x-tracer/internal/agentmanager"
//	"github.com/jroimartin/gocui"
//	"io"
//	"fmt"
	"strings"
	"github.com/ITRI-ICL-Peregrine/x-tracer/getval"
)


func StartAgent(p string,probes string) error {
	cs := GetClientSet()
	var containerId []string

	containerId = GetPodContainersID(p)

	targetNode := GetTargetNode(p)

	nodeIp := GetNodeIp()

	if probes == "All Probes" {

		pn := getval.GetProbeNames()
		allpn := strings.Join(pn, ",")
		agent := agentmanager.New(containerId[0], targetNode, nodeIp, cs, allpn)

		//Start x-agent Pod
		//fmt.Fprintln(o, "Starting x-agent Pod...")

		agent.ApplyAgentPod()

		//fmt.Fprintln(o, "Starting x-agent Service...")
		agent.ApplyAgentService()

		agent.SetupCloseHandler()

	} else if probes == "All TCP Probes" {
		pn := getval.GetTcpProbeNames()
		tcppn := strings.Join(pn, ",")
		agent := agentmanager.New(containerId[0], targetNode, nodeIp, cs, tcppn)

		//fmt.Fprintln(o, "Starting x-agent Pod...")

		agent.ApplyAgentPod()

		//fmt.Fprintln(o, "Starting x-agent Service...")
		agent.ApplyAgentService()

		agent.SetupCloseHandler()

	} else {
		agent := agentmanager.New(containerId[0], targetNode, nodeIp, cs, probes)

		//Start x-agent Pod
		//fmt.Fprintln(o, "Starting x-agent Pod...")

		agent.ApplyAgentPod()

		//fmt.Fprintln(o, "Starting x-agent Service...")
		agent.ApplyAgentService()

		agent.SetupCloseHandler()

	}

	return nil




}
