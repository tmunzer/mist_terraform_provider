package resource_gatewaytemplate

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func idpProfileMatchingSeverityTerraformToSdk(ctx context.Context, list basetypes.ListValue) []mistsdkgo.IdpProfileMatchingSeverityValue {
	var items []mistsdkgo.IdpProfileMatchingSeverityValue
	for _, item := range list.Elements() {
		s, _ := mistsdkgo.NewIdpProfileMatchingSeverityValueFromValue(item.String())
		items = append(items, *s)
	}
	return items
}

func idpProfileMatchingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.IdpProfileMatching {
	data := *mistsdkgo.NewIdpProfileMatching()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		var d_interface interface{} = d
		plan := d_interface.(IpdProfileOverwriteMatchingValue)

		data.SetAttackName(mist_transform.ListOfStringTerraformToSdk(ctx, plan.AttackName))
		data.SetDstSubnet(mist_transform.ListOfStringTerraformToSdk(ctx, plan.DstSubnet))
		data.SetSeverity(idpProfileMatchingSeverityTerraformToSdk(ctx, plan.Severity))

		return data
	}
}

func idpProfileOverwritesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.IdpProfileOverwrite {
	var data_list []mistsdkgo.IdpProfileOverwrite
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(OverwritesValue)
		data := mistsdkgo.NewIdpProfileOverwrite()
		data.SetAction(mistsdkgo.IdpProfileAction(plan.Action.ValueString()))
		data.SetMatching(idpProfileMatchingTerraformToSdk(ctx, diags, plan.IpdProfileOverwriteMatching))

		data_list = append(data_list, *data)
	}
	return data_list
}

func idpProfileTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.IdpProfile {
	data_map := make(map[string]mistsdkgo.IdpProfile)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(IdpProfilesValue)

		overwrites := idpProfileOverwritesTerraformToSdk(ctx, diags, plan.Overwrites)

		data := *mistsdkgo.NewIdpProfile()
		data.SetBaseProfile(mistsdkgo.IdpProfileBaseProfile(plan.BaseProfile.ValueString()))
		data.SetName(plan.Name.ValueString())
		data.SetOverwrites(overwrites)

		data_map[k] = data
	}
	return data_map
}
