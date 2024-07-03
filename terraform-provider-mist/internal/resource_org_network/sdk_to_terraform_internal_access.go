package resource_org_network

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func InternalAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.NetworkInternalAccess) InternalAccessValue {

	data_map_attr_type := InternalAccessValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled": types.BoolValue(d.GetEnabled()),
	}
	data, e := NewInternalAccessValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
