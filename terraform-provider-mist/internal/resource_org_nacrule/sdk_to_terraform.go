package resource_org_nacrule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"
)

func SdkToTerraform(ctx context.Context, data *mistapigo.NacRule) (OrgNacruleModel, diag.Diagnostics) {
	var state OrgNacruleModel
	var diags diag.Diagnostics

	state.Id = types.StringValue(data.GetId())
	state.OrgId = types.StringValue(data.GetOrgId())
	state.Name = types.StringValue(data.GetName())

	state.Action = types.StringValue(string(data.GetAction()))
	state.ApplyTags = mist_transform.ListOfStringSdkToTerraform(ctx, data.GetApplyTags())
	state.Enabled = types.BoolValue(data.GetEnabled())
	state.Matching = matchingSdkToTerraform(ctx, &diags, data.GetMatching())
	state.NotMatching = notMatchingSdkToTerraform(ctx, &diags, data.GetMatching())
	state.Order = types.Int64Value(int64(data.GetOrder()))

	return state, diags
}
