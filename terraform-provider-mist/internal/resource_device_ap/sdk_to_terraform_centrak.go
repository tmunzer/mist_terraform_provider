package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func centrakSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.ApCentrak) CentrakValue {
	tflog.Debug(ctx, "centrakSdkToTerraform")

	r_attr_type := CentrakValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"enabled": types.BoolValue(d.GetEnabled()),
	}
	r, e := NewCentrakValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
