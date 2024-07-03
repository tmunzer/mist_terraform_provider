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

### ORG NAC TAGS
resource "mist_org_nactag" "vlan_sta" {
  vlan   = "10"
  name   = "vlan.sta"
  type   = "vlan"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "ssid_mln_do1x" {
  values = [
    "MlN.1X"
  ]
  name   = "ssid.MlN.1X"
  type   = "match"
  match  = "ssid"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "ssid_mln_nac" {
  values = [
    "MlN.NAC"
  ]
  name   = "ssid.MlN.NAC"
  type   = "match"
  match  = "ssid"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "ssid_mln" {
  values = [
    "MlN"
  ]
  name   = "ssid.MlN"
  type   = "match"
  match  = "ssid"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "mac_printer" {
  values = [
    "deadbeefdead"
  ]
  name   = "mac.printer"
  type   = "match"
  match  = "client_mac"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "vlan_admin" {
  vlan   = "20"
  name   = "vlan.admins"
  type   = "vlan"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "vlan_guest" {
  vlan   = "12"
  name   = "vlan.guest"
  type   = "vlan"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "ssor_corp_dot1x" {
  values = [
    "Corp"
  ]
  name   = "ssid.Corp.1X"
  type   = "match"
  match  = "ssid"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "role_admin" {
  radius_group = "admin"
  name         = "role.admin"
  type         = "radius_group"
  org_id       = mist_org.terraform_test.id
}
resource "mist_org_nactag" "role_user" {
  radius_group = "user"
  name         = "role.user"
  type         = "radius_group"
  org_id       = mist_org.terraform_test.id
}
resource "mist_org_nactag" "role_byod" {
  radius_group = "byod"
  name         = "role.byod"
  type         = "radius_group"
  org_id       = mist_org.terraform_test.id
}
resource "mist_org_nactag" "vlan_srv" {
  vlan   = "20"
  name   = "vlan.srv"
  type   = "vlan"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "cert_stag" {
  values = [
    "/DC=one/DC=stag/CN=stag-DC-CA"
  ]
  name   = "cert.Stag CA"
  type   = "match"
  match  = "cert_issuer"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "role_corp" {
  radius_group = "corporate"
  name         = "role.corporate"
  type         = "radius_group"
  org_id       = mist_org.terraform_test.id
}
resource "mist_org_nactag" "group_dot11" {
  values = [
    "dot11"
  ]
  name   = "group.dot11"
  type   = "match"
  match  = "idp_role"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "cert_mist" {
  values = [
    "/C=US/O=Mist/OU=OrgCA/CN=203d3d02-dbc0-4c1b-9f41-76896a3330f4"
  ]
  name   = "cert.Mist"
  type   = "match"
  match  = "cert_issuer"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "role_ap" {
  radius_group = "wireless"
  name         = "role.AP"
  type         = "radius_group"
  org_id       = mist_org.terraform_test.id
}
resource "mist_org_nactag" "session_7d" {
  session_timeout = 604800
  name            = "session.7d"
  type            = "session_timeout"
  org_id          = mist_org.terraform_test.id
}
resource "mist_org_nactag" "vman_mgmt" {
  vlan   = "172"
  name   = "vlan.mgmt"
  type   = "vlan"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "session_1d" {
  session_timeout = 86440
  name            = "session.1d"
  type            = "session_timeout"
  org_id          = mist_org.terraform_test.id
}
resource "mist_org_nactag" "mac_pc" {
  values = [
    "04d9f5c8b26e",
    "04d9f5c8b26d"
  ]
  name   = "mac.pc"
  type   = "match"
  match  = "client_mac"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "group_admins" {
  values = [
    "administrators"
  ]
  name   = "group.admins"
  type   = "match"
  match  = "idp_role"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "vlan_reg" {
  vlan   = "12"
  name   = "vlan.reg"
  type   = "vlan"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "cert_mist_eu" {
  values = [
    "/C=US/O=Mist/OU=OrgCA/CN=39ce2088-1dbe-4346-987a-1a5a88bab5ee"
  ]
  name   = "cert.Mist-EU"
  type   = "match"
  match  = "cert_issuer"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "trunk_ap" {
  egress_vlan_names = [
    "2mgt",
    "1corp",
    "1iot",
    "1lab",
    "1lan",
    "1srv",
    "1wan",
    "1reg"
  ]
  name   = "trunk-ap"
  type   = "egress_vlan_names"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "mac_srv" {
  values = [
    "dca6321e2f86",
    "e45f018e7054",
    "b8cb29d96468",
    "f4e2c6553b32",
    "b827ebfd31a4"
  ]
  name   = "mac.srv"
  type   = "match"
  match  = "client_mac"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "mdm_compliant" {
  values = [
    "compliant"
  ]
  name   = "mdm.compliant"
  type   = "match"
  match  = "mdm_status"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "mdm_non_compliant" {
  values = [
    "non-compliant",
    "unknown"
  ]
  name   = "mdm.non_compliant"
  type   = "match"
  match  = "mdm_status"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "mac_ap" {
  values = [
    "5c5b35*"
  ]
  name   = "mac.ap"
  type   = "match"
  match  = "client_mac"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_nactag" "vlan_dmz" {
  vlan   = "192"
  name   = "vlan.dmz"
  type   = "vlan"
  org_id = mist_org.terraform_test.id
}

### ORG NAC RULES
resource "mist_org_nacrule" "wired_man" {
  matching = {
    port_types = [
      "wired"
    ]
    auth_type = "mab"
    nactags = [
      mist_org_nactag.mac_srv.id
    ]
  }
  apply_tags = [
    mist_org_nactag.vlan_srv.id,
    mist_org_nactag.session_1d.id
  ]
  action  = "allow"
  enabled = true
  name    = "Wired MAB - SRV"
  org_id  = mist_org.terraform_test.id
  order   = 9
}
resource "mist_org_nacrule" "default_wired" {
  matching = {
    port_types = [
      "wired"
    ]
  }
  apply_tags = [
    mist_org_nactag.vlan_reg.id
  ]
  action  = "allow"
  enabled = true
  name    = "Default-Wired"
  org_id  = mist_org.terraform_test.id
  order   = 11
}
resource "mist_org_nacrule" "wired_mac_ap_stag" {
  matching = {
    auth_type = "mab"
    nactags = [
      mist_org_nactag.mac_ap.id
    ]
    port_types = [
      "wired"
    ]
  }
  apply_tags = []
  action     = "allow"
  enabled    = true
  name       = "Wired MAB - AP Staging"
  org_id     = mist_org.terraform_test.id
  order      = 7
}
resource "mist_org_nacrule" "wifi_tls_minc_sta" {
  matching = {
    auth_type = "cert"
    nactags = [
      mist_org_nactag.cert_stag.id,
      mist_org_nactag.group_dot11.id,
      mist_org_nactag.ssid_mln_nac.id,
      mist_org_nactag.mdm_non_compliant.id
    ]
    port_types = [
      "wireless"
    ]
  }
  apply_tags = [
    mist_org_nactag.vlan_sta.id,
    mist_org_nactag.role_byod.id,
    mist_org_nactag.session_1d.id
  ]
  action  = "allow"
  enabled = true
  name    = "Wi-Fi TLS MINC - STA"
  org_id  = mist_org.terraform_test.id
  order   = 4
}
resource "mist_org_nacrule" "wifi_tls_sta" {
  matching = {
    auth_type = "cert"
    nactags = [
      mist_org_nactag.cert_stag.id,
      mist_org_nactag.ssid_mln_nac.id
    ]
    port_types = [
      "wireless"
    ]
  }
  apply_tags = [
    mist_org_nactag.vlan_sta.id
  ]
  action  = "allow"
  enabled = true
  name    = "Wi-Fi TLS - STA"
  org_id  = mist_org.terraform_test.id
  order   = 5
}
resource "mist_org_nacrule" "wired_tls_ap" {
  matching = {
    port_types = [
      "wired"
    ]
    nactags = [
      mist_org_nactag.cert_mist.id,
      "cert",
      mist_org_nactag.cert_mist_eu.id
    ]
    auth_type = "cert"
  }
  apply_tags = [
    mist_org_nactag.trunk_ap.id,
    mist_org_nactag.session_7d.id
  ]
  action  = "allow"
  enabled = true
  name    = "Wired TLS - AP Prod"
  org_id  = mist_org.terraform_test.id
  order   = 6
}
resource "mist_org_nacrule" "test" {
  matching = {
    port_types = [
      "wired"
    ]
    nactags = [
      mist_org_nactag.ssid_mln_nac.id
    ]
    auth_type = "eap-teap"
  }
  action  = "allow"
  enabled = true
  name    = "test"
  org_id  = mist_org.terraform_test.id
  order   = 0
}
resource "mist_org_nacrule" "wifi_teap_sta" {
  matching = {
    port_types = [
      "wired"
    ]
    auth_type = "cert"
  }
  apply_tags = [
    mist_org_nactag.session_1d.id,
    mist_org_nactag.vlan_srv.id
  ]
  action  = "allow"
  enabled = true
  name    = "Wired TLS - Corp"
  org_id  = mist_org.terraform_test.id
  order   = 8
}
resource "mist_org_nacrule" "wifi_teap_corp" {
  matching = {
    auth_type = "eap-teap"
    nactags = [
      mist_org_nactag.ssid_mln_nac.id
    ]
    port_types = [
      "wireless"
    ]
  }
  apply_tags = [
    mist_org_nactag.vlan_sta.id,
    mist_org_nactag.role_byod.id,
    mist_org_nactag.session_1d.id
  ]
  action  = "allow"
  enabled = true
  name    = "Wi-Fi TEAP - STA"
  org_id  = mist_org.terraform_test.id
  order   = 1
}
resource "mist_org_nacrule" "wifi_tsl_mic_sta" {
  matching = {
    nactags = [
      mist_org_nactag.cert_stag.id,
      mist_org_nactag.group_dot11.id,
      mist_org_nactag.ssid_mln_nac.id,
      mist_org_nactag.mdm_compliant.id
    ]
    auth_type = "eap-tls"
    port_types = [
      "wireless"
    ]
  }
  apply_tags = [
    mist_org_nactag.vlan_sta.id,
    mist_org_nactag.role_byod.id,
    mist_org_nactag.session_1d.id
  ]
  action  = "allow"
  enabled = true
  name    = "Wi-Fi TLS MIC - STA"
  org_id  = mist_org.terraform_test.id
  order   = 3
}
resource "mist_org_nacrule" "wired_mac_sta" {
  matching = {
    auth_type = "mab"
    port_types = [
      "wired"
    ]
    nactags = [
      mist_org_nactag.mac_pc.id
    ]
  }
  apply_tags = [
    mist_org_nactag.vlan_sta.id
  ]
  action  = "allow"
  enabled = true
  name    = "Wired MAB - STA"
  org_id  = mist_org.terraform_test.id
  order   = 10
}