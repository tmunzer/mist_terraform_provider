package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"
)

func dhcpSnoopingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d DhcpSnoopingValue) models.DhcpSnooping {
	data := models.NewDhcpSnooping()
	data.SetAllNetworks(d.AllNetworks.ValueBool())
	data.SetEnableArpSpoofCheck(d.EnableArpSpoofCheck.ValueBool())
	data.SetEnableIpSourceGuard(d.EnableIpSourceGuard.ValueBool())
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetNetworks(mist_transform.ListOfStringTerraformToSdk(ctx, d.Networks))
	return *data
}
