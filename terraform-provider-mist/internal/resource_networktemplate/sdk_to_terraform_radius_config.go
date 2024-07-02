package resource_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func radiusServersAcctSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistsdkgo.RadiusAcctServer) basetypes.ListValue {
	var acct_value_list []attr.Value
	acct_value_list_type := AcctServersValue{}.AttributeTypes(ctx)
	for _, srv_data := range d {
		rc_srv_state_value := map[string]attr.Value{
			"host":            types.StringValue(srv_data.GetHost()),
			"port":            types.Int64Value(int64(srv_data.GetPort())),
			"secret":          types.StringValue(srv_data.GetSecret()),
			"keywrap_enabled": types.BoolValue(srv_data.GetKeywrapEnabled()),
			"keywrap_format":  types.StringValue(string(srv_data.GetKeywrapFormat())),
			"keywrap_kek":     types.StringValue(srv_data.GetKeywrapKek()),
			"keywrap_mack":    types.StringValue(srv_data.GetKeywrapMack()),
		}
		acct_server, e := NewAcctServersValue(acct_value_list_type, rc_srv_state_value)
		diags.Append(e...)

		acct_value_list = append(acct_value_list, acct_server)
	}

	acct_state_list_type := AcctServersValue{}.Type(ctx)
	acct_state_list, e := types.ListValueFrom(ctx, acct_state_list_type, acct_value_list)
	diags.Append(e...)

	return acct_state_list
}

func radiusServersAuthSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistsdkgo.RadiusAuthServer) basetypes.ListValue {
	var auth_value_list []attr.Value
	auth_value_list_type := AuthServersValue{}.AttributeTypes(ctx)
	for _, srv_data := range d {
		rc_srv_state_value := map[string]attr.Value{
			"host":            types.StringValue(srv_data.GetHost()),
			"port":            types.Int64Value(int64(srv_data.GetPort())),
			"secret":          types.StringValue(srv_data.GetSecret()),
			"keywrap_enabled": types.BoolValue(srv_data.GetKeywrapEnabled()),
			"keywrap_format":  types.StringValue(string(srv_data.GetKeywrapFormat())),
			"keywrap_kek":     types.StringValue(srv_data.GetKeywrapKek()),
			"keywrap_mack":    types.StringValue(srv_data.GetKeywrapMack()),
		}
		auth_server, e := NewAuthServersValue(auth_value_list_type, rc_srv_state_value)
		diags.Append(e...)

		auth_value_list = append(auth_value_list, auth_server)
	}

	auth_state_list_type := AuthServersValue{}.Type(ctx)
	auth_state_list, e := types.ListValueFrom(ctx, auth_state_list_type, auth_value_list)
	diags.Append(e...)

	return auth_state_list
}

func radiusConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.RadiusConfig) RadiusConfigValue {
	acct_state_result := radiusServersAcctSdkToTerraform(ctx, diags, d.AcctServers)
	auth_state_result := radiusServersAuthSdkToTerraform(ctx, diags, d.AuthServers)
	radius_config_type := RadiusConfigValue{}.AttributeTypes(ctx)
	radius_config_map := map[string]attr.Value{
		"acct_interim_interval": types.Int64Value(int64(d.GetAcctInterimInterval())),
		"auth_servers_retries":  types.Int64Value(int64(d.GetAuthServersRetries())),
		"auth_servers_timeout":  types.Int64Value(int64(d.GetAuthServersTimeout())),
		"coa_enabled":           types.BoolValue(d.GetCoaEnabled()),
		"coa_port":              types.Int64Value(int64(d.GetCoaPort())),
		"network":               types.StringValue(d.GetNetwork()),
		"source_ip":             types.StringValue(d.GetSourceIp()),
		"acct_servers":          acct_state_result,
		"auth_servers":          auth_state_result,
	}

	state_result, e := NewRadiusConfigValue(radius_config_type, radius_config_map)
	diags.Append(e...)
	return state_result
}
