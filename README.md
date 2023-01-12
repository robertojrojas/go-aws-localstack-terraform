# go-aws-localstack-terraform
This repo contains AWS SDK Go v1/v2 using Localstack and provisioning with Terraform.

# Requirements
1- Go
2- Docker


# Run
```
cd scripts
./run_localstack.sh
./create_dynamodb_simple_crud.sh

cd cd ../aws-sdk/go/dynamodb/simple_crud 
go build .
AWS_ENDPOINT_URL=http://dynamodb.localhost.localstack.cloud:8566 ./dynamodb_simple_crud \
               -d simple_crud -y 2022 -t "NOPE" -r 1.0  -p 'hungry alien'

```

# Clean up
```
cd scripts
./clean.sh
```
