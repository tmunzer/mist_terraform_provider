package resource_gatewaytemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func ipConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.GatewayTemplateIpConfig) basetypes.MapValue {
	tflog.Debug(ctx, "ipConfigsSdkToTerraform")
	port_usage_type := IpConfigsValue{}.AttributeTypes(ctx)
	state_value_map := make(map[string]attr.Value)
	for k, v := range d {
		var port_usage_state = map[string]attr.Value{
			"ip":            types.StringValue(v.GetIp()),
			"netmask":       types.StringValue(v.GetNetmask()),
			"secondary_ips": mist_transform.ListOfStringSdkToTerraform(ctx, v.GetSecondaryIps()),
		}
		port_usage_object, e := NewIpConfigsValue(port_usage_type, port_usage_state)
		diags.Append(e...)
		state_value_map[k] = port_usage_object
	}
	state_type := IpConfigsValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}
