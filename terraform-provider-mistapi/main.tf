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
  name = "Terraform Testing"
}

resource "mistapi_site" "terraform_site" {
  org_id       = mistapi_org.terraform_test.id
  name         = "terraform_site"
  country_code = "FR"
  timezone     = "Europe/Paris"
  address      = "77 Terrasse de l'Universit\u00e9, 92000 Nanterre, France"
  notes        = "Created with Terraform, Updated with Terraform"
  latlng = {
    lat = 48.899268
    lng = 2.214447
  }
  sitegroup_ids      = [mistapi_sitegroup.test_group.id, mistapi_sitegroup.test_group2.id]
  networktemplate_id = mistapi_networktemplate.switch_template.id
}


resource "mistapi_sitegroup" "test_group" {
  org_id = mistapi_org.terraform_test.id
  name   = "test group"
}
resource "mistapi_sitegroup" "test_group2" {
  org_id = mistapi_org.terraform_test.id
  name   = "test group2b"
}

resource "mistapi_networktemplate" "switch_template" {
  name                   = "test switch"
  org_id                 = mistapi_org.terraform_test.id
  dns_servers            = ["10.3.51.222"]
  dns_suffix             = ["stag.one"]
  ntp_servers            = ["10.3.51.222"]
  additional_config_cmds = ["set system hostnam test", "set system services ssh root-login allow"]
  networks = {
    test = {
      subnet  = "1.2.3.4"
      vlan_id = 10
    },
    test2 = {
      subnet  = "1.2.3.4"
      vlan_id = 11

    }
  }
  radius_config = {
    acct_interim_interval = 60
    coa_enabled           = true
    network               = "test"
  }
}
