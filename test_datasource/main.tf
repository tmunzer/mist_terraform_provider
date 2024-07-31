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
data "mist_device_ap_stats" "example" {
  org_id  = "992bf4b9-c900-4850-9992-107b2f9df928"
}

data "mist_device_switch_stats" "example2" {
  org_id  = "992bf4b9-c900-4850-9992-107b2f9df928"
}
data "mist_device_gateway_stats" "example3" {
  org_id  = "992bf4b9-c900-4850-9992-107b2f9df928"
}
data "mist_org_deviceprofiles_ap" "dpa"{
  org_id  = "992bf4b9-c900-4850-9992-107b2f9df928"
}
data "mist_org_deviceprofiles_gateway" "dpg"{
  org_id  = "992bf4b9-c900-4850-9992-107b2f9df928"
}
data "mist_org_vpns" "dpg"{
  org_id  = "992bf4b9-c900-4850-9992-107b2f9df928"
}
data "mist_org_sitegroups" "test" {
   org_id  = "992bf4b9-c900-4850-9992-107b2f9df928"
}
data "mist_sites" "test" {
   org_id  = "992bf4b9-c900-4850-9992-107b2f9df928"
}
data "mist_org_servicepolicies" "servicepolicies" {
   org_id  = "992bf4b9-c900-4850-9992-107b2f9df928"
}
data "mist_org_idpprofiles" "idpprofiles" {
   org_id  = "992bf4b9-c900-4850-9992-107b2f9df928"
}

data "mist_const_countries" "test" {}
data "mist_const_app_categories" "test" {}
data "mist_const_app_sub_categories" "test" {}
//data "mist_const_gateway_applications" "test" {}
data "mist_const_traffic_types" "test" {}
data "mist_const_applications" "test" {}

# output "mist_device_ap_stats" {
#   value = data.mist_device_ap_stats.example
# }