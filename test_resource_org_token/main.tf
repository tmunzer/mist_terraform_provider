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
  //proxy = "test_proxy:Jun1per!Mist@10.3.18.11:3128"
   //username = local.envs["USERNAME"]
  // password = local.envs["PASSWORD"]
}

### ORG

resource "mist_org_servicepolicy" "test1org" {
  org_id = "992bf4b9-c900-4850-9992-107b2f9df928"
  tenants = [
    "PRD-Core"
  ]
  services = [
    "app_mqtt"
  ]
  action = "allow"
  idp = {
    enabled = false
  }
  name = "Policy-1"
}