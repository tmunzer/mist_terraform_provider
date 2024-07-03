package resource_org_gatewaytemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func bgpConfigCommunitiesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.BgpConfigCommunity {
	tflog.Debug(ctx, "bgpConfigCommunitiesTerraformToSdk")

	var data_list []mistapigo.BgpConfigCommunity
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(CommunitiesValue)
		data := mistapigo.NewBgpConfigCommunity()
		data.SetId(plan.Id.ValueString())
		data.SetLocalPreference(int32(plan.LocalPreference.ValueInt64()))
		data.SetVpnName(plan.VpnName.ValueString())

		data_list = append(data_list, *data)
	}
	return data_list
}

func bgpConfigNeighborsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.BgpConfigNeighbors {
	tflog.Debug(ctx, "bgpConfigNeighborsTerraformToSdk")
	data_map := make(map[string]mistapigo.BgpConfigNeighbors)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(NeighborsValue)

		data := *mistapigo.NewBgpConfigNeighbors()
		data.SetDisabled(plan.Disabled.ValueBool())
		data.SetExportPolicy(plan.ExportPolicy.ValueString())
		data.SetHoldTime(int32(plan.HoldTime.ValueInt64()))
		data.SetImportPolicy(plan.ImportPolicy.ValueString())
		data.SetMultihopTtl(int32(plan.MultihopTtl.ValueInt64()))
		data.SetNeighborAs(int32(plan.NeighborAs.ValueInt64()))

		data_map[k] = data
	}
	return data_map
}

func bgpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.BgpConfig {
	tflog.Debug(ctx, "bgpConfigTerraformToSdk")
	data_map := make(map[string]mistapigo.BgpConfig)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(BgpConfigValue)

		communities := bgpConfigCommunitiesTerraformToSdk(ctx, diags, plan.Communities)
		neighbors := bgpConfigNeighborsTerraformToSdk(ctx, diags, plan.Neighbors)

		data := *mistapigo.NewBgpConfig()
		data.SetAuthKey(plan.AuthKey.ValueString())
		data.SetBfdMinimumInterval(int32(plan.BfdMinimumInterval.ValueInt64()))
		data.SetBfdMultiplier(int32(plan.BfdMultiplier.ValueInt64()))
		data.SetCommunities(communities)
		data.SetDisableBfd(plan.DisableBfd.ValueBool())
		data.SetExport(plan.Export.ValueString())
		data.SetExportPolicy(plan.ExportPolicy.ValueString())
		data.SetExtendedV4Nexthop(plan.ExtendedV4Nexthop.ValueBool())
		data.SetGracefulRestartTime(int32(plan.GracefulRestartTime.ValueInt64()))
		data.SetHoldTime(int32(plan.HoldTime.ValueInt64()))
		data.SetImport(plan.Import.ValueString())
		data.SetImportPolicy(plan.ImportPolicy.ValueString())
		data.SetNeighborAs(int32(plan.NeighborAs.ValueInt64()))
		data.SetNeighbors(neighbors)
		data.SetNetworks(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Networks))
		data.SetNoReadvertiseToOverlay(plan.NoReadvertiseToOverlay.ValueBool())
		data.SetType(mistapigo.BgpConfigType(plan.BgpConfigType.ValueString()))
		data.SetVia(mistapigo.BgpConfigVia(plan.Via.ValueString()))
		data.SetVpnName(plan.VpnName.ValueString())
		data.SetWanName(plan.WanName.ValueString())

		data_map[k] = data
	}
	return data_map
}
