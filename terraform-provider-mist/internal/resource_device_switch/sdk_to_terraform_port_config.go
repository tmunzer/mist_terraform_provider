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

func switchMatchingRulesPortMirroringSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.SwitchPortMirroringProperty) basetypes.MapValue {
	map_item_value := make(map[string]attr.Value)
	map_item_type := PortMirroringValue{}.Type(ctx)
	item_type := PortMirroringValue{}.AttributeTypes(ctx)
	for k, v := range d {

		var item_value = map[string]attr.Value{
			"input_networks_ingress": mist_transform.ListOfStringSdkToTerraform(ctx, v.GetInputNetworksIngress()),
			"input_port_ids_egress":  mist_transform.ListOfStringSdkToTerraform(ctx, v.GetInputPortIdsEgress()),
			"input_port_ids_ingress": mist_transform.ListOfStringSdkToTerraform(ctx, v.GetInputPortIdsIngress()),
			"output_port_id":         types.StringValue(v.GetOutputPortId()),
		}
		item_obj, e := NewPortMirroringValue(item_type, item_value)
		diags.Append(e...)
		map_item_value[k] = item_obj
	}
	r, e := types.MapValueFrom(ctx, map_item_type, map_item_value)
	diags.Append(e...)
	return r
}
func portConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.JunosPortConfig) basetypes.MapValue {
	map_item_value := make(map[string]attr.Value)
	map_item_type := PortConfigValue{}.Type(ctx)
	item_type := PortConfigValue{}.AttributeTypes(ctx)
	for k, v := range d {

		var item_value = map[string]attr.Value{
			"ae_disable_lacp":    types.BoolValue(v.GetAeDisableLacp()),
			"ae_idx":             types.Int64Value(int64(v.GetAeIdx())),
			"ae_lacp_slow":       types.BoolValue(v.GetAeLacpSlow()),
			"aggregated":         types.BoolValue(v.GetAggregated()),
			"critical":           types.BoolValue(v.GetCritical()),
			"description":        types.StringValue(v.GetDescription()),
			"disable_autoneg":    types.BoolValue(v.GetDisableAutoneg()),
			"duplex":             types.StringValue(string(v.GetDuplex())),
			"dynamic_usage":      types.StringValue(v.GetDynamicUsage()),
			"esilag":             types.BoolValue(v.GetEsilag()),
			"mtu":                types.Int64Value(int64(v.GetMtu())),
			"no_local_overwrite": types.BoolValue(v.GetNoLocalOverwrite()),
			"poe_disabled":       types.BoolValue(v.GetPoeDisabled()),
			"speed":              types.StringValue(string(v.GetSpeed())),
			"usage":              types.StringValue(v.GetUsage()),
		}
		item_obj, e := NewPortConfigValue(item_type, item_value)
		diags.Append(e...)
		map_item_value[k] = item_obj
	}
	r, e := types.MapValueFrom(ctx, map_item_type, map_item_value)
	diags.Append(e...)
	return r
}
