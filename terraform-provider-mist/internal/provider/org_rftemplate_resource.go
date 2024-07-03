package provider

import (
	"context"
	"fmt"

	mistsdkgo "terraform-provider-mist/github.com/tmunzer/mist-sdk-go"
	"terraform-provider-mist/internal/resource_rftemplate"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgRfTemplateResource{}
	_ resource.ResourceWithConfigure = &orgRfTemplateResource{}
)

func NewOrgRfTemplate() resource.Resource {
	return &orgRfTemplateResource{}
}

type orgRfTemplateResource struct {
	client *mistsdkgo.APIClient
}

func (r *orgRfTemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist RfTemplate client")
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
func (r *orgRfTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_rftemplate"
}

func (r *orgRfTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_rftemplate.RftemplateResourceSchema(ctx)
}

func (r *orgRfTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting RfTemplate Create")
	var plan, state resource_rftemplate.RftemplateModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	rftemplate, diags := resource_rftemplate.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, _, err := r.client.OrgsRFTemplatesAPI.CreateOrgRfTemplate(ctx, plan.OrgId.ValueString()).RfTemplate(rftemplate).Execute()
	if err != nil {
		//url, _ := httpr.Location()
		resp.Diagnostics.AddError(
			"Error creating RfTemplate",
			"Could not create RfTemplate, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_rftemplate.SdkToTerraform(ctx, data)
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

func (r *orgRfTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_rftemplate.RftemplateModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting RfTemplate Read: rftemplate_id "+state.Id.ValueString())
	data, _, err := r.client.OrgsRFTemplatesAPI.GetOrgRfTemplate(ctx, state.OrgId.ValueString(), state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting RfTemplate",
			"Could not get RfTemplate, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_rftemplate.SdkToTerraform(ctx, data)
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

func (r *orgRfTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_rftemplate.RftemplateModel
	tflog.Info(ctx, "Starting RfTemplate Update")

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

	rftemplate, diags := resource_rftemplate.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting RfTemplate Update for RfTemplate "+state.Id.ValueString())
	data, _, err := r.client.OrgsRFTemplatesAPI.
		UpdateOrgRfTemplate(ctx, state.OrgId.ValueString(), state.Id.ValueString()).
		RfTemplate(rftemplate).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating RfTemplate",
			"Could not create RfTemplate, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_rftemplate.SdkToTerraform(ctx, data)
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

func (r *orgRfTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_rftemplate.RftemplateModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting RfTemplate Delete: rftemplate_id "+state.Id.ValueString())
	_, err := r.client.OrgsRFTemplatesAPI.DeleteOrgRfTemplate(ctx, state.OrgId.ValueString(), state.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating RfTemplate",
			"Could not delete RfTemplate, unexpected error: "+err.Error(),
		)
		return
	}
}
