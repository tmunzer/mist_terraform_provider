package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"
)

func dhcpSnoopingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d DhcpSnoopingValue) *models.DhcpSnooping {
	data := models.DhcpSnooping{}
	data.AllNetworks = models.ToPointer(d.AllNetworks.ValueBool())
	data.EnableArpSpoofCheck = models.ToPointer(d.EnableArpSpoofCheck.ValueBool())
	data.EnableIpSourceGuard = models.ToPointer(d.EnableIpSourceGuard.ValueBool())
	data.Enabled = models.ToPointer(d.Enabled.ValueBool())
	data.Networks = mist_transform.ListOfStringTerraformToSdk(ctx, d.Networks)
	return &data
}
