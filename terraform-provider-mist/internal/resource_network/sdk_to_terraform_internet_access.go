package resource_network

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func InternetAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.NetworkInternetAccess) InternetAccessValue {

	data_map_attr_type := InternetAccessValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"create_simple_service_policy": types.BoolValue(d.GetCreateSimpleServicePolicy()),
		"destination_nat":              destinationNatSdkToTerraform(ctx, diags, d.GetDestinationNat()),
		"enabled":                      types.BoolValue(d.GetEnabled()),
		"restricted":                   types.BoolValue(d.GetRestricted()),
		"static_nat":                   staticNatSdkToTerraform(ctx, diags, d.GetStaticNat()),
	}
	data, e := NewInternetAccessValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
