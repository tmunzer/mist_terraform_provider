import json
import sys
import re

SPEC_IN = "./provider-code-spec.json"


CUSTOM_DEFAULT_LIST_OF_STR = """
                  {"custom": {"imports": [
                    {"path": "github.com/hashicorp/terraform-plugin-framework/attr"},
                    {"path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"},
                    {"path": "github.com/hashicorp/terraform-plugin-framework/types"}
                  ],
                  "schema_definition": "listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{}))"}}
"""
CUSTOM_DEFAULT_LIST_OF_INT = """
                   {"custom": {"imports": [
                    {"path": "github.com/hashicorp/terraform-plugin-framework/attr"},
                    {"path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"},
                    {"path": "github.com/hashicorp/terraform-plugin-framework/types"}
                  ],
                  "schema_definition": "listdefault.StaticValue(types.ListValueMust(types.Int64Type, []attr.Value{}))"}}
"""

RE_LIST = r"\"list\": {"
with open(SPEC_IN, "r") as f:
    RAW = f.read()

out = re.sub(RE_LIST, '"list": { "default": ' + CUSTOM_DEFAULT_LIST_OF_STR + ",", RAW)

with open(SPEC_IN, "w") as f:
    f.write(out)

RENAME = [
    {
        "name": "device_ap",
        "get": ["schema", "attributes"],
        "next": [

            {
                "name": "ble_config",
                "get": ["single_nested", "attributes"],
                "next": [
                    {
                        "name": "beam_disabled",
                        "get": ["list"],
                        "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
                    }
                ],
            },
            {
                "name": "radio_config",
                "get": ["single_nested", "attributes"],
                "next": [
                    {
                        "name": "band_24",
                        "get": ["single_nested", "attributes"],
                        "next": [
                            {
                                "name": "channels",
                                "get": ["list"],
                                "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
                            },
                        ],
                    },                    
                    {
                        "name": "band_5",
                        "get": ["single_nested", "attributes"],
                        "next": [
                            {
                                "name": "channels",
                                "get": ["list"],
                                "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
                            },
                        ],
                    },                    
                    {
                        "name": "band_5_on_24_radio",
                        "get": ["single_nested", "attributes"],
                        "next": [
                            {
                                "name": "channels",
                                "get": ["list"],
                                "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
                            },
                        ],
                    },                    
                    {
                        "name": "band_6",
                        "get": ["single_nested", "attributes"],
                        "next": [
                            {
                                "name": "channels",
                                "get": ["list"],
                                "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
                            },
                        ],
                    },                    
                ],
            },
        ],
    },
    {
        "name": "org_gatewaytemplate",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "port_config",
                "get": ["map_nested", "nested_object", "attributes"],
                "next": [
                    {
                        "name": "traffic_shaping",
                        "get": ["single_nested", "attributes"],
                        "next": [
                            {
                                "name": "class_percentages",
                                "get": ["list"],
                                "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
                            },
                        ],
                    },
                    {
                        "name": "vpn_paths",
                        "get": ["map_nested", "nested_object", "attributes"],
                        "next": [
                            {
                                "name": "traffic_shaping",
                                "get": ["single_nested", "attributes"],
                                "next": [
                                    {
                                        "name": "class_percentages",
                                        "get": ["list"],
                                        "default": json.loads(
                                            CUSTOM_DEFAULT_LIST_OF_INT
                                        ),
                                    },
                                ],
                            },
                        ],
                    },
                ],
            },
        ],
    },
    {
        "name": "org_networktemplate",
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
            {
                "name": "snmp_config",
                "get": ["single_nested", "attributes"],
                "next": [
                    {
                        "name": "v3_config",
                        "get": ["single_nested", "attributes"],
                        "next": [
                            {
                                "name": "notify_filter",
                                "get": ["list_nested", "nested_object", "attributes"],
                                "next": [
                                    {
                                        "name": "contents",
                                        "rename": "snmpv3_contents",
                                    }
                                ],
                            },
                            {
                                "name": "usm",
                                "get": ["single_nested", "attributes"],
                                "next": [
                                    {
                                        "name": "users",
                                        "rename": "snmpv3_users",
                                    }
                                ],
                            },
                            {
                                "name": "vacm",
                                "get": ["single_nested", "attributes"],
                                "next": [
                                    {
                                        "name": "security_to_group",
                                        "get": ["single_nested", "attributes"],
                                        "next": [
                                            {
                                                "name": "content",
                                                "rename": "snmpv3_vacm_content",
                                            }
                                        ],
                                    },
                                ],
                            },
                        ],
                    },
                ],
            },
            {
                "name": "org_id",
                "get": ["string"],
                "computed_optional_required": "required",
            },
        ],
    },
    {
        "name": "org_sitegroup",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "name",
                "get": ["string"],
                "computed_optional_required": "required",
            },
            {
                "name": "org_id",
                "get": ["string"],
                "computed_optional_required": "required",
            },
        ],
    },
    {
        "name": "org_wlantemplate",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "name",
                "get": ["string"],
                "computed_optional_required": "required",
            },
            {
                "name": "org_id",
                "get": ["string"],
                "computed_optional_required": "required",
            },
            {
                "name": "template_id",
                "get": ["string"],
                "computed_optional_required": "required",
            },
        ],
    },
    {
        "name": "org_wlan",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "ssid",
                "get": ["string"],
                "computed_optional_required": "required",
            },
            {
                "name": "org_id",
                "get": ["string"],
                "computed_optional_required": "required",
            },
            {
                "name": "dynamic_psk",
                "get": ["single_nested", "attributes"],
                "next": [
                    {
                        "name": "vlan_ids",
                        "get": ["list"],
                        "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
                    }
                ],
            },
            {
                "name": "dynamic_vlan",
                "get": ["single_nested", "attributes"],
                "next": [
                    {
                        "name": "local_vlan_ids",
                        "get": ["list"],
                        "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
                    }
                ],
            },
            {
                "name": "bonjour",
                "get": ["single_nested", "attributes"],
                "next": [
                    {
                        "name": "additional_vlan_ids",
                        "get": ["list"],        
                        "no_default": True,
                    }
                ],
            },
            {
                "name": "vlan_ids",
                "get": ["list"],           
                        "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),  
            },
        ],
    },
    {
        "name": "org_service",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "name",
                "get": ["string"],
                "computed_optional_required": "required",
            },
            {
                "name": "org_id",
                "get": ["string"],
                "computed_optional_required": "required",
            },
        ],
    },
    {
        "name": "org_network",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "name",
                "get": ["string"],
                "computed_optional_required": "required",
            },
            {
                "name": "org_id",
                "get": ["string"],
                "computed_optional_required": "required",
            },
        ],
    },
    {
        "name": "org_nactag",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "name",
                "get": ["string"],
                "computed_optional_required": "required",
            },
            {
                "name": "org_id",
                "get": ["string"],
                "computed_optional_required": "required",
            },
        ],
    },
    {
        "name": "org_wxrule",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "org_id",
                "get": ["string"],
                "computed_optional_required": "required",
            },
            {
                "name": "src_wxtags",
                "get": ["list"],           
                        "no_default": True,  
            },
        ],
    },
    {
        "name": "org_wxtag",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "name",
                "get": ["string"],
                "computed_optional_required": "required",
            },
            {
                "name": "org_id",
                "get": ["string"],
                "computed_optional_required": "required",
            },
        ],
    },
    {
        "name": "site_wxrule",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "site_id",
                "get": ["string"],
                "computed_optional_required": "required",
            },
            {
                "name": "src_wxtags",
                "get": ["list"],           
                        "no_default": True,  
            },
        ],
    },
    {
        "name": "site_wxtag",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "name",
                "get": ["string"],
                "computed_optional_required": "required",
            },
            {
                "name": "site_id",
                "get": ["string"],
                "computed_optional_required": "required",
            },
        ],
    },
    {
        "name": "org_rftemplate",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "name",
                "get": ["string"],
                "computed_optional_required": "required",
            },
            {
                "name": "org_id",
                "get": ["string"],
                "computed_optional_required": "required",
            },

                    {
                        "name": "band_24",
                        "get": ["single_nested", "attributes"],
                        "next": [
                            {
                                "name": "channels",
                                "get": ["list"],
                                "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
                            },
                        ],
                    },                    
                    {
                        "name": "band_5",
                        "get": ["single_nested", "attributes"],
                        "next": [
                            {
                                "name": "channels",
                                "get": ["list"],
                                "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
                            },
                        ],
                    },                    
                    {
                        "name": "band_5_on_24_radio",
                        "get": ["single_nested", "attributes"],
                        "next": [
                            {
                                "name": "channels",
                                "get": ["list"],
                                "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
                            },
                        ],
                    },                    
                    {
                        "name": "band_6",
                        "get": ["single_nested", "attributes"],
                        "next": [
                            {
                                "name": "channels",
                                "get": ["list"],
                                "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
                            },
                        ],
                    },    
                    {"name": "model_specific",
                     "get": ["map_nested", "nested_object", "attributes"],
                     "next":[
                             {
                        "name": "band_24",
                        "get": ["single_nested", "attributes"],
                        "next": [
                            {
                                "name": "channels",
                                "get": ["list"],
                                "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
                            },
                        ],
                    },                    
                    {
                        "name": "band_5",
                        "get": ["single_nested", "attributes"],
                        "next": [
                            {
                                "name": "channels",
                                "get": ["list"],
                                "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
                            },
                        ],
                    },                    
                    {
                        "name": "band_5_on_24_radio",
                        "get": ["single_nested", "attributes"],
                        "next": [
                            {
                                "name": "channels",
                                "get": ["list"],
                                "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
                            },
                        ],
                    },                    
                    {
                        "name": "band_6",
                        "get": ["single_nested", "attributes"],
                        "next": [
                            {
                                "name": "channels",
                                "get": ["list"],
                                "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
                            },
                        ],
                    },    
                     ]
                     }
        ],
    },
    {
        "name": "org_nacrule",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "name",
                "get": ["string"],
                "computed_optional_required": "required",
            },
            {
                "name": "org_id",
                "get": ["string"],
                "computed_optional_required": "required",
            },
        ],
    },
    {
        "name": "site_wlan",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "ssid",
                "get": ["string"],
                "computed_optional_required": "required",
            },
            {
                "name": "site_id",
                "get": ["string"],
                "computed_optional_required": "required",
            },
            {
                "name": "dynamic_psk",
                "get": ["single_nested", "attributes"],
                "next": [
                    {
                        "name": "vlan_ids",
                        "get": ["list"],
                        "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
                    }
                ],
            },
            {
                "name": "dynamic_vlan",
                "get": ["single_nested", "attributes"],
                "next": [
                    {
                        "name": "local_vlan_ids",
                        "get": ["list"],
                        "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
                    }
                ],
            },
            {
                "name": "bonjour",
                "get": ["single_nested", "attributes"],
                "next": [
                    {
                        "name": "additional_vlan_ids",
                        "get": ["list"],                        
                        "no_default": True,
                    }
                ],
            },
            {
                "name": "vlan_ids",
                "get": ["list"],
                "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
            },
        ],
    },
    {
        "name": "site_networktemplate",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "id",
                "rename": "site_id",
            },
            {
                "name": "site_id",
                "get": ["string"],
                "computed_optional_required": "required",
            },
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
            {
                "name": "snmp_config",
                "get": ["single_nested", "attributes"],
                "next": [
                    {
                        "name": "v3_config",
                        "get": ["single_nested", "attributes"],
                        "next": [
                            {
                                "name": "notify_filter",
                                "get": ["list_nested", "nested_object", "attributes"],
                                "next": [
                                    {
                                        "name": "contents",
                                        "rename": "snmpv3_contents",
                                    }
                                ],
                            },
                            {
                                "name": "usm",
                                "get": ["single_nested", "attributes"],
                                "next": [
                                    {
                                        "name": "users",
                                        "rename": "snmpv3_users",
                                    }
                                ],
                            },
                            {
                                "name": "vacm",
                                "get": ["single_nested", "attributes"],
                                "next": [
                                    {
                                        "name": "security_to_group",
                                        "get": ["single_nested", "attributes"],
                                        "next": [
                                            {
                                                "name": "content",
                                                "rename": "snmpv3_vacm_content",
                                            }
                                        ],
                                    },
                                ],
                            },
                        ],
                    },
                ],
            },
        ],
    },
    {
        "name": "site_setting",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "site_id",
                "get": ["string"],
                "computed_optional_required": "required",
            },
            {
                "name": "ble_config",
                "get": ["single_nested", "attributes"],
                "next": [
                    {
                        "name": "beam_disabled",
                        "get": ["list"],
                        "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
                    }
                ],
            },
            {
                "name": "synthetic_test",
                "get": ["single_nested", "attributes"],
                "next": [
                    {
                        "name": "vlans",
                        "get": ["list_nested", "nested_object", "attributes"],
                        "next": [
                            {
                                "name": "vlan_ids",
                                "get": ["list"],
                                "default": json.loads(CUSTOM_DEFAULT_LIST_OF_INT),
                            }
                        ],
                    }
                ],
            },
        ],
    },
    {
        "name": "org_gatewaytemplate",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "org_id",
                "get": ["string"],
                "computed_optional_required": "required",
            },
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
        plan_modifiers = entry.get("plan_modifiers")
        computed_optional_required = entry.get("computed_optional_required")
        default = entry.get("default")
        no_default = entry.get("no_default")
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
            if default:
                sub_data["default"] = default
            if no_default and sub_data.get("default"):
                del sub_data["default"]
            if plan_modifiers:
                sub_data["plan_modifiers"] = plan_modifiers
            if computed_optional_required:
                sub_data["computed_optional_required"] = computed_optional_required
            if n:
                next_item(sub_data, n, curr_path)
        except:
            print(f"not found {'.'.join(curr_path)}")


with open(SPEC_IN, "r") as f_in:
    DATA = json.load(f_in)

next_item(DATA["resources"], RENAME, [])


with open(SPEC_IN, "w") as f_out:
    json.dump(DATA, f_out, indent=4)
