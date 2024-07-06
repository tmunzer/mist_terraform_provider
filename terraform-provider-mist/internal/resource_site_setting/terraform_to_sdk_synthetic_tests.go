package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"
)

func syntheticTestVlansTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, l basetypes.ListValue) []models.SyntheticTestProperties {
	tflog.Debug(ctx, "syntheticTestVlansTerraformToSdk")
	var data_list []models.SyntheticTestProperties
	for _, v := range l.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(VlansValue)
		data := models.NewSyntheticTestProperties()
		data.SetCustomTestUrls(mist_transform.ListOfStringTerraformToSdk(ctx, plan.CustomTestUrls))
		data.SetDisabled(plan.Disabled.ValueBool())
		data.SetVlanIds(mist_transform.ListOfIntTerraformToSdk(ctx, plan.VlanIds))

		data_list = append(data_list, *data)
	}
	return data_list
}

func syntheticTestTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SyntheticTestValue) models.SyntheticTestConfig {
	tflog.Debug(ctx, "syntheticTestTerraformToSdk")
	data := models.NewSyntheticTestConfig()

	data.SetDisabled(d.Disabled.ValueBool())

	vlans := syntheticTestVlansTerraformToSdk(ctx, diags, d.Vlans)
	data.SetVlans(vlans)

	return *data
}
