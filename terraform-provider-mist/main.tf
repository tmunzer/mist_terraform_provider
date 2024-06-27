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

resource "mist_org" "terraform_test" {
  name = "Terraform Testing"
}

resource "mist_site" "terraform_site" {
  org_id       = mist_org.terraform_test.id
  name         = "terraform_site"
  country_code = "FR"
  timezone     = "Europe/Paris"
  address      = "77 Terrasse de l'Universit\u00e9, 92000 Nanterre, France"
  notes        = "Created with Terraform, Updated with Terraform"
  latlng = {
    lat = 48.899268
    lng = 2.214447
  }
  sitegroup_ids      = [mist_sitegroup.test_group.id, mist_sitegroup.test_group2.id]
  networktemplate_id = mist_networktemplate.switch_template.id
}


resource "mist_sitegroup" "test_group" {
  org_id = mist_org.terraform_test.id
  name   = "test group"
}
resource "mist_sitegroup" "test_group2" {
  org_id = mist_org.terraform_test.id
  name   = "test group2b"
}

resource "mist_service" "lab" {
  org_id = mist_org.terraform_test.id
  addresses = [
    "10.3.0.0/24", "10.4.0.0/24"
  ]
  description = "the lab network"
  name        = "lab_network"
  type        = "custom"
  specs = [
    {
      protocol   = "tcp"
      port_range = "443"
    }
  ]
}


resource "mist_network" "corp" {
  org_id                 = mist_org.terraform_test.id
  disallow_mist_services = false
  name                   = "prd_corp"
  subnet                 = "10.3.0.0"
}

resource "mist_networktemplate" "switch_template" {
  name                   = "test switch"
  org_id                 = mist_org.terraform_test.id
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


resource "mist_network" "mgmt" {
  org_id  = mist_org.terraform_test.id
  vlan_id = 172
  subnet  = "10.3.172.0/24"
  gateway = "10.3.172.9"
  vpn_access = {
    "OrgOverlay" : {
      routed                    = true
      no_readvertise_to_overlay = false
      no_readvertise_to_lan_bgp = false
    }
  }
  isolation              = true
  disallow_mist_services = false
  name                   = "SRX-Mgmt"
}
resource "mist_network" "reg" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 12
  subnet                 = "10.3.12.0/24"
  disallow_mist_services = false
  name                   = "SRX-REG"
}
resource "mist_network" "ssr" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 128
  subnet                 = "10.128.100.0/16"
  disallow_mist_services = false
  name                   = "SRX-Core-128T"
}
resource "mist_network" "mxe" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 1000
  subnet                 = "10.10.0.0/24"
  disallow_mist_services = false
  tenants = {
    "mxe-ext" = {
      addresses = [
        "10.10.0.10"
      ]
    }
  }
  internet_access = {
    static_nat = {}
    destination_nat = {
      "192.168.1.9:500" : {
        name        = "ike"
        internal_ip = "10.10.0.10"
        port        = "500"
      }
      "192.168.1.9:4500" : {
        name        = "isakmp"
        internal_ip = "10.10.0.10"
        port        = "4500"
      }
    }
  }
  name = "SRX-MXE-tt-in"
}
resource "mist_network" "iot" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 8
  subnet                 = "10.3.8.0/24"
  disallow_mist_services = true
  tenants = {
    netatmo = {
      addresses = [
        "10.3.8.201/32"
      ]
    }
    "cctv-e" = {
      addresses = [
        "10.3.8.40/32"
      ]
    }
    "cctv-b" = {
      addresses = [
        "10.3.8.41/32"
      ]
    }
    "cctv-c" = {
      addresses = [
        "10.3.8.42/32"
      ]
    }
    "cctv-x" = {
      addresses = [
        "10.3.8.44/32"
      ]
    }
  }
  name = "SRX-IoT"
}
resource "mist_network" "dmz" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3
  subnet                 = "10.3.3.0/24"
  disallow_mist_services = true
  tenants = {
    aws = {
      addresses = [
        "10.3.3.3"
      ]
    }
  }
  internet_access = {
    static_nat = {}
    destination_nat = {
      "192.168.1.9:5431" : {
        name        = "aws-wg"
        internal_ip = "10.3.3.3"
        port        = "51820"
      }
    }
  }
  name = "SRX-DMZ"
}
