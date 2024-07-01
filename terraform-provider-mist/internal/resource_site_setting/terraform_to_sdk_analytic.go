package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func analyticTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d AnalyticValue) *mistsdkgo.SiteSettingAnalytic {
	tflog.Debug(ctx, "analyticTerraformToSdk")
	data := mistsdkgo.NewSiteSettingAnalytic()

	data.SetEnabled(d.Enabled.ValueBool())

	return data
}
