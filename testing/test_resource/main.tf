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
  name = "Terraform Testing"
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
  mist_nac = {
    default_idp_id = mist_org_nacidp.idp_azure.id
    idps = [{
      id          = mist_org_nacidp.idp_azure.id
      user_realms = ["nacidpazure.com"]
    }]
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
  ui_idle_timeout = 0
}

resource "mist_org_alarmtemplate" "alarmtemplate_one" {
  org_id = mist_org.terraform_test.id
  name   = "alarmtemplate_one"
  delivery = {
    enabled           = true
    to_org_admins     = true
    additional_emails = ["tmunzer@juniper.net"]
  }
  rules = {
    dhcp_failure : {
      enabled : true
      delivery : {
        enabled : true
      }
    }
    health_check_failed : {
      enabled : true
    },
    insufficient_capacity : {
      enabled : true
    },
    insufficient_coverage : {
      enabled : true
    },
    infra_arp_failure : {
      enabled : true
    },
    arp_failure : {
      enabled : true
    },
    authentication_failure : {
      enabled : true
    },
    bad_wan_uplink : {
      enabled : true
    },
    bad_cable : {
      enabled : true
    },
    gw_bad_cable : {
      enabled : true
    },
    ap_bad_cable : {
      enabled : true
    },
    rogue_client : {
      enabled : true
    },
    infra_dhcp_failure : {
      enabled : true
    },
    infra_dns_failure : {
      enabled : true
    },
    dns_failure : {
      enabled : true
    },
    wan_device_problem : {
      enabled : true
    },
    missing_vlan : {
      enabled : true
    },
    mist_edge_cpu_usage_high : {
      enabled : true
    },
    mist_edge_disconnected : {
      enabled : true
    },
    mist_edge_disk_usage_high : {
      enabled : true
    },
    mist_edge_memory_usage_high : {
      enabled : true
    },
    mist_edge_powerinput_disconnected : {
      enabled : true
    },
    mist_edge_service_failed : {
      enabled : true
    },
    mist_edge_psu_unplugged : {
      enabled : true
    },
    negotiation_mismatch : {
      enabled : true
    },
    gw_negotiation_mismatch : {
      enabled : true
    },
    non_compliant : {
      enabled : true
    },
    ap_offline : {
      enabled : true
    },
    port_stuck : {
      enabled : true
    },
    rogue_ap : {
      enabled : true
    },
    switch_stp_loop : {
      enabled : true
    },
    vpn_path_down : {
      enabled : true
    },
    vc_backup_failed : {
      enabled : true
    },
    vc_master_changed : {
      enabled : true
    },
    vc_member_deleted : {
      enabled : true
    },
    sw_vc_port_down : {
      enabled : true
    },
    infra_arp_success : {
      enabled : true
    },
    air_magnet_scan : {
      enabled : true
    },
    sw_bgp_neighbor_state_changed : {
      enabled : true
    },
    sw_bgp_neighbor_up : {
      enabled : true
    },
    cellular_edge_connected_to_ncm : {
      enabled : true
    },
    cellular_edge_disconnected_from_ncm : {
      enabled : true
    },
    cellular_edge_firmware_upgraded : {
      enabled : true
    },
    cellular_edge_login_failed : {
      enabled : true
    },
    cellular_edge_login_succeeded : {
      enabled : true
    },
    cellular_edge_rebooted : {
      enabled : true
    },
    cellular_edge_sim_door_closed : {
      enabled : true
    },
    cellular_edge_sim_door_opened : {
      enabled : true
    },
    cellular_edge_modem_wan_connected : {
      enabled : true
    },
    cellular_edge_wan_service_type_changed : {
      enabled : true
    },
    cellular_edge_ethernet_wan_connected : {
      enabled : true
    },
    cellular_edge_ethernet_wan_plugged : {
      enabled : true
    },
    sw_critical_port_up : {
      enabled : true
    },
    gw_critical_port_up : {
      enabled : true
    },
    infra_dhcp_success : {
      enabled : true
    },
    infra_dns_success : {
      enabled : true
    },
    device_reconnected : {
      enabled : true
    },
    device_restarted : {
      enabled : true
    },
    eap_handshake_flood : {
      enabled : true
    },
    ha_control_link_up : {
      enabled : true
    },
    tt_inactive_vlans : {
      enabled : true
    },
    mist_edge_connected : {
      enabled : true
    },
    mist_edge_cpu_usage_normal : {
      enabled : true
    },
    mist_edge_disk_usage_normal : {
      enabled : true
    },
    mist_edge_memory_usage_normal : {
      enabled : true
    },
    mist_edge_psu_plugged : {
      enabled : true
    },
    mist_edge_powerinput_connected : {
      enabled : true
    },
    tt_tunnels_up : {
      enabled : true
    },
    switch_reconnected : {
      enabled : true
    },
    switch_restarted : {
      enabled : true
    },
    vpn_peer_up : {
      enabled : true
    },
    vc_member_added : {
      enabled : true
    },
    sw_vc_port_up : {
      enabled : true
    },
    gw_bgp_neighbor_up : {
      enabled : true
    },
    gateway_reconnected : {
      enabled : true
    },
    watched_station : {
      enabled : true
    },
    adhoc_network : {
      enabled : true
    },
    tt_ports_all_dropped_from_lacp : {
      enabled : true
    },
    tt_tunnels_lost : {
      enabled : true
    },
    sw_bgp_neighbor_down : {
      enabled : true
    },
    bssid_spoofing : {
      enabled : true
    },
    cellular_edge_modem_wan_disconnected : {
      enabled : true
    },
    cellular_edge_ethernet_wan_disconnected : {
      enabled : true
    },
    cellular_edge_ethernet_wan_unplugged : {
      enabled : true
    },
    sw_critical_port_down : {
      enabled : true
    },
    gw_critical_port_down : {
      enabled : true
    },
    device_down : {
      enabled : true
    },
    disassociation_flood : {
      enabled : true
    },
    eap_dictionary_attack : {
      enabled : true
    },
    eap_failure_injection : {
      enabled : true
    },
    eap_spoofed_success : {
      enabled : true
    },
    eapol_logoff_attack : {
      enabled : true
    },
    esl_hung : {
      enabled : true
    },
    esl_recovered : {
      enabled : true
    },
    essid_jack : {
      enabled : true
    },
    sw_evpn_duplicate_mac : {
      enabled : true
    },
    excessive_client : {
      enabled : true
    },
    excessive_eapol_start : {
      enabled : true
    },
    beacon_flood : {
      enabled : true
    },
    ha_control_link_down : {
      enabled : true
    },
    honeypot_ssid : {
      enabled : true
    },
    idp_attack_detected : {
      enabled : true
    },
    tt_port_last_dropped_from_lacp : {
      enabled : true
    },
    loop_detected_by_ap : {
      enabled : true
    },
    mist_edge_service_crashed : {
      enabled : true
    },
    monkey_jack : {
      enabled : true
    },
    out_of_sequence : {
      enabled : true
    },
    port_flap : {
      enabled : true
    },
    repeated_auth_failures : {
      enabled : true
    },
    krack_attack : {
      enabled : true
    },
    ssid_injection : {
      enabled : true
    },
    secpolicy_violation : {
      enabled : true
    },
    sw_bpdu_error : {
      enabled : true
    },
    sw_bad_optics : {
      enabled : true
    },
    sw_dhcp_pool_exhausted : {
      enabled : true
    },
    sw_alarm_chassis_hot : {
      enabled : true
    },
    sw_alarm_chassis_pem : {
      enabled : true
    },
    sw_alarm_chassis_poe : {
      enabled : true
    },
    sw_alarm_chassis_psu : {
      enabled : true
    },
    sw_alarm_chassis_partition : {
      enabled : true
    },
    switch_down : {
      enabled : true
    },
    tkip_icv_attack : {
      enabled : true
    },
    tunnel_down : {
      enabled : true
    },
    url_blocked : {
      enabled : true
    },
    vpn_peer_down : {
      enabled : true
    },
    vendor_ie_missing : {
      enabled : true
    },
    vc_member_restarted : {
      enabled : true
    },
    gw_bgp_neighbor_down : {
      enabled : true
    },
    gw_dhcp_pool_exhausted : {
      enabled : true
    },
    gw_src_nat_pool_threshold_exceeded : {
      enabled : true
    },
    gateway_down : {
      enabled : true
    },
    zero_ssid_association : {
      enabled : true
    }
  }
}
# resource "mist_site_psk" "psk_one" {
#   site_id    = mist_site.terraform_site.id
#   name       = "JNP-FR-PAR"
#   passphrase = "secretone"
#   ssid       = mist_site_wlan.wlan_one.ssid
#   usage      = "multi"
# }
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
  # sitegroup_ids = [mist_org_sitegroup.test_group.id, mist_org_sitegroup.test_group2.id]
  # networktemplate_id = mist_org_networktemplate.switch_template.id
  # rftemplate_id      = mist_org_rftemplate.test_rf.id
  # gatewaytemplate_id = mist_org_gatewaytemplate.test-api.id
}
resource "mist_org_inventory" "inventory" {
  # {
  #   claim_code = "C9QQFW2B9NKCS33"
  #   //site_id    = mist_site.terraform_site.id
  # },site_id    = mist_site.terraform_site.id
  # {
  #   mac     = "2c21311c37b0"
  #   site_id = mist_site.terraform_site.id
  # },
  org_id = mist_org.terraform_test.id
  # devices = [
  #   {
  #     claim_code = "CPKL2EXN8JY98AC"
  #     site_id    = mist_site.terraform_site.id
  #   },
  #   {
  #     claim_code             = "G87JHBFXZJSFNMX"
  #     site_id                = mist_site.terraform_site.id
  #     unclaim_when_destroyed = true
  #   },
  #   {
  #     claim_code = "CV4YAS8DQWYLL6M"
  #     site_id                = mist_site.terraform_site.id
  #   },
  #   {
  #     mac                    = (local.node0)
  #     site_id                = mist_site.terraform_site.id
  #     unclaim_when_destroyed = false
  #   },
  #   {
  #     mac                    = (local.node1)
  #     site_id                = mist_site.terraform_site.id
  #     unclaim_when_destroyed = false
  #   },
  #   {
  #     mac     = "4c9614026d00"
  #     site_id = mist_site.terraform_site.id
  #   },
  #   {
  #     mac     = "4c961418c000"
  #     site_id = mist_site.terraform_site.id
  #   },
  #   {
  #     mac     = "020004bbf189"
  #     //site_id = mist_site.terraform_site.id
  #   }
  # ]
  inventory = {
    "CPKL2EXN8JY98AC" = {
      site_id = mist_site.terraform_site.id
    }
    "G87JHBFXZJSFNMX" = {
      site_id                = mist_site.terraform_site.id
      unclaim_when_destroyed = true
    }
    "CV4YAS8DQWYLL6M" = {
      site_id = mist_site.terraform_site.id
    }
    (local.node0) = {
      site_id                = mist_site.terraform_site.id
      unclaim_when_destroyed = false
    }
    (local.node1) = {
      site_id                = mist_site.terraform_site.id
      unclaim_when_destroyed = false
    }
    "4c9614026d00" = {
      site_id = mist_site.terraform_site.id
    }
    "4c9614026d00" = {
      site_id                = mist_site.terraform_site.id
      unclaim_when_destroyed = false
    }
    "4c961418c000" = {
      site_id                = mist_site.terraform_site.id
      unclaim_when_destroyed = false
    }
  }
}



resource "mist_org_webhook" "webhook_one" {
  org_id  = mist_org.terraform_test.id
  name    = "test"
  url     = "https://test.com"
  topics  = ["audits"]
  enabled = false
  type    = "http-post"
}
resource "mist_site_webhook" "webhook_one" {
  site_id = mist_site.terraform_site.id
  name    = "test"
  url     = "https://test.com"
  topics  = ["audits"]
  enabled = false
  type    = "http-post"
}
resource "mist_org_rftemplate" "example" {
  org_id       = mist_org.terraform_test.id
  name         = "Integration_Template_Test"
  country_code = "US"
}
resource "mist_site_wlan" "wlan_one" {
  ssid    = "McDonalds Free WiFi"
  site_id = mist_site.terraform_site.id
  bands   = ["24", "5"]
  auth = {
    type = "open"
  }
  client_limit_down         = 5000
  client_limit_up           = 5000
  client_limit_down_enabled = true
  client_limit_up_enabled   = true
  isolation                 = true
  l2_isolation              = true
  portal = {
    forward     = true
    forward_url = "https://www.mcdonalds.com/us/en-us/wifi.html"
  }
  vlan_enabled            = true
  vlan_id                 = 450
  wlan_limit_down         = 100000
  wlan_limit_down_enabled = true
  wlan_limit_up           = 100000
  wlan_limit_up_enabled   = true
}

resource "mist_site_wlan" "wlan_two" {
  ssid    = "eBOS"
  site_id = mist_site.terraform_site.id
  auth = {
    enable_mac_auth = true
    key_idx         = 1
    type            = "open"
  }
  bands = ["24", "5"]
  mist_nac = {
    enabled = true
  }
  vlan_enabled = true
  vlan_id      = 451
}

resource "mist_site_wlan" "wlan_three" {
  ssid    = "McD-HHOT"
  site_id = mist_site.terraform_site.id
  auth = {
    type = "eap"
  }
  bands                   = ["24", "5"]
  block_blacklist_clients = true
  hide_ssid               = true
  limit_bcast             = true
  mist_nac = {
    enabled = true
  }
  vlan_enabled = true
  vlan_id      = 420
}

resource "mist_site_wlan" "wlan_four" {
  ssid    = "McD-IoT"
  site_id = mist_site.terraform_site.id
  auth = {
    multi_psk_only = true
    type           = "psk"
  }
  bands                   = ["24", "5"]
  block_blacklist_clients = true
  isolation               = true
  l2_isolation            = true
  limit_bcast             = true
  mist_nac = {
    enabled = true
  }
  vlan_enabled = true
  vlan_id      = 494
}


resource "mist_org_servicepolicy" "test1" {
  org_id = mist_org.terraform_test.id
  tenants = [
    "CORE_VLAN410"
  ]
  services = [
    "any"
  ]
  action = "allow"
  idp = {
    enabled = false
  }
  name = "Policy-1"
}
resource "mist_org_service" "service11" {
  org_id       = mist_org.terraform_test.id
  traffic_type = "default"
  app_categories = [
    "Advertisement",
    "FileSharing"
  ]
  name = "public_url_block"
  type = "app_categories"
}
resource "mist_org_servicepolicy" "test2" {
  org_id = mist_org.terraform_test.id
  tenants = [
    "MANAGEMENT_VLAN430"
  ]
  services = [
    "any"
  ]
  action = "allow"
  name   = "Policy-3"
}
resource "mist_org_servicepolicy" "test3" {
  org_id = mist_org.terraform_test.id
  tenants = [
    "cctv-b.PRD-IoT"
  ]
  services = [
    "any"
  ]
  action = "allow"
  name   = "Policy-2"
}
resource "mist_org_servicepolicy" "test4" {
  org_id = mist_org.terraform_test.id
  tenants = [
    "saltlake-lan2"
  ]
  services = [
    "Internet"
  ]
  action = "allow"
  name   = "UDP-ZERO-DATA-DNS-GOOGLE"
}
resource "mist_org_servicepolicy" "test5" {
  org_id = mist_org.terraform_test.id
  tenants = [
    "corp_datacenter1",
    "corp_datacenter2",
    "core_dc2_network",
    "mgmt_datacenter1",
    "mgmt_datacenter2",
    "corp_network",
    "mgmt_network"
  ]
  services = [
    "rfc1918"
  ]
  action = "allow"
  idp = {
    enabled = false
  }
  name = "access-to-rfc1918"
}
resource "mist_org_servicepolicy" "test6" {
  org_id = mist_org.terraform_test.id
  tenants = [
    "guest"
  ]
  services = [
    "guest-internet"
  ]
  action = "allow"
  idp = {
    enabled    = true,
    profile    = "standard",
    alert_only = true
  }
  name = "Guest-IDP"
}
resource "mist_org_service" "service143" {
  org_id      = mist_org.terraform_test.id
  description = "App for detecting office 365"
  apps = [
    "office365"
  ]
  traffic_type = "default"
  name         = "XX-office365"
  type         = "apps"
}

resource "mist_org_idpprofile" "test1" {
  org_id       = mist_org.terraform_test.id
  base_profile = "standard"
  overwrites = [
    {
      name   = "r1"
      action = "drop"
      matching = {
        severity   = []
        dst_subnet = []
        attack_name = [
          "SSL:OVERFLOW:KEY-ARG-NO-ENTROPY"
        ]
      }
    },
    {
      name = "r2"
      matching = {
        severity   = []
        dst_subnet = []
        attack_name = [
          "UDP:ZERO-DATA"
        ]
      }
    }
  ]
  name = "custom-standard"
}
resource "mist_org_nactag" "nactag_VLAN_Crew451" {
  org_id = mist_org.terraform_test.id
  vlan   = "451"
  name   = "VLAN-Crew451"
  type   = "vlan"
}
resource "mist_org_idpprofile" "test2" {
  org_id       = mist_org.terraform_test.id
  base_profile = "standard"
  overwrites = [
    {
      name   = "UDP-ZERO-DATA-AllowGoogle1"
      action = "alert"
      matching = {
        severity = [],
        dst_subnet = [
          "8.8.8.8/32"
        ],
        attack_name = [
          "UDP:ZERO-DATA"
        ]
      }
    }
  ]
  name = "UDP-ZERO-DATA"
}
resource "mist_org_idpprofile" "test3" {
  org_id       = mist_org.terraform_test.id
  base_profile = "critical"
  overwrites = [
    {
      name = "server_bypass"
      matching = {
        severity   = []
        dst_subnet = []
        attack_name = [
          "SSL:OVERFLOW:KEY-ARG-NO-ENTROPY"
        ]
      }
    },
    {
      name = "aide-demo-guest-bypass"
      matching = {
        severity   = []
        dst_subnet = []
        attack_name = [
          "SSL:OVERFLOW:KEY-ARG-NO-ENTROPY"
        ]
      }
    }
  ]
  name = "aide-demo-dmz-bypass"
}
resource "mist_org_idpprofile" "test4" {
  org_id       = mist_org.terraform_test.id
  base_profile = "critical"
  name         = "nn-test"
}
resource "mist_org_idpprofile" "test5" {
  org_id       = mist_org.terraform_test.id
  base_profile = "standard"
  overwrites = [
    {
      name   = "zero-data"
      action = "close"
      matching = {
        severity = [],
        dst_subnet = [
          "9.9.9.9/32",
          "1.0.0.1/32"
        ]
        attack_name = [
          "UDP:ZERO-DATA"
        ]
      }
    }
  ]
  name = "zero-data"
}
resource "mist_org_idpprofile" "test6" {
  org_id       = mist_org.terraform_test.id
  base_profile = "standard"
  overwrites = [
    {
      name   = "CVE"
      action = "alert"
      matching = {
        severity = []
        dst_subnet = [
          "192.168.94.2/32",
          "93.184.215.14/32",
          "13.89.179.10/32",
          "1.0.0.1/32",
          "9.9.9.9/32",
          "1.1.1.1/32",
          "142.251.33.99/32",
          "23.102.110.248/32",
          "20.189.173.10/32",
          "13.89.179.14/32",
          "142.251.211.227/32"
        ]
        attack_name = [
          "HTTP:INFO-LEAK:NO-SP-AFTER-FLD",
          "HTTP:OVERFLOW:HEADER",
          "HTTP:INVALID:HDR-FIELD",
          "HTTP:INVALID:INV-CONT-ENC",
          "SSL:AUDIT:EXPORT-GRADE-CIPHER",
          "PORTMAPPER:ERR:SHORT-READ",
          "SSL:OVERFLOW:KEY-ARG-NO-ENTROPY"
        ]
      }
    }
  ]
  name = "CVE"
}


resource "mist_org_wxtag" "test_1" {
  values = [
    "10.3.0.0/16"
  ]
  op     = "in"
  name   = "Internal Network"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "ip_range_subnet"
}
resource "mist_org_wxtag" "test_2" {
  values = [
    "53"
  ]
  op     = "in"
  name   = "DNS"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "port"
}
resource "mist_org_wxtag" "test_3" {
  values = [
    "443"
  ]
  op     = "in"
  name   = "HTTPS"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "port"
}
resource "mist_org_wxtag" "test_4" {
  values = [
    "80"
  ]
  op     = "in"
  name   = "HTTP"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "port"
}
resource "mist_org_wxtag" "test_5" {
  op = "in"
  values = [
    "www.myapp.com:443"
  ]
  name   = "my app"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "hostname"
}
resource "mist_org_wxtag" "test_6" {
  op = "in"
  values = [
    "test1"
  ]
  name   = "test1"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "radius_group"
}
resource "mist_org_wxtag" "test_7" {
  op = "in"
  values = [
    "10.3.20.0/24"
  ]
  name   = "Subnet SRV"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "ip_range_subnet"
}
resource "mist_org_wxtag" "test_8" {
  specs = [
    {
      subnets = [
        "10.3.20.105"
      ]
      protocol   = "tcp"
      port_range = "443-443"
    }
  ]
  name   = "IP traefik.stag.one:443"
  org_id = mist_org.terraform_test.id
  type   = "spec"
}
resource "mist_org_wxtag" "test_81" {
  specs = [
    {
      protocol = "udp"
      subnets  = ["0.0.0.0/0"]
    }
  ]
  name   = "test_81"
  org_id = mist_org.terraform_test.id
  type   = "spec"
}
resource "mist_org_wxtag" "test_9" {
  op = "in"
  values = [
    "10.3.20.105"
  ]
  name   = "IP traefik.stag.one"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "ip_range_subnet"
}
resource "mist_org_wxtag" "test_10" {
  name   = "ma station"
  org_id = mist_org.terraform_test.id
  type   = "client"
  mac    = "9a4fccaa8f82"
}
resource "mist_org_wxtag" "test_12" {
  op = "in"
  values = [
    "iot"
  ]
  name   = "iot_label"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "radius_group"
}
resource "mist_org_wxtag" "test_13" {
  op = "in"
  values = [
    "station"
  ]
  name   = "station"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "radius_group"
}
resource "mist_org_wxtag" "test_14" {
  op     = "not_in"
  name   = "AP_test"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "ap_id"
  values = [
    "0000000"
  ]
}
resource "mist_org_wxtag" "test_15" {
  op = "in"
  values = [
    "test"
  ]
  name   = "test"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "radius_group"
}
resource "mist_org_wxtag" "test_16" {
  values = [
    "test"
  ]
  name   = "VIP"
  op     = "in"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "psk_role"
}
resource "mist_org_wxtag" "test_16_1" {
  op = "in"
  values = [
    "10.3.20.164"
  ]
  name   = "vcsa.stag.one_ip"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "ip_range_subnet"
}
resource "mist_org_wxtag" "test_17" {
  op = "not_in"
  values = [
    "ringcentral",
    "gcp",
    "gsuite",
    "teamviewer",
    "ms-teams"
  ]
  name   = "Allowed Applications"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "app"
}
resource "mist_org_wxtag" "test_18" {
  op = "in"
  values = [
    "0.0.0.0/0"
  ]
  name   = "any"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "ip_range_subnet"
}
resource "mist_org_wxtag" "test_19" {
  op = "in"
  values = [
    "manage.mist.com",
    "api.mist.com"
  ]
  name   = "manage.mist.com"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "hostname"
}
resource "mist_org_wxtag" "test_19_1" {
  op = "in"
  values = [
    "demopsk"
  ]
  name   = "demopsk"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "radius_group"
}
resource "mist_org_wxtag" "test_20" {
  op = "in"
  values = [
    "manage.mist.com:443"
  ]
  name   = "resource_mist:443"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "hostname"
}
resource "mist_org_wxtag" "test_21" {
  specs = [
    {
      subnets = [
        "0.0.0.0/0"
      ]
      protocol   = "tcp"
      port_range = "443"
    }
  ]
  name   = "internet:443"
  org_id = mist_org.terraform_test.id
  type   = "spec"
}
resource "mist_org_wxtag" "test_22" {
  values = [
    "53"
  ]
  op     = "in"
  name   = "DNS_all"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "port"
}
resource "mist_org_wxtag" "test_23" {
  specs = [
    {
      subnets = [
        "10.3.20.201",
        "10.3.51.222"
      ]
      protocol   = "any"
      port_range = "53"
    }
  ]
  name   = "DNS_internal"
  org_id = mist_org.terraform_test.id
  type   = "spec"
}
resource "mist_org_wxtag" "test_24" {
  op = "in"
  values = [
    "netflix.com"
  ]
  name   = "netflix"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "hostname"
}
resource "mist_org_wxtag" "test_25" {
  op = "in"
  values = [
    "10.0.0.0/8",
    "172.16.0.0/22",
    "192.168.0.0/16"
  ]
  name   = "RFC1918"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "ip_range_subnet"
}
resource "mist_org_wxtag" "test_26" {
  op = "in"
  values = [
    "10.3.8.0/24"
  ]
  name   = "iot_subnet"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "ip_range_subnet"
}
resource "mist_org_wxtag" "test_27" {
  op = "in"
  values = [
    ".teamviewer.com:5938",
    ".teamviewer.com:443"
  ]
  name   = "teamviewer"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "hostname"
}
resource "mist_org_wxtag" "test_29" {
  op     = "in"
  name   = "my ssid"
  org_id = mist_org.terraform_test.id
  type   = "match"
  values = ["test"]
  match  = "wlan_id"
}
resource "mist_org_wxtag" "test_30" {
  op = "in"
  values = [
    "cdn.segment.com"
  ]
  name   = "cdn.segment.com"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "hostname"
}
resource "mist_org_wxtag" "test_32" {
  op     = "in"
  name   = "MlN.1X"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "wlan_id"
  values = ["test"]
}
resource "mist_org_wxrule" "test_02" {
  org_id      = mist_org.terraform_test.id
  template_id = mist_org_wlantemplate.test101.id
  order       = 2
  action      = "allow"
  enabled     = false
}

resource "mist_org_wxrule" "test_01" {
  org_id      = mist_org.terraform_test.id
  template_id = mist_org_wlantemplate.test101.id
  action      = "allow"
  dst_deny_wxtags = [
    mist_org_wxtag.test_30.id
  ]
  enabled = true
  src_wxtags = [
    mist_org_wxtag.test_32.id
  ]
  order = 1
}
resource "mist_org_wxrule" "test_00" {
  org_id      = mist_org.terraform_test.id
  template_id = mist_org_wlantemplate.test101.id
  action      = "block"
  order       = -1
}
# resource "mist_org_vpn" "vpn_one" {
#   org_id = mist_org.terraform_test.id
#   name   = "VpnOrgOverlay"
#   paths = {
#     "AWS_Hub_Profile1-WAN1" : {
#       bfd_profile = "broadband"
#     },
#     "AWS_Hub_Profile1-WAN2" : {},
#   }
# }

resource "mist_device_gateway_cluster" "cluster_one" {
  site_id = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, local.node0).site_id
  nodes = [
    { mac = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, local.node0).mac },
    { mac = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, local.node1).mac },
  ]
}
resource "mist_device_gateway" "cluster_one" {
  device_id = resource.mist_device_gateway_cluster.cluster_one.id
  site_id   = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, local.node0).site_id
  name      = "cluster_one"
  managed   = true
  bgp_config = {
    "test" = {
      auth_key             = "secret"
      bfd_minimum_interval = 4
      bfd_multiplier       = 3
      communities = [
        {
          id               = "1234"
          local_preference = 23
        }
      ]
    }
  }
  dhcpd_config = {
    config = {
      "test" = {}
      "CORE_VLAN410" = {
        dns_servers = ["8.8.8.8"]
        dns_suffix  = ["mycorp.local"]
        fixed_bindings = {
          "deadbeefdead" = {
            ip   = "10.4.171.11"
            name = "c0ffee"
          }
        }
        gateway    = "10.4.171.1"
        ip_end     = "10.4.171.100"
        ip_start   = "10.4.171.10"
        lease_time = "5463"
        options = {
          "42" = {
            type  = "string"
            value = "2.2.2.2"
          }
        }
        type = "local"
      }
      "RESTRICTED_VLAN420" = {
        type    = "relay"
        servers = ["1.2.3.4"]
      }
    }
  }
  vrf_instances = {
    VRF_ONE = {
      networks = [
        "CORE_VLAN410",
        "RESTRICTED_VLAN420"
      ]
    }
  }
  vrf_config = {
    enabled = true
  }
  port_config = {
    "ge-0/0/7,ge-5/0/7" = {
      usage     = "lan"
      redundant = true
      networks = [
        "CORE_VLAN410",
        "RESTRICTED_VLAN420",
        "AUTOMATION_VLAN460",
        "UNTRUSTED_VLAN490",
        "MANAGEMENT_VLAN430",
      ]
      reth_idx  = "3"
      reth_node = "node0"
    }
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
      }
      disable_autoneg = false
      speed           = "auto"
      duplex          = "auto"
      wan_source_nat = {
        disabled = false
      }
    }
  }
  ip_configs = {
    "CORE_VLAN410" = {
      type    = "static"
      ip      = "10.4.171.1"
      netmask = "/24"
    }
    "RESTRICTED_VLAN420" = {
      type    = "static"
      ip      = "10.5.171.1"
      netmask = "/24"
    }
    "AUTOMATION_VLAN460" = {
      type    = "static"
      ip      = "10.6.171.1"
      netmask = "/24"
    }
    "UNTRUSTED_VLAN490" = {
      type    = "static"
      ip      = "10.7.171.1"
      netmask = "/24"
    }
    "MANAGEMENT_VLAN430" = {
      type    = "static"
      ip      = "10.8.171.1"
      netmask = "/24"
    }
  }
  path_preferences = {
    HUB = {
      paths = [
        {
          type = "wan"
          name = "FTTH"
        }
      ]
    }
  }
  service_policies = [
    {
      name            = "Policy-14"
      tenants         = ["MANAGEMENT_VLAN430"]
      services        = ["any"]
      action          = "allow"
      path_preference = "HUB"
      idp = {
        enabled    = true
        profile    = "critical"
        alert_only = false
      }
    },
    {
      name            = "Policy-15"
      tenants         = ["MANAGEMENT_VLAN430"]
      services        = ["any"]
      action          = "allow"
      path_preference = "HUB"
      idp = {
        enabled       = true
        idpprofile_id = mist_org_idpprofile.test1.id
        profile       = "Custom"
      }
    },
    {
      servicepolicy_id = mist_org_servicepolicy.test1.id
      path_preference  = "HUB"
    }
  ]
  tunnel_configs = {
    "default" = {
      ike_lifetime = 28800
      ike_mode     = "main"
      ike_proposals = [{
        auth_algo = "sha1"
        dh_group  = "14"
        enc_algo  = "3des"
      }]
      ipsec_lifetime = 28800
      ipsec_proposals = [{
        auth_algo = "sha1"
        dh_group  = "14"
        enc_algo  = "3des"
      }]
      mode     = "active-active"
      provider = "custom-ipsec"
    }
  }
}




resource "mist_device_ap" "test_ap" {
  device_id = provider::mist::search_inventory_by_claimcode(resource.mist_org_inventory.inventory, "CPKL2EXN8JY98AC").id
  site_id   = provider::mist::search_inventory_by_claimcode(resource.mist_org_inventory.inventory, "CPKL2EXN8JY98AC").site_id
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

# # # ### ORG LEVEL
# resource "mist_org_service" "lab" {
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
# resource "mist_org_network" "networks14" {
#   org_id    = mist_org.terraform_test.id
#   isolation = true
#   subnet    = "0.0.0.0/0"
#   internet_access = {
#   }
#   vpn_access = {
#     OrgOverlay = {
#       routed = true
#     }
#   }
#   disallow_mist_services = true
#   name                   = "any"
# }
# resource "mist_org_network" "networks17" {
#   org_id    = mist_org.terraform_test.id
#   isolation = true
#   subnet    = "{{VLAN494}}.0/24"
#   internet_access = {
#   }
#   vpn_access = {
#     OrgOverlay = {
#       routed = true
#     }
#   }
#   vlan_id                = "{{VLAN494}}"
#   disallow_mist_services = true
#   name                   = "VLAN494_Network"
# }
# resource "mist_org_network" "networks1" {
#   org_id    = mist_org.terraform_test.id
#   isolation = true
#   subnet    = "{{cnf_peer_subnet}}/{{cnf_peer_prefix}}"
#   tenants = {
#     ANY = {
#       addresses = [
#         "0.0.0.0/0"
#       ]
#     }
#   }
#   internet_access = {
#   }
#   name = "CNF-Peer"
# }
# resource "mist_org_network" "corp" {
#   org_id                 = mist_org.terraform_test.id
#   disallow_mist_services = false
#   name                   = "prd_corp"
#   subnet                 = "10.4.30.0/{{test}}"
# }

# resource "mist_org_network" "mgmt" {
#   org_id  = mist_org.terraform_test.id
#   vlan_id = "{{test}}"
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
# resource "mist_org_network" "reg" {
#   org_id                 = mist_org.terraform_test.id
#   isolation              = true
#   vlan_id                = "12"
#   subnet                 = "10.3.12.0/24"
#   disallow_mist_services = false
#   name                   = "SRX-REG"
# }
# resource "mist_org_network" "ssr" {
#   org_id                 = mist_org.terraform_test.id
#   isolation              = true
#   vlan_id                = 128
#   subnet                 = "10.128.0.0/16"
#   disallow_mist_services = false
#   name                   = "SRX-Core-128T"
# }
# resource "mist_org_network" "mxe" {
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
# resource "mist_org_network" "iot" {
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
# resource "mist_org_network" "dmz" {
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


# resource "mist_org_gatewaytemplate" "test-api" {

#   type   = "spoke"
#   name   = "test-api"
#   org_id = mist_org.terraform_test.id
#   tunnel_configs = {
#     Prisma-Tunnel = {
#       provider = "custom-ipsec"
#       protocol = "ipsec"
#       local_id = "{{sc_local_id}}"
#       psk      = "psktest"
#       primary = {
#         hosts = [
#           "{{service_ip_1}}"
#         ]
#         remote_ids = [
#           "{{sc_remote_id}}"
#         ]
#         wan_names = [
#           "WAN_0"
#         ]
#       }
#       ike_proposals = [
#         {
#           auth_algo = "sha2"
#           enc_algo  = "aes256"
#           dh_group  = "19"
#         }
#       ]
#       ike_lifetime = 28800
#       ipsec_proposals = [
#         {
#           auth_algo = "sha2"
#           enc_algo  = "aes256"
#           dh_group  = "19"
#         }
#       ]
#       ipsec_lifetime = 3600
#       version        = "2"
#     }
#   }
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
#       servicepolicy_id = mist_org_servicepolicy.test1.id
#       path_preference  = "TO-INTERNET"
#     },
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
# }

# resource "mist_org_gatewaytemplate" "sdwan-westford" {
#   ntp_override = true
#   service_policies = [
#     {
#       name = "inbound-hub-to-corp-net"
#       idp = {
#         enabled = false
#       },
#       path_preference = "to-corp-lan"
#       services = [
#         "spoke-corp-network"
#       ],
#       action = "allow"
#       tenants = [
#         "corp_datacenter1",
#         "corp_datacenter2",
#         "core_dc2_network",
#         "mgmt_datacenter1",
#         "mgmt_datacenter2"
#       ]
#     },
#     {
#       name = "inbound-remote-spoke-to-corp-net"
#       tenants = [
#         "corp_network",
#         "mgmt_network"
#       ],
#       services = [
#         "spoke-corp-network"
#       ],
#       action          = "allow"
#       path_preference = "to-corp-lan"
#       idp = {
#         enabled = false
#       }
#     },
#     {
#       name = "outbound-to-hub-workloads"
#       tenants = [
#         "corp_network"
#       ],
#       services = [
#         "datacenter1-mgmt-network",
#         "datacenter1-workloads",
#         "datacenter2-lan",
#         "datacenter2-mgmt-network",
#         "datacenter2-workloads"
#       ],
#       action          = "allow"
#       path_preference = "broadband-overlay"
#       idp = {
#         enabled = false
#       }
#     },
#     {
#       name = "outbound-corp-internet"
#       tenants = [
#         "corp_network"
#       ],
#       services = [
#         "corp-internet"
#       ],
#       action          = "allow",
#       path_preference = "internet-breakout"
#       idp = {
#         enabled = false
#       }
#     },
#     {
#       name = "application-test"
#       tenants = [
#         "corp_network"
#       ],
#       services = [
#         "youtube"
#       ],
#       action = "allow"
#       idp = {
#         enabled = false
#       },
#       path_preference = "internet-breakout"
#     },
#     {
#       name = "url-filtering"
#       tenants = [
#         "corp_network"
#       ],
#       services = [
#         "blocked-categories"
#       ],
#       action        = "deny"
#       local_routing = true
#       idp = {
#         enabled = false
#       }
#     }
#   ]
#   dns_servers = [
#     "8.8.8.8"
#   ]
#   port_config = {
#     "ge-0/0/1" : {
#       name       = "wan1-broadband"
#       usage      = "wan"
#       aggregated = false
#       redundant  = false
#       critical   = false
#       disabled   = false
#       wan_type   = "broadband"
#       ip_config = {
#         type    = "static"
#         ip      = "{{broadband_ip_var}}"
#         netmask = "/{{broadband_prefix_var}}"
#         gateway = "{{broadband_gw_var}}"
#       },
#       disable_autoneg = false
#       wan_source_nat = {
#         disabled = false
#       },
#       vpn_paths = {
#         "sdwan_newyork_hub-external1.OrgOverlay" : {
#           key         = 0
#           role        = "spoke"
#           bfd_profile = "broadband"
#         }
#       }
#     },
#     "ge-0/0/3" : {
#       networks = [
#         "corp_network"
#       ],
#       usage      = "lan"
#       aggregated = false
#       redundant  = false
#       critical   = false
#       disabled   = false
#     }
#   }
#   path_preferences = {
#     "broadband-overlay" : {
#       strategy = "ordered"
#       paths = [
#         {
#           name = "sdwan_newyork_hub-external1.OrgOverlay"
#           type = "vpn"
#         }
#       ]
#     },
#     "internet-breakout" : {
#       strategy = "ordered"
#       paths = [
#         {
#           name = "wan1-broadband"
#           type = "wan"
#         }
#       ]
#     },
#     "to-corp-lan" : {
#       strategy = "ordered"
#       paths = [
#         {
#           type = "local"
#           networks = [
#             "corp_network"
#           ]
#         }
#       ]
#     }
#   }
#   ntp_servers = [
#     "pool.ntp.org"
#   ]
#   routing_policies = {
#     "ibgp-hub-preference" : {
#       terms = [
#         {
#           matching = {
#             prefix = [
#               "172.76.128.0/24",
#               "172.36.128.0/24",
#               "192.168.94.0/24",
#               "192.168.93.0/24",
#               "192.168.92.0/24",
#               "192.168.95.0/24",
#               "192.168.91.0/24",
#               "172.46.128.0/24",
#               "172.66.128.0/24",
#               "172.56.128.0/24",
#               "24.0.0.0/29",
#               "192.168.90.0/24",
#               "192.168.97.0/24",
#               "192.168.98.0/24",
#               "192.168.94.0/24"
#             ],
#             protocol = [
#               "bgp"
#             ],
#             vpn_path = [
#               "NewYork-newyork-broadband1.OrgOverlay",
#               "NewYork-newyork-broadband1-lte.OrgOverlay",
#               "SanFrancisco-sanfran-broadband1.OrgOverlay"
#             ]
#           },
#           actions = {
#             accept = true
#           }
#         }
#       ]
#     }
#   }
#   ip_configs = {
#     corp_network = {
#       type    = "static"
#       ip      = "{{corp_ip}}"
#       netmask = "/{{corp_CIDR}}"
#     }
#   }
#   dns_override = true
#   dhcpd_config = {
#     enabled = true
#     config = {
#       corp_network = {
#         type     = "local"
#         ip_start = "{{corp_dhcp_start}}"
#         ip_end   = "{{corp_dhcp_end}}"
#         gateway  = "{{corp_ip}}"
#         dns_servers = [
#           "8.8.8.8",
#           "1.1.1.1"
#         ]
#       }
#     }
#   }
#   bgp_config = {
#     "bgp-group1" : {
#       via           = "vpn"
#       vpn_name      = "OrgOverlay"
#       import_policy = "ibgp-hub-preference"
#     }
#   }
#   type   = "spoke"
#   name   = "sdwan-westford"
#   org_id = mist_org.terraform_test.id

# }


# resource "mist_org_rftemplate" "test_rf" {
#   band_24_usage = "auto"
#   band_5 = {
#     ant_gain = 2
#     power    = 8
#     channels = [
#       60,
#       104,
#       132
#     ]
#     bandwidth = 20
#   }
#   band_6 = {
#     ant_gain = 2
#     power    = 8
#   }
#   band_24 = {
#     ant_gain          = 1
#     allow_rrm_disable = true
#     power_min         = 18
#     power_max         = 18
#     bandwidth         = 20
#   }
#   ant_gain_5   = 2
#   ant_gain_6   = 2
#   ant_gain_24  = 1
#   country_code = "FR"
#   name         = "tf"
#   org_id       = mist_org.terraform_test.id
# }
resource "mist_site_setting" "test" {
  site_id = mist_site.terraform_site.id
  # analytic = {
  #   enabled = true
  # }
  gateway_mgmt = {
    root_password = "pwd123"
    app_usage     = true
    auto_signature_update = {
      enable      = true
      time_of_day = "02:00"
    }
    protect_re = {
      enabled          = true
      allowed_services = ["icmp", "ssh"]
      custom = [{
        protocol = "icmp"
        // port_range = "4000"
        subnets = ["1.2.3.4"]
      }]
      trusted_hosts = [
        "1.2.3.4", "10.0.0.1/24"
      ]
    }
    app_probing = {
      enabled = true
      custom_apps = [
        {
          name      = "test1"
          hostnames = ["1.2.3.5"]
          protocol  = "icmp"
        },
        {
          name      = "test2"
          hostnames = ["test.com"]
          protocol  = "icmp"
        },
        {
          name      = "test3"
          hostnames = ["http://example.com"]
          protocol  = "http"
        },
        {
          name      = "test4"
          hostnames = ["https://example.com"]
          protocol  = "http"
        }
      ]
    }
  }
  vars = {
  }
  ap_updown_threshold     = 5
  device_updown_threshold = 5
  auto_upgrade = {
    enabled     = true
    day_of_week = "tue"
    time_of_day = "02:00"
    version     = "beta"
  }
  config_auto_revert       = true
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

# resource "mist_org_wlan" "wlan_cwp" {
#   ssid    = "MlN.test"
#   bands   = ["5"]
#   vlan_id = "141"
#   portal = {
#     enabled                = true
#     bypass_when_cloud_down = true
#     auth                   = "sso"
#     privacy                = false
#     sso_issuer             = "https://sts.windows.net/f2532c2f-938c-4529-b6e4-aa26992b6b62/"
#     sso_nameid_format      = "email"
#     sso_idp_sign_algo      = "sha256"
#     sso_idp_cert           = "-----BEGIN CERTIFICATE-----\nMIIC8DCCAdigAwIBAgIQE5pOI9W1DZFHbB9m2Q7ADzANBgkqhkiG9w0BAQsFADA0MTIwMAYDVQQD\nEylNaWNyb3NvZnQgQXp1cmUgRmVkZXJhdGVkIFNTTyBDZXJ0aWZpY2F0ZTAeFw0yMjAyMDIxNDEz\nMTNaFw0yNTAyMDIxNDEzMTNaMDQxMjAwBgNVBAMTKU1pY3Jvc29mdCBBenVyZSBGZWRlcmF0ZWQg\nU1NPIENlcnRpZmljYXRlMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5gQTCccB3oE7\nelNYH2+11Q69Iq/2f3qf5KUZEQKwL++HyoBCOAM3wL3uLWwvRaih4+qpAeZvNsuShNIyB08SDcWN\nYsqVxaUsLYfzDD0c9VG9mwAx0Kh01S2JvtaLCaFveac7UXVfn/E/QbPXibS1EQvHUj0hwNXMrdS4\nh4TOk4D1Q70+OnCWyy7ykG1/RuO8UerIfqkQEy4C3QFb3Cyo4E7bEaYQo0NiCqD9IoM3B0wZib8Y\n3yRGJKdzXyDxuVJFb5rF7XMAHTWWAbxaN4KOLhZnjaJla7Pu/sFAj2Npm8Hm5pYEYBaUz4fc/8kg\nIwakFb3mnbnYw0xQwf+aJss1vQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQCF+oKuLmnooDzALwaE\nbFVI7PVGhU7/UZzAnq6HHI9ngF0Af2+uIrvAz6rdUM1bsGhRbj3SV2oaj26pe/1TDrGescXWhTPw\nKcXOwBnVmFr8FlMkozwpHRNzCQyFYGiTAztgQcmtwF7pilVndOmEc+p3LvCdI5JZB+LtMM/o9V+1\n+Yhm4MEWO6wTSY+j7goc/vi5f76TDZPN6PkRv17+EkybEudJuTOuIoNiqsAbNB52bVNHtxFHGIwb\nH9iS45QJ4/RG1WUr91xe3Vzh/fp1BkiHZVL4iOywOIF0TYcW7h958JEf+q0HD5LUMO47NPEbc/Cd\n+fVCTXWzABXXy4D+S8gA\n-----END CERTIFICATE-----"
#     sso_idp_sso_url        = "https://login.microsoftonline.com/f2532c2f-938c-4529-b6e4-aa26992b6b62/saml2"
#     email_enabled          = true
#   }
#   portal_allowed_hostnames = [
#     "login.microsoftonline.com",
#     "portal.mist.com",
#     "login.live.com",
#     "aadcdn.msauth.net",
#     "logincdn.msauth.net"
#   ]
#   auth = {
#     type = "psk"
#     psk  = "Juniper123"
#   }
#   apply_to    = "site"
#   org_id      = mist_org.terraform_test.id
#   template_id = mist_org_wlantemplate.wlantemplate_one.id
#   interface   = "all"
#   dynamic_vlan = {
#     enabled = true
#     type    = "standard"
#     vlans = {
#       460 = ""
#       492 = ""
#       494 = ""
#     }
#     default_vlan_ids = ["1-10", "{{dddd}}", "123"]
#   }
# }
resource "mist_site_wxtag" "test_2" {
  values = [
    "53"
  ]
  op      = "in"
  name    = "DNS"
  site_id = mist_site.terraform_site.id
  type    = "match"
  match   = "port"
}
resource "mist_site_wxtag" "test_6" {
  op = "in"
  values = [
    "test1"
  ]
  name    = "test1"
  site_id = mist_site.terraform_site.id
  type    = "match"
  match   = "radius_group"
}

resource "mist_site_wxrule" "test_01" {
  site_id = mist_site.terraform_site.id
  order   = 2
  action  = "allow"
  enabled = false
}
resource "mist_site_wxrule" "test_1" {
  action = "allow"
  dst_deny_wxtags = [
    mist_site_wxtag.test_2.id
  ]
  enabled = true
  src_wxtags = [
    mist_site_wxtag.test_6.id
  ]
  site_id = mist_site.terraform_site.id
  order   = 1
}

resource "mist_org_wlantemplate" "wlantemplate_one" {
  name   = "wlantemplate"
  org_id = mist_org.terraform_test.id
}
resource "mist_site_wlan" "test_open" {
  ssid        = "demo"
  site_id     = mist_site.terraform_site.id
  limit_bcast = true
  allow_ssdp  = true
  portal = {
    enabled = false
  }
  roam_mode = "NONE"
}

resource "mist_org_wlan" "test_open" {
  ssid         = "test"
  org_id       = mist_org.terraform_test.id
  template_id  = mist_org_wlantemplate.wlantemplate_one.id
  bands        = ["5", "6"]
  vlan_enabled = true
  vlan_id      = 160
  auth = {
    type = "open"
  }
  portal = {
    email_enabled = true
    enabled       = true
    expire        = 60
  }
  rateset = {
    "5" = {
      template = "custom"
      min_rssi = -70
      legacy   = ["6", "9", "12", "18", "24b", "36", "48", "54"]
    },
    "24" = {
      template = "high-density"
      min_rssi = 0
    }
  }

  apply_to  = "site"
  interface = "all"
}
resource "mist_org_wlan_portal_image" "image1" {
  org_id  = mist_org.terraform_test.id
  wlan_id = mist_org_wlan.test_open.id
  file    = "/Users/tmunzer/OneDrive/data/demo/IMG_0049.jpg"
}
# # # # ################### SITE LEVEL


resource "mist_org_networktemplate" "switch_template" {
  name   = "switch_template"
  org_id = mist_org.terraform_test.id
  acl_policies = [{
    actions = [{
      action  = "allow"
      dst_tag = "test2"
    }]
    name     = "acl test"
    src_tags = ["test"]
  }]
  acl_tags = {
    "test" = {
      type = "mac"
      macs = ["deadbeefc0ffee"]
    },
    "test2" = {
      type = "resource"
    }
  }
  //additional_config_cmds = []
  dhcp_snooping = {
    all_networks           = false
    enable_arp_spoof_check = true
    enable_ip_source_guard = true
    enabled                = true
    networks               = ["test"]
  }
  dns_servers = ["10.3.51.222"]
  dns_suffix  = ["stag.one"]
  extra_routes = {
    "1.2.0.0/24" = {
      via = "1.2.3.4"
    }
  }
  networks = {
    test = {
      subnet  = "1.2.3.0/24"
      vlan_id = 10
    }
    test2 = {
      subnet  = "1.2.3.0/24"
      vlan_id = 11
    }
  }
  ntp_servers = ["10.3.51.222"]
  port_mirroring = {
    "mcd-sw1-span" : {
      input_networks_ingress : [
        "CORE"
      ]
      output_port_id : "ge-0/0/42"
    },
    "test" = {
      output_port_id = "ge-0/0/10"
      input_port_ids_ingress = [
        "ge-0/0/3"
      ]
      input_port_ids_egress = [
        "ge-0/0/3"
      ]
    },
    "test2" = {
      output_network = "prx"
      input_networks_ingress = [
        "default"
      ]
    }
    "test3" = {
      output_port_id = "ge-0/0/10"
      input_networks_ingress = [
        "default"
      ]
    }
  }
  port_usages = {
    disabled_port = {
      mode         = "access"
      disabled     = true
      port_network = "default"
    },
    trunk = {
      mode     = "trunk"
      networks = []
    }
    test = {
      allow_dhcpd                                     = false
      allow_multiple_supplicants                      = true
      bypass_auth_when_server_down                    = true
      bypass_auth_when_server_down_for_unkonwn_client = false
      dynamic_vlan_networks                           = []
      description : "test"
      disable_autoneg   = false
      disabled          = false
      enable_mac_auth   = true
      enable_qos        = true
      guest_network     = "test"
      inter_switch_link = false
      mac_auth_only     = false
      mac_auth_protocol = "pap"
      mac_limit         = 10
      mode              = "access"
      mtu               = 1500
      persist_mac       = false
      poe_disabled      = false
      port_auth         = "dot1x"
      port_network      = "test2"
      storm_control = {
        no_broadcast            = true
        no_multicast            = true
        no_registered_multicast = true
        no_unknown_unicast      = true
        percentage              = 75
      }
      stp_edge         = true
      stp_no_root_port = false
      stp_p2p          = false
      voip_network     = "test2"
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
  switch_matching = {
    enable = true
    rules = [
      {
        port_mirroring = {
          "test" = {
            output_port_id = "ge-0/0/10"
            input_port_ids_ingress = [
              "ge-0/0/3"
            ]
            input_port_ids_egress = [
              "ge-0/0/3"
            ]
          },
          "test2" = {
            output_network = "prx"
            input_networks_ingress = [
              "default"
            ]
          }
          "test3" = {
            output_port_id = "ge-0/0/1"
            input_networks_ingress = [
              "default"
            ]
          }
        }
        ip_config = {
          network = "test"
          type    = "static"
        }
        oob_ip_config = {
          type = "dhcp"
        }
        match_name = "abc"
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
      },
      {
        match_role = "issue"
        name       = "issue"
        # additional_config_cmds = [
        #   "set system name-server 8.8.8.8"
        # ]
        port_config = {
          "ge-0/0/0-10" = {
            usage = "trunk"
          }
        }
      }
    ]
  }
  switch_mgmt = {
    ap_affinity_threshold = 12
    cli_banner            = "Weclome!"
    cli_idle_timeout      = 10
    config_revert_timer   = 5
    protect_re = {
      enabled          = true
      allowed_services = ["icmp", "ssh"]
      custom = [{
        protocol = "icmp"
        // port_range = "4000"
        subnets = ["1.2.3.4"]
      }]
      trusted_hosts = [
        "1.2.3.4", "10.0.0.1/24"
      ]
    }
    root_password = "Juniper123!"
  }
  vrf_config = {
    enabled = true
  }
  vrf_instances = {
    "test" = {
      networks = ["test", "test2"]
      extra_routes = {
        "1.2.3.0/24" : {
          via : "10.10.10.10"
        }
      }
    }
  }

}

# # # ################### SITE LEVEL


resource "mist_site_networktemplate" "site_switch_template" {
  site_id                             = mist_site.terraform_site.id
  dns_servers                         = ["10.3.51.222"]
  dns_suffix                          = ["stag.one"]
  ntp_servers                         = ["10.3.51.222"]
  additional_config_cmds              = ["set system hostnam test", "set system services ssh root-login allow"]
  disabled_system_defined_port_usages = ["ap"]
  networks = {
    test = {
      subnet  = "1.2.3.0/24"
      vlan_id = 10
    }
    test2 = {
      subnet  = "1.2.3.0/24"
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
      storm_control = {
        percentage = 15
      }
    }
    disabled_port = {
      mode         = "access"
      disabled     = true
      port_network = "default"
    }
  }
  port_mirroring = {
    "mcd-sw1-span" : {
      input_networks_ingress : [
        "CORE"
      ]
      output_port_id : "ge-0/0/42"
    },
    "test" = {
      output_port_id = "ge-0/0/10"
      input_port_ids_ingress = [
        "ge-0/0/3"
      ]
      input_port_ids_egress = [
        "ge-0/0/3"
      ]
    },
    "test2" = {
      output_network = "prx"
      input_networks_ingress = [
        "default"
      ]
    }
    "test3" = {
      output_port_id = "ge-0/0/10"
      input_networks_ingress = [
        "default"
      ]
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
        port_mirroring = {
          "test" = {
            output_port_id = "ge-0/0/11"
            input_port_ids_ingress = [
              "ge-0/0/3"
            ]
            input_port_ids_egress = [
              "ge-0/0/3"
            ]
          },
          "test2" = {
            output_network = "prx"
            input_networks_ingress = [
              "default"
            ]
          }
          "test3" = {
            output_port_id = "ge-0/0/10"
            input_networks_ingress = [
              "default"
            ]
          }
        }
        match_name = "abc"
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
  vrf_instances = {
    "fds" = {
      networks = [
        "test"
      ],
      extra_routes = {
        "1.2.0.0/24" = {
          via = "1.2.3.4"
        }
      }
    }
  }
  vrf_config = {
    enabled = true
  }
}


# resource "mist_org_wlan" "wlan_one" {
#   ssid        = "wlan_one"
#   org_id      = mist_org.terraform_test.id
#   template_id = mist_org_wlantemplate.test101.id
#   bands       = ["5", "6"]
#   vlan_id     = 143
#   auth = {
#     type = "psk"
#     psk  = "secretpsk"
#   }
#   interface = "all"
#   dynamic_vlan = {
#     enabled = true
#     type    = "standard"
#     vlans = {
#       460 = ""
#       492 = ""
#       494 = ""
#     }
#     default_vlan_ids = ["494"]
#   }
# }

# resource "mist_site_wlan" "wlan_cwp2" {
#   ssid    = "MlN.test"
#   bands   = ["5"]
#   vlan_id = 143
#   portal = {
#     enabled                = true
#     bypass_when_cloud_down = true
#     auth                   = "sso"
#     privacy                = false
#     sso_issuer             = "https://sts.windows.net/f2532c2f-938c-4529-b6e4-aa26992b6b62/"
#     sso_nameid_format      = "email"
#     sso_idp_sign_algo      = "sha256"
#     sso_idp_cert           = "-----BEGIN CERTIFICATE-----\nMIIC8DCCAdigAwIBAgIQE5pOI9W1DZFHbB9m2Q7ADzANBgkqhkiG9w0BAQsFADA0MTIwMAYDVQQD\nEylNaWNyb3NvZnQgQXp1cmUgRmVkZXJhdGVkIFNTTyBDZXJ0aWZpY2F0ZTAeFw0yMjAyMDIxNDEz\nMTNaFw0yNTAyMDIxNDEzMTNaMDQxMjAwBgNVBAMTKU1pY3Jvc29mdCBBenVyZSBGZWRlcmF0ZWQg\nU1NPIENlcnRpZmljYXRlMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5gQTCccB3oE7\nelNYH2+11Q69Iq/2f3qf5KUZEQKwL++HyoBCOAM3wL3uLWwvRaih4+qpAeZvNsuShNIyB08SDcWN\nYsqVxaUsLYfzDD0c9VG9mwAx0Kh01S2JvtaLCaFveac7UXVfn/E/QbPXibS1EQvHUj0hwNXMrdS4\nh4TOk4D1Q70+OnCWyy7ykG1/RuO8UerIfqkQEy4C3QFb3Cyo4E7bEaYQo0NiCqD9IoM3B0wZib8Y\n3yRGJKdzXyDxuVJFb5rF7XMAHTWWAbxaN4KOLhZnjaJla7Pu/sFAj2Npm8Hm5pYEYBaUz4fc/8kg\nIwakFb3mnbnYw0xQwf+aJss1vQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQCF+oKuLmnooDzALwaE\nbFVI7PVGhU7/UZzAnq6HHI9ngF0Af2+uIrvAz6rdUM1bsGhRbj3SV2oaj26pe/1TDrGescXWhTPw\nKcXOwBnVmFr8FlMkozwpHRNzCQyFYGiTAztgQcmtwF7pilVndOmEc+p3LvCdI5JZB+LtMM/o9V+1\n+Yhm4MEWO6wTSY+j7goc/vi5f76TDZPN6PkRv17+EkybEudJuTOuIoNiqsAbNB52bVNHtxFHGIwb\nH9iS45QJ4/RG1WUr91xe3Vzh/fp1BkiHZVL4iOywOIF0TYcW7h958JEf+q0HD5LUMO47NPEbc/Cd\n+fVCTXWzABXXy4D+S8gA\n-----END CERTIFICATE-----"
#     sso_idp_sso_url        = "https://login.microsoftonline.com/f2532c2f-938c-4529-b6e4-aa26992b6b62/saml2"
#     email_enabled          = true
#   }
#   portal_allowed_hostnames = [
#     "login.microsoftonline.com",
#     "portal.mist.com",
#     "login.live.com",
#     "aadcdn.msauth.net",
#     "logincdn.msauth.net"
#   ]
#   auth = {
#     type = "psk"
#     psk  = "Juniper123"
#   }
#   roam_mode = "11r"
#   apply_to  = "site"
#   interface = "all"
#   site_id   = mist_site.terraform_site.id
# }

# resource "mist_device_image" "switch_image_one" {
#   # for_each = { 
#   #   for  device in resource.mist_org_inventory.inventory.inventory : device.mac => device if device.mac == "2c21311c37b0" 
#   #   }
#   # device_id =each.value.id
#   # site_id = each.value.site_id
#   device_id = provider::mist::search_inventory_rs(resource.mist_org_inventory.inventory.inventory, "2c21311c37b0").id
#   site_id =  provider::mist::search_inventory_rs(resource.mist_org_inventory.inventory.inventory, "2c21311c37b0").site_id
#   file         = "/Users/tmunzer/OneDrive/data/demo/IMG_0049.jpg"
#   image_number = 1
# }

resource "mist_device_switch" "test_switch" {
  device_id = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, "2c21311c37b0").id
  site_id   = provider::mist::search_inventory_by_mac(resource.mist_org_inventory.inventory, "2c21311c37b0").site_id
  name      = "demo-ex"
  managed   = true
  role      = "test2"
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
    dns     = ["1.2.3.4"]
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
        "ge-0/0/3"
      ]
      input_port_ids_egress = [
        "ge-0/0/3"
      ]
    },
    "test2" = {
      output_network = "prx"
      input_networks_ingress = [
        "default"
      ]
    }
    "test3" = {
      output_port_id = "ge-0/0/10"
      input_networks_ingress = [
        "default"
      ]
    }
  }
  mist_nac = {
    enabled = true
  }
  vrf_instances = {
    "fds" = {
      networks = [
        "prx"
      ],
      extra_routes = {
        "1.2.0.0/24" = {
          via = "1.2.3.4"
        }
      }
    }
  }
  vrf_config = {
    enabled = true
  }
  dhcpd_config = {
    enabled = true
    config = {
      "prx" = {
        type        = "server"
        ip_start    = "10.3.18.10"
        ip_end      = "10.3.18.20"
        gateway     = "10.3.18.1"
        dns_servers = ["1.2.3.4"]

      }
    }
  }
  ospf_areas = {
    "0" = {
      type = "default"
      networks = {
        "prx" = {
          passive              = false
          interface_type       = "p2p"
          hello_interval       = 10
          dead_interval        = 40
          auth_type            = "password"
          auth_password        = "hpihpi"
          metric               = 1
          bfd_minimum_interval = 1
        }
      },
      include_loopback = false
    }
  }

  router_id = "1.2.3.4"
  oob_ip_config = {
    type = "dhcp"
    # ip      = "2.2.2.2"
    # netmask = "/24"
    # gateway = "2.2.2.1"
  }
}


# # resource "mist_device_gateway" "srx" {

# #   name      = "srx"
# #   device_id = mist_org_inventory.inventory.devices[2].id
# #   site_id   = mist_org_inventory.inventory.devices[2].site_id

# #   port_config = {
# #     "ge-0/0/3" = {
# #       name       = "FTTH"
# #       usage      = "wan"
# #       aggregated = false
# #       redundant  = false
# #       critical   = false
# #       wan_type   = "broadband"
# #       ip_config = {
# #         type    = "static"
# #         ip      = "192.168.1.8"
# #         netmask = "/24"
# #         gateway = "192.168.1.1"
# #       },
# #       disable_autoneg = false
# #       speed           = "auto"
# #       duplex          = "auto"
# #       wan_source_nat = {
# #         disabled = false
# #       },
# #       vpn_paths = {
# #         "SSR_HUB_DC-MPLS.OrgOverlay" = {
# #           key         = 0
# #           role        = "spoke"
# #           bfd_profile = "broadband"
# #         }
# #       }
# #     },
# #     "ge-0/0/4" = {
# #       name       = "LTE"
# #       usage      = "wan"
# #       aggregated = false
# #       redundant  = false
# #       critical   = false
# #       wan_type   = "broadband"
# #       ip_config = {
# #         type = "dhcp"
# #       },
# #       disable_autoneg = false
# #       speed           = "auto"
# #       duplex          = "auto"
# #       wan_source_nat = {
# #         disabled = false
# #       },
# #       vpn_paths = {
# #         "SSR_HUB_DC-MPLS.OrgOverlay" = {
# #           key         = 0
# #           role        = "spoke"
# #           bfd_profile = "broadband"
# #         }
# #       }
# #     },
# #     "ge-0/0/5" = {
# #       usage            = "lan"
# #       critical         = false
# #       aggregated       = true
# #       ae_disable_lacp  = false
# #       ae_lacp_force_up = true
# #       ae_idx           = 0
# #       redundant        = false
# #       networks = [
# #         "PRD-Core"
# #       ]
# #     },
# #     "ge-0/0/7" = {
# #       usage      = "lan"
# #       critical   = false
# #       aggregated = false
# #       redundant  = false
# #       networks = [
# #         "PRD-Mgmt"
# #       ]
# #     },
# #     "ge-0/0/6" = {
# #       usage      = "lan"
# #       critical   = false
# #       aggregated = false
# #       redundant  = false
# #       networks = [
# #         "PRD-Lab"
# #       ]
# #     }
# #   }
# #   ip_configs = {
# #     "PRD-Core" = {
# #       type    = "static"
# #       ip      = "10.3.100.9"
# #       netmask = "/24"
# #     },
# #     "PRD-Mgmt" = {
# #       type    = "static"
# #       ip      = "10.3.172.1"
# #       netmask = "/24"
# #     },
# #     "PRD-Lab" = {
# #       type    = "static"
# #       ip      = "10.3.171.1"
# #       netmask = "/24"
# #     }
# #   }
# #   dhcpd_config = {
# #     enable = true
# #   }
# #   path_preferences = {
# #     "HUB" = {
# #       strategy = "ordered"
# #       paths = [
# #         {
# #           name = "SSR_HUB_DC-MPLS.OrgOverlay"
# #           type = "vpn"
# #         }
# #       ]
# #     },
# #     "HUB-ORDERED" = {
# #       strategy = "ordered"
# #       paths = [
# #         {
# #           name     = "SSR_HUB_DC-MPLS.OrgOverlay"
# #           wan_name = "FTTH"
# #           type     = "vpn"
# #         },
# #         {
# #           name     = "SSR_HUB_DC-MPLS.OrgOverlay"
# #           wan_name = "LTE"
# #           type     = "vpn"
# #         }
# #       ]
# #     },
# #     "HUB-ECMP" = {
# #       strategy = "weighted"
# #       paths = [
# #         {
# #           name     = "SSR_HUB_DC-MPLS.OrgOverlay"
# #           wan_name = "LTE"
# #           cost     = 30
# #           type     = "vpn"
# #         },
# #         {
# #           name     = "SSR_HUB_DC-MPLS.OrgOverlay"
# #           wan_name = "FTTH"
# #           cost     = 30
# #           type     = "vpn"
# #         }
# #       ]
# #     }
# #   }
# #   service_policies = [
# #     {
# #       name = "Policy-14"
# #       tenants = [
# #         "PRD-Core"
# #       ],
# #       services = [
# #         "any"
# #       ],
# #       action          = "allow"
# #       path_preference = "HUB"
# #       idp = {
# #         enabled    = true
# #         profile    = "critical"
# #         alert_only = false
# #       }
# #     },
# #     {
# #       name = "Policy-2"
# #       tenants = [
# #         "PRD-Mgmt"
# #       ],
# #       services = [
# #         "any"
# #       ],
# #       action          = "allow",
# #       path_preference = "HUB-ECMP"
# #       idp = {
# #         enabled    = true,
# #         profile    = "standard"
# #         alert_only = true
# #       }
# #     },
# #     {
# #       name = "Policy-3"
# #       tenants = [
# #         "PRD-Lab"
# #       ],
# #       services = [
# #         "any"
# #       ],
# #       action          = "allow"
# #       path_preference = "HUB-ORDERED"
# #       idp = {
# #         enabled = false
# #       }
# #     }
# #   ]
# # }
# resource "mist_org_deviceprofile_assign" "hub_one" {
#   deviceprofile_id = mist_org_deviceprofile_gateway.hub_one.id
#   org_id = mist_org.terraform_test.id
#   macs = [mist_org_inventory.inventory.devices[4].mac]

# }
resource "mist_device_gateway" "hub_one" {
  device_id              = provider::mist::search_inventory_by_claimcode(resource.mist_org_inventory.inventory, "CV4YAS8DQWYLL6M").id
  site_id                = provider::mist::search_inventory_by_claimcode(resource.mist_org_inventory.inventory, "CV4YAS8DQWYLL6M").site_id
  name                   = "hub_one2"
  additional_config_cmds = ["#", "###"]
  dns_servers            = ["8.8.8.8"]
  managed                = true
  dhcpd_config = {
    enabled = true,
    config = { # This code block seems to be the cause, if I remove it, I stop seeing the constant diff
      mykey = {}
    }
  }
  oob_ip_config = {
  }
  port_config = {
    "ge-0/0/10" = {
      name  = "WAN_0"
      usage = "wan"
    }
  }
  tunnel_configs = {
    default = {
      ike_lifetime = 28811,
      ike_proposals = [
        {
          auth_algo = "sha2",
          dh_group  = "19",
          enc_algo  = "aes256",
        },
      ],
      ipsec_lifetime = 3600,
      ipsec_proposals = [
        {
          auth_algo = "sha2",
          dh_group  = "19",
          enc_algo  = "aes256",
        },
      ],
      local_id = "test-local-id",
      primary = {
        hosts      = ["192.0.2.1"],
        remote_ids = ["192.0.2.2"],
        wan_names  = ["WAN_0"],
      },
      protocol = "ipsec",
      provider = "custom-ipsec",
      psk      = "my-secret",
      version  = "2",
    },
  }
}
# resource "mist_device_switch" "switch_one" {
#   site_id = mist_org_inventory.inventory.devices[3].site_id
#   device_id = mist_org_inventory.inventory.devices[3].id
#   name = "hub_one"
# }

resource "mist_org_deviceprofile_gateway" "hub_one" {
  name   = "hub_one"
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
      usage = "lan"
      networks = [
        "STK-CORE"
      ]
    },
    "ge-0/0/6" = {
      usage      = "lan"
      critical   = false
      aggregated = false
      redundant  = false
      networks = [
        "STK-USER"
      ]
    }
  }
  ip_configs = {
    "STK-CORE" = {
      type    = "static"
      ip      = "172.16.0.1"
      netmask = "/24"
    },
    "STK-USER" = {
      type    = "static"
      ip      = "172.16.20.1"
      netmask = "/24"
    },
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
        "STK-CORE"
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
        "STK-USER"
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
    }
  ]
  routing_policies = {
    aj-dc-in = {
      terms = [
        {
          actions = {
            accept = true
            community = [
              "65534:200"
            ]
          }
        }
      ]
    }
  }
}



resource "mist_org_network" "CORE_VLAN410" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "0.0.0.0/0"
  internet_access = {
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  disallow_mist_services = true
  name                   = "CORE_VLAN410"
}
resource "mist_org_network" "RESTRICTED_VLAN420" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "0.0.0.0/0"
  internet_access = {
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  disallow_mist_services = true
  name                   = "RESTRICTED_VLAN420"
}
resource "mist_org_network" "AUTOMATION_VLAN460" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "0.0.0.0/0"
  internet_access = {
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  disallow_mist_services = true
  name                   = "AUTOMATION_VLAN460"
}
resource "mist_org_network" "UNTRUSTED_VLAN490" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "0.0.0.0/0"
  internet_access = {
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  disallow_mist_services = true
  name                   = "UNTRUSTED_VLAN490"
}
resource "mist_org_network" "MANAGEMENT_VLAN430" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "0.0.0.0/0"
  internet_access = {
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  disallow_mist_services = true
  name                   = "MANAGEMENT_VLAN430"
}

resource "mist_org_gatewaytemplate" "gatewaytemplate_one" {
  type                   = "spoke"
  name                   = "gatewaytemplate_one"
  org_id                 = mist_org.terraform_test.id
  additional_config_cmds = ["hello world"]
  bgp_config = {
    "test" = {
      auth_key             = "secret"
      bfd_minimum_interval = 4
      bfd_multiplier       = 3
      communities = [
        {
          id               = "1234"
          local_preference = 23
        }
      ]
    }
  }
  dhcpd_config = {
    config = {
      "test" = {}
      "CORE_VLAN410" = {
        dns_servers = ["8.8.8.8"]
        dns_suffix  = ["mycorp.local"]
        fixed_bindings = {
          "deadbeefdead" = {
            ip   = "10.4.171.11"
            name = "c0ffee"
          }
        }
        gateway    = "10.4.171.1"
        ip_end     = "10.4.171.100"
        ip_start   = "10.4.171.10"
        lease_time = "5463"
        options = {
          "42" = {
            type  = "string"
            value = "2.2.2.2"
          }
        }
        type = "local"
      }
      "RESTRICTED_VLAN420" = {
        type    = "relay"
        servers = ["1.2.3.4"]
      }
    }
    enabled = true
  }
  vrf_instances = {
    VRF_ONE = {
      networks = [
        "CORE_VLAN410",
        "RESTRICTED_VLAN420"
      ]
    }
  }
  vrf_config = {
    enabled = true
  }
  port_config = {
    "ge-0/0/7,ge-5/0/7" = {
      usage     = "lan"
      redundant = true
      networks = [
        "CORE_VLAN410",
        "RESTRICTED_VLAN420",
        "AUTOMATION_VLAN460",
        "UNTRUSTED_VLAN490",
        "MANAGEMENT_VLAN430",
      ]
      reth_idx  = "3"
      reth_node = "node0"
    }
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
      }
      disable_autoneg = false
      speed           = "auto"
      duplex          = "auto"
      wan_source_nat = {
        disabled = false
      }
    }
  }
  ip_configs = {
    "CORE_VLAN410" = {
      type    = "static"
      ip      = "10.4.171.1"
      netmask = "/24"
    }
    "RESTRICTED_VLAN420" = {
      type    = "static"
      ip      = "10.5.171.1"
      netmask = "/24"
    }
    "AUTOMATION_VLAN460" = {
      type    = "static"
      ip      = "10.6.171.1"
      netmask = "/24"
    }
    "UNTRUSTED_VLAN490" = {
      type    = "static"
      ip      = "10.7.171.1"
      netmask = "/24"
    }
    "MANAGEMENT_VLAN430" = {
      type    = "static"
      ip      = "10.8.171.1"
      netmask = "/24"
    }
  }
  path_preferences = {
    HUB = {
      paths = [
        {
          type = "wan"
          name = "FTTH"
        }
      ]
    }
  }
  service_policies = [
    {
      name            = "Policy-14"
      tenants         = ["MANAGEMENT_VLAN430"]
      services        = ["any"]
      action          = "allow"
      path_preference = "HUB"
      idp = {
        enabled    = true
        profile    = "critical"
        alert_only = false
      }
    },
    {
      name            = "Policy-15"
      tenants         = ["MANAGEMENT_VLAN430"]
      services        = ["any"]
      action          = "allow"
      path_preference = "HUB"
      idp = {
        enabled       = true
        idpprofile_id = mist_org_idpprofile.test1.id
        profile       = "Custom"
      }
    },
    {
      servicepolicy_id = mist_org_servicepolicy.test1.id
      path_preference  = "HUB"
    }
  ]
  tunnel_configs = {
    "default" = {
      ike_lifetime = 28800
      ike_mode     = "main"
      ike_proposals = [{
        auth_algo = "sha1"
        dh_group  = "14"
        enc_algo  = "3des"
      }]
      ipsec_lifetime = 28800
      ipsec_proposals = [{
        auth_algo = "sha1"
        dh_group  = "14"
        enc_algo  = "3des"
      }]
      mode     = "active-active"
      provider = "custom-ipsec"
    }
  }
}

# resource "mist_org_nactag" "test222" {
#   values = [
#     "MlN.1X"
#   ]
#   name   = "ssid.MlN.1X"
#   org_id = mist_org.terraform_test.id
#   type   = "match"
#   match  = "ssid"
# }

# resource "mist_org_nacrule" "nacrule_Wireless_EAP_TLS_McD_IoT" {
#   org_id  = mist_org.terraform_test.id
#   enabled = true
#   matching = {
#     auth_type = "eap-tls"
#     nactags = [
#       mist_org_nactag.test222.id
#     ]
#     port_types = [
#       "wireless"
#     ]
#   }
#   apply_tags = [
#     mist_org_nactag.nactag_VLAN_Crew451.id
#   ]
#   action = "allow"
#   name   = "Wireless-EAP-TLS-McD-IoT"
#   order  = 15
# }

# resource "mist_org_wlan_portal_template" "wlan_one" {
#   org_id  = mist_org.terraform_test.id
#   wlan_id = mist_org_wlan.wlan_cwp.id
#   portal_template = {
#     sms_message_format    = "Code {{code}} expires in {{duration}} minutes."
#     sms_validity_duration = "10"
#     page_title            = "Welcome To My Demo Portal"
#     locales = {
#       "fr-FR" = {
#         page_title = "Bienvenue sur mon portail de démo"
#       }
#     }
#   }
# }

resource "mist_org_wlan_portal_template" "test" {
  org_id  = mist_org.terraform_test.id
  wlan_id = mist_org_wlan.test_open.id
  portal_template = {
    sms_message_format = "Code {{code}} expires in {{duration}} minutes."
    // return null
    sms_validity_duration = "10"
    smsIsTwilio           = false
    hidePoweredBy         = true
    optout                = true
    optOutDefault         = false
    responsiveLayout      = true
    alignment             = "right"
    privacy               = false
    color                 = "#1074bc"
    cross_site            = false
    tos                   = true
    tosText               = "<< provide your Terms of Service here >>"
    privacyPolicyText     = "<< provide your Privacy Terms here >>"
    email                 = true
    name                  = true
    company               = false
    field1                = false
    // return null
    field1required = false
    field2         = false
    // return null
    field2required = false
    field3         = false
    // return null
    field3required = false
    field4         = false
    // return null
    field4Required           = false
    poweredBy                = false
    pageTitle                = "Welcome To My Portal"
    tosAcceptLabel           = "I accept the Terms of Service"
    tosLink                  = "Terms of Service"
    privacyPolicyAcceptLabel = "I accept the Privacy Terms"
    privacyPolicyLink        = "Privacy Terms"
    optoutLabel              = "Do Not Store My Personal Information"
    // return null
    back_link                = "Back to Sign In"
    tos_error                = "Please review and accept the Terms of Service"
    privacy_policy_error     = "Please review and accept the Privacy Terms"
    required_field_label     = "required"
    nameLabel                = "Name"
    nameError                = "Please provide your name"
    emailLabel               = "Email"
    emailError               = "Please provide valid email"
    companyLabel             = "Company"
    companyError             = "Please provide your company name"
    field1Label              = "Custom Field 1"
    field1Error              = "Please provide Custom Field 1"
    field2Label              = "Custom Field 2"
    field2Error              = "Please provide Custom Field 2"
    field3Label              = "Custom Field 3"
    field3Error              = "Please provide Custom Field 3"
    field4Label              = "Custom Field 4"
    field4Error              = "Please provide Custom Field 4"
    signInLabel              = "Sign In"
    message                  = "Sign in to get online"
    authLabel                = "Connect to WiFi with"
    authButtonPassphrase     = "Sign in with Passphrase"
    passphraseTitle          = "Sign in with Passphrase"
    passphraseMessage        = "Enter the secret passphrase to access the WiFi network."
    passphraseLabel          = "Passphrase"
    passphraseError          = "Invalid Passphrase"
    passphraseSubmit         = "Sign In"
    passphraseCancel         = "Cancel"
    authButtonGoogle         = "Sign in with Google"
    authButtonFacebook       = "Sign in with Facebook"
    authButtonAmazon         = "Sign in with Amazon"
    authButtonMicrosoft      = "Sign in with Microsoft"
    authButtonAzure          = "Sign in with Azure"
    authButtonSms            = "Sign in with Text Message"
    authButtonEmail          = "Sign in with Email"
    authButtonSponsor        = "Sign in as Guest"
    emailTitle               = "Sign in with Email"
    emailMessage             = "We will email you an authentication code which you can use to connect to the WiFi network."
    accessCodeAlternateEmail = "Use alternate email address"
    emailFieldLabel          = "Enter your email address"
    emailSubmit              = "Send Access Code"
    emailCancel              = "Cancel"
    emailCodeCancel          = "I did not receive the code"
    emailCodeError           = "Please provide valid alternate email"
    emailCodeFieldLabel      = "AccessCode"
    emailCodeMessage         = "Enter the access number that was sent to your Email address."
    emailCodeSubmit          = "Sign In"
    emailCodeTitle           = "Access Code"
    smsNumberTitle           = "Sign in with Text Message"
    smsNumberMessage         = "We will send an access code to your mobile number which you can use to connect to the WiFi network. Message and data rates may apply."
    smsNumberFieldLabel      = "Mobile Number"
    smsCarrierFieldLabel     = "Mobile Carrier"
    smsNumberSubmit          = "Send Access Code"
    smsNumberCancel          = "Cancel"
    smsNumberFormat          = "2125551212 (digits only)"
    smsUsernameFormat        = "username"
    smsNumberError           = "Invalid Mobile Number"
    smsCountryFormat         = "+1"
    smsCountryFieldLabel     = "Country Code"
    smsCodeTitle             = "Access Code"
    smsCodeMessage           = "Enter the access number that was sent to your mobile number."
    smsCodeFieldLabel        = "Access Code"
    smsCodeSubmit            = "Sign In"
    smsCodeCancel            = "I did not receive the code"
    smsCodeError             = "Invalid Access Code"
    smsCarrierDefault        = "Please Select"
    sms_have_access_code     = "I have an access code"
    smsCarrierError          = "Please select a mobile carrier"
    emailAccessDomainError   = "Email Access Domain Error"
    sponsorName              = "Sponsor Name"
    sponsorEmail             = "Sponsor Email TEST"
    sponsorSubmit            = "Request WiFi Access"
    sponsorRequestAccess     = "Request WiFi Access"
    sponsorCancel            = "Cancel"
    sponsorsFieldLabel       = "Sponsors"
    sponsorNameError         = "Please provide sponsor name"
    sponsorEmailError        = "Please provide valid sponsor email TEST"
    sponsorStatusPending     = "Notification Sent"
    sponsor_status_approved  = "Your request was approved"
    sponsorStatusDenied      = "Your request was denied"
    sponsorInfoPending       = "Your notification has been sent to"
    sponsorInfoApproved      = "Your request was approved by"
    sponsorInfoDenied        = "Your request was denied by"
    sponsorNotePending       = "Please wait for them to acknowledge."
    sponsorBackLink          = "Go back and edit request form"
    sponsors_error           = "Please select a sponsor"
    sponsorEmailTemplate     = ""
    multiAuth                = false
    //logo                     = "/Users/SynologyDrive/logos/Mist/Juniper_en_Mist-00.png"
    locales = {
      "fr-FR" = {
        sponsor_status_approved = "test portal 2!"
      },
      "ar" = {
        sms_have_access_code = "a"
      },
      "ca-ES" = {
        sms_have_access_code = "a"
      },
      "cs-CZ" = {
        sms_have_access_code = "a"
      },
      "da-DK" = {
        sms_have_access_code = "a"
      },
      "de-DE" = {
        sms_have_access_code = "a"
      },
      "el-GR" = {
        sms_have_access_code = "a"
      },
      "en-GB" = {
        sms_have_access_code = "a"
      },
      "en-US" = {
        sms_have_access_code = "a"
      },
      "es-ES" = {
        sms_have_access_code = "a"
      },
      "fi-FI" = {
        sms_have_access_code = "a"
      },
      "he-IL" = {
        sms_have_access_code = "a"
      },
      "zh-Hant" = {
        sms_have_access_code = "a"
      },
      "zh-Hans" = {
        sms_have_access_code = "a"
      },
      "vi-VN" = {
        sms_have_access_code = "a"
      },
      "uk-UA" = {
        sms_have_access_code = "a"
      },
      "tr-TR" = {
        sms_have_access_code = "a"
      },
      "th-TH" = {
        sms_have_access_code = "a"
      },
      "sv-SE" = {
        sms_have_access_code = "a"
      },
      "sk-SK" = {
        sms_have_access_code = "a"
      },
      "ru-RU" = {
        sms_have_access_code = "a"
      },
      "ro-RO" = {
        sms_have_access_code = "a"
      },
      "pt-PT" = {
        sms_have_access_code = "a"
      },
      "pt-BR" = {
        sms_have_access_code = "a"
      },
      "pl-PL" = {
        sms_have_access_code = "a"
      },
      "nl-NL" = {
        sms_have_access_code = "a"
      },
      "nb-NO" = {
        sms_have_access_code = "a"
      },
      "ms-MY" = {
        sms_have_access_code = "a"
      },
      "ko-KR" = {
        sms_have_access_code = "a"
      },
      "ja-JP" = {
        sms_have_access_code = "a"
      },
      "it-IT" = {
        sms_have_access_code = "a"
      },
      "hi-IN" = {
        sms_have_access_code = "a"
      },
      "hr-HR" = {
        sms_have_access_code = "a"
      },
      "hu-HU" = {
        sms_have_access_code = "a"
      },
      "id-ID" = {
        sms_have_access_code = "a"
      }
    }
  }
}

resource "mist_org_apitoken" "test_one" {
  org_id = mist_org.terraform_test.id
  name   = "test_one_reest_and_updated"
  privileges = [
    {
      scope   = "site"
      role    = "admin"
      site_id = mist_site.terraform_site.id
    },
    {
      scope   = "site"
      role    = "read"
      site_id = mist_site.site_two.id
    },
  ]
}
resource "mist_site" "site_two" {
  org_id       = mist_org.terraform_test.id
  name         = "terraform_site_two"
  country_code = "FR"
  timezone     = "Europe/Paris"
  address      = "77 Terrasse de l'Universit\u00e9, 92000 Nanterre, France"
  notes        = "Created with Terraform, Updated with Terraform"
  latlng = {
    lat = 48.899268
    lng = 2.214447
  }
  # sitegroup_ids = [mist_org_sitegroup.test_group.id, mist_org_sitegroup.test_group2.id]
  # networktemplate_id = mist_org_networktemplate.switch_template.id
  # rftemplate_id      = mist_org_rftemplate.test_rf.id
  # gatewaytemplate_id = mist_org_gatewaytemplate.test-api.id
}


resource "mist_site_setting" "site_two" {
  site_id                 = mist_site.site_two.id
  ap_updown_threshold     = 5
  device_updown_threshold = 5
  gateway_mgmt = {
    app_probing = {
      custom_apps = [{
        hostnames = [
          "example.com"
        ]
        protocol = "http"
        name     = "value"
      }]
    }

    protect_re = {
      enabled          = true
      allowed_services = ["icmp", "ssh"]
      custom = [{
        protocol = "icmp"
        // port_range = "4000"
        subnets = ["1.2.3.4"]
      }]
      trusted_hosts = [
        "1.2.3.4", "10.0.10.1/24"
      ]
    }
  }
}


resource "mist_org_sso" "admin_sso" {
  org_id            = mist_org.terraform_test.id
  name              = "admin_sso"
  custom_logout_url = "https://idp.com/logout"
  idp_cert          = "-----BEGIN CERTIFICATE-----MIIF0jCCA7qgAwIBAgIBATANBgkqhkiG9w0BAQsFADBbMQswCQYDVQQGEwJVUzENMAsGA1UECgwETWlzdDEOMAwGA1UECwwFT3JnQ0ExLTArBgNVBAMMJDk5MmJmNGI5LWM5MDAtNDg1MC05OTkyLTEwN2IyZjlkZjkyODAeFw0yNDA3MTExMjIxNTBaFw0zNDA3MDkxMjIxNTBaMFsxCzAJBgNVBAYTAlVTMQ0wCwYDVQQKDARNaXN0MQ4wDAYDVQQLDAVPcmdDQTEtMCsGA1UEAwwkOTkyYmY0YjktYzkwMC00ODUwLTk5OTItMTA3YjJmOWRmOTI4MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEApfdNlmu78olZo5BqJw7KSXhhTDnN0XvYrUQJphRIBrC26/zNPj4/+mH9koM/nprV8KFRxbg4Elu0n32K8iCKtegeLrPMzkbkqtjdR22C8x38X+JnQgxoZfASGros+PkHhFiLumbKYaMlnu4zs5BXiXBqobpKLC2MSelsJxYu3GJwJDOnxKb757b4bdcueC+609wfz15cxRCsKvvmO/ZcPidPZsZU4so3G5PPYBDFCzvFv7b3i43g8zHREZ8knMZsgM78yNTEPDoVBe9Nvvaxqv4g4zZ7NA7fJ5mCbK7CBsuHhnAt8lkDJlRVTBjDE9Z4RW7ChLKMvq17HMaxr4AizOclD3jhlOYMlRskd4gh2VEYKZRUbZrwulC3AWrIDG+60FPY6WCPyFC7uP8k4pq2pgTsmOctWvS/fJmCz8Yi62D36HhMtzb8ddOq4wpOTnL7WuhxKUblsXNuzjEbQwB9EYvHVmyxr2s1Uu+RGOaE/YHu0PsNNESIMgeUmzkDjM17qfHVEoR8VKzYxol9nlKG6f9YjFGa637dPpd+SPxv0fH9SUuUCPJcMB8AT22ufC0n3CEAp/6QOm5NjOvrMM51FRgfWU13KINmv3EHiDo6EMJrYwRrxCDLoS+8gB/j3ErcHLm1FJ3oIkOfkeUqETA9mdvIBMBtPcgW/JFpe+MQdIsCAwEAAaOBoDCBnTAvBgNVHREEKDAmgiQ5OTJiZjRiOS1jOTAwLTQ4NTAtOTk5Mi0xMDdiMmY5ZGY5MjgwDwYDVR0TAQH/BAUwAwEB/zBZBgNVHR8EUjBQME6gTKBKhkhodHRwOi8vYXBpLm1pc3QuY29tL2FwaS92MS9vcmdzLzk5MmJmNGI5LWM5MDAtNDg1MC05OTkyLTEwN2IyZjlkZjkyOC9jcmwwDQYJKoZIhvcNAQELBQADggIBAAzhXvOWj+sOqCFddBBsezW19nY1KgEgO3MFZHVuwrNpcfHVsXaY4YMmp93OawYKJKOBYpyZXRtajb/3kuGdPHHxX5+aOG0X+HNeCPzxI9xOTG3wjZ/Zr5xDKV7VsxuMyLQf0QdciLMk5q1TbZftzxSKvd0JlRO/cQNw1zocP0bcnp9t1loNcLlw7joh3qibdTj9bpHm958DgCk5WaRdsVHABLHPaPIHx07kDFedY6pGo2QILe1/6sPidxJek1yszOdWcwXeSJ+UEl4UsWNftM5E2QBib4iVr2fhKUuiOHeVnpWPCVAKG1SObfK/2uWTYItjl9ZjgIydo2rz5V0Oa6p1ftx8h9GEkAHUAJjMcHzL7u1vv+t3J029z3W1GN3Ndhox9XzEWxoQf2WdKWmAZErgZFwfDURwW0tyP1A+jAH6UK0dQJweVAo5a0HdY21x/fQz8MoLwa2oG8Y1CcrOJywZG48ll+6As0HkIrLZ1ciwJJS5s4gBu8q4/2y/XSdxaK+LmM1WdsJgYM+7PBuW5Xkf9j+n4tIMRj851cwgMk4CU3PNQP2BqOcFHKaf+v2JystzYdM2k1H9zcQMmNIMpo2VMaq+nJJYF0f4myVCEyO5T3kQLEAQWPnVJua1mXUhOSQqRBO4A6Uk8TWKyshvo0tERubwpP6iTTJWV/fSCGx7-----END CERTIFICATE-----"
  idp_sign_algo     = "sha512"
  idp_sso_url       = "https://idp.com/login"
  issuer            = "idp_issuer"
  nameid_format     = "email"
}

resource "mist_org_sso_role" "sso_role_one" {
  org_id = mist_org.terraform_test.id
  name   = "admin_sso"
  privileges = [
    {
      scope   = "site"
      role    = "read"
      site_id = mist_site.terraform_site.id
    }
  ]
}

resource "mist_org_sso_role" "sso_role_two" {
  org_id = mist_org.terraform_test.id
  name   = "admin_sso_two"
  privileges = [
    {
      scope   = "site"
      role    = "read"
      site_id = mist_site.terraform_site.id
    },
    {
      scope   = "site"
      role    = "write"
      site_id = mist_site.site_two.id
    }
  ]
}

resource "mist_org_nacidp" "idp_azure" {
  org_id                   = mist_org.terraform_test.id
  name                     = "idp_azure"
  idp_type                 = "oauth"
  oauth_cc_client_id       = "client_id"
  oauth_cc_client_secret   = "-----BEGIN RSA PRIVATE KEY-----MIIF0jCCA7qgAwIBAgIBATANBgkqhkiG9w0BAQsFADBbMQswCQYDVQQGEwJVUzENMAsGA1UECgwETWlzdDEOMAwGA1UECwwFT3JnQ0ExLTArBgNVBAMMJDk5MmJmNGI5LWM5MDAtNDg1MC05OTkyLTEwN2IyZjlkZjkyODAeFw0yNDA3MTExMjIxNTBaFw0zNDA3MDkxMjIxNTBaMFsxCzAJBgNVBAYTAlVTMQ0wCwYDVQQKDARNaXN0MQ4wDAYDVQQLDAVPcmdDQTEtMCsGA1UEAwwkOTkyYmY0YjktYzkwMC00ODUwLTk5OTItMTA3YjJmOWRmOTI4MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEApfdNlmu78olZo5BqJw7KSXhhTDnN0XvYrUQJphRIBrC26/zNPj4/+mH9koM/nprV8KFRxbg4Elu0n32K8iCKtegeLrPMzkbkqtjdR22C8x38X+JnQgxoZfASGros+PkHhFiLumbKYaMlnu4zs5BXiXBqobpKLC2MSelsJxYu3GJwJDOnxKb757b4bdcueC+609wfz15cxRCsKvvmO/ZcPidPZsZU4so3G5PPYBDFCzvFv7b3i43g8zHREZ8knMZsgM78yNTEPDoVBe9Nvvaxqv4g4zZ7NA7fJ5mCbK7CBsuHhnAt8lkDJlRVTBjDE9Z4RW7ChLKMvq17HMaxr4AizOclD3jhlOYMlRskd4gh2VEYKZRUbZrwulC3AWrIDG+60FPY6WCPyFC7uP8k4pq2pgTsmOctWvS/fJmCz8Yi62D36HhMtzb8ddOq4wpOTnL7WuhxKUblsXNuzjEbQwB9EYvHVmyxr2s1Uu+RGOaE/YHu0PsNNESIMgeUmzkDjM17qfHVEoR8VKzYxol9nlKG6f9YjFGa637dPpd+SPxv0fH9SUuUCPJcMB8AT22ufC0n3CEAp/6QOm5NjOvrMM51FRgfWU13KINmv3EHiDo6EMJrYwRrxCDLoS+8gB/j3ErcHLm1FJ3oIkOfkeUqETA9mdvIBMBtPcgW/JFpe+MQdIsCAwEAAaOBoDCBnTAvBgNVHREEKDAmgiQ5OTJiZjRiOS1jOTAwLTQ4NTAtOTk5Mi0xMDdiMmY5ZGY5MjgwDwYDVR0TAQH/BAUwAwEB/zBZBgNVHR8EUjBQME6gTKBKhkhodHRwOi8vYXBpLm1pc3QuY29tL2FwaS92MS9vcmdzLzk5MmJmNGI5LWM5MDAtNDg1MC05OTkyLTEwN2IyZjlkZjkyOC9jcmwwDQYJKoZIhvcNAQELBQADggIBAAzhXvOWj+sOqCFddBBsezW19nY1KgEgO3MFZHVuwrNpcfHVsXaY4YMmp93OawYKJKOBYpyZXRtajb/3kuGdPHHxX5+aOG0X+HNeCPzxI9xOTG3wjZ/Zr5xDKV7VsxuMyLQf0QdciLMk5q1TbZftzxSKvd0JlRO/cQNw1zocP0bcnp9t1loNcLlw7joh3qibdTj9bpHm958DgCk5WaRdsVHABLHPaPIHx07kDFedY6pGo2QILe1/6sPidxJek1yszOdWcwXeSJ+UEl4UsWNftM5E2QBib4iVr2fhKUuiOHeVnpWPCVAKG1SObfK/2uWTYItjl9ZjgIydo2rz5V0Oa6p1ftx8h9GEkAHUAJjMcHzL7u1vv+t3J029z3W1GN3Ndhox9XzEWxoQf2WdKWmAZErgZFwfDURwW0tyP1A+jAH6UK0dQJweVAo5a0HdY21x/fQz8MoLwa2oG8Y1CcrOJywZG48ll+6As0HkIrLZ1ciwJJS5s4gBu8q4/2y/XSdxaK+LmM1WdsJgYM+7PBuW5Xkf9j+n4tIMRj851cwgMk4CU3PNQP2BqOcFHKaf+v2JystzYdM2k1H9zcQMmNIMpo2VMaq+nJJYF0f4myVCEyO5T3kQLEAQWPnVJua1mXUhOSQqRBO4A6Uk8TWKyshvo0tERubwpP6iTTJWV/fSCGx7-----END RSA PRIVATE KEY-----"
  oauth_ropc_client_id     = "ropc_client_id"
  oauth_ropc_client_secret = "ropc_client_secret"
  oauth_tenant_id          = "tenant_id"
  oauth_type               = "azure"
}

data "mist_org_nacidp_metadata" "idp_azure" {
  org_id    = mist_org.terraform_test.id
  nacidp_id = mist_org_nacidp.idp_azure.id

}

resource "mist_org_nacidp" "idp_ldap" {
  org_id             = mist_org.terraform_test.id
  name               = "idp_ldap"
  idp_type           = "ldap"
  ldap_type          = "custom"
  group_filter       = "memberOf"
  member_filter      = "memberOf"
  ldap_user_filter   = "(mail=%s)"
  ldap_server_hosts  = ["ldap.mycorp.com", "1.2.3.4"]
  ldap_base_dn       = "DC=abc,DC=com"
  ldap_bind_dn       = "CN=admin,CN=users,DC=abc,DC=com"
  ldap_bind_password = "secret!password"
  ldap_cacerts = [
    "-----BEGIN CERTIFICATE-----\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\n-----END CERTIFICATE-----",
    "-----BEGIN CERTIFICATE-----\nBhMCRVMxFDASBgNVBAoMC1N0YXJ0Q29tIENBMSwwKgYDVn-----END CERTIFICATE-----"
  ]
  ldap_client_cert = "-----BEGIN CERTIFICATE-----\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\n-----END CERTIFICATE-----"
  ldap_client_key  = "-----BEGIN PRI..."
}

resource "mist_org_networktemplate" "networktemplate_one" {
  name   = "My_Provider_Network_Template"
  org_id = mist_org.terraform_test.id
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
      mode         = "trunk"
      port_network = "guest_vlan"
    }
  }
  radius_config = {
    network = "network_one"
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
        match_name = "ex"
        port_config = {
          "ge-0/0/0-10" = {
            usage = "guest"
          }
        }
      }
    ]
  }
}
