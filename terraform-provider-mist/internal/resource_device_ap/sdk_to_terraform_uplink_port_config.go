package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func uplinkPortConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.ApUplinkPortConfig) UplinkPortConfigValue {
	tflog.Debug(ctx, "uplinkPortConfigSdkToTerraform")

	r_attr_type := UplinkPortConfigValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"dot1x":                 types.BoolValue(d.GetDot1x()),
		"keep_wlans_up_if_down": types.BoolValue(d.GetKeepWlansUpIfDown()),
	}
	r, e := NewUplinkPortConfigValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
