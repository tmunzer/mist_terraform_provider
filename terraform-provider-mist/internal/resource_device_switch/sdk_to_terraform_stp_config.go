package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

// ////////////////// MIST NAC ///////////////////////
func stpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistapigo.SwitchStpConfig) StpConfigValue {
	mist_nac_attr_type := StpConfigValue{}.AttributeTypes(ctx)
	mist_nac_attr_value := map[string]attr.Value{
		"type": types.StringValue(string(d.GetType())),
	}

	r, e := NewStpConfigValue(mist_nac_attr_type, mist_nac_attr_value)
	diags.Append(e...)

	return r
}
