# person-svc-go

This is a really simple and basic implementation of a go lang based
microservice. It's main purpose is to demonstrate how to containerize
it, how to work with this locally and how to run this on that Kubernetes
thing.



## Building

We'll experiment with different ways of building this very simple service.

We'll try the following tools:

- buildah
- podman-compose
- GitHub Actions
- Openshift BuildConfig (s2i)

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
- kompose or podman 3.0
- metagraf / mg

### On my local machine

#### podman

Podman command to run the Postgres instance our person-svc-go service needs:

```bash
podman run --env-file configs/envfiles/postgres.env docker.io/library/postgres
# or with --detach to get your shell back
podman run --detach --env-file configs/envfiles/postgres.env docker.io/library/postgres
```

Podman command to run the person-svc-go container we built:

```bash
# Run the locally built image
podman run <imageid>
# or
podman run localhost/person-svc-go
```

#### podman-compose

```bash
# or
podman-compose up
# Need to issue restart since the go application starts before the database is ready.
# A deployment to Kubernetes would take care of this on it's own.
podman-compose restart person-svc-go
```

### For Kubernetes

#### kompose

We can use a tool called kompose to transform a docker-compose.yaml to 
Kubernetes manifests.

This is a quick way to get some basic Kubernetes manifests going.

```bash
#todo
kompose ....
```

#### metaGraf - Generating manifests with mg

We'll use the mg tool from the https://github.com/laetho/metagraf project
and the person-svc-go-db.metagraf.json we wrote. We're using the *mg* specific
config files in configs/mg/ as input for some of the commands.

To generate an empty .properties file from a metaGraf specification you
can do the following:

```bash
> mg create properties person-svc-go-db.metagraf.json > configs/mg/person-svc-go-db.properties 

> mg create properties person-svc-go.metagraf.json > configs/mg/person-svc-go.properties

```

The *mg* tool can also validate your *.properties* file against the metaGraf specification:

```bash
> mg inspect properties person-svc-go.metagraf.json configs/mg/person-svc-go.properties 
The configs/mg/person-svc-go.properties configuration is valid for this metaGraf specification.
```

Generate kubernetes manifests for the ephemeral PosgreSQL instance.

Snippet from: [scripts/mg-generate-manifests.sh](https://github.com/laetho/person-svc-go/blob/master/scripts/mg-generate-manifests.sh):
```bash
...
# Service{}
mg create service \
 --output \
 --dryrun \
 --namespace person \
 -o yaml \
 person-svc-go-db.metagraf.json \
 | tee deployments/mg/person-svc-go-db.service.yaml

# Deployment{}
mg create deployment \
 --output \
 --dryrun \
 --namespace person \
 --disable-aliasing \
 -o yaml \
 --cvfile configs/mg/person-go-svc-db.properties \
 person-svc-go-db.metagraf.json \
 | tee deployments/mg/person-svc-go-db.deployment.yaml
 ...
```
Produced:
- [deployments/mg/person-svc-go-db.deployment.yaml](https://github.com/laetho/person-svc-go/blob/master/deployments/mg/person-svc-go-db.deployment.yaml)
- [deployments/mg/person-svc-go-db.service.yaml](https://github.com/laetho/person-svc-go/blob/master/deployments/mg/person-svc-go-db.service.yaml)

Generate Kubernetes manifests for the person-svc-go service:

Snippet from: [scripts/mg-generate-manifests.sh](https://github.com/laetho/person-svc-go/blob/master/scripts/mg-generate-manifests.sh):
```bash
# person-svc-gov1 Service{}
mg create service \
 --output \
 --dryrun \
 --namespace person \
 -o yaml \
 person-svc-go.metagraf.json \
 | tee deployments/mg/person-svc-go.service.yaml

# person-svc-gov1 Deployment{}
# --disable-aliasing is used because default behaviour is to expect a retag image with the
# application name and not the upstream name on your internal registry. This should be the
# other way around as default.
mg create deployment \
 --output \
 --dryrun \
 --namespace person \
 --disable-aliasing \
 -o yaml \
 --cvfile configs/mg/person-svc-go.properties \
 person-svc-go.metagraf.json \
 | tee deployments/mg/person-svc-go.deployment.yaml
```
Produced:
- [deployments/mg/person-svc-go.deployment.yaml](https://github.com/laetho/person-svc-go/blob/master/deployments/mg/person-svc-go.deployment.yaml)
- [deployments/mg/person-svc-go.service.yaml](https://github.com/laetho/person-svc-go/blob/master/deployments/mg/person-svc-go.service.yaml)


## Performance testing

My old and trusty workstation from yesteryears:

```bash
...
model name      : Intel(R) Xeon(R) CPU           W3530  @ 2.80GHz
...
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
