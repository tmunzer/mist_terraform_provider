package provider

import (
	"context"
	"fmt"

	"terraform-provider-mist/internal/resource_device_gateway_cluster"

	"mistapi"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &deviceGatewayClusterResource{}
	_ resource.ResourceWithConfigure = &deviceGatewayClusterResource{}
)

func NewDeviceGatewayClusterResource() resource.Resource {
	return &deviceGatewayClusterResource{}
}

type deviceGatewayClusterResource struct {
	client mistapi.ClientInterface
}

func (r *deviceGatewayClusterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist DeviceGatewayCluster client")
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
func (r *deviceGatewayClusterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_gateway_cluster"
}

func (r *deviceGatewayClusterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_device_gateway_cluster.DeviceGatewayClusterResourceSchema(ctx)
}

func (r *deviceGatewayClusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting DeviceGatewayCluster Create")
	var plan, state resource_device_gateway_cluster.DeviceGatewayClusterModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device_gateway_cluster, diags := resource_device_gateway_cluster.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway_cluster site_id from plan",
			"Could not get device_gateway_cluster site_id, unexpected error: "+err.Error(),
		)
		return
	}

	deviceId, err := uuid.Parse(plan.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway_cluster device_id from plan",
			"Could not get device_gateway_cluster device_id, unexpected error: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, "Starting DeviceGatewayCluster Create on Site "+plan.SiteId.ValueString()+" for device "+plan.DeviceId.ValueString())
	data, err := r.client.SitesDevicesWANCluster().CreateSiteDeviceHaCluster(ctx, siteId, deviceId, device_gateway_cluster)

	if data.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error creating device_gateway_cluster",
			"Could not create device_gateway_cluster, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_device_gateway_cluster.SdkToTerraform(ctx, siteId, deviceId, &data.Data)
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

func (r *deviceGatewayClusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_device_gateway_cluster.DeviceGatewayClusterModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceGatewayCluster Read: device_gateway_cluster_id "+state.DeviceId.ValueString())
	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway_cluster site_id from state",
			"Could not get device_gateway_cluster site_id, unexpected error: "+err.Error(),
		)
		return
	}

	deviceId, err := uuid.Parse(state.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway_cluster device_id from state",
			"Could not get device_gateway_cluster device_id, unexpected error: "+err.Error(),
		)
		return
	}

	data, err := r.client.SitesDevicesWANCluster().GetSiteDeviceHaClusterNode(ctx, siteId, deviceId)
	if data.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway_cluster",
			"Could not get device_gateway_cluster, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_device_gateway_cluster.SdkToTerraform(ctx, siteId, deviceId, &data.Data)
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

func (r *deviceGatewayClusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state resource_device_gateway_cluster.DeviceGatewayClusterModel
	// 	tflog.Info(ctx, "Starting DeviceGatewayCluster Update")

	// 	diags := resp.State.Get(ctx, &state)
	// 	resp.Diagnostics.Append(diags...)
	// 	if resp.Diagnostics.HasError() {
	// 		return
	// 	}

	// 	diags = req.Plan.Get(ctx, &plan)
	// 	resp.Diagnostics.Append(diags...)
	// 	if resp.Diagnostics.HasError() {
	// 		return
	// 	}

	// 	device_gateway_cluster, diags := resource_device_gateway_cluster.TerraformToSdk(ctx, &plan)
	// 	resp.Diagnostics.Append(diags...)
	// 	if resp.Diagnostics.HasError() {
	// 		return
	// 	}

	// 	tflog.Info(ctx, "Starting DeviceGatewayCluster Update for DeviceGatewayCluster "+state.DeviceId.ValueString())

	// 	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	// 	if err != nil {
	// 		resp.Diagnostics.AddError(
	// 			"Error getting device_gateway_cluster site_id from state",
	// 			"Could not get device_gateway_cluster site_id, unexpected error: "+err.Error(),
	// 		)
	// 		return
	// 	}

	// 	deviceId, err := uuid.Parse(plan.DeviceId.ValueString())
	// 	if err != nil {
	// 		resp.Diagnostics.AddError(
	// 			"Error getting device_gateway_cluster device_id from state",
	// 			"Could not get device_gateway_cluster device_id, unexpected error: "+err.Error(),
	// 		)
	// 		return
	// 	}

	// 	data, err := r.client.SitesDevices().UpdateSiteDevice(ctx, siteId, deviceId, &device_gateway_cluster)

	// 	if data.Response.StatusCode != 200 && err != nil {
	// 		resp.Diagnostics.AddError(
	// 			"Error updating device_gateway_cluster",
	// 			"Could not update device_gateway_cluster, unexpected error: "+err.Error(),
	// 		)
	// 		return
	// 	}

	// 	body, _ := io.ReadAll(data.Response.Body)
	// 	mist_gateway := models.DeviceGatewayCluster{}
	// 	json.Unmarshal(body, &mist_gateway)
	// 	state, diags = resource_device_gateway_cluster.SdkToTerraform(ctx, &mist_gateway)
	// 	resp.Diagnostics.Append(diags...)
	// 	if resp.Diagnostics.HasError() {
	// 		return
	// 	}

	diags := resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *deviceGatewayClusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_device_gateway_cluster.DeviceGatewayClusterModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceGatewayCluster Delete: device_gateway_cluster_id "+state.DeviceId.ValueString())

	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway_cluster site_id from state",
			"Could not get device_gateway_cluster site_id, unexpected error: "+err.Error(),
		)
		return
	}

	deviceId, err := uuid.Parse(state.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway_cluster device_id from state",
			"Could not get device_gateway_cluster device_id, unexpected error: "+err.Error(),
		)
		return
	}

	_, err = r.client.SitesDevicesWANCluster().DeleteSiteDeviceHaCluster(ctx, siteId, deviceId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating device_gateway_cluster",
			"Could not delete device_gateway_cluster, unexpected error: "+err.Error(),
		)
		return
	}
}