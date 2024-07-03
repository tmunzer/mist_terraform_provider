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

### SITE
resource "mist_site" "terraform_site" {
  org_id       = mist_org.terraform_org.id
  name         = "terraform_site2"
}

### INVENTORY
### NOTES:
# Adding the device (magic is required) to the list will claim it to the org
# Adding the site_id to a device will assign it to the site
# Changing the site_id from a device will reassign it to the new site
# Removing the site_id from a device will unassign it from the site
# Removing the device from the list will unclaim is from the Org
resource "mist_org_inventory" "inventory" {
  org_id       = mist_org.terraform_org.id
  devices = [
    {
    magic = "DEVICE_CLAIM_CODE_1"
    site_id = mist_site.terraform_site.id
    },
    {
    magic = "DEVICE_CLAIM_CODE_2"
    },
  ]
}