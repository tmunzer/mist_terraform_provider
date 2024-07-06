package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func ipConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.JunosIpConfig) IpConfigValue {

	state_value_map_attr_type := IpConfigValue{}.AttributeTypes(ctx)

	state_value_map_attr_value := map[string]attr.Value{
		"dns":        mist_transform.ListOfStringSdkToTerraform(ctx, d.GetDns()),
		"dns_suffix": mist_transform.ListOfStringSdkToTerraform(ctx, d.GetDnsSuffix()),
		"gateway":    types.StringValue(d.GetGateway()),
		"ip":         types.StringValue(d.GetIp()),
		"netmask":    types.StringValue(d.GetNetmask()),
		"network":    types.StringValue(d.GetNetwork()),
		"type":       types.StringValue(string(d.GetType())),
	}
	r, e := NewIpConfigValue(state_value_map_attr_type, state_value_map_attr_value)
	diags.Append(e...)

	return r
}
