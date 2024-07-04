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
        },
        "dedup": ["DestinationNat", "TrafficShaping", "StaticNat"],
    },
    "./terraform-provider-mist/internal/resource_org_networktemplate/org_networktemplate_resource_gen.go": {
        "rename": {
            "matching_rules": "rules",
            "tacacct_servers": "acct_servers",
            "snmpv3_contents": "contents",
            "snmpv3_users": "users",
            "snmpv3_vacm_content": "content"
        },
        "dedup": [
            "ExtraRoutes",
            "Archive",
            "Contents",
            "Users",
        ],
    },
    "./terraform-provider-mist/internal/resource_site_networktemplate/site_networktemplate_resource_gen.go": {
        "rename": {
            "matching_rules": "rules",
            "tacacct_servers": "acct_servers",
            "snmpv3_contents": "contents",
            "snmpv3_users": "users",
            "snmpv3_vacm_content": "content"
        },
        "dedup": [
            "ExtraRoutes",
            "Archive",
            "Contents",
            "Users",
        ],
    },
    "./terraform-provider-mist/internal/resource_org_network/org_network_resource_gen.go": {
        "dedup": ["DestinationNat", "StaticNat"]
    },
    "./terraform-provider-mist/internal/resource_org_rftemplate/org_rftemplate_resource_gen.go": {
        "dedup": [
            "Band24",
            "Band5",
            "Band6",
        ]
    },
    "./terraform-provider-mist/internal/resource_site_setting/site_setting_resource_gen.go": {
        "dedup": [
            "Hours",
        ]
    },
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
                if UNSKIP_NEXT:
                    UNSKIP_NEXT = False
                    SKIP = False
                if line.startswith("}"):
                    UNSKIP_NEXT = True
                if line.startswith("type") or line.startswith("func"):
                    for dedup_string in dedup:
                        if f"{dedup_string}Type" in line or f"{dedup_string}Value" in line:
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
                    for src_string, dst_string in rename.items():
                        line = line.replace(src_string, dst_string)
                    f_out.write(f"{line}")
