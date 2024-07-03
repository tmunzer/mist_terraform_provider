package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func portUsageStormControlSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.SwitchPortUsageStormControl) basetypes.ObjectValue {
	storm_control_attr_type := StormControlValue{}.AttributeTypes(ctx)
	storm_control_attr_value := map[string]attr.Value{
		"no_broadcast":            types.BoolValue(d.GetNoBroadcast()),
		"no_multicast":            types.BoolValue(d.GetNoMulticast()),
		"no_registered_multicast": types.BoolValue(d.GetNoRegisteredMulticast()),
		"no_unknown_unicast":      types.BoolValue(d.GetNoUnknownUnicast()),
		"percentage":              types.Int64Value(int64(d.GetPercentage())),
	}

	r, e := basetypes.NewObjectValue(storm_control_attr_type, storm_control_attr_value)
	diags.Append(e...)

	return r
}

func portUsageRulesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistsdkgo.SwitchPortUsageDynamicRule) basetypes.ListValue {
	var value_list []attr.Value
	value_list_type := RulesValue{}.AttributeTypes(ctx)
	for _, v := range d {
		value_list_item := map[string]attr.Value{
			"equals":     types.StringValue(v.GetEquals()),
			"equals_any": mist_transform.ListOfStringSdkToTerraform(ctx, v.GetEqualsAny()),
			"expression": types.StringValue(v.GetExpression()),
			"src":        types.StringValue(string(v.GetSrc())),
			"usage":      types.StringValue(v.GetUsage()),
		}
		data, e := NewRulesValue(value_list_type, value_list_item)
		diags.Append(e...)

		value_list = append(value_list, data)
	}

	state_list_type := RulesValue{}.Type(ctx)
	state_list, e := types.ListValueFrom(ctx, state_list_type, value_list)
	diags.Append(e...)

	return state_list
}

func portUsagesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.SwitchPortUsage) basetypes.MapValue {
	port_usage_type := PortUsagesValue{}.AttributeTypes(ctx)
	state_value_map := make(map[string]attr.Value)
	for k, v := range d {
		storm_control_state := portUsageStormControlSdkToTerraform(ctx, diags, mistsdkgo.SwitchPortUsageStormControl(v.GetStormControl()))
		rules_state := portUsageRulesSdkToTerraform(ctx, diags, v.GetRules())
		var port_usage_state = map[string]attr.Value{
			"all_networks":                                    types.BoolValue(v.GetAllNetworks()),
			"allow_dhcpd":                                     types.BoolValue(v.GetAllowDhcpd()),
			"allow_multiple_supplicants":                      types.BoolValue(v.GetAllowMultipleSupplicants()),
			"bypass_auth_when_server_down":                    types.BoolValue(v.GetBypassAuthWhenServerDown()),
			"bypass_auth_when_server_down_for_unkonwn_client": types.BoolValue(v.GetBypassAuthWhenServerDownForUnkonwnClient()),
			"description":                                     types.StringValue(v.GetDescription()),
			"disable_autoneg":                                 types.BoolValue(v.GetDisableAutoneg()),
			"disabled":                                        types.BoolValue(v.GetDisabled()),
			"duplex":                                          types.StringValue(string(v.GetDuplex())),
			"dynamic_vlan_networks":                           mist_transform.ListOfStringSdkToTerraform(ctx, v.GetDynamicVlanNetworks()),
			"enable_mac_auth":                                 types.BoolValue(v.GetEnableMacAuth()),
			"enable_qos":                                      types.BoolValue(v.GetEnableQos()),
			"guest_network":                                   types.StringValue(v.GetGuestNetwork()),
			"inter_switch_link":                               types.BoolValue(v.GetInterSwitchLink()),
			"mac_auth_only":                                   types.BoolValue(v.GetMacAuthOnly()),
			"mac_auth_protocol":                               types.StringValue(string(v.GetMacAuthProtocol())),
			"mac_limit":                                       types.Int64Value(int64(v.GetMacLimit())),
			"mode":                                            types.StringValue(string(v.GetMode())),
			"mtu":                                             types.Int64Value(int64(v.GetMtu())),
			"networks":                                        mist_transform.ListOfStringSdkToTerraform(ctx, v.GetNetworks()),
			"persist_mac":                                     types.BoolValue(v.GetPersistMac()),
			"poe_disabled":                                    types.BoolValue(v.GetPoeDisabled()),
			"port_auth":                                       types.StringValue(v.GetPortAuth()),
			"port_network":                                    types.StringValue(v.GetPortNetwork()),
			"reauth_interval":                                 types.Int64Value(int64(v.GetReauthInterval())),
			"rejected_network":                                types.StringValue(v.GetRejectedNetwork()),
			"rules":                                           rules_state,
			"reset_default_when":                              types.StringValue(string(v.GetResetDefaultWhen())),
			"speed":                                           types.StringValue(v.GetSpeed()),
			"storm_control":                                   storm_control_state,
			"stp_edge":                                        types.BoolValue(v.GetStpEdge()),
			"voip_network":                                    types.StringValue(v.GetVoipNetwork()),
		}
		port_usage_object, e := NewPortUsagesValue(port_usage_type, port_usage_state)
		diags.Append(e...)
		state_value_map[k] = port_usage_object
	}
	state_type := PortUsagesValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}
