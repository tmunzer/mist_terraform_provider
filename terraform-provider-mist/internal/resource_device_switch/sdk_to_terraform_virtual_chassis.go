package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func virtualChassisMembersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistapigo.SwitchVirtualChassisMember) basetypes.ListValue {

	var data_list = []MembersValue{}
	for _, v := range d {
		data_map_value := map[string]attr.Value{
			"mac":       types.StringValue(v.GetMac()),
			"vc_role":   types.StringValue(string(v.GetVcRole())),
			"member_id": types.Int64Value(int64(v.GetMemberId())),
		}
		data, e := NewMembersValue(MembersValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, MembersValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

func virtualChassisSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.SwitchVirtualChassis) VirtualChassisValue {

	state_value_map_attr_type := VirtualChassisValue{}.AttributeTypes(ctx)

	state_value_map_attr_value := map[string]attr.Value{
		"members":        virtualChassisMembersSdkToTerraform(ctx, diags, d.GetMembers()),
		"preprovisioned": types.BoolValue(d.GetPreprovisioned()),
	}
	r, e := NewVirtualChassisValue(state_value_map_attr_type, state_value_map_attr_value)
	diags.Append(e...)

	return r
}
