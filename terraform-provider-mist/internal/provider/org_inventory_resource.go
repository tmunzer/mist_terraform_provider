package provider

import (
	"context"
	"fmt"
	"strings"

	"terraform-provider-mist/internal/resource_org_inventory"

	"mistapi"
	"mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgInventoryResource{}
	_ resource.ResourceWithConfigure = &orgInventoryResource{}
)

func NewOrgInventory() resource.Resource {
	return &orgInventoryResource{}
}

type orgInventoryResource struct {
	client mistapi.ClientInterface
}

func (r *orgInventoryResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Inventory client")
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
func (r *orgInventoryResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_inventory"
}

func (r *orgInventoryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_org_inventory.OrgInventoryResourceSchema(ctx)
}

func (r *orgInventoryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Inventory Create")
	var plan, state resource_org_inventory.OrgInventoryModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var magics []string
	for _, v := range plan.Devices.Elements() {
		var vi interface{} = v
		device := vi.(resource_org_inventory.DevicesValue)
		magics = append(magics, device.Magic.ValueString())

	}

	orgId := uuid.MustParse(plan.OrgId.ValueString())
	response_body, err := r.client.OrgsInventory().AddOrgInventory(ctx, orgId, magics)
	tflog.Info(ctx, "response for API Call to claim devices:", map[string]interface{}{
		"added":      strings.Join(response_body.Data.Added, ", "),
		"duplicated": strings.Join(response_body.Data.Duplicated, ", "),
		"error":      strings.Join(response_body.Data.Error, ", "),
	})
	if err != nil {
		//url, _ := httpr.Location()
		resp.Diagnostics.AddError(
			"Error creating Inventory",
			"Could not create Inventory, unexpected error: "+err.Error(),
		)
		return
	}

	tflog.Info(ctx, "Starting Inventory state refresh: org_id  "+state.OrgId.ValueString())
	var serial string
	var model string
	var mType models.DeviceTypeEnum
	var mac string
	var siteId string
	var vcMac string
	var vc bool = true
	var unassigned bool
	var limit int = 1000
	var page int
	tflog.Info(ctx, "Starting Inventory Read: org_id  "+state.OrgId.ValueString())
	data, err := r.client.OrgsInventory().GetOrgInventory(ctx, orgId, &serial, &model, &mType, &mac, &siteId, &vcMac, &vc, &unassigned, &limit, &page)
	if err != nil {
		//url, _ := httpr.Location()
		resp.Diagnostics.AddError(
			"Error creating Inventory",
			"Could not create Inventory, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_inventory.SdkToTerraform(ctx, plan.OrgId.ValueString(), data.Data)
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

func (r *orgInventoryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_inventory.OrgInventoryModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId := uuid.MustParse(state.OrgId.ValueString())
	tflog.Info(ctx, "Starting Inventory Read: org_id  "+state.OrgId.ValueString())
	var serial string
	var model string
	var mType models.DeviceTypeEnum
	var mac string
	var siteId string
	var vcMac string
	var vc bool = true
	var unassigned bool
	var limit int = 1000
	var page int
	tflog.Info(ctx, "Starting Inventory Read: org_id  "+state.OrgId.ValueString())
	data, err := r.client.OrgsInventory().GetOrgInventory(ctx, orgId, &serial, &model, &mType, &mac, &siteId, &vcMac, &vc, &unassigned, &limit, &page)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting Inventory",
			"Could not get Inventory, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_inventory.SdkToTerraform(ctx, state.OrgId.ValueString(), data.Data)
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

func (r *orgInventoryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_inventory.OrgInventoryModel
	tflog.Info(ctx, "Starting Inventory Update")

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

	orgId := uuid.MustParse(plan.OrgId.ValueString())
	var claim_devices []string
	var unclaim_devices []string
	var unassign_devices []string
	asign_claimed_devices := make(map[string]string)
	assign_devices := make(map[string][]string)

	for _, dev_plan_attr := range plan.Devices.Elements() {
		var dpi interface{} = dev_plan_attr
		var dev_plan = dpi.(resource_org_inventory.DevicesValue)
		var op string = ""
		var device_mac = ""
		already_claimed := false
		for _, dev_state_attr := range state.Devices.Elements() {
			var dsi interface{} = dev_state_attr
			var dev_state = dsi.(resource_org_inventory.DevicesValue)
			if dev_plan.Magic == dev_state.Magic {
				already_claimed = true
				// Site ID not in state but planed > must be assigned
				if (dev_state.SiteId.IsNull() || dev_state.SiteId.IsUnknown() || len(dev_state.SiteId.ValueString()) == 0) &&
					!(dev_plan.SiteId.IsNull() || dev_plan.SiteId.IsUnknown() || len(dev_plan.SiteId.ValueString()) == 0) {
					op = "assign"
					device_mac = dev_state.Mac.ValueString()
					// Site ID in state but not planed > must be unassigned
				} else if !(dev_state.SiteId.IsNull() || dev_state.SiteId.IsUnknown() || len(dev_state.SiteId.ValueString()) == 0) &&
					(dev_plan.SiteId.IsNull() || dev_plan.SiteId.IsUnknown() || len(dev_plan.SiteId.ValueString()) == 0) {
					op = "unassign"
					device_mac = dev_state.Mac.ValueString()
					// State Site ID != Plan Site ID > must be reassigned
				} else if !dev_state.SiteId.Equal(dev_plan.SiteId) {
					device_mac = dev_state.Mac.ValueString()
					op = "reassign"
				}
			}
		}
		if !already_claimed {
			claim_devices = append(claim_devices, dev_plan.Magic.ValueString())
			if !dev_plan.SiteId.IsNull() && !dev_plan.SiteId.IsUnknown() {
				tflog.Debug(ctx, "Device "+dev_plan.Magic.ValueString()+" will be claimed then assigned to "+dev_plan.SiteId.ValueString())
				asign_claimed_devices[dev_plan.Magic.ValueString()] = dev_plan.SiteId.ValueString()
			}
		}
		switch op {
		case "assign":
			assign_devices[dev_plan.SiteId.ValueString()] = append(assign_devices[dev_plan.SiteId.ValueString()], device_mac)
		case "reassign":
			assign_devices[dev_plan.SiteId.ValueString()] = append(assign_devices[dev_plan.SiteId.ValueString()], device_mac)
		case "unassign":
			unassign_devices = append(unassign_devices, device_mac)
		}
	}

	for _, dev_state_attr := range state.Devices.Elements() {
		var dsi interface{} = dev_state_attr
		var dev_state = dsi.(resource_org_inventory.DevicesValue)
		to_unclaim := true
		for _, dev_plan_attr := range plan.Devices.Elements() {
			var dpi interface{} = dev_plan_attr
			var dev_plan = dpi.(resource_org_inventory.DevicesValue)
			if dev_state.Magic.Equal(dev_plan.Magic) {
				to_unclaim = false
			}
		}
		if to_unclaim {
			unclaim_devices = append(unclaim_devices, dev_state.Serial.ValueString())
		}
	}

	/////////////////////// CLAIM
	if len(claim_devices) > 0 {
		tflog.Info(ctx, "Starting to Claim devices")
		claim_response, err := r.client.OrgsInventory().AddOrgInventory(ctx, orgId, claim_devices)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Claiming Devices to the Org Inventory",
				"Could not Claim devices, unexpected error: "+err.Error(),
			)
			return
		}
		inventory_added := claim_response.Data.InventoryAdded
		for _, v := range inventory_added {
			site_id, ok := asign_claimed_devices[v.Magic]
			if ok {
				assign_devices[site_id] = append(assign_devices[site_id], v.Mac)
			}

		}
	}
	/////////////////////// UNCLAIM
	if len(unclaim_devices) > 0 {
		tflog.Info(ctx, "Starting to Unclaim devices: ", map[string]interface{}{"serials": strings.Join(unclaim_devices, ", ")})

		unclaim_body := models.InventoryUpdate{}
		unclaim_body.Op = models.InventoryUpdateOperationEnum_DELETE
		unclaim_body.Serials = unclaim_devices
		unclaim_response, err := r.client.OrgsInventory().UpdateOrgInventoryAssignment(ctx, orgId, &unclaim_body)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Unclaiming Devices from the Org Inventory",
				"Could not Unclaim devices, unexpected error: "+err.Error(),
			)
			return
		}
		tflog.Info(ctx, "response for API Call to unclaim devices:", map[string]interface{}{
			"Error":   strings.Join(unclaim_response.Data.Error, ", "),
			"Reason":  strings.Join(unclaim_response.Data.Reason, ", "),
			"Success": strings.Join(unclaim_response.Data.Success, ", "),
		})
	}
	/////////////////////// UNASSIGN
	if len(unassign_devices) > 0 {
		tflog.Info(ctx, "Starting to Unassign devices: ", map[string]interface{}{"macs": strings.Join(unassign_devices, ", ")})

		unassign_body := models.InventoryUpdate{}
		unassign_body.Op = models.InventoryUpdateOperationEnum_UNASSIGN
		unassign_body.Macs = unassign_devices
		unassign_response, err := r.client.OrgsInventory().UpdateOrgInventoryAssignment(ctx, orgId, &unassign_body)
		tflog.Info(ctx, "response for API Call to claim devices:", map[string]interface{}{
			"Error":   strings.Join(unassign_response.Data.Error, ", "),
			"Reason":  strings.Join(unassign_response.Data.Reason, ", "),
			"Success": strings.Join(unassign_response.Data.Success, ", "),
		})

		if err != nil {
			resp.Diagnostics.AddError(
				"Error Unassigning Devices from the Org Inventory",
				"Could not Unassign devices, unexpected error: "+err.Error(),
			)
			return
		}
	}
	/////////////////////// ASSIGN
	if len(assign_devices) > 0 {
		for k, v := range assign_devices {
			tflog.Info(ctx, "Starting to Assign devices to site "+k+": ", map[string]interface{}{"macs": strings.Join(v, ", ")})

			assign_body := models.InventoryUpdate{}
			assign_body.Op = models.InventoryUpdateOperationEnum_ASSIGN
			assign_body.Macs = unassign_devices
			assign_body.SiteId = models.ToPointer(uuid.MustParse(k))

			assign_response, err := r.client.OrgsInventory().UpdateOrgInventoryAssignment(ctx, orgId, &assign_body)
			if err != nil {
				resp.Diagnostics.AddError(
					"Error Assigning Devices to the Org Inventory",
					"Could not Assign devices, unexpected error: "+err.Error(),
				)
				return
			}
			tflog.Info(ctx, "response for API Call to assign devices:", map[string]interface{}{
				"Error":   strings.Join(assign_response.Data.Error, ", "),
				"Reason":  strings.Join(assign_response.Data.Reason, ", "),
				"Success": strings.Join(assign_response.Data.Success, ", "),
			})
		}
	}
	/////////////////////// Check
	tflog.Info(ctx, "Starting Inventory state refresh: org_id  "+state.OrgId.ValueString())
	var serial string
	var model string
	var mType models.DeviceTypeEnum
	var mac string
	var siteId string
	var vcMac string
	var vc bool = true
	var unassigned bool
	var limit int = 1000
	var page int
	tflog.Info(ctx, "Starting Inventory Read: org_id  "+state.OrgId.ValueString())
	data, err := r.client.OrgsInventory().GetOrgInventory(ctx, orgId, &serial, &model, &mType, &mac, &siteId, &vcMac, &vc, &unassigned, &limit, &page)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error refreshing Inventory",
			"Could not get Inventory, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_inventory.SdkToTerraform(ctx, state.OrgId.ValueString(), data.Data)
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

func (r *orgInventoryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_inventory.OrgInventoryModel
	tflog.Info(ctx, "Starting Inventory Delete: wxtag_id ")

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var serials []string
	for _, v := range state.Devices.Elements() {
		var vi interface{} = v
		var dev_state = vi.(resource_org_inventory.DevicesValue)
		serials = append(serials, dev_state.Serial.ValueString())
	}

	orgId := uuid.MustParse(state.OrgId.ValueString())
	unclaim_body := models.InventoryUpdate{}
	unclaim_body.Op = models.InventoryUpdateOperationEnum_DELETE
	unclaim_body.Serials = serials
	unclaim_response, err := r.client.OrgsInventory().UpdateOrgInventoryAssignment(ctx, orgId, &unclaim_body)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Unclaiming Devices from the Org Inventory",
			"Could not Unclaim devices, unexpected error: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, "response for API Call to unclaim devices:", map[string]interface{}{
		"Error":   strings.Join(unclaim_response.Data.Error, ", "),
		"Reason":  strings.Join(unclaim_response.Data.Reason, ", "),
		"Success": strings.Join(unclaim_response.Data.Success, ", "),
	})
}
