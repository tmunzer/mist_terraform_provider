package resource_org_gatewaytemplate

import (
	"context"

	"mistapi/models"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/google/uuid"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func servicePolicyAppqoeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.ServicePolicyAppqoe {
	tflog.Debug(ctx, "servicePolicyAppqoeTerraformToSdk")
	data := models.ServicePolicyAppqoe{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan := NewAppqoeValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.Enabled = models.ToPointer(plan.Enabled.ValueBool())
		return &data
	}
}

func servicePolicyEwfRuleTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.ServicePolicyEwfRule {
	tflog.Debug(ctx, "servicePolicyEwfRuleTerraformToSdk")
	var data_list []models.ServicePolicyEwfRule
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(EwfValue)
		data := models.ServicePolicyEwfRule{}
		data.AlertOnly = models.ToPointer(plan.AlertOnly.ValueBool())
		data.BlockMessage = models.ToPointer(plan.BlockMessage.ValueString())
		data.Enabled = models.ToPointer(plan.Enabled.ValueBool())
		data.Profile = models.ToPointer(models.ServicePolicyEwfRuleProfileEnum(plan.Profile.ValueString()))

		data_list = append(data_list, data)
	}
	return data_list
}

func idpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.IdpConfig {
	tflog.Debug(ctx, "idpConfigTerraformToSdk")
	data := models.IdpConfig{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan := NewIdpValueMust(d.AttributeTypes(ctx), d.Attributes())

		if len(plan.IdpprofileId.ValueString()) > 0 {
			idp_profile_id, e := uuid.Parse(plan.IdpprofileId.ValueString())
			if e != nil {
				diags.AddError("Unable to convert IdpprofileId", e.Error())
			} else {
				data.IdpprofileId = models.ToPointer(idp_profile_id)
			}
		}

		data.AlertOnly = models.ToPointer(plan.AlertOnly.ValueBool())
		data.Enabled = models.ToPointer(plan.Enabled.ValueBool())
		data.Profile = models.ToPointer(plan.Profile.ValueString())
		return &data
	}
}

func servicePoliciesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.ServicePolicy {
	tflog.Debug(ctx, "servicePoliciesTerraformToSdk")
	var data_list []models.ServicePolicy
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(ServicePoliciesValue)
		data := models.ServicePolicy{}

		data.Action = models.ToPointer(models.AllowDenyEnum(plan.Action.ValueString()))

		appqoe := servicePolicyAppqoeTerraformToSdk(ctx, diags, plan.Appqoe)
		data.Appqoe = appqoe

		ewf := servicePolicyEwfRuleTerraformToSdk(ctx, diags, plan.Ewf)
		data.Ewf = ewf

		idp := idpConfigTerraformToSdk(ctx, diags, plan.Idp)
		data.Idp = idp

		data.LocalRouting = models.ToPointer(plan.LocalRouting.ValueBool())
		data.Name = models.ToPointer(plan.Name.ValueString())
		data.PathPreferences = models.ToPointer(plan.PathPreferences.ValueString())
		if len(plan.ServicepolicyId.ValueString()) > 0 {
			service_policy_id, e := uuid.Parse(plan.ServicepolicyId.ValueString())
			if e != nil {
				diags.AddError("Unable to convert ServicepolicyId", e.Error())
			} else {
				data.ServicepolicyId = models.ToPointer(service_policy_id)
			}
		}

		data.Services = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Services)
		data.Tenants = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Tenants)

		data_list = append(data_list, data)
	}
	return data_list
}
