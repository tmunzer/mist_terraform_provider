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

# resource "mist_site" "terraform_site" {
#   org_id       = mist_org.terraform_test.id
#   name         = "terraform_site"
#   country_code = "FR"
#   timezone     = "Europe/Paris"
#   address      = "77 Terrasse de l'Universit\u00e9, 92000 Nanterre, France"
#   notes        = "Created with Terraform, Updated with Terraform"
#   latlng = {
#     lat = 48.899268
#     lng = 2.214447
#   }
#   sitegroup_ids      = [mist_sitegroup.test_group.id, mist_sitegroup.test_group2.id]
#   networktemplate_id = mist_networktemplate.switch_template.id
# }


# resource "mist_sitegroup" "test_group" {
#   org_id = mist_org.terraform_test.id
#   name   = "test group"
# }
# resource "mist_sitegroup" "test_group2" {
#   org_id = mist_org.terraform_test.id
#   name   = "test group2b"
# }

# resource "mist_service" "lab" {
#   org_id = mist_org.terraform_test.id
#   addresses = [
#     "10.3.0.0/24", "10.4.0.0/24"
#   ]
#   description = "the lab network"
#   name        = "lab_network"
#   type        = "custom"
#   specs = [
#     {
#       protocol   = "tcp"
#       port_range = "443"
#     }
#   ]
# }


# resource "mist_network" "corp" {
#   org_id                 = mist_org.terraform_test.id
#   disallow_mist_services = false
#   name                   = "prd_corp"
#   subnet                 = "10.3.0.0"
# }

# resource "mist_networktemplate" "switch_template" {
#   name                   = "test switch"
#   org_id                 = mist_org.terraform_test.id
#   dns_servers            = ["10.3.51.222"]
#   dns_suffix             = ["stag.one"]
#   ntp_servers            = ["10.3.51.222"]
#   additional_config_cmds = ["set system hostnam test", "set system services ssh root-login allow"]
#   networks = {
#     test = {
#       subnet  = "1.2.3.4"
#       vlan_id = 10
#     }
#     test2 = {
#       subnet  = "1.2.3.4"
#       vlan_id = 11

#     }
#   }
#   radius_config = {
#     acct_interim_interval = 60
#     coa_enabled           = true
#     network               = "test"
#     acct_servers = [
#       {
#         host   = "1.2.3.4"
#         secret = "secret"
#       }
#     ]
#     auth_servers = [
#       {
#         host   = "1.2.3.4"
#         secret = "secret"
#       }
#     ]
#   }
#   port_usages = {
#     trunk = {
#       all_networks  = true
#       description   = "profile for trunk ports"
#       enable_qos    = true
#       mode          = "trunk"
#       port_networks = "test2"
#     }
#   }
#   remote_syslog = {
#     archive = {
#       files = 2
#       size  = "5m"
#     }
#     console = {
#       contents = [
#         {
#           facility = "kernel"
#           severity = "alert"
#         }
#       ]
#     }
#     enabled             = true
#     network             = "test"
#     send_to_all_servers = false
#     servers = [
#       {
#         contents = [
#           {
#             facility = "any"
#             severity = "any"
#           }
#         ]
#         host = "1.2.3.4"
#       }
#     ]
#     structured_data = true
#   }
#   switch_mgmt = {
#     config_revert = 5
#     protect_re = {
#       enabled = true
#     }
#     root_password = "Juniper123"
#   }
#   switch_matching = {
#     enable = true
#     rules = [
#       {
#         match_type  = "match_name[0:3]"
#         match_value = "abc"
#         additional_config_cmds = [
#           "set system name-server 8.8.8.8"
#         ]
#         match_role = "access"
#         name       = "access"
#         port_config = {
#           "ge-0/0/0-10" = {
#             usage = "trunk"
#           }
#         }
#       },
#       {
#         additional_config_cmds = [
#           "set system name-server 8.8.8.8"
#         ]
#         match_role = "core"
#         name       = "core"
#         port_config = {
#           "ge-0/0/0-10" = {
#             usage = "trunk"
#           }
#         }
#       }
#     ]
#   }
# }


# resource "mist_network" "mgmt" {
#   org_id  = mist_org.terraform_test.id
#   vlan_id = 172
#   subnet  = "10.3.172.0/24"
#   gateway = "10.3.172.9"
#   vpn_access = {
#     "OrgOverlay" : {
#       routed                    = true
#       no_readvertise_to_overlay = false
#       no_readvertise_to_lan_bgp = false
#     }
#   }
#   isolation              = true
#   disallow_mist_services = false
#   name                   = "SRX-Mgmt"
# }
# resource "mist_network" "reg" {
#   org_id                 = mist_org.terraform_test.id
#   isolation              = true
#   vlan_id                = 12
#   subnet                 = "10.3.12.0/24"
#   disallow_mist_services = false
#   name                   = "SRX-REG"
# }
# resource "mist_network" "ssr" {
#   org_id                 = mist_org.terraform_test.id
#   isolation              = true
#   vlan_id                = 128
#   subnet                 = "10.128.100.0/16"
#   disallow_mist_services = false
#   name                   = "SRX-Core-128T"
# }
# resource "mist_network" "mxe" {
#   org_id                 = mist_org.terraform_test.id
#   isolation              = true
#   vlan_id                = 1000
#   subnet                 = "10.10.0.0/24"
#   disallow_mist_services = false
#   tenants = {
#     "mxe-ext" = {
#       addresses = [
#         "10.10.0.10"
#       ]
#     }
#   }
#   internet_access = {
#     static_nat = {}
#     destination_nat = {
#       "192.168.1.9:500" : {
#         name        = "ike"
#         internal_ip = "10.10.0.10"
#         port        = "500"
#       }
#       "192.168.1.9:4500" : {
#         name        = "isakmp"
#         internal_ip = "10.10.0.10"
#         port        = "4500"
#       }
#     }
#   }
#   name = "SRX-MXE-tt-in"
# }
# resource "mist_network" "iot" {
#   org_id                 = mist_org.terraform_test.id
#   isolation              = true
#   vlan_id                = 8
#   subnet                 = "10.3.8.0/24"
#   disallow_mist_services = true
#   tenants = {
#     netatmo = {
#       addresses = [
#         "10.3.8.201/32"
#       ]
#     }
#     "cctv-e" = {
#       addresses = [
#         "10.3.8.40/32"
#       ]
#     }
#     "cctv-b" = {
#       addresses = [
#         "10.3.8.41/32"
#       ]
#     }
#     "cctv-c" = {
#       addresses = [
#         "10.3.8.42/32"
#       ]
#     }
#     "cctv-x" = {
#       addresses = [
#         "10.3.8.44/32"
#       ]
#     }
#   }
#   name = "SRX-IoT"
# }
# resource "mist_network" "dmz" {
#   org_id                 = mist_org.terraform_test.id
#   isolation              = true
#   vlan_id                = 3
#   subnet                 = "10.3.3.0/24"
#   disallow_mist_services = true
#   tenants = {
#     aws = {
#       addresses = [
#         "10.3.3.3"
#       ]
#     }
#   }
#   internet_access = {
#     static_nat = {}
#     destination_nat = {
#       "192.168.1.9:5431" : {
#         name        = "aws-wg"
#         internal_ip = "10.3.3.3"
#         port        = "51820"
#       }
#     }
#   }
#   name = "SRX-DMZ"
# }


# resource "mist_gatewaytemplate" "stag" {

#   port_config = {
#     "ge-0/0/3" = {
#       name       = "FTTH"
#       usage      = "wan"
#       aggregated = false
#       redundant  = false
#       critical   = false
#       wan_type   = "broadband"
#       ip_config = {
#         type    = "static"
#         ip      = "192.168.1.8"
#         netmask = "/24"
#         gateway = "192.168.1.1"
#       },
#       disable_autoneg = false
#       speed           = "auto"
#       duplex          = "auto"
#       wan_source_nat = {
#         disabled = false
#       },
#       vpn_paths = {
#         "SSR_HUB_DC-MPLS.OrgOverlay" = {
#           key         = 0
#           role        = "spoke"
#           bfd_profile = "broadband"
#         }
#       }
#     },
#     "ge-0/0/4" = {
#       name       = "LTE"
#       usage      = "wan"
#       aggregated = false
#       redundant  = false
#       critical   = false
#       wan_type   = "broadband"
#       ip_config = {
#         type = "dhcp"
#       },
#       disable_autoneg = false
#       speed           = "auto"
#       duplex          = "auto"
#       wan_source_nat = {
#         disabled = false
#       },
#       vpn_paths = {
#         "SSR_HUB_DC-MPLS.OrgOverlay" = {
#           key         = 0
#           role        = "spoke"
#           bfd_profile = "broadband"
#         }
#       }
#     },
#     "ge-0/0/5" = {
#       usage            = "lan"
#       critical         = false
#       aggregated       = true
#       ae_disable_lacp  = false
#       ae_lacp_force_up = true
#       ae_idx           = 0
#       redundant        = false
#       networks = [
#         "PRD-Core"
#       ]
#     },
#     "ge-0/0/7" = {
#       usage      = "lan"
#       critical   = false
#       aggregated = false
#       redundant  = false
#       networks = [
#         "PRD-Mgmt"
#       ]
#     },
#     "ge-0/0/6" = {
#       usage      = "lan"
#       critical   = false
#       aggregated = false
#       redundant  = false
#       networks = [
#         "PRD-Lab"
#       ]
#     }
#   }
#   ip_configs = {
#     "PRD-Core" = {
#       type    = "static"
#       ip      = "10.3.100.9"
#       netmask = "/24"
#     },
#     "PRD-Mgmt" = {
#       type    = "static"
#       ip      = "10.3.172.1"
#       netmask = "/24"
#     },
#     "PRD-Lab" = {
#       type    = "static"
#       ip      = "10.3.171.1"
#       netmask = "/24"
#     }
#   }
#   dhcpd_config = {
#     enable = true
#   }
#   path_preferences = {
#     "HUB" = {
#       strategy = "ordered"
#       paths = [
#         {
#           name = "SSR_HUB_DC-MPLS.OrgOverlay"
#           type = "vpn"
#         }
#       ]
#     },
#     "HUB-ORDERED" = {
#       strategy = "ordered"
#       paths = [
#         {
#           name     = "SSR_HUB_DC-MPLS.OrgOverlay"
#           wan_name = "FTTH",
#           type     = "vpn"
#         },
#         {
#           name     = "SSR_HUB_DC-MPLS.OrgOverlay"
#           wan_name = "LTE"
#           type     = "vpn"
#         }
#       ]
#     },
#     "HUB-ECMP" = {
#       strategy = "weighted"
#       paths = [
#         {
#           name     = "SSR_HUB_DC-MPLS.OrgOverlay"
#           wan_name = "LTE"
#           cost     = 30
#           type     = "vpn"
#         },
#         {
#           name     = "SSR_HUB_DC-MPLS.OrgOverlay"
#           wan_name = "FTTH"
#           cost     = 30
#           type     = "vpn"
#         }
#       ]
#     }
#   }
#   service_policies = [
#     {
#       name = "Policy-14"
#       tenants = [
#         "PRD-Core"
#       ],
#       services = [
#         "any"
#       ],
#       action          = "allow"
#       path_preference = "HUB"
#       idp = {
#         enabled    = true
#         profile    = "critical"
#         alert_only = false
#       }
#     },
#     {
#       name = "Policy-2"
#       tenants = [
#         "PRD-Mgmt"
#       ],
#       services = [
#         "any"
#       ],
#       action          = "allow",
#       path_preference = "HUB-ECMP"
#       idp = {
#         enabled    = true,
#         profile    = "standard"
#         alert_only = true
#       }
#     },
#     {
#       name = "Policy-3"
#       tenants = [
#         "PRD-Lab"
#       ],
#       services = [
#         "any"
#       ],
#       action          = "allow"
#       path_preference = "HUB-ORDERED"
#       idp = {
#         enabled = false
#       }
#     }
#   ]
#   type   = "spoke"
#   name   = "test-api"
#   org_id = mist_org.terraform_test.id
# }

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
  type = "spoke"
  name   = "sdwan-westford"
  org_id = mist_org.terraform_test.id

}
