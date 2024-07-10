terraform {
  required_providers {
    mist = {
      source = "registry.terraform.io/juniper/mist"
    }
  }
}

provider "mist" {
  host     = local.envs["HOST"]
  //apitoken = local.envs["APITOKEN"]
  username = local.envs["USERNAME"]
  password = local.envs["PASSWORD"]
}

### ORG
resource "mist_org" "terraform_test" {
  name = "Terraform Testing"
}
resource "mist_org_inventory" "inventory" {
  org_id = mist_org.terraform_test.id
  devices = [
    {
      claim_code = "CPKL2EXN8JY98AC"
      site_id    = mist_site.terraform_site.id
    },
    {
      claim_code = "G87JHBFXZJSFNMX"
      site_id    = mist_site.terraform_site.id
    },
    {
      claim_code = "CV4YAS8DQWYLL6M"
      site_id    = mist_site.terraform_site.id
    }
  ]
}
### SITES
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
  sitegroup_ids      = [mist_org_sitegroup.test_group.id, mist_org_sitegroup.test_group2.id]
  networktemplate_id = mist_org_networktemplate.switch_template.id
  rftemplate_id      = mist_org_rftemplate.test_rf.id
  gatewaytemplate_id = mist_org_gatewaytemplate.test-api.id
}

resource "mist_device_ap" "test_ap" {
  device_id = mist_org_inventory.inventory.devices[0].id
  site_id   = mist_org_inventory.inventory.devices[0].site_id
  name      = "test_ap"
}


resource "mist_site" "terraform_site2" {
  org_id       = mist_org.terraform_test.id
  name         = "terraform_site2"
  country_code = "FR"
  timezone     = "Europe/Paris"
  address      = "77 Terrasse de l'Universit\u00e9, 92000 Nanterre, France"
  notes        = "Created with Terraform, Updated with Terraform"
  latlng = {
    lat = 48.899268
    lng = 2.214447
  }
  sitegroup_ids = [mist_org_sitegroup.test_group.id, mist_org_sitegroup.test_group2.id]
}

resource "mist_org_sitegroup" "test_group" {
  org_id = mist_org.terraform_test.id
  name   = "test group"
}
resource "mist_org_sitegroup" "test_group2" {
  org_id = mist_org.terraform_test.id
  name   = "test group2b"
}

# # ### ORG LEVEL
resource "mist_org_service" "lab" {
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

resource "mist_org_network" "corp" {
  org_id                 = mist_org.terraform_test.id
  disallow_mist_services = false
  name                   = "prd_corp"
  subnet                 = "10.3.0.0"
}

resource "mist_org_network" "mgmt" {
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
resource "mist_org_network" "reg" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 12
  subnet                 = "10.3.12.0/24"
  disallow_mist_services = false
  name                   = "SRX-REG"
}
resource "mist_org_network" "ssr" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 128
  subnet                 = "10.128.100.0/16"
  disallow_mist_services = false
  name                   = "SRX-Core-128T"
}
resource "mist_org_network" "mxe" {
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
resource "mist_org_network" "iot" {
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
resource "mist_org_network" "dmz" {
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


resource "mist_org_gatewaytemplate" "test-api" {

  type   = "spoke"
  name   = "test-api"
  org_id = mist_org.terraform_test.id

  port_config = {
    "ge-0/0/3" = {
      name       = "FTTH"
      usage      = "wan"
      aggregated = false
      redundant  = false
      critical   = false
      wan_type   = "broadband"
      ip_config = {
        type    = "static"
        ip      = "192.168.1.8"
        netmask = "/24"
        gateway = "192.168.1.1"
      },
      disable_autoneg = false
      speed           = "auto"
      duplex          = "auto"
      wan_source_nat = {
        disabled = false
      },
      vpn_paths = {
        "SSR_HUB_DC-MPLS.OrgOverlay" = {
          key         = 0
          role        = "spoke"
          bfd_profile = "broadband"
        }
      }
    },
    "ge-0/0/4" = {
      name       = "LTE"
      usage      = "wan"
      aggregated = false
      redundant  = false
      critical   = false
      wan_type   = "broadband"
      ip_config = {
        type = "dhcp"
      },
      disable_autoneg = false
      speed           = "auto"
      duplex          = "auto"
      wan_source_nat = {
        disabled = false
      },
      vpn_paths = {
        "SSR_HUB_DC-MPLS.OrgOverlay" = {
          key         = 0
          role        = "spoke"
          bfd_profile = "broadband"
        }
      }
    },
    "ge-0/0/5" = {
      usage            = "lan"
      critical         = false
      aggregated       = true
      ae_disable_lacp  = false
      ae_lacp_force_up = true
      ae_idx           = 0
      redundant        = false
      networks = [
        "PRD-Core"
      ]
    },
    "ge-0/0/7" = {
      usage      = "lan"
      critical   = false
      aggregated = false
      redundant  = false
      networks = [
        "PRD-Mgmt"
      ]
    },
    "ge-0/0/6" = {
      usage      = "lan"
      critical   = false
      aggregated = false
      redundant  = false
      networks = [
        "PRD-Lab"
      ]
    }
  }
  ip_configs = {
    "PRD-Core" = {
      type    = "static"
      ip      = "10.3.100.9"
      netmask = "/24"
    },
    "PRD-Mgmt" = {
      type    = "static"
      ip      = "10.3.172.1"
      netmask = "/24"
    },
    "PRD-Lab" = {
      type    = "static"
      ip      = "10.3.171.1"
      netmask = "/24"
    }
  }
  dhcpd_config = {
    enable = true
  }
  path_preferences = {
    "HUB" = {
      strategy = "ordered"
      paths = [
        {
          name = "SSR_HUB_DC-MPLS.OrgOverlay"
          type = "vpn"
        }
      ]
    },
    "HUB-ORDERED" = {
      strategy = "ordered"
      paths = [
        {
          name     = "SSR_HUB_DC-MPLS.OrgOverlay"
          wan_name = "FTTH",
          type     = "vpn"
        },
        {
          name     = "SSR_HUB_DC-MPLS.OrgOverlay"
          wan_name = "LTE"
          type     = "vpn"
        }
      ]
    },
    "HUB-ECMP" = {
      strategy = "weighted"
      paths = [
        {
          name     = "SSR_HUB_DC-MPLS.OrgOverlay"
          wan_name = "LTE"
          cost     = 30
          type     = "vpn"
        },
        {
          name     = "SSR_HUB_DC-MPLS.OrgOverlay"
          wan_name = "FTTH"
          cost     = 30
          type     = "vpn"
        }
      ]
    }
  }
  service_policies = [
    {
      name = "Policy-14"
      tenants = [
        "PRD-Core"
      ],
      services = [
        "any"
      ],
      action          = "allow"
      path_preference = "HUB"
      idp = {
        enabled    = true
        profile    = "critical"
        alert_only = false
      }
    },
    {
      name = "Policy-2"
      tenants = [
        "PRD-Mgmt"
      ],
      services = [
        "any"
      ],
      action          = "allow",
      path_preference = "HUB-ECMP"
      idp = {
        enabled    = true,
        profile    = "standard"
        alert_only = true
      }
    },
    {
      name = "Policy-3"
      tenants = [
        "PRD-Lab"
      ],
      services = [
        "any"
      ],
      action          = "allow"
      path_preference = "HUB-ORDERED"
      idp = {
        enabled = false
      }
    }
  ]
}

resource "mist_org_gatewaytemplate" "sdwan-westford" {
  ntp_override = true
  service_policies = [
    {
      name = "inbound-hub-to-corp-net"
      idp = {
        enabled = false
      },
      path_preference = "to-corp-lan"
      services = [
        "spoke-corp-network"
      ],
      action = "allow"
      tenants = [
        "corp_datacenter1",
        "corp_datacenter2",
        "core_dc2_network",
        "mgmt_datacenter1",
        "mgmt_datacenter2"
      ]
    },
    {
      name = "inbound-remote-spoke-to-corp-net"
      tenants = [
        "corp_network",
        "mgmt_network"
      ],
      services = [
        "spoke-corp-network"
      ],
      action          = "allow"
      path_preference = "to-corp-lan"
      idp = {
        enabled = false
      }
    },
    {
      name = "outbound-to-hub-workloads"
      tenants = [
        "corp_network"
      ],
      services = [
        "datacenter1-mgmt-network",
        "datacenter1-workloads",
        "datacenter2-lan",
        "datacenter2-mgmt-network",
        "datacenter2-workloads"
      ],
      action          = "allow"
      path_preference = "broadband-overlay"
      idp = {
        enabled = false
      }
    },
    {
      name = "outbound-corp-internet"
      tenants = [
        "corp_network"
      ],
      services = [
        "corp-internet"
      ],
      action          = "allow",
      path_preference = "internet-breakout"
      idp = {
        enabled = false
      }
    },
    {
      name = "application-test"
      tenants = [
        "corp_network"
      ],
      services = [
        "youtube"
      ],
      action = "allow"
      idp = {
        enabled = false
      },
      path_preference = "internet-breakout"
    },
    {
      name = "url-filtering"
      tenants = [
        "corp_network"
      ],
      services = [
        "blocked-categories"
      ],
      action        = "deny"
      local_routing = true
      idp = {
        enabled = false
      }
    }
  ]
  dns_servers = [
    "8.8.8.8"
  ]
  port_config = {
    "ge-0/0/1" : {
      name       = "wan1-broadband"
      usage      = "wan"
      aggregated = false
      redundant  = false
      critical   = false
      disabled   = false
      wan_type   = "broadband"
      ip_config = {
        type    = "static"
        ip      = "{{broadband_ip_var}}"
        netmask = "/{{broadband_prefix_var}}"
        gateway = "{{broadband_gw_var}}"
      },
      disable_autoneg = false
      wan_source_nat = {
        disabled = false
      },
      vpn_paths = {
        "sdwan_newyork_hub-external1.OrgOverlay" : {
          key         = 0
          role        = "spoke"
          bfd_profile = "broadband"
        }
      }
    },
    "ge-0/0/3" : {
      networks = [
        "corp_network"
      ],
      usage      = "lan"
      aggregated = false
      redundant  = false
      critical   = false
      disabled   = false
    }
  }
  path_preferences = {
    "broadband-overlay" : {
      strategy = "ordered"
      paths = [
        {
          name = "sdwan_newyork_hub-external1.OrgOverlay"
          type = "vpn"
        }
      ]
    },
    "internet-breakout" : {
      strategy = "ordered"
      paths = [
        {
          name = "wan1-broadband"
          type = "wan"
        }
      ]
    },
    "to-corp-lan" : {
      strategy = "ordered"
      paths = [
        {
          type = "local"
          networks = [
            "corp_network"
          ]
        }
      ]
    }
  }
  ntp_servers = [
    "pool.ntp.org"
  ]
  routing_policies = {
    "ibgp-hub-preference" : {
      terms = [
        {
          matching = {
            prefix = [
              "172.76.128.0/24",
              "172.36.128.0/24",
              "192.168.94.0/24",
              "192.168.93.0/24",
              "192.168.92.0/24",
              "192.168.95.0/24",
              "192.168.91.0/24",
              "172.46.128.0/24",
              "172.66.128.0/24",
              "172.56.128.0/24",
              "24.0.0.0/29",
              "192.168.90.0/24",
              "192.168.97.0/24",
              "192.168.98.0/24",
              "192.168.94.0/24"
            ],
            protocol = [
              "bgp"
            ],
            vpn_path = [
              "NewYork-newyork-broadband1.OrgOverlay",
              "NewYork-newyork-broadband1-lte.OrgOverlay",
              "SanFrancisco-sanfran-broadband1.OrgOverlay"
            ]
          },
          actions = {
            accept = true
          }
        }
      ]
    }
  }
  ip_configs = {
    corp_network = {
      type    = "static"
      ip      = "{{corp_ip}}"
      netmask = "/{{corp_CIDR}}"
    }
  }
  dns_override = true
  dhcpd_config = {
    enabled = true
    config = {
      corp_network = {
        type     = "local"
        ip_start = "{{corp_dhcp_start}}"
        ip_end   = "{{corp_dhcp_end}}"
        gateway  = "{{corp_ip}}"
        dns_servers = [
          "8.8.8.8",
          "1.1.1.1"
        ]
        options = {}
      }
    }
  }
  bgp_config = {
    "bgp-group1" : {
      via           = "vpn"
      vpn_name      = "OrgOverlay"
      import_policy = "ibgp-hub-preference"
    }
  }
  type   = "spoke"
  name   = "sdwan-westford"
  org_id = mist_org.terraform_test.id

}


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

resource "mist_site_setting" "test" {
  site_id = mist_site.terraform_site.id
  # analytic = {
  #   enabled = true
  # }
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

resource "mist_org_wlantemplate" "test101" {
  org_id = mist_org.terraform_test.id
  name   = "test101"
  applies = {
    site_ids = [
      mist_site.terraform_site.id
    ]
  }
}

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







# # ################### SITE LEVEL


resource "mist_org_networktemplate" "switch_template" {
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

# ################### SITE LEVEL


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
      all_networks = true
      description  = "profile for trunk ports"
      enable_qos   = true
      mode         = "trunk"
      port_network = "test2"
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

resource "mist_site_wlan" "wlan_cwp2" {
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
  apply_to  = "site"
  interface = "all"
  site_id   = mist_site.terraform_site2.id
}

resource "mist_device_switch" "test_switch" {
  managed = true
  role    = "test"
  networks = {
    "prx" = {
      vlan_id = "18"
    }
  }
  port_usages = {
    "prx" = {
      mode         = "trunk"
      disabled     = false
      port_network = "default"
      stp_edge     = false
      all_networks = false
      networks = [
        "default",
        "prx"
      ],
      speed         = "auto"
      duplex        = "auto"
      mac_limit     = 0
      persist_mac   = false
      poe_disabled  = false
      enable_qos    = false
      storm_control = {}
      description   = ""
    },
    "autogenerated_ge-0_0_10" = {
      mode         = "trunk"
      all_networks = false
      networks = [
        "default",
        "prx"
      ],
      port_network = "prx"
      autoneg      = true
      disabled     = false
      poe_disabled = false
      duplex       = "auto"
      speed        = "auto"
    },
    "autogenerated_ge-0_0_11" = {
      mode         = "trunk"
      all_networks = false
      networks = [
        "default",
        "prx"
      ],
      port_network = "prx"
      autoneg      = true
      disabled     = false
      poe_disabled = false
      duplex       = "auto"
      speed        = "auto"
    }
  }
  additional_config_cmds = [
    "set groups top system proxy password \"$9$VpsgajHmFnCq.pBIEeK4aZUDin6Ctu136eW\"",
    "",
    "annotate system \" -- custom-auth -- Template level --\"",
    "delete apply-groups custom-auth",
    "delete groups custom-auth",
    "set groups custom-auth",
    "set groups custom-auth  system login user tmunzer class super-user",
    "set groups custom-auth  system login user tmunzer authentication ssh-rsa \"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCxk1CXmrc/wNlzuelQoBgEdnA3Mq8r05UhXgBUwqFdcbwkEmgYnciQWXYeFdA+UwR3SWX4VtyeenQi/S9pSOjvoZQ5OYBWNqak+AWjXXzX7sO8cf59TSDDN2lybdH6kWbvRKYzMbfpZvwluYwQX7ftN/D4/iS7288i5knaMVklhDJGdcuh6Xlve6PueZ5ov2twPzpplPVdWOCqgaCMhvVplG3oNTvBpeJ6BCXzhgJdkFNinFehPnvUR6MRwdjD/MvS1r+yrggHGQzJ57VEDsw2qH8wRaUUn0+Nd5hMYO4vamfintlm5hEAW9+my6HKNzBmd/TMeLctKUOg/q4Bc9TfYB51HEghHZ1i3ClJ6nE/1M6Ev9B+YwBJpeRaieWHbOHWrmMKM5rnYpX+uShPobvHKC8Z5HU5UkAtWvRyXVqbeeVv7xlGDTKnm0UQtzv9UDR/BN3VSnh8JKGLeiPDRZ0oDdgqK1ijxXBCFKb6FXAUiORUQLtI/7dmV4YxK4oS+ZRU/uxCCNDx6R90o08RFDFpWNBPY1yj332T/6JnA7nzlAYdQ1s1hzgzamZoEZuO/a+NObV9gQbJC13ZtZEWKXSLHPvEsORc5nb7hN76RNI7yKnAWNy5Zj2aaPlTGofbdr9M8+Tez4wldmYB/k86M4lENv5fZRx9A6VmfXFDoSI33w== tmunzer@stag.one\"",
    "set apply-groups custom-auth"
  ]
  ip_config = {
    type    = "static"
    ip      = "10.3.18.99"
    netmask = "255.255.255.0"
    network = "prx"
    gateway = "10.3.18.11"
  }
  # routing_policies= {
  #     "tyes"= {
  #         terms= [
  #             {
  #                 matching= {
  #                     prefix= [
  #                         "0.0.0.0/0"
  #                     ],
  #                     protocol= [
  #                         "ospf"
  #                     ]
  #                 },
  #                 actions= {
  #                     accept= true
  #                 },
  #                 name= "gfds"
  #             }
  #         ]
  #     }
  # }
  port_config = {
    "ge-0/0/0" = {
      usage                = "prx"
      critical             = false
      description          = ""
      "no_local_overwrite" = true
    },
    "ge-0/0/10" = {
      usage              = "autogenerated_ge-0_0_10"
      no_local_overwrite = false
      speed              = "100m"
    },
    "ge-0/0/11" = {
      usage                = "default"
      port_network         = "prx"
      critical             = false
      description          = ""
      "no_local_overwrite" = false
    }
  }
  port_mirroring = {
    "test" = {
      output_port_id = "ge-0/0/10"
      input_port_ids_ingress = [
        "ge-0/0/2"
      ],
      input_port_ids_egress = [
        "ge-0/0/2"
      ],
      input_networks_ingress = [
        "default"
      ]
    },
    "test2" = {
      output_network         = "prx"
      input_port_ids_ingress = []
      input_port_ids_egress = [
        "ge-0/0/6"
      ],
      input_networks_ingress = []
    }
  }
  mist_nac = {
    enabled = true
  }
  vrf_instances= {
      "fds"= {
          networks= [
              "prx"
          ],
          extra_routes= {
              "1.2.0.0/24"= {
                  via= "1.2.3.4"
              }
          }
      }
  }
  vrf_config = {
    enabled = true
  }
  dhcpd_config = {
    enabled = true
    "prx" = {
      type = "relay"
      servers = [
        "1.2.3.4"
      ]
    }
  }
  # ospf_areas= {
  #     "0"= {
  #         type= "default"
  #         networks= {
  #             "prx"= {
  #                 passive= false
  #                 interface_type= "p2p"
  #                 hello_interval= 10
  #                 dead_interval= 40
  #                 auth_type= "password"
  #                 auth_password= "hpihpi"
  #                 metric= 1
  #                 bfd_minimum_interval= 1
  #             }
  #         },
  #         include_loopback= false
  #     }
  # }
  ospf_config = {
    enabled             = true
    reference_bandwidth = "10M"
    areas = {
      "0" = {
        no_summary = false
      }
    }
  }
  router_id = "1.2.3.4"
  oob_ip_config = {
    type                      = "static"
    ip                        = "2.2.2.2"
    netmask                   = "/24"
    use_mgmt_vrf              = true
    use_mgmt_vrf_for_host_out = true
    gateway                   = "2.2.2.1"
  }
  device_id = mist_org_inventory.inventory.devices[1].id
  name      = "demo-ex"
  site_id   = mist_org_inventory.inventory.devices[1].site_id
}


resource "mist_device_gateway" "srx" {

  name   = "srx"
  device_id= mist_org_inventory.inventory.devices[2].id
  site_id = mist_org_inventory.inventory.devices[2].site_id

  port_config = {
    "ge-0/0/3" = {
      name       = "FTTH"
      usage      = "wan"
      aggregated = false
      redundant  = false
      critical   = false
      wan_type   = "broadband"
      ip_config = {
        type    = "static"
        ip      = "192.168.1.8"
        netmask = "/24"
        gateway = "192.168.1.1"
      },
      disable_autoneg = false
      speed           = "auto"
      duplex          = "auto"
      wan_source_nat = {
        disabled = false
      },
      vpn_paths = {
        "SSR_HUB_DC-MPLS.OrgOverlay" = {
          key         = 0
          role        = "spoke"
          bfd_profile = "broadband"
        }
      }
    },
    "ge-0/0/4" = {
      name       = "LTE"
      usage      = "wan"
      aggregated = false
      redundant  = false
      critical   = false
      wan_type   = "broadband"
      ip_config = {
        type = "dhcp"
      },
      disable_autoneg = false
      speed           = "auto"
      duplex          = "auto"
      wan_source_nat = {
        disabled = false
      },
      vpn_paths = {
        "SSR_HUB_DC-MPLS.OrgOverlay" = {
          key         = 0
          role        = "spoke"
          bfd_profile = "broadband"
        }
      }
    },
    "ge-0/0/5" = {
      usage            = "lan"
      critical         = false
      aggregated       = true
      ae_disable_lacp  = false
      ae_lacp_force_up = true
      ae_idx           = 0
      redundant        = false
      networks = [
        "PRD-Core"
      ]
    },
    "ge-0/0/7" = {
      usage      = "lan"
      critical   = false
      aggregated = false
      redundant  = false
      networks = [
        "PRD-Mgmt"
      ]
    },
    "ge-0/0/6" = {
      usage      = "lan"
      critical   = false
      aggregated = false
      redundant  = false
      networks = [
        "PRD-Lab"
      ]
    }
  }
  ip_configs = {
    "PRD-Core" = {
      type    = "static"
      ip      = "10.3.100.9"
      netmask = "/24"
    },
    "PRD-Mgmt" = {
      type    = "static"
      ip      = "10.3.172.1"
      netmask = "/24"
    },
    "PRD-Lab" = {
      type    = "static"
      ip      = "10.3.171.1"
      netmask = "/24"
    }
  }
  dhcpd_config = {
    enable = true
  }
  path_preferences = {
    "HUB" = {
      strategy = "ordered"
      paths = [
        {
          name = "SSR_HUB_DC-MPLS.OrgOverlay"
          type = "vpn"
        }
      ]
    },
    "HUB-ORDERED" = {
      strategy = "ordered"
      paths = [
        {
          name     = "SSR_HUB_DC-MPLS.OrgOverlay"
          wan_name = "FTTH",
          type     = "vpn"
        },
        {
          name     = "SSR_HUB_DC-MPLS.OrgOverlay"
          wan_name = "LTE"
          type     = "vpn"
        }
      ]
    },
    "HUB-ECMP" = {
      strategy = "weighted"
      paths = [
        {
          name     = "SSR_HUB_DC-MPLS.OrgOverlay"
          wan_name = "LTE"
          cost     = 30
          type     = "vpn"
        },
        {
          name     = "SSR_HUB_DC-MPLS.OrgOverlay"
          wan_name = "FTTH"
          cost     = 30
          type     = "vpn"
        }
      ]
    }
  }
  service_policies = [
    {
      name = "Policy-14"
      tenants = [
        "PRD-Core"
      ],
      services = [
        "any"
      ],
      action          = "allow"
      path_preference = "HUB"
      idp = {
        enabled    = true
        profile    = "critical"
        alert_only = false
      }
    },
    {
      name = "Policy-2"
      tenants = [
        "PRD-Mgmt"
      ],
      services = [
        "any"
      ],
      action          = "allow",
      path_preference = "HUB-ECMP"
      idp = {
        enabled    = true,
        profile    = "standard"
        alert_only = true
      }
    },
    {
      name = "Policy-3"
      tenants = [
        "PRD-Lab"
      ],
      services = [
        "any"
      ],
      action          = "allow"
      path_preference = "HUB-ORDERED"
      idp = {
        enabled = false
      }
    }
  ]
}
