package resource_org_wlan

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func coaServersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []mistapigo.CoaServer) basetypes.ListValue {

	var data_list = []CoaServersValue{}
	for _, v := range data {
		data_map_value := map[string]attr.Value{
			"disable_event_timestamp_check": types.BoolValue(v.GetDisableEventTimestampCheck()),
			"enabled":                       types.BoolValue(v.GetEnabled()),
			"ip":                            types.StringValue(v.GetIp()),
			"port":                          types.Int64Value(int64(v.GetPort())),
			"secret":                        types.StringValue(v.GetSecret()),
		}
		data, e := NewCoaServersValue(CoaServersValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, CoaServersValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r

}
