terraform {
  required_providers {
    mist = {
      source = "registry.terraform.io/juniper/mist"
    }
  }
}

provider "mist" {
  host = local.envs["HOST"]
  // apitoken = local.envs["APITOKEN"]
  // proxy = "test_proxy:Jun1per!Mist@10.3.18.11:3128"
  username = local.envs["USERNAME"]
  password = local.envs["PASSWORD"]
}


resource "mist_site_networktemplate" "site_evpn" {
  site_id = (local.site_id)
  networks = {
    "corp" = {
      vlan_id = 100
      subnet  = "10.99.100.0/24"
      gateway = "10.99.100.1"
    },
    "mgmt" = {
      vlan_id = 110
      subnet  = "10.99.110.0/24"
      gateway = "10.99.110.1"
    }
  }
  vrf_instances = {
    "evpn_vrf" = {
      networks = ["corp"]
      extra_routes = {
        "10.10.0.0/16" : {
          via = "10.99.110.254"
        }
      }
    }
  }
}
# 020004a6c1f7
# 02000467ac7f

# 020004b87dbd
# 020004f4b1bd

# 020004a5c06d
# 0200044539d0
resource "mist_org_inventory" "inventory" {

  org_id = (local.org_id)
  inventory = {
    (local.sw0) = {
      site_id                = (local.site_id)
      unclaim_when_destroyed = false
    },
    (local.sw1) = {
      site_id                = (local.site_id)
      unclaim_when_destroyed = false
    },
    (local.sw2) = {
      site_id                = (local.site_id)
      unclaim_when_destroyed = false
    },
    (local.sw3) = {
      site_id                = (local.site_id)
      unclaim_when_destroyed = false
    },
    (local.sw4) = {
      site_id                = (local.site_id)
      unclaim_when_destroyed = false
    },
    (local.sw5) = {
      site_id                = (local.site_id)
      unclaim_when_destroyed = false
    }
  }
}




resource "mist_site_evpn_topology" "ip_clos" {
  site_id = (local.site_id)
  name    = "evpn_one"
  evpn_options = {
    routed_at = "core"
    overlay = {
      as = 65000
    },
    core_as_border        = true
    auto_loopback_subnet  = "172.16.192.0/24"
    auto_loopback_subnet6 = "fd33:ab00:2::/64"
    per_vlan_vga_v4_mac   = false
    underlay = {
      as_base  = 65001
      use_ipv6 = false
      subnet   = "10.255.240.0/20"
    }
    auto_router_id_subnet = "172.16.254.0/23"
  }
  switches = {
    (local.sw0) = {
      role = "core"
    },
    (local.sw1) = {
      role = "core"
    },
    (local.sw2) = {
      role = "distribution"
    },
    (local.sw3) = {
      role = "distribution"
    },
    (local.sw4) = {
      role = "access"
    },
    (local.sw5) = {
      role = "access"
    }
  }
}

resource "mist_device_switch" "switch_core_01" {
  device_id = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw0)).id
  site_id   = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw0)).site_id
  name      = "core-01"
  managed   = true
  port_config = {
    "et-0/0/0,et-0/0/1" = {
      usage = "evpn_downlink"
    }
  }
  other_ip_configs = {
    "corp" = {
      type    = "static"
      ip      = "10.99.100.2"
      netmask = "255.255.255.0"
    }
    "mgmt" = {
      type    = "static"
      ip      = "10.99.110.2"
      netmask = "255.255.255.0"
    }
  }
}
resource "mist_device_switch" "switch_core_02" {
  device_id = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw1)).id
  site_id   = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw1)).site_id
  name      = "core-02"
  managed   = true
  port_config = {
    "et-0/0/0,et-0/0/1" = {
      usage = "evpn_downlink"
    }
  }
  other_ip_configs = {
    "corp" = {
      type    = "static"
      ip      = "10.99.100.3"
      netmask = "255.255.255.0"
    }
    "mgmt" = {
      type    = "static"
      ip      = "10.99.110.3"
      netmask = "255.255.255.0"
    }
  }
}
resource "mist_device_switch" "switch_distri_01" {
  device_id = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw2)).id
  site_id   = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw2)).site_id
  name      = "distri-01"
  managed   = true
  port_config = {
    "et-0/0/0,et-0/0/1" = {
      usage = "evpn_uplink"
    },
    "et-0/0/2,et-0/0/3" = {
      usage = "evpn_downlink"
    }
  }
}
resource "mist_device_switch" "switch_distri_02" {
  device_id = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw3)).id
  site_id   = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw3)).site_id
  name      = "distri-02"
  managed   = true
  port_config = {
    "et-0/0/0,et-0/0/1" = {
      usage = "evpn_uplink"
    },
    "et-0/0/2,et-0/0/3" = {
      usage = "evpn_downlink"
    }
  }
}
resource "mist_device_switch" "switch_access_01" {
  device_id = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw4)).id
  site_id   = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw4)).site_id
  name      = "access-01"
  managed   = true
  port_config = {
    "et-0/0/0,et-0/0/1" = {
      usage = "evpn_uplink"
    }
  }
}
resource "mist_device_switch" "switch_access_02" {
  device_id = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw5)).id
  site_id   = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw5)).site_id
  name      = "access-02"
  managed   = true
  port_config = {
    "et-0/0/0,et-0/0/1" = {
      usage = "evpn_uplink"
    }
  }
}
