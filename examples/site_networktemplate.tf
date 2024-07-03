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

### SITE NETWORK CONFIGURATION
resource "mist_site_networktemplate" "site_switch_template" {
  site_id                = mist_site.terraform_site.id
  dns_servers            = ["10.3.51.222"]
  dns_suffix             = ["stag.one"]
  ntp_servers            = ["10.3.51.222"]
  additional_config_cmds = ["set system hostnam test", "set system services ssh root-login allow"]
  networks = {
    test = {
      subnet  = "1.2.3.4"
      vlan_id = 10
    }
    test2 = {
      subnet  = "1.2.3.4"
      vlan_id = 11

    }
  }
  radius_config = {
    acct_interim_interval = 60
    coa_enabled           = true
    network               = "test"
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
  port_usages = {
    trunk = {
      all_networks  = true
      description   = "profile for trunk ports"
      enable_qos    = true
      mode          = "trunk"
      port_networks = "test2"
    }
  }
  remote_syslog = {
    archive = {
      files = 2
      size  = "5m"
    }
    console = {
      contents = [
        {
          facility = "kernel"
          severity = "alert"
        }
      ]
    }
    enabled             = true
    network             = "test"
    send_to_all_servers = false
    servers = [
      {
        contents = [
          {
            facility = "any"
            severity = "any"
          }
        ]
        host = "1.2.3.4"
      }
    ]
    structured_data = true
  }
  switch_mgmt = {
    config_revert = 5
    protect_re = {
      enabled = true
    }
    root_password = "Juniper123"
  }
  switch_matching = {
    enable = true
    rules = [
      {
        match_type  = "match_name[0:3]"
        match_value = "abc"
        additional_config_cmds = [
          "set system name-server 8.8.8.8"
        ]
        match_role = "access"
        name       = "access"
        port_config = {
          "ge-0/0/0-10" = {
            usage = "trunk"
          }
        }
      },
      {
        additional_config_cmds = [
          "set system name-server 8.8.8.8"
        ]
        match_role = "core"
        name       = "core"
        port_config = {
          "ge-0/0/0-10" = {
            usage = "trunk"
          }
        }
      }
    ]
  }
}
