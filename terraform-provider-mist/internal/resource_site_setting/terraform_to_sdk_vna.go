package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"mistapi/models"
)

func vnaTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d VnaValue) models.SiteSettingVna {
	tflog.Debug(ctx, "vnaTerraformToSdk")
	data := models.NewSiteSettingVna()

	data.SetEnabled(d.Enabled.ValueBool())

	return *data
}
