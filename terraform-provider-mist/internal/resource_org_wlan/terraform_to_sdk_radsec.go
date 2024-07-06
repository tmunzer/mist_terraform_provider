package resource_org_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"mistapi/models"
)

func radsecServersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.RadsecServer {
	var data_list []models.RadsecServer
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(ServersValue)
		data := models.NewRadsecServer()
		data.SetHost(plan.Host.ValueString())
		data.SetPort(int32(plan.Port.ValueInt64()))

		data_list = append(data_list, *data)
	}
	return data_list
}

func radsecTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d RadsecValue) models.Radsec {
	data := models.NewRadsec()

	data.SetCoaEnabled(d.CoaEnabled.ValueBool())
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetIdleTimeout(int32(d.IdleTimeout.ValueInt64()))
	data.SetMxclusterIds(mist_transform.ListOfStringTerraformToSdk(ctx, d.MxclusterIds))
	data.SetProxyHosts(mist_transform.ListOfStringTerraformToSdk(ctx, d.ProxyHosts))
	data.SetServerName(d.ServerName.ValueString())

	servers := radsecServersTerraformToSdk(ctx, diags, d.Servers)
	data.SetServers(servers)

	data.SetUseMxedge(d.UseMxedge.ValueBool())
	data.SetUseSiteMxedge(d.UseSiteMxedge.ValueBool())

	return *data
}
