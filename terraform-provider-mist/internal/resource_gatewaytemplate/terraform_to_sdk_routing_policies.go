package resource_gatewaytemplate

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func routingPolicyTermActionTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.RoutingPolicyTermAction {
	tflog.Debug(ctx, "routingPolicyTermActionTerraformToSdk")
	data := *mistsdkgo.NewRoutingPolicyTermAction()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewActionValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetAccept(plan.Accept.ValueBool())
		data.SetAddCommunity(mist_transform.ListOfStringTerraformToSdk(ctx, plan.AddCommunity))
		data.SetAddTargetVrfs(mist_transform.ListOfStringTerraformToSdk(ctx, plan.AddTargetVrfs))
		data.SetCommunity(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Community))
		data.SetExcludeAsPath(mist_transform.ListOfStringTerraformToSdk(ctx, plan.ExcludeAsPath))
		data.SetExcludeCommunity(mist_transform.ListOfStringTerraformToSdk(ctx, plan.ExcludeCommunity))
		data.SetLocalPreference(plan.LocalPreference.ValueString())
		data.SetPrependAsPath(mist_transform.ListOfStringTerraformToSdk(ctx, plan.PrependAsPath))
		return data
	}
}

func routingPolicyTermMatchingRouteExistsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.RoutingPolicyTermMatchingRouteExists {
	tflog.Debug(ctx, "routingPolicyTermMatchingRouteExistsTerraformToSdk")
	data := *mistsdkgo.NewRoutingPolicyTermMatchingRouteExists()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewRouteExistsValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetRoute(plan.Route.ValueString())
		data.SetVrfName(plan.VrfName.ValueString())
		return data
	}
}

func routingPolicyTermMatchingVpnPathSlaExistsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.RoutingPolicyTermMatchingVpnPathSla {
	tflog.Debug(ctx, "routingPolicyTermMatchingVpnPathSlaExistsTerraformToSdk")
	data := *mistsdkgo.NewRoutingPolicyTermMatchingVpnPathSla()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewVpnPathSlaValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetMaxJitter(int32(plan.MaxJitter.ValueInt64()))
		data.SetMaxLatency(int32(plan.MaxLatency.ValueInt64()))
		data.SetMaxLoss(int32(plan.MaxLoss.ValueInt64()))
		return data
	}
}

func routingPolicyTermMatchingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.RoutingPolicyTermMatching {
	tflog.Debug(ctx, "routingPolicyTermMatchingTerraformToSdk")
	data := *mistsdkgo.NewRoutingPolicyTermMatching()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewRoutingPolicyTermMatchingValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetAsPath(mist_transform.ListOfStringTerraformToSdk(ctx, plan.AsPath))
		data.SetCommunity(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Community))
		data.SetNetwork(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Network))
		data.SetPrefix(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Prefix))
		data.SetProtocol(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Protocol))

		route_exists := routingPolicyTermMatchingRouteExistsTerraformToSdk(ctx, diags, plan.RouteExists)
		data.SetRouteExists(route_exists)

		data.SetVpnNeighborMac(mist_transform.ListOfStringTerraformToSdk(ctx, plan.VpnNeighborMac))
		data.SetVpnPath(mist_transform.ListOfStringTerraformToSdk(ctx, plan.VpnPath))

		vpn_path_sla := routingPolicyTermMatchingVpnPathSlaExistsTerraformToSdk(ctx, diags, plan.VpnPathSla)
		data.SetVpnPathSla(vpn_path_sla)
		return data
	}
}

func routingPolicyTermerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.RoutingPolicyTerm {
	tflog.Debug(ctx, "routingPolicyTermerraformToSdk")
	var data_list []mistsdkgo.RoutingPolicyTerm
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(TermsValue)
		data := mistsdkgo.NewRoutingPolicyTerm()

		action := routingPolicyTermActionTerraformToSdk(ctx, diags, plan.Action)
		data.SetAction(action)

		routing_policy_term_matching := routingPolicyTermMatchingTerraformToSdk(ctx, diags, plan.RoutingPolicyTermMatching)
		data.SetMatching(routing_policy_term_matching)

		data_list = append(data_list, *data)
	}
	return data_list
}

func routingPoliciesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.RoutingPolicy {
	tflog.Debug(ctx, "routingPoliciesTerraformToSdk")
	data_map := make(map[string]mistsdkgo.RoutingPolicy)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(RoutingPoliciesValue)

		data := *mistsdkgo.NewRoutingPolicy()
		terms := routingPolicyTermerraformToSdk(ctx, diags, plan.Terms)
		data.SetTerms(terms)

		data_map[k] = data
	}
	return data_map
}
