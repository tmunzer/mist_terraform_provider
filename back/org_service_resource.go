package provider

import (
	"context"
	"fmt"
	"mistapi"

	"terraform-provider-mist/internal/resource_service"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgServiceResource{}
	_ resource.ResourceWithConfigure = &orgServiceResource{}
)

func NewOrgServiceResource() resource.Resource {
	return &orgServiceResource{}
}

type orgServiceResource struct {
	client mistapi.ClientInterface
}

func (r *orgServiceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Service client")
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(mistapi.ClientInterface)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *mistapigo.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *orgServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_service"
}
func (r *orgServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_service.ServiceResourceSchema(ctx)
}

func (r *orgServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Service Create")
	var plan, state resource_service.ServiceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	service, diags := resource_service.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	orgId := service.GetOrgId()
	tflog.Info(ctx, "Starting Service Create for Org "+orgId)
	data, err := r.client.OrgsServicesAPI.CreateOrgService(ctx, orgId).Service(service

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating service",
			"Could not create service, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_service.SdkToTerraform(ctx, data.Data)
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

func (r *orgServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_service.ServiceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Service Read: service_id "+state.Id.ValueString())
	data, err := r.client.OrgsServicesAPI.GetOrgService(ctx, state.OrgId.ValueString(), state.Id.ValueString()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting service",
			"Could not get service, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_service.SdkToTerraform(ctx, data.Data)
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

func (r *orgServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_service.ServiceModel
	tflog.Info(ctx, "Starting Service Update")

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

	service, diags := resource_service.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Service Update for Service "+service.GetId())
	data, err := r.client.OrgsServicesAPI.UpdateOrgService(ctx, service.GetOrgId(), service.GetId()).Service(service

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating service",
			"Could not update service, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_service.SdkToTerraform(ctx, data.Data)
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

func (r *orgServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_service.ServiceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Service Delete: service_id "+state.Id.ValueString())
	_, err := r.client.OrgsServicesAPI.DeleteOrgService(ctx, state.OrgId.ValueString(), state.Id.ValueString()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating service",
			"Could not delete service, unexpected error: "+err.Error(),
		)
		return
	}
}
