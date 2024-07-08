package resource_org

import (
	"mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(plan *OrgModel) (*models.Org, diag.Diagnostics) {
	var diags diag.Diagnostics

	var alarmtemplate_id uuid.UUID
	if !plan.AlarmtemplateId.IsNull() && !plan.AlarmtemplateId.IsUnknown() {
		alarmtemplate_id, _ = uuid.Parse(plan.AlarmtemplateId.ValueString())
	}
	data := models.Org{
		AlarmtemplateId: models.NewOptional(models.ToPointer(alarmtemplate_id)),
		AllowMist:       models.ToPointer(plan.AllowMist.ValueBool()),
		Name:            plan.Name.ValueString(),
		SessionExpiry:   models.ToPointer(int(plan.SessionExpiry.ValueInt64())),
	}

	return &data, diags
}
