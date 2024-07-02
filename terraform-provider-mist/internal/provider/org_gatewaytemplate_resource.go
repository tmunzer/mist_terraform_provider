package provider

import (
	"context"
	"fmt"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	"terraform-provider-mist/internal/resource_gatewaytemplate"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgGatewaytemplateResource{}
	_ resource.ResourceWithConfigure = &orgGatewaytemplateResource{}
)

func NewOrgGatewayTemplate() resource.Resource {
	return &orgGatewaytemplateResource{}
}

type orgGatewaytemplateResource struct {
	client *mistsdkgo.APIClient
}

func (r *orgGatewaytemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist GatewayTemplate client")
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
func (r *orgGatewaytemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_gatewaytemplate"
}

func (r *orgGatewaytemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_gatewaytemplate.GatewaytemplateResourceSchema(ctx)
}

func (r *orgGatewaytemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting GatewayTemplate Create")
	var plan, state resource_gatewaytemplate.GatewaytemplateModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	gatewaytemplate, diags := resource_gatewaytemplate.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, _, err := r.client.OrgsGatewayTemplatesAPI.CreateOrgGatewayTemplate(ctx, plan.OrgId.ValueString()).GatewayTemplate(gatewaytemplate).Execute()
	if err != nil {
		//url, _ := httpr.Location()
		resp.Diagnostics.AddError(
			"Error creating GatewayTemplate",
			"Could not create GatewayTemplate, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_gatewaytemplate.SdkToTerraform(ctx, data)
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

func (r *orgGatewaytemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_gatewaytemplate.GatewaytemplateModel
	tflog.Info(ctx, "Starting GatewayTemplate Read: gatewaytemplate_id "+state.Id.ValueString())

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, _, err := r.client.OrgsGatewayTemplatesAPI.GetOrgGatewayTemplate(ctx, state.OrgId.ValueString(), state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting GatewayTemplate",
			"Could not get GatewayTemplate, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_gatewaytemplate.SdkToTerraform(ctx, data)
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

func (r *orgGatewaytemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_gatewaytemplate.GatewaytemplateModel
	tflog.Info(ctx, "Starting GatewayTemplate Update")

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

	gatewaytemplate, diags := resource_gatewaytemplate.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting GatewayTemplate Update for GatewayTemplate "+state.Id.ValueString())
	data, _, err := r.client.OrgsGatewayTemplatesAPI.
		UpdateOrgGatewayTemplate(ctx, state.OrgId.ValueString(), state.Id.ValueString()).
		GatewayTemplate(gatewaytemplate).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating GatewayTemplate",
			"Could not create GatewayTemplate, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_gatewaytemplate.SdkToTerraform(ctx, data)
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

func (r *orgGatewaytemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_gatewaytemplate.GatewaytemplateModel
	tflog.Info(ctx, "Starting GatewayTemplate Delete: gatewaytemplate_id "+state.Id.ValueString())

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.OrgsGatewayTemplatesAPI.DeleteOrgGatewayTemplate(ctx, state.OrgId.ValueString(), state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating GatewayTemplate",
			"Could not delete GatewayTemplate, unexpected error: "+err.Error(),
		)
		return
	}
}
