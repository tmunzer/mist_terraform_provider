package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"
)

func aclPolicyActionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.AclPolicyAction {

	var data []models.AclPolicyAction
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(ActionsValue)
		data_item := models.AclPolicyAction{}
		data_item.Action = models.ToPointer(models.AllowDenyEnum(v_plan.Action.ValueString()))
		data_item.DstTag = models.ToPointer(v_plan.DstTag.ValueString())
		data = append(data, data_item)
	}
	return data
}

func aclPoliciesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.AclPolicy {

	var data []models.AclPolicy
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(AclPoliciesValue)
		actions := aclPolicyActionsTerraformToSdk(ctx, diags, v_plan.Actions)
		data_item := models.AclPolicy{}
		data_item.Name = models.ToPointer(v_plan.Name.ValueString())
		data_item.Actions = actions
		data_item.SrcTags = mist_transform.ListOfStringTerraformToSdk(ctx, v_plan.SrcTags)

		data = append(data, data_item)
	}
	return data
}
