// Provider configuration
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

// Create the Mist Organization
resource "mist_org" "org_one" {
  name = "My Provider Organization"
}

// Create the Site Group
resource "mist_org_sitegroup" "sitegroup_one" {
  org_id = mist_org.org_one.id
  name   = "My Provider Site Group"
}

// Create the WLAN Template and the WLAN
resource "mist_org_wlantemplate" "wlantempalte_one" {
  name   = "My Provider WLAN Template"
  org_id = mist_org.org_one.id
  applies = {
    sitegroup_ids = [
      mist_org_sitegroup.sitegroup_one.id
    ]
  }
}

resource "mist_org_wlan" "wlan_one" {
  ssid        = "My Provider WLAN"
  org_id      = mist_org.org_one.id
  template_id = mist_org_wlantemplate.wlantempalte_one.id
  auth = {
    type = "psk"
    psk  = "secretpsk"
  }
}

resource "mist_org_networktemplate" "networktemplate_one" {
  name   = "MyProviderNetworkTemplate"
  org_id = mist_org.org_one.id
  networks = {
    guest_vlan = {
      vlan_id = 10
    }
    user_vlan = {
      vlan_id = 11
    }
  }
  port_usages = {
    guest = {
      mode         = "access"
      port_network = "guest_vlan"
    }
    user = {
      mode      = "access"
      port_auth = "dot1x"
    }
    trunk = {
      all_networks = true
      enable_qos   = true
      mode         = "trunk"
      port_network = "guest_vlan"
    }
  }
  radius_config = {
    acct_servers = [
      {
        host   = "1.2.3.4"
        secret = "secret"
      }
    ]
    auth_servers = [
      {
        host   = "1.2.3.4"
        secret = "secret"
      }
    ]
  }
  switch_matching = {
    enable = true
    rules = [
      {
        name       = "switch_rule_one"
        match_name = "abc"
        port_config = {
          "ge-0/0/0-10" = {
            usage = "guest"
          }
        }
      }
    ]
  }
}

resource "mist_site" "site_one" {
  org_id       = mist_org.org_one.id
  name         = "Site One"
  address      = "41 rue de Villiers, 92100 Neully sur Seine, France"
  country_code = "FR"
  sitegroup_ids = [
    mist_org_sitegroup.sitegroup_one.id
  ]
  networktemplate_id = mist_org_networktemplate.networktemplate_one.id
}
