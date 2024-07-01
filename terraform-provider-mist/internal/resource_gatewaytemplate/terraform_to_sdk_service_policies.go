package resource_gatewaytemplate

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	mist_transform "terraform-provider-mist/internal/provider/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func servicePolicyAppqoeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.ServicePolicyAppqoe {
	data := *mistsdkgo.NewServicePolicyAppqoe()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewAppqoeValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetEnabled(plan.Enabled.ValueBool())
		return data
	}
}

func servicePolicyEwfRuleTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.ServicePolicyEwfRule {
	var data_list []mistsdkgo.ServicePolicyEwfRule
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(EwfValue)
		data := mistsdkgo.NewServicePolicyEwfRule()
		data.SetAlertOnly(plan.AlertOnly.ValueBool())
		data.SetBlockMessage(plan.BlockMessage.ValueString())
		data.SetEnabled(plan.Enabled.ValueBool())
		data.SetProfile(mistsdkgo.ServicePolicyEwfRuleProfile(plan.Profile.ValueString()))

		data_list = append(data_list, *data)
	}
	return data_list
}

func idpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.IdpConfig {
	data := *mistsdkgo.NewIdpConfig()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewIdpValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetAlertOnly(plan.AlertOnly.ValueBool())
		data.SetEnabled(plan.Enabled.ValueBool())
		data.SetIdpprofileId(plan.IdpprofileId.ValueString())
		data.SetProfile(plan.Profile.ValueString())
		return data
	}
}

func servicePoliciesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.ServicePolicy {
	var data_list []mistsdkgo.ServicePolicy
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(ServicePoliciesValue)
		data := mistsdkgo.NewServicePolicy()

		data.SetAction(mistsdkgo.AllowDeny(plan.Action.ValueString()))

		appqoe := servicePolicyAppqoeTerraformToSdk(ctx, diags, plan.Appqoe)
		data.SetAppqoe(appqoe)

		ewf := servicePolicyEwfRuleTerraformToSdk(ctx, diags, plan.Ewf)
		data.SetEwf(ewf)

		idp := idpConfigTerraformToSdk(ctx, diags, plan.Idp)
		data.SetIdp(idp)

		data.SetLocalRouting(plan.LocalRouting.ValueBool())
		data.SetName(plan.Name.ValueString())
		data.SetPathPreferences(plan.PathPreferences.ValueString())
		data.SetServicepolicyId(plan.ServicepolicyId.ValueString())
		data.SetServices(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Services))
		data.SetTenants(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Tenants))

		data_list = append(data_list, *data)
	}
	return data_list
}
