package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"mistapi/models"
)

func wanVnaTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d WanVnaValue) models.SiteSettingWanVna {
	tflog.Debug(ctx, "wanVnaTerraformToSdk")
	data := models.NewSiteSettingWanVna()

	data.SetEnabled(d.Enabled.ValueBool())

	return *data
}
