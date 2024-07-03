package resource_org_gatewaytemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func tunnelConfigsAutoProvisionPrimaryTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistapigo.TunnelConfigsAutoProvisionNode {
	tflog.Debug(ctx, "tunnelConfigsAutoProvisionPrimaryTerraformToSdk")
	data := *mistapigo.NewTunnelConfigsAutoProvisionNode()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewAutoProvisionPrimaryValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetNumHosts(plan.NumHosts.ValueString())
		data.SetWanNames(mist_transform.ListOfStringTerraformToSdk(ctx, plan.WanNames))

		return data
	}
}

func tunnelConfigsAutoProvisionSecondaryTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistapigo.TunnelConfigsAutoProvisionNode {
	tflog.Debug(ctx, "tunnelConfigsAutoProvisionSecondaryTerraformToSdk")
	data := *mistapigo.NewTunnelConfigsAutoProvisionNode()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewAutoProvisionSecondaryValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetNumHosts(plan.NumHosts.ValueString())
		data.SetWanNames(mist_transform.ListOfStringTerraformToSdk(ctx, plan.WanNames))

		return data
	}
}

func tunnelConfigsAutoProvisionTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistapigo.TunnelConfigsAutoProvision {
	tflog.Debug(ctx, "tunnelConfigsAutoProvisionTerraformToSdk")
	data := *mistapigo.NewTunnelConfigsAutoProvision()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewAutoProvisionValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetEnable(plan.Enable.ValueBool())

		var plan_latlng_interface interface{} = plan.Latlng
		plan_latlng := plan_latlng_interface.(LatlngValue)

		var latlng mistapigo.LatLng
		latlng.SetLat(plan_latlng.Lat.ValueFloat64())
		latlng.SetLng(plan_latlng.Lng.ValueFloat64())
		data.SetLatlng(latlng)

		primary := tunnelConfigsAutoProvisionPrimaryTerraformToSdk(ctx, diags, plan.AutoProvisionPrimary)
		data.SetPrimary(primary)

		data.SetRegion(mistapigo.TunnelConfigsAutoProvisionRegion(plan.Region.ValueString()))

		secondary := tunnelConfigsAutoProvisionSecondaryTerraformToSdk(ctx, diags, plan.AutoProvisionSecondary)
		data.SetSecondary(secondary)

		return data
	}
}

func gatewayTemplateTunnelIkeProposalTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.GatewayTemplateTunnelIkeProposal {
	tflog.Debug(ctx, "gatewayTemplateTunnelIkeProposalTerraformToSdk")
	var data_list []mistapigo.GatewayTemplateTunnelIkeProposal
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(IkeProposalsValue)
		data := mistapigo.NewGatewayTemplateTunnelIkeProposal()
		data.SetAuthAlgo(mistapigo.TunnelConfigsAuthAlgo(plan.AuthAlgo.ValueString()))
		data.SetDhGroup(mistapigo.GatewayTemplateTunnelIkeDhGroup(plan.DhGroup.ValueString()))
		data.SetEncAlgo(mistapigo.TunnelConfigsEncAlgo(plan.EncAlgo.ValueString()))

		data_list = append(data_list, *data)
	}
	return data_list
}

func gatewayTemplateTunnelIpsecProposalTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistapigo.GatewayTemplateTunnelIpsecProposal {
	tflog.Debug(ctx, "gatewayTemplateTunnelIpsecProposalTerraformToSdk")
	var data_list []mistapigo.GatewayTemplateTunnelIpsecProposal
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(IpsecProposalsValue)
		data := mistapigo.NewGatewayTemplateTunnelIpsecProposal()
		data.SetAuthAlgo(mistapigo.TunnelConfigsAuthAlgo(plan.AuthAlgo.ValueString()))
		data.SetDhGroup(mistapigo.TunnelConfigsDhGroup(plan.DhGroup.ValueString()))
		data.SetEncAlgo(mistapigo.TunnelConfigsEncAlgo(plan.EncAlgo.ValueString()))

		data_list = append(data_list, *data)
	}
	return data_list
}

func gatewayTemplateTunnelProbeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistapigo.GatewayTemplateTunnelProbe {
	tflog.Debug(ctx, "gatewayTemplateTunnelProbeTerraformToSdk")
	data := *mistapigo.NewGatewayTemplateTunnelProbe()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewProbeValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetInterval(int32(plan.Interval.ValueInt64()))
		data.SetThreshold(int32(plan.Threshold.ValueInt64()))
		data.SetTimeout(int32(plan.Timeout.ValueInt64()))
		data.SetType(mistapigo.GatewayTemplateProbeType(plan.ProbeType.ValueString()))
		return data
	}
}

func gatewayTemplateTunnelPrimaryProbeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistapigo.GatewayTemplateTunnelNode {
	tflog.Debug(ctx, "gatewayTemplateTunnelPrimaryProbeTerraformToSdk")
	data := *mistapigo.NewGatewayTemplateTunnelNode()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewPrimaryValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetHosts(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Hosts))
		data.SetInternalIps(mist_transform.ListOfStringTerraformToSdk(ctx, plan.InternalIps))
		data.SetProbeIps(mist_transform.ListOfStringTerraformToSdk(ctx, plan.ProbeIps))
		data.SetRemoteIds(mist_transform.ListOfStringTerraformToSdk(ctx, plan.RemoteIds))
		data.SetWanNames(mist_transform.ListOfStringTerraformToSdk(ctx, plan.WanNames))
		return data
	}
}

func gatewayTemplateTunnelSecondaryProbeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistapigo.GatewayTemplateTunnelNode {
	tflog.Debug(ctx, "gatewayTemplateTunnelSecondaryProbeTerraformToSdk")
	data := *mistapigo.NewGatewayTemplateTunnelNode()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewSecondaryValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetHosts(mist_transform.ListOfStringTerraformToSdk(ctx, plan.Hosts))
		data.SetInternalIps(mist_transform.ListOfStringTerraformToSdk(ctx, plan.InternalIps))
		data.SetProbeIps(mist_transform.ListOfStringTerraformToSdk(ctx, plan.ProbeIps))
		data.SetRemoteIds(mist_transform.ListOfStringTerraformToSdk(ctx, plan.RemoteIds))
		data.SetWanNames(mist_transform.ListOfStringTerraformToSdk(ctx, plan.WanNames))
		return data
	}
}

func tunnelConfigsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistapigo.TunnelConfigs {
	tflog.Debug(ctx, "tunnelConfigsTerraformToSdk")
	data_map := make(map[string]mistapigo.TunnelConfigs)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(TunnelConfigsValue)

		data := *mistapigo.NewTunnelConfigs()

		auto_provision := tunnelConfigsAutoProvisionTerraformToSdk(ctx, diags, plan.AutoProvision)
		data.SetAutoProvision(auto_provision)

		data.SetIkeLifetime(int32(plan.IkeLifetime.ValueInt64()))
		data.SetIkeMode(mistapigo.GatewayTemplateTunnelIkeMode(plan.IkeMode.ValueString()))

		ike_proposals := gatewayTemplateTunnelIkeProposalTerraformToSdk(ctx, diags, plan.IkeProposals)
		data.SetIkeProposals(ike_proposals)

		data.SetIpsecLifetime(int32(plan.IpsecLifetime.ValueInt64()))

		primary := gatewayTemplateTunnelPrimaryProbeTerraformToSdk(ctx, diags, plan.Primary)
		data.SetPrimary(primary)

		ipsec_proposals := gatewayTemplateTunnelIpsecProposalTerraformToSdk(ctx, diags, plan.IpsecProposals)
		data.SetIpsecProposals(ipsec_proposals)

		data.SetLocalId(plan.LocalId.ValueString())

		probe := gatewayTemplateTunnelProbeTerraformToSdk(ctx, diags, plan.Probe)
		data.SetProbe(probe)

		data.SetProtocol(mistapigo.GatewayTemplateTunnelProtocol(plan.Protocol.ValueString()))
		data.SetProvider(mistapigo.TunnelProviderOptionsName(plan.Provider.ValueString()))
		data.SetPsk(plan.Psk.ValueString())

		secondary := gatewayTemplateTunnelSecondaryProbeTerraformToSdk(ctx, diags, plan.Secondary)
		data.SetSecondary(secondary)

		data.SetVersion(mistapigo.GatewayTemplateTunnelVersion(plan.Version.ValueString()))

		data_map[k] = data
	}
	return data_map
}
