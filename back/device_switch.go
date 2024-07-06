package provider

import (
	"context"
	"fmt"
	"mistapi"

	"terraform-provider-mist/internal/resource_device_switch"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &device_switchResource{}
	_ resource.ResourceWithConfigure = &device_switchResource{}
)

func NewDeviceSwitchResource() resource.Resource {
	return &device_switchResource{}
}

type device_switchResource struct {
	client mistapi.ClientInterface
}

func (r *device_switchResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist DeviceSwitch client")
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
func (r *device_switchResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_switch"
}

func (r *device_switchResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_device_switch.DeviceSwitchResourceSchema(ctx)
}

func (r *device_switchResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting DeviceSwitch Create")
	var plan, state resource_device_switch.DeviceSwitchModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device_switch, diags := resource_device_switch.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Starting DeviceSwitch Create on Site "+plan.SiteId.ValueString()+" for device "+plan.Id.ValueString())
	data, err := r.client.SitesDevicesAPI.UpdateSiteDevice(ctx, plan.SiteId.ValueString(), plan.Id.ValueString()).MistDevice(device_switch

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating device_switch",
			"Could not create device_switch, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_device_switch.SdkToTerraform(ctx, data.Data)
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

func (r *device_switchResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_device_switch.DeviceSwitchModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceSwitch Read: device_switch_id "+state.Id.ValueString())
	data, err := r.client.SitesDevicesAPI.GetSiteDevice(ctx, state.SiteId.ValueString(), state.Id.ValueString()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_switch",
			"Could not get device_switch, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_device_switch.SdkToTerraform(ctx, data.Data)
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

func (r *device_switchResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_device_switch.DeviceSwitchModel
	tflog.Info(ctx, "Starting DeviceSwitch Update")

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

	device_switch, diags := resource_device_switch.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceSwitch Update for DeviceSwitch "+state.Id.ValueString())
	data, err := r.client.SitesDevicesAPI.UpdateSiteDevice(ctx, state.SiteId.ValueString(), state.Id.ValueString()).MistDevice(device_switch

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating device_switch",
			"Could not update device_switch, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_device_switch.SdkToTerraform(ctx, data.Data)
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

func (r *device_switchResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_device_switch.DeviceSwitchModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device_switch, e := resource_device_switch.DeleteTerraformToSdk(ctx)
	resp.Diagnostics.Append(e...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceSwitch Delete: device_switch_id "+state.Id.ValueString())
	_, _, err := r.client.SitesDevicesAPI.UpdateSiteDevice(ctx, state.SiteId.ValueString(), state.Id.ValueString()).MistDevice(device_switch
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating device_switch",
			"Could not delete device_switch, unexpected error: "+err.Error(),
		)
		return
	}
}
