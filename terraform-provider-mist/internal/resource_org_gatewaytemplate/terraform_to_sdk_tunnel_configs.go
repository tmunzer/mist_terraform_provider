package resource_org_gatewaytemplate

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func tunnelConfigsAutoProvisionPrimaryTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.TunnelConfigsAutoProvisionNode {
	tflog.Debug(ctx, "tunnelConfigsAutoProvisionPrimaryTerraformToSdk")
	data := *mistsdkgo.NewTunnelConfigsAutoProvisionNode()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewAutoProvisionPrimaryValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetNumHosts(plan.NumHosts.ValueString())
		data.SetWanNames(mist_transform.ListOfStringTerraformToSdk(ctx, plan.WanNames))

		return data
	}
}

func tunnelConfigsAutoProvisionSecondaryTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.TunnelConfigsAutoProvisionNode {
	tflog.Debug(ctx, "tunnelConfigsAutoProvisionSecondaryTerraformToSdk")
	data := *mistsdkgo.NewTunnelConfigsAutoProvisionNode()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewAutoProvisionSecondaryValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetNumHosts(plan.NumHosts.ValueString())
		data.SetWanNames(mist_transform.ListOfStringTerraformToSdk(ctx, plan.WanNames))

		return data
	}
}

func tunnelConfigsAutoProvisionTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.TunnelConfigsAutoProvision {
	tflog.Debug(ctx, "tunnelConfigsAutoProvisionTerraformToSdk")
	data := *mistsdkgo.NewTunnelConfigsAutoProvision()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewAutoProvisionValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetEnable(plan.Enable.ValueBool())

		var plan_latlng_interface interface{} = plan.Latlng
		plan_latlng := plan_latlng_interface.(LatlngValue)

		var latlng mistsdkgo.LatLng
		latlng.SetLat(plan_latlng.Lat.ValueFloat64())
		latlng.SetLng(plan_latlng.Lng.ValueFloat64())
		data.SetLatlng(latlng)

		primary := tunnelConfigsAutoProvisionPrimaryTerraformToSdk(ctx, diags, plan.AutoProvisionPrimary)
		data.SetPrimary(primary)

		data.SetRegion(mistsdkgo.TunnelConfigsAutoProvisionRegion(plan.Region.ValueString()))

		secondary := tunnelConfigsAutoProvisionSecondaryTerraformToSdk(ctx, diags, plan.AutoProvisionSecondary)
		data.SetSecondary(secondary)

		return data
	}
}

func gatewayTemplateTunnelIkeProposalTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.GatewayTemplateTunnelIkeProposal {
	tflog.Debug(ctx, "gatewayTemplateTunnelIkeProposalTerraformToSdk")
	var data_list []mistsdkgo.GatewayTemplateTunnelIkeProposal
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(IkeProposalsValue)
		data := mistsdkgo.NewGatewayTemplateTunnelIkeProposal()
		data.SetAuthAlgo(mistsdkgo.TunnelConfigsAuthAlgo(plan.AuthAlgo.ValueString()))
		data.SetDhGroup(mistsdkgo.GatewayTemplateTunnelIkeDhGroup(plan.DhGroup.ValueString()))
		data.SetEncAlgo(mistsdkgo.TunnelConfigsEncAlgo(plan.EncAlgo.ValueString()))

		data_list = append(data_list, *data)
	}
	return data_list
}

func gatewayTemplateTunnelIpsecProposalTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []mistsdkgo.GatewayTemplateTunnelIpsecProposal {
	tflog.Debug(ctx, "gatewayTemplateTunnelIpsecProposalTerraformToSdk")
	var data_list []mistsdkgo.GatewayTemplateTunnelIpsecProposal
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(IpsecProposalsValue)
		data := mistsdkgo.NewGatewayTemplateTunnelIpsecProposal()
		data.SetAuthAlgo(mistsdkgo.TunnelConfigsAuthAlgo(plan.AuthAlgo.ValueString()))
		data.SetDhGroup(mistsdkgo.TunnelConfigsDhGroup(plan.DhGroup.ValueString()))
		data.SetEncAlgo(mistsdkgo.TunnelConfigsEncAlgo(plan.EncAlgo.ValueString()))

		data_list = append(data_list, *data)
	}
	return data_list
}

func gatewayTemplateTunnelProbeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.GatewayTemplateTunnelProbe {
	tflog.Debug(ctx, "gatewayTemplateTunnelProbeTerraformToSdk")
	data := *mistsdkgo.NewGatewayTemplateTunnelProbe()
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewProbeValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.SetInterval(int32(plan.Interval.ValueInt64()))
		data.SetThreshold(int32(plan.Threshold.ValueInt64()))
		data.SetTimeout(int32(plan.Timeout.ValueInt64()))
		data.SetType(mistsdkgo.GatewayTemplateProbeType(plan.ProbeType.ValueString()))
		return data
	}
}

func gatewayTemplateTunnelPrimaryProbeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.GatewayTemplateTunnelNode {
	tflog.Debug(ctx, "gatewayTemplateTunnelPrimaryProbeTerraformToSdk")
	data := *mistsdkgo.NewGatewayTemplateTunnelNode()
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

func gatewayTemplateTunnelSecondaryProbeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) mistsdkgo.GatewayTemplateTunnelNode {
	tflog.Debug(ctx, "gatewayTemplateTunnelSecondaryProbeTerraformToSdk")
	data := *mistsdkgo.NewGatewayTemplateTunnelNode()
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

func tunnelConfigsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]mistsdkgo.TunnelConfigs {
	tflog.Debug(ctx, "tunnelConfigsTerraformToSdk")
	data_map := make(map[string]mistsdkgo.TunnelConfigs)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(TunnelConfigsValue)

		data := *mistsdkgo.NewTunnelConfigs()

		auto_provision := tunnelConfigsAutoProvisionTerraformToSdk(ctx, diags, plan.AutoProvision)
		data.SetAutoProvision(auto_provision)

		data.SetIkeLifetime(int32(plan.IkeLifetime.ValueInt64()))
		data.SetIkeMode(mistsdkgo.GatewayTemplateTunnelIkeMode(plan.IkeMode.ValueString()))

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

		data.SetProtocol(mistsdkgo.GatewayTemplateTunnelProtocol(plan.Protocol.ValueString()))
		data.SetProvider(mistsdkgo.TunnelProviderOptionsName(plan.Provider.ValueString()))
		data.SetPsk(plan.Psk.ValueString())

		secondary := gatewayTemplateTunnelSecondaryProbeTerraformToSdk(ctx, diags, plan.Secondary)
		data.SetSecondary(secondary)

		data.SetVersion(mistsdkgo.GatewayTemplateTunnelVersion(plan.Version.ValueString()))

		data_map[k] = data
	}
	return data_map
}
