package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func clientAuthBridgeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.ApClientBridgeAuth) basetypes.ObjectValue {

	r_attr_type := AuthValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"psk":  types.StringValue(d.GetPsk()),
		"type": types.StringValue(string(d.GetType())),
	}
	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
func clientBridgeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.ApClientBridge) ClientBridgeValue {
	tflog.Debug(ctx, "clientBridgeSdkToTerraform")

	r_attr_type := ClientBridgeValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"auth":    clientAuthBridgeSdkToTerraform(ctx, diags, d.GetAuth()),
		"enabled": types.BoolValue(d.GetEnabled()),
		"ssid":    types.StringValue(d.GetSsid()),
	}
	r, e := NewClientBridgeValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
