package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func widsAuthFailureTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, o basetypes.ObjectValue) mistapigo.SiteWidsRepeatedAuthFailures {
	tflog.Debug(ctx, "widsAuthFailureTerraformToSdk")
	data := mistapigo.NewSiteWidsRepeatedAuthFailures()
	if o.IsNull() || o.IsUnknown() {
		return *data
	} else {
		d := NewRepeatedAuthFailuresValueMust(o.AttributeTypes(ctx), o.Attributes())
		data.SetDuration(int32(d.Duration.ValueInt64()))
		data.SetThreshold(int32(d.Threshold.ValueInt64()))
		return *data
	}
}

func widsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d WidsValue) mistapigo.SiteWids {
	tflog.Debug(ctx, "widsTerraformToSdk")
	data := mistapigo.NewSiteWids()

	if !d.IsNull() {
		repeated_auth_failures := widsAuthFailureTerraformToSdk(ctx, diags, d.RepeatedAuthFailures)
		data.SetRepeatedAuthFailures(repeated_auth_failures)
	}

	return *data
}
