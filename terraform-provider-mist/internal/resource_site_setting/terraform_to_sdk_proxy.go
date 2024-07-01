package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func proxyTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ProxyValue) mistsdkgo.Proxy {
	tflog.Debug(ctx, "proxyTerraformToSdk")
	data := mistsdkgo.NewProxy()

	data.SetUrl(d.Url.ValueString())

	return *data
}
