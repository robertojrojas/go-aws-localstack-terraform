#!/bin/bash

cd ../terraform/dynamodb
terraform init 
terraform validate
terraform apply --auto-approve

cd -