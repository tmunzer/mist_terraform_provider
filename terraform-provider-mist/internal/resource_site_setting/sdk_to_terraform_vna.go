package resource_site_setting

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func vnaSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.SiteSettingVna) VnaValue {
	tflog.Debug(ctx, "vnaSdkToTerraform")

	r_attr_type := VnaValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"enabled": types.BoolValue(d.GetEnabled()),
	}
	r, e := NewVnaValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
