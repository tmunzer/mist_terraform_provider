---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "mist_org_inventory Resource - terraform-provider-mist"
subcategory: ""
description: |-
  
---

# mist_org_inventory (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `org_id` (String)

### Optional

- `devices` (Attributes List) (see [below for nested schema](#nestedatt--devices))

<a id="nestedatt--devices"></a>
### Nested Schema for `devices`

Required:

- `magic` (String)

Optional:

- `org_id` (String)
- `site_id` (String) site id if assigned, null if not assigned

Read-Only:

- `device_id` (String) Mist Device ID
- `hostname` (String) Device Hostname
- `mac` (String) MAC address
- `model` (String) device model
- `serial` (String) device serial
- `type` (String)
- `vc_mac` (String) Virtual Chassis MAC Address