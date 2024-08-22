# __generated__ by Terraform
# Please review these resources and move them into your main configuration files.

# __generated__ by Terraform from "992bf4b9-c900-4850-9992-107b2f9df928"
resource "mist_org_inventory" "inventory" {
  devices = [
  ]
  org_id = "992bf4b9-c900-4850-9992-107b2f9df928"
}

# __generated__ by Terraform from "992bf4b9-c900-4850-9992-107b2f9df928"
resource "mist_org" "terraform_test" {
  alarmtemplate_id = null
  allow_mist       = true
  msp_logo_url     = null
  msp_name         = null
  name             = "Terraform Testing"
  session_expiry   = 1440
}

# __generated__ by Terraform from "be11a751-4b5e-46f6-a867-023763df369c"
resource "mist_site" "terraform_site" {
  address            = "77 Terrasse de l'Université, 92000 Nanterre, France"
  alarmtemplate_id   = null
  aptemplate_id      = null
  country_code       = "FR"
  gatewaytemplate_id = "cb545f33-254f-4a86-9d73-c067faba58a2"
  latlng = {
    lat = 48.899268
    lng = 2.214447
  }
  name               = "terraform_site"
  networktemplate_id = "ee8095cd-4eb2-4a25-a37d-22f601dbbb71"
  notes              = "Created with Terraform, Updated with Terraform"
  org_id             = "992bf4b9-c900-4850-9992-107b2f9df928"
  rftemplate_id      = "7bdfd85e-5a87-4d84-bb08-f5fa79ab603d"
  secpolicy_id       = null
  sitegroup_ids      = ["f53ec5f4-5115-4e5b-a3f1-2a3ffd5bab53", "0e6350fc-c35a-4c02-939f-bb87fa453d55"]
  sitetemplate_id    = null
  timezone           = "Europe/Paris"
}
