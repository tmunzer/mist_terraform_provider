terraform {
  required_providers {
    mist = {
      source = "registry.terraform.io/juniper/mist"
    }
  }
}

provider "mist" {
  host     = local.envs["HOST"]
  apitoken = local.envs["APITOKEN"]
}

### ORG
resource "mist_org" "terraform_org" {
  name = "Terraform Testing"
}

resource "mist_site" "terraform_site" {
  org_id       = mist_org.terraform_org.id
  name         = "terraform_site2"
}

resource "mist_org_inventory" "inventory" {
  org_id       = mist_org.terraform_org.id
  devices = [{
    magic = "DEVICE_CLAIM_CODE"
    site_id = mist_site.terraform_site.id
    }
  ]
}