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
  org_id        = mist_org.terraform_org.id
  name          = "terraform_site"
  rftemplate_id = mist_org_rftemplate.test_rf.id
}

### ORG RF TEMPLATE
resource "mist_org_rftemplate" "test_rf" {
  band_24_usage = "auto"
  band_5 = {
    ant_gain = 2
    power    = 8
    channels = [
      60,
      104,
      132
    ]
    bandwidth = 20
  }
  band_6 = {
    ant_gain = 2
    power    = 8
  }
  band_24 = {
    ant_gain          = 1
    allow_rrm_disable = true
    power_min         = 18
    power_max         = 18
    bandwidth         = 20
  }
  ant_gain_5   = 2
  ant_gain_6   = 2
  ant_gain_24  = 1
  country_code = "FR"
  name         = "tf"
  org_id       = mist_org.terraform_test.id
}
