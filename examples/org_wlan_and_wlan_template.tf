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

### ORG WLAN TEMPLATE
resource "mist_org_wlantemplate" "test101" {
  org_id = mist_org.terraform_org.id
  name   = "test101"
  applies = {
    site_ids = [
      mist_site.terraform_site.id
    ]
  }
}

### ORG WLAN
resource "mist_org_wlan" "wlan_cwp" {
  ssid              = "MlN.test"
  bands             = ["5"]
  vlan_id           = 143
  wlan_limit_up     = 10000
  wlan_limit_down   = 20000
  client_limit_up   = 512
  client_limit_down = 1000
  portal = {
    enabled                = true
    bypass_when_cloud_down = true
    auth                   = "sso"
    privacy                = false
    sso_issuer             = "https://sts.windows.net/f2532c2f-938c-4529-b6e4-aa26992b6b62/"
    sso_nameid_format      = "email"
    sso_idp_sign_algo      = "sha256"
    sso_idp_cert           = "-----BEGIN CERTIFICATE-----\nMIIC8DCCAdigAwIBAgIQE5pOI9W1DZFHbB9m2Q7ADzANBgkqhkiG9w0BAQsFADA0MTIwMAYDVQQD\nEylNaWNyb3NvZnQgQXp1cmUgRmVkZXJhdGVkIFNTTyBDZXJ0aWZpY2F0ZTAeFw0yMjAyMDIxNDEz\nMTNaFw0yNTAyMDIxNDEzMTNaMDQxMjAwBgNVBAMTKU1pY3Jvc29mdCBBenVyZSBGZWRlcmF0ZWQg\nU1NPIENlcnRpZmljYXRlMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5gQTCccB3oE7\nelNYH2+11Q69Iq/2f3qf5KUZEQKwL++HyoBCOAM3wL3uLWwvRaih4+qpAeZvNsuShNIyB08SDcWN\nYsqVxaUsLYfzDD0c9VG9mwAx0Kh01S2JvtaLCaFveac7UXVfn/E/QbPXibS1EQvHUj0hwNXMrdS4\nh4TOk4D1Q70+OnCWyy7ykG1/RuO8UerIfqkQEy4C3QFb3Cyo4E7bEaYQo0NiCqD9IoM3B0wZib8Y\n3yRGJKdzXyDxuVJFb5rF7XMAHTWWAbxaN4KOLhZnjaJla7Pu/sFAj2Npm8Hm5pYEYBaUz4fc/8kg\nIwakFb3mnbnYw0xQwf+aJss1vQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQCF+oKuLmnooDzALwaE\nbFVI7PVGhU7/UZzAnq6HHI9ngF0Af2+uIrvAz6rdUM1bsGhRbj3SV2oaj26pe/1TDrGescXWhTPw\nKcXOwBnVmFr8FlMkozwpHRNzCQyFYGiTAztgQcmtwF7pilVndOmEc+p3LvCdI5JZB+LtMM/o9V+1\n+Yhm4MEWO6wTSY+j7goc/vi5f76TDZPN6PkRv17+EkybEudJuTOuIoNiqsAbNB52bVNHtxFHGIwb\nH9iS45QJ4/RG1WUr91xe3Vzh/fp1BkiHZVL4iOywOIF0TYcW7h958JEf+q0HD5LUMO47NPEbc/Cd\n+fVCTXWzABXXy4D+S8gA\n-----END CERTIFICATE-----"
    sso_idp_sso_url        = "https://login.microsoftonline.com/f2532c2f-938c-4529-b6e4-aa26992b6b62/saml2"
    email_enabled          = true
  }
  portal_allowed_hostnames = [
    "login.microsoftonline.com",
    "portal.mist.com",
    "login.live.com",
    "aadcdn.msauth.net",
    "logincdn.msauth.net"
  ]
  auth = {
    type = "psk"
    psk  = "Juniper123"
  }
  apply_to    = "site"
  org_id      = mist_org.terraform_test.id
  template_id = mist_org_wlantemplate.test101.id
  interface   = "all"
}

