terraform {
  backend "s3" {
    bucket = "rickenbacker-terraform"
    key    = "terraform.tfstate"
    region = "ap-northeast-1"
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.27"
    }
  }
  required_version = ">= 0.14.9"
}
