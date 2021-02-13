#!/usr/bin/env bash

if [[ $(pwd) =~ "person-svc-go/scripts" ]]; then 
	cd ..
fi

# person-svc-go-dbv1 Service{}
mg create service \
 --output \
 --dryrun \
 --namespace person \
 -o yaml \
 person-svc-go-db.metagraf.json \
 | tee deployments/mg/person-svc-go-db.service.yaml

# person-svc-go-dbv1 Deployment{}
# --disable-aliasing is used because default behaviour is to expect a retag image with the
# application name and not the upstream name on your internal registry. This should be the
# other way around as default.
mg create deployment \
 --output \
 --dryrun \
 --namespace person \
 --disable-aliasing \
 -o yaml \
 --cvfile configs/mg/person-svc-go-db.properties \
 person-svc-go-db.metagraf.json \
 | tee deployments/mg/person-svc-go-db.deployment.yaml 

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