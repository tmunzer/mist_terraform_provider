package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func oobIpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SwitchOobIpConfig) OobIpConfigValue {

	state_value_map_attr_type := OobIpConfigValue{}.AttributeTypes(ctx)

	state_value_map_attr_value := map[string]attr.Value{
		"ip":                        types.StringValue(d.GetIp()),
		"netmask":                   types.StringValue(d.GetNetmask()),
		"network":                   types.StringValue(d.GetNetwork()),
		"type":                      types.StringValue(string(d.GetType())),
		"use_mgmt_vrf":              types.BoolValue(d.GetUseMgmtVrf()),
		"use_mgmt_vrf_for_host_out": types.BoolValue(d.GetUseMgmtVrfForHostOut()),
	}
	r, e := NewOobIpConfigValue(state_value_map_attr_type, state_value_map_attr_value)
	diags.Append(e...)

	return r
}
