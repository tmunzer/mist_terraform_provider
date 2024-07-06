package resource_org_networktemplate

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

		data_item := models.ProtectReCustom{}
		data_item.PortRange = models.ToPointer(item_obj.PortRange.ValueString())
		data_item.Protocol = models.ToPointer(models.ProtectReCustomProtocolEnum(item_obj.Protocol.ValueString()))
		data_item.Subnet = mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.Subnet)

		data = append(data, data_item)
	}
	return data
}
func switchMgmtProtectReTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.ProtectRe {
	data := models.ProtectRe{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		item, e := NewProtectReValue(ProtectReValue{}.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)
		var item_interface interface{} = item
		item_obj := item_interface.(ProtectReValue)

		data.AllowedServices = mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.AllowedServices)
		data.Custom = switchMgmtProtectReCustomTerraformToSdk(ctx, diags, item_obj.Custom)
		data.Enabled = models.ToPointer(item_obj.Enabled.ValueBool())
		data.TrustedHosts = mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.TrustedHosts)
		return &data
	}
}
func TacacsAcctServersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.TacacsAcctServer {

	var data []models.TacacsAcctServer
	for _, plan_attr := range d.Elements() {
		var srv_plan_interface interface{} = plan_attr
		srv_plan := srv_plan_interface.(models.TacacsAcctServer)

		srv_data := models.TacacsAcctServer{}
		srv_data.Host = srv_plan.Host
		srv_data.Port = srv_plan.Port
		srv_data.Secret = srv_plan.Secret
		srv_data.Timeout = models.ToPointer(int(*srv_plan.Timeout))
		data = append(data, srv_data)
	}
	return data
}
func TacacsAuthServersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.TacacsAuthServer {

	var data []models.TacacsAuthServer
	for _, plan_attr := range d.Elements() {
		var srv_plan_interface interface{} = plan_attr
		srv_plan := srv_plan_interface.(TacplusServersValue)

		srv_data := models.TacacsAuthServer{}
		srv_data.Host = models.ToPointer(srv_plan.Host.ValueString())
		srv_data.Port = models.ToPointer(srv_plan.Port.ValueString())
		srv_data.Secret = models.ToPointer(srv_plan.Secret.ValueString())
		srv_data.Timeout = models.ToPointer(int(srv_plan.Timeout.ValueInt64()))
		data = append(data, srv_data)
	}
	return data
}
func switchMgmtTacacsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.Tacacs {

	data := models.Tacacs{}

	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		item, e := NewProtectReValue(TacacsValue{}.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)
		var item_interface interface{} = item
		item_obj := item_interface.(TacacsValue)

		data.Enabled = models.ToPointer(item_obj.Enabled.ValueBool())
		data.Network = models.ToPointer(item_obj.Network.ValueString())
		data.AcctServers = TacacsAcctServersTerraformToSdk(ctx, diags, item_obj.TacacctServers)
		data.TacplusServers = TacacsAuthServersTerraformToSdk(ctx, diags, item_obj.TacplusServers)
		data.DefaultRole = models.ToPointer(models.TacacsDefaultRoleEnum(item_obj.DefaultRole.ValueString()))

		return &data
	}
}
func switchMgmtTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SwitchMgmtValue) *models.SwitchMgmt {

	data := models.SwitchMgmt{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {

		data.ConfigRevert = models.ToPointer(int(d.ConfigRevert.ValueInt64()))
		data.ProtectRe = switchMgmtProtectReTerraformToSdk(ctx, diags, d.ProtectRe)
		data.RootPassword = models.ToPointer(d.RootPassword.ValueString())
		data.Tacacs = switchMgmtTacacsTerraformToSdk(ctx, diags, d.Tacacs)

		return &data
	}

}
