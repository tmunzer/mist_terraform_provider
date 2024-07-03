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
  sitegroup_ids = [mist_org_sitegroup.test_group.id, mist_org_sitegroup.test_group2.id]
}


resource "mist_site" "terraform_site2" {
  org_id       = mist_org.terraform_org.id
  name         = "terraform_site2"
  sitegroup_ids = [mist_org_sitegroup.test_group.id, mist_org_sitegroup.test_group2.id]
}

resource "mist_org_sitegroup" "test_group" {
  org_id = mist_org.terraform_org.id
  name   = "test group"
}
resource "mist_org_sitegroup" "test_group2" {
  org_id = mist_org.terraform_org.id
  name   = "test group2b"
}