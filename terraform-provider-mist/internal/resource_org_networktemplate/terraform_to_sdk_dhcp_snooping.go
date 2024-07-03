package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func dhcpSnoopingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d DhcpSnoopingValue) mistapigo.DhcpSnooping {
	data := mistapigo.NewDhcpSnooping()
	data.SetAllNetworks(d.AllNetworks.ValueBool())
	data.SetEnableArpSpoofCheck(d.EnableArpSpoofCheck.ValueBool())
	data.SetEnableIpSourceGuard(d.EnableIpSourceGuard.ValueBool())
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetNetworks(mist_transform.ListOfStringTerraformToSdk(ctx, d.Networks))
	return *data
}
