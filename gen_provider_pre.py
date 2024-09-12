import re
import yaml

SPEC_IN = "./mist_openapi/workdir/mist.openapi.yml"
SPEC_OUT = "./mist.provider.yml"
SPEC_OUT_JSON = "./mist.sdk.json"
REG = r"^( *)enum:$"

with open(SPEC_IN, "r") as f_in:
    with open(SPEC_OUT, "w") as f_out:
        f_lines = f_in.readlines()
        for line in f_lines:
            f_out.write(line)
            res = re.match(REG, line)
            if res:
                f_out.write(f'{res.groups()[0]}- ""\n')


with open(SPEC_OUT, "r") as f:
    DATA = yaml.load(f, Loader=yaml.loader.SafeLoader)

## switch_matching_rule
if (
    DATA.get("components", {})
    .get("schemas", {})
    .get("switch_matching_rule", {})
    .get("additionalProperties")
):
    del DATA["components"]["schemas"]["switch_matching_rule"]["additionalProperties"]
    DATA["components"]["schemas"]["switch_matching_rule"]["properties"][
        "match_type"
    ] = {
        "type": "string",
        "description": "'property key define the type of matching, value is the string to match. e.g: `match_name[0:3]`, `match_name[2:6]`, `match_model`,  `match_model[0-6]`",
        "example": "match_name[0:3]",
    }
    DATA["components"]["schemas"]["switch_matching_rule"]["properties"][
        "match_value"
    ] = {"type": "string"}

## dhcpd_config
if (
    DATA.get("components", {})
    .get("schemas", {})
    .get("dhcpd_config", {})
    .get("additionalProperties")
):
    del DATA["components"]["schemas"]["dhcpd_config"]["additionalProperties"]

    DATA["components"]["schemas"]["dhcpd_config"]["properties"]["config"] = {
        "$ref": "#/components/schemas/dhcpd_config_fix"
    }
    DATA["components"]["schemas"]["dhcpd_config_fix"] = {
        "type": "object",
        "description": "Property key is the network name",
        "additionalProperties": {"$ref": "#/components/schemas/dhcpd_config_property"},
    }

## wlan.portal.sponsors
## remove anyOf to remove support of the backward compatibility
## (only support object, remove array)
if (
    DATA.get("components", {})
    .get("schemas", {})
    .get("wlan_portal", {})
    .get("properties", {})
    .get("sponsors", {})
):
    del DATA["components"]["schemas"]["wlan_portal"]["properties"]["sponsors"]
    DATA["components"]["schemas"]["wlan_portal"]["properties"]["sponsors"] = {
        "additionalProperties": {"type": "string"},
        "default": {},
        "description": """object of allowed sponsors email with name. Required if `sponsor_enabled`
            is `true` and `sponsor_email_domains` is empty.

            Property key is the sponsor email, Property value is the sponsor name""",
        "example": {
            "sponsor1@company.com": "FirstName1 LastName1",
            "sponsor2@company.com": "FirstName2 LastName2",
        },
        "type": "object",
    }

## wlan_bonjour.additional_vlan_ids
## replace string of VLANS with array of VLANS
if (
    DATA.get("components", {})
    .get("schemas", {})
    .get("wlan_bonjour", {})
    .get("properties", {})
    .get("additional_vlan_ids", {})
):
    del DATA["components"]["schemas"]["wlan_bonjour"]["properties"][
        "additional_vlan_ids"
    ]
    DATA["components"]["schemas"]["wlan_bonjour"]["properties"][
        "additional_vlan_ids"
    ] = {"$ref": "#/components/schemas/wlan_bonjour_additional_vlan_ids"}
    DATA["components"]["schemas"]["wlan_bonjour_additional_vlan_ids"] = {
        "default": [],
        "description": "additional VLAN IDs (on the LAN side or from other WLANs) should we be forwarding bonjour queries/responses",
        "items": {"$ref": "#/components/schemas/vlan_id_with_variable"},
        "type": "array",
    }

## portal templates
locales = [            "ar", "ca-ES", "cs-CZ", "da-DK", "de-DE", "el-GR", "en-GB", "en-US", "es-ES", 
            "fi-FI", "fr-FR", "he-IL", "hi-IN", "hr-HR", "hu-HU", "id-ID", "it-IT", "ja-JP", 
            "ko-KR", "ms-MY", "nb-NO", "nl-NL", "pl-PL", "pt-BR", "pt-PT", "ro-RO", "ru-RU", 
            "sk-SK", "sv-SE", "th-TH", "tr-TR", "uk-UA", "vi-VN", "zh-Hans", "zh-Hant"]
if (
    DATA.get("components", {})
    .get("schemas", {})
    .get("wlan_portal_template_setting", {})
    .get("properties")
):
    for locale in locales:
        if (
            DATA.get("components", {})
            .get("schemas", {})
            .get("wlan_portal_template_setting", {})
            .get("properties", {})
            .get(locale)
        ):
            del DATA["components"]["schemas"]["wlan_portal_template_setting"]["properties"][locale]
    DATA["components"]["schemas"]["wlan_portal_template_setting"]["properties"]["locales"] = {
        "type": "object",
        "additionalProperties": {
            "$ref": '#/components/schemas/wlan_portal_template_setting_locale'
        }
    }

## remove webhooks samples
del DATA["paths"]["/webhook_example/_alarm_"]
del DATA["components"]["schemas"]["webhook_alarms"]
del DATA["components"]["schemas"]["webhook_alarms_events"]
del DATA["components"]["schemas"]["webhook_alarm_event"]

del DATA["paths"]["/webhook_example/_asset_raw_"]
del DATA["components"]["schemas"]["webhook_asset_raw"]
del DATA["components"]["schemas"]["webhook_asset_raw_events"]
del DATA["components"]["schemas"]["webhook_asset_raw_event"]

del DATA["paths"]["/webhook_example/_audit_"]
del DATA["components"]["schemas"]["webhook_audits"]
del DATA["components"]["schemas"]["webhook_audits_events"]
del DATA["components"]["schemas"]["webhook_audit_event"]

del DATA["paths"]["/webhook_example/_client_info_"]
del DATA["components"]["schemas"]["webhook_client_info"]
del DATA["components"]["schemas"]["webhook_client_info_events"]
del DATA["components"]["schemas"]["webhook_client_info_event"]

del DATA["paths"]["/webhook_example/_client_join_"]
del DATA["components"]["schemas"]["webhook_client_join"]
del DATA["components"]["schemas"]["webhook_client_join_events"]
del DATA["components"]["schemas"]["webhook_client_join_event"]

del DATA["paths"]["/webhook_example/_client_latency_"]
del DATA["components"]["schemas"]["webhook_client_latency"]
del DATA["components"]["schemas"]["webhook_client_latency_events"]
del DATA["components"]["schemas"]["webhook_client_latency_event"]

del DATA["paths"]["/webhook_example/_client_sessions_"]
del DATA["components"]["schemas"]["webhook_client_sessions"]
del DATA["components"]["schemas"]["webhook_client_sessions_events"]
del DATA["components"]["schemas"]["webhook_client_sessions_event"]

del DATA["paths"]["/webhook_example/_device_events_"]
del DATA["components"]["schemas"]["webhook_device_events"]
del DATA["components"]["schemas"]["webhook_device_events_events"]
del DATA["components"]["schemas"]["webhook_device_events_event"]

del DATA["paths"]["/webhook_example/_device_updowns_"]
del DATA["components"]["schemas"]["webhook_device_updowns"]
del DATA["components"]["schemas"]["webhook_device_updowns_events"]
del DATA["components"]["schemas"]["webhook_device_updowns_event"]

del DATA["paths"]["/webhook_example/_discovered_raw_rssi_"]
del DATA["components"]["schemas"]["webhook_discovered_raw_rssi"]
del DATA["components"]["schemas"]["webhook_discovered_raw_rssi_events"]
del DATA["components"]["schemas"]["webhook_discovered_raw_rssi_event"]
del DATA["components"]["schemas"]["webhook_discovered_raw_rssi_event_ap_loc"]

del DATA["paths"]["/webhook_example/_location_"]
del DATA["components"]["schemas"]["webhook_location"]
del DATA["components"]["schemas"]["webhook_location_events"]
del DATA["components"]["schemas"]["webhook_location_event"]
del DATA["components"]["schemas"]["wifi_beacon_extended_info"]
del DATA["components"]["schemas"]["wifi_beacon_extended_info_items"]

del DATA["paths"]["/webhook_example/_location_asset_"]
del DATA["components"]["schemas"]["webhook_location_asset"]
del DATA["components"]["schemas"]["webhook_location_asset_events"]
del DATA["components"]["schemas"]["webhook_location_asset_event"]

del DATA["paths"]["/webhook_example/_location_centrak_"]
del DATA["components"]["schemas"]["webhook_location_centrak"]
del DATA["components"]["schemas"]["webhook_location_centrak_events"]
del DATA["components"]["schemas"]["webhook_location_centrak_event"]

del DATA["paths"]["/webhook_example/_location_client_"]
del DATA["components"]["schemas"]["webhook_location_client"]
del DATA["components"]["schemas"]["webhook_location_client_events"]
del DATA["components"]["schemas"]["webhook_location_client_event"]

del DATA["paths"]["/webhook_example/_location_sdk_"]
del DATA["components"]["schemas"]["webhook_location_sdk"]
del DATA["components"]["schemas"]["webhook_location_sdk_events"]
del DATA["components"]["schemas"]["webhook_location_sdk_event"]

del DATA["paths"]["/webhook_example/_location_unclient_"]
del DATA["components"]["schemas"]["webhook_location_unclient"]
del DATA["components"]["schemas"]["webhook_location_unclient_events"]
del DATA["components"]["schemas"]["webhook_location_unclient_event"]

del DATA["paths"]["/webhook_example/_nac_accounting_"]
del DATA["components"]["schemas"]["webhook_nac_accounting"]
del DATA["components"]["schemas"]["webhook_nac_accounting_events"]
del DATA["components"]["schemas"]["webhook_nac_accounting_event"]

del DATA["paths"]["/webhook_example/_nac_events_"]
del DATA["components"]["schemas"]["webhook_nac_events"]
del DATA["components"]["schemas"]["webhook_nac_events_events"]
del DATA["components"]["schemas"]["webhook_nac_events_event"]
del DATA["components"]["schemas"]["webhook_nac_events_event_idp_role"]
del DATA["components"]["schemas"]["webhook_nac_events_event_resp_attrs"]

del DATA["paths"]["/webhook_example/_occupancy_alerts_"]
del DATA["components"]["schemas"]["webhook_occupancy_alerts"]
del DATA["components"]["schemas"]["webhook_occupancy_alerts_events"]
del DATA["components"]["schemas"]["webhook_occupancy_alerts_event"]
del DATA["components"]["schemas"]["webhook_occupancy_alerts_event_alert_events"]
del DATA["components"]["schemas"]["webhook_occupancy_alerts_event_alert_events_items"]
del DATA["components"]["schemas"]["webhook_occupancy_alert_type"]

del DATA["paths"]["/webhook_example/_ping_"]
del DATA["components"]["schemas"]["webhook_ping"]
del DATA["components"]["schemas"]["webhook_ping_events"]
del DATA["components"]["schemas"]["webhook_ping_event"]

del DATA["paths"]["/webhook_example/_sdkclient_scan_data"]
del DATA["components"]["schemas"]["webhook_sdkclient_scan_data"]
del DATA["components"]["schemas"]["webhook_sdkclient_scan_data_events"]
del DATA["components"]["schemas"]["webhook_sdkclient_scan_data_event"]
del DATA["components"]["schemas"]["webhook_sdkclient_scan_data_topic"]
del DATA["components"]["schemas"]["webhook_sdkclient_scan_data_event_scan_data"]
del DATA["components"]["schemas"]["webhook_sdkclient_scan_data_event_scan_data_item"]
del DATA["components"]["schemas"]["scan_data_item_band"]

del DATA["paths"]["/webhook_example/_site_sle_"]
del DATA["components"]["schemas"]["webhook_site_sle"]
del DATA["components"]["schemas"]["webhook_site_sle_events"]
del DATA["components"]["schemas"]["webhook_site_sle_event"]
del DATA["components"]["schemas"]["webhook_site_sle_event_sle"]

del DATA["paths"]["/webhook_example/_zone_"]
del DATA["components"]["schemas"]["webhook_zone"]
del DATA["components"]["schemas"]["webhook_zone_events"]
del DATA["components"]["schemas"]["webhook_zone_event"]
del DATA["components"]["schemas"]["webhook_zone_event_trigger"]

with open(SPEC_OUT, "w") as f:
    yaml.dump({"openapi": DATA["openapi"]}, f)
    yaml.dump({"info": DATA["info"]}, f)
    yaml.dump({"servers": DATA["servers"]}, f)
    yaml.dump({"security": DATA["security"]}, f)
    yaml.dump({"tags": DATA["tags"]}, f)
    yaml.dump({"paths": DATA["paths"]}, f)
    yaml.dump({"components": DATA["components"]}, f)
