package provider

import (
	"context"
	"fmt"
	"terraform-provider-mist/internal/resource_wlantemplate"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgWlanTemplateResource{}
	_ resource.ResourceWithConfigure = &orgWlanTemplateResource{}
)

func NewOrgWlanTemplate() resource.Resource {
	return &orgWlanTemplateResource{}
}

type orgWlanTemplateResource struct {
	client *mistsdkgo.APIClient
}

func (r *orgWlanTemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist WlanTemplate client")
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
func (r *orgWlanTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_wlantemplate"
}

func (r *orgWlanTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_wlantemplate.WlantemplateResourceSchema(ctx)
}

func (r *orgWlanTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting WlanTemplate Create")
	var plan, state resource_wlantemplate.WlantemplateModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wlantemplate, diags := resource_wlantemplate.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, _, err := r.client.OrgsWLANTemplatesAPI.CreateOrgTemplate(ctx, plan.OrgId.ValueString()).Template(wlantemplate).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating WlanTemplate",
			"Could not create WlanTemplate, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_wlantemplate.SdkToTerraform(ctx, data)
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

func (r *orgWlanTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_wlantemplate.WlantemplateModel
	tflog.Info(ctx, "Starting WlanTemplate Read: wlantemplate_id "+state.Id.ValueString())

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, _, err := r.client.OrgsWLANTemplatesAPI.GetOrgTemplate(ctx, state.OrgId.ValueString(), state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting WlanTemplate",
			"Could not get WlanTemplate, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_wlantemplate.SdkToTerraform(ctx, data)
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

func (r *orgWlanTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_wlantemplate.WlantemplateModel
	tflog.Info(ctx, "Starting WlanTemplate Update")

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

	wlantemplate, diags := resource_wlantemplate.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting WlanTemplate Update for WlanTemplate "+state.Id.ValueString())
	data, _, err := r.client.OrgsWLANTemplatesAPI.
		UpdateOrgTemplate(ctx, state.OrgId.ValueString(), state.Id.ValueString()).
		Template(wlantemplate).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating WlanTemplate",
			"Could not create WlanTemplate, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_wlantemplate.SdkToTerraform(ctx, data)
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

func (r *orgWlanTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_wlantemplate.WlantemplateModel
	tflog.Info(ctx, "Starting WlanTemplate Delete: wlantemplate_id "+state.Id.ValueString())

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.OrgsWLANTemplatesAPI.DeleteOrgTemplate(ctx, state.OrgId.ValueString(), state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating WlanTemplate",
			"Could not delete WlanTemplate, unexpected error: "+err.Error(),
		)
		return
	}
}