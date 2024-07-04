package provider

import (
	"context"
	"fmt"

	"terraform-provider-mist/internal/resource_device_ap"

	mistapigo "github.com/tmunzer/mistapi-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &device_apResource{}
	_ resource.ResourceWithConfigure = &device_apResource{}
)

func NewDeviceApResource() resource.Resource {
	return &device_apResource{}
}

type device_apResource struct {
	client *mistapigo.APIClient
}

func (r *device_apResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist DeviceAp client")
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*mistapigo.APIClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *mistapigo.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}
func (r *device_apResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_ap"
}

func (r *device_apResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_device_ap.DeviceApResourceSchema(ctx)
}

func (r *device_apResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting DeviceAp Create")
	var plan, state resource_device_ap.DeviceApModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device_ap, diags := resource_device_ap.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Starting DeviceAp Create on Site "+plan.SiteId.ValueString()+" for device "+plan.Id.ValueString())
	data, _, err := r.client.SitesDevicesAPI.UpdateSiteDevice(ctx, plan.SiteId.ValueString(), plan.Id.ValueString()).MistDevice(device_ap).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating device_ap",
			"Could not create device_ap, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_device_ap.SdkToTerraform(ctx, data)
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

func (r *device_apResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_device_ap.DeviceApModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceAp Read: device_ap_id "+state.Id.ValueString())
	data, _, err := r.client.SitesDevicesAPI.GetSiteDevice(ctx, state.SiteId.ValueString(), state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_ap",
			"Could not get device_ap, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_device_ap.SdkToTerraform(ctx, data)
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

func (r *device_apResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_device_ap.DeviceApModel
	tflog.Info(ctx, "Starting DeviceAp Update")

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

	device_ap, diags := resource_device_ap.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceAp Update for DeviceAp "+state.Id.ValueString())
	data, _, err := r.client.SitesDevicesAPI.UpdateSiteDevice(ctx, state.SiteId.ValueString(), state.Id.ValueString()).MistDevice(device_ap).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating device_ap",
			"Could not update device_ap, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_device_ap.SdkToTerraform(ctx, data)
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

func (r *device_apResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_device_ap.DeviceApModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device_ap, e := resource_device_ap.DeleteTerraformToSdk(ctx)
	resp.Diagnostics.Append(e...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceAp Delete: device_ap_id "+state.Id.ValueString())
	_, _, err := r.client.SitesDevicesAPI.UpdateSiteDevice(ctx, state.SiteId.ValueString(), state.Id.ValueString()).MistDevice(device_ap).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating device_ap",
			"Could not delete device_ap, unexpected error: "+err.Error(),
		)
		return
	}
}
