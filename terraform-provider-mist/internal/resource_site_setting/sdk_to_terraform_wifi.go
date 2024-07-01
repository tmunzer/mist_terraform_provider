package resource_site_setting

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func wifiSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.SiteWifi) WifiValue {
	tflog.Debug(ctx, "wifiSdkToTerraform")

	r_attr_type := WifiValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"cisco_enabled":                         types.BoolValue(d.GetCiscoEnabled()),
		"disable_11k":                           types.BoolValue(d.GetDisable11k()),
		"disable_radios_when_power_constrained": types.BoolValue(d.GetDisableRadiosWhenPowerConstrained()),
		"enable_arp_spoof_check":                types.BoolValue(d.GetEnableArpSpoofCheck()),
		"enable_shared_radio_scanning":          types.BoolValue(d.GetEnableSharedRadioScanning()),
		"enabled":                               types.BoolValue(d.GetEnabled()),
		"locate_connected":                      types.BoolValue(d.GetLocateConnected()),
		"locate_unconnected":                    types.BoolValue(d.GetLocateUnconnected()),
		"mesh_allow_dfs":                        types.BoolValue(d.GetMeshAllowDfs()),
		"mesh_enable_crm":                       types.BoolValue(d.GetMeshEnableCrm()),
		"mesh_enabled":                          types.BoolValue(d.GetMeshEnabled()),
		"mesh_psk":                              types.StringValue(d.GetMeshPsk()),
		"mesh_ssid":                             types.StringValue(d.GetMeshSsid()),
		"proxy_arp":                             types.StringValue(string(d.GetProxyArp())),
	}
	r, e := NewWifiValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
