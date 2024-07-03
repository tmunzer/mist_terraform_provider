package provider

import (
	"context"
	"fmt"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	"terraform-provider-mist/internal/resource_org_network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgNetworkResource{}
	_ resource.ResourceWithConfigure = &orgNetworkResource{}
)

func NewOrgNetworkResource() resource.Resource {
	return &orgNetworkResource{}
}

type orgNetworkResource struct {
	client *mistsdkgo.APIClient
}

func (r *orgNetworkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Network client")
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
func (r *orgNetworkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_network"
}

func (r *orgNetworkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_org_network.OrgNetworkResourceSchema(ctx)
}

func (r *orgNetworkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Network Create")
	var plan, state resource_org_network.OrgNetworkModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	network, diags := resource_org_network.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Network Create for Org "+plan.OrgId.ValueString())
	data, _, err := r.client.OrgsNetworksAPI.CreateOrgNetwork(ctx, plan.OrgId.ValueString()).Network(network).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating network",
			"Could not create network, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_network.SdkToTerraform(ctx, data)
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

func (r *orgNetworkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_network.OrgNetworkModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Network Read: network_id "+state.Id.ValueString())
	data, _, err := r.client.OrgsNetworksAPI.GetOrgNetwork(ctx, state.OrgId.ValueString(), state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting network",
			"Could not get network, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_network.SdkToTerraform(ctx, data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgNetworkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_network.OrgNetworkModel
	tflog.Info(ctx, "Starting Network Update")

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

	network, diags := resource_org_network.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Network Update for Network "+state.Id.ValueString())
	data, _, err := r.client.OrgsNetworksAPI.UpdateOrgNetwork(ctx, state.OrgId.ValueString(), state.Id.ValueString()).Network(network).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating network",
			"Could not update network, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_network.SdkToTerraform(ctx, data)
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

func (r *orgNetworkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_network.OrgNetworkModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Network Delete: network_id "+state.Id.ValueString())
	_, err := r.client.OrgsNetworksAPI.DeleteOrgNetwork(ctx, state.OrgId.ValueString(), state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating network",
			"Could not delete network, unexpected error: "+err.Error(),
		)
		return
	}
}
