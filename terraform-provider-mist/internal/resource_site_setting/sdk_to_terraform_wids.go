package resource_site_setting

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func widsAuthFailureSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SiteWidsRepeatedAuthFailures) basetypes.ObjectValue {
	tflog.Debug(ctx, "widsAuthFailureSdkToTerraform")

	r_attr_type := RepeatedAuthFailuresValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"duration":  types.Int64Value(int64(d.GetDuration())),
		"threshold": types.Int64Value(int64(d.GetThreshold())),
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

func widsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SiteWids) WidsValue {
	tflog.Debug(ctx, "widsSdkToTerraform")

	r_attr_type := WidsValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"repeated_auth_failures": widsAuthFailureSdkToTerraform(ctx, diags, d.GetRepeatedAuthFailures()),
	}
	r, e := NewWidsValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
