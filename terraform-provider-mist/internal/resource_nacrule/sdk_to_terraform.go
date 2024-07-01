package resource_nacrule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	mist_transform "terraform-provider-mist/internal/provider/commons/utils"
)

func SdkToTerraform(ctx context.Context, data *mistsdkgo.NacRule) (NacruleModel, diag.Diagnostics) {
	var state NacruleModel
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
