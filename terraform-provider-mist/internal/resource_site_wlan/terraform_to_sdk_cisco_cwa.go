package resource_site_wlan

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func ciscoCwaTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan CiscoCwaValue) models.WlanCiscoCwa {

	data := models.NewWlanCiscoCwa()
	data.SetAllowedHostnames(mist_transform.ListOfStringTerraformToSdk(ctx, plan.AllowedHostnames))
	data.SetAllowedSubnets(mist_transform.ListOfStringTerraformToSdk(ctx, plan.AllowedSubnets))
	data.SetBlockedSubnets(mist_transform.ListOfStringTerraformToSdk(ctx, plan.BlockedSubnets))
	data.SetEnabled(plan.Enabled.ValueBool())

	return *data
}
