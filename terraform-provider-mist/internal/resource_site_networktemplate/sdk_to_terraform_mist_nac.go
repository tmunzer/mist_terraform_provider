package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
)

// ////////////////// MIST NAC ///////////////////////
func mistNacSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d mistsdkgo.NetworkTemplateMistNac) MistNacValue {
	mist_nac_attr_type := MistNacValue{}.AttributeTypes(ctx)
	mist_nac_attr_value := map[string]attr.Value{
		"enabled": types.BoolValue(d.GetEnabled()),
		"network": types.StringValue(d.GetNetwork()),
	}

	r, e := NewMistNacValue(mist_nac_attr_type, mist_nac_attr_value)
	diags.Append(e...)

	return r
}
