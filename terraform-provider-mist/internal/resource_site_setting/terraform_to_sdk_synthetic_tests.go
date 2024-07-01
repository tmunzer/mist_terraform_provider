package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func syntheticTestVlansTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, l basetypes.ListValue) []mistsdkgo.SyntheticTestProperties {
	tflog.Debug(ctx, "syntheticTestVlansTerraformToSdk")
	var data_list []mistsdkgo.SyntheticTestProperties
	for _, v := range l.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(VlansValue)
		data := mistsdkgo.NewSyntheticTestProperties()
		data.SetCustomTestUrls(mist_transform.ListOfStringTerraformToSdk(ctx, plan.CustomTestUrls))
		data.SetDisabled(plan.Disabled.ValueBool())
		data.SetVlanIds(mist_transform.ListOfIntTerraformToSdk(ctx, plan.VlanIds))

		data_list = append(data_list, *data)
	}
	return data_list
}

func syntheticTestTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SyntheticTestValue) mistsdkgo.SyntheticTestConfig {
	tflog.Debug(ctx, "syntheticTestTerraformToSdk")
	data := mistsdkgo.NewSyntheticTestConfig()

	data.SetDisabled(d.Disabled.ValueBool())

	vlans := syntheticTestVlansTerraformToSdk(ctx, diags, d.Vlans)
	data.SetVlans(vlans)

	return *data
}
