package resource_site_wlan

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func bonjourServicesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.MapValue) map[string]mistapigo.WlanBonjourServiceProperties {
	data_map := make(map[string]mistapigo.WlanBonjourServiceProperties)
	for k, v := range plan.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(ServicesValue)
		v_data := mistapigo.NewWlanBonjourServiceProperties()
		v_data.SetDisableLocal(v_plan.DisableLocal.ValueBool())
		v_data.SetRadiusGroups(mist_transform.ListOfStringTerraformToSdk(ctx, v_plan.RadiusGroups))
		v_data.SetScope(mistapigo.WlanBonjourServicePropertiesScope(v_plan.Scope.ValueString()))

		data_map[k] = *v_data
	}
	return data_map
}
func bonjourTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan BonjourValue) mistapigo.WlanBonjour {

	additional_vlan_ids := mist_transform.ListOfIntTerraformToSdk(ctx, plan.AdditionalVlanIds)
	services := bonjourServicesTerraformToSdk(ctx, diags, plan.Services)
	data := mistapigo.NewWlanBonjour(additional_vlan_ids, services)
	data.SetEnabled(plan.Enabled.ValueBool())

	return *data
}
