package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func portMirroringSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.SwitchPortMirroring) basetypes.MapValue {
	data_map_attr_type := PortMirroringValue{}.AttributeTypes(ctx)
	data_map_value := make(map[string]attr.Value)
	for k, v := range d {
		var data_map_item = map[string]attr.Value{
			"input_networks_ingress": mist_transform.ListOfStringSdkToTerraform(ctx, v.GetInputNetworksIngress()),
			"input_port_ids_egress":  mist_transform.ListOfStringSdkToTerraform(ctx, v.GetInputPortIdsEgress()),
			"input_port_ids_ingress": mist_transform.ListOfStringSdkToTerraform(ctx, v.GetInputPortIdsIngress()),
			"output_network":         types.StringValue(v.GetOutputNetwork()),
			"output_port_id":         types.StringValue(v.GetOutputPortId()),
		}
		data_map_item_object, e := NewPortMirroringValue(data_map_attr_type, data_map_item)
		diags.Append(e...)
		data_map_value[k] = data_map_item_object
	}
	state_type := PortMirroringValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, data_map_value)
	diags.Append(e...)
	return state_result
}
