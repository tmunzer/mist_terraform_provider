FILES = {
    "./terraform-provider-mist/internal/resource_org_gatewaytemplate/org_gatewaytemplate_resource_gen.go": {
        "rename": {
            "ipd_profile_overwrite_matching": "matching",
            "routing_policy_term_matching": "matching",
            "auto_provision_primary": "primary",
            "auto_provision_secondary": "secondary",
            "ip_end4": "ip_end",
            "ip_start4": "ip_start",
            "servers4": "servers",
            "type4": "type",
            "port_ip_config": "ip_config",
        },
        "dedup": ["DestinationNat", "TrafficShaping", "StaticNat"],
    },
    "./terraform-provider-mist/internal/resource_org_deviceprofile_gateway/org_deviceprofile_gateway_resource_gen.go": {
        "rename": {
            "ipd_profile_overwrite_matching": "matching",
            "routing_policy_term_matching": "matching",
            "auto_provision_primary": "primary",
            "auto_provision_secondary": "secondary",
            "ip_end4": "ip_end",
            "ip_start4": "ip_start",
            "servers4": "servers",
            "type4": "type",
            "port_ip_config": "ip_config",
        },
        "dedup": ["DestinationNat", "TrafficShaping", "StaticNat"],
    },
    "./terraform-provider-mist/internal/resource_device_gateway/device_gateway_resource_gen.go": {
        "rename": {
            "ipd_profile_overwrite_matching": "matching",
            "routing_policy_term_matching": "matching",
            "auto_provision_primary": "primary",
            "auto_provision_secondary": "secondary",
            "ip_end4": "ip_end",
            "ip_start4": "ip_start",
            "servers4": "servers",
            "type4": "type",
            "port_ip_config": "ip_config",
        },
        "dedup": ["DestinationNat", "TrafficShaping", "StaticNat"],
    },
    "./terraform-provider-mist/internal/resource_org_networktemplate/org_networktemplate_resource_gen.go": {
        "rename": {
            "matching_rules": "rules",
            "tacacct_servers": "acct_servers",
            "vrf_extra_routes": "extra_routes",
            "snmpv3_contents": "contents",
            "snmpv3_users": "users",
            "snmpv3_vacm_content": "content",
            "ospf_networks": "networks",
        },
        "dedup": [
            "ExtraRoutes",
            "Archive",
            "Contents",
            "Users",
            "PortMirroring",
            "NextQualified",
            "RadiusConfig",
            "AcctServers",
            "AuthServers",
        ],
    },
    "./terraform-provider-mist/internal/datasource_org_networktemplates/org_networktemplates_data_source_gen.go": {
        "rename": {
            "matching_rules": "rules",
            "tacacct_servers": "acct_servers",
            "vrf_extra_routes": "extra_routes",
            "snmpv3_contents": "contents",
            "snmpv3_users": "users",
            "snmpv3_vacm_content": "content",
            "ospf_networks": "networks",
        },
        "dedup": [
            "ExtraRoutes",
            "Archive",
            "Contents",
            "Users",
            "PortMirroring",
            "NextQualified",
            "RadiusConfig",
            "AcctServers",
            "AuthServers",
        ],
    },
    "./terraform-provider-mist/internal/resource_site_networktemplate/site_networktemplate_resource_gen.go": {
        "rename": {
            "matching_rules": "rules",
            "tacacct_servers": "acct_servers",
            "vrf_extra_routes": "extra_routes",
            "snmpv3_contents": "contents",
            "snmpv3_users": "users",
            "snmpv3_vacm_content": "content",
        },
        "dedup": [
            "ExtraRoutes",
            "Archive",
            "Contents",
            "Users",
            "PortMirroring",
            "NextQualified",
            "RadiusConfig",
            "AcctServers",
            "AuthServers",
        ],
    },
    "./terraform-provider-mist/internal/resource_org_network/org_network_resource_gen.go": {
        "dedup": ["DestinationNat", "StaticNat"]
    },
    "./terraform-provider-mist/internal/resource_org_rftemplate/org_rftemplate_resource_gen.go": {
        "dedup": ["Band24", "Band5", "Band6", "Band5On24Radio"]
    },
    "./terraform-provider-mist/internal/resource_site_setting/site_setting_resource_gen.go": {
        "dedup": [
            "Hours",
        ]
    },
    "./terraform-provider-mist/internal/resource_org_setting/org_setting_resource_gen.go": {
        "rename": {
            "synthetictest": "synthetic_test",
        }
    },
    "./terraform-provider-mist/internal/resource_device_switch/device_switch_resource_gen.go": {
        "rename": {
            "matching_rules": "rules",
            "tacacct_servers": "acct_servers",
            "vrf_extra_routes": "vrf_extra_routes",
            "snmpv3_contents": "contents",
            "snmpv3_users": "users",
            "snmpv3_vacm_content": "content",
            "ospf_networks": "networks",
        },
        "dedup": [
            "ExtraRoutes",
            "Archive",
            "Contents",
            "Users",
            "PortMirroring",
            "NextQualified",
        ],
    },
    "./terraform-provider-mist/internal/resource_org_inventory/org_inventory_resource_gen.go": {
        "rename": {
            "magic": "claim_code",
        },
    },
    "./terraform-provider-mist/internal/resource_org_alarmtemplate/org_alarmtemplate_resource_gen.go": {
        "dedup": [
            "Delivery",
        ]
    },
    "./terraform-provider-mist/internal/datasource_device_gateway_stats/device_gateway_stats_data_source_gen.go": {
        "rename": {
            "Cpu2StatValue": "CpuStatValue",
            "Cpu2StatType": "CpuStatType",
            "Dhcpd2StatValue": "DhcpdStatValue",
            "Dhcpd2StatType": "DhcpdStatType",
            "If2StatValue": "IfStatValue",
            "If2StatType": "IfStatType",
            "Ip2StatValue": "IpStatValue",
            "Ip2StatType": "IpStatType",
            "Memory2StatValue": "MemoryStatValue",
            "Memory2StatType": "MemoryStatType",
            "Module2StatValue": "ModuleStatValue",
            "Module2StatType": "ModuleStatType",
            "Service2StatValue": "ServiceStatValue",
            "Service2StatType": "ServiceStatType",
            "Spu2StatValue": "SpuStatValue",
            "Spu2StatType": "SpuStatType",
        },
        "dedup": [
            "ServpInfo",
            "Errors",
            "Fans",
            "Pics",
            "PortGroups",
            "Poe",
            "Psus",
            "Temperatures",
            "VcLinks",
            "CpuStat",
            "DhcpdStat",
            "IfStat",
            "IpStat",
            "MemoryStat",
            "ModuleStat",
            "ServiceStat",
            "SpuStat",
        ],
    },
    "./terraform-provider-mist/internal/datasource_device_ap_stats/device_ap_stats_data_source_gen.go": {
        "rename": {
            "Band24Value": "BandValue",
            "Band24Type": "BandType",
            "Band5Value": "BandValue",
            "Band5Type": "BandType",
            "Band6Value": "BandValue",
            "Band6Type": "BandType",
        },
        "dedup": [
            "Band",
        ],
    },
    # "./terraform-provider-mist/internal/resource_org_wlan_portal_template/org_wlan_portal_template_resource_gen.go": {
    #     "rename": {
    #         "Ar": "Locale",
    #         "CaEs": "Locale",
    #         "CsCz": "Locale",
    #         "DaDk": "Locale",
    #         "DeDe": "Locale",
    #         "ElGr": "Locale",
    #         "EnGb": "Locale",
    #         "EnUs": "Locale",
    #         "EsRs": "Locale",
    #         "FiFi": "Locale",
    #         "FrFr": "Locale",
    #         "HeIl": "Locale",
    #         "HiIn": "Locale",
    #         "HrHr": "Locale",
    #         "HuHu": "Locale",
    #         "IdId": "Locale",
    #         "ItIt": "Locale",
    #         "JaJp": "Locale",
    #         "KoKr": "Locale",
    #         "MsMy": "Locale",
    #         "NbNo": "Locale",
    #         "NlNl": "Locale",
    #         "PlPl": "Locale",
    #         "PtBr": "Locale",
    #         "PtPt": "Locale",
    #         "RoRo": "Locale",
    #         "RuRu": "Locale",
    #         "SkSk": "Locale",
    #         "SvSe": "Locale",
    #         "ThTh": "Locale",
    #         "TrTr": "Locale",
    #         "UkUa": "Locale",
    #         "ViVn": "Locale",
    #         "ZhHns": "Locale",
    #         "ZhHant": "Locale",
    #     },
    #     "dedup": [
    #         "Locale",
    #     ],
    # },
}


for file, actions in FILES.items():
    with open(file, "r") as f_in:
        with open(f"{file}.bak", "w") as f_bak:
            f_bak.writelines(f_in)
    with open(file, "w") as f_out:
        rename = actions.get("rename", {})
        dedup = actions.get("dedup", [])
        with open(f"{file}.bak", "r") as f_in:
            SKIP = False
            UNSKIP_NEXT = False
            DONE = []
            for line in f_in.readlines():
                for src_string, dst_string in rename.items():
                    line = line.replace(src_string, dst_string)
                if UNSKIP_NEXT:
                    UNSKIP_NEXT = False
                    SKIP = False
                if line.startswith("}"):
                    UNSKIP_NEXT = True
                if line.startswith("type") or line.startswith("func"):
                    for dedup_string in dedup:
                        if (
                            f"{dedup_string}Type" in line
                            or f"{dedup_string}Value" in line
                        ):
                            if not line in DONE:
                                DONE.append(line)
                            else:
                                SKIP = True
                                UNSKIP_NEXT = False
                elif line.startswith("var"):
                    if line not in DONE:
                        DONE.append(line)
                    else:
                        SKIP = True
                        UNSKIP_NEXT = True
                if not SKIP:
                    f_out.write(f"{line}")
