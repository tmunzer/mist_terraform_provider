package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func pwrConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.ApPwrConfig) PwrConfigValue {
	tflog.Debug(ctx, "pwrConfigSdkToTerraform")

	r_attr_type := PwrConfigValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"base":                 types.Int64Value(int64(d.GetBase())),
		"prefer_usb_over_wifi": types.BoolValue(d.GetPreferUsbOverWifi()),
	}
	r, e := NewPwrConfigValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
