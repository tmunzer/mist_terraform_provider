package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func wiredVnaTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d WiredVnaValue) mistapigo.SiteSettingWiredVna {
	tflog.Debug(ctx, "wiredVnaTerraformToSdk")
	data := mistapigo.NewSiteSettingWiredVna()

	data.SetEnabled(d.Enabled.ValueBool())

	return *data
}
