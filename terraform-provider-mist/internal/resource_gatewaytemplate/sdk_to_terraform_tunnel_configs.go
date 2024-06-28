package resource_gatewaytemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/utils/transform"
)

func tunnelConfigAutoProvNodeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.TunnelConfigsAutoProvisionNode, r_type map[string]attr.Type) basetypes.ObjectValue {
	r_attr_value := map[string]attr.Value{
		"num_hosts": types.StringValue(d.GetNumHosts()),
		"wan_names": mist_transform.ListOfStringSdkToTerraform(ctx, d.GetWanNames()),
	}
	r, e := basetypes.NewObjectValue(r_type, r_attr_value)
	diags.Append(e...)
	return r
}

func tunnelConfigAutoProvSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.TunnelConfigsAutoProvision) basetypes.ObjectValue {
	r_attr_type := AutoProvisionValue{}.AttributeTypes(ctx)

	latlng_type := map[string]attr.Type{
		"lat": basetypes.Float64Type{},
		"lng": basetypes.Float64Type{},
	}
	latlng_value := map[string]attr.Value{
		"lat": types.Float64Value(d.Latlng.GetLat()),
		"lng": types.Float64Value(d.Latlng.GetLng()),
	}
	latlng, e := NewLatlngValue(latlng_type, latlng_value)
	diags.Append(e...)

	primary := tunnelConfigAutoProvNodeSdkToTerraform(ctx, diags, d.GetPrimary(), PrimaryValue{}.AttributeTypes(ctx))
	secondary := tunnelConfigAutoProvNodeSdkToTerraform(ctx, diags, d.GetSecondary(), SecondaryValue{}.AttributeTypes(ctx))

	r_attr_value := map[string]attr.Value{
		"enable":    types.BoolValue(d.GetEnable()),
		"latlng":    latlng,
		"primary":   primary,
		"region":    types.StringValue(string(d.GetRegion())),
		"secondary": secondary,
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

func tunnelConfigIkeProposalSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistsdkgo.GatewayTemplateTunnelIkeProposal) basetypes.ListValue {
	var data_list = []IkeProposalsValue{}
	for _, v := range d {
		data_map_attr_type := IkeProposalsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"auth_algo": types.StringValue(string(v.GetAuthAlgo())),
			"dh_group":  types.StringValue(string(v.GetDhGroup())),
			"enc_algo":  types.StringValue(string(v.GetEncAlgo())),
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

func tunnelConfigIpsecProposalSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistsdkgo.GatewayTemplateTunnelIpsecProposal) basetypes.ListValue {
	var data_list = []IpsecProposalsValue{}
	for _, v := range d {
		data_map_attr_type := IpsecProposalsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"auth_algo": types.StringValue(string(v.GetAuthAlgo())),
			"dh_group":  types.StringValue(string(v.GetDhGroup())),
			"enc_algo":  types.StringValue(string(v.GetEncAlgo())),
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

func tunnelConfigNodeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.GatewayTemplateTunnelNode, r_type map[string]attr.Type) basetypes.ObjectValue {
	r_attr_value := map[string]attr.Value{
		"hosts":        mist_transform.ListOfStringSdkToTerraform(ctx, d.GetHosts()),
		"internal_ips": mist_transform.ListOfStringSdkToTerraform(ctx, d.GetInternalIps()),
		"probe_ips":    mist_transform.ListOfStringSdkToTerraform(ctx, d.GetProbeIps()),
		"remote_ids":   mist_transform.ListOfStringSdkToTerraform(ctx, d.GetRemoteIds()),
		"wan_names":    mist_transform.ListOfStringSdkToTerraform(ctx, d.GetWanNames()),
	}
	r, e := basetypes.NewObjectValue(r_type, r_attr_value)
	diags.Append(e...)
	return r
}

func tunnelConfigProbeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.GatewayTemplateTunnelProbe) basetypes.ObjectValue {
	r_attr_type := ProbeValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"interval":  types.Int64Value(int64(d.GetInterval())),
		"threshold": types.Int64Value(int64(d.GetThreshold())),
		"timeout":   types.Int64Value(int64(d.GetTimeout())),
		"type":      types.StringValue(string(d.GetType())),
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

func tunnelConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]mistsdkgo.TunnelConfigs) basetypes.MapValue {
	r_type := TunnelConfigsValue{}.AttributeTypes(ctx)
	state_value_map := make(map[string]attr.Value)
	for k, v := range d {
		auto_provision := tunnelConfigAutoProvSdkToTerraform(ctx, diags, v.GetAutoProvision())
		ike_proposals := tunnelConfigIkeProposalSdkToTerraform(ctx, diags, v.GetIkeProposals())
		ipsec_proposals := tunnelConfigIpsecProposalSdkToTerraform(ctx, diags, v.GetIpsecProposals())
		primary := tunnelConfigNodeSdkToTerraform(ctx, diags, v.GetPrimary(), PrimaryValue{}.AttributeTypes(ctx))
		probe := tunnelConfigProbeSdkToTerraform(ctx, diags, v.GetProbe())
		secondary := tunnelConfigNodeSdkToTerraform(ctx, diags, v.GetSecondary(), SecondaryValue{}.AttributeTypes(ctx))
		var r_state = map[string]attr.Value{
			"auto_provision":  auto_provision,
			"ike_lifetime":    types.Int64Value(int64(v.GetIkeLifetime())),
			"ike_mode":        types.StringValue(string(v.GetIkeMode())),
			"ike_proposals":   ike_proposals,
			"ipsec_lifetime":  types.Int64Value(int64(v.GetIpsecLifetime())),
			"ipsec_proposals": ipsec_proposals,
			"local_id":        types.StringValue(v.GetLocalId()),
			"mode":            types.StringValue(string(v.GetMode())),
			"primary":         primary,
			"probe":           probe,
			"protocol":        types.StringValue(string(v.GetProtocol())),
			"provider":        types.StringValue(string(v.GetProvider())),
			"psk":             types.StringValue(v.GetPsk()),
			"secondary":       secondary,
			"version":         types.StringValue(string(v.GetVersion())),
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
