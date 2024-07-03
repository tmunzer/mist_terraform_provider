package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func wifiTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d WifiValue) mistapigo.SiteWifi {
	tflog.Debug(ctx, "wifiTerraformToSdk")
	data := mistapigo.NewSiteWifi()

	data.SetCiscoEnabled(d.CiscoEnabled.ValueBool())
	data.SetDisable11k(d.Disable11k.ValueBool())
	data.SetDisableRadiosWhenPowerConstrained(d.DisableRadiosWhenPowerConstrained.ValueBool())
	data.SetEnableArpSpoofCheck(d.EnableArpSpoofCheck.ValueBool())
	data.SetEnableSharedRadioScanning(d.EnableSharedRadioScanning.ValueBool())
	data.SetEnabled(d.Enabled.ValueBool())
	data.SetLocateConnected(d.LocateConnected.ValueBool())
	data.SetLocateUnconnected(d.LocateUnconnected.ValueBool())
	data.SetMeshAllowDfs(d.MeshAllowDfs.ValueBool())
	data.SetMeshEnableCrm(d.MeshEnableCrm.ValueBool())
	data.SetMeshEnabled(d.MeshEnabled.ValueBool())
	data.SetMeshPsk(d.MeshPsk.ValueString())
	data.SetMeshSsid(d.MeshSsid.ValueString())
	data.SetProxyArp(mistapigo.SiteWifiProxyArp(d.ProxyArp.ValueString()))

	return *data
}
