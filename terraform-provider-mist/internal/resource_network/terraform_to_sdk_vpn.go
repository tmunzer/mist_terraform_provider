package resource_network

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/utils/transform"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func VpnTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.NetworkVpnAccessConfig {
	data_map := make(map[string]mistsdkgo.NetworkVpnAccessConfig)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(VpnAccessValue)

		destination_nat := destinationNatTerraformToSdk(ctx, diags, v_plan.DestinationNat)
		source_nat := sourceNatTerraformToSdk(ctx, diags, v_plan.SourceNat)
		static_nat := staticNatTerraformToSdk(ctx, diags, v_plan.StaticNat)

		data := *mistsdkgo.NewNetworkVpnAccessConfig()
		data.SetAdvertisedSubnet(v_plan.AdvertisedSubnet.ValueString())
		data.SetAllowPing(v_plan.AllowPing.ValueBool())
		data.SetDestinationNat(destination_nat)
		data.SetNatPool(v_plan.NatPool.ValueString())
		data.SetNoReadvertiseToLanBgp(v_plan.NoReadvertiseToLanBgp.ValueBool())
		data.SetNoReadvertiseToLanOspf(v_plan.NoReadvertiseToLanOspf.ValueBool())
		data.SetNoReadvertiseToOverlay(v_plan.NoReadvertiseToOverlay.ValueBool())
		data.SetOtherVrfs(mist_transform.ListOfStringTerraformToSdk(ctx, v_plan.OtherVrfs))
		data.SetRouted(v_plan.Routed.ValueBool())
		data.SetSourceNat(source_nat)
		data.SetStaticNat(static_nat)
		data.SetSummarizedSubnet(v_plan.SummarizedSubnet.ValueString())
		data.SetSummarizedSubnetToLanBgp(v_plan.SummarizedSubnetToLanBgp.ValueString())
		data.SetSummarizedSubnetToLanOspf(v_plan.SummarizedSubnetToLanOspf.ValueString())
		data_map[k] = data
	}
	return data_map
}
