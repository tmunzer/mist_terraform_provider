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
    vlan1099 = {
      vlan_id = 1099
      subnet  = "10.99.99.0/24"
      gateway = "10.99.99.1"
    },
    vlan1088 = {
      vlan_id = 1088
      subnet  = "10.88.88.0/24"
      gateway = "10.88.88.1"
    },
    vlan1033 = {
      vlan_id = 1033
      subnet  = "10.33.33.0/24"
      gateway = "10.33.33.1"
    }

  }
  vrf_instances = {
    "customera" = {
      networks = ["vlan1099"]
      extra_routes = {
        "0.0.0.0/0" : {
          via = "10.99.99.254"
        }
      }
    }
    "customerb" = {
      networks = ["vlan1088"]
      extra_routes = {
        "0.0.0.0/0" : {
          via = "10.88.88.254"
        }
      }
    }
    "devices" = {
      networks = ["vlan1033"]
      extra_routes = {
        "0.0.0.0/0" : {
          via = "10.33.33.254"
        }
      }
    }
  }
  port_usages = {
    l2fabricexit = {
      mode         = "trunk"
      all_networks = false
      networks = [
        "vlan1099",
        "vlan1088",
        "vlan1033"
      ]
      mtu = "9018"
    }
    vlan1099 = {
      mode         = "access"
      port_network = "vlan1099"
    },
    vlan1088 = {
      mode         = "access"
      port_network = "vlan1088"
    },
    vlan1033 = {
      mode         = "access"
      port_network = "vlan1033"
    }
    x-esilag = {
      mode         = "trunk"
      all_networks = false
      networks = [
        "vlan1099",
        "vlan1088",
        "vlan1033"
      ]
      poe_disabled = true
      mtu          = "9100"
      ui_evpntopo_id = mist_site_evpn_topology.code_distri.id
    }
  }
}

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

resource "mist_site_evpn_topology" "code_distri" {
  site_id = (local.site_id)
  name    = "evpn_one"
  evpn_options = {
    routed_at = "core"
    per_vlan_vga_v4_mac = true
    overlay = {
      as = 65000
    },
    core_as_border        = true
    auto_loopback_subnet  = "172.16.192.0/24"
    auto_loopback_subnet6 = "fd33:ab00:2::/64"
    underlay = {
      as_base  = 65001
      use_ipv6 = false
      subnet   = "10.255.240.0/20"
    }
  }
  pod_names = {
    "1" = "Pod 1"
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
      pod  = 1
    },
    (local.sw3) = {
      role = "distribution"
      pod  = 1
    },
    (local.sw4) = {
      role = "esilag-access"
      pod  = 1
    },
    (local.sw5) = {
      role = "esilag-access"
      pod  = 1
    }
  }
}

resource "mist_device_switch" "switch_core_01" {
  device_id = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw0)).id
  site_id   = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw0)).site_id
  name      = "core-01"
  managed   = true
  port_config = {
    "ge-0/0/0" = {
      usage        = "l2fabricexit"
      aggregated   = true
      esilag       = true
      ae_idx       = 0
    }
    "ge-0/0/2,ge-0/0/1" = {
      usage = "evpn_downlink"
    }
  }
  other_ip_configs = {
    "vlan1099" = {
      type         = "static"
      ip           = "10.99.99.2"
      netmask      = "255.255.255.0"
      evpn_anycast = true
    }
    "vlan1088" = {
      type         = "static"
      ip           = "10.88.88.2"
      netmask      = "255.255.255.0"
      evpn_anycast = true
    }
    "vlan1033" = {
      type         = "static"
      ip           = "10.33.33.2"
      netmask      = "255.255.255.0"
      evpn_anycast = true
    }
  }
  vrf_config = {
    enabled = true
  }
}
resource "mist_device_switch" "switch_core_02" {
  device_id = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw1)).id
  site_id   = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw1)).site_id
  name      = "core-02"
  managed   = true
  port_config = {
    "ge-0/0/0" = {
      usage        = "l2fabricexit"
      aggregated   = true
      esilag       = true
      ae_idx       = 0
    }
    "ge-0/0/2,ge-0/0/1" = {
      usage = "evpn_downlink"
    }
  }
  other_ip_configs = {
    "vlan1099" = {
      type         = "static"
      ip           = "10.99.99.3"
      netmask      = "255.255.255.0"
      evpn_anycast = true
    }
    "vlan1088" = {
      type         = "static"
      ip           = "10.88.88.3"
      netmask      = "255.255.255.0"
      evpn_anycast = true
    }
    "vlan1033" = {
      type         = "static"
      ip           = "10.33.33.3"
      netmask      = "255.255.255.0"
      evpn_anycast = true
    }
  }
  vrf_config = {
    enabled = true
  }
}
resource "mist_device_switch" "switch_distri_01" {
  device_id = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw2)).id
  site_id   = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw2)).site_id
  name      = "distri-01"
  managed   = true
  port_config = {
    "ge-0/0/1" = {
      usage = "x-esilag"
      aggregated   = true
      esilag       = true
      ae_idx       = 10
    },
    "ge-0/0/2" = {
      usage = "x-esilag"
      aggregated   = true
      esilag       = true
      ae_idx       = 11
    },
    "ge-0/0/3,ge-0/0/4" = {
      usage = "evpn_uplink"
    }
  }
}
resource "mist_device_switch" "switch_distri_02" {
  device_id = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw3)).id
  site_id   = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw3)).site_id
  name      = "distri-02"
  managed   = true
  port_config = {
    "ge-0/0/1" = {
      usage = "x-esilag"
      aggregated   = true
      esilag       = true
      ae_idx       = 10
    },
    "ge-0/0/2" = {
      usage = "x-esilag"
      aggregated   = true
      esilag       = true
      ae_idx       = 11
    },
    "ge-0/0/3,ge-0/0/4" = {
      usage = "evpn_uplink"
    }
  }
}
resource "mist_device_switch" "switch_access_01" {
  device_id = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw4)).id
  site_id   = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw4)).site_id
  name      = "access-01"
  managed   = true
  port_config = {
    "ge-0/0/1,ge-0/0/2" = {
      usage = "x-esilag"
      aggregated   = true
      esilag       = true
      ae_idx       = 10
    }
    "ge-0/0/3" = {
      usage = "vlan1099"
    }
  }

}
resource "mist_device_switch" "switch_access_02" {
  device_id = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw5)).id
  site_id   = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, (local.sw5)).site_id
  name      = "access-02"
  managed   = true
  port_config = {
    "ge-0/0/1,ge-0/0/2" = {
      usage = "x-esilag"
      aggregated   = true
      esilag       = true
      ae_idx       = 11
    }
    "ge-0/0/3" = {
      usage = "vlan1088"
    }
  }
}
