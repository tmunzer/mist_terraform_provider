package resource_site_networktemplate

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
		if !srv_plan.KeywrapEnabled.IsNull() && !srv_plan.KeywrapEnabled.IsUnknown() {
			srv_data.KeywrapEnabled = models.ToPointer(srv_plan.KeywrapEnabled.ValueBool())
		}
		if !srv_plan.KeywrapFormat.IsNull() && !srv_plan.KeywrapFormat.IsUnknown() {
			srv_data.KeywrapFormat = models.ToPointer(models.RadiusKeywrapFormatEnum(srv_plan.KeywrapFormat.ValueString()))
		}
		if !srv_plan.KeywrapKek.IsNull() && !srv_plan.KeywrapKek.IsUnknown() {
			srv_data.KeywrapKek = models.ToPointer(srv_plan.KeywrapKek.ValueString())
		}
		if !srv_plan.KeywrapMack.IsNull() && !srv_plan.KeywrapMack.IsUnknown() {
			srv_data.KeywrapMack = models.ToPointer(srv_plan.KeywrapMack.ValueString())
		}
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
		if !srv_plan.KeywrapEnabled.IsNull() && !srv_plan.KeywrapEnabled.IsUnknown() {
			srv_data.KeywrapEnabled = models.ToPointer(srv_plan.KeywrapEnabled.ValueBool())
		}
		if !srv_plan.KeywrapFormat.IsNull() && !srv_plan.KeywrapFormat.IsUnknown() {
			srv_data.KeywrapFormat = models.ToPointer(models.RadiusKeywrapFormatEnum(srv_plan.KeywrapFormat.ValueString()))
		}
		if !srv_plan.KeywrapKek.IsNull() && !srv_plan.KeywrapKek.IsUnknown() {
			srv_data.KeywrapKek = models.ToPointer(srv_plan.KeywrapKek.ValueString())
		}
		if !srv_plan.KeywrapMack.IsNull() && !srv_plan.KeywrapMack.IsUnknown() {
			srv_data.KeywrapMack = models.ToPointer(srv_plan.KeywrapMack.ValueString())
		}
		data = append(data, srv_data)
	}
	return data
}

func radiusConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d RadiusConfigValue) *models.RadiusConfig {

	data := models.RadiusConfig{}
	if !d.AcctInterimInterval.IsNull() && !d.AcctInterimInterval.IsUnknown() {
		data.AcctInterimInterval = models.ToPointer(int(d.AcctInterimInterval.ValueInt64()))
	}
	if !d.AcctServers.IsNull() && !d.AcctServers.IsUnknown() {
		data.AcctServers = radiusAcctServersTerraformToSdk(ctx, diags, d.AcctServers)
	}
	if !d.AuthServersRetries.IsNull() && !d.AuthServersRetries.IsUnknown() {
		data.AuthServersRetries = models.ToPointer(int(d.AuthServersRetries.ValueInt64()))
	}
	if !d.AuthServersTimeout.IsNull() && !d.AuthServersTimeout.IsUnknown() {
		data.AuthServersTimeout = models.ToPointer(int(d.AuthServersTimeout.ValueInt64()))
	}
	if !d.CoaEnabled.IsNull() && !d.CoaEnabled.IsUnknown() {
		data.CoaEnabled = models.ToPointer(d.CoaEnabled.ValueBool())
	}
	if !d.CoaPort.IsNull() && !d.CoaPort.IsUnknown() {
		data.CoaPort = models.ToPointer(int(d.CoaPort.ValueInt64()))
	}
	if !d.Network.IsNull() && !d.Network.IsUnknown() {
		data.Network = models.ToPointer(d.Network.ValueString())
	}
	if !d.SourceIp.IsNull() && !d.SourceIp.IsUnknown() {
		data.SourceIp = models.ToPointer(d.SourceIp.ValueString())
	}
	if !d.AuthServers.IsNull() && !d.AuthServers.IsUnknown() {
		data.AuthServers = radiusAuthServersTerraformToSdk(ctx, diags, d.AuthServers)
	}

	return &data
}
