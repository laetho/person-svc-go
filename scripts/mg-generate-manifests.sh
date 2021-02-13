#!/usr/bin/env bash

if [[ $(pwd) =~ "person-svc-go/scripts" ]]; then 
	cd ..
fi

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


