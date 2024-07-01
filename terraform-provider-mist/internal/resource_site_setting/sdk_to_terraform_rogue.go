package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func rogueSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.SiteRogue) RogueValue {
	tflog.Debug(ctx, "rogueSdkToTerraform")

	r_attr_type := RogueValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"enabled":            types.BoolValue(d.GetEnabled()),
		"honeypot_enabled":   types.BoolValue(d.GetHoneypotEnabled()),
		"min_duration":       types.Int64Value(int64(d.GetMinDuration())),
		"min_rssi":           types.Int64Value(int64(d.GetMinRssi())),
		"whitelisted_bssids": mist_transform.ListOfStringSdkToTerraform(ctx, d.GetWhitelistedBssids()),
		"whitelisted_ssids":  mist_transform.ListOfStringSdkToTerraform(ctx, d.GetWhitelistedSsids()),
	}
	r, e := NewRogueValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
