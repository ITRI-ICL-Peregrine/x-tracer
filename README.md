[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) 

[![Build Status](https://travis-ci.com/Sheenam3/x-tracer.svg?branch=master)](https://travis-ci.com/Sheenam3/x-tracer)

# x-tracer

In the era of kubernetes and containerization, there is a need to scale applications and understand its working inside a pod or a container. Kubernetes provision two metric pipelines to evaluate the applications performance and where bottlenecks can be erased for further enhancement:


1. Resource Metric Pipeline
According to the kubernetes documentation, a metric-server(need to deploy seprately), which discovers all the nodes and calculate its CPU and memory usage. Kubelet fetches individual container usage statistics in run time using ```kubectl top``` utility.

2. Full Metrics Pipeline
Like Prometheus, tool for event monitoring and alerting when crash occurs, by checking memory checks continuously. In short, it monitors linux/window servers, apache server, single application, and services with units like cpu status, memory usage, requests counts etc.


Here we are introducing a tool by ITRI named x-tracer , which traces every process log inside a pod and stream the logs to the x-tracer server in real time.


x-tracer includes 7 ebp tools(BCC),probes to trace the process events: 
1. Tcp connections: closed, active, established, life 
2. Block device I/O
3. New executed processes 
4. Cache kernel function calls


<h2>Basic Architecture/Flow of x-tracer:</h2>

![alt text](https://sheenampathak.com/wp-content/uploads/2020/06/Screenshot-from-2020-06-10-13-48-07.png)

<b>x-tracer flow:</b>
1. x-tracer server is deployed on the master node
2. x-agent client deploys on the worker node(in which our target pod is running)
3. x-agent creation triggers a go module named ```probeparser```, which executes 7 different probes(ebpf tools)
4. 7 probes traces the logs of the target_pod's processes using namespace ID(as every process PID in container belongs to the same namespace ID) 
5. These generated logs from the probes are channelized to the x-tracer server in real time


<b>The following logs are streamed at the x-tracer server: </b>
  
<pre>
Choose pod : 2
---------------------------------------------
The pod you chose is testpod
Container ID is ...
19fb910a711f5eabf2cad6569a01db3702752e2f8155059654130711b8bf2c8f
2020/04/01 12:45:25 Hostname :  dad
Start Agent Pod
Start Agent Service

{Probe:tcptracer |Sys_Time: 04:03:39 |T: 25.724 | PID:20656 | PNAME:iperf3 |IP->4 | SADDR:127.0.0.1 | DADDR:127.0.0.1 | SPORT:42334 | DPORT:6001 
{Probe:tcpconnect |Sys_Time: 04:03:40 |T: 28.857 | PID:20656 | PNAME:iperf3 | IP:4 | SADDR:127.0.0.1 | DADDR:127.0.0.1 | DPORT:6001 
{Probe:tcptracer |Sys_Time: 04:03:39 |T: 25.724 | PID:8592 | PNAME:iperf3 |IP->6 | SADDR:[::] | DADDR:[0:ffff:7f00:1::] | SPORT:0 | DPORT:65535 
{Probe:tcpaccept |Sys_Time: 04:03:40 |T: 28.863 | PID:8592 | PNAME:iperf3 | IP:6 | RADDR:::ffff:127.0.0.1 | RPORT:42336 | LADDR:::ffff:127.0.0.1 | LPORT:6001 
{Probe:tcptracer |Sys_Time: 04:03:40 |T: 25.767 | PID:20656 | PNAME:iperf3 |IP->4 | SADDR:127.0.0.1 | DADDR:127.0.0.1 | SPORT:42336 | DPORT:6001 
</pre>


