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
}
