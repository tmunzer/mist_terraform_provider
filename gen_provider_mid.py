import json

SPEC_IN = "./provider-code-spec.json"


RENAME = [
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
        ]
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
        ]
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
        ]
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
        ]
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
        ]
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
        ]
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
        ]
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
        ]
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
        ]
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
        ]
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
        ]
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
        ]
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
        ]
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
                "name": "analytic",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "auto_upgrade",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "ble_config",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "config_push_policy",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "critical_url_monitoring",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "engagement",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "led",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "occupancy",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "simple_alert",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "skyatp",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "srx_app",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "ssr",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "synthetic_test",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "vna",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "wan_vna",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "wids",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "wifi",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "wired_vna",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "zone_occupancy_alert",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "vars",
                "get": ["map"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
                                }
                            ],
                            "schema_definition": "mapplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "disabled_system_defined_port_usages",
                "get": ["list"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
                                }
                            ],
                            "schema_definition": "listplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "ssh_keys",
                "get": ["list"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
                                }
                            ],
                            "schema_definition": "listplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "org_id",
                "get": ["string"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
                                }
                            ],
                            "schema_definition": "stringplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "blacklist_url",
                "get": ["string"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
                                }
                            ],
                            "schema_definition": "stringplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "watched_station_url",
                "get": ["string"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
                                }
                            ],
                            "schema_definition": "stringplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "whitelist_url",
                "get": ["string"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
                                }
                            ],
                            "schema_definition": "stringplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "gateway_updown_threshold",
                "get": ["int64"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
                                }
                            ],
                            "schema_definition": "int64planmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "switch_updown_threshold",
                "get": ["int64"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
                                }
                            ],
                            "schema_definition": "int64planmodifier.UseStateForUnknown()",
                        }
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
    {
        "name": "org_wlan",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "org_id",
                "get": ["string"],
                "computed_optional_required": "required",                
            },
            {
                "name": "site_id",
                "get": ["string"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
                                }
                            ],
                            "schema_definition": "stringplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "id",
                "get": ["string"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
                                }
                            ],
                            "schema_definition": "stringplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "portal_sso_url",
                "get": ["string"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
                                }
                            ],
                            "schema_definition": "stringplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "mxtunnel_name",
                "get": ["list"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
                                }
                            ],
                            "schema_definition": "listplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "mxtunnel_ids",
                "get": ["list"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
                                }
                            ],
                            "schema_definition": "listplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "airwatch",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "ap_ids",
                "get": ["list"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
                                }
                            ],
                            "schema_definition": "listplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "app_limit",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "app_qos",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "qos",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "apply_to",
                "get": ["string"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
                                }
                            ],
                            "schema_definition": "stringplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "auth_servers_nas_id",
                "get": ["string"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
                                }
                            ],
                            "schema_definition": "stringplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "id",
                "get": ["string"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
                                }
                            ],
                            "schema_definition": "stringplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "auth_servers_nas_ip",
                "get": ["string"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
                                }
                            ],
                            "schema_definition": "stringplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "auth",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "bonjour",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "cisco_cwa",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "dns_server_rewrite",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "dynamic_psk",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "dynamic_vlan",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "hotspot20",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "inject_dhcp_option_82",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "mist_nac",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "portal",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "radsec",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "schedule",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
        ],
    },
    {
        "name": "site_wlan",
        "get": ["schema", "attributes"],
        "next": [
            {
                "name": "org_id",
                "get": ["string"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
                                }
                            ],
                            "schema_definition": "stringplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "site_id",
                "get": ["string"],
                "computed_optional_required": "required",
            },
            {
                "name": "id",
                "get": ["string"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
                                }
                            ],
                            "schema_definition": "stringplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "portal_sso_url",
                "get": ["string"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
                                }
                            ],
                            "schema_definition": "stringplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "mxtunnel_name",
                "get": ["list"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
                                }
                            ],
                            "schema_definition": "listplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "mxtunnel_ids",
                "get": ["list"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
                                }
                            ],
                            "schema_definition": "listplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "airwatch",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "ap_ids",
                "get": ["list"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
                                }
                            ],
                            "schema_definition": "listplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "app_limit",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "app_qos",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "qos",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "apply_to",
                "get": ["string"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
                                }
                            ],
                            "schema_definition": "stringplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "auth_servers_nas_id",
                "get": ["string"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
                                }
                            ],
                            "schema_definition": "stringplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "id",
                "get": ["string"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
                                }
                            ],
                            "schema_definition": "stringplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "auth_servers_nas_ip",
                "get": ["string"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
                                }
                            ],
                            "schema_definition": "stringplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "auth",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "bonjour",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "cisco_cwa",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "dns_server_rewrite",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "dynamic_psk",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "dynamic_vlan",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "hotspot20",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "inject_dhcp_option_82",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "mist_nac",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "portal",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "radsec",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
                    }
                ],
            },
            {
                "name": "schedule",
                "get": ["single_nested"],
                "plan_modifiers": [
                    {
                        "custom": {
                            "imports": [
                                {
                                    "path": "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
                                }
                            ],
                            "schema_definition": "objectplanmodifier.UseStateForUnknown()",
                        }
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
    json.dump(DATA, f_out)
