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

### SITES
resource "mist_site" "terraform_site" {
  org_id       = mist_org.terraform_org.id
  name         = "terraform_site"
}

### SITE SETTINGS
resource "mist_site_setting" "test" {
  site_id = mist_site.terraform_site.id
  analytic = {
    enabled = true
  }
  ap_updown_threshold     = 5
  device_updown_threshold = 5
  auto_upgrade = {
    enabled     = true
    day_of_week = "tue"
    time_of_day = "02:00"
    version     = "beta"
  }
  config_auto_revert = true

  persist_config_on_device = true
  proxy = {
    url = "http://myproxy:3128"
  }
  rogue = {
    enabled          = true
    honeypot_enabled = true
    min_duration     = 5
  }
}