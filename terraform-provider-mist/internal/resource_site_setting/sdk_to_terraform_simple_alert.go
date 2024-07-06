package resource_site_setting

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func simpleAlertArpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SimpleAlertArpFailure) basetypes.ObjectValue {
	tflog.Debug(ctx, "simpleAlertArpSdkToTerraform")

	r_attr_type := ArpFailureValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"client_count":   types.Int64Value(int64(d.GetClientCount())),
		"duration":       types.Int64Value(int64(d.GetDuration())),
		"incident_count": types.Int64Value(int64(d.GetIncidentCount())),
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
func simpleAlertDnsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SimpleAlertDnsFailure) basetypes.ObjectValue {
	tflog.Debug(ctx, "simpleAlertDnsSdkToTerraform")

	r_attr_type := DnsFailureValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"client_count":   types.Int64Value(int64(d.GetClientCount())),
		"duration":       types.Int64Value(int64(d.GetDuration())),
		"incident_count": types.Int64Value(int64(d.GetIncidentCount())),
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
func simpleAlertDhcpSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SimpleAlertDhcpFailure) basetypes.ObjectValue {
	tflog.Debug(ctx, "simpleAlertDhcpSdkToTerraform")

	r_attr_type := DhcpFailureValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"client_count":   types.Int64Value(int64(d.GetClientCount())),
		"duration":       types.Int64Value(int64(d.GetDuration())),
		"incident_count": types.Int64Value(int64(d.GetIncidentCount())),
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}

func simpleAlertSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SimpleAlert) SimpleAlertValue {
	tflog.Debug(ctx, "simpleAlertSdkToTerraform")

	r_attr_type := SimpleAlertValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"arp_failure":  simpleAlertArpSdkToTerraform(ctx, diags, d.GetArpFailure()),
		"dhcp_failure": simpleAlertDhcpSdkToTerraform(ctx, diags, d.GetDhcpFailure()),
		"dns_failure":  simpleAlertDnsSdkToTerraform(ctx, diags, d.GetDnsFailure()),
	}

	r, e := NewSimpleAlertValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
