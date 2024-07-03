package resource_site_wlan

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func radsecServersSkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []mistsdkgo.RadsecServer) basetypes.ListValue {
	var data_list = []ServersValue{}
	for _, v := range data {
		data_map_value := map[string]attr.Value{
			"host": types.StringValue(v.GetHost()),
			"port": types.Int64Value(int64(v.GetPort())),
		}
		data, e := NewServersValue(ServersValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, ServersValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}
func radsecSkToTerraform(ctx context.Context, diags *diag.Diagnostics, data mistsdkgo.Radsec) RadsecValue {

	plan_attr := map[string]attr.Value{
		"coa_enabled":     types.BoolValue(data.GetCoaEnabled()),
		"enabled":         types.BoolValue(data.GetEnabled()),
		"idle_timeout":    types.Int64Value(int64(data.GetIdleTimeout())),
		"mxcluster_ids":   mist_transform.ListOfStringSdkToTerraform(ctx, data.GetMxclusterIds()),
		"proxy_hosts":     mist_transform.ListOfStringSdkToTerraform(ctx, data.GetProxyHosts()),
		"server_name":     types.StringValue(data.GetServerName()),
		"servers":         radsecServersSkToTerraform(ctx, diags, data.GetServers()),
		"use_mxedge":      types.BoolValue(data.GetUseMxedge()),
		"use_site_mxedge": types.BoolValue(data.GetUseSiteMxedge()),
	}
	r, e := NewRadsecValue(RadsecValue{}.AttributeTypes(ctx), plan_attr)
	diags.Append(e...)

	return r

}
