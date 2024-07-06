package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"mistapi/models"
)

func simpleAlertArpTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, o basetypes.ObjectValue) models.SimpleAlertArpFailure {
	tflog.Debug(ctx, "simpleAlertArpTerraformToSdk")
	data := models.NewSimpleAlertArpFailure()
	if o.IsNull() || o.IsUnknown() {
		return *data
	} else {
		d := NewArpFailureValueMust(o.AttributeTypes(ctx), o.Attributes())

		data.SetClientCount(int32(d.ClientCount.ValueInt64()))
		data.SetDuration(int32(d.Duration.ValueInt64()))
		data.SetIncidentCount(int32(d.IncidentCount.ValueInt64()))

		return *data
	}
}

func simpleAlertDhcpTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, o basetypes.ObjectValue) models.SimpleAlertDhcpFailure {
	tflog.Debug(ctx, "simpleAlertDhcpTerraformToSdk")
	data := models.NewSimpleAlertDhcpFailure()
	if o.IsNull() || o.IsUnknown() {
		return *data
	} else {
		d := NewDhcpFailureValueMust(o.AttributeTypes(ctx), o.Attributes())

		data.SetClientCount(int32(d.ClientCount.ValueInt64()))
		data.SetDuration(int32(d.Duration.ValueInt64()))
		data.SetIncidentCount(int32(d.IncidentCount.ValueInt64()))

		return *data
	}
}

func simpleAlertDnsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, o basetypes.ObjectValue) models.SimpleAlertDnsFailure {
	tflog.Debug(ctx, "simpleAlertDnsTerraformToSdk")
	data := models.NewSimpleAlertDnsFailure()
	if o.IsNull() || o.IsUnknown() {
		return *data
	} else {
		d := NewDnsFailureValueMust(o.AttributeTypes(ctx), o.Attributes())

		data.SetClientCount(int32(d.ClientCount.ValueInt64()))
		data.SetDuration(int32(d.Duration.ValueInt64()))
		data.SetIncidentCount(int32(d.IncidentCount.ValueInt64()))

		return *data
	}
}

func simpleAlertTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SimpleAlertValue) models.SimpleAlert {
	tflog.Debug(ctx, "simpleAlertTerraformToSdk")
	data := models.NewSimpleAlert()

	arp := simpleAlertArpTerraformToSdk(ctx, diags, d.ArpFailure)
	data.SetArpFailure(arp)

	dhcp := simpleAlertDhcpTerraformToSdk(ctx, diags, d.DhcpFailure)
	data.SetDhcpFailure(dhcp)

	dns := simpleAlertDnsTerraformToSdk(ctx, diags, d.DnsFailure)
	data.SetDnsFailure(dns)

	return *data
}
