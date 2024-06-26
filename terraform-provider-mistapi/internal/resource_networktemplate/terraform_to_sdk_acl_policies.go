package resource_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mistapi/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mistapi/internal/provider/utils/transform"
)

func aclPolicyActionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.AclPolicyAction {

	var data []mistsdkgo.AclPolicyAction
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(ActionsValue)
		data_item := mistsdkgo.NewAclPolicyAction()
		data_item.SetAction(mistsdkgo.AllowDeny(v_plan.Action.ValueString()))
		data_item.SetDstTag(v_plan.DstTag.ValueString())
		data = append(data, *data_item)
	}
	return data
}

func aclPoliciesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.AclPolicy {

	var data []mistsdkgo.AclPolicy
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(AclPoliciesValue)
		actions := aclPolicyActionsTerraformToSdk(ctx, diags, v_plan.Actions)
		data_item := mistsdkgo.NewAclPolicy()
		data_item.SetName(v_plan.Name.ValueString())
		data_item.SetActions(actions)
		data_item.SetSrcTags(mist_transform.ListOfStringTerraformToSdk(ctx, v_plan.SrcTags))

		data = append(data, *data_item)
	}
	return data
}
