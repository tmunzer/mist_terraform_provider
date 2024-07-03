package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func proxyTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ProxyValue) mistapigo.Proxy {
	tflog.Debug(ctx, "proxyTerraformToSdk")
	data := mistapigo.NewProxy()

	data.SetUrl(d.Url.ValueString())

	return *data
}
