package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func eslSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.ApEslConfig) EslConfigValue {
	tflog.Debug(ctx, "eslSdkToTerraform")

	r_attr_type := EslConfigValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"cacert":      types.StringValue(d.GetCacert()),
		"channel":     types.Int64Value(int64(d.GetChannel())),
		"enabled":     types.BoolValue(d.GetEnabled()),
		"host":        types.StringValue(d.GetHost()),
		"port":        types.Int64Value(int64(d.GetPort())),
		"type":        types.StringValue(string(d.GetType())),
		"verify_cert": types.BoolValue(d.GetVerifyCert()),
		"vlan_id":     types.Int64Value(int64(d.GetVlanId())),
	}
	r, e := NewEslConfigValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
