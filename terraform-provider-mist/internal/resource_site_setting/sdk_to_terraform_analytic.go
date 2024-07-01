package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func analyticSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.SiteSettingAnalytic) AnalyticValue {
	tflog.Debug(ctx, "analyticSdkToTerraform")

	r_attr_type := AnalyticValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"enabled": types.BoolValue(d.GetEnabled()),
	}
	r, e := NewAnalyticValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
