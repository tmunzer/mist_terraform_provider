# __generated__ by Terraform
# Please review these resources and move them into your main configuration files.

# __generated__ by Terraform
resource "mist_org_deviceprofile_assign" "test" {
  deviceprofile_id = "d53575f0-253e-42cd-9674-5c0523fff3ca"
  macs             = null
  org_id           = "992bf4b9-c900-4850-9992-107b2f9df928"
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

# __generated__ by Terraform from "992bf4b9-c900-4850-9992-107b2f9df928.e24ff180-2646-4e0c-84f0-2e594ef2f17c"
resource "mist_org_wxtag" "aaa" {
  mac     = null
  match   = "radius_username"
  name    = "AAA Attribute"
  op      = "in"
  org_id  = "992bf4b9-c900-4850-9992-107b2f9df928"
  specs   = null
  type    = "match"
  values  = ["test_val"]
  vlan_id = null
}

# __generated__ by Terraform from "992bf4b9-c900-4850-9992-107b2f9df928.97ae31f9-6f50-413f-ae4d-80a92d068c58"
resource "mist_org_rftemplate" "test" {
  ant_gain_24 = 0
  ant_gain_5  = 0
  ant_gain_6  = 0
  band_24 = {
    allow_rrm_disable = false
    ant_gain          = 0
    antenna_mode      = null
    bandwidth         = 20
    channels          = []
    disabled          = false
    power             = null
    power_max         = 18
    power_min         = 8
    preamble          = "short"
  }
  band_24_usage = jsonencode(24)
  band_5 = {
    allow_rrm_disable = false
    ant_gain          = 0
    antenna_mode      = null
    bandwidth         = 40
    channels          = []
    disabled          = false
    power             = null
    power_max         = 17
    power_min         = 8
    preamble          = "short"
  }
  band_5_on_24_radio = null
  band_6 = {
    allow_rrm_disable = false
    ant_gain          = 0
    antenna_mode      = null
    bandwidth         = 80
    channels          = []
    disabled          = false
    power             = null
    power_max         = 17
    power_min         = 8
    preamble          = "short"
    standard_power    = null
  }
  country_code     = null
  model_specific   = null
  name             = "test"
  org_id           = "992bf4b9-c900-4850-9992-107b2f9df928"
  scanning_enabled = null
}

# __generated__ by Terraform from "992bf4b9-c900-4850-9992-107b2f9df928"
resource "mist_org_inventory" "inventory" {
  devices = [
    {
      claim_code = "CV4YAS8DQWYLL6M"
      mac        = "e8a24550e732"
      site_id    = "19e06ed0-9ee9-4ac7-9d5c-dd1ed551a95c"
    },
    {
      claim_code = "CPKL2EXN8JY98AC"
      mac        = "003e7316ff9e"
      site_id    = "19e06ed0-9ee9-4ac7-9d5c-dd1ed551a95c"
    },
    {
      claim_code = "G87JHBFXZJSFNMX"
      mac        = "2c21311c37b0"
      site_id    = null
    },
  ]
  org_id = "992bf4b9-c900-4850-9992-107b2f9df928"
}
