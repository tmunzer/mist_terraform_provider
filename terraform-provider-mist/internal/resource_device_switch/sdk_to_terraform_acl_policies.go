package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func aclPolicyActionsToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.AclPolicyAction) basetypes.ListValue {
	var data_list []attr.Value
	data_list_type := ActionsValue{}.AttributeTypes(ctx)
	for _, v := range d {
		rc_srv_state_value := map[string]attr.Value{
			"action":  types.StringValue(string(v.GetAction())),
			"dst_tag": types.StringValue(v.GetDstTag()),
		}
		acct_server, e := NewActionsValue(data_list_type, rc_srv_state_value)
		diags.Append(e...)

		data_list = append(data_list, acct_server)
	}

	state_list_type := ActionsValue{}.Type(ctx)
	state_list, e := types.ListValueFrom(ctx, state_list_type, data_list)
	diags.Append(e...)
	return state_list
}

func aclPoliciesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.AclPolicy) basetypes.ListValue {
	var data_list []attr.Value
	data_list_type := AclPoliciesValue{}.AttributeTypes(ctx)
	for _, v := range d {
		actions := aclPolicyActionsToTerraform(ctx, diags, v.GetActions())
		rc_srv_state_value := map[string]attr.Value{
			"actions":  actions,
			"name":     types.StringValue(v.GetName()),
			"src_tags": mist_transform.ListOfStringSdkToTerraform(ctx, v.GetSrcTags()),
		}
		acct_server, e := NewAclPoliciesValue(data_list_type, rc_srv_state_value)
		diags.Append(e...)

		data_list = append(data_list, acct_server)
	}

	state_list_type := AclPoliciesValue{}.Type(ctx)
	state_list, e := types.ListValueFrom(ctx, state_list_type, data_list)
	diags.Append(e...)

	return state_list
}