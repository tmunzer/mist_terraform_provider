package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func HoursSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.Hours) basetypes.ObjectValue {
	tflog.Debug(ctx, "HoursSdkToTerraform")
	r_attr_type := HoursValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"mon": types.StringValue(d.GetMon()),
		"tue": types.StringValue(d.GetTue()),
		"wed": types.StringValue(d.GetWed()),
		"thu": types.StringValue(d.GetThu()),
		"fri": types.StringValue(d.GetFri()),
		"sat": types.StringValue(d.GetSat()),
		"sun": types.StringValue(d.GetSun()),
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
