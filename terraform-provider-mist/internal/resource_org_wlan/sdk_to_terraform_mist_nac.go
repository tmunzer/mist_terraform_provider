package resource_org_wlan

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func mistNacdSkToTerraform(ctx context.Context, diags *diag.Diagnostics, data models.WlanMistNac) MistNacValue {

	plan_attr := map[string]attr.Value{
		"enabled": types.BoolValue(data.GetEnabled()),
	}
	r, e := NewMistNacValue(MistNacValue{}.AttributeTypes(ctx), plan_attr)
	diags.Append(e...)

	return r

}
