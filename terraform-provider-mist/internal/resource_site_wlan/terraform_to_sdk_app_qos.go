package resource_site_wlan

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func appQosAppsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.MapValue) map[string]mistsdkgo.WlanAppQosAppsProperties {
	data_map := make(map[string]mistsdkgo.WlanAppQosAppsProperties)
	for k, v := range plan.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(AppsValue)
		data := mistsdkgo.NewWlanAppQosAppsProperties()
		data.SetDscp(int32(v_plan.Dscp.ValueInt64()))
		data.SetDstSubnet(v_plan.DstSubnet.ValueString())
		data.SetSrcSubnet(v_plan.SrcSubnet.ValueString())
		data_map[k] = *data
	}
	return data_map
}
func appQosOthersTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ListValue) []mistsdkgo.WlanAppQosOthersItem {
	var data_list []mistsdkgo.WlanAppQosOthersItem
	for _, v := range plan.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(OthersValue)
		data := mistsdkgo.NewWlanAppQosOthersItem()
		data.SetDscp(int32(v_plan.Dscp.ValueInt64()))
		data.SetDstSubnet(v_plan.DstSubnet.ValueString())
		data.SetPortRanges(v_plan.PortRanges.ValueString())
		data.SetProtocol(v_plan.Protocol.ValueString())
		data.SetSrcSubnet(v_plan.SrcSubnet.ValueString())
		data_list = append(data_list, *data)
	}
	return data_list
}

func appQosTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan AppQosValue) mistsdkgo.WlanAppQos {

	data := *mistsdkgo.NewWlanAppQos()

	apps := appQosAppsTerraformToSdk(ctx, diags, plan.Apps)
	data.SetApps(apps)

	data.SetEnabled(plan.Enabled.ValueBool())

	others := appQosOthersTerraformToSdk(ctx, diags, plan.Others)
	data.SetOthers(others)

	return data
}
