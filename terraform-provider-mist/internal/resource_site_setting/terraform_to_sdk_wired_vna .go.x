package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"mistapi/models"
)

func wiredVnaTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d WiredVnaValue) *models.SiteSettingWiredVna {
	tflog.Debug(ctx, "wiredVnaTerraformToSdk")
	data := models.SiteSettingWiredVna{}

	data.Enabled = d.Enabled.ValueBool()

	return &data
}