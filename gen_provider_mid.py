import json

SPEC_IN = "./provider-code-spec.json"


RENAME = [
    {
        "name": "networktemplate",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "switch_matching",
                "get": ["single_nested", "attributes"],
                "next": [
                    {
                        "name": "rules",
                        "rename": "matching_rules",
                    }
                ],
            },
            {
                "name": "switch_mgmt",
                "get": ["single_nested", "attributes"],
                "next": [
                    {
                        "name": "tacacs",
                        "get": ["single_nested", "attributes"],
                        "next": [
                            {
                                "name": "acct_servers",
                                "rename": "tacacct_servers",
                            }
                        ],
                    },
                ],
            },
        ],
    },
    # {
    #     "name": "sitesetting",
    #     "get": ["schema", "attributes"],
    #     "next": [
    #         {
    #             "name": "ap_matching",
    #             "get": ["single_nested", "attributes"],
    #             "next": [
    #                 {
    #                     "name": "rules",
    #                     "rename": "ap_matching_rules",
    #                 }
    #             ],
    #         },
    #         {
    #             "name": "gateway_matching",
    #             "get": ["single_nested", "attributes"],
    #             "next": [
    #                 {
    #                     "name": "rules",
    #                     "rename": "gateway_matching_rules",
    #                 }                    
    #             ],
    #         },
    #     ],
    # },
    {
        "name": "gatewaytemplate",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "idp_profiles",
                "get": ["map_nested", "nested_object", "attributes"],
                "next": [
                    {
                        "name": "overwrites",
                        "get": ["list_nested", "nested_object", "attributes"],
                        "next": [
                            {
                                "name": "matching",
                                "rename": "ipd_profile_overwrite_matching",
                            }
                        ],
                    }
                ],
            },
            {
                "name": "routing_policies",
                "get": ["map_nested", "nested_object", "attributes"],
                "next": [
                    {
                        "name": "terms",
                        "get": ["list_nested", "nested_object", "attributes"],
                        "next": [
                            {
                                "name": "matching",
                                "rename": "routing_policy_term_matching",
                            }
                        ],
                    }
                ],
            },
            {
                "name": "tunnel_configs",
                "get": ["map_nested", "nested_object", "attributes"],
                "next": [
                    {
                        "name": "auto_provision",
                        "get": ["single_nested", "attributes"],
                        "next": [
                            {
                                "name": "primary",
                                "rename": "auto_provision_primary",
                            },
                            {
                                "name": "secondary",
                                "rename": "auto_provision_secondary",
                            },
                        ],
                    }
                ],
            },
            {
                "name": "dhcpd_config",
                "get": ["single_nested", "attributes"],
                "next": [
                    {
                        "name": "config",
                        "get": ["map_nested", "nested_object", "attributes"],
                        "next": [
                            {
                                "name": "ip_end",
                                "rename": "ip_end4",
                            },
                            {
                                "name": "ip_start",
                                "rename": "ip_start4",
                            },
                            {
                                "name": "servers",
                                "rename": "servers4",
                            },
                            {
                                "name": "type",
                                "rename": "type4",
                            },
                        ],
                    }
                ],
            },
        ],
    },
]


def next_item(data: dict, entries: list, path: list):
    for entry in entries:
        name = entry.get("name")
        get = entry.get("get")
        n = entry.get("next")
        rename = entry.get("rename")
        curr_path = path.copy()
        curr_path.append(name)
        try:
            sub_data = next(i for i in data if i["name"] == name)
            if get:
                for g in get:
                    sub_data = sub_data.get(g)
                    curr_path.append(g)
                    if not sub_data:
                        print(f"not able to get {'.'.join(curr_path)}")
            if rename:
                sub_data["name"] = rename
            if n:
                next_item(sub_data, n, curr_path)
        except:
            print(f"not found {'.'.join(curr_path)}")


with open(SPEC_IN, "r") as f_in:
    DATA = json.load(f_in)

next_item(DATA["resources"], RENAME, [])


with open(SPEC_IN, "w") as f_out:
    json.dump(DATA, f_out)