package resource_site_setting

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func wanVnaSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.SiteSettingWanVna) WanVnaValue {
	tflog.Debug(ctx, "wanVnaSdkToTerraform")

	r_attr_type := WanVnaValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"enabled": types.BoolValue(d.GetEnabled()),
	}
	r, e := NewWanVnaValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
