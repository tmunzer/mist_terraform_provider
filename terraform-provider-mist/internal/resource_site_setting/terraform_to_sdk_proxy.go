package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"mistapi/models"
)

func proxyTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ProxyValue) models.Proxy {
	tflog.Debug(ctx, "proxyTerraformToSdk")
	data := models.NewProxy()

	data.SetUrl(d.Url.ValueString())

	return *data
}
