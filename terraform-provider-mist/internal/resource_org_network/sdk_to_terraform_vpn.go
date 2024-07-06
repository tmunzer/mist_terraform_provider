package resource_org_network

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func VpnSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.NetworkVpnAccessConfig) basetypes.MapValue {

	state_value_map_attr_type := VpnAccessValue{}.AttributeTypes(ctx)
	state_value_map_value := make(map[string]attr.Value)
	for k, v := range d {
		state_value_map_attr_value := map[string]attr.Value{
			"advertised_subnet":             types.StringValue(*v.AdvertisedSubnet),
			"allow_ping":                    types.BoolValue(*v.AllowPing),
			"destination_nat":               destinationNatSdkToTerraform(ctx, diags, v.DestinationNat),
			"nat_pool":                      types.StringValue(*v.NatPool),
			"no_readvertise_to_lan_bgp":     types.BoolValue(*v.NoReadvertiseToLanBgp),
			"no_readvertise_to_lan_ospf":    types.BoolValue(*v.NoReadvertiseToLanOspf),
			"no_readvertise_to_overlay":     types.BoolValue(*v.NoReadvertiseToOverlay),
			"other_vrfs":                    mist_transform.ListOfStringSdkToTerraform(ctx, v.OtherVrfs),
			"routed":                        types.BoolValue(*v.Routed),
			"source_nat":                    sourceNatSdkToTerraform(ctx, diags, *v.SourceNat),
			"static_nat":                    staticNatSdkToTerraform(ctx, diags, v.StaticNat),
			"summarized_subnet":             types.StringValue(*v.SummarizedSubnet),
			"summarized_subnet_to_lan_bgp":  types.StringValue(*v.SummarizedSubnetToLanBgp),
			"summarized_subnet_to_lan_ospf": types.StringValue(*v.SummarizedSubnetToLanOspf),
		}
		n, e := NewVpnAccessValue(state_value_map_attr_type, state_value_map_attr_value)
		diags.Append(e...)
		state_value_map_value[k] = n
	}
	state_result_map_type := VpnAccessValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}
