package resource_device_ap

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func ipConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.ApIpConfig) IpConfigValue {
	tflog.Debug(ctx, "ipConfigSdkToTerraform")

	r_attr_type := IpConfigValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"dns":        mist_transform.ListOfStringSdkToTerraform(ctx, d.GetDns()),
		"dns_suffix": mist_transform.ListOfStringSdkToTerraform(ctx, d.GetDnsSuffix()),
		"gateway":    types.StringValue(d.GetGateway()),
		"gateway6":   types.StringValue(d.GetGateway6()),
		"ip":         types.StringValue(d.GetIp()),
		"ip6":        types.StringValue(d.GetIp6()),
		"mtu":        types.Int64Value(int64(d.GetMtu())),
		"netmask":    types.StringValue(d.GetNetmask()),
		"netmask6":   types.StringValue(d.GetNetmask6()),
		"type":       types.StringValue(string(d.GetType())),
		"type6":      types.StringValue(string(d.GetType6())),
		"vlan_id":    types.Int64Value(int64(d.GetVlanId())),
	}
	r, e := NewIpConfigValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
