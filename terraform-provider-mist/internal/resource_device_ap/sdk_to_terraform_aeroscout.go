package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func aeroscoutSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.ApAeroscout) AeroscoutValue {
	tflog.Debug(ctx, "aeroscoutSdkToTerraform")

	r_attr_type := AeroscoutValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"enabled":          types.BoolValue(d.GetEnabled()),
		"host":             types.StringValue(d.GetHost()),
		"locate_connected": types.BoolValue(d.GetLocateConnected()),
	}
	r, e := NewAeroscoutValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
