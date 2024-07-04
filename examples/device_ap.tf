provider "mist" {
  host     = local.envs["HOST"]
  apitoken = local.envs["APITOKEN"]
}

### ORG
resource "mist_org" "terraform_test" {
  name = "Terraform Testing"
}
resource "mist_org_inventory" "inventory" {
  org_id       = mist_org.terraform_test.id
  devices = [
    {
      magic = "DEVICE_CLAIM_CODE"
      site_id = mist_site.terraform_site.id

    }
  ]
}
### SITES
resource "mist_site" "terraform_site" {
  org_id       = mist_org.terraform_org.id
  name         = "terraform_site"
}
resource "mist_device_ap" "test_ap" {
  id = mist_org_inventory.inventory.devices[0].device_id
  site_id  = mist_org_inventory.inventory.devices[0].site_id
  name = "test"
}