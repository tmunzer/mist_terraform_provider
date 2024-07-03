package resource_org_wlan

import (
	"context"
	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	"terraform-provider-mist/internal/commons/hours"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func scheduleSkToTerraform(ctx context.Context, diags *diag.Diagnostics, data mistsdkgo.WlanSchedule) ScheduleValue {

	plan_attr := map[string]attr.Value{
		"enabled": types.BoolValue(data.GetEnabled()),
		"hours":   hours.HoursSdkToTerraform(ctx, diags, data.GetHours()),
	}
	r, e := NewScheduleValue(ScheduleValue{}.AttributeTypes(ctx), plan_attr)
	diags.Append(e...)

	return r

}
