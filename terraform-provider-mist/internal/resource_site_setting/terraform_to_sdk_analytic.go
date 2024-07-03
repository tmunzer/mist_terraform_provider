package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func analyticTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d AnalyticValue) *mistapigo.SiteSettingAnalytic {
	tflog.Debug(ctx, "analyticTerraformToSdk")
	data := mistapigo.NewSiteSettingAnalytic()

	data.SetEnabled(d.Enabled.ValueBool())

	return data
}
