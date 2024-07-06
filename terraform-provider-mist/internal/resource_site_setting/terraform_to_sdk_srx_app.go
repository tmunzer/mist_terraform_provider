package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"mistapi/models"
)

func srxAppTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SrxAppValue) models.SiteSettingSrxApp {
	tflog.Debug(ctx, "srxAppTerraformToSdk")
	data := models.NewSiteSettingSrxApp()

	data.SetEnabled(d.Enabled.ValueBool())

	return *data
}
