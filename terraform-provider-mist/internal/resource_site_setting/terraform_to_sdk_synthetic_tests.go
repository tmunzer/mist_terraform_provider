package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func syntheticTestVlansTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, l basetypes.ListValue) []mistapigo.SyntheticTestProperties {
	tflog.Debug(ctx, "syntheticTestVlansTerraformToSdk")
	var data_list []mistapigo.SyntheticTestProperties
	for _, v := range l.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(VlansValue)
		data := mistapigo.NewSyntheticTestProperties()
		data.SetCustomTestUrls(mist_transform.ListOfStringTerraformToSdk(ctx, plan.CustomTestUrls))
		data.SetDisabled(plan.Disabled.ValueBool())
		data.SetVlanIds(mist_transform.ListOfIntTerraformToSdk(ctx, plan.VlanIds))

		data_list = append(data_list, *data)
	}
	return data_list
}

func syntheticTestTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SyntheticTestValue) mistapigo.SyntheticTestConfig {
	tflog.Debug(ctx, "syntheticTestTerraformToSdk")
	data := mistapigo.NewSyntheticTestConfig()

	data.SetDisabled(d.Disabled.ValueBool())

	vlans := syntheticTestVlansTerraformToSdk(ctx, diags, d.Vlans)
	data.SetVlans(vlans)

	return *data
}
