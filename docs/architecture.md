
<h2>Basic Architecture/Flow of x-tracer:</h2>

![alt text](https://sheenampathak.com/wp-content/uploads/2020/06/Screenshot-from-2020-06-10-13-48-07.png)

<b>x-tracer flow:</b>
1. x-tracer server is deployed on the master node
2. x-agent client deploys on the worker node(in which our target pod is running)
3. x-agent creation triggers a go module named ```probeparser```, which executes 7 different probes(ebpf tools)
4. 7 probes traces the logs of the target_pod's processes using namespace ID(as every process PID in container belongs to the same namespace ID) 
5. These generated logs from the probes are channelized to the x-tracer server in real time

