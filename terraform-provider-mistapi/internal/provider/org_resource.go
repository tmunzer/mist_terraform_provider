package provider

import (
	"context"
	"fmt"
	"math/big"
	mistsdkgo "terraform-provider-mistapi/github.com/tmunzer/mist-sdk-go"
	"terraform-provider-mistapi/internal/resource_org"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgResource{}
	_ resource.ResourceWithConfigure = &orgResource{}
)

func NewOrgResource() resource.Resource {
	return &orgResource{}
}

type orgResource struct {
	client *mistsdkgo.APIClient
}

func (r *orgResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Site client")
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*mistsdkgo.APIClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *mistsdkgo.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}
func (r *orgResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org"
}

func (r *orgResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_org.OrgResourceSchema(ctx)
}

func (r *orgResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Org Create")
	var plan resource_org.OrgModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data := processOrgPlan(&plan)
	data, _, err := r.client.OrgsAPI.CreateOrg(ctx).Org(*data).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating org",
			"Could not create org, unexpected error: "+err.Error(),
		)
		return
	}

	plan = processOrgData(data)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *orgResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org.OrgModel
	tflog.Info(ctx, "Starting Org Read: org_id "+state.Id.ValueString())

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, _, err := r.client.OrgsAPI.GetOrg(ctx, state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org",
			"Could not get org, unexpected error: "+err.Error(),
		)
		return
	}
	state = processOrgData(data)
	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data resource_org.OrgModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	//resp.Diagnostics.Append(callOrgAPI(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *orgResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org.OrgModel
	tflog.Info(ctx, "Starting Org Delete: org_id "+state.Id.ValueString())

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.OrgsAPI.DeleteOrg(ctx, state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating org",
			"Could not delete org, unexpected error: "+err.Error(),
		)
		return
	}
}

func processOrgData(data *mistsdkgo.Org) resource_org.OrgModel {
	var plan resource_org.OrgModel

	plan.AlarmtemplateId = types.StringValue(data.GetAlarmtemplateId())
	plan.AllowMist = types.BoolValue(data.GetAllowMist())
	plan.ModifiedTime = types.NumberValue(big.NewFloat(float64(data.GetModifiedTime())))
	plan.CreatedTime = types.NumberValue(big.NewFloat(float64(data.GetCreatedTime())))
	plan.Id = types.StringValue(data.GetId())
	plan.LogoUrl = types.StringValue(data.GetLogoUrl())
	plan.MspId = types.StringValue(data.GetMspId())
	plan.MspName = types.StringValue(data.GetMspName())
	plan.Name = types.StringValue(data.GetName())
	plan.SessionExpiry = types.NumberValue(big.NewFloat(float64(data.GetSessionExpiry())))
	var items []attr.Value
	for _, item := range data.GetOrggroupIds() {
		tmp := types.StringValue(item)
		items = append(items, tmp)
	}
	plan.OrggroupIds, _ = types.ListValue(types.StringType, items)
	return plan
}

func processOrgPlan(plan *resource_org.OrgModel) *mistsdkgo.Org {
	var data *mistsdkgo.Org
	data.SetName(plan.Name.ValueString())
	return data
}