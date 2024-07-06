package resource_org_wlan

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func radiusServersAcctSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.RadiusAcctServer) basetypes.ListValue {
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

func radiusServersAuthSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []models.RadiusAuthServer) basetypes.ListValue {
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
