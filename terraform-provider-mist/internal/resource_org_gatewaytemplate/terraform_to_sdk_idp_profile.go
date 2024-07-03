package resource_org_gatewaytemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func idpProfileMatchingSeverityTerraformToSdk(ctx context.Context, list basetypes.ListValue) []mistapigo.IdpProfileMatchingSeverityValue {
	tflog.Debug(ctx, "idpProfileMatchingSeverityTerraformToSdk")
	var items []mistapigo.IdpProfileMatchingSeverityValue
	for _, item := range list.Elements() {
		s, _ := mistapigo.NewIdpProfileMatchingSeverityValueFromValue(item.String())
		items = append(items, *s)
	}
	return items
}

func idpProfileMatchingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistapigo.IdpProfileMatching {
	tflog.Debug(ctx, "idpProfileMatchingTerraformToSdk")
	data := *mistapigo.NewIdpProfileMatching()
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

func idpProfileOverwritesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.IdpProfileOverwrite {
	tflog.Debug(ctx, "idpProfileOverwritesTerraformToSdk")
	var data_list []mistapigo.IdpProfileOverwrite
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(OverwritesValue)
		data := mistapigo.NewIdpProfileOverwrite()
		data.SetAction(mistapigo.IdpProfileAction(plan.Action.ValueString()))
		data.SetMatching(idpProfileMatchingTerraformToSdk(ctx, diags, plan.IpdProfileOverwriteMatching))

		data_list = append(data_list, *data)
	}
	return data_list
}

func idpProfileTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.IdpProfile {
	tflog.Debug(ctx, "idpProfileTerraformToSdk")
	data_map := make(map[string]mistapigo.IdpProfile)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(IdpProfilesValue)

		overwrites := idpProfileOverwritesTerraformToSdk(ctx, diags, plan.Overwrites)

		data := *mistapigo.NewIdpProfile()
		data.SetBaseProfile(mistapigo.IdpProfileBaseProfile(plan.BaseProfile.ValueString()))
		data.SetName(plan.Name.ValueString())
		data.SetOverwrites(overwrites)

		data_map[k] = data
	}
	return data_map
}
