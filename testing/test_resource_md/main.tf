terraform {
  required_providers {
    mist = {
      source = "registry.terraform.io/juniper/mist"
    }
  }
}

provider "mist" {
  host = local.envs["HOST"]
  //apitoken = local.envs["APITOKEN"]
  //proxy = "test_proxy:Jun1per!Mist@10.3.18.11:3128"
  username = local.envs["USERNAME"]
  password = local.envs["PASSWORD"]
}

### ORG
resource "mist_org" "terraform_test" {
  name = "Terraform Testing - MD"
}

resource "mist_org_setting" "terraform_test" {
  org_id                   = mist_org.terraform_test.id
  ap_updown_threshold      = 10
  device_updown_threshold  = 10
  disable_pcap             = false
  disable_remote_shell     = true
  gateway_updown_threshold = 10
  mxedge_mgmt = {
    mist_password = "mist_secret_passowrd"
    root_password = "root_secret_password"
    oob_ip_type   = "dhcp"
    oob_ip_type6  = "disabled"
  }
  password_policy = {
    enabled                  = false
    freshness                = 180
    min_length               = 12
    requires_special_char    = true
    requires_two_factor_auth = false
  }
  security = {
    disable_local_ssh = false
  }
  switch_updown_threshold = 10
  synthetic_test = {
    disabled = false
    vlans = [{
      vlan_ids         = ["10", "30"]
      custom_test_urls = ["http://www.abc.com/", "https://10.3.5.1:8080/about"]
      }, {
      vlan_ids = ["20"]
      disabled = true
    }]
  }
}

resource "mist_site" "terraform_site" {
  name    = "terraform - md"
  org_id  = mist_org.terraform_test.id
  address = "41 rue de villiers, 92100 Neuilly sur seine, France"
}

resource "mist_site_setting" "terraform_site" {
  site_id = mist_site.terraform_site.id
  auto_upgrade = {
    custom_versions = {
      AP32 = "0.14.29411",
      AP43 = "0.12.27139",
      AP64 = "0.14.29411"
    },
    day_of_week = "sun",
    enabled     = true,
    time_of_day = "02:00",
    version     = "custom"
  }
  ap_updown_threshold     = 0
  config_auto_revert      = false
  device_updown_threshold = 0
  gateway_mgmt = {
    app_probing = {
      apps = [
        "icloud",
        "aws",
        "facebook",
        "azure",
        "ms-teams",
        "office365",
        "okta",
        "webex"
      ],
      custom_apps = [
        {
          hostnames = [
            "https://vhqmcd.verifone.com"
          ],
          name     = "Verifone",
          protocol = "http"
        }
      ],
      enabled = true
    },
    app_usage           = true,
    config_revert_timer = 10,
    disable_console     = false,
    disable_oob         = false,
    protect_re = {
      allowed_services = [
        "icmp",
        "ssh"
      ],
      enabled = true
    },
    root_password = "VftgfW7L13"
  }
  gateway_updown_threshold = 0
  persist_config_on_device = true
  remove_existing_configs  = false
  report_gatt              = false
  rogue = {
    enabled          = true,
    honeypot_enabled = true,
    min_duration     = 10,
    min_rssi         = -80
  }
  switch_updown_threshold = 0
  track_anonymous_devices = false
  vars = {
    branch_ip                          = "10.220.93",
    branch_no_routable                 = "172.24.151",
    crew_network                       = "29.220.93",
    iot_net_wired                      = "30.220.93",
    iot_net_wired_upper                = "192.168.93",
    iot_net_wireless                   = "30.220.93",
    iot_net_wireless_upper             = "30.220.93",
    local_ASN                          = "64777",
    peer_ASN                           = "64781",
    private_net                        = "29.220.93",
    public_dns1                        = "64.134.255.10",
    public_dns2                        = "64.134.255.2",
    public_dns3                        = "208.67.222.222",
    public_dns4                        = "208.67.220.220",
    public_net                         = "192.168.4",
    public_net_upper                   = "192.168.5",
    quarantine_net                     = "29.220.93",
    sc_local_id_1                      = "US_Prod_04270_1@accenture.net",
    sc_local_id_2                      = "US_Prod_04270_2@accenture.net",
    sc_remote_id_1                     = "prismacentral_1@palo.net",
    sc_remote_id_2                     = "prismacentral_2@palo.net",
    secondary_crew_network             = "192.168.6",
    secondary_iot_net_wireless         = "192.168.94",
    secondary_iot_wired                = "192.168.92",
    secondary_private_net              = "192.168.91",
    secondary_untrusted_net            = "172.24.151",
    secondary_untrusted_net_full_lower = "172.24.151.64",
    secondary_untrusted_net_full_upper = "172.24.151.65",
    service_ip_1                       = "137.83.229.89",
    service_ip_2                       = "140.209.233.142",
    store_nat                          = "7.0.0",
    tunnel_monitor_ip                  = "10.220.93.62",
    untrusted_net                      = "29.220.93"
  }
}
resource "mist_site_networktemplate" "terraform_site" {
  site_id = mist_site.terraform_site.id
  additional_config_cmds = [
    "set chassis alarm management-ethernet link-down ignore",
    "#Decouple ARP table and Authentication sessions",
    "set protocols dot1x authenticator no-mac-table-binding",
    "set protocols dot1x authenticator interface dot1x-mab_only reauthentication 43200"
  ]
  dhcp_snooping = {
    all_networks           = false,
    enable_arp_spoof_check = false,
    enable_ip_source_guard = false,
    enabled                = true,
    networks = [
      "PUBLIC"
    ]
  }
  dns_servers = [
    "{{public_dns1}}",
    "{{public_dns2}}"
  ]
  mist_nac = {
    enabled = true
  }
  networks = {
    AUTOMATION = {
      isolation = false,
      vlan_id   = "460"
    },
    CORE = {
      isolation = false,
      vlan_id   = "410"
    },
    CREW = {
      isolation = false,
      vlan_id   = "451"
    },
    Cradlepoint = {
      isolation = false,
      vlan_id   = "4000"
    },
    DESSERT_CART = {
      isolation = false,
      vlan_id   = "900"
    },
    HA = {
      isolation = false,
      vlan_id   = "10"
    },
    IOT_WIRED = {
      isolation = false,
      vlan_id   = "492"
    },
    IOT_WIRELESS = {
      isolation = false,
      vlan_id   = "494"
    },
    MANAGEMENT = {
      isolation = false,
      vlan_id   = "430"
    },
    PCI = {
      isolation = false,
      vlan_id   = "440"
    },
    PRIVATE = {
      isolation = false,
      vlan_id   = "491"
    },
    PUBLIC = {
      isolation = false,
      vlan_id   = "450"
    },
    QUARANTINE = {
      isolation = false,
      vlan_id   = "495"
    },
    RESTRICTED = {
      isolation = false,
      vlan_id   = "420"
    },
    UNTRUSTED = {
      isolation = false,
      vlan_id   = "490"
    }
  }
  ntp_servers = [
    "129.6.15.28",
    "129.6.15.29",
    "129.6.15.30"
  ]
  port_usages = {
    automation = {
      all_networks                                    = false,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "access",
      persist_mac                                     = false,
      poe_disabled                                    = false,
      port_network                                    = "AUTOMATION",
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      storm_control = {
        no_broadcast            = false,
        no_multicast            = false,
        no_registered_multicast = false,
        no_unknown_unicast      = false,
        percentage              = 5
      },
      stp_edge         = true,
      stp_no_root_port = false,
      stp_p2p          = false,
      use_vstp         = false
    },
    blackhole = {
      all_networks                                    = false,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "access",
      persist_mac                                     = false,
      poe_disabled                                    = false,
      port_network                                    = "QUARANTINE",
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      storm_control = {
        no_broadcast            = false,
        no_multicast            = false,
        no_registered_multicast = false,
        no_unknown_unicast      = false,
        percentage              = 5
      },
      stp_edge         = true,
      stp_no_root_port = false,
      stp_p2p          = false,
      use_vstp         = false
    },
    core = {
      all_networks                                    = false,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "access",
      persist_mac                                     = false,
      poe_disabled                                    = false,
      port_network                                    = "CORE",
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      storm_control = {
        no_broadcast            = false,
        no_multicast            = false,
        no_registered_multicast = false,
        no_unknown_unicast      = false,
        percentage              = 5
      },
      stp_edge         = true,
      stp_no_root_port = false,
      stp_p2p          = false,
      use_vstp         = false
    },
    corporate = {
      all_networks                                    = false,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "access",
      persist_mac                                     = false,
      poe_disabled                                    = false,
      port_network                                    = "RESTRICTED",
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      storm_control = {
        no_broadcast            = false,
        no_multicast            = false,
        no_registered_multicast = false,
        no_unknown_unicast      = false,
        percentage              = 5
      },
      stp_edge         = true,
      stp_no_root_port = false,
      stp_p2p          = false,
      use_vstp         = false
    },
    crew = {
      all_networks                                    = false,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "access",
      persist_mac                                     = false,
      poe_disabled                                    = false,
      port_network                                    = "CREW",
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      storm_control = {
        no_broadcast            = false,
        no_multicast            = false,
        no_registered_multicast = false,
        no_unknown_unicast      = false,
        percentage              = 5
      },
      stp_edge         = true,
      stp_no_root_port = false,
      stp_p2p          = false,
      use_vstp         = false
    },
    disabled_port = {
      all_networks                                    = false,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = true,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "access",
      persist_mac                                     = false,
      poe_disabled                                    = false,
      port_network                                    = "default",
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      stp_edge                                        = false,
      stp_no_root_port                                = false,
      stp_p2p                                         = false,
      use_vstp                                        = false
    },
    dot1x = {
      all_networks                                    = false,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "access",
      persist_mac                                     = false,
      poe_disabled                                    = false,
      port_auth                                       = "dot1x",
      port_network                                    = "QUARANTINE",
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      storm_control = {
        no_broadcast            = false,
        no_multicast            = false,
        no_registered_multicast = false,
        no_unknown_unicast      = false,
        percentage              = 5
      },
      stp_edge         = true,
      stp_no_root_port = false,
      stp_p2p          = false,
      use_vstp         = false
    },
    dot1x-mab_only = {
      all_networks                                    = false,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = true,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = true,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_only                                   = true,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "access",
      persist_mac                                     = false,
      poe_disabled                                    = false,
      port_auth                                       = "dot1x",
      port_network                                    = "QUARANTINE",
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      storm_control = {
        no_broadcast            = false,
        no_multicast            = false,
        no_registered_multicast = false,
        no_unknown_unicast      = false,
        percentage              = 5
      },
      stp_edge         = true,
      stp_no_root_port = false,
      stp_p2p          = false,
      use_vstp         = false
    },
    dynamic = {
      all_networks                                    = false,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "dynamic",
      persist_mac                                     = false,
      poe_disabled                                    = false,
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      stp_edge                                        = false,
      stp_no_root_port                                = false,
      stp_p2p                                         = false,
      use_vstp                                        = false
    },
    google_edge = {
      all_networks                                    = false,
      allow_dhcpd                                     = true,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "trunk",
      networks = [
        "CORE"
      ],
      persist_mac        = false,
      poe_disabled       = true,
      port_network       = "AUTOMATION",
      reauth_interval    = 3600,
      reset_default_when = "link_down",
      storm_control = {
        no_broadcast            = false,
        no_multicast            = false,
        no_registered_multicast = false,
        no_unknown_unicast      = false,
        percentage              = 5
      },
      stp_edge         = true,
      stp_no_root_port = false,
      stp_p2p          = false,
      use_vstp         = false
    },
    iot_wired = {
      all_networks                                    = false,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "access",
      persist_mac                                     = false,
      poe_disabled                                    = false,
      port_network                                    = "IOT_WIRED",
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      storm_control = {
        no_broadcast            = false,
        no_multicast            = false,
        no_registered_multicast = false,
        no_unknown_unicast      = false,
        percentage              = 5
      },
      stp_edge         = true,
      stp_no_root_port = false,
      stp_p2p          = false,
      use_vstp         = false
    },
    iot_wireless = {
      all_networks                                    = false,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "access",
      persist_mac                                     = false,
      poe_disabled                                    = false,
      port_network                                    = "IOT_WIRELESS",
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      storm_control = {
        no_broadcast            = false,
        no_multicast            = false,
        no_registered_multicast = false,
        no_unknown_unicast      = false,
        percentage              = 5
      },
      stp_edge         = true,
      stp_no_root_port = false,
      stp_p2p          = false,
      use_vstp         = false
    },
    mgmt = {
      all_networks                                    = false,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "access",
      persist_mac                                     = false,
      poe_disabled                                    = false,
      port_network                                    = "MANAGEMENT",
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      storm_control = {
        no_broadcast            = false,
        no_multicast            = false,
        no_registered_multicast = false,
        no_unknown_unicast      = false,
        percentage              = 5
      },
      stp_edge         = true,
      stp_no_root_port = false,
      stp_p2p          = false,
      use_vstp         = false
    },
    pci = {
      all_networks                                    = false,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "access",
      persist_mac                                     = false,
      poe_disabled                                    = false,
      port_network                                    = "PCI",
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      storm_control = {
        no_broadcast            = false,
        no_multicast            = false,
        no_registered_multicast = false,
        no_unknown_unicast      = false,
        percentage              = 5
      },
      stp_edge         = true,
      stp_no_root_port = false,
      stp_p2p          = false,
      use_vstp         = false
    },
    private = {
      all_networks                                    = false,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "access",
      persist_mac                                     = false,
      poe_disabled                                    = false,
      port_network                                    = "PRIVATE",
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      storm_control = {
        no_broadcast            = false,
        no_multicast            = false,
        no_registered_multicast = false,
        no_unknown_unicast      = false,
        percentage              = 5
      },
      stp_edge         = true,
      stp_no_root_port = false,
      stp_p2p          = false,
      use_vstp         = false
    },
    public = {
      all_networks                                    = false,
      allow_dhcpd                                     = false,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "access",
      persist_mac                                     = false,
      poe_disabled                                    = false,
      port_network                                    = "PUBLIC",
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      storm_control = {
        no_broadcast            = false,
        no_multicast            = false,
        no_registered_multicast = false,
        no_unknown_unicast      = false,
        percentage              = 15
      },
      stp_edge         = true,
      stp_no_root_port = false,
      stp_p2p          = false,
      use_vstp         = false
    },
    restricted = {
      all_networks                                    = false,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "access",
      persist_mac                                     = false,
      poe_disabled                                    = false,
      port_network                                    = "RESTRICTED",
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      storm_control = {
        no_broadcast            = false,
        no_multicast            = false,
        no_registered_multicast = false,
        no_unknown_unicast      = false,
        percentage              = 15
      },
      stp_edge         = true,
      stp_no_root_port = false,
      stp_p2p          = false,
      use_vstp         = false
    },
    server-esxi = {
      all_networks                                    = true,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "trunk",
      persist_mac                                     = false,
      poe_disabled                                    = false,
      port_network                                    = "default",
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      storm_control = {
        no_broadcast            = false,
        no_multicast            = false,
        no_registered_multicast = false,
        no_unknown_unicast      = false,
        percentage              = 10
      },
      stp_edge         = false,
      stp_no_root_port = false,
      stp_p2p          = false,
      use_vstp         = false
    },
    servers = {
      all_networks                                    = true,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "trunk",
      persist_mac                                     = false,
      poe_disabled                                    = false,
      port_network                                    = "default",
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      stp_edge                                        = false,
      stp_no_root_port                                = false,
      stp_p2p                                         = false,
      use_vstp                                        = false
    },
    store-ap = {
      all_networks                                    = false,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "trunk",
      networks = [
        "CREW",
        "IOT_WIRELESS",
        "PUBLIC",
        "RESTRICTED",
        "CORE"
      ],
      persist_mac        = false,
      poe_disabled       = false,
      port_network       = "default",
      reauth_interval    = 3600,
      reset_default_when = "link_down",
      storm_control = {
        no_broadcast            = false,
        no_multicast            = false,
        no_registered_multicast = false,
        no_unknown_unicast      = false,
        percentage              = 5
      },
      stp_edge         = true,
      stp_no_root_port = false,
      stp_p2p          = false,
      use_vstp         = false
    },
    untrusted = {
      all_networks                                    = false,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "access",
      persist_mac                                     = false,
      poe_disabled                                    = false,
      port_network                                    = "UNTRUSTED",
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      storm_control = {
        no_broadcast            = false,
        no_multicast            = false,
        no_registered_multicast = false,
        no_unknown_unicast      = false,
        percentage              = 5
      },
      stp_edge         = true,
      stp_no_root_port = false,
      stp_p2p          = false,
      use_vstp         = false
    },
    uplink-srx-wan = {
      all_networks                                    = true,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "trunk",
      persist_mac                                     = false,
      poe_disabled                                    = false,
      port_network                                    = "default",
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      storm_control = {
        no_broadcast            = false,
        no_multicast            = false,
        no_registered_multicast = false,
        no_unknown_unicast      = false,
        percentage              = 5
      },
      stp_edge         = false,
      stp_no_root_port = false,
      stp_p2p          = false,
      use_vstp         = false
    },
    uplink-switch = {
      all_networks                                    = true,
      allow_multiple_supplicants                      = false,
      bypass_auth_when_server_down                    = false,
      bypass_auth_when_server_down_for_unkonwn_client = false,
      disable_autoneg                                 = false,
      disabled                                        = false,
      duplex                                          = "auto",
      enable_mac_auth                                 = false,
      enable_qos                                      = false,
      inter_switch_link                               = false,
      mac_auth_protocol                               = "eap-md5",
      mac_limit                                       = 0,
      mode                                            = "trunk",
      persist_mac                                     = false,
      poe_disabled                                    = true,
      port_network                                    = "default",
      reauth_interval                                 = 3600,
      reset_default_when                              = "link_down",
      storm_control = {
        no_broadcast            = false,
        no_multicast            = false,
        no_registered_multicast = false,
        no_unknown_unicast      = false,
        percentage              = 2
      },
      stp_edge         = false,
      stp_no_root_port = false,
      stp_p2p          = false,
      use_vstp         = false
    }
  }
  switch_matching = {
    enable = true,
    rules = [
      {
        additional_config_cmds = [
          "set protocols rstp bridge-priority 28k",
          "set protocols rstp interface uplink-switch no-root-port"
        ],
        match_role = "usa-style1-switch-1",
        name       = "usa-style1-switch-1",
        port_config = {
          "ge-0/0/0-1" = {
            ae_lacp_slow       = true,
            aggregated         = false,
            critical           = true,
            disable_autoneg    = false,
            duplex             = "auto",
            mtu                = 1514,
            no_local_overwrite = true,
            poe_disabled       = false,
            speed              = "auto",
            usage              = "core"
          },
          "ge-0/0/2-39" = {
            ae_lacp_slow       = true,
            aggregated         = false,
            critical           = false,
            disable_autoneg    = false,
            duplex             = "auto",
            mtu                = 1514,
            no_local_overwrite = false,
            poe_disabled       = false,
            speed              = "auto",
            usage              = "dot1x-mab_only"
          },
          "ge-0/0/44-45" = {
            ae_lacp_slow       = true,
            aggregated         = false,
            critical           = true,
            disable_autoneg    = false,
            duplex             = "auto",
            mtu                = 1514,
            no_local_overwrite = true,
            poe_disabled       = false,
            speed              = "auto",
            usage              = "store-ap"
          },
          "ge-0/0/46-47" = {
            ae_disable_lacp    = true,
            ae_idx             = 0,
            ae_lacp_slow       = false,
            aggregated         = true,
            critical           = true,
            disable_autoneg    = false,
            duplex             = "auto",
            mtu                = 1514,
            no_local_overwrite = true,
            poe_disabled       = false,
            speed              = "auto",
            usage              = "uplink-srx-wan"
          },
          "xe-0/1/0-1" = {
            ae_lacp_slow       = true,
            aggregated         = false,
            critical           = true,
            disable_autoneg    = false,
            duplex             = "auto",
            mtu                = 1514,
            no_local_overwrite = true,
            poe_disabled       = false,
            speed              = "auto",
            usage              = "uplink-switch"
          },
          "xe-0/1/2-3, ge-0/0/40-42" = {
            ae_lacp_slow       = true,
            aggregated         = false,
            critical           = false,
            disable_autoneg    = false,
            duplex             = "auto",
            mtu                = 1514,
            no_local_overwrite = true,
            poe_disabled       = false,
            speed              = "auto",
            usage              = "disabled_port"
          }
        },
        port_mirroring = {
          mcd-sw1-span = {
            input_networks_ingress = [
              "CORE"
            ],
            input_port_ids_egress  = [],
            input_port_ids_ingress = [],
            output_port_id         = "ge-0/0/42"
          }
        }
      },
      {
        additional_config_cmds = [
          ""
        ],
        match_role = "usa-style1-switch-2",
        name       = "usa-style1-switch-2",
        port_config = {
          "ge-0/0/0-39" = {
            ae_lacp_slow       = true,
            aggregated         = false,
            critical           = false,
            disable_autoneg    = false,
            duplex             = "auto",
            mtu                = 1514,
            no_local_overwrite = false,
            poe_disabled       = false,
            speed              = "auto",
            usage              = "dot1x-mab_only"
          },
          "ge-0/0/40-42, xe-0/1/2-3" = {
            ae_lacp_slow       = true,
            aggregated         = false,
            critical           = false,
            disable_autoneg    = false,
            duplex             = "auto",
            mtu                = 1514,
            no_local_overwrite = true,
            poe_disabled       = false,
            speed              = "auto",
            usage              = "disabled_port"
          },
          "ge-0/0/44-45" = {
            ae_lacp_slow       = true,
            aggregated         = false,
            critical           = true,
            disable_autoneg    = false,
            duplex             = "auto",
            mtu                = 1514,
            no_local_overwrite = true,
            poe_disabled       = false,
            speed              = "auto",
            usage              = "store-ap"
          },
          "ge-0/0/46-47" = {
            ae_disable_lacp    = true,
            ae_idx             = 0,
            ae_lacp_slow       = false,
            aggregated         = true,
            critical           = true,
            disable_autoneg    = false,
            duplex             = "auto",
            mtu                = 1514,
            no_local_overwrite = true,
            poe_disabled       = false,
            speed              = "auto",
            usage              = "uplink-srx-wan"
          },
          "xe-0/1/0-1" = {
            ae_lacp_slow       = true,
            aggregated         = false,
            critical           = true,
            disable_autoneg    = false,
            duplex             = "auto",
            mtu                = 1514,
            no_local_overwrite = true,
            poe_disabled       = false,
            speed              = "auto",
            usage              = "uplink-switch"
          }
        }
      },
      {
        additional_config_cmds = [
          "set protocols rstp bridge-priority 36k"
        ],
        match_role = "usa-style1-switch-3",
        name       = "usa-style1-switch-3",
        port_config = {
          "ge-0/0/0-39" = {
            ae_lacp_slow       = true,
            aggregated         = false,
            critical           = false,
            disable_autoneg    = false,
            duplex             = "auto",
            mtu                = 1514,
            no_local_overwrite = false,
            poe_disabled       = false,
            speed              = "auto",
            usage              = "dot1x-mab_only"
          },
          "ge-0/0/40-42, ge-0/0/46-47, xe-0/1/2-3" = {
            ae_lacp_slow       = true,
            aggregated         = false,
            critical           = false,
            disable_autoneg    = false,
            duplex             = "auto",
            mtu                = 1514,
            no_local_overwrite = true,
            poe_disabled       = false,
            speed              = "auto",
            usage              = "disabled_port"
          },
          "ge-0/0/44-45" = {
            ae_lacp_slow       = true,
            aggregated         = false,
            critical           = true,
            disable_autoneg    = false,
            duplex             = "auto",
            mtu                = 1514,
            no_local_overwrite = true,
            poe_disabled       = false,
            speed              = "auto",
            usage              = "store-ap"
          },
          "xe-0/1/0-1" = {
            ae_lacp_slow       = true,
            aggregated         = false,
            critical           = true,
            disable_autoneg    = false,
            duplex             = "auto",
            mtu                = 1514,
            no_local_overwrite = true,
            poe_disabled       = false,
            speed              = "auto",
            usage              = "uplink-switch"
          }
        }
      }
    ]
  }
  switch_mgmt = {
    ap_affinity_threshold = 10,
    config_revert_timer   = 10,
    dhcp_option_fqdn      = false,
    mxedge_proxy_port     = 2222,
    protect_re = {
      allowed_services = [
        "icmp",
        "ssh"
      ],
      enabled = true
    },
    root_password = "5T156XC0Ou"
  }
}

############# WXRULES
resource "mist_site_wxrule" "wxrule_wxrule_0" {
  action           = "allow"
  dst_allow_wxtags = []
  dst_deny_wxtags = [
    mist_site_wxtag.wxtag_4.id
  ]
  dst_wxtags = []
  enabled    = true
  src_wxtags = [
    mist_site_wxtag.wxtag_1.id
  ]
  site_id = mist_site.terraform_site.id
  order   = 1
}
############# WXTAGS
resource "mist_site_wxtag" "wxtag_0" {
  op = "in"
  values = [
    "8.8.8.8/32"
  ]
  name         = "google 8.8.8.8"
  site_id      = mist_site.terraform_site.id
  type         = "match"
  match        = "ip_range_subnet"
  mac          = null
}
resource "mist_site_wxtag" "wxtag_1" {
  op = "in"
  values = [
    mist_site_wlan.wlan_0.id
  ]
  name         = "Free-Wi-Fi"
  site_id      = mist_site.terraform_site.id
  type         = "match"
  match        = "wlan_id"
  mac          = null
}
resource "mist_site_wxtag" "wxtag_2" {
  op = "in"
  values = [
    "IoT"
  ]
  name         = "IoT"
  site_id      = mist_site.terraform_site.id
  type         = "match"
  match        = "radius_group"
  mac          = null
}
resource "mist_site_wxtag" "wxtag_3" {
  op = "in"
  values = [
    "Employee"
  ]
  name         = "Employee"
  site_id      = mist_site.terraform_site.id
  type         = "match"
  match        = "radius_group"
  mac          = null
}
resource "mist_site_wxtag" "wxtag_4" {
  op = "in"
  values = [
    "10.0.0.0/8",
    "172.16.0.0/12",
    "192.168.0.0/16"
  ]
  name         = "RFC-1918"
  site_id      = mist_site.terraform_site.id
  type         = "match"
  match        = "ip_range_subnet"
  mac          = null
}

############# WLAN

resource "mist_site_wlan" "wlan_0" {
  acct_immediate_update = false
  acct_interim_interval = 0
  allow_ipv6_ndp        = false
  allow_mdns            = false
  allow_ssdp            = false
  apply_to              = "site"
  arp_filter            = true
  auth = {
    anticlog_threshold    = 16
    eap_reauth            = false
    enable_mac_auth       = false
    key_idx               = 1
    multi_psk_only        = false
    owe                   = "disabled"
    psk                   = ""
    type                  = "open"
    wep_as_secondary_auth = false
  }
  auth_server_selection                  = "ordered"
  auth_servers_retries                   = 2
  auth_servers_timeout                   = 5
  band_steer                             = false
  band_steer_force_band5                 = false
  bands                                  = ["24", "5"]
  block_blacklist_clients                = true
  client_limit_down                      = 5000
  client_limit_down_enabled              = true
  client_limit_up                        = 5000
  client_limit_up_enabled                = true
  disable_11ax                           = false
  disable_ht_vht_rates                   = false
  disable_uapsd                          = false
  disable_v1_roam_notify                 = false
  disable_v2_roam_notify                 = false
  disable_wmm                            = false
  dtim                                   = 2
  enable_local_keycaching                = false
  enable_wireless_bridging               = false
  enable_wireless_bridging_dhcp_tracking = false
  enabled                                = true
  fast_dot1x_timers                      = false
  hide_ssid                              = false
  hostname_ie                            = false
  isolation                              = true
  l2_isolation                           = true
  legacy_overds                          = false
  limit_bcast                            = true
  limit_probe_response                   = false
  max_idletime                           = 1800
  mist_nac                               = { enabled = false }
  no_static_dns                          = false
  no_static_ip                           = false
  portal = {
    amazon_client_id               = ""
    amazon_client_secret           = ""
    amazon_enabled                 = false
    amazon_expire                  = 0
    auth                           = "none"
    azure_client_id                = ""
    azure_client_secret            = ""
    azure_enabled                  = false
    azure_expire                   = 0
    azure_tenant_id                = ""
    broadnet_password              = ""
    broadnet_sid                   = "MIST"
    broadnet_user_id               = ""
    bypass_when_cloud_down         = true
    clickatell_api_key             = ""
    cross_site                     = false
    email_enabled                  = false
    enabled                        = true
    expire                         = 240
    external_portal_url            = ""
    facebook_client_id             = ""
    facebook_client_secret         = ""
    facebook_enabled               = false
    facebook_expire                = 0
    forward                        = true
    forward_url                    = "https://www.mcdonalds.com/us/en-us/wifi.html"
    google_client_id               = ""
    google_client_secret           = ""
    google_enabled                 = false
    google_expire                  = 0
    gupshup_password               = ""
    gupshup_userid                 = ""
    microsoft_client_id            = ""
    microsoft_client_secret        = ""
    microsoft_enabled              = false
    microsoft_expire               = 0
    passphrase_enabled             = false
    passphrase_expire              = 0
    password                       = ""
    predefined_sponsors_enabled    = false
    privacy                        = false
    puzzel_password                = ""
    puzzel_service_id              = ""
    puzzel_username                = ""
    smsMessageFormat               = "Code {{code}} expires in {{duration}} minutes."
    sms_enabled                    = false
    sms_expire                     = 0
    sms_provider                   = "manual"
    sponsor_auto_approve           = false
    sponsor_enabled                = false
    sponsor_expire                 = 0
    sponsor_link_validity_duration = "60"
    sponsor_notify_all             = true
    sponsor_status_notify          = false
    sponsors                       = {}
    sso_default_role               = ""
    sso_forced_role                = ""
    sso_idp_cert                   = ""
    sso_idp_sign_algo              = "sha1"
    sso_idp_sso_url                = ""
    sso_issuer                     = ""
    sso_nameid_format              = "email"
    telstra_client_id              = ""
    telstra_client_secret          = ""
    twilio_auth_token              = ""
    twilio_phone_number            = ""
    twilio_sid                     = ""
  }
  roam_mode               = "NONE"
  sle_excluded            = true
  ssid                    = "McDonalds Free WiFi"
  use_eapol_v1            = false
  vlan_enabled            = true
  vlan_id                 = "450"
  vlan_pooling            = false
  wlan_limit_down         = 100000
  wlan_limit_down_enabled = true
  wlan_limit_up           = 100000
  wlan_limit_up_enabled   = true
  wxtunnel_remote_id      = ""
  site_id                 = mist_site.terraform_site.id
  interface               = "all"
}
resource "mist_site_wlan" "wlan_1" {
  acct_immediate_update = false
  acct_interim_interval = 0
  allow_ipv6_ndp        = false
  allow_mdns            = false
  allow_ssdp            = false
  apply_to              = "site"
  arp_filter            = true
  auth = {
    anticlog_threshold    = 16
    eap_reauth            = false
    enable_mac_auth       = true
    key_idx               = 1
    multi_psk_only        = false
    owe                   = "disabled"
    psk                   = ""
    type                  = "open"
    wep_as_secondary_auth = false
  }
  auth_server_selection                  = "ordered"
  auth_servers_retries                   = 2
  auth_servers_timeout                   = 5
  band_steer                             = false
  band_steer_force_band5                 = false
  bands                                  = ["24", "5"]
  block_blacklist_clients                = true
  client_limit_down_enabled              = false
  client_limit_up_enabled                = false
  disable_11ax                           = false
  disable_ht_vht_rates                   = false
  disable_uapsd                          = false
  disable_v1_roam_notify                 = false
  disable_v2_roam_notify                 = false
  disable_wmm                            = false
  dtim                                   = 2
  enable_local_keycaching                = false
  enable_wireless_bridging               = false
  enable_wireless_bridging_dhcp_tracking = false
  enabled                                = true
  fast_dot1x_timers                      = false
  hide_ssid                              = false
  hostname_ie                            = false
  isolation                              = false
  l2_isolation                           = false
  legacy_overds                          = false
  limit_bcast                            = true
  limit_probe_response                   = false
  max_idletime                           = 1800
  mist_nac                               = { enabled = true }
  no_static_dns                          = false
  no_static_ip                           = false
  portal = {
    amazon_client_id               = ""
    amazon_client_secret           = ""
    amazon_enabled                 = false
    amazon_expire                  = 0
    auth                           = "none"
    azure_client_id                = ""
    azure_client_secret            = ""
    azure_enabled                  = false
    azure_expire                   = 0
    azure_tenant_id                = ""
    broadnet_password              = ""
    broadnet_sid                   = "MIST"
    broadnet_user_id               = ""
    bypass_when_cloud_down         = true
    clickatell_api_key             = ""
    cross_site                     = false
    email_enabled                  = false
    enabled                        = false
    expire                         = 1440
    external_portal_url            = ""
    facebook_client_id             = ""
    facebook_client_secret         = ""
    facebook_enabled               = false
    facebook_expire                = 0
    forward                        = false
    forward_url                    = ""
    google_client_id               = ""
    google_client_secret           = ""
    google_enabled                 = false
    google_expire                  = 0
    gupshup_password               = ""
    gupshup_userid                 = ""
    microsoft_client_id            = ""
    microsoft_client_secret        = ""
    microsoft_enabled              = false
    microsoft_expire               = 0
    passphrase_enabled             = false
    passphrase_expire              = 0
    password                       = ""
    predefined_sponsors_enabled    = false
    privacy                        = false
    puzzel_password                = ""
    puzzel_service_id              = ""
    puzzel_username                = ""
    smsMessageFormat               = "Code {{code}} expires in {{duration}} minutes."
    sms_enabled                    = false
    sms_expire                     = 0
    sms_provider                   = "manual"
    sponsor_auto_approve           = false
    sponsor_enabled                = false
    sponsor_expire                 = 0
    sponsor_link_validity_duration = "60"
    sponsor_notify_all             = true
    sponsor_status_notify          = false
    sponsors                       = {}
    sso_default_role               = ""
    sso_forced_role                = ""
    sso_idp_cert                   = ""
    sso_idp_sign_algo              = "sha1"
    sso_idp_sso_url                = ""
    sso_issuer                     = ""
    sso_nameid_format              = "email"
    telstra_client_id              = ""
    telstra_client_secret          = ""
    twilio_auth_token              = ""
    twilio_phone_number            = ""
    twilio_sid                     = ""
  }
  roam_mode               = "NONE"
  sle_excluded            = false
  ssid                    = "eBOS"
  use_eapol_v1            = false
  vlan_enabled            = true
  vlan_id                 = "451"
  vlan_pooling            = false
  wlan_limit_down         = 20000
  wlan_limit_down_enabled = false
  wlan_limit_up         = 10000
  wlan_limit_up_enabled   = false
  wxtunnel_remote_id      = ""
  site_id                 = mist_site.terraform_site.id
  wxtunnel_id             = null
  interface               = "all"
}
resource "mist_site_wlan" "wlan_2" {
  acct_immediate_update = false
  acct_interim_interval = 0
  allow_ipv6_ndp        = false
  allow_mdns            = false
  allow_ssdp            = false
  arp_filter            = true
  auth = {
    anticlog_threshold    = 16
    eap_reauth            = false
    enable_mac_auth       = false
    key_idx               = 1
    multi_psk_only        = false
    owe                   = "disabled"
    psk                   = ""
    type                  = "eap"
    wep_as_secondary_auth = false
  }
  auth_server_selection                  = "ordered"
  auth_servers_retries                   = 2
  auth_servers_timeout                   = 5
  band_steer                             = false
  band_steer_force_band5                 = false
  bands                                  = ["24", "5"]
  block_blacklist_clients                = true
  client_limit_down_enabled              = false
  client_limit_up_enabled                = false
  disable_11ax                           = false
  disable_ht_vht_rates                   = false
  disable_uapsd                          = false
  disable_v1_roam_notify                 = false
  disable_v2_roam_notify                 = false
  disable_wmm                            = false
  dtim                                   = 2
  enable_local_keycaching                = false
  enable_wireless_bridging               = false
  enable_wireless_bridging_dhcp_tracking = false
  enabled                                = true
  fast_dot1x_timers                      = false
  hide_ssid                              = true
  hostname_ie                            = false
  isolation                              = false
  l2_isolation                           = false
  legacy_overds                          = false
  limit_bcast                            = true
  limit_probe_response                   = false
  max_idletime                           = 1800
  mist_nac                               = { enabled = true }
  no_static_dns                          = false
  no_static_ip                           = false
  portal = {
    amazon_client_id               = ""
    amazon_client_secret           = ""
    amazon_enabled                 = false
    amazon_expire                  = 0
    auth                           = "none"
    azure_client_id                = ""
    azure_client_secret            = ""
    azure_enabled                  = false
    azure_expire                   = 0
    azure_tenant_id                = ""
    broadnet_password              = ""
    broadnet_sid                   = "MIST"
    broadnet_user_id               = ""
    bypass_when_cloud_down         = false
    clickatell_api_key             = ""
    cross_site                     = false
    email_enabled                  = false
    enabled                        = false
    expire                         = 1440
    external_portal_url            = ""
    facebook_client_id             = ""
    facebook_client_secret         = ""
    facebook_enabled               = false
    facebook_expire                = 0
    forward                        = false
    forward_url                    = ""
    google_client_id               = ""
    google_client_secret           = ""
    google_enabled                 = false
    google_expire                  = 0
    gupshup_password               = ""
    gupshup_userid                 = ""
    microsoft_client_id            = ""
    microsoft_client_secret        = ""
    microsoft_enabled              = false
    microsoft_expire               = 0
    passphrase_enabled             = false
    passphrase_expire              = 0
    password                       = ""
    predefined_sponsors_enabled    = true
    privacy                        = false
    puzzel_password                = ""
    puzzel_service_id              = ""
    puzzel_username                = ""
    smsMessageFormat               = "Code {{code}} expires in {{duration}} minutes."
    sms_enabled                    = false
    sms_expire                     = 0
    sms_provider                   = "manual"
    sponsor_auto_approve           = false
    sponsor_enabled                = false
    sponsor_expire                 = 0
    sponsor_link_validity_duration = "60"
    sponsor_notify_all             = false
    sponsor_status_notify          = false
    sponsors                       = {}
    sso_default_role               = ""
    sso_forced_role                = ""
    sso_idp_cert                   = ""
    sso_idp_sign_algo              = "sha1"
    sso_idp_sso_url                = ""
    sso_issuer                     = ""
    sso_nameid_format              = "email"
    telstra_client_id              = ""
    telstra_client_secret          = ""
    twilio_auth_token              = ""
    twilio_phone_number            = ""
    twilio_sid                     = ""
  }
  roam_mode               = "NONE"
  sle_excluded            = false
  ssid                    = "McD-HHOT"
  use_eapol_v1            = false
  vlan_enabled            = true
  vlan_id                 = "420"
  vlan_pooling            = false
  wlan_limit_down         = 20000
  wlan_limit_down_enabled = false
  wlan_limit_up         = 10000
  wlan_limit_up_enabled   = false
  wxtunnel_remote_id      = ""
  site_id                 = mist_site.terraform_site.id
  wxtunnel_id             = null
  interface               = "all"
}
resource "mist_site_wlan" "wlan_3" {
  acct_immediate_update = false
  acct_interim_interval = 0
  allow_ipv6_ndp        = false
  allow_mdns            = false
  allow_ssdp            = false
  apply_to              = "site"
  arp_filter            = true
  auth = {
    anticlog_threshold    = 16
    eap_reauth            = false
    enable_mac_auth       = false
    key_idx               = 1
    multi_psk_only        = true
    owe                   = "disabled"
    psk                   = ""
    type                  = "psk"
    wep_as_secondary_auth = false
  }
  auth_server_selection                  = "ordered"
  auth_servers_retries                   = 2
  auth_servers_timeout                   = 5
  band_steer                             = false
  band_steer_force_band5                 = false
  bands                                  = ["24", "5"]
  block_blacklist_clients                = true
  client_limit_down_enabled              = false
  client_limit_up_enabled                = false
  disable_11ax                           = false
  disable_ht_vht_rates                   = false
  disable_uapsd                          = false
  disable_v1_roam_notify                 = false
  disable_v2_roam_notify                 = false
  disable_wmm                            = false
  dtim                                   = 2
  enable_local_keycaching                = false
  enable_wireless_bridging               = false
  enable_wireless_bridging_dhcp_tracking = false
  enabled                                = true
  fast_dot1x_timers                      = false
  hide_ssid                              = false
  hostname_ie                            = false
  isolation                              = true
  l2_isolation                           = true
  legacy_overds                          = false
  limit_bcast                            = true
  limit_probe_response                   = false
  max_idletime                           = 1800
  mist_nac = {
    enabled = true
  }
  no_static_dns = false
  no_static_ip  = false
  portal = {
    amazon_client_id               = ""
    amazon_client_secret           = ""
    amazon_enabled                 = false
    amazon_expire                  = 0
    auth                           = "none"
    azure_client_id                = ""
    azure_client_secret            = ""
    azure_enabled                  = false
    azure_expire                   = 0
    azure_tenant_id                = ""
    broadnet_password              = ""
    broadnet_sid                   = "MIST"
    broadnet_user_id               = ""
    bypass_when_cloud_down         = true
    clickatell_api_key             = ""
    cross_site                     = false
    email_enabled                  = false
    enabled                        = false
    expire                         = 1440
    external_portal_url            = ""
    facebook_client_id             = ""
    facebook_client_secret         = ""
    facebook_enabled               = false
    facebook_expire                = 0
    forward                        = false
    forward_url                    = ""
    google_client_id               = ""
    google_client_secret           = ""
    google_enabled                 = false
    google_expire                  = 0
    gupshup_password               = ""
    gupshup_userid                 = ""
    microsoft_client_id            = ""
    microsoft_client_secret        = ""
    microsoft_enabled              = false
    microsoft_expire               = 0
    passphrase_enabled             = false
    passphrase_expire              = 0
    password                       = ""
    predefined_sponsors_enabled    = false
    privacy                        = false
    puzzel_password                = ""
    puzzel_service_id              = ""
    puzzel_username                = ""
    smsMessageFormat               = "Code {{code}} expires in {{duration}} minutes."
    sms_enabled                    = false
    sms_expire                     = 0
    sms_provider                   = "manual"
    sponsor_auto_approve           = false
    sponsor_enabled                = false
    sponsor_expire                 = 0
    sponsor_link_validity_duration = "60"
    sponsor_notify_all             = true
    sponsor_status_notify          = false
    sponsors                       = {}
    sso_default_role               = ""
    sso_forced_role                = ""
    sso_idp_cert                   = ""
    sso_idp_sign_algo              = "sha1"
    sso_idp_sso_url                = ""
    sso_issuer                     = ""
    sso_nameid_format              = "email"
    telstra_client_id              = ""
    telstra_client_secret          = ""
    twilio_auth_token              = ""
    twilio_phone_number            = ""
    twilio_sid                     = ""
  }
  roam_mode               = "NONE"
  sle_excluded            = false
  ssid                    = "McD-IoT"
  use_eapol_v1            = false
  vlan_enabled            = true
  vlan_id                 = "494"
  vlan_pooling            = false
  wlan_limit_down         = 20000
  wlan_limit_down_enabled = false
  wlan_limit_up         = 10000
  wlan_limit_up_enabled   = false
  wxtunnel_remote_id      = ""
  site_id                 = mist_site.terraform_site.id
  wxtunnel_id             = null
  interface               = "all"
}
############# SERVICES

resource "mist_org_service" "service_0" {
  apps = [
    "office365"
  ]
  failover_policy                   = "revertable"
  sle_enabled                       = false
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "LBO-L7-APPS"
  org_id                            = mist_org.terraform_test.id
  type                              = "apps"
}
resource "mist_org_service" "service_1" {
  apps = [
    "office365"
  ]
  description                       = "App for detecting office 365"
  failover_policy                   = "revertable"
  sle_enabled                       = false
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "XX-office365"
  org_id                            = mist_org.terraform_test.id
  type                              = "apps"
}
resource "mist_org_service" "service_2" {
  apps = [
    "speedtest"
  ]
  description                       = "to demonstrate traffic shaping"
  failover_policy                   = "revertable"
  sle_enabled                       = false
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "voip_audio"
  name                              = "XX-Speedtest-classification"
  org_id                            = mist_org.terraform_test.id
  type                              = "apps"
}
resource "mist_org_service" "service_3" {
  addresses = [
    "{{branch_ip}}.128/27"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "8123-8123"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_128-27-TCP_8123"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_4" {
  addresses = [
    "{{iot_net_wired}}.50/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "icmp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_50-ICMP"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_5" {
  addresses = [
    "{{branch_ip}}.132/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "5016-5016"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_132-TCP_5016"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_6" {
  addresses = [
    "{{iot_net_wireless}}.0/23"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "22-22"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_94_0-23-TCP_22"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_7" {
  addresses = [
    "{{iot_net_wired}}.200/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "22-22"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_200-TCP_22"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_8" {
  addresses = [
    "{{branch_ip}}.160/27"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "3333-3333"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_160-27-TCP_3333"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_9" {
  addresses = [
    "{{branch_no_routable}}.130/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "2000-2000"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_90_130-TCP_2000"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_10" {
  addresses = [
    "0.0.0.0/0"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "8112-8112"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "ANY-TCP_8112"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_11" {
  addresses = [
    "{{iot_net_wired}}.0/26"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "80-80"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_0-24-TCP_80"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_12" {
  addresses = [
    "{{iot_net_wired_upper}}.0/24"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "8883-8883"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_93_0-24-TCP_8883"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_13" {
  addresses = [
    "{{iot_net_wired_upper}}.0/24"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "2222-2222"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_93_0-24-UDP_2222"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_14" {
  addresses = [
    "{{iot_net_wired_upper}}.0/24"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "443-443"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_93_0-24-TCP_443"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_15" {
  addresses = [
    "{{iot_net_wired}}.126/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "123-123"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_1-UDP_123"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_16" {
  addresses = [
    "{{iot_net_wired}}.168/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "10000-10000"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_168-TCP_10000"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_17" {
  addresses = [
    "{{branch_ip}}.160/27"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "22500-22500"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_160-27-TCP_22500"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_18" {
  addresses = [
    "{{iot_net_wired}}.0/24"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "44818-44818"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_0-24-TCP_44818"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_19" {
  addresses = [
    "{{iot_net_wired}}.0/24"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "8883-8883"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_0-24-TCP_8883"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_20" {
  addresses = [
    "{{iot_net_wired_upper}}.0/24"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "1883-1883"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_93_0-24-TCP_1883"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_21" {
  addresses = [
    "0.0.0.0/0"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "65535-65535"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "ANY-UDP_65535"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_22" {
  addresses = [
    "{{branch_ip}}.160/27"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "225-225"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_160-27-TCP_225"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_23" {
  app_categories = [
    "Adult"
  ]
  description                       = "Block the following categories Enhanced_Adult_Material Enhanced_Nudity Enhanced_Adult_Content Enhanced_Sex Enhanced_Sex_Education Enhanced_Lingerie_and_Swimsuit"
  failover_policy                   = "revertable"
  sle_enabled                       = false
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "public_url_block"
  org_id                            = mist_org.terraform_test.id
  type                              = "app_categories"
}
resource "mist_org_service" "service_24" {
  addresses = [
    "{{branch_ip}}.129/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "123-123"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_129-UDP_123"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_25" {
  addresses = [
    "{{iot_net_wired_upper}}.0/24"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "23-23"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_93_0-24-TCP_23"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_26" {
  addresses = [
    "{{VLAN430}}.160/27"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "any"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "MANAGEMENT_VLAN430"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_27" {
  addresses = [
    "{{branch_ip}}.17/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "5172-5172"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_17-TCP_5172"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_28" {
  addresses = [
    "{{branch_ip}}.160/27"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "22600-22600"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_160-27-TCP_22600"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_29" {
  addresses = [
    "0.0.0.0/0"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "443-443"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "ANY-TCP_443"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_30" {
  addresses = [
    "{{branch_ip}}.0/24"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "any"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_0-24-ANY"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_31" {
  addresses = [
    "0.0.0.0/0"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "80-80"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "ANY-TCP_80"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_32" {
  addresses = [
    "{{branch_no_routable}}.130/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "443-443"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_90_130-TCP_443"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_33" {
  addresses = [
    "{{branch_no_routable}}.130/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "20000-20000"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_90_130-TCP_20000"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_34" {
  addresses = [
    "{{branch_ip}}.172/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "225-225"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_172-TCP_225"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_35" {
  addresses = [
    "{{branch_ip}}.225/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "123-123"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_225-UDP_123"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_36" {
  addresses = [
    "{{branch_ip}}.131/32",
    "{{branch_ip}}.132/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "5016-5016"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "RESTRICTED-TCP-5016"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_37" {
  addresses = [
    "{{branch_ip}}.171/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "225-225"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_171-TCP_225"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_38" {
  addresses = [
    "{{branch_ip}}.51/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "icmp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_51-ICMP"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_39" {
  addresses = [
    "{{iot_net_wired_upper}}.0/26"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "44818-44818"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_93_0-24-TCP_44818"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_40" {
  addresses = [
    "{{branch_ip}}.131/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "5016-5016"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_131-TCP_5016"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_41" {
  addresses = [
    "{{branch_ip}}.17/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "5172-5172"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_18-TCP-5172"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_42" {
  addresses = [
    "{{branch_no_routable}}.130/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "2002-2002"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_90_130-TCP_2002"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_43" {
  addresses = [
    "{{branch_ip}}.25/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "123-123"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_25-UDP_123"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_44" {
  addresses = [
    "{{branch_ip}}.132/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "5015-5015"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_132-TCP_5015"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_45" {
  addresses = [
    "0.0.0.0/0"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "80-80"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "ANY_TCP_80"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_46" {
  addresses = [
    "0.0.0.0/0"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "67-67"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "ANY-UDP_67"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_47" {
  addresses = [
    "{{iot_net_wired_upper}}.0/24"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "80-80"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_93_0-24-TCP_80"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_48" {
  addresses = [
    "0.0.0.0/0"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "1163-1163"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "ANY-TCP_1163"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_49" {
  addresses = [
    "{{branch_ip}}.131/32\n{{branch_ip}}.132/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "5015-5015"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "CORE-TCP-5015"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_50" {
  addresses = [
    "0.0.0.0/0"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "8211-8211"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "ANY-UDP_8211"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_51" {
  addresses = [
    "{{branch_ip}}.160/27"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "22700-22700"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_160-27-TCP_22700"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_52" {
  addresses = [
    "0.0.0.0/0"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "123-123"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "NTP"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_53" {
  addresses = [
    "{{branch_ip}}.160/27"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "any"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_160-27-ANY"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_54" {
  addresses = [
    "192.168.0.0/16"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "icmp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_0_0-16-ICMP"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_55" {
  addresses = [
    "{{private_net}}.0/24"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "any"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_91_0-24-ANY"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_56" {
  addresses = [
    "{{iot_net_wired}}.198/32",
    "{{secondary_iot_wired}}.198/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "1883-1883"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_198-TCP_1883"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_57" {
  addresses = [
    "208.67.220.220/32",
    "208.67.222.222/32"
  ]
  description     = "Cisco Umbrella - Puertos TCP UDP for DNS and HTTPS"
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "53-53"
      protocol   = "tcp"
    },
    {
      port_range = "53-53"
      protocol   = "udp"
    },
    {
      port_range = "80-80"
      protocol   = "tcp"
    },
    {
      port_range = "443-443"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "DNS_HTTP_HTTPS_External"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_58" {
  addresses = [
    "{{iot_net_wired}}.160/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "443-443"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_160-TCP_443"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_59" {
  addresses = [
    "{{branch_ip}}.62/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "53-53"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_62-UDP_53"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_60" {
  addresses = [
    "{{branch_ip}}.129/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "53-53"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_129-UDP_53"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_61" {
  addresses = [
    "{{iot_net_wired}}.126/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "53-53"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_1-UDP_53"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_62" {
  addresses = [
    "{{iot_net_wired}}.0/24"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "2222-2222"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_0-24-UDP_2222"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_63" {
  addresses = [
    "{{iot_net_wired}}.199/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "443-443"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_199-TCP_443"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_64" {
  addresses = [
    "{{branch_no_routable}}.130/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "80-80"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_90_130-TCP_80"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_65" {
  addresses = [
    "0.0.0.0/0"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "68-68"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "ANY-UDP_68"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_66" {
  addresses = [
    "0.0.0.0/0"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "22-22"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "ANY-TCP_22"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_67" {
  addresses = [
    "{{iot_net_wired}}.0/24"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "443-443"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_0-24-TCP_443"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_68" {
  addresses = [
    "{{branch_ip}}.161/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "53-53"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_161-UDP_53"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_69" {
  addresses = [
    "{{iot_net_wired}}.198/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "22-22"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_198-TCP_22"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_70" {
  addresses = [
    "{{branch_ip}}.0/25"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "22501-22501"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_0-25-TCP_22501"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_71" {
  addresses = [
    "{{branch_ip}}.160/27"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "22501-22501"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_160-27-TCP_22501"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_72" {
  addresses = [
    "{{private_net}}.1/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "123-123"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_91_1-UDP_123"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_73" {
  addresses = [
    "{{branch_ip}}.25/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "5173-5173"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_25-TCP_5173"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_74" {
  addresses = [
    "{{branch_ip}}.160/27"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "80-80"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_160-27-TCP_80"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_75" {
  addresses = [
    "{{branch_ip}}.25/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "8123-8123"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_25-TCP_8123"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_76" {
  addresses = [
    "{{branch_ip}}.0/25"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "any"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_0-25-ANY"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_77" {
  addresses = [
    "{{iot_net_wired}}.160/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "3255-3255"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_160-TCP_3255"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_78" {
  addresses = [
    "{{iot_net_wired}}.198/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "8883-8883"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_198-TCP_8883"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_79" {
  addresses = [
    "{{branch_ip}}.0/25"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "22500-22500"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_0-25-TCP_22500"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_80" {
  addresses = [
    "172.16.0.0/16"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "any"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "172_16_0_0-16-ANY"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_81" {
  addresses = [
    "{{iot_net_wired}}.198/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "443-443"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_198-TCP_443"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_82" {
  addresses = [
    "{{branch_ip}}.168/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "225-225"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_168-TCP_225"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_83" {
  addresses = [
    "{{branch_ip}}.225/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "53-53"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_225-UDP_53"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_84" {
  addresses = [
    "{{branch_ip}}.25/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "8883-8883"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_25-TCP_8883"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_85" {
  addresses = [
    "{{branch_ip}}.131/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "5015-5015"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_131-TCP_5015"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_86" {
  addresses = [
    "{{branch_ip}}.25/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "1883-1883"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_25-TCP_1883"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_87" {
  addresses = [
    "{{crew_network}}.0/24"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "any"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_6_0-24-ANY"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_88" {
  addresses = [
    "{{branch_ip}}.160/27"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "8888-8888"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_160-27-TCP_8888"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_89" {
  addresses = [
    "{{branch_ip}}.25/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "icmp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_25-ICMP"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_90" {
  addresses = [
    "{{branch_ip}}.51/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "123-123"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_51-UDP_123"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_91" {
  addresses = [
    "{{iot_net_wired}}.168/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "65535-65535"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_168-UDP_65535"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_92" {
  addresses = [
    "{{iot_net_wired}}.168/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "10001-10001"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_168-TCP_10001"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_93" {
  addresses = [
    "{{branch_no_routable}}.129/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "icmp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_90_129-ICMP"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_94" {
  addresses = [
    "{{branch_no_routable}}.130/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "20000-20000"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_90_130-UDP_20000"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_95" {
  addresses = [
    "{{iot_net_wired}}.166/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "80-80"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_166-TCP_80"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_96" {
  addresses = [
    "{{branch_ip}}.1/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "8123-8123"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_1-TCP_8123"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_97" {
  addresses = [
    "{{branch_ip}}.160/27"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "4983-4983"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_160-27-TCP_4983"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_98" {
  addresses = [
    "{{iot_net_wired}}.0/24"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "502-502"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_0-24-TCP_502"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_99" {
  addresses = [
    "192.168.0.0/16"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "any"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_0_0-16-ANY"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_100" {
  addresses = [
    "{{branch_ip}}.50/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "icmp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_50-ICMP"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_101" {
  addresses = [
    "{{branch_ip}}.62/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "123-123"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_62-UDP_123"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_102" {
  addresses = [
    "0.0.0.0/0"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "8209-8209"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "ANY-UDP_8209"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_103" {
  addresses = [
    "{{iot_net_wired}}.200/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "443-443"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_200-TCP_443"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_104" {
  addresses = [
    "{{branch_no_routable}}.0/26"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "any"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_90_0-24-ANY"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_105" {
  addresses = [
    "{{branch_ip}}.131/32",
    "{{branch_ip}}.132/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "5015-5015"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "RESTRICTED-TCP-5015"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_106" {
  addresses = [
    "{{branch_ip}}.1/32",
    "{{branch_ip}}.25/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "8123-8123"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "CORE_1_25-TCP-8123"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_107" {
  addresses = [
    "{{iot_net_wired}}.192/26",
    "{{iot_net_wired}}.180/32",
    "{{iot_net_wired}}.181/32",
    "{{iot_net_wired}}.182/32",
    "{{iot_net_wired}}.183/32",
    "{{iot_net_wired}}.184/32",
    "{{iot_net_wired}}.185/32",
    "{{iot_net_wired}}.186/32",
    "{{iot_net_wired}}.187/32",
    "{{iot_net_wired}}.188/32",
    "{{iot_net_wired}}.189/32",
    "{{iot_net_wired}}.190/32",
    "{{iot_net_wired}}.191/32",
    "{{secondary_iot_wired}}.192/26",
    "{{secondary_iot_wired}}.180/32",
    "{{secondary_iot_wired}}.181/32",
    "{{secondary_iot_wired}}.182/32",
    "{{secondary_iot_wired}}.183/32",
    "{{secondary_iot_wired}}.184/32",
    "{{secondary_iot_wired}}.185/32",
    "{{secondary_iot_wired}}.186/32",
    "{{secondary_iot_wired}}.187/32",
    "{{secondary_iot_wired}}.188/32",
    "{{secondary_iot_wired}}.189/32",
    "{{secondary_iot_wired}}.190/32",
    "{{secondary_iot_wired}}.191/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "22501-22501"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_192-26-TCP_22501"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_108" {
  addresses = [
    "{{iot_net_wired}}.199/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "8883-8883"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_199-TCP_8883"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_109" {
  addresses = [
    "{{iot_net_wireless}}.254/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "53-53"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_94_1-UDP_53"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_110" {
  addresses = [
    "{{iot_net_wired_upper}}.0/24"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "502-502"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_93_0-24-TCP_502"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_111" {
  addresses = [
    "{{private_net}}.1/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "53-53"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_91_1-UDP_53"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_112" {
  addresses = [
    "{{public_net}}.0/23"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "any"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_4_0-23-ANY"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_113" {
  addresses = [
    "{{iot_net_wired}}.199/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "22-22"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_199-TCP_22"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_114" {
  addresses = [
    "{{branch_ip}}.249/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "44818-44818"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_249-TCP_44818"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_115" {
  addresses = [
    "{{iot_net_wired}}.169/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "80-80"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_169-TCP_80"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_116" {
  addresses = [
    "{{VLAN410}}.0/25",
    "{{VLAN420}}.128/27",
    "{{VLAN430}}.160/27",
    "{{VLAN440}}.192/27",
    "{{VLAN490}}.128/26"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "icmp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "ICMP_VLAN410to490"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_117" {
  addresses = [
    "{{branch_ip}}.25/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "5172-5172"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_25-TCP_5172"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_118" {
  addresses = [
    "{{iot_net_wired}}.165/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "80-80"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_165-TCP_80"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_119" {
  addresses = [
    "{{iot_net_wired}}.0/24"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "23-23"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_0-24-TCP_23"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_120" {
  addresses = [
    "{{iot_net_wired}}.160/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "80-80"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_160-TCP_80"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_121" {
  addresses = [
    "{{iot_net_wireless}}.254/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "123-123"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_94_1-UDP_123"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_122" {
  addresses = [
    "{{iot_net_wired}}.200/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "8883-8883"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_200-TCP_8883"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_123" {
  addresses = [
    "{{iot_net_wired}}.54/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "44818-44818"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_54-TCP_44818"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_124" {
  addresses = [
    "{{branch_ip}}.17/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "5173-5173"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_17-TCP_5173"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_125" {
  addresses = [
    "{{branch_ip}}.160/27"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "icmp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_160-27-ICMP"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_126" {
  addresses = [
    "{{branch_ip}}.161/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "123-123"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_161-UDP_123"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_127" {
  addresses = [
    "{{branch_ip}}.169/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "225-225"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_169-TCP_225"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_128" {
  addresses = [
    "0.0.0.0/0"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "123-123"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "ANY-UDP_123"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_129" {
  addresses = [
    "{{branch_ip}}.18/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "5173-5173"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_X_Y_18-TCP_5173"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_130" {
  addresses = [
    "{{iot_net_wired}}.168/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "10002-10002"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_168-TCP_10002"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_131" {
  addresses = [
    "10.0.0.0/8"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "any"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "10_0_0_0-8-ANY"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_132" {
  addresses = [
    "{{iot_net_wired}}.168/32",
    "{{secondary_iot_wired}}.168/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "1883-1883"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_168-TCP_1883"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_133" {
  addresses = [
    "10.118.25.101/32"
  ]
  description     = "TEMPORARY - FOR CAPLAB ONLY"
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "any"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "CAPLAB-10_118_25_101-ANY"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_134" {
  addresses = [
    "0.0.0.0/0"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      protocol = "icmp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "ICMP_ANY"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_135" {
  failover_policy = "revertable"
  hostnames = [
    "github.com"
  ]
  sle_enabled = false
  specs = [
    {
      port_range = "80-80"
      protocol   = "tcp"
    },
    {
      port_range = "443-443"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "GitHub_TCP_80_443"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_136" {
  addresses = [
    "{{iot_net_wired}}.199/32",
    "{{secondary_iot_wired}}.199/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "1883-1883"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_199-TCP_1883"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_137" {
  addresses = [
    "{{branch_ip}}.17/32",
    "{{branch_ip}}.18/32",
    "{{branch_ip}}.25/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "5172-5172"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "CORE-TCP-5172"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_138" {
  addresses = [
    "{{iot_net_wired}}.0/24",
    "{{secondary_iot_wired}}.0/24"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "1883-1883"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_0-24-TCP_1883"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_139" {
  addresses = [
    "208.67.220.220/32",
    "208.67.222.222/32",
    "8.8.8.8/32",
    "1.1.1.1/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "53-53"
      protocol   = "tcp"
    },
    {
      port_range = "53-53"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "MANAGEMENT_VLAN430_DNS"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_140" {
  addresses = [
    "{{iot_net_wired}}.200/32",
    "{{secondary_iot_wired}}.200/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "1883-1883"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_200-TCP_1883"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_141" {
  addresses = [
    "{{iot_net_wired}}.160/32",
    "{{iot_net_wired}}.165/32",
    "{{iot_net_wired}}.166/32",
    "{{iot_net_wired}}.169/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "80-80"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "Gr_92-160_165_166_169-TCP_80"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_142" {
  addresses = [
    "40.122.111.34/32",
    "40.91.202.159/32"
  ]
  failover_policy = "revertable"
  hostnames = [
    "gas.mcd.com"
  ]
  sle_enabled = false
  specs = [
    {
      port_range = "80-80"
      protocol   = "tcp"
    },
    {
      port_range = "443-443"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "GAS_MCD_HTTP_HTTPS"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_143" {
  addresses = [
    "{{branch_ip}}.17/32",
    "{{branch_ip}}.18/32",
    "{{branch_ip}}.25/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "5173-5173"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "CORE-TCP-5173"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_144" {
  addresses = [
    "{{iot_net_wired}}.200/32",
    "{{iot_net_wired}}.199/32",
    "{{iot_net_wired}}.198/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "8883-8883"
      protocol   = "tcp"
    },
    {
      port_range = "1883-1883"
      protocol   = "tcp"
    },
    {
      port_range = "443-443"
      protocol   = "tcp"
    },
    {
      port_range = "22-22"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "Gr2_IoT_Wired"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_145" {
  addresses = [
    "0.0.0.0/0"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "80-80"
      protocol   = "tcp"
    },
    {
      port_range = "443-443"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "TCP_80_443"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_146" {
  addresses = [
    "{{iot_net_wired}}.160/32",
    "{{iot_net_wired}}.165/32",
    "{{iot_net_wired}}.166/32",
    "{{iot_net_wired}}.169/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "80-80"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "Gr_224-192_168_92-TCP_80"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_147" {
  addresses = [
    "{{iot_net_wired}}.192/26",
    "{{iot_net_wired}}.180/32",
    "{{iot_net_wired}}.181/32",
    "{{iot_net_wired}}.182/32",
    "{{iot_net_wired}}.183/32",
    "{{iot_net_wired}}.184/32",
    "{{iot_net_wired}}.185/32",
    "{{iot_net_wired}}.186/32",
    "{{iot_net_wired}}.187/32",
    "{{iot_net_wired}}.188/32",
    "{{iot_net_wired}}.189/32",
    "{{iot_net_wired}}.190/32",
    "{{iot_net_wired}}.191/32",
    "{{secondary_iot_wired}}.192/26",
    "{{secondary_iot_wired}}.180/32",
    "{{secondary_iot_wired}}.181/32",
    "{{secondary_iot_wired}}.182/32",
    "{{secondary_iot_wired}}.183/32",
    "{{secondary_iot_wired}}.184/32",
    "{{secondary_iot_wired}}.185/32",
    "{{secondary_iot_wired}}.186/32",
    "{{secondary_iot_wired}}.187/32",
    "{{secondary_iot_wired}}.188/32",
    "{{secondary_iot_wired}}.189/32",
    "{{secondary_iot_wired}}.190/32",
    "{{secondary_iot_wired}}.191/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "22500-22500"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "192_168_92_192-26-TCP_22500"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_148" {
  addresses = [
    "{{branch_ip}}.62/32",
    "{{branch_ip}}.129/32",
    "{{branch_ip}}.161/32",
    "{{branch_ip}}.225/32"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "53-53"
      protocol   = "udp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "Gr-10_X_Y-UDP_53"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}
resource "mist_org_service" "service_149" {
  addresses = [
    "0.0.0.0/0"
  ]
  failover_policy = "revertable"
  sle_enabled     = false
  specs = [
    {
      port_range = "53-53"
      protocol   = "udp"
    },
    {
      port_range = "53-53"
      protocol   = "tcp"
    }
  ]
  ssr_relaxed_tcp_state_enforcement = false
  traffic_class                     = "best_effort"
  traffic_type                      = "default"
  name                              = "ANY-UDP_53"
  org_id                            = mist_org.terraform_test.id
  type                              = "custom"
}


############# SERVICE POLICIES

resource "mist_org_servicepolicy" "servicepolicy_0" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_92_50-ICMP"
  ]
  tenants = [
    "CORE_25.CORE_VLAN410"
  ]
  name   = "022_410-492-bi-com-3-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_1" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_168-TCP_225"
  ]
  tenants = [
    "CORE_12.CORE_VLAN410"
  ]
  name   = "089_410-cod-6-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_2" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_128-27-TCP_8123"
  ]
  tenants = [
    "CORE_1.CORE_VLAN410",
    "CORE_25.CORE_VLAN410"
  ]
  name   = "071_410-420-bi-com-0-VLAN420"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_3" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "CORE_1_25-TCP-8123"
  ]
  tenants = [
    "RESTRICTED_VLAN420"
  ]
  name   = "018_410-420-bi-com-1-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_4" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_169-TCP_225"
  ]
  tenants = [
    "CORE_12.CORE_VLAN410"
  ]
  name   = "090_410-cod-7-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_5" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_94_1-UDP_53"
  ]
  tenants = [
    "PED_DEVICES-CORE-DNS.CORE_VLAN410",
    "PED_DEVICES-RESTRICTED-DNS.RESTRICTED_VLAN420"
  ]
  name   = "002f-ped-devices-rules-0-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_6" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_67",
    "ANY-UDP_68"
  ]
  tenants = [
    "PUBLIC_VLAN450"
  ]
  name   = "100_450-control-1-VLAN450"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_7" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_92_192-26-TCP_22501"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "025_410-492-bi-com-6-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_8" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_92_50-ICMP"
  ]
  tenants = [
    "CORE_51.CORE_VLAN410"
  ]
  name   = "023_410-492-bi-com-4-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_9" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_171-TCP_225"
  ]
  tenants = [
    "CORE_12.CORE_VLAN410"
  ]
  name   = "091_410-cod-8-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_10" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_25-ICMP"
  ]
  tenants = [
    "IOT_WIRED_50.IOT_WIRED_VLAN492"
  ]
  name   = "185_410-492-bi-comm-1-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_11" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_25-ICMP"
  ]
  tenants = [
    "IOT_WIRED_50.IOT_WIRED_VLAN492"
  ]
  name   = "020_410-492-bi-com-1-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_12" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_171-TCP_225"
  ]
  tenants = [
    "CORE_25.CORE_VLAN410"
  ]
  name   = "036_410-cod-4-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_13" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_25-UDP_123"
  ]
  tenants = [
    "IOT_WIRED_VLAN492"
  ]
  name   = "173_492-control-4-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_14" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-TCP_443"
  ]
  tenants = [
    "CORE_23.CORE_VLAN410",
    "CORE_24.CORE_VLAN410"
  ]
  name   = "016_qsrsoft-vm-rules-4-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_15" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "Gr2_IoT_Wired"
  ]
  tenants = [
    "IOT_WIRELESS_VLAN494"
  ]
  name   = "177_pk-test-494-492-r0-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_16" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_90_130-TCP_2002"
  ]
  tenants = [
    "CORE_51.CORE_VLAN410"
  ]
  name   = "063_livestore-sec-cam-3-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_17" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_171-TCP_225"
  ]
  tenants = [
    "CORE_12.CORE_VLAN410"
  ]
  name   = "040_410-cod-8-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_18" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_169-TCP_225"
  ]
  tenants = [
    "CORE_14.CORE_VLAN410"
  ]
  name   = "094_410-cod-11-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_19" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_161-UDP_53"
  ]
  tenants = [
    "IOT_WIRED_VLAN492"
  ]
  name   = "176c_492-control-7-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_20" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "Gr2_IoT_Wired"
  ]
  tenants = [
    "IOT_WIRELESS_VLAN494"
  ]
  name   = "178_pk-test-494-492-r1-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_21" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_92_192-26-TCP_22500"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "028_410-492-2331-0-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_22" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_51-UDP_123"
  ]
  tenants = [
    "IOT_WIRED_VLAN492"
  ]
  name   = "174_492-control-5-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_23" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "Gr2_IoT_Wired"
  ]
  tenants = [
    "IOT_WIRELESS_VLAN494"
  ]
  name   = "179_pk-test-494-492-r2-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_24" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_92_192-26-TCP_22501"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "029_410-492-2331-1-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_25" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_171-TCP_225"
  ]
  tenants = [
    "CORE_14.CORE_VLAN410"
  ]
  name   = "095_410-cod-12-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_26" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "Gr2_IoT_Wired"
  ]
  tenants = [
    "IOT_WIRELESS_VLAN494"
  ]
  name   = "180_pk-test-494-492-r3-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_27" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_0-25-TCP_22500"
  ]
  tenants = [
    "Gr_IoT_WIRED.IOT_WIRED_VLAN492"
  ]
  name   = "030_410-492-2331-2-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_28" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_62-UDP_123"
  ]
  tenants = [
    "IOT_WIRED_VLAN492"
  ]
  name   = "175a_492-control-6-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_29" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_90_129-ICMP"
  ]
  tenants = [
    "UNTRUSTED_VLAN490"
  ]
  name   = "150_control-vlan-490-2-VLAN490"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_30" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_0-25-TCP_22501"
  ]
  tenants = [
    "Gr_IoT_WIRED.IOT_WIRED_VLAN492"
  ]
  name   = "031_410-492-2331-3-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_31" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_169-TCP_225"
  ]
  tenants = [
    "CORE_12.CORE_VLAN410"
  ]
  name   = "039_410-cod-7-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_32" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_172-TCP_225"
  ]
  tenants = [
    "CORE_14.CORE_VLAN410"
  ]
  name   = "096_410-cod-13-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_33" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_90_130-TCP_80"
  ]
  tenants = [
    "CORE_51.CORE_VLAN410"
  ]
  name   = "151_livestore-sec-cam-0-VLAN490"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_34" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-TCP_225"
  ]
  tenants = [
    "Gr_CORE_12_14_25.CORE_VLAN410"
  ]
  name   = "032_410-cod-0-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_35" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_90_130-TCP_20000"
  ]
  tenants = [
    "CORE_51.CORE_VLAN410"
  ]
  name   = "064_livestore-sec-cam-4-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_36" {
  action = "deny"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_8209"
  ]
  tenants = [
    "PUBLIC_VLAN450"
  ]
  name   = "097_papi-block-0-VLAN450"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_37" {
  action = "deny"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_8209"
  ]
  tenants = [
    "CREW_VLAN451"
  ]
  name   = "104_papi-block-0-VLAN451"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_38" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "any"
  ]
  tenants = [
    "CREW_VLAN451"
  ]
  name   = "112_int-src-nat-pri4-0-VLAN451"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_39" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-TCP_4983"
  ]
  tenants = [
    "Gr_CORE_12_14_25.CORE_VLAN410"
  ]
  name   = "033_410-cod-1-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_40" {
  action = "deny"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_0_0-16-ANY"
  ]
  tenants = [
    "UNTRUSTED_VLAN490"
  ]
  name   = "157_deny-rfc-privateip-0-VLAN490"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_41" {
  action = "deny"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_8211"
  ]
  tenants = [
    "PUBLIC_VLAN450"
  ]
  name   = "098_papi-block-1-VLAN450"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_42" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_25-TCP_1883"
  ]
  tenants = [
    "Gr2_IoT_Wired.IOT_WIRED_VLAN492"
  ]
  name   = "182_pk-test-492-410-r0-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_43" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_168-TCP_225"
  ]
  tenants = [
    "CORE_25.CORE_VLAN410"
  ]
  name   = "034_410-cod-2-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_44" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_53"
  ]
  tenants = [
    "PUBLIC_VLAN450"
  ]
  name   = "101_450-control-2-VLAN450"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_45" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_129-UDP_123"
  ]
  tenants = [
    "CREW_VLAN451"
  ]
  name   = "109b_control-limited-4-VLAN451"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_46" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "any"
  ]
  tenants = [
    "CORE_VLAN410",
    "RESTRICTED_VLAN420",
    "MANAGEMENT_VLAN430",
    "PCI_VLAN440",
    "PRIVATE_VLAN491",
    "IOT_WIRED_VLAN492",
    "IOT_WIRELESS_VLAN494",
    "SECONDARY_UNTRUSTED_VLAN490.UNTRUSTED_VLAN490",
    "SECONDARY_PRIVATE_VLAN491.PRIVATE_VLAN491",
    "SECONDARY_IOT_WIRELESS_VLAN494.IOT_WIRELESS_VLAN494",
    "SECONDARY_CREW_VLAN451.CREW_VLAN451",
    "SECONDARY_IOT_WIRED_VLAN492.IOT_WIRED_VLAN492"
  ]
  name   = "Default-rule"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_47" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-TCP_1163",
    "ANY-TCP_8112"
  ]
  tenants = [
    "PUBLIC_VLAN450"
  ]
  name   = "099_450-control-0-VLAN450"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_48" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_169-TCP_225"
  ]
  tenants = [
    "CORE_25.CORE_VLAN410"
  ]
  name   = "035_410-cod-3-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_49" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_129-UDP_53"
  ]
  tenants = [
    "IOT_WIRED_VLAN492"
  ]
  name   = "176b_492-control-7-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_50" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_94_1-UDP_123"
  ]
  tenants = [
    "IOT_WIRED_VLAN492"
  ]
  name   = "175f_492-control-7-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_51" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_0-25-TCP_22501"
  ]
  tenants = [
    "Gr_IoT_WIRED.IOT_WIRED_VLAN492"
  ]
  name   = "196_410-492-2331-3-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_52" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_172-TCP_225"
  ]
  tenants = [
    "CORE_25.CORE_VLAN410"
  ]
  name   = "037_410-cod-5-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_53" {
  action = "deny"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "public_url_block"
  ]
  tenants = [
    "PUBLIC_VLAN450"
  ]
  name   = "103_URL-block"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_54" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_53"
  ]
  tenants = [
    "CORE_23.CORE_VLAN410",
    "CORE_24.CORE_VLAN410"
  ]
  name   = "013_qsrsoft-vm-rules-1-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_55" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_168-TCP_225"
  ]
  tenants = [
    "CORE_12.CORE_VLAN410"
  ]
  name   = "038_410-cod-6-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_56" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_123"
  ]
  tenants = [
    "CORE_23.CORE_VLAN410",
    "CORE_24.CORE_VLAN410"
  ]
  name   = "014_qsrsoft-vm-rules-2-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_57" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_172-TCP_225"
  ]
  tenants = [
    "CORE_12.CORE_VLAN410"
  ]
  name   = "041_410-cod-9-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_58" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_92_1-UDP_123"
  ]
  tenants = [
    "PRIVATE_VLAN491"
  ]
  name   = "166e_491-control-6-VLAN491"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_59" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-TCP_80"
  ]
  tenants = [
    "CORE_23.CORE_VLAN410",
    "CORE_24.CORE_VLAN410"
  ]
  name   = "015_qsrsoft-vm-rules-3-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_60" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_0-25-TCP_22501"
  ]
  tenants = [
    "Gr_IoT_WIRED.IOT_WIRED_VLAN492"
  ]
  name   = "192_410-492-bi-comm-8-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_61" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_168-TCP_225"
  ]
  tenants = [
    "CORE_14.CORE_VLAN410"
  ]
  name   = "042_410-cod-10-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_62" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_92_192-26-TCP_22500"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "193_410-492-2331-0-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_63" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_169-TCP_225"
  ]
  tenants = [
    "CORE_14.CORE_VLAN410"
  ]
  name   = "043_410-cod-11-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_64" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_92_192-26-TCP_22501"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "194_410-492-2331-1-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_65" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_171-TCP_225"
  ]
  tenants = [
    "CORE_14.CORE_VLAN410"
  ]
  name   = "044_410-cod-12-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_66" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_0-25-TCP_22500"
  ]
  tenants = [
    "Gr_IoT_WIRED.IOT_WIRED_VLAN492"
  ]
  name   = "195_410-492-2331-2-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_67" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_91_1-UDP_123"
  ]
  tenants = [
    "IOT_WIRELESS_VLAN494"
  ]
  name   = "232g_494-control-6-VLAN494"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_68" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_172-TCP_225"
  ]
  tenants = [
    "CORE_14.CORE_VLAN410"
  ]
  name   = "045_410-cod-13-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_69" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "any"
  ]
  tenants = [
    "PUBLIC_VLAN450"
  ]
  name   = "103_int-src-nat-pri4-0-VLAN450"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_70" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-TCP_4983"
  ]
  tenants = [
    "CORE_25.CORE_VLAN410"
  ]
  name   = "051_410-dmb-0-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_71" {
  action = "deny"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_8209"
  ]
  tenants = [
    "IOT_WIRELESS_VLAN494"
  ]
  name   = "225_papi-block-0-VLAN494"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_72" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_0-25-ANY"
  ]
  tenants = [
    "RESTRICTED_VLAN420"
  ]
  name   = "052_410-dmb-1-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_73" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-TCP_22700"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "056_410-dmb-5-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_74" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_90_130-UDP_20000"
  ]
  tenants = [
    "CORE_51.CORE_VLAN410"
  ]
  name   = "065_livestore-sec-cam-5-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_75" {
  action = "deny"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_8211"
  ]
  tenants = [
    "CREW_VLAN451"
  ]
  name   = "105_papi-block-1-VLAN451"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_76" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-TCP_80"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "053_410-dmb-2-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_77" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ICMP_ANY"
  ]
  tenants = [
    "IOT_WIRELESS_VLAN494"
  ]
  name   = "229_494-control-3-VLAN494"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_78" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-TCP_22500"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "054_410-dmb-3-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_79" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_62-UDP_123"
  ]
  tenants = [
    "CREW_VLAN451"
  ]
  name   = "109a_control-limited-4-VLAN451"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_80" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-TCP_3333"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "057_410-dmb-6-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_81" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_51-UDP_123"
  ]
  tenants = [
    "IOT_WIRELESS_VLAN494"
  ]
  name   = "231_494-control-5-VLAN494"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_82" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-TCP_3333"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "058_410-dmb-7-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_83" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_62-UDP_123"
  ]
  tenants = [
    "IOT_WIRELESS_VLAN494"
  ]
  name   = "232a_494-control-6-VLAN494"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_84" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-TCP_8888"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "059_410-dmb-8-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_85" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_90_130-TCP_2000"
  ]
  tenants = [
    "CORE_51.CORE_VLAN410"
  ]
  name   = "153_livestore-sec-cam-2-VLAN490"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_86" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "any"
  ]
  tenants = [
    "UNTRUSTED_VLAN490"
  ]
  name   = "158_int-src-nat-pri4-0-VLAN490"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_87" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_129-UDP_123"
  ]
  tenants = [
    "IOT_WIRELESS_VLAN494"
  ]
  name   = "232b_494-control-6-VLAN494"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_88" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_90_130-TCP_80"
  ]
  tenants = [
    "CORE_51.CORE_VLAN410"
  ]
  name   = "060_livestore-sec-cam-0-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_89" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_161-UDP_123"
  ]
  tenants = [
    "CREW_VLAN451"
  ]
  name   = "109c_control-limited-4-VLAN451"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_90" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_90_130-TCP_443"
  ]
  tenants = [
    "CORE_51.CORE_VLAN410"
  ]
  name   = "061_livestore-sec-cam-1-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_91" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_91_1-UDP_123"
  ]
  tenants = [
    "CREW_VLAN451"
  ]
  name   = "109e_control-limited-4-VLAN451"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_92" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_90_130-TCP_2000"
  ]
  tenants = [
    "CORE_51.CORE_VLAN410"
  ]
  name   = "062_livestore-sec-cam-2-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_93" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_92_1-UDP_123"
  ]
  tenants = [
    "CREW_VLAN451"
  ]
  name   = "109f_control-limited-4-VLAN451"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_94" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_0-25-TCP_22501"
  ]
  tenants = [
    "Gr_IoT_WIRED.IOT_WIRED_VLAN492"
  ]
  name   = "027_410-492-bi-com-8-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_95" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_94_1-UDP_123"
  ]
  tenants = [
    "CREW_VLAN451"
  ]
  name   = "109g_control-limited-4-VLAN451"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_96" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_90_130-TCP_2002"
  ]
  tenants = [
    "CORE_51.CORE_VLAN410"
  ]
  name   = "154_livestore-sec-cam-3-VLAN490"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_97" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_92_1-UDP_123"
  ]
  tenants = [
    "IOT_WIRELESS_VLAN494"
  ]
  name   = "232h_494-control-6-VLAN494"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_98" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_90_130-TCP_20000"
  ]
  tenants = [
    "CORE_51.CORE_VLAN410"
  ]
  name   = "155_livestore-sec-cam-4-VLAN490"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_99" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_62-UDP_53"
  ]
  tenants = [
    "IOT_WIRELESS_VLAN494"
  ]
  name   = "233a_494-control-7-VLAN494"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_100" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_90_130-UDP_20000"
  ]
  tenants = [
    "CORE_51.CORE_VLAN410"
  ]
  name   = "156_livestore-sec-cam-5-VLAN490"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_101" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ICMP_ANY"
  ]
  tenants = [
    "PRIVATE_VLAN491"
  ]
  name   = "163_491-control-3-VLAN491"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_102" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_129-UDP_53"
  ]
  tenants = [
    "IOT_WIRELESS_VLAN494"
  ]
  name   = "233b_494-control-7-VLAN494"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_103" {
  action = "deny"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_8209"
  ]
  tenants = [
    "PRIVATE_VLAN491"
  ]
  name   = "159_papi-block-0-VLAN491"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_104" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_161-UDP_53"
  ]
  tenants = [
    "IOT_WIRELESS_VLAN494"
  ]
  name   = "233c_494-control-7-VLAN494"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_105" {
  action = "deny"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_8211"
  ]
  tenants = [
    "PRIVATE_VLAN491"
  ]
  name   = "160_papi-block-1-VLAN491"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_106" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_91_1-UDP_53"
  ]
  tenants = [
    "IOT_WIRELESS_VLAN494"
  ]
  name   = "233g_494-control-7-VLAN494"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_107" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_94_1-UDP_53"
  ]
  tenants = [
    "IOT_WIRED_VLAN492"
  ]
  name   = "176f_492-control-7-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_108" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_25-UDP_123"
  ]
  tenants = [
    "PRIVATE_VLAN491"
  ]
  name   = "164_491-control-4-VLAN491"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_109" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_92_1-UDP_53"
  ]
  tenants = [
    "IOT_WIRELESS_VLAN494"
  ]
  name   = "233h_494-control-7-VLAN494"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_110" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "CORE-TCP-5173"
  ]
  tenants = [
    "PED_DEVICES-CORE-DNS.CORE_VLAN410",
    "PED_DEVICES-RESTRICTED-DNS.RESTRICTED_VLAN420"
  ]
  name   = "004_ped-devices-rules-3-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_111" {
  action = "deny"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_8209"
  ]
  tenants = [
    "UNTRUSTED_VLAN490"
  ]
  name   = "146_papi-block-0-VLAN490"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_112" {
  action = "deny"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_8211"
  ]
  tenants = [
    "UNTRUSTED_VLAN490"
  ]
  name   = "147_papi-block-1-VLAN490"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_113" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_0_0-16-ICMP"
  ]
  tenants = [
    "CREW_VLAN451"
  ]
  name   = "108_control-limited-3-VLAN451"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_114" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_51-UDP_123"
  ]
  tenants = [
    "PRIVATE_VLAN491"
  ]
  name   = "165_491-control-5-VLAN491"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_115" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_53"
  ]
  tenants = [
    "CREW_VLAN451"
  ]
  name   = "110_control-limited-5-VLAN451"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_116" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_92_50-ICMP"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "188_410-492-bi-comm-4-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_117" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_62-UDP_123"
  ]
  tenants = [
    "PRIVATE_VLAN491"
  ]
  name   = "166a_491-control-6-VLAN491"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_118" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_91_1-UDP_53"
  ]
  tenants = [
    "IOT_WIRED_VLAN492"
  ]
  name   = "176e_492-control-7-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_119" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_129-UDP_123"
  ]
  tenants = [
    "PRIVATE_VLAN491"
  ]
  name   = "166b_491-control-6-VLAN491"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_120" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_161-UDP_53"
  ]
  tenants = [
    "PED_DEVICES-CORE-DNS.CORE_VLAN410",
    "PED_DEVICES-RESTRICTED-DNS.RESTRICTED_VLAN420"
  ]
  name   = "002b-ped-devices-rules-0-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_121" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_161-UDP_123"
  ]
  tenants = [
    "PRIVATE_VLAN491"
  ]
  name   = "166c_491-control-6-VLAN491"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_122" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_91_1-UDP_53"
  ]
  tenants = [
    "PED_DEVICES-CORE-DNS.CORE_VLAN410",
    "PED_DEVICES-RESTRICTED-DNS.RESTRICTED_VLAN420"
  ]
  name   = "002d-ped-devices-rules-0-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_123" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_94_1-UDP_53"
  ]
  tenants = [
    "PED_DEVICES-CORE-DNS.CORE_VLAN410",
    "PED_DEVICES-RESTRICTED-DNS.RESTRICTED_VLAN420"
  ]
  name   = "066g_ped-devices-rules-0-VLAN420"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_124" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_94_1-UDP_123"
  ]
  tenants = [
    "PRIVATE_VLAN491"
  ]
  name   = "166f_491-control-6-VLAN491"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_125" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_92_1-UDP_53"
  ]
  tenants = [
    "PED_DEVICES-CORE-DNS.CORE_VLAN410",
    "PED_DEVICES-RESTRICTED-DNS.RESTRICTED_VLAN420"
  ]
  name   = "002e-ped-devices-rules-0-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_126" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_92_1-UDP_53"
  ]
  tenants = [
    "PRIVATE_VLAN491"
  ]
  name   = "167e_491-control-6-VLAN491"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_127" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "CORE-TCP-5172"
  ]
  tenants = [
    "PED_DEVICES-RESTRICTED-DNS.RESTRICTED_VLAN420",
    "PED_DEVICES-CORE-DNS.CORE_VLAN410"
  ]
  name   = "003_ped-devices-rules-2-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_128" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_25-TCP_8123"
  ]
  tenants = [
    "IOT_WIRED_50.IOT_WIRED_VLAN492"
  ]
  name   = "184_410-492-bi-comm-0-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_129" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_94_1-UDP_53"
  ]
  tenants = [
    "PRIVATE_VLAN491"
  ]
  name   = "167f_491-control-6-VLAN491"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_130" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_92_1-UDP_53"
  ]
  tenants = [
    "PED_DEVICES-CORE-DNS.CORE_VLAN410",
    "PED_DEVICES-RESTRICTED-DNS.RESTRICTED_VLAN420"
  ]
  name   = "066f_ped-devices-rules-0-VLAN420"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_131" {
  action = "deny"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_8209"
  ]
  tenants = [
    "IOT_WIRED_VLAN492"
  ]
  name   = "168_papi-block-0-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_132" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-TCP_1163",
    "ANY-TCP_8112"
  ]
  tenants = [
    "CREW_VLAN451"
  ]
  name   = "106_control-limited-0-VLAN451"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_133" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_25-TCP_8883"
  ]
  tenants = [
    "Gr2_IoT_Wired.IOT_WIRED_VLAN492"
  ]
  name   = "183_pk-test-492-410-r1-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_134" {
  action = "deny"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_8211"
  ]
  tenants = [
    "IOT_WIRED_VLAN492"
  ]
  name   = "169_papi-block-1-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_135" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_128-27-TCP_8123"
  ]
  tenants = [
    "CORE_1.CORE_VLAN410",
    "CORE_25.CORE_VLAN410"
  ]
  name   = "017_410-420-bi-com-0-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_136" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ICMP_ANY"
  ]
  tenants = [
    "IOT_WIRED_VLAN492"
  ]
  name   = "172_492-control-3-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_137" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_62-UDP_53"
  ]
  tenants = [
    "PED_DEVICES-CORE-DNS.CORE_VLAN410",
    "PED_DEVICES-RESTRICTED-DNS.RESTRICTED_VLAN420"
  ]
  name   = "066a_ped-devices-rules-0-VLAN420"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_138" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "RESTRICTED-TCP-5015"
  ]
  tenants = [
    "PED_DEVICES-CORE-TCP-5015.CORE_VLAN410"
  ]
  name   = "069_ped-devices-rules-4-VLAN420"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_139" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_129-UDP_53"
  ]
  tenants = [
    "PED_DEVICES-CORE-DNS.CORE_VLAN410",
    "PED_DEVICES-RESTRICTED-DNS.RESTRICTED_VLAN420"
  ]
  name   = "066b_ped-devices-rules-0-VLAN420"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_140" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "RESTRICTED-TCP-5016"
  ]
  tenants = [
    "PED_DEVICES-CORE-TCP-5015.CORE_VLAN410"
  ]
  name   = "070_ped-devices-rules-5-VLAN420"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_141" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_62-UDP_53"
  ]
  tenants = [
    "IOT_WIRED_VLAN492"
  ]
  name   = "176a_492-control-7-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_142" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_172-TCP_225"
  ]
  tenants = [
    "CORE_12.CORE_VLAN410"
  ]
  name   = "092_410-cod-9-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_143" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_161-UDP_123"
  ]
  tenants = [
    "IOT_WIRELESS_VLAN494"
  ]
  name   = "232c_494-control-6-VLAN494"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_144" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_161-UDP_53"
  ]
  tenants = [
    "PED_DEVICES-CORE-DNS.CORE_VLAN410",
    "PED_DEVICES-RESTRICTED-DNS.RESTRICTED_VLAN420"
  ]
  name   = "066c_ped-devices-rules-0-VLAN420"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_145" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_67",
    "ANY-UDP_68"
  ]
  tenants = [
    "CREW_VLAN451"
  ]
  name   = "107_control-limited-2-VLAN451"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_146" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_92_50-ICMP"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "187_410-492-bi-comm-3-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_147" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "CORE_1_25-TCP-8123"
  ]
  tenants = [
    "RESTRICTED_VLAN420"
  ]
  name   = "072_410-420-bi-com-1-VLAN420"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_148" {
  action = "deny"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_8211"
  ]
  tenants = [
    "IOT_WIRELESS_VLAN494"
  ]
  name   = "226_papi-block-1-VLAN494"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_149" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_91_1-UDP_53"
  ]
  tenants = [
    "PED_DEVICES-CORE-DNS.CORE_VLAN410",
    "PED_DEVICES-RESTRICTED-DNS.RESTRICTED_VLAN420"
  ]
  name   = "066e_ped-devices-rules-0-VLAN420"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_150" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_92_192-26-TCP_22500"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "189_410-492-bi-comm-5-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_151" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-TCP_4983"
  ]
  tenants = [
    "CORE_25.CORE_VLAN410"
  ]
  name   = "073_410-dmb-0-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_152" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_91_1-UDP_123"
  ]
  tenants = [
    "IOT_WIRED_VLAN492"
  ]
  name   = "175e_492-control-7-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_153" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_92_192-26-TCP_22501"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "190_410-492-bi-comm-6-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_154" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_129-UDP_53"
  ]
  tenants = [
    "PED_DEVICES-CORE-DNS.CORE_VLAN410"
  ]
  name   = "002a-ped-devices-rules-0-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_155" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_0-25-ANY"
  ]
  tenants = [
    "RESTRICTED_VLAN420"
  ]
  name   = "074_410-dmb-1-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_156" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_0-25-TCP_22500"
  ]
  tenants = [
    "Gr_IoT_WIRED.IOT_WIRED_VLAN492"
  ]
  name   = "191_410-492-bi-comm-7-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_157" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-TCP_1163",
    "ANY-TCP_8112"
  ]
  tenants = [
    "UNTRUSTED_VLAN490"
  ]
  name   = "148_control-vlan-490-0-VLAN490"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_158" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "RESTRICTED-TCP-5015"
  ]
  tenants = [
    "PED_DEVICES-CORE-TCP-5015.CORE_VLAN410"
  ]
  name   = "005_ped-devices-rules-4-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_159" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-TCP_22500"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "076_410-dmb-3-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_160" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_67",
    "ANY-UDP_68"
  ]
  tenants = [
    "UNTRUSTED_VLAN490"
  ]
  name   = "149_control-vlan-490-1-VLAN490"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_161" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_25-TCP_1883"
  ]
  tenants = [
    "Gr2_IoT_Wired.IOT_WIRED_VLAN492"
  ]
  name   = "007_pk-test-492-410-r0-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_162" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-TCP_22501"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "077_410-dmb-4-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_163" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_129-UDP_123"
  ]
  tenants = [
    "IOT_WIRED_VLAN492"
  ]
  name   = "175b_492-control-5-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_164" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-TCP_1163",
    "ANY-TCP_8112"
  ]
  tenants = [
    "PRIVATE_VLAN491"
  ]
  name   = "161_491-control-0-VLAN491"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_165" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_51-ICMP"
  ]
  tenants = [
    "IOT_WIRED_50.IOT_WIRED_VLAN492"
  ]
  name   = "186_410-492-bi-comm-2-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_166" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "RESTRICTED-TCP-5016"
  ]
  tenants = [
    "PED_DEVICES-CORE-TCP-5016.CORE_VLAN410"
  ]
  name   = "006_ped-devices-rules-5-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_167" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_129-UDP_53"
  ]
  tenants = [
    "PRIVATE_VLAN491"
  ]
  name   = "167b_491-control-6-VLAN491"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_168" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-TCP_22600"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "078_410-dmb-5-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_169" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-TCP_4983"
  ]
  tenants = [
    "Gr_CORE_12_14_25.CORE_VLAN410"
  ]
  name   = "084_410-cod-1-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_170" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_67",
    "ANY-UDP_68"
  ]
  tenants = [
    "PRIVATE_VLAN491"
  ]
  name   = "162_491-control-2-VLAN491"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_171" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_25-TCP_8123"
  ]
  tenants = [
    "IOT_WIRED_50.IOT_WIRED_VLAN492"
  ]
  name   = "019_410-492-bi-com-0-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_172" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-TCP_22700"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "079_410-dmb-6-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_173" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-TCP_1163",
    "ANY-TCP_8112"
  ]
  tenants = [
    "IOT_WIRED_VLAN492"
  ]
  name   = "170_492-control-0-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_174" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_25-TCP_8883"
  ]
  tenants = [
    "Gr2_IoT_Wired.IOT_WIRED_VLAN492"
  ]
  name   = "008_pk-test-492-410-r1-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_175" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_0-25-TCP_22500"
  ]
  tenants = [
    "Gr_IoT_WIRED.IOT_WIRED_VLAN492"
  ]
  name   = "026_410-492-bi-com-7-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_176" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-TCP_3333"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "080_410-dmb-7-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_177" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-TCP_80"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "075_410-dmb-2-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_178" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_67",
    "ANY-UDP_68"
  ]
  tenants = [
    "IOT_WIRED_VLAN492"
  ]
  name   = "171_492-control-2-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_179" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_90_130-TCP_443"
  ]
  tenants = [
    "CORE_51.CORE_VLAN410"
  ]
  name   = "152_livestore-sec-cam-1-VLAN490"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_180" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_53"
  ]
  tenants = [
    "CORE_51.CORE_VLAN410"
  ]
  name   = "009_410-bos-isp-0-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_181" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-TCP_22501"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "055_410-dmb-4-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_182" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-TCP_8888"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "081_410-dmb-8-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_183" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-TCP_1163",
    "ANY-TCP_8112"
  ]
  tenants = [
    "IOT_WIRELESS_VLAN494"
  ]
  name   = "227_494-control-0-VLAN494"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_184" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-TCP_443"
  ]
  tenants = [
    "CORE_51.CORE_VLAN410"
  ]
  name   = "010_410-bos-isp-1-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_185" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-ANY"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "082_410-dmb-9-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_186" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-UDP_67",
    "ANY-UDP_68"
  ]
  tenants = [
    "IOT_WIRELESS_VLAN494"
  ]
  name   = "228_494-control-2-VLAN494"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_187" {
  action        = "allow"
  local_routing = false
  services = [
    "any"
  ]
  tenants = [
    "AUTOMATION_VLAN460"
  ]
  name   = "460_to_any_prisma_bypass"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_188" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_160-27-TCP_225"
  ]
  tenants = [
    "Gr_CORE_12_14_25.CORE_VLAN410"
  ]
  name   = "083_410-cod-0-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_189" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_168-TCP_225"
  ]
  tenants = [
    "CORE_14.CORE_VLAN410"
  ]
  name   = "093_410-cod-10-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_190" {
  action = "deny"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "172_16_0_0-16-ANY",
    "192_168_0_0-16-ANY",
    "10_0_0_0-8-ANY"
  ]
  tenants = [
    "PUBLIC_VLAN450"
  ]
  name   = "102_450-control-5-VLAN450"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_191" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-TCP_80"
  ]
  tenants = [
    "CORE_51.CORE_VLAN410"
  ]
  name   = "011_410-bos-isp-2-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_192" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_161-UDP_53"
  ]
  tenants = [
    "PRIVATE_VLAN491"
  ]
  name   = "167c_491-control-6-VLAN491"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_193" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_168-TCP_225"
  ]
  tenants = [
    "CORE_25.CORE_VLAN410"
  ]
  name   = "085_410-cod-2-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_194" {
  action = "deny"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "172_16_0_0-16-ANY",
    "192_168_0_0-16-ANY",
    "10_0_0_0-8-ANY"
  ]
  tenants = [
    "CREW_VLAN451"
  ]
  name   = "111_deny-rfc-privateip-0-VLAN451"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_195" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_94_0-23-TCP_22"
  ]
  tenants = [
    "Gr2_IoT_Wired.IOT_WIRED_VLAN492"
  ]
  name   = "181_pk-test-494-492-r4-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_196" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "any"
  ]
  tenants = [
    "MGMT"
  ]
  name   = "IN_BAND_MGMT"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_197" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_62-UDP_53"
  ]
  tenants = [
    "PRIVATE_VLAN491"
  ]
  name   = "167a_491-control-6-VLAN491"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_198" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_169-TCP_225"
  ]
  tenants = [
    "CORE_25.CORE_VLAN410"
  ]
  name   = "086_410-cod-3-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_199" {
  action = "deny"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_6_0-24-ANY",
    "192_168_90_0-24-ANY",
    "192_168_91_0-24-ANY",
    "10_X_Y_0-24-ANY",
    "192_168_4_0-23-ANY"
  ]
  tenants = [
    "PUBLIC_VLAN450",
    "CREW_VLAN451",
    "UNTRUSTED_VLAN490",
    "CORE_VLAN410",
    "AUTOMATION_VLAN460"
  ]
  name   = "Deny-Any-Any"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_200" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "ANY-TCP_22"
  ]
  tenants = [
    "Gr_CORE_12_14_25.CORE_VLAN410"
  ]
  name   = "012_qsrsoft-vm-rules-0-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_201" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "192_168_92_192-26-TCP_22500"
  ]
  tenants = [
    "CORE_VLAN410"
  ]
  name   = "024_410-492-bi-com-5-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_202" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_171-TCP_225"
  ]
  tenants = [
    "CORE_25.CORE_VLAN410"
  ]
  name   = "087_410-cod-4-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_203" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "CORE-TCP-5173"
  ]
  tenants = [
    "PED_DEVICES-CORE-DNS.CORE_VLAN410",
    "PED_DEVICES-RESTRICTED-DNS.RESTRICTED_VLAN420"
  ]
  name   = "068_ped-devices-rules-3-VLAN420"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_204" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_51-ICMP"
  ]
  tenants = [
    "IOT_WIRED_50.IOT_WIRED_VLAN492"
  ]
  name   = "021_410-492-bi-com-2-VLAN410"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_205" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_172-TCP_225"
  ]
  tenants = [
    "CORE_25.CORE_VLAN410"
  ]
  name   = "088_410-cod-5-VLAN430"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_206" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "10_X_Y_161-UDP_123"
  ]
  tenants = [
    "IOT_WIRED_VLAN492"
  ]
  name   = "175c_492-control-5-VLAN492"
  org_id = mist_org.terraform_test.id
}
resource "mist_org_servicepolicy" "servicepolicy_207" {
  action = "allow"
  idp = {
    enabled = false
    profile = "strict"
  }
  local_routing = false
  services = [
    "CORE-TCP-5172"
  ]
  tenants = [
    "PED_DEVICES-CORE-DNS.CORE_VLAN410",
    "PED_DEVICES-RESTRICTED-DNS.RESTRICTED_VLAN420"
  ]
  name   = "067_ped-devices-rules-2-VLAN420"
  org_id = mist_org.terraform_test.id
}




############# GW
# resource "mist_device_gateway" "gateway_1" {
#   device_id = ""
#   site_id   = mist_site.terraform_site.id
#   name      = ""
#   additional_config_cmds = [
#     "# US_Prod_04270 - CLI Commands - 11/21/2024",
#     "# 4 Tunnels, 4 LAN Links, Re-IP LAN Networks, Quarantine",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# Increase the number of rollback configs, configure tcp-mss and avoid rtlog-conn-err (CPU)",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "set system max-configuration-rollbacks 48",
#     "set security flow tcp-mss all-tcp mss 1350",
#     "set security flow tcp-mss ipsec-vpn mss 1200",
#     "",
#     "#deactivate groups mist-script event-options policy rtlog-conn-err",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# Configure security policy logging",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "set groups top security policies from-zone <*> to-zone <*> policy <*> then log session-init",
#     "#set security log mode stream-event",
#     "#set groups top system syslog file traffic-log any any",
#     "#set groups top system syslog file traffic-log archive size 10m",
#     "#set groups top system syslog file traffic-log archive files 10",
#     "#set groups top system syslog file traffic-log match \"RT_FLOW:|RT_UTM\"",
#     "",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# Configure security policy logging",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "delete security log mode stream-event",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# BGP to Prisma-Tunnel",
#     "# Configure BGP inside IPSec to Prisma-Tunnel",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "set interfaces st0 unit 1000 family inet address 7.4.0.112/31",
#     "set groups top routing-instances tun_Prisma-Tunnel protocols bgp group SSE1 type external",
#     "set groups top routing-instances tun_Prisma-Tunnel protocols bgp group SSE1 multihop",
#     "set groups top routing-instances tun_Prisma-Tunnel protocols bgp group SSE1 export EXP-TO-PRISMA-PRIMARY",
#     "set groups top routing-instances tun_Prisma-Tunnel protocols bgp group SSE1 peer-as 64781",
#     "set groups top routing-instances tun_Prisma-Tunnel protocols bgp group SSE1 local-as 64777",
#     "set groups top routing-instances tun_Prisma-Tunnel protocols bgp group SSE1 neighbor 7.4.0.113 import IMPORT-FROM-PRISMA-PRIMARY",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# BGP to Prisma-Tunnel-II",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "set interfaces st0 unit 1001 family inet address 7.4.0.118/31",
#     "set groups top routing-instances tun_Prisma-Tunnel-II protocols bgp group SSE2 type external",
#     "set groups top routing-instances tun_Prisma-Tunnel-II protocols bgp group SSE2 multihop",
#     "set groups top routing-instances tun_Prisma-Tunnel-II protocols bgp group SSE2 export EXP-TO-PRISMA-SECONDARY",
#     "set groups top routing-instances tun_Prisma-Tunnel-II protocols bgp group SSE2 peer-as 64781",
#     "set groups top routing-instances tun_Prisma-Tunnel-II protocols bgp group SSE2 local-as 64777",
#     "set groups top routing-instances tun_Prisma-Tunnel-II protocols bgp group SSE2 neighbor 7.4.0.119 import IMPORT-FROM-PRISMA-SECONDARY",
#     "",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# BGP to Prisma-Tunnel-III",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "",
#     "set interfaces st0 unit 1002 family inet address 7.4.0.116/31",
#     "set groups top routing-instances tun_Prisma-Tunnel-III protocols bgp group SSE3 type external",
#     "set groups top routing-instances tun_Prisma-Tunnel-III protocols bgp group SSE3 multihop",
#     "set groups top routing-instances tun_Prisma-Tunnel-III protocols bgp group SSE3 log-updown",
#     "set groups top routing-instances tun_Prisma-Tunnel-III protocols bgp group SSE3 export EXP-TO-PRISMA-III",
#     "set groups top routing-instances tun_Prisma-Tunnel-III protocols bgp group SSE3 peer-as 64781",
#     "set groups top routing-instances tun_Prisma-Tunnel-III protocols bgp group SSE3 local-as 64777",
#     "set groups top routing-instances tun_Prisma-Tunnel-III protocols bgp group SSE3 neighbor 7.4.0.117 import IMPORT-FROM-PRISMA-III",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# BGP to Prisma-Tunnel-IV",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "set interfaces st0 unit 1003 family inet address 7.4.0.114/31",
#     "set groups top routing-instances tun_Prisma-Tunnel-IV protocols bgp group SSE4 type external",
#     "set groups top routing-instances tun_Prisma-Tunnel-IV protocols bgp group SSE4 multihop",
#     "set groups top routing-instances tun_Prisma-Tunnel-IV protocols bgp group SSE4 log-updown",
#     "set groups top routing-instances tun_Prisma-Tunnel-IV protocols bgp group SSE4 export EXP-TO-PRISMA-IV",
#     "set groups top routing-instances tun_Prisma-Tunnel-IV protocols bgp group SSE4 peer-as 64781",
#     "set groups top routing-instances tun_Prisma-Tunnel-IV protocols bgp group SSE4 local-as 64777",
#     "set groups top routing-instances tun_Prisma-Tunnel-IV protocols bgp group SSE4 neighbor 7.4.0.115 import IMPORT-FROM-PRISMA-IV",
#     "",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# BGP EXPORT/IMPORT Policy to PRISMA-PRIMARY",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-PRIMARY term from-lan from instance LAN",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-PRIMARY term from-lan from protocol direct",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-PRIMARY term from-lan from protocol aggregate",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-PRIMARY term from-lan from route-filter 10.220.93.0/24 exact",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-PRIMARY term from-lan from route-filter 172.24.151.64/26 exact",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-PRIMARY term from-lan from route-filter 192.168.0.0/16 orlonger",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-PRIMARY term from-lan from route-filter 29.220.93.0/24 orlonger",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-PRIMARY term from-lan from route-filter 30.220.93.0/24 orlonger",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-PRIMARY term from-lan then accept",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-PRIMARY term from-local from protocol static",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-PRIMARY term from-local then accept",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-PRIMARY term reject then reject",
#     "",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA-PRIMARY term from-bgp from protocol bgp",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA-PRIMARY term from-bgp then accept",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA-PRIMARY term from-bgp then local-preference 330",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA-PRIMARY term reject then reject",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# BGP EXPORT/IMPORT Policy to PRISMA-SECONDARY",
#     "#                                               BGP path prepend to prefer primary path",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-SECONDARY term from-lan from instance LAN",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-SECONDARY term from-lan from protocol direct",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-SECONDARY term from-lan from protocol aggregate",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-SECONDARY term from-lan from route-filter 10.220.93.0/24 exact",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-SECONDARY term from-lan from route-filter 172.24.151.64/26 exact",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-SECONDARY term from-lan from route-filter 192.168.0.0/16 orlonger",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-SECONDARY term from-lan from route-filter 29.220.93.0/24 orlonger",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-SECONDARY term from-lan from route-filter 30.220.93.0/24 orlonger",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-SECONDARY term from-lan then accept",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-SECONDARY term from-lan then as-path-prepend \"64777 64777 64777\"",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-SECONDARY term from-local from protocol static",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-SECONDARY term from-local then accept",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-SECONDARY term from-local then as-path-prepend \"64777 64777 64777\"",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-SECONDARY term reject then reject",
#     "",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA-SECONDARY term from-bgp from protocol bgp",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA-SECONDARY term from-bgp then accept",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA-SECONDARY term from-bgp then local-preference 320",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA-SECONDARY term reject then reject",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# BGP EXPORT/IMPORT Policy to PRISMA-III",
#     "#                                               BGP path prepend to prefer primary path",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-III term from-lan from instance LAN",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-III term from-lan from protocol direct",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-III term from-lan from protocol aggregate",
#     "",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-III term from-lan from route-filter 10.220.93.0/24 exact",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-III term from-lan from route-filter 172.24.151.64/26 exact",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-III term from-lan from route-filter 192.168.0.0/16 orlonger",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-III term from-lan from route-filter 29.220.93.0/24 orlonger",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-III term from-lan from route-filter 30.220.93.0/24 orlonger",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-III term from-lan then accept",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-III term from-lan then as-path-prepend \"64777 64777 64777 64777 64777 64777\"",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-III term from-local from protocol static",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-III term from-local then accept",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-III term from-local then as-path-prepend \"64777 64777 64777 64777 64777 64777\"",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-III term reject then reject",
#     "",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA-III term from-bgp from protocol bgp",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA-III term from-bgp then accept",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA-III term from-bgp then local-preference 310",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA-III term reject then reject",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# BGP EXPORT/IMPORT Policy to PRISMA-IV",
#     "#                                               BGP path prepend to prefer primary path",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-IV term from-lan from instance LAN",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-IV term from-lan from protocol direct",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-IV term from-lan from protocol aggregate",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-IV term from-lan from route-filter 10.220.93.0/24 exact",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-IV term from-lan from route-filter 172.24.151.64/26 exact",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-IV term from-lan from route-filter 192.168.0.0/16 orlonger",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-IV term from-lan from route-filter 29.220.93.0/24 orlonger",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-IV term from-lan from route-filter 30.220.93.0/24 orlonger",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-IV term from-lan then accept",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-IV term from-lan then as-path-prepend \"64777 64777 64777 64777 64777 64777 64777 64777 64777\"",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-IV term from-local from protocol static",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-IV term from-local then accept",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-IV term from-local then as-path-prepend \"64777 64777 64777 64777 64777 64777 64777 64777 64777\"",
#     "set groups top policy-options policy-statement EXP-TO-PRISMA-IV term reject then reject",
#     "",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA-IV term from-bgp from protocol bgp",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA-IV term from-bgp then accept",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA-IV term from-bgp then local-preference 300",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA-IV term reject then reject",
#     "",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# Export policy that advertise only summary route to tun_Prisma-Tunnel side - REQUIRED",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "set groups top routing-instances LAN routing-options aggregate route 10.220.93.0/24",
#     "set groups top routing-instances LAN routing-options aggregate route 29.220.93.0/24",
#     "set groups top routing-instances LAN routing-options aggregate route 30.220.93.0/24",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# Leaking LAN Prefixes to the PRISMA's Routing Instance so they can be advertised to PRISMA",
#     "# Also need to have the LAN Prefixes in the PRISMA's RI's for Incoming Traffic from tun_Prisma-Tunnel",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "set groups top policy-options policy-statement LAN-RTs term from-lan from instance LAN",
#     "set groups top policy-options policy-statement LAN-RTs term from-lan from protocol direct",
#     "set groups top policy-options policy-statement LAN-RTs term from-lan from protocol aggregate",
#     "set groups top policy-options policy-statement LAN-RTs term from-lan then accept",
#     "set groups top policy-options policy-statement LAN-RTs term reject then reject",
#     "",
#     "set groups top routing-instances tun_Prisma-Tunnel routing-options instance-import LAN-RTs",
#     "set groups top routing-instances tun_Prisma-Tunnel-II routing-options instance-import LAN-RTs",
#     "set groups top routing-instances tun_Prisma-Tunnel-III routing-options instance-import LAN-RTs",
#     "set groups top routing-instances tun_Prisma-Tunnel-IV routing-options instance-import LAN-RTs",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# IMPORT-FROM-PRISMA - Policy to leak routes advertised from PRISMA to the LAN Routing Instance",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA term 01 from instance tun_Prisma-Tunnel",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA term 01 from protocol bgp",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA term 01 then accept",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA term 02 from instance tun_Prisma-Tunnel-II",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA term 02 from protocol bgp",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA term 02 then accept",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA term 03 from instance tun_Prisma-Tunnel-III",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA term 03 from protocol bgp",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA term 03 then accept",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA term 04 from instance tun_Prisma-Tunnel-IV",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA term 04 from protocol bgp",
#     "set groups top policy-options policy-statement IMPORT-FROM-PRISMA term 04 then accept",
#     "#set groups top policy-options policy-statement IMPORT-FROM-PRISMA term 05 then reject",
#     "",
#     "#set groups top routing-instances LAN routing-options instance-import IMPORT-FROM-PRISMA",
#     "set groups top routing-instances apbr_TO-PRISMA-SSE routing-options instance-import IMPORT-FROM-PRISMA",
#     "#insert groups top routing-instances apbr_TO-PRISMA-SSE routing-options instance-import IMPORT-FROM-PRISMA before tun_Prisma-Tunnel-IV_direct",
#     "insert groups top routing-instances apbr_TO-PRISMA-SSE routing-options instance-import IMPORT-FROM-PRISMA after apbr_TO-PRISMA-SSE_wan",
#     "",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# Prevent chassis alarm due to dedicated management Ethernet interface being down (we're not using it)",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "set chassis alarm management-ethernet link-down ignore",
#     "",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# Allow ping and traceroute on CORE",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "set groups top security zones security-zone <*_*> host-inbound-traffic system-services ping",
#     "set groups top security zones security-zone <*_*> host-inbound-traffic system-services traceroute",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# Traffic from Prisma",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "set security policies global policy from_prisma match source-address any",
#     "set security policies global policy from_prisma match destination-address any",
#     "set security policies global policy from_prisma match application any",
#     "set security policies global policy from_prisma match dynamic-application any",
#     "set security policies global policy from_prisma match from-zone tun_Prisma-Tunnel",
#     "set security policies global policy from_prisma match from-zone tun_Prisma-Tunnel-II",
#     "set security policies global policy from_prisma match from-zone tun_Prisma-Tunnel-III",
#     "set security policies global policy from_prisma match from-zone tun_Prisma-Tunnel-IV",
#     "set security policies global policy from_prisma match to-zone CORE_VLAN410",
#     "set security policies global policy from_prisma match to-zone RESTRICTED_VLAN420",
#     "set security policies global policy from_prisma match to-zone MANAGEMENT_VLAN430",
#     "set security policies global policy from_prisma match to-zone PCI_VLAN440",
#     "set security policies global policy from_prisma match to-zone AUTOMATION_VLAN460",
#     "",
#     "set security policies global policy from_prisma match to-zone UNTRUSTED_VLAN490",
#     "set security policies global policy from_prisma match to-zone CREW_VLAN451",
#     "set security policies global policy from_prisma match to-zone PRIVATE_VLAN491",
#     "set security policies global policy from_prisma match to-zone IOT_WIRED_VLAN492",
#     "set security policies global policy from_prisma match to-zone IOT_WIRELESS_VLAN494",
#     "set security policies global policy from_prisma then permit",
#     "set security policies global policy from_prisma then log session-init",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# reject prisma in WAN",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "set groups top policy-options policy-statement apbr_TO-INTERNET_wan term everything-else then reject",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# Reth Interfaces and RG's - Separate RG for WAN_0, WAN_1 and LAN",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "delete chassis cluster redundancy-group 1",
#     "delete chassis cluster redundancy-group 2",
#     "delete chassis cluster redundancy-group 3",
#     "set chassis cluster redundancy-group 0 node 0 priority 100",
#     "set chassis cluster redundancy-group 0 node 1 priority 1",
#     "set chassis cluster redundancy-group 1 node 0 priority 100",
#     "set chassis cluster redundancy-group 1 node 1 priority 1",
#     "set chassis cluster redundancy-group 2 node 0 priority 100",
#     "set chassis cluster redundancy-group 2 node 1 priority 1",
#     "set chassis cluster redundancy-group 3 node 0 priority 100",
#     "set chassis cluster redundancy-group 3 node 1 priority 1",
#     "set chassis cluster redundancy-group 1 preempt",
#     "set chassis cluster redundancy-group 1 hold-down-interval 10",
#     "set chassis cluster redundancy-group 1 interface-monitor ge-0/0/0 weight 255",
#     "set chassis cluster redundancy-group 1 interface-monitor ge-5/0/0 weight 255",
#     "set chassis cluster redundancy-group 2 preempt",
#     "set chassis cluster redundancy-group 2 hold-down-interval 10",
#     "set chassis cluster redundancy-group 2 interface-monitor ge-0/0/3 weight 255",
#     "set chassis cluster redundancy-group 2 interface-monitor ge-5/0/3 weight 255",
#     "set chassis cluster redundancy-group 3 preempt",
#     "set chassis cluster redundancy-group 3 hold-down-interval 10",
#     "set chassis cluster redundancy-group 3 interface-monitor ge-0/0/6 weight 128",
#     "set chassis cluster redundancy-group 3 interface-monitor ge-0/0/7 weight 128",
#     "set chassis cluster redundancy-group 3 interface-monitor ge-5/0/6 weight 128",
#     "set chassis cluster redundancy-group 3 interface-monitor ge-5/0/7 weight 128",
#     "set interfaces reth1 redundant-ether-options redundancy-group 1",
#     "set interfaces reth2 redundant-ether-options redundancy-group 2",
#     "set interfaces reth3 redundant-ether-options redundancy-group 3",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# MGMT - Permit SSH access from the EX switches for when connectivity to Mist is lost",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "set security zones security-zone MGMT host-inbound-traffic system-services ssh",
#     "set security zones security-zone MGMT host-inbound-traffic system-services ping",
#     "set system services ssh root-login allow",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# WAN_0 Better Preference - Default route in inet.0 to use Broadband rather than CradlePoint",
#     "# e.g. SRX outbound-ssh, DNS, etc - And place traffic in control queue rather than default",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "set groups top policy-options policy-statement wan_default term 02_WAN_0_default then preference 11",
#     "set class-of-service host-outbound-traffic forwarding-class network-control",
#     "set class-of-service host-outbound-traffic dscp-code-point cs6",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# Security Log stream - McD custom-syslog",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "",
#     "set groups top services ssl initiation profile custom-tls protocol-version all",
#     "set groups top services ssl initiation profile custom-tls trusted-ca all",
#     "set groups top services ssl initiation profile custom-tls client-certificate mist-device-cert",
#     "set groups top services ssl initiation profile custom-tls actions ignore-server-auth-failure",
#     "set groups top services ssl initiation profile custom-tls actions crl disable",
#     "set groups top security log stream custom-syslog category all",
#     "set groups top security log stream custom-syslog host default.usrestaurants.friendly-goldwasser-jhy7mtd.cribl.cloud",
#     "set groups top security log stream custom-syslog host port 6514",
#     "set groups top security log stream custom-syslog transport tcp-connections 1",
#     "set groups top security log stream custom-syslog transport protocol tls",
#     "set groups top security log stream custom-syslog transport tls-profile custom-tls",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# DHCP exclude range",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "set groups top routing-instances LAN access address-assignment pool CORE_VLAN410 family inet excluded-range exclude1 low 10.220.93.24",
#     "set groups top routing-instances LAN access address-assignment pool CORE_VLAN410 family inet excluded-range exclude1 high 10.220.93.24",
#     "",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# DNS Proxy with second loopback IP",
#     "# DNS Proxy enabled in lo0 and using DNAT towards lo0. DNS Requested sent over Prisma Tunnels to OpenDNS 208.67.222.222",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "",
#     "set groups DNS_PROXY system services dns dns-proxy interface lo0.0",
#     "set groups DNS_PROXY system services dns forwarders 208.67.222.222",
#     "set groups DNS_PROXY system services dns forwarders 208.67.220.220",
#     "set groups DNS_PROXY routing-options static route 10.220.93.0/25 next-table LAN.inet.0",
#     "set groups DNS_PROXY routing-options static route 208.67.222.222/32 next-table apbr_TO-PRISMA-SSE.inet.0",
#     "set groups DNS_PROXY routing-options static route 208.67.220.220/32 next-table apbr_TO-PRISMA-SSE.inet.0",
#     "set apply-groups DNS_PROXY",
#     "set security nat destination pool pool1 address 100.10.0.1/32",
#     "set security nat destination rule-set qsrsoft from zone CORE_VLAN410",
#     "set security nat destination rule-set qsrsoft rule r1 match source-address 10.220.93.23/32",
#     "set security nat destination rule-set qsrsoft rule r1 match source-address 10.220.93.24/32",
#     "set security nat destination rule-set qsrsoft rule r1 match destination-address 10.220.93.62/32",
#     "set security nat destination rule-set qsrsoft rule r1 match destination-port 53",
#     "set security nat destination rule-set qsrsoft rule r1 then destination-nat pool pool1",
#     "set groups top routing-instances LAN routing-options instance-import lo0_to_LAN",
#     "set groups DNS_PROXY policy-options policy-statement lo0_to_LAN term lo0 from instance master",
#     "set groups DNS_PROXY policy-options policy-statement lo0_to_LAN term lo0 from interface lo0.0",
#     "set groups DNS_PROXY policy-options policy-statement lo0_to_LAN term lo0 then accept",
#     "set groups DNS_PROXY policy-options policy-statement lo0_to_LAN term 2 then reject",
#     "set security nat source rule-set host_to_Prisma from zone junos-host",
#     "set security nat source rule-set host_to_Prisma to zone tun_Prisma-Tunnel",
#     "set security nat source rule-set host_to_Prisma to zone tun_Prisma-Tunnel-II",
#     "set security nat source rule-set host_to_Prisma to zone tun_Prisma-Tunnel-III",
#     "set security nat source rule-set host_to_Prisma to zone tun_Prisma-Tunnel-IV",
#     "set security nat source rule-set host_to_Prisma rule rule1 match source-address 100.10.0.1/32",
#     "set security nat source rule-set host_to_Prisma rule rule1 match destination-address 208.67.222.222/32",
#     "set security nat source rule-set host_to_Prisma rule rule1 then source-nat interface",
#     "set security nat source rule-set host_to_Prisma rule rule2 match source-address 100.10.0.1/32",
#     "set security nat source rule-set host_to_Prisma rule rule2 match destination-address 208.67.220.220/32",
#     "set security nat source rule-set host_to_Prisma rule rule2 then source-nat interface",
#     "set groups DNS_PROXY routing-instances tun_Prisma-Tunnel   routing-options static route 100.10.0.1/32 next-table inet.0",
#     "set groups DNS_PROXY routing-instances tun_Prisma-Tunnel-II  routing-options static route 100.10.0.1/32 next-table inet.0",
#     "set groups DNS_PROXY routing-instances tun_Prisma-Tunnel-III routing-options static route 100.10.0.1/32 next-table inet.0",
#     "set groups DNS_PROXY routing-instances tun_Prisma-Tunnel-IV  routing-options static route 100.10.0.1/32 next-table inet.0",
#     "set security zones security-zone CORE_VLAN410 host-inbound-traffic system-services dns",
#     "set groups top firewall family inet filter protect_re term dns_proxy from protocol udp",
#     "set groups top firewall family inet filter protect_re term dns_proxy from port 53",
#     "set groups top firewall family inet filter protect_re term dns_proxy then accept",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# Additional Loopback address 100.10.0.1 - lower than 100.100.0.X",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "",
#     "set groups top interfaces lo0 unit 0 family inet address 100.10.0.1/32",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# Modify Protect_RE filter to permit DNS:",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "",
#     "set groups top firewall family inet filter protect_re term dns_proxy from protocol udp",
#     "set groups top firewall family inet filter protect_re term dns_proxy from port 53",
#     "set groups top firewall family inet filter protect_re term dns_proxy then accept",
#     "insert groups top firewall family inet filter protect_re term dns_proxy before term otherwise",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# DNS Proxy with second loopback IP",
#     "# DNS Proxy enabled in lo0 and using DNAT towards lo0. DNS Requested sent over Prisma Tunnels to OpenDNS 208.67.222.222",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "",
#     "set groups DNS_PROXY system services dns dns-proxy interface lo0.0",
#     "set groups DNS_PROXY system services dns forwarders 208.67.222.222",
#     "set groups DNS_PROXY system services dns forwarders 208.67.220.220",
#     "set groups DNS_PROXY routing-options static route 10.220.93.0/25 next-table LAN.inet.0",
#     "set groups DNS_PROXY routing-options static route 208.67.222.222/32 next-table apbr_TO-PRISMA-SSE.inet.0",
#     "set groups DNS_PROXY routing-options static route 208.67.220.220/32 next-table apbr_TO-PRISMA-SSE.inet.0",
#     "set apply-groups DNS_PROXY",
#     "set security nat destination pool pool1 address 100.10.0.1/32",
#     "set security nat destination rule-set qsrsoft from zone CORE_VLAN410",
#     "set security nat destination rule-set qsrsoft rule r1 match source-address 10.220.93.23/32",
#     "set security nat destination rule-set qsrsoft rule r1 match source-address 10.220.93.24/32",
#     "set security nat destination rule-set qsrsoft rule r1 match destination-address 10.220.93.62/32",
#     "set security nat destination rule-set qsrsoft rule r1 match destination-port 53",
#     "set security nat destination rule-set qsrsoft rule r1 then destination-nat pool pool1",
#     "set groups top routing-instances LAN routing-options instance-import lo0_to_LAN",
#     "set groups DNS_PROXY policy-options policy-statement lo0_to_LAN term lo0 from instance master",
#     "set groups DNS_PROXY policy-options policy-statement lo0_to_LAN term lo0 from interface lo0.0",
#     "set groups DNS_PROXY policy-options policy-statement lo0_to_LAN term lo0 then accept",
#     "set groups DNS_PROXY policy-options policy-statement lo0_to_LAN term 2 then reject",
#     "set security nat source rule-set host_to_Prisma from zone junos-host",
#     "set security nat source rule-set host_to_Prisma to zone tun_Prisma-Tunnel",
#     "set security nat source rule-set host_to_Prisma to zone tun_Prisma-Tunnel-II",
#     "set security nat source rule-set host_to_Prisma to zone tun_Prisma-Tunnel-III",
#     "set security nat source rule-set host_to_Prisma to zone tun_Prisma-Tunnel-IV",
#     "set security nat source rule-set host_to_Prisma rule rule1 match source-address 100.10.0.1/32",
#     "set security nat source rule-set host_to_Prisma rule rule1 match destination-address 208.67.222.222/32",
#     "set security nat source rule-set host_to_Prisma rule rule1 then source-nat interface",
#     "set security nat source rule-set host_to_Prisma rule rule2 match source-address 100.10.0.1/32",
#     "set security nat source rule-set host_to_Prisma rule rule2 match destination-address 208.67.220.220/32",
#     "set security nat source rule-set host_to_Prisma rule rule2 then source-nat interface",
#     "set groups DNS_PROXY routing-instances tun_Prisma-Tunnel   routing-options static route 100.10.0.1/32 next-table inet.0",
#     "set groups DNS_PROXY routing-instances tun_Prisma-Tunnel-II  routing-options static route 100.10.0.1/32 next-table inet.0",
#     "set groups DNS_PROXY routing-instances tun_Prisma-Tunnel-III routing-options static route 100.10.0.1/32 next-table inet.0",
#     "set groups DNS_PROXY routing-instances tun_Prisma-Tunnel-IV  routing-options static route 100.10.0.1/32 next-table inet.0",
#     "set security zones security-zone CORE_VLAN410 host-inbound-traffic system-services dns",
#     "set groups top firewall family inet filter protect_re term dns_proxy from protocol udp",
#     "set groups top firewall family inet filter protect_re term dns_proxy from port 53",
#     "set groups top firewall family inet filter protect_re term dns_proxy then accept",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# Additional Loopback address 100.10.0.1 - lower than 100.100.0.X",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "",
#     "set groups top interfaces lo0 unit 0 family inet address 100.10.0.1/32",
#     "",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "# Modify Protect_RE filter to permit DNS:",
#     "# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ #",
#     "",
#     "set groups top firewall family inet filter protect_re term dns_proxy from protocol udp",
#     "set groups top firewall family inet filter protect_re term dns_proxy from port 53",
#     "set groups top firewall family inet filter protect_re term dns_proxy then accept",
#     "insert groups top firewall family inet filter protect_re term dns_proxy before term otherwise",
#     ""
#   ]
#   dhcpd_config = {
#     AUTOMATION_VLAN460 = {
#       dns_servers = [
#         "{{public_dns3}}",
#         "{{public_dns4}}"
#       ],
#       gateway            = "{{branch_ip}}.254"
#       ip_end             = "{{branch_ip}}.241"
#       ip_start           = "{{branch_ip}}.225"
#       lease_time         = 86400
#       server_id_override = false
#       type               = "local"
#       type6              = "none"
#     },
#     CORE_VLAN410 = {
#       dns_servers = [
#         "{{public_dns3}}",
#         "{{public_dns4}}"
#       ],
#       fixed_bindings = {
#         "0007328b5e13" = {
#           ip   = "10.220.93.89"
#           name = "MAC-1"
#         },
#         "00155d3ac800" = {
#           ip   = "10.220.93.25"
#           name = "MAC-2"
#         },
#         "00155d719800" = {
#           ip   = "10.220.93.23"
#           name = "MAC-3"
#         },
#         "002324d0f216" = {
#           ip   = "10.220.93.81"
#           name = "MAC-4"
#         },
#         "047bcb07abdd" = {
#           ip   = "10.220.93.2"
#           name = "MAC-5"
#         },
#         "047bcb0832a8" = {
#           ip   = "10.220.93.1"
#           name = "MAC-6"
#         },
#         "088fc3a3583e" = {
#           ip   = "10.220.93.85"
#           name = "MAC-7"
#         },
#         "088fc3a4c940" = {
#           ip   = "10.220.93.51"
#           name = "MAC-8"
#         },
#         "282986883174" = {
#           ip   = "10.220.93.56"
#           name = "MAC-9"
#         },
#         "300ed53e5d80" = {
#           ip   = "10.220.93.15"
#           name = "MAC-10"
#         },
#         "300ed53e5f33" = {
#           ip   = "10.220.93.14"
#           name = "MAC-11"
#         },
#         "300ed53e6016" = {
#           ip   = "10.220.93.13"
#           name = "MAC-12"
#         },
#         "300ed53e6218" = {
#           ip   = "10.220.93.12"
#           name = "MAC-13"
#         },
#         "843a5b1756c9" = {
#           ip   = "10.220.93.7"
#           name = "MAC-14"
#         },
#         "843a5b17ab18" = {
#           ip   = "10.220.93.6"
#           name = "MAC-15"
#         },
#         "843a5b1dcdf3" = {
#           ip   = "10.220.93.9"
#           name = "MAC-16"
#         },
#         "847bebd29657" = {
#           ip   = "10.220.93.10"
#           name = "MAC-17"
#         },
#         "847bebd29708" = {
#           ip   = "10.220.93.8"
#           name = "MAC-18"
#         },
#         "a4601193bea0" = {
#           ip   = "10.220.93.103"
#           name = "MAC-19"
#         },
#         "e04f4399641d" = {
#           ip   = "10.220.93.36"
#           name = "MAC-20"
#         },
#         "e0be03434ba0" = {
#           ip   = "10.220.93.34"
#           name = "MAC-21"
#         },
#         "e0be03434c0b" = {
#           ip   = "10.220.93.32"
#           name = "MAC-22"
#         },
#         "e0be03434c70" = {
#           ip   = "10.220.93.33"
#           name = "MAC-23"
#         },
#         "e0be0347f6c4" = {
#           ip   = "10.220.93.31"
#           name = "MAC-24"
#         },
#         "e0be03692b3b" = {
#           ip   = "10.220.93.69"
#           name = "MAC-25"
#         },
#         "e0be03692dfe" = {
#           ip   = "10.220.93.67"
#           name = "MAC-26"
#         },
#         "e0be0369ff5f" = {
#           ip   = "10.220.93.35"
#           name = "MAC-27"
#         },
#         "e0be0369ffde" = {
#           ip   = "10.220.93.68"
#           name = "MAC-28"
#         },
#         "e0be036a114b" = {
#           ip   = "10.220.93.41"
#           name = "MAC-29"
#         },
#         "e0be036a1158" = {
#           ip   = "10.220.93.40"
#           name = "MAC-30"
#         },
#         "e805dc6e94f6" = {
#           ip   = "10.220.93.105"
#           name = "MAC-31"
#         },
#         "e805dc8c9110" = {
#           ip   = "10.220.93.100"
#           name = "MAC-32"
#         },
#         "e805dc8cec4f" = {
#           ip   = "10.220.93.97"
#           name = "MAC-33"
#         },
#         "e805dc8cec59" = {
#           ip   = "10.220.93.96"
#           name = "MAC-34"
#         },
#         "e805dc8cecf8" = {
#           ip   = "10.220.93.98"
#           name = "MAC-35"
#         },
#         "e805dc8ced99" = {
#           ip   = "10.220.93.92"
#           name = "MAC-36"
#         },
#         "e805dc8cedab" = {
#           ip   = "10.220.93.101"
#           name = "MAC-37"
#         },
#         "e805dc8cedce" = {
#           ip   = "10.220.93.99"
#           name = "MAC-38"
#         }
#       },
#       gateway            = "{{branch_ip}}.62"
#       ip_end             = "{{branch_ip}}.126"
#       ip_start           = "{{branch_ip}}.1"
#       lease_time         = 86400
#       server_id_override = false
#       type               = "local"
#       type6              = "none"
#     },
#     CREW_VLAN451 = {
#       dns_servers = [
#         "{{public_dns1}}",
#         "{{public_dns2}}"
#       ],
#       gateway            = "{{crew_network}}.222"
#       ip_end             = "{{crew_network}}.221"
#       ip_start           = "{{crew_network}}.194"
#       lease_time         = 86400
#       server_id_override = false
#       type               = "local"
#       type6              = "none"
#     },
#     IOT_WIRED_VLAN492 = {
#       dns_servers = [
#         "{{public_dns3}}",
#         "{{public_dns4}}"
#       ],
#       gateway            = "{{iot_net_wired}}.126"
#       ip_end             = "{{iot_net_wired}}.125"
#       ip_start           = "{{iot_net_wired}}.2"
#       lease_time         = 86400
#       server_id_override = false
#       type               = "local"
#       type6              = "none"
#     },
#     IOT_WIRELESS_VLAN494 = {
#       dns_servers = [
#         "{{public_dns3}}",
#         "{{public_dns4}}"
#       ],
#       gateway            = "{{iot_net_wireless}}.254"
#       ip_end             = "{{iot_net_wireless}}.253"
#       ip_start           = "{{iot_net_wireless}}.130"
#       lease_time         = 86400
#       server_id_override = false
#       type               = "local"
#       type6              = "none"
#     },
#     MANAGEMENT_VLAN430 = {
#       dns_servers = [
#         "{{public_dns3}}",
#         "{{public_dns4}}"
#       ],
#       fixed_bindings = {
#         "00012e832b1b" = {
#           ip   = "10.220.93.167"
#           name = "MAC-1"
#         },
#         "00012e832b6a" = {
#           ip   = "10.220.93.169"
#           name = "MAC-2"
#         },
#         "00012e832c42" = {
#           ip   = "10.220.93.171"
#           name = "MAC-3"
#         },
#         "00012e832c76" = {
#           ip   = "10.220.93.168"
#           name = "MAC-4"
#         },
#         "00012e832c80" = {
#           ip   = "10.220.93.172"
#           name = "MAC-5"
#         },
#         "00012e833286" = {
#           ip   = "10.220.93.170"
#           name = "MAC-6"
#         },
#         "00012ea04369" = {
#           ip   = "10.220.93.163"
#           name = "MAC-7"
#         },
#         "00012ea386cf" = {
#           ip   = "10.220.93.162"
#           name = "MAC-8"
#         }
#       },
#       gateway            = "{{branch_ip}}.161"
#       ip_end             = "{{branch_ip}}.190"
#       ip_start           = "{{branch_ip}}.162"
#       lease_time         = 86400
#       server_id_override = false
#       type               = "local"
#       type6              = "none"
#     },
#     MGMT = {
#       dns_servers = [
#         "{{public_dns1}}",
#         "{{public_dns2}}"
#       ],
#       gateway            = "192.168.2.1"
#       ip_end             = "192.168.2.100"
#       ip_start           = "192.168.2.2"
#       lease_time         = 86400
#       server_id_override = false
#       type               = "local"
#       type6              = "none"
#     },
#     PRIVATE_VLAN491 = {
#       dns_servers = [
#         "{{public_dns3}}",
#         "{{public_dns4}}"
#       ],
#       gateway            = "{{private_net}}.190"
#       ip_end             = "{{private_net}}.189"
#       ip_start           = "{{private_net}}.130"
#       lease_time         = 86400
#       server_id_override = false
#       type               = "local"
#       type6              = "none"
#     },
#     PUBLIC_VLAN450 = {
#       dns_servers = [
#         "{{public_dns1}}",
#         "{{public_dns2}}"
#       ],
#       gateway            = "{{public_net}}.1"
#       ip_end             = "{{public_net_upper}}.254"
#       ip_start           = "{{public_net}}.2"
#       lease_time         = 86400
#       server_id_override = false
#       type               = "local"
#       type6              = "none"
#     },
#     QUARANTINE_VLAN495 = {
#       gateway            = "{{quarantine_net}}.254"
#       ip_end             = "{{quarantine_net}}.253"
#       ip_start           = "{{quarantine_net}}.226"
#       lease_time         = 86400
#       server_id_override = false
#       type               = "local"
#       type6              = "none"
#     },
#     RESTRICTED_VLAN420 = {
#       dns_servers = [
#         "{{public_dns3}}",
#         "{{public_dns4}}"
#       ],
#       gateway            = "{{branch_ip}}.129"
#       ip_end             = "{{branch_ip}}.158"
#       ip_start           = "{{branch_ip}}.130"
#       lease_time         = 86400
#       server_id_override = false
#       type               = "local"
#       type6              = "none"
#     },
#     UNTRUSTED_VLAN490 = {
#       dns_servers = [
#         "{{public_dns1}}",
#         "{{public_dns2}}"
#       ],
#       gateway            = "{{untrusted_net}}.126"
#       ip_end             = "{{untrusted_net}}.125"
#       ip_start           = "{{untrusted_net}}.2"
#       lease_time         = 86400
#       server_id_override = false
#       type               = "local"
#       type6              = "none"
#     }
#   }
#   dns_servers = [
#     "8.8.8.8",
#     "8.8.4.4"
#   ]
#   ip_configs = {
#     AUTOMATION_VLAN460 = {
#       ip      = "{{branch_ip}}.254"
#       netmask = "/27"
#       type    = "static"
#     },
#     CORE_VLAN410 = {
#       ip      = "{{branch_ip}}.62"
#       netmask = "/25"
#       type    = "static"
#     },
#     CREW_VLAN451 = {
#       ip      = "{{crew_network}}.222"
#       netmask = "/27"
#       secondary_ips = [
#         "{{secondary_crew_network}}.1/24"
#       ],
#       type = "static"
#     },
#     IOT_WIRED_VLAN492 = {
#       ip      = "{{iot_net_wired}}.126"
#       netmask = "/25"
#       secondary_ips = [
#         "{{secondary_iot_wired}}.1/23"
#       ],
#       type = "static"
#     },
#     IOT_WIRELESS_VLAN494 = {
#       ip      = "{{iot_net_wireless}}.254"
#       netmask = "/25"
#       secondary_ips = [
#         "{{secondary_iot_net_wireless}}.1/23"
#       ],
#       type = "static"
#     },
#     MANAGEMENT_VLAN430 = {
#       ip      = "{{branch_ip}}.161"
#       netmask = "/27"
#       type    = "static"
#     },
#     MGMT = {
#       ip      = "192.168.2.1"
#       netmask = "/24"
#       type    = "static"
#     },
#     PCI_VLAN440 = {
#       ip      = "{{branch_ip}}.193"
#       netmask = "/27"
#       type    = "static"
#     },
#     PRIVATE_VLAN491 = {
#       ip      = "{{private_net}}.190"
#       netmask = "/26"
#       secondary_ips = [
#         "{{secondary_private_net}}.1/24"
#       ],
#       type = "static"
#     },
#     PUBLIC_VLAN450 = {
#       ip      = "{{public_net}}.1"
#       netmask = "/23"
#       type    = "static"
#     },
#     QUARANTINE_VLAN495 = {
#       ip      = "{{quarantine_net}}.254"
#       netmask = "/27"
#       type    = "static"
#     },
#     RESTRICTED_VLAN420 = {
#       ip      = "{{branch_ip}}.129"
#       netmask = "/27"
#       type    = "static"
#     },
#     UNTRUSTED_VLAN490 = {
#       ip      = "{{untrusted_net}}.126"
#       netmask = "/25"
#       secondary_ips = [
#         "{{secondary_untrusted_net_full_upper}}/26"
#       ],
#       type = "static"
#     }
#   }
#   ntp_servers = [
#     "129.6.15.28",
#     "129.6.15.29",
#     "129.6.15.30"
#   ]
#   oob_ip_config = {
#     type         = "dhcp"
#     use_mgmt_vrf = false
#     node1 = {
#       type = "dhcp"
#     }
#   }
#   path_preferences = {
#     ANY-TO-ANY = {
#       paths = [
#         {
#           networks = [
#             "CORE_VLAN410"
#           ],
#           type = "local"
#         },
#         {
#           networks = [
#             "RESTRICTED_VLAN420"
#           ],
#           type = "local"
#         },
#         {
#           networks = [
#             "MANAGEMENT_VLAN430"
#           ],
#           type = "local"
#         },
#         {
#           networks = [
#             "PCI_VLAN440"
#           ],
#           type = "local"
#         },
#         {
#           networks = [
#             "UNTRUSTED_VLAN490"
#           ],
#           type = "local"
#         },
#         {
#           networks = [
#             "PRIVATE_VLAN491"
#           ],
#           type = "local"
#         },
#         {
#           networks = [
#             "IOT_WIRED_VLAN492"
#           ],
#           type = "local"
#         },
#         {
#           networks = [
#             "IOT_WIRELESS_VLAN494"
#           ],
#           type = "local"
#         },
#         {
#           networks = [
#             "PUBLIC_VLAN450"
#           ],
#           type = "local"
#         },
#         {
#           networks = [
#             "CREW_VLAN451"
#           ],
#           type = "local"
#         },
#         {
#           name = "Prisma-Tunnel"
#           type = "tunnel"
#         },
#         {
#           name = "Prisma-Tunnel-II"
#           type = "tunnel"
#         },
#         {
#           name = "Prisma-Tunnel-III"
#           type = "tunnel"
#         },
#         {
#           name = "Prisma-Tunnel-IV"
#           type = "tunnel"
#         },
#         {
#           name = "WAN_0"
#           type = "wan"
#         },
#         {
#           name = "WAN_1"
#           type = "wan"
#         }
#       ],
#       strategy = "ordered"
#     },
#     TO-CORE = {
#       paths = [
#         {
#           networks = [
#             "CORE_VLAN410"
#           ],
#           type = "local"
#         }
#       ],
#       strategy = "ordered"
#     },
#     TO-INTERNET = {
#       paths = [
#         {
#           name = "WAN_0"
#           type = "wan"
#         },
#         {
#           name = "WAN_1"
#           type = "wan"
#         }
#       ],
#       strategy = "ordered"
#     },
#     TO-IOT_WIRED = {
#       paths = [
#         {
#           networks = [
#             "IOT_WIRED_VLAN492"
#           ],
#           type = "local"
#         }
#       ],
#       strategy = "ordered"
#     },
#     TO-IOT_WIRELESS = {
#       paths = [
#         {
#           networks = [
#             "IOT_WIRELESS_VLAN494"
#           ],
#           type = "local"
#         }
#       ],
#       strategy = "ordered"
#     },
#     TO-MANAGEMENT = {
#       paths = [
#         {
#           networks = [
#             "MANAGEMENT_VLAN430"
#           ],
#           type = "local"
#         }
#       ],
#       strategy = "ordered"
#     },
#     TO-PRIO-CELL = {
#       paths = [
#         {
#           name = "WAN_1"
#           type = "wan"
#         },
#         {
#           name = "WAN_0"
#           type = "wan"
#         }
#       ],
#       strategy = "ordered"
#     },
#     TO-PRISMA-SSE = {
#       paths = [
#         {
#           name = "Prisma-Tunnel"
#           type = "tunnel"
#         },
#         {
#           name = "Prisma-Tunnel-II"
#           type = "tunnel"
#         },
#         {
#           name = "Prisma-Tunnel-III"
#           type = "tunnel"
#         },
#         {
#           name = "Prisma-Tunnel-IV"
#           type = "tunnel"
#         },
#         {
#           name = "WAN_0"
#           type = "wan"
#         },
#         {
#           name = "WAN_1"
#           type = "wan"
#         }
#       ],
#       strategy = "ordered"
#     },
#     TO-PRIVATE = {
#       paths = [
#         {
#           networks = [
#             "PRIVATE_VLAN491"
#           ],
#           type = "local"
#         }
#       ],
#       strategy = "ordered"
#     },
#     TO-RESTRICTED = {
#       paths = [
#         {
#           networks = [
#             "RESTRICTED_VLAN420"
#           ],
#           type = "local"
#         }
#       ],
#       strategy = "ordered"
#     },
#     TO-UNTRUSTED = {
#       paths = [
#         {
#           networks = [
#             "UNTRUSTED_VLAN490"
#           ],
#           type = "local"
#         }
#       ],
#       strategy = "ordered"
#     }
#   }
#   port_config = {
#     "ge-0/0/0,ge-5/0/0" = {
#       ae_disable_lacp  = false
#       ae_lacp_force_up = false
#       aggregated       = false
#       critical         = true
#       description      = "WAN_0 First WAN Interface"
#       disable_autoneg  = false
#       disabled         = false
#       dsl_type         = "vdsl"
#       dsl_vci          = 35
#       dsl_vpi          = 0
#       duplex           = "auto"
#       ip_config = {
#         pppoe_auth = "none"
#         type       = "dhcp"
#       },
#       lte_auth           = "none"
#       name               = "WAN_0"
#       poe_disabled       = false
#       preserve_dscp      = true
#       redundant          = true
#       reth_idx           = 1
#       reth_node          = "node0"
#       speed              = "auto"
#       ssr_no_virtual_mac = false
#       svr_port_range     = "none"
#       usage              = "wan"
#       wan_arp_policer    = "default"
#       wan_extra_routes   = {}
#       wan_source_nat     = {}
#       wan_type           = "broadband"
#       vpn_paths          = {}
#     },
#     "ge-0/0/3,ge-5/0/3" = {
#       ae_disable_lacp  = false
#       ae_lacp_force_up = false
#       aggregated       = false
#       critical         = true
#       description      = "WAN_1 Second WAN Interface"
#       disable_autoneg  = false
#       disabled         = false
#       dsl_type         = "vdsl"
#       dsl_vci          = 35
#       dsl_vpi          = 0
#       duplex           = "auto"
#       ip_config = {
#         pppoe_auth = "none"
#         type       = "dhcp"
#       },
#       lte_auth           = "none"
#       name               = "WAN_1"
#       poe_disabled       = false
#       preserve_dscp      = true
#       redundant          = true
#       reth_idx           = 2
#       reth_node          = "node0"
#       speed              = "auto"
#       ssr_no_virtual_mac = false
#       svr_port_range     = "none"
#       usage              = "wan"
#       wan_arp_policer    = "default"
#       wan_extra_routes   = {}
#       wan_source_nat     = {}
#       wan_type           = "broadband"
#       vpn_paths          = {}
#     },
#     "ge-0/0/6-7,ge-5/0/6-7" = {
#       ae_disable_lacp  = false
#       ae_lacp_force_up = false
#       aggregated       = false
#       critical         = false
#       disable_autoneg  = false
#       disabled         = false
#       dsl_type         = "vdsl"
#       dsl_vci          = 35
#       dsl_vpi          = 0
#       duplex           = "auto"
#       lte_auth         = "none"
#       networks = [
#         "IOT_WIRED_VLAN492",
#         "RESTRICTED_VLAN420",
#         "PUBLIC_VLAN450",
#         "MANAGEMENT_VLAN430",
#         "CREW_VLAN451",
#         "CORE_VLAN410",
#         "PCI_VLAN440",
#         "QUARANTINE_VLAN495",
#         "UNTRUSTED_VLAN490",
#         "PRIVATE_VLAN491",
#         "IOT_WIRELESS_VLAN494",
#         "AUTOMATION_VLAN460"
#       ],
#       poe_disabled       = false
#       port_network       = "MGMT"
#       preserve_dscp      = true
#       redundant          = true
#       reth_idx           = 3
#       reth_node          = "node0"
#       speed              = "auto"
#       ssr_no_virtual_mac = false
#       svr_port_range     = "none"
#       usage              = "lan"
#       vpn_paths          = {}
#       wan_arp_policer    = "default"
#       wan_extra_routes   = {}
#       wan_type           = "broadband"
#     }
#   }
#   service_policies = [
#     {
#       servicepolicy_id = mist_org_servicepolicy.servicepolicy_196.id
#       path_preference  = "TO-INTERNET"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_154.id
#       path_preference  = "TO-RESTRICTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_120.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_122.id
#       path_preference  = "TO-PRIVATE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_125.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_5.id
#       path_preference  = "TO-IOT_WIRELESS"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_127.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_110.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_158.id
#       path_preference  = "TO-RESTRICTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_166.id
#       path_preference  = "TO-RESTRICTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_161.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_174.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_180.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_184.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_191.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_200.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_54.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_56.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_59.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_14.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_135.id
#       path_preference  = "TO-RESTRICTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_3.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_171.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_11.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_204.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_0.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_8.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_201.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_7.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_175.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_94.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_21.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_24.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_27.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_30.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_34.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_39.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_43.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_48.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_12.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_52.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_55.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_31.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_17.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_57.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_61.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_63.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_65.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_68.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_70.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_72.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_76.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_78.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_181.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_73.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_80.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_82.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_84.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_88.id
#       path_preference  = "TO-UNTRUSTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_90.id
#       path_preference  = "TO-UNTRUSTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_92.id
#       path_preference  = "TO-UNTRUSTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_16.id
#       path_preference  = "TO-UNTRUSTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_35.id
#       path_preference  = "TO-UNTRUSTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_74.id
#       path_preference  = "TO-UNTRUSTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_137.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_139.id
#       path_preference  = "TO-RESTRICTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_144.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_149.id
#       path_preference  = "TO-PRIVATE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_130.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_123.id
#       path_preference  = "TO-IOT_WIRELESS"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_207.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_203.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_138.id
#       path_preference  = "TO-RESTRICTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_140.id
#       path_preference  = "TO-RESTRICTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_2.id
#       path_preference  = "TO-RESTRICTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_147.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_151.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_155.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_177.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_159.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_162.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_168.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_172.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_176.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_182.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_185.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_188.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_169.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_193.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_198.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_202.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_205.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_1.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_4.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_9.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_142.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_189.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_18.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_25.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_32.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_36.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_41.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_47.id
#       path_preference  = "TO-INTERNET"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_6.id
#       path_preference  = "TO-INTERNET"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_44.id
#       path_preference  = "TO-INTERNET"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_190.id
#       path_preference  = "TO-INTERNET"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_53.id
#       path_preference  = "TO-INTERNET"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_69.id
#       path_preference  = "TO-INTERNET"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_37.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_75.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_132.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_145.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_113.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_79.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_45.id
#       path_preference  = "TO-RESTRICTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_89.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_91.id
#       path_preference  = "TO-PRIVATE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_93.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_95.id
#       path_preference  = "TO-IOT_WIRELESS"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_115.id
#       path_preference  = "TO-INTERNET"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_194.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_38.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_111.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_112.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_157.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_160.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_29.id
#       path_preference  = "TO-UNTRUSTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_33.id
#       path_preference  = "TO-UNTRUSTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_179.id
#       path_preference  = "TO-UNTRUSTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_85.id
#       path_preference  = "TO-UNTRUSTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_96.id
#       path_preference  = "TO-UNTRUSTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_98.id
#       path_preference  = "TO-UNTRUSTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_100.id
#       path_preference  = "TO-UNTRUSTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_40.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_86.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_103.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_105.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_164.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_170.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_101.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_108.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_114.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_117.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_119.id
#       path_preference  = "TO-RESTRICTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_121.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_58.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_124.id
#       path_preference  = "TO-IOT_WIRELESS"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_197.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_167.id
#       path_preference  = "TO-RESTRICTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_192.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_126.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_129.id
#       path_preference  = "TO-IOT_WIRELESS"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_131.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_134.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_173.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_178.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_136.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_13.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_22.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_28.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_163.id
#       path_preference  = "TO-RESTRICTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_206.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_152.id
#       path_preference  = "TO-PRIVATE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_50.id
#       path_preference  = "TO-IOT_WIRELESS"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_141.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_49.id
#       path_preference  = "TO-RESTRICTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_19.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_118.id
#       path_preference  = "TO-PRIVATE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_107.id
#       path_preference  = "TO-IOT_WIRELESS"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_15.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_20.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_23.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_26.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_195.id
#       path_preference  = "TO-IOT_WIRELESS"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_42.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_133.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_128.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_10.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_165.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_146.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_116.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_150.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_153.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_156.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_60.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_62.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_64.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_66.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_51.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_71.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_148.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_183.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_186.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_77.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_81.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_83.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_87.id
#       path_preference  = "TO-RESTRICTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_143.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_67.id
#       path_preference  = "TO-PRIVATE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_97.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_99.id
#       path_preference  = "TO-CORE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_102.id
#       path_preference  = "TO-RESTRICTED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_104.id
#       path_preference  = "TO-MANAGEMENT"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_106.id
#       path_preference  = "TO-PRIVATE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_109.id
#       path_preference  = "TO-IOT_WIRED"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_199.id
#       path_preference  = "ANY-TO-ANY"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_46.id
#       path_preference  = "TO-PRISMA-SSE"
#     },
#     { servicepolicy_id = mist_org_servicepolicy.servicepolicy_187.id
#       path_preference  = "TO-INTERNET"
#     }
#   ]
#   tunnel_configs = {
#     Prisma-Tunnel = {
#       ike_lifetime = 28800
#       ike_mode     = "main"
#       ike_proposals = [
#         {
#           auth_algo = "sha2"
#           dh_group  = "19"
#           enc_algo  = "aes256"
#         }
#       ],
#       ipsec_lifetime = 3600
#       ipsec_proposals = [
#         {
#           auth_algo = "sha2"
#           dh_group  = "19"
#           enc_algo  = "aes256"
#         }
#       ],
#       local_id = "{{sc_local_id_1}}"
#       mode     = "active-standby"
#       primary = {
#         hosts = [
#           "{{service_ip_1}}"
#         ],
#         remote_ids = [
#           "{{sc_remote_id_1}}"
#         ],
#         wan_names = [
#           "WAN_0"
#         ]
#       },
#       protocol = "ipsec"
#       provider = "custom-ipsec"
#       psk      = "IEpg6tpOMGLBauoT5nodeEacximEpF"
#       version  = "2"
#     },
#     Prisma-Tunnel-II = {
#       ike_lifetime = 28800
#       ike_mode     = "main"
#       ike_proposals = [
#         {
#           auth_algo = "sha2"
#           dh_group  = "19"
#           enc_algo  = "aes256"
#         }
#       ],
#       ipsec_lifetime = 3600
#       ipsec_proposals = [
#         {
#           auth_algo = "sha2"
#           dh_group  = "19"
#           enc_algo  = "aes256"
#         }
#       ],
#       local_id = "{{sc_local_id_2}}"
#       mode     = "active-standby"
#       primary = {
#         hosts = [
#           "{{service_ip_1}}"
#         ],
#         remote_ids = [
#           "{{sc_remote_id_2}}"
#         ],
#         wan_names = [
#           "WAN_1"
#         ]
#       },
#       protocol = "ipsec"
#       provider = "custom-ipsec"
#       psk      = "IEpg6tpOMGLBauoT5nodeEacximEpF"
#       version  = "2"
#     },
#     Prisma-Tunnel-III = {
#       ike_lifetime = 28800
#       ike_mode     = "main"
#       ike_proposals = [
#         {
#           auth_algo = "sha2"
#           dh_group  = "19"
#           enc_algo  = "aes256"
#         }
#       ],
#       ipsec_lifetime = 3600
#       ipsec_proposals = [
#         {
#           auth_algo = "sha2"
#           dh_group  = "19"
#           enc_algo  = "aes256"
#         }
#       ],
#       local_id = "{{sc_local_id_1}}"
#       mode     = "active-standby"
#       primary = {
#         hosts = [
#           "{{service_ip_2}}"
#         ],
#         remote_ids = [
#           "{{sc_remote_id_1}}"
#         ],
#         wan_names = [
#           "WAN_0"
#         ]
#       },
#       protocol = "ipsec"
#       provider = "custom-ipsec"
#       psk      = "IEpg6tpOMGLBauoT5nodeEacximEpF"
#       version  = "2"
#     },
#     Prisma-Tunnel-IV = {
#       ike_lifetime = 28800
#       ike_mode     = "main"
#       ike_proposals = [
#         {
#           auth_algo = "sha2"
#           dh_group  = "19"
#           enc_algo  = "aes256"
#         }
#       ],
#       ipsec_lifetime = 3600
#       ipsec_proposals = [
#         {
#           auth_algo = "sha2"
#           dh_group  = "19"
#           enc_algo  = "aes256"
#         }
#       ],
#       local_id = "{{sc_local_id_2}}"
#       mode     = "active-standby"
#       primary = {
#         hosts = [
#           "{{service_ip_2}}"
#         ],
#         remote_ids = [
#           "{{sc_remote_id_2}}"
#         ],
#         wan_names = [
#           "WAN_1"
#         ]
#       },
#       protocol = "ipsec"
#       provider = "custom-ipsec"
#       psk      = "IEpg6tpOMGLBauoT5nodeEacximEpF"
#       version  = "2"
#     }
#   }
#   vrf_config = {
#     enabled = true
#   }
#   vrf_instances = {
#     LAN = {
#       networks = [
#         "CORE_VLAN410",
#         "RESTRICTED_VLAN420",
#         "MANAGEMENT_VLAN430",
#         "PRIVATE_VLAN491",
#         "PCI_VLAN440",
#         "IOT_WIRELESS_VLAN494",
#         "IOT_WIRED_VLAN492",
#         "PUBLIC_VLAN450",
#         "CREW_VLAN451",
#         "UNTRUSTED_VLAN490",
#         "QUARANTINE_VLAN495"
#       ]
#     },
#     MANAGEMENT = {
#       networks = [
#         "MGMT"
#       ]
#     }
#   }
# }
