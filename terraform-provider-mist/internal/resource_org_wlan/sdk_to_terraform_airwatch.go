package resource_org_wlan

import (
	"context"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func airwatchSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data mistapigo.WlanAirwatch) AirwatchValue {

	plan_attr := map[string]attr.Value{
		"api_key":     types.StringValue(data.GetApiKey()),
		"console_url": types.StringValue(data.GetConsoleUrl()),
		"enabled":     types.BoolValue(data.GetEnabled()),
		"password":    types.StringValue(data.GetPassword()),
		"username":    types.StringValue(data.GetUsername()),
	}
	state, e := NewAirwatchValue(AirwatchValue{}.AttributeTypes(ctx), plan_attr)
	diags.Append(e...)

	return state

}
