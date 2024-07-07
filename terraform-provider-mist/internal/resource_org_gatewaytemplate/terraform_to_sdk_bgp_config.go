package resource_org_gatewaytemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func bgpConfigCommunitiesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.BgpConfigCommunity {
	tflog.Debug(ctx, "bgpConfigCommunitiesTerraformToSdk")

	var data_list []models.BgpConfigCommunity
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(CommunitiesValue)
		data := models.BgpConfigCommunity{}
		data.Id = plan.Id.ValueStringPointer()
		data.LocalPreference = models.ToPointer(int(plan.LocalPreference.ValueInt64()))
		data.VpnName = plan.VpnName.ValueStringPointer()

		data_list = append(data_list, data)
	}
	return data_list
}

func bgpConfigNeighborsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.BgpConfigNeighbors {
	tflog.Debug(ctx, "bgpConfigNeighborsTerraformToSdk")
	data_map := make(map[string]models.BgpConfigNeighbors)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(NeighborsValue)

		data := models.BgpConfigNeighbors{}
		data.Disabled = models.ToPointer(plan.Disabled.ValueBool())
		data.ExportPolicy = models.ToPointer(plan.ExportPolicy.ValueString())
		data.HoldTime = models.ToPointer(int(plan.HoldTime.ValueInt64()))
		data.ImportPolicy = models.ToPointer(plan.ImportPolicy.ValueString())
		data.MultihopTtl = models.ToPointer(int(plan.MultihopTtl.ValueInt64()))
		data.NeighborAs = models.ToPointer(int(plan.NeighborAs.ValueInt64()))

		data_map[k] = data
	}
	return data_map
}

func bgpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.BgpConfig {
	tflog.Debug(ctx, "bgpConfigTerraformToSdk")
	data_map := make(map[string]models.BgpConfig)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(BgpConfigValue)

		communities := bgpConfigCommunitiesTerraformToSdk(ctx, diags, plan.Communities)
		neighbors := bgpConfigNeighborsTerraformToSdk(ctx, diags, plan.Neighbors)

		data := models.BgpConfig{}
		data.AuthKey = models.ToPointer(plan.AuthKey.ValueString())
		data.BfdMinimumInterval = models.NewOptional(models.ToPointer(int(plan.BfdMinimumInterval.ValueInt64())))
		data.BfdMultiplier = models.NewOptional(models.ToPointer(int(plan.BfdMultiplier.ValueInt64())))
		data.Communities = communities
		data.DisableBfd = models.ToPointer(plan.DisableBfd.ValueBool())
		data.Export = models.ToPointer(plan.Export.ValueString())
		data.ExportPolicy = models.ToPointer(plan.ExportPolicy.ValueString())
		data.ExtendedV4Nexthop = models.ToPointer(plan.ExtendedV4Nexthop.ValueBool())
		data.GracefulRestartTime = models.ToPointer(int(plan.GracefulRestartTime.ValueInt64()))
		data.HoldTime = models.ToPointer(int(plan.HoldTime.ValueInt64()))
		data.Import = models.ToPointer(plan.Import.ValueString())
		data.ImportPolicy = models.ToPointer(plan.ImportPolicy.ValueString())
		data.NeighborAs = models.ToPointer(int(plan.NeighborAs.ValueInt64()))
		data.Neighbors = neighbors
		data.Networks = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Networks)
		data.NoReadvertiseToOverlay = models.ToPointer(plan.NoReadvertiseToOverlay.ValueBool())
		data.Type = models.ToPointer(models.BgpConfigTypeEnum(plan.BgpConfigType.ValueString()))
		data.Via = models.ToPointer(models.BgpConfigViaEnum(plan.Via.ValueString()))
		data.VpnName = models.ToPointer(plan.VpnName.ValueString())
		data.WanName = models.ToPointer(plan.WanName.ValueString())

		data_map[k] = data
	}
	return data_map
}
