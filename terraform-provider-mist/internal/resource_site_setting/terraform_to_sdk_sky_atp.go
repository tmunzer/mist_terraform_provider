package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"mistapi/models"
)

func skyAtpTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SkyatpValue) models.SiteSettingSkyatp {
	tflog.Debug(ctx, "skyAtpTerraformToSdk")
	data := models.NewSiteSettingSkyatp()

	data.SetEnabled(d.Enabled.ValueBool())
	data.SetSendIpMacMapping(d.SendIpMacMapping.ValueBool())

	return *data
}
