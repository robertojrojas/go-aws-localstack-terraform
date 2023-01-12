terraform {
  required_version = ">=0.13"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

provider "aws" {

  access_key                  = "mock_access_key"
  secret_key                  = "mock_secret_key"
  region                      = "us-east-1"

  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true

  endpoints {
    dynamodb             = "http://dynamodb.localhost.localstack.cloud:8566"
  }
}

resource "aws_dynamodb_table" "conformance_test_table" {
  name           = "simple_crud"
  billing_mode   = "PROVISIONED"
  read_capacity  = "10"
  write_capacity = "10"
  
  attribute {
    name = "Title"
    type = "S"
  } 
  hash_key = "Title"

  attribute {
    name = "Year"
    type = "N"
  }
  range_key = "Year"
}

