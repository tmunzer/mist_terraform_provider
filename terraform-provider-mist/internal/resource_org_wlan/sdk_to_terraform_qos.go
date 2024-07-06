package resource_org_wlan

import (
	"context"

	"mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func qosSkToTerraform(ctx context.Context, diags *diag.Diagnostics, data models.WlanQos) QosValue {

	plan_attr := map[string]attr.Value{
		"class":     types.StringValue(string(data.GetClass())),
		"overwrite": types.BoolValue(data.GetOverwrite()),
	}
	r, e := NewQosValue(QosValue{}.AttributeTypes(ctx), plan_attr)
	diags.Append(e...)

	return r

}
