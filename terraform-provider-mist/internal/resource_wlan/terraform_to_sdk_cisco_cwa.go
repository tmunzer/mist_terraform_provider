package resource_wlan

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func ciscoCwaTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan CiscoCwaValue) mistsdkgo.WlanCiscoCwa {

	data := mistsdkgo.NewWlanCiscoCwa()
	data.SetAllowedHostnames(mist_transform.ListOfStringTerraformToSdk(ctx, plan.AllowedHostnames))
	data.SetAllowedSubnets(mist_transform.ListOfStringTerraformToSdk(ctx, plan.AllowedSubnets))
	data.SetBlockedSubnets(mist_transform.ListOfStringTerraformToSdk(ctx, plan.BlockedSubnets))
	data.SetEnabled(plan.Enabled.ValueBool())

	return *data
}
