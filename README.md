[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) 

[![Build Status](https://travis-ci.com/ITRI-ICL-Peregrine/x-tracer.svg?branch=master)](https://travis-ci.com/ITRI-ICL-Peregrine/x-tracer)

# x-tracer
x-tracer provides the process-level visibility inside the Kubernetes pod in real time. As containers are isolated, lightweight which makes it difficult to look inside what 
could go wrong because they totally run on a separate namespace and PIDs inside containers are different as from their hosts which doesn't provide us with a clear view of 
what processes are running inside a particular pod or container. To make it simple, x-tracer is beautifully integrated with BCC tools to provide such kind of visibility.

Currently, x-tracer includes 7 ebp tools(BCC),probes to trace the process events:
1. Tcp connections: closed, active, established, life
2. Block device I/O 
3. New executed processes
4. Cache kernel function calls
  
## Architecture

[Architecture](docs/Architecture.md)

## Installation Steps

[Installation](docs/x-tracer-installation.md)



