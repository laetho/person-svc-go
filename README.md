# person-svc-go


## Building

We'll experiment with different ways of building this rather service.

We'll try the following tools:

- buildah
- podman-compose
- docker
- docker-compose
- GitHub Actions

### buildah
Experimenting with buildah and podman instead of running a docker 
daemon. Testing podman-compose as well. 

```bash
buildah bud -f Dockerfile -t person-svc-go
```

### podman-compose

```bash
podman-compose build
```


## Running

In this section we'll explore different ways of running the container we 
built and other supporting containers we need to test or run this application. 

We'll go into the following tools for running it or generating manifests to run
it on kuberntes using diffrent tools like:

- podman
- podman-compose
- docker
- docker-compose
- kompose
- metagraf / mg

### On my local machine

```bash
# Run just the built image
podman run <imageid>
# or
podman-compose up
# Need to issue restart since the go application starts before the database is ready.
# A deployment to Kubernetes would take care of this on it's own.
podman-compose restart person-svc-go
```

### On kubernetes

## Performance testing

My old and trusty workstation from yesteryears:

```bash
processor       : 0
vendor_id       : GenuineIntel
cpu family      : 6
model           : 26
model name      : Intel(R) Xeon(R) CPU           W3530  @ 2.80GHz
stepping        : 5
microcode       : 0x1d
cpu MHz         : 1644.764
cache size      : 8192 KB
physical id     : 0
siblings        : 8
core id         : 0
cpu cores       : 4
apicid          : 0
initial apicid  : 0
fpu             : yes
fpu_exception   : yes
cpuid level     : 11
wp              : yes
flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ht tm pbe syscall nx rdtscp lm constant_tsc arch_perfmon pebs bts rep_good nopl xtopology nonstop_tsc cpuid aperfmperf pni dtes64 monitor ds_cpl vmx est tm2 ssse3 cx16 xtpr pdcm dca sse4_1 sse4_2 popcnt lahf_lm pti ssbd ibrs ibpb stibp tpr_shadow vnmi flexpriority ept vpid dtherm ida flush_l1d
vmx flags       : vnmi preemption_timer invvpid ept_x_only flexpriority tsc_offset vtpr mtf vapic ept vpid
bugs            : cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit
bogomips        : 5602.25
clflush size    : 64
cache_alignment : 64
address sizes   : 36 bits physical, 48 bits virtual
power management:
```

Run that connects to Postgres and queries for persons and returns the results.

```bash
wrk -c 100 -d 60 -t 8 http://localhost:8080/persons
Running 1m test @ http://localhost:8080/persons
  8 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.48ms    3.90ms  45.51ms   63.03%
    Req/Sec     3.02k   157.93     4.21k    68.25%
  1442493 requests in 1.00m, 537.89MB read
Requests/sec:  24031.57
Transfer/sec:      8.96MB
```


Run that calls a status enpoint with a static JSON response.

```bash
wrk -c 100 -d 60 -t 8 http://localhost:8080/status
Running 1m test @ http://localhost:8080/status
  8 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.17ms    3.11ms  57.00ms   87.35%
    Req/Sec     9.80k     2.23k   23.82k    68.21%
  4682922 requests in 1.00m, 553.78MB read
Requests/sec:  77947.51
Transfer/sec:      9.22MB
```
