package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

func ledSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.ApLed) LedValue {

	r_attr_type := LedValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"brightness": types.Int64Value(int64(d.GetBrightness())),
		"enabled":    types.BoolValue(d.GetEnabled()),
	}
	r, e := NewLedValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
