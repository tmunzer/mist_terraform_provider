package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func dhcpSnoopingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d DhcpSnoopingValue) mistsdkgo.DhcpSnooping {
	data := mistsdkgo.NewDhcpSnooping()
	data.SetAllNetworks(d.AllNetworks.ValueBool())
	data.SetEnableArpSpoofCheck(d.EnableArpSpoofCheck.ValueBool())
	data.SetEnableIpSourceGuard(d.EnableIpSourceGuard.ValueBool())
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetNetworks(mist_transform.ListOfStringTerraformToSdk(ctx, d.Networks))
	return *data
}
