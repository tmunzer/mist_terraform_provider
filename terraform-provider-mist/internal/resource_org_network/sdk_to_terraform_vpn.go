package resource_org_network

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func VpnSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.NetworkVpnAccessConfig) basetypes.MapValue {

	state_value_map_attr_type := VpnAccessValue{}.AttributeTypes(ctx)
	state_value_map_value := make(map[string]attr.Value)
	for k, v := range d {
		state_value_map_attr_value := map[string]attr.Value{
			"advertised_subnet":             types.StringValue(v.GetAdvertisedSubnet()),
			"allow_ping":                    types.BoolValue(v.GetAllowPing()),
			"destination_nat":               destinationNatSdkToTerraform(ctx, diags, v.GetDestinationNat()),
			"nat_pool":                      types.StringValue(v.GetNatPool()),
			"no_readvertise_to_lan_bgp":     types.BoolValue(v.GetNoReadvertiseToLanBgp()),
			"no_readvertise_to_lan_ospf":    types.BoolValue(v.GetNoReadvertiseToLanOspf()),
			"no_readvertise_to_overlay":     types.BoolValue(v.GetNoReadvertiseToOverlay()),
			"other_vrfs":                    mist_transform.ListOfStringSdkToTerraform(ctx, v.GetOtherVrfs()),
			"routed":                        types.BoolValue(v.GetRouted()),
			"source_nat":                    sourceNatSdkToTerraform(ctx, diags, v.GetSourceNat()),
			"static_nat":                    staticNatSdkToTerraform(ctx, diags, v.GetStaticNat()),
			"summarized_subnet":             types.StringValue(v.GetSummarizedSubnet()),
			"summarized_subnet_to_lan_bgp":  types.StringValue(v.GetSummarizedSubnetToLanBgp()),
			"summarized_subnet_to_lan_ospf": types.StringValue(v.GetSummarizedSubnetToLanOspf()),
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
