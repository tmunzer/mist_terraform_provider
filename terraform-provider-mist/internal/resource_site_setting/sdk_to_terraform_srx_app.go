package resource_site_setting

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func srxAppSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SiteSettingSrxApp) SrxAppValue {
	tflog.Debug(ctx, "srxAppSdkToTerraform")

	r_attr_type := SrxAppValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"enabled": types.BoolValue(d.GetEnabled()),
	}
	r, e := NewSrxAppValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
