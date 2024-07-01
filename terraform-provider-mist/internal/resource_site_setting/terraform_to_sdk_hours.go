package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func HoursConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.Hours {
	tflog.Debug(ctx, "HoursConfigTerraformToSdk")
	data := mistsdkgo.NewHours()
	if d.IsNull() || d.IsUnknown() {
		return *data
	} else {
		v := NewHoursValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetMon(v.Mon.ValueString())
		data.SetThu(v.Tue.ValueString())
		data.SetWed(v.Wed.ValueString())
		data.SetThu(v.Thu.ValueString())
		data.SetFri(v.Fri.ValueString())
		data.SetSat(v.Sat.ValueString())
		data.SetSun(v.Sun.ValueString())

		return *data
	}
}
