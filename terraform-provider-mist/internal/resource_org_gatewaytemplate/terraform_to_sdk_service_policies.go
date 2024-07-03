package resource_org_gatewaytemplate

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func servicePolicyAppqoeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistapigo.ServicePolicyAppqoe {
	tflog.Debug(ctx, "servicePolicyAppqoeTerraformToSdk")
	data := *mistapigo.NewServicePolicyAppqoe()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewAppqoeValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetEnabled(plan.Enabled.ValueBool())
		return data
	}
}

func servicePolicyEwfRuleTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.ServicePolicyEwfRule {
	tflog.Debug(ctx, "servicePolicyEwfRuleTerraformToSdk")
	var data_list []mistapigo.ServicePolicyEwfRule
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(EwfValue)
		data := mistapigo.NewServicePolicyEwfRule()
		data.SetAlertOnly(plan.AlertOnly.ValueBool())
		data.SetBlockMessage(plan.BlockMessage.ValueString())
		data.SetEnabled(plan.Enabled.ValueBool())
		data.SetProfile(mistapigo.ServicePolicyEwfRuleProfile(plan.Profile.ValueString()))

		data_list = append(data_list, *data)
	}
	return data_list
}

func idpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistapigo.IdpConfig {
	tflog.Debug(ctx, "idpConfigTerraformToSdk")
	data := *mistapigo.NewIdpConfig()
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

func servicePoliciesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.ServicePolicy {
	tflog.Debug(ctx, "servicePoliciesTerraformToSdk")
	var data_list []mistapigo.ServicePolicy
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(ServicePoliciesValue)
		data := mistapigo.NewServicePolicy()

		data.SetAction(mistapigo.AllowDeny(plan.Action.ValueString()))

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
