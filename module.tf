provider "aws" {
  region = "eu-central-1"
}

module "etcd" {
  source    = "modules/ebs"
  instances = 3
  zone      = "eu-central-1a"
}


