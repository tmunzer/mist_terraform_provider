package provider

import (
	"context"
	"fmt"
	"terraform-provider-mist/internal/resource_wlan"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgWlanResource{}
	_ resource.ResourceWithConfigure = &orgWlanResource{}
)

func NewOrgWlan() resource.Resource {
	return &orgWlanResource{}
}

type orgWlanResource struct {
	client *mistsdkgo.APIClient
}

func (r *orgWlanResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Wlan client")
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

func (r *orgWlanResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_wlan"
}

func (r *orgWlanResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_wlan.WlanResourceSchema(ctx)
}

func (r *orgWlanResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Wlan Create")
	var plan, state resource_wlan.WlanModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wlan, diags := resource_wlan.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, _, err := r.client.OrgsWlansAPI.CreateOrgWlan(ctx, plan.OrgId.ValueString()).Wlan(wlan).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Wlan",
			"Could not create Wlan, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_wlan.SdkToTerraform(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *orgWlanResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_wlan.WlanModel
	tflog.Info(ctx, "Starting Wlan Read: wlan_id "+state.Id.ValueString())

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, _, err := r.client.OrgsWlansAPI.GetOrgWLAN(ctx, state.OrgId.ValueString(), state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting Wlan",
			"Could not get Wlan, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_wlan.SdkToTerraform(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgWlanResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_wlan.WlanModel
	tflog.Info(ctx, "Starting Wlan Update")

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	diags = req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wlan, diags := resource_wlan.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Wlan Update for Wlan "+state.Id.ValueString())
	data, _, err := r.client.OrgsWlansAPI.
		UpdateOrgWlan(ctx, state.OrgId.ValueString(), state.Id.ValueString()).
		Wlan(wlan).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Wlan",
			"Could not create Wlan, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_wlan.SdkToTerraform(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *orgWlanResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_wlan.WlanModel
	tflog.Info(ctx, "Starting Wlan Delete: wlan_id "+state.Id.ValueString())

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.OrgsWlansAPI.DeleteOrgWlan(ctx, state.OrgId.ValueString(), state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Wlan",
			"Could not delete Wlan, unexpected error: "+err.Error(),
		)
		return
	}
}
