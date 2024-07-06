package resource_org_gatewaytemplate

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func tunnelConfigAutoProvNodeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.TunnelConfigsAutoProvisionNode, r_type map[string]attr.Type) basetypes.ObjectValue {
	tflog.Debug(ctx, "tunnelConfigAutoProvNodeSdkToTerraform")

	var num_hosts basetypes.StringValue
	var wan_names basetypes.ListValue

	if d.NumHosts != nil {
		num_hosts = types.StringValue(*d.NumHosts)
	}
	if d.WanNames != nil {
		wan_names = mist_transform.ListOfStringSdkToTerraform(ctx, d.WanNames)
	}

	r_attr_value := map[string]attr.Value{
		"num_hosts": num_hosts,
		"wan_names": wan_names,
	}
	r, e := basetypes.NewObjectValue(r_type, r_attr_value)
	diags.Append(e...)
	return r
}

func tunnelConfigAutoProvSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.TunnelConfigsAutoProvision) basetypes.ObjectValue {
	tflog.Debug(ctx, "tunnelConfigAutoProvSdkToTerraform")
	r_attr_type := AutoProvisionValue{}.AttributeTypes(ctx)

	latlng_type := map[string]attr.Type{
		"lat": basetypes.Float64Type{},
		"lng": basetypes.Float64Type{},
	}
	latlng_value := map[string]attr.Value{
		"lat": types.Float64Value(d.Latlng.Lat),
		"lng": types.Float64Value(d.Latlng.Lng),
	}
	latlng, e := NewLatlngValue(latlng_type, latlng_value)
	diags.Append(e...)

	primary := tunnelConfigAutoProvNodeSdkToTerraform(ctx, diags, *d.Primary, PrimaryValue{}.AttributeTypes(ctx))
	secondary := tunnelConfigAutoProvNodeSdkToTerraform(ctx, diags, *d.Secondary, SecondaryValue{}.AttributeTypes(ctx))

	r_attr_value := map[string]attr.Value{
		"enable":    types.BoolValue(*d.Enable),
		"latlng":    latlng,
		"primary":   primary,
		"region":    types.StringValue(string(*d.Region)),
		"secondary": secondary,
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

func tunnelConfigIkeProposalSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.GatewayTemplateTunnelIkeProposal) basetypes.ListValue {
	tflog.Debug(ctx, "tunnelConfigIkeProposalSdkToTerraform")
	var data_list = []IkeProposalsValue{}
	for _, v := range d {
		data_map_attr_type := IkeProposalsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"auth_algo": types.StringValue(string(*v.AuthAlgo)),
			"dh_group":  types.StringValue(string(*v.DhGroup)),
			"enc_algo":  types.StringValue(string(*v.EncAlgo.Value())),
		}
		data, e := NewIkeProposalsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := IkeProposalsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func tunnelConfigIpsecProposalSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.GatewayTemplateTunnelIpsecProposal) basetypes.ListValue {
	tflog.Debug(ctx, "tunnelConfigIpsecProposalSdkToTerraform")
	var data_list = []IpsecProposalsValue{}
	for _, v := range d {
		data_map_attr_type := IpsecProposalsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"auth_algo": types.StringValue(string(*v.AuthAlgo)),
			"dh_group":  types.StringValue(string(*v.DhGroup)),
			"enc_algo":  types.StringValue(string(*v.EncAlgo.Value())),
		}
		data, e := NewIpsecProposalsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := IpsecProposalsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func tunnelConfigNodeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.GatewayTemplateTunnelNode, r_type map[string]attr.Type) basetypes.ObjectValue {
	tflog.Debug(ctx, "tunnelConfigNodeSdkToTerraform")
	r_attr_value := map[string]attr.Value{
		"hosts":        mist_transform.ListOfStringSdkToTerraform(ctx, d.Hosts),
		"internal_ips": mist_transform.ListOfStringSdkToTerraform(ctx, d.InternalIps),
		"probe_ips":    mist_transform.ListOfStringSdkToTerraform(ctx, d.ProbeIps),
		"remote_ids":   mist_transform.ListOfStringSdkToTerraform(ctx, d.RemoteIds),
		"wan_names":    mist_transform.ListOfStringSdkToTerraform(ctx, d.WanNames),
	}
	r, e := basetypes.NewObjectValue(r_type, r_attr_value)
	diags.Append(e...)
	return r
}

func tunnelConfigProbeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.GatewayTemplateTunnelProbe) basetypes.ObjectValue {
	tflog.Debug(ctx, "tunnelConfigProbeSdkToTerraform")
	r_attr_type := ProbeValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"interval":  types.Int64Value(int64(*d.Interval)),
		"threshold": types.Int64Value(int64(*d.Threshold)),
		"timeout":   types.Int64Value(int64(*d.Timeout)),
		"type":      types.StringValue(string(*d.Type)),
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

func tunnelConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.TunnelConfigs) basetypes.MapValue {
	tflog.Debug(ctx, "tunnelConfigsSdkToTerraform")
	r_type := TunnelConfigsValue{}.AttributeTypes(ctx)
	state_value_map := make(map[string]attr.Value)
	for k, v := range d {
		auto_provision := tunnelConfigAutoProvSdkToTerraform(ctx, diags, *v.AutoProvision)
		ike_proposals := tunnelConfigIkeProposalSdkToTerraform(ctx, diags, v.IkeProposals)
		ipsec_proposals := tunnelConfigIpsecProposalSdkToTerraform(ctx, diags, v.IpsecProposals)
		primary := tunnelConfigNodeSdkToTerraform(ctx, diags, *v.Primary, PrimaryValue{}.AttributeTypes(ctx))
		probe := tunnelConfigProbeSdkToTerraform(ctx, diags, *v.Probe)
		secondary := tunnelConfigNodeSdkToTerraform(ctx, diags, *v.Secondary, SecondaryValue{}.AttributeTypes(ctx))
		var r_state = map[string]attr.Value{
			"auto_provision":  auto_provision,
			"ike_lifetime":    types.Int64Value(int64(*v.IkeLifetime)),
			"ike_mode":        types.StringValue(string(*v.IkeMode)),
			"ike_proposals":   ike_proposals,
			"ipsec_lifetime":  types.Int64Value(int64(*v.IpsecLifetime)),
			"ipsec_proposals": ipsec_proposals,
			"local_id":        types.StringValue(*v.LocalId),
			"mode":            types.StringValue(string(*v.Mode)),
			"primary":         primary,
			"probe":           probe,
			"protocol":        types.StringValue(string(*v.Protocol)),
			"provider":        types.StringValue(string(*v.Provider)),
			"psk":             types.StringValue(*v.Psk),
			"secondary":       secondary,
			"version":         types.StringValue(string(*v.Version)),
		}
		r_object, e := NewTunnelConfigsValue(r_type, r_state)
		diags.Append(e...)
		state_value_map[k] = r_object
	}
	state_type := TunnelConfigsValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}
