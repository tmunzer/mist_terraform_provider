package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"
)

func switchMgmtProtectReCustomTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.ProtectReCustom {
	var data []models.ProtectReCustom
	for _, item := range d.Elements() {
		var item_interface interface{} = item
		item_obj := item_interface.(CustomValue)

		data_item := models.NewProtectReCustom()
		data_item.SetPortRange(item_obj.PortRange.ValueString())
		data_item.SetProtocol(models.ProtectReCustomProtocol(item_obj.Protocol.ValueString()))
		data_item.SetSubnet(mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.Subnet))

		data = append(data, *data_item)
	}
	return data
}
func switchMgmtProtectReTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.ProtectRe {
	data := models.NewProtectRe()
	if d.IsNull() || d.IsUnknown() {
		return *data
	} else {
		item, e := NewProtectReValue(ProtectReValue{}.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)
		var item_interface interface{} = item
		item_obj := item_interface.(ProtectReValue)

		customRe := switchMgmtProtectReCustomTerraformToSdk(ctx, diags, item_obj.Custom)

		data.SetAllowedServices(mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.AllowedServices))
		data.SetCustom(customRe)
		data.SetEnabled(item_obj.Enabled.ValueBool())
		data.SetTrustedHosts(mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.TrustedHosts))
		return *data
	}
}
func TacacsAcctServersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.TacacsAcctServer {

	var data []models.TacacsAcctServer
	for _, plan_attr := range d.Elements() {
		var srv_plan_interface interface{} = plan_attr
		srv_plan := srv_plan_interface.(models.TacacsAcctServer)
		srv_data := models.NewTacacsAcctServer()
		srv_data.SetHost(srv_plan.GetHost())
		srv_data.SetPort(srv_plan.GetPort())
		srv_data.SetSecret(srv_plan.GetSecret())
		srv_data.SetTimeout(int32(srv_plan.GetTimeout()))
		data = append(data, *srv_data)
	}
	return data
}
func TacacsAuthServersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.TacacsAuthServer {

	var data []models.TacacsAuthServer
	for _, plan_attr := range d.Elements() {
		var srv_plan_interface interface{} = plan_attr
		srv_plan := srv_plan_interface.(TacplusServersValue)
		srv_data := models.NewTacacsAuthServer()
		srv_data.SetHost(srv_plan.Host.ValueString())
		srv_data.SetPort(srv_plan.Port.ValueString())
		srv_data.SetSecret(srv_plan.Secret.ValueString())
		srv_data.SetTimeout(int32(srv_plan.Timeout.ValueInt64()))
		data = append(data, *srv_data)
	}
	return data
}
func switchMgmtTacacsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.Tacacs {

	data := models.NewTacacs()

	if d.IsNull() || d.IsUnknown() {
		return *data
	} else {
		item, e := NewProtectReValue(TacacsValue{}.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)
		var item_interface interface{} = item
		item_obj := item_interface.(TacacsValue)

		acct_servers := TacacsAcctServersTerraformToSdk(ctx, diags, item_obj.TacacctServers)
		auth_servers := TacacsAuthServersTerraformToSdk(ctx, diags, item_obj.TacplusServers)

		data.SetEnabled(item_obj.Enabled.ValueBool())
		data.SetNetwork(item_obj.Network.ValueString())
		data.SetAcctServers(acct_servers)
		data.SetTacplusServers(auth_servers)
		data.SetDefaultRole(models.TacacsDefaultRole(item_obj.DefaultRole.ValueString()))

		return *data
	}
}
func switchMgmtTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SwitchMgmtValue) models.SwitchMgmt {

	data := models.NewSwitchMgmt()
	if d.IsNull() || d.IsUnknown() {
		return *data
	} else {
		protectRe := switchMgmtProtectReTerraformToSdk(ctx, diags, d.ProtectRe)
		tacas := switchMgmtTacacsTerraformToSdk(ctx, diags, d.Tacacs)

		data.SetConfigRevert(int32(d.ConfigRevert.ValueInt64()))
		data.SetProtectRe(protectRe)
		data.SetRootPassword(d.RootPassword.ValueString())
		data.SetTacacs(tacas)

		return *data
	}

}