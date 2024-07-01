package resource_wlan

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func bonjourServicesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.MapValue) map[string]mistsdkgo.WlanBonjourServiceProperties {
	data_map := make(map[string]mistsdkgo.WlanBonjourServiceProperties)
	for k, v := range plan.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(ServicesValue)
		v_data := mistsdkgo.NewWlanBonjourServiceProperties()
		v_data.SetDisableLocal(v_plan.DisableLocal.ValueBool())
		v_data.SetRadiusGroups(mist_transform.ListOfStringTerraformToSdk(ctx, v_plan.RadiusGroups))
		v_data.SetScope(mistsdkgo.WlanBonjourServicePropertiesScope(v_plan.Scope.ValueString()))

		data_map[k] = *v_data
	}
	return data_map
}
func bonjourTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan BonjourValue) mistsdkgo.WlanBonjour {

	additional_vlan_ids := mist_transform.ListOfIntTerraformToSdk(ctx, plan.AdditionalVlanIds)
	services := bonjourServicesTerraformToSdk(ctx, diags, plan.Services)
	data := mistsdkgo.NewWlanBonjour(additional_vlan_ids, services)
	data.SetEnabled(plan.Enabled.ValueBool())

	return *data
}
