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
  # username = local.envs["USERNAME"]
  # password = local.envs["PASSWORD"]
}

### ORG
resource "mist_org" "terraform_test" {
  name = "Terraform Testing"
}

# data "mist_device_ap_stats" "example" {
#   #org_id = mist_org.terraform_test.id
#   org_id  = "9777c1a0-6ef6-11e6-8bbf-02e208b2d34f"
# }

# data "mist_device_switch_stats" "example2" {
#   #org_id = mist_org.terraform_test.id
#   org_id  = "9777c1a0-6ef6-11e6-8bbf-02e208b2d34f"
# }
# data "mist_device_gateway_stats" "example3" {
#   org_id  = "9777c1a0-6ef6-11e6-8bbf-02e208b2d34f"
# }
# output "mist_device_ap_stats" {
#   value = data.mist_device_ap_stats.example
# }
data "mist_org_deviceprofiles_ap" "dpa"{
  org_id  = "9777c1a0-6ef6-11e6-8bbf-02e208b2d34f"
}
data "mist_org_deviceprofiles_gateway" "dpg"{
  org_id  = "9777c1a0-6ef6-11e6-8bbf-02e208b2d34f"
}

# data "mist_countries" "test" {}
# output "mist_countries" {
#   value = 
# }

data "mist_org_sitegroups" "test" {
   org_id  = "9777c1a0-6ef6-11e6-8bbf-02e208b2d34f"
}