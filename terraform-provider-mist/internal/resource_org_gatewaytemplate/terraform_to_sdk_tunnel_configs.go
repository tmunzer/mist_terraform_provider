package resource_org_gatewaytemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func tunnelConfigsAutoProvisionPrimaryTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.TunnelConfigsAutoProvisionNode {
	tflog.Debug(ctx, "tunnelConfigsAutoProvisionPrimaryTerraformToSdk")
	data := models.TunnelConfigsAutoProvisionNode{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewAutoProvisionPrimaryValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.NumHosts = models.ToPointer(plan.NumHosts.ValueString())
		data.WanNames = mist_transform.ListOfStringTerraformToSdk(ctx, plan.WanNames)

		return data
	}
}

func tunnelConfigsAutoProvisionSecondaryTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.TunnelConfigsAutoProvisionNode {
	tflog.Debug(ctx, "tunnelConfigsAutoProvisionSecondaryTerraformToSdk")
	data := models.TunnelConfigsAutoProvisionNode{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewAutoProvisionSecondaryValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.NumHosts = models.ToPointer(plan.NumHosts.ValueString())
		data.WanNames = mist_transform.ListOfStringTerraformToSdk(ctx, plan.WanNames)

		return data
	}
}

func tunnelConfigsAutoProvisionTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.TunnelConfigsAutoProvision {
	tflog.Debug(ctx, "tunnelConfigsAutoProvisionTerraformToSdk")
	data := models.TunnelConfigsAutoProvision{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewAutoProvisionValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.Enable = models.ToPointer(plan.Enable.ValueBool())

		var plan_latlng_interface interface{} = plan.Latlng
		plan_latlng := plan_latlng_interface.(LatlngValue)

		var latlng models.LatLng
		latlng.Lat = plan_latlng.Lng.ValueFloat64()
		latlng.Lng = plan_latlng.Lng.ValueFloat64()
		data.Latlng = models.ToPointer(latlng)

		primary := tunnelConfigsAutoProvisionPrimaryTerraformToSdk(ctx, diags, plan.AutoProvisionPrimary)
		data.Primary = &primary

		data.Region = models.ToPointer(models.TunnelConfigsAutoProvisionRegionEnum(plan.Region.ValueString()))

		secondary := tunnelConfigsAutoProvisionSecondaryTerraformToSdk(ctx, diags, plan.AutoProvisionSecondary)
		data.Secondary = &secondary

		return data
	}
}

func gatewayTemplateTunnelIkeProposalTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.GatewayTemplateTunnelIkeProposal {
	tflog.Debug(ctx, "gatewayTemplateTunnelIkeProposalTerraformToSdk")
	var data_list []models.GatewayTemplateTunnelIkeProposal
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(IkeProposalsValue)
		data := models.GatewayTemplateTunnelIkeProposal{}
		data.AuthAlgo = models.ToPointer(models.TunnelConfigsAuthAlgoEnum(plan.AuthAlgo.ValueString()))
		data.DhGroup = models.ToPointer(models.GatewayTemplateTunnelIkeDhGroupEnum(plan.DhGroup.ValueString()))
		data.EncAlgo = models.NewOptional(models.ToPointer(models.TunnelConfigsEncAlgoEnum(plan.EncAlgo.ValueString())))

		data_list = append(data_list, data)
	}
	return data_list
}

func gatewayTemplateTunnelIpsecProposalTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.GatewayTemplateTunnelIpsecProposal {
	tflog.Debug(ctx, "gatewayTemplateTunnelIpsecProposalTerraformToSdk")
	var data_list []models.GatewayTemplateTunnelIpsecProposal
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(IpsecProposalsValue)
		data := models.GatewayTemplateTunnelIpsecProposal{}
		data.AuthAlgo = models.ToPointer(models.TunnelConfigsAuthAlgoEnum(plan.AuthAlgo.ValueString()))
		data.DhGroup = models.ToPointer(models.TunnelConfigsDhGroupEnum(plan.DhGroup.ValueString()))
		data.EncAlgo = models.NewOptional(models.ToPointer(models.TunnelConfigsEncAlgoEnum(plan.EncAlgo.ValueString())))

		data_list = append(data_list, data)
	}
	return data_list
}

func gatewayTemplateTunnelProbeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.GatewayTemplateTunnelProbe {
	tflog.Debug(ctx, "gatewayTemplateTunnelProbeTerraformToSdk")
	data := models.GatewayTemplateTunnelProbe{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewProbeValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.Interval = models.ToPointer(int(plan.Interval.ValueInt64()))
		data.Threshold = models.ToPointer(int(plan.Threshold.ValueInt64()))
		data.Timeout = models.ToPointer(int(plan.Timeout.ValueInt64()))
		data.Type = models.ToPointer(models.GatewayTemplateProbeTypeEnum(plan.ProbeType.ValueString()))
		return data
	}
}

func gatewayTemplateTunnelPrimaryProbeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.GatewayTemplateTunnelNode {
	tflog.Debug(ctx, "gatewayTemplateTunnelPrimaryProbeTerraformToSdk")
	data := models.GatewayTemplateTunnelNode{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewPrimaryValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.Hosts = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Hosts)
		data.InternalIps = mist_transform.ListOfStringTerraformToSdk(ctx, plan.InternalIps)
		data.ProbeIps = mist_transform.ListOfStringTerraformToSdk(ctx, plan.ProbeIps)
		data.RemoteIds = mist_transform.ListOfStringTerraformToSdk(ctx, plan.RemoteIds)
		data.WanNames = mist_transform.ListOfStringTerraformToSdk(ctx, plan.WanNames)
		return data
	}
}

func gatewayTemplateTunnelSecondaryProbeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.GatewayTemplateTunnelNode {
	tflog.Debug(ctx, "gatewayTemplateTunnelSecondaryProbeTerraformToSdk")
	data := models.GatewayTemplateTunnelNode{}
	if d.IsNull() || d.IsUnknown() {
		return data
	} else {
		plan := NewSecondaryValueMust(d.AttributeTypes(ctx), d.Attributes())
		data.Hosts = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Hosts)
		data.InternalIps = mist_transform.ListOfStringTerraformToSdk(ctx, plan.InternalIps)
		data.ProbeIps = mist_transform.ListOfStringTerraformToSdk(ctx, plan.ProbeIps)
		data.RemoteIds = mist_transform.ListOfStringTerraformToSdk(ctx, plan.RemoteIds)
		data.WanNames = mist_transform.ListOfStringTerraformToSdk(ctx, plan.WanNames)
		return data
	}
}

func tunnelConfigsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.TunnelConfigs {
	tflog.Debug(ctx, "tunnelConfigsTerraformToSdk")
	data_map := make(map[string]models.TunnelConfigs)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(TunnelConfigsValue)

		data := models.TunnelConfigs{}

		auto_provision := tunnelConfigsAutoProvisionTerraformToSdk(ctx, diags, plan.AutoProvision)
		data.AutoProvision = &auto_provision

		data.IkeLifetime = models.ToPointer(int(plan.IkeLifetime.ValueInt64()))
		data.IkeMode = models.ToPointer(models.GatewayTemplateTunnelIkeModeEnum(plan.IkeMode.ValueString()))

		ike_proposals := gatewayTemplateTunnelIkeProposalTerraformToSdk(ctx, diags, plan.IkeProposals)
		data.IkeProposals = ike_proposals

		data.IpsecLifetime = models.ToPointer(int(plan.IpsecLifetime.ValueInt64()))

		primary := gatewayTemplateTunnelPrimaryProbeTerraformToSdk(ctx, diags, plan.Primary)
		data.Primary = &primary

		ipsec_proposals := gatewayTemplateTunnelIpsecProposalTerraformToSdk(ctx, diags, plan.IpsecProposals)
		data.IpsecProposals = ipsec_proposals

		data.LocalId = models.ToPointer(plan.LocalId.ValueString())

		probe := gatewayTemplateTunnelProbeTerraformToSdk(ctx, diags, plan.Probe)
		data.Probe = &probe

		data.Protocol = models.ToPointer(models.GatewayTemplateTunnelProtocolEnum(plan.Protocol.ValueString()))
		data.Provider = models.ToPointer(models.TunnelProviderOptionsNameEnum(plan.Provider.ValueString()))
		data.Psk = models.ToPointer(plan.Psk.ValueString())

		secondary := gatewayTemplateTunnelSecondaryProbeTerraformToSdk(ctx, diags, plan.Secondary)
		data.Secondary = &secondary

		data.Version = models.ToPointer(models.GatewayTemplateTunnelVersionEnum(plan.Version.ValueString()))

		data_map[k] = data
	}
	return data_map
}
