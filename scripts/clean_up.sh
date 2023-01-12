#!/bin/bash

cd ../terraform/dynamodb
terraform destroy --auto-approve
rm -rf .terraform* terraform*

cd ../../docker
rm -rf docker.sock volume

cd -
docker-compose -f ../docker/localstack.yml down
