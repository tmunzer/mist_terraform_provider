package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func dhcpSnoopingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.DhcpSnooping) DhcpSnoopingValue {
	data_attr_type := DhcpSnoopingValue{}.AttributeTypes(ctx)
	data_attr_value := map[string]attr.Value{
		"all_networks":           types.BoolValue(d.GetAllNetworks()),
		"enable_arp_spoof_check": types.BoolValue(d.GetEnableArpSpoofCheck()),
		"enable_ip_source_guard": types.BoolValue(d.GetEnableIpSourceGuard()),
		"enabled":                types.BoolValue(d.GetEnabled()),
		"networks":               mist_transform.ListOfStringSdkToTerraform(ctx, d.GetNetworks()),
	}

	r, e := NewDhcpSnoopingValue(data_attr_type, data_attr_value)
	diags.Append(e...)

	return r
}
