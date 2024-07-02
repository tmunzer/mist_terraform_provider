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
  rftemplate_id = mist_rftemplate.test_rf.id
  gatewaytemplate_id = mist_gatewaytemplate.stag.id
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


resource "mist_gatewaytemplate" "stag" {

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
  type   = "spoke"
  name   = "test-api"
  org_id = mist_org.terraform_test.id
}

resource "mist_gatewaytemplate" "demo" {
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


resource "mist_nactag" "vlan_sta" {
  vlan   = "10"
  name   = "vlan.sta"
  type   = "vlan"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "ssid_mln_do1x" {
  values = [
    "MlN.1X"
  ]
  name   = "ssid.MlN.1X"
  type   = "match"
  match  = "ssid"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "ssid_mln_nac" {
  values = [
    "MlN.NAC"
  ]
  name   = "ssid.MlN.NAC"
  type   = "match"
  match  = "ssid"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "ssid_mln" {
  values = [
    "MlN"
  ]
  name   = "ssid.MlN"
  type   = "match"
  match  = "ssid"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "mac_printer" {
  values = [
    "deadbeefdead"
  ]
  name   = "mac.printer"
  type   = "match"
  match  = "client_mac"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "vlan_admin" {
  vlan   = "20"
  name   = "vlan.admins"
  type   = "vlan"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "vlan_guest" {
  vlan   = "12"
  name   = "vlan.guest"
  type   = "vlan"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "ssor_corp_dot1x" {
  values = [
    "Corp"
  ]
  name   = "ssid.Corp.1X"
  type   = "match"
  match  = "ssid"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "role_admin" {
  radius_group = "admin"
  name         = "role.admin"
  type         = "radius_group"
  org_id       = mist_org.terraform_test.id
}
resource "mist_nactag" "role_user" {
  radius_group = "user"
  name         = "role.user"
  type         = "radius_group"
  org_id       = mist_org.terraform_test.id
}
resource "mist_nactag" "role_byod" {
  radius_group = "byod"
  name         = "role.byod"
  type         = "radius_group"
  org_id       = mist_org.terraform_test.id
}
resource "mist_nactag" "vlan_srv" {
  vlan   = "20"
  name   = "vlan.srv"
  type   = "vlan"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "cert_stag" {
  values = [
    "/DC=one/DC=stag/CN=stag-DC-CA"
  ]
  name   = "cert.Stag CA"
  type   = "match"
  match  = "cert_issuer"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "cert_juniper" {
  values = [
    "/DC=net/DC=jnpr/CN=Juniper Networks Issuing AWS1 CA"
  ]
  name   = "cert.Juniper CA"
  type   = "match"
  match  = "cert_issuer"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "cert_tm" {
  values = [
    "tmunzer@juniper.net"
  ]
  name   = "cert.tmunzer@jun"
  type   = "match"
  match  = "cert_san"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "role_corp" {
  radius_group = "corporate"
  name         = "role.corporate"
  type         = "radius_group"
  org_id       = mist_org.terraform_test.id
}
resource "mist_nactag" "group_dot11" {
  values = [
    "dot11"
  ]
  name   = "group.dot11"
  type   = "match"
  match  = "idp_role"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "cert_mist" {
  values = [
    "/C=US/O=Mist/OU=OrgCA/CN=203d3d02-dbc0-4c1b-9f41-76896a3330f4"
  ]
  name   = "cert.Mist"
  type   = "match"
  match  = "cert_issuer"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "role_ap" {
  radius_group = "wireless"
  name         = "role.AP"
  type         = "radius_group"
  org_id       = mist_org.terraform_test.id
}
resource "mist_nactag" "session_7d" {
  session_timeout = 604800
  name            = "session.7d"
  type            = "session_timeout"
  org_id          = mist_org.terraform_test.id
}
resource "mist_nactag" "vman_mgmt" {
  vlan   = "172"
  name   = "vlan.mgmt"
  type   = "vlan"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "session_1d" {
  session_timeout = 86440
  name            = "session.1d"
  type            = "session_timeout"
  org_id          = mist_org.terraform_test.id
}
resource "mist_nactag" "mac_pc" {
  values = [
    "04d9f5c8b26e",
    "04d9f5c8b26d"
  ]
  name   = "mac.pc"
  type   = "match"
  match  = "client_mac"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "group_admins" {
  values = [
    "administrators"
  ]
  name   = "group.admins"
  type   = "match"
  match  = "idp_role"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "vlan_reg" {
  vlan   = "12"
  name   = "vlan.reg"
  type   = "vlan"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "cert_mist_eu" {
  values = [
    "/C=US/O=Mist/OU=OrgCA/CN=39ce2088-1dbe-4346-987a-1a5a88bab5ee"
  ]
  name   = "cert.Mist-EU"
  type   = "match"
  match  = "cert_issuer"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "trunk_ap" {
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
resource "mist_nactag" "mac_srv" {
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
resource "mist_nactag" "mdm_compliant" {
  values = [
    "compliant"
  ]
  name   = "mdm.compliant"
  type   = "match"
  match  = "mdm_status"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "mdm_non_compliant" {
  values = [
    "non-compliant",
    "unknown"
  ]
  name   = "mdm.non_compliant"
  type   = "match"
  match  = "mdm_status"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "mac_ap" {
  values = [
    "5c5b35*"
  ]
  name   = "mac.ap"
  type   = "match"
  match  = "client_mac"
  org_id = mist_org.terraform_test.id
}
resource "mist_nactag" "vlan_dmz" {
  vlan   = "192"
  name   = "vlan.dmz"
  type   = "vlan"
  org_id = mist_org.terraform_test.id
}

resource "mist_nacrule" "wired_man" {
  matching = {
    port_types = [
      "wired"
    ]
    auth_type = "mab"
    nactags = [
      mist_nactag.mac_srv.id
    ]
  }
  apply_tags = [
    mist_nactag.vlan_srv.id,
    mist_nactag.session_1d.id
  ]
  action  = "allow"
  enabled = true
  name    = "Wired MAB - SRV"
  org_id  = mist_org.terraform_test.id
  order   = 9
}
resource "mist_nacrule" "default_wired" {
  matching = {
    port_types = [
      "wired"
    ]
  }
  apply_tags = [
    mist_nactag.vlan_reg.id
  ]
  action  = "allow"
  enabled = true
  name    = "Default-Wired"
  org_id  = mist_org.terraform_test.id
  order   = 11
}
resource "mist_nacrule" "wired_mac_ap_stag" {
  matching = {
    auth_type = "mab"
    nactags = [
      mist_nactag.mac_ap.id
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
resource "mist_nacrule" "wifi_tls_minc_sta" {
  matching = {
    auth_type = "cert"
    nactags = [
      mist_nactag.cert_stag.id,
      mist_nactag.group_dot11.id,
      mist_nactag.ssid_mln_nac.id,
      mist_nactag.mdm_non_compliant.id
    ]
    port_types = [
      "wireless"
    ]
  }
  apply_tags = [
    mist_nactag.vlan_sta.id,
    mist_nactag.role_byod.id,
    mist_nactag.session_1d.id
  ]
  action  = "allow"
  enabled = true
  name    = "Wi-Fi TLS MINC - STA"
  org_id  = mist_org.terraform_test.id
  order   = 4
}
resource "mist_nacrule" "wifi_tls_sta" {
  matching = {
    auth_type = "cert"
    nactags = [
      mist_nactag.cert_stag.id,
      mist_nactag.ssid_mln_nac.id
    ]
    port_types = [
      "wireless"
    ]
  }
  apply_tags = [
    mist_nactag.vlan_sta.id
  ]
  action  = "allow"
  enabled = true
  name    = "Wi-Fi TLS - STA"
  org_id  = mist_org.terraform_test.id
  order   = 5
}
resource "mist_nacrule" "wired_tls_ap" {
  matching = {
    port_types = [
      "wired"
    ]
    nactags = [
      mist_nactag.cert_mist.id,
      "cert",
      mist_nactag.cert_mist_eu.id
    ]
    auth_type = "cert"
  }
  apply_tags = [
    mist_nactag.trunk_ap.id,
    mist_nactag.session_7d.id
  ]
  action  = "allow"
  enabled = true
  name    = "Wired TLS - AP Prod"
  org_id  = mist_org.terraform_test.id
  order   = 6
}
resource "mist_nacrule" "test" {
  matching = {
    port_types = [
      "wired"
    ]
    nactags = [
      mist_nactag.ssid_mln_nac.id
    ]
    auth_type = "eap-teap"
  }
  action     = "allow"
  enabled    = true
  name       = "test"
  org_id     = mist_org.terraform_test.id
  order      = 0
}
resource "mist_nacrule" "wifi_teap_sta" {
  matching = {
    port_types = [
      "wired"
    ]
    auth_type = "cert"
  }
  apply_tags = [
    mist_nactag.session_1d.id,
    mist_nactag.vlan_srv.id
  ]
  action  = "allow"
  enabled = true
  name    = "Wired TLS - Corp"
  org_id  = mist_org.terraform_test.id
  order   = 8
}
resource "mist_nacrule" "wifi_teap_corp" {
  matching = {
    auth_type = "eap-teap"
    nactags = [
      mist_nactag.ssid_mln_nac.id
    ]
    port_types = [
      "wireless"
    ]
  }
  apply_tags = [
    mist_nactag.vlan_sta.id,
    mist_nactag.role_byod.id,
    mist_nactag.session_1d.id
  ]
  action  = "allow"
  enabled = true
  name    = "Wi-Fi TEAP - STA"
  org_id  = mist_org.terraform_test.id
  order   = 1
}
resource "mist_nacrule" "wifi_tls_corp" {
  matching = {
    auth_type = "cert"
    port_types = [
      "wireless"
    ]
    nactags = [
      mist_nactag.ssid_mln_nac.id,
      mist_nactag.cert_juniper.id,
      mist_nactag.cert_tm.id
    ]
  }
  apply_tags = [
    mist_nactag.vlan_sta.id,
    mist_nactag.role_corp.id,
    mist_nactag.session_1d.id
  ]
  action  = "allow"
  enabled = true
  name    = "Wi-Fi TLS - Corporate"
  org_id  = mist_org.terraform_test.id
  order   = 2
}
resource "mist_nacrule" "wifi_tsl_mic_sta" {
  matching = {
    nactags = [
      mist_nactag.cert_stag.id,
      mist_nactag.group_dot11.id,
      mist_nactag.ssid_mln_nac.id,
      mist_nactag.mdm_compliant.id
    ]
    auth_type = "eap-tls"
    port_types = [
      "wireless"
    ]
  }
  apply_tags = [
    mist_nactag.vlan_sta.id,
    mist_nactag.role_byod.id,
    mist_nactag.session_1d.id
  ]
  action  = "allow"
  enabled = true
  name    = "Wi-Fi TLS MIC - STA"
  org_id  = mist_org.terraform_test.id
  order   = 3
}
resource "mist_nacrule" "wired_mac_sta" {
  matching = {
    auth_type = "mab"
    port_types = [
      "wired"
    ]
    nactags = [
      mist_nactag.mac_pc.id
    ]
  }
  apply_tags = [
    mist_nactag.vlan_sta.id
  ]
  action  = "allow"
  enabled = true
  name    = "Wired MAB - STA"
  org_id  = mist_org.terraform_test.id
  order   = 10
}

resource "mist_rftemplate" "test_rf" {
    band_24_usage = "auto"
    band_5 = {
        ant_gain = 2
        power = 8
        channels = [
            60,
            104,
            132
        ]
        bandwidth = 20
    }
    band_6 = {
        ant_gain = 2
        power = 8
    }
    band_24 = {
        ant_gain = 1
        allow_rrm_disable = true
        power_min = 18
        power_max = 18
        bandwidth = 20
    }
    ant_gain_5 = 2
    ant_gain_6 = 2
    ant_gain_24 = 1
    country_code = "FR"
    name = "tf"
    org_id = mist_org.terraform_test.id
}

resource "mist_site_setting" "test" {
  site_id = mist_site.terraform_site.id
  analytic = {
    enabled = true
  }
  ap_updown_threshold = 5
  device_updown_threshold = 5
  auto_upgrade= {
    enabled = true
    day_of_week = "tue"
    time_of_day = "02:00"
    version = "beta"
  }
  config_auto_revert = true
  dns_servers = [
    "8.8.8.8",
    "1.1.1.1"
  ]
  persist_config_on_device = true
  proxy = {
    url ="http://myproxy:3128"
  }
  rogue = {
    enabled = true
    honeypot_enabled = true
    min_duration = 5
  }
}

resource "mist_wlantemplate" "test101" {
  org_id = mist_org.terraform_test.id
  name = "test101"
  applies = {
    site_ids= [
      mist_site.terraform_site.id
    ]
  }
}
resource "mist_wlan" "wlan_cwp"     {
        ssid= "MlN.test"
        bands= ["5"]
        vlan_id= 143
        wlan_limit_up= 10000
        wlan_limit_down= 20000
        client_limit_up= 512
        client_limit_down= 1000
        portal= {
            enabled= true
            bypass_when_cloud_down= true
            auth= "sso"
            privacy= false
            sso_issuer= "https://sts.windows.net/f2532c2f-938c-4529-b6e4-aa26992b6b62/"
            sso_nameid_format= "email"
            sso_idp_sign_algo= "sha256"
            sso_idp_cert= "-----BEGIN CERTIFICATE-----\nMIIC8DCCAdigAwIBAgIQE5pOI9W1DZFHbB9m2Q7ADzANBgkqhkiG9w0BAQsFADA0MTIwMAYDVQQD\nEylNaWNyb3NvZnQgQXp1cmUgRmVkZXJhdGVkIFNTTyBDZXJ0aWZpY2F0ZTAeFw0yMjAyMDIxNDEz\nMTNaFw0yNTAyMDIxNDEzMTNaMDQxMjAwBgNVBAMTKU1pY3Jvc29mdCBBenVyZSBGZWRlcmF0ZWQg\nU1NPIENlcnRpZmljYXRlMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5gQTCccB3oE7\nelNYH2+11Q69Iq/2f3qf5KUZEQKwL++HyoBCOAM3wL3uLWwvRaih4+qpAeZvNsuShNIyB08SDcWN\nYsqVxaUsLYfzDD0c9VG9mwAx0Kh01S2JvtaLCaFveac7UXVfn/E/QbPXibS1EQvHUj0hwNXMrdS4\nh4TOk4D1Q70+OnCWyy7ykG1/RuO8UerIfqkQEy4C3QFb3Cyo4E7bEaYQo0NiCqD9IoM3B0wZib8Y\n3yRGJKdzXyDxuVJFb5rF7XMAHTWWAbxaN4KOLhZnjaJla7Pu/sFAj2Npm8Hm5pYEYBaUz4fc/8kg\nIwakFb3mnbnYw0xQwf+aJss1vQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQCF+oKuLmnooDzALwaE\nbFVI7PVGhU7/UZzAnq6HHI9ngF0Af2+uIrvAz6rdUM1bsGhRbj3SV2oaj26pe/1TDrGescXWhTPw\nKcXOwBnVmFr8FlMkozwpHRNzCQyFYGiTAztgQcmtwF7pilVndOmEc+p3LvCdI5JZB+LtMM/o9V+1\n+Yhm4MEWO6wTSY+j7goc/vi5f76TDZPN6PkRv17+EkybEudJuTOuIoNiqsAbNB52bVNHtxFHGIwb\nH9iS45QJ4/RG1WUr91xe3Vzh/fp1BkiHZVL4iOywOIF0TYcW7h958JEf+q0HD5LUMO47NPEbc/Cd\n+fVCTXWzABXXy4D+S8gA\n-----END CERTIFICATE-----"
            sso_idp_sso_url= "https://login.microsoftonline.com/f2532c2f-938c-4529-b6e4-aa26992b6b62/saml2"
            email_enabled= true
        }
        portal_allowed_hostnames= [
            "login.microsoftonline.com",
            "portal.mist.com",
            "login.live.com",
            "aadcdn.msauth.net",
            "logincdn.msauth.net"
        ]
        auth= {
            type= "psk"
            psk= "Juniper123"
        }
        apply_to= "site"
        template_id= mist_wlantemplate.test101.id
        interface= "all"
    }