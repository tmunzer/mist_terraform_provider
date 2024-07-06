package resource_org_gatewaytemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func routingPolicyTermActionTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.RoutingPolicyTermAction {
	tflog.Debug(ctx, "routingPolicyTermActionTerraformToSdk")
	data := models.RoutingPolicyTermAction{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan := NewActionValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.Accept = models.ToPointer(plan.Accept.ValueBool())
		data.AddCommunity = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AddCommunity)
		data.AddTargetVrfs = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AddTargetVrfs)
		data.Community = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Community)
		data.ExcludeAsPath = mist_transform.ListOfStringTerraformToSdk(ctx, plan.ExcludeAsPath)
		data.ExcludeCommunity = mist_transform.ListOfStringTerraformToSdk(ctx, plan.ExcludeCommunity)
		data.LocalPreference = models.ToPointer(plan.LocalPreference.ValueString())
		data.PrependAsPath = mist_transform.ListOfStringTerraformToSdk(ctx, plan.PrependAsPath)
		return &data
	}
}

func routingPolicyTermMatchingRouteExistsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.RoutingPolicyTermMatchingRouteExists {
	tflog.Debug(ctx, "routingPolicyTermMatchingRouteExistsTerraformToSdk")
	data := models.RoutingPolicyTermMatchingRouteExists{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan := NewRouteExistsValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.Route = models.ToPointer(plan.Route.ValueString())
		data.VrfName = models.ToPointer(plan.VrfName.ValueString())
		return &data
	}
}

func routingPolicyTermMatchingVpnPathSlaExistsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.RoutingPolicyTermMatchingVpnPathSla {
	tflog.Debug(ctx, "routingPolicyTermMatchingVpnPathSlaExistsTerraformToSdk")
	data := models.RoutingPolicyTermMatchingVpnPathSla{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan := NewVpnPathSlaValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.MaxJitter.SetValue(models.ToPointer(int(plan.MaxJitter.ValueInt64())))
		data.MaxLatency.SetValue(models.ToPointer(int(plan.MaxLatency.ValueInt64())))
		data.MaxLoss.SetValue(models.ToPointer(int(plan.MaxLoss.ValueInt64())))
		return &data
	}
}

func routingPolicyTermMatchingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.RoutingPolicyTermMatching {
	tflog.Debug(ctx, "routingPolicyTermMatchingTerraformToSdk")
	data := models.RoutingPolicyTermMatching{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan := NewRoutingPolicyTermMatchingValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.AsPath = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AsPath)
		data.Community = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Community)
		data.Network = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Network)
		data.Prefix = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Prefix)
		data.Protocol = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Protocol)

		route_exists := routingPolicyTermMatchingRouteExistsTerraformToSdk(ctx, diags, plan.RouteExists)
		data.RouteExists = route_exists

		data.VpnNeighborMac = mist_transform.ListOfStringTerraformToSdk(ctx, plan.VpnNeighborMac)
		data.VpnPath = mist_transform.ListOfStringTerraformToSdk(ctx, plan.VpnPath)

		vpn_path_sla := routingPolicyTermMatchingVpnPathSlaExistsTerraformToSdk(ctx, diags, plan.VpnPathSla)
		data.VpnPathSla = vpn_path_sla
		return &data
	}
}

func routingPolicyTermerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.RoutingPolicyTerm {
	tflog.Debug(ctx, "routingPolicyTermerraformToSdk")
	var data_list []models.RoutingPolicyTerm
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(TermsValue)
		data := models.RoutingPolicyTerm{}

		action := routingPolicyTermActionTerraformToSdk(ctx, diags, plan.Action)
		data.Action = action

		routing_policy_term_matching := routingPolicyTermMatchingTerraformToSdk(ctx, diags, plan.RoutingPolicyTermMatching)
		data.Matching = routing_policy_term_matching

		data_list = append(data_list, data)
	}
	return data_list
}

func routingPoliciesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.RoutingPolicy {
	tflog.Debug(ctx, "routingPoliciesTerraformToSdk")
	data_map := make(map[string]models.RoutingPolicy)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(RoutingPoliciesValue)

		data := models.RoutingPolicy{}
		terms := routingPolicyTermerraformToSdk(ctx, diags, plan.Terms)
		data.Terms = terms

		data_map[k] = data
	}
	return data_map
}
