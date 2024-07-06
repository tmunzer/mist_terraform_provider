package resource_org_gatewaytemplate

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func ipConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.GatewayTemplateIpConfig) basetypes.MapValue {
	tflog.Debug(ctx, "ipConfigsSdkToTerraform")
	port_usage_type := IpConfigsValue{}.AttributeTypes(ctx)
	state_value_map := make(map[string]attr.Value)
	for k, v := range d {
		var port_usage_state = map[string]attr.Value{
			"ip":            types.StringValue(*v.Ip),
			"netmask":       types.StringValue(*v.Netmask),
			"secondary_ips": mist_transform.ListOfStringSdkToTerraform(ctx, v.SecondaryIps),
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
