terraform {
  required_providers {
    mistapi = {
      source = "registry.terraform.io/juniper/mistapi"
    }
  }
}

provider "mistapi" {
  host     = local.envs["HOST"]
  apitoken = local.envs["APITOKEN"]
}

resource "mistapi_org" "terraform_test" {
  name       = "Terraform Testing"
}

