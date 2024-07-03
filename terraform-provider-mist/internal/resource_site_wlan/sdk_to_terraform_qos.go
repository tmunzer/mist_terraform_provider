package resource_site_wlan

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func qosSkToTerraform(ctx context.Context, diags *diag.Diagnostics, data mistapigo.WlanQos) QosValue {

	plan_attr := map[string]attr.Value{
		"class":     types.StringValue(string(data.GetClass())),
		"overwrite": types.BoolValue(data.GetOverwrite()),
	}
	r, e := NewQosValue(QosValue{}.AttributeTypes(ctx), plan_attr)
	diags.Append(e...)

	return r

}
