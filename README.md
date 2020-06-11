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

