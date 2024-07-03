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
  gatewaytemplate_id = mist_org_gatewaytemplate.my_template.id
}

### ORG LEVEL SERVICES
resource "mist_org_service" "lab" {
  org_id = mist_org.terraform_org.id
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


### ORG LEVEL NETWORKS
resource "mist_org_network" "corp" {
  org_id                 = mist_org.terraform_org.id
  disallow_mist_services = false
  name                   = "prd_corp"
  subnet                 = "10.3.0.0"
}

resource "mist_org_network" "mgmt" {
  org_id  = mist_org.terraform_org.id
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
  org_id                 = mist_org.terraform_org.id
  isolation              = true
  vlan_id                = 12
  subnet                 = "10.3.12.0/24"
  disallow_mist_services = false
  name                   = "SRX-REG"
}
resource "mist_org_network" "ssr" {
  org_id                 = mist_org.terraform_org.id
  isolation              = true
  vlan_id                = 128
  subnet                 = "10.128.100.0/16"
  disallow_mist_services = false
  name                   = "SRX-Core-128T"
}
resource "mist_org_network" "mxe" {
  org_id                 = mist_org.terraform_org.id
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
  org_id                 = mist_org.terraform_org.id
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
  org_id                 = mist_org.terraform_org.id
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


### ORG GATEWAY TEMPLATE
resource "mist_org_gatewaytemplate" "my_template" {
  org_id = mist_org.terraform_org.id
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
}
