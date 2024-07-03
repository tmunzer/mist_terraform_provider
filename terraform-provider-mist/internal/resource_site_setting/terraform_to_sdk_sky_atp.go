package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func skyAtpTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SkyatpValue) mistapigo.SiteSettingSkyatp {
	tflog.Debug(ctx, "skyAtpTerraformToSdk")
	data := mistapigo.NewSiteSettingSkyatp()

	data.SetEnabled(d.Enabled.ValueBool())
	data.SetSendIpMacMapping(d.SendIpMacMapping.ValueBool())

	return *data
}
