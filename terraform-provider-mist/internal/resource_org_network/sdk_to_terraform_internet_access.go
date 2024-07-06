package resource_org_network

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func InternetAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.NetworkInternetAccess) InternetAccessValue {
	var create_simple_service_policy basetypes.BoolValue
	var destination_nat basetypes.MapValue = destinationNatSdkToTerraform(ctx, diags, d.DestinationNat)
	var enabled basetypes.BoolValue
	var restricted basetypes.BoolValue
	var static_nac basetypes.MapValue = staticNatSdkToTerraform(ctx, diags, d.StaticNat)

	if d.CreateSimpleServicePolicy != nil {
		create_simple_service_policy = types.BoolValue(*d.CreateSimpleServicePolicy)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Restricted != nil {
		restricted = types.BoolValue(*d.Restricted)
	}

	data_map_attr_type := InternetAccessValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"create_simple_service_policy": create_simple_service_policy,
		"destination_nat":              destination_nat,
		"enabled":                      enabled,
		"restricted":                   restricted,
		"static_nat":                   static_nac,
	}
	data, e := NewInternetAccessValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
