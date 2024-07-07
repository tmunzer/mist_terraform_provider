package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"mistapi/models"
)

func radiusAcctServersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.RadiusAcctServer {

	var data []models.RadiusAcctServer
	for _, plan_attr := range d.Elements() {
		var srv_plan_interface interface{} = plan_attr
		srv_plan := srv_plan_interface.(AcctServersValue)

		srv_data := models.RadiusAcctServer{}
		srv_data.Host = srv_plan.Host.ValueString()
		srv_data.Port = models.ToPointer(int(srv_plan.Port.ValueInt64()))
		srv_data.Secret = srv_plan.Secret.ValueString()
		srv_data.KeywrapEnabled = models.ToPointer(srv_plan.KeywrapEnabled.ValueBool())
		srv_data.KeywrapFormat = models.ToPointer(models.RadiusKeywrapFormatEnum(srv_plan.KeywrapFormat.ValueString()))
		srv_data.KeywrapKek = models.ToPointer(srv_plan.KeywrapKek.ValueString())
		srv_data.KeywrapMack = models.ToPointer(srv_plan.KeywrapMack.ValueString())
		data = append(data, srv_data)
	}
	return data
}

func radiusAuthServersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.RadiusAuthServer {

	var data []models.RadiusAuthServer
	for _, plan_attr := range d.Elements() {
		var srv_plan_interface interface{} = plan_attr
		srv_plan := srv_plan_interface.(AuthServersValue)

		srv_data := models.RadiusAuthServer{}
		srv_data.Host = srv_plan.Host.ValueString()
		srv_data.Port = models.ToPointer(int(srv_plan.Port.ValueInt64()))
		srv_data.Secret = srv_plan.Secret.ValueString()
		srv_data.KeywrapEnabled = models.ToPointer(srv_plan.KeywrapEnabled.ValueBool())
		srv_data.KeywrapFormat = models.ToPointer(models.RadiusKeywrapFormatEnum(srv_plan.KeywrapFormat.ValueString()))
		srv_data.KeywrapKek = models.ToPointer(srv_plan.KeywrapKek.ValueString())
		srv_data.KeywrapMack = models.ToPointer(srv_plan.KeywrapMack.ValueString())
		data = append(data, srv_data)
	}
	return data
}

func radiusConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d RadiusConfigValue) *models.RadiusConfig {

	data := models.RadiusConfig{}
	data.AcctInterimInterval = models.ToPointer(int(d.AcctInterimInterval.ValueInt64()))
	data.AcctServers = radiusAcctServersTerraformToSdk(ctx, diags, d.AcctServers)
	data.AuthServersRetries = models.ToPointer(int(d.AuthServersRetries.ValueInt64()))
	data.AuthServersTimeout = models.ToPointer(int(d.AuthServersTimeout.ValueInt64()))
	data.CoaEnabled = models.ToPointer(d.CoaEnabled.ValueBool())
	data.CoaPort = models.ToPointer(int(d.CoaPort.ValueInt64()))
	data.Network = models.ToPointer(d.Network.ValueString())
	data.SourceIp = models.ToPointer(d.SourceIp.ValueString())
	data.AuthServers = radiusAuthServersTerraformToSdk(ctx, diags, d.AuthServers)

	return &data
}
