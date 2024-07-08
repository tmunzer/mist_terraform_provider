package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"
)

func dhcpSnoopingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d DhcpSnoopingValue) *models.DhcpSnooping {
	data := models.DhcpSnooping{}
	if !d.AllNetworks.IsNull() && !d.AllNetworks.IsUnknown() {
		data.AllNetworks = models.ToPointer(d.AllNetworks.ValueBool())
	}
	if !d.EnableArpSpoofCheck.IsNull() && !d.EnableArpSpoofCheck.IsUnknown() {
		data.EnableArpSpoofCheck = models.ToPointer(d.EnableArpSpoofCheck.ValueBool())
	}
	if !d.EnableIpSourceGuard.IsNull() && !d.EnableIpSourceGuard.IsUnknown() {
		data.EnableIpSourceGuard = models.ToPointer(d.EnableIpSourceGuard.ValueBool())
	}
	if !d.Enabled.IsNull() && !d.Enabled.IsUnknown() {
		data.Enabled = models.ToPointer(d.Enabled.ValueBool())
	}
	if !d.Networks.IsNull() && !d.Networks.IsUnknown() {
		data.Networks = mist_transform.ListOfStringTerraformToSdk(ctx, d.Networks)
	}
	return &data
}
