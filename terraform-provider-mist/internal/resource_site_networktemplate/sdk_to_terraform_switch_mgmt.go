package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func switchMgmtProtecCustomtReSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistapigo.ProtectReCustom) basetypes.ListValue {
	tflog.Debug(ctx, "switchMgmtProtecCustomtReSdkToTerraform")
	var data_list = []CustomValue{}

	for _, item := range d {
		data_map_attr_type := CustomValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"port_range": types.StringValue(item.GetPortRange()),
			"protocol":   types.StringValue(string(*item.Protocol)),
			"subnet":     mist_transform.ListOfStringSdkToTerraform(ctx, item.GetSubnet()),
		}

		data, e := NewCustomValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	data_list_type := CustomValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)

	return r
}
func switchMgmtProtectReSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.ProtectRe) basetypes.ObjectValue {
	tflog.Debug(ctx, "switchMgmtProtectReSdkToTerraform")

	custom_re := switchMgmtProtecCustomtReSdkToTerraform(ctx, diags, d.GetCustom())

	data_map_attr_type := ProtectReValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"allowed_services": mist_transform.ListOfStringSdkToTerraform(ctx, d.GetAllowedServices()),
		"custom":           custom_re,
		"enabled":          types.BoolValue(d.GetEnabled()),
		"trusted_hosts":    mist_transform.ListOfStringSdkToTerraform(ctx, d.GetTrustedHosts()),
	}

	r, e := NewProtectReValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	o, e := r.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}

func switchMgmtTacacsAcctSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistapigo.TacacsAcctServer) basetypes.ListValue {
	tflog.Debug(ctx, "switchMgmtTacacsAcctSdkToTerraform")

	var acct_value_list []attr.Value
	acct_value_list_type := TacacctServersValue{}.AttributeTypes(ctx)
	for _, srv_data := range d {
		rc_srv_state_value := map[string]attr.Value{
			"host":    types.StringValue(srv_data.GetHost()),
			"port":    types.StringValue(srv_data.GetPort()),
			"secret":  types.StringValue(srv_data.GetSecret()),
			"timeout": types.StringValue(srv_data.GetSecret()),
		}
		acct_server, e := NewTacacctServersValue(acct_value_list_type, rc_srv_state_value)
		diags.Append(e...)

		acct_value_list = append(acct_value_list, acct_server)
	}

	acct_state_list_type := TacacctServersValue{}.Type(ctx)
	acct_state_list, e := types.ListValueFrom(ctx, acct_state_list_type, acct_value_list)
	diags.Append(e...)

	return acct_state_list
}
func switchMgmtTacacsAuthSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d []mistapigo.TacacsAuthServer) basetypes.ListValue {
	tflog.Debug(ctx, "switchMgmtTacacsAuthSdkToTerraform")

	var acct_value_list []attr.Value
	acct_value_list_type := TacplusServersValue{}.AttributeTypes(ctx)
	for _, srv_data := range d {
		rc_srv_state_value := map[string]attr.Value{
			"host":    types.StringValue(srv_data.GetHost()),
			"port":    types.StringValue(srv_data.GetPort()),
			"secret":  types.StringValue(srv_data.GetSecret()),
			"timeout": types.StringValue(srv_data.GetSecret()),
		}
		acct_server, e := NewTacplusServersValue(acct_value_list_type, rc_srv_state_value)
		diags.Append(e...)

		acct_value_list = append(acct_value_list, acct_server)
	}

	acct_state_list_type := TacplusServersValue{}.Type(ctx)
	acct_state_list, e := types.ListValueFrom(ctx, acct_state_list_type, acct_value_list)
	diags.Append(e...)

	return acct_state_list
}
func switchMgmtTacacsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.Tacacs) basetypes.ObjectValue {
	tflog.Debug(ctx, "switchMgmtTacacsSdkToTerraform")

	tacacs_acct_servers := switchMgmtTacacsAcctSdkToTerraform(ctx, diags, d.GetAcctServers())
	tacacs_auth_servers := switchMgmtTacacsAuthSdkToTerraform(ctx, diags, d.GetTacplusServers())

	data_map_attr_type := TacacsValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"acct_servers":    tacacs_acct_servers,
		"enabled":         types.BoolValue(d.GetEnabled()),
		"network":         types.StringValue(d.GetNetwork()),
		"tacplus_servers": tacacs_auth_servers,
		"default_role":    types.StringValue(string(d.GetDefaultRole())),
	}

	r, e := NewTacacsValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	o, e := r.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}
func switchMgmtSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.SwitchMgmt) SwitchMgmtValue {
	tflog.Debug(ctx, "switchMgmtSdkToTerraform")

	switch_mgmt_protect_re := switchMgmtProtectReSdkToTerraform(ctx, diags, d.GetProtectRe())
	switch_mgmt_tacacs := switchMgmtTacacsSdkToTerraform(ctx, diags, d.GetTacacs())

	data_map_attr_type := SwitchMgmtValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"config_revert": types.Int64Value(int64(d.GetConfigRevert())),
		"protect_re":    switch_mgmt_protect_re,
		"root_password": types.StringValue(d.GetRootPassword()),
		"tacacs":        switch_mgmt_tacacs,
	}

	state_result, e := NewSwitchMgmtValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return state_result
}
