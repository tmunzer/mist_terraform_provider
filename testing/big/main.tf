terraform {
  required_providers {
    mist = {
      source = "registry.terraform.io/juniper/mist"
    }
  }
}

provider "mist" {
  host     = local.envs["HOST"]
  apitoken = local.envs["APITOKEN"]
}

### ORG
resource "mist_org" "terraform_test" {
  name = "Debug_TF_TM"
}

resource "mist_org_network" "network_vlan_11" {
  org_id    = mist_org.terraform_test.id
  vlan_id   = 1
  subnet    = "{{Class_B}}{{Class_C_Vlan_1}}{{CIDR}}"
  gateway   = "{{Class_B}}{{Class_C_Vlan_1}}{{Gateway}}"
  isolation = false
  internet_access = {
    enabled    = true
    restricted = false
  }
  disallow_mist_services = false
  name                   = "vlan_1"
}
resource "mist_org_network" "network_Test_vlan32" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3
  subnet                 = "3.3.3.0/30"
  disallow_mist_services = true
  name                   = "Test_vlan3"
}
resource "mist_org_network" "network_SC_TABLETS_VL2013" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 201
  subnet                 = "{{sc_tablets}}/{{sc_tablets_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_TABLETS_VL201"
}
resource "mist_org_network" "network_GUEST4" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 72
  subnet                 = "192.168.254.0/24"
  disallow_mist_services = false
  name                   = "GUEST"
}
resource "mist_org_network" "network_CTR_WIRED_TABLETS_VL215" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 21
  subnet                 = "{{ctr_wired_tablets}}/{{ctr_wired_tablets_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_WIRED-TABLETS_VL21"
}
resource "mist_org_network" "network_PC_RF_GUNS6" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3
  subnet                 = "{{pc_rf_guns}}/{{pc_rf_guns_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_RF-GUNS"
}
resource "mist_org_network" "network_PC_MGMT7" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "{{pc_default}}/{{pc_default_cidr}}"
  disallow_mist_services = false
  tenants = {
    PC_Lab_Mgmt = {
      addresses = [
        "10.167.186.0/25"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_MGMT"
}
resource "mist_org_network" "network_CTR_WIRED_CFD8" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 21
  subnet                 = "{{ctr_wired_cfd}}/{{ctr_wired_cfd_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_WIRED-CFD"
}
resource "mist_org_network" "network_CTR_MCAFEE_SERVERS_109" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "10.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    CTR_MCAFEE_SERVER1 = {
      addresses = [
        "10.126.57.98/32"
      ]
    }
    CTR_MCAFEE_SERVER2 = {
      addresses = [
        "10.126.57.100/32"
      ]
    }
    CTR_MCAFEE_SERVER3 = {
      addresses = [
        "10.126.57.101/32"
      ]
    }
    CTR_MCAFEE_SERVER4 = {
      addresses = [
        "10.126.57.122/32"
      ]
    }
    CTR_MCAFEE_SERVER5 = {
      addresses = [
        "10.126.57.123/32"
      ]
    }
    CTR_MCAFEE_SERVER6 = {
      addresses = [
        "10.2.38.0/23"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_MCAFEE_SERVERS_10"
}
resource "mist_org_network" "network_IOT10" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 290
  subnet                 = "192.168.255.0/24"
  disallow_mist_services = true
  name                   = "IOT"
}
resource "mist_org_network" "network_PS_THIRD_PARTY_VL70011" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 700
  subnet                 = "{{ps_third_party}}/{{ps_third_party_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PS_THIRD-PARTY_VL700"
}
resource "mist_org_network" "network_NEXPOSE_1012" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "10.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    CTC_NEXPOSE1 = {
      addresses = [
        "10.2.36.21/32",
        "10.2.36.29/32",
        "10.127.20.201/32",
        "10.18.27.135/32",
        "10.18.27.136/32",
        "10.100.249.30/32",
        "10.100.249.31/32",
        "10.100.249.32/32",
        "10.100.249.33/32",
        "10.100.249.34/32",
        "10.100.249.35/32",
        "10.100.249.36/32",
        "10.100.249.37/32",
        "10.127.26.200/32",
        "10.127.26.201/32",
        "10.127.26.202/32",
        "10.127.26.203/32",
        "10.127.26.204/32",
        "10.127.26.205/32",
        "10.127.26.206/32",
        "10.127.26.207/32",
        "10.127.26.208/32",
        "10.127.26.209/32",
        "10.127.26.210/32",
        "10.127.26.211/32",
        "10.127.26.212/32",
        "10.127.26.224/32",
        "10.127.26.225/32",
        "10.127.26.226/32",
        "10.127.26.227/32",
        "10.127.26.228/32",
        "10.127.26.229/32",
        "10.127.26.230/32",
        "10.127.26.231/32",
        "10.127.26.232/32",
        "10.127.26.233/32"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "NEXPOSE-10"
}
resource "mist_org_network" "network_PC_SRV_MGMT13" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 5
  subnet                 = "{{pc_srv_mgmt}}/{{pc_srv_mgmt_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_SRV-MGMT"
}
resource "mist_org_network" "network_PC_D_VOIP14" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 63
  subnet                 = "{{pc_d_voip}}/{{pc_d_voip_cidr}}"
  disallow_mist_services = true
  name                   = "PC_D-VOIP"
}
resource "mist_org_network" "network_CTR_EFD_VOIP15" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 31
  subnet                 = "{{ctr_efd_voip}}/{{ctr_efd_voip_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_EFD-VOIP"
}
resource "mist_org_network" "network_CTR_RF_GUNS16" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3
  subnet                 = "{{ctr_rf_guns}}/{{ctr_rf_guns_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_RF-GUNS"
}
resource "mist_org_network" "network_MWW_VOIP17" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 500
  subnet                 = "{{mww_voip}}/{{mww_voip_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "MWW_VOIP"
}
resource "mist_org_network" "network_standalone_320_edge_vl18" {
  org_id                 = mist_org.terraform_test.id
  vlan_id                = 1
  subnet                 = "{{Class_B}}{{Class_C_Vlan_1}}{{CIDR}}/24"
  gateway                = "{{Class_B}}{{Class_C_Vlan_1}}{{Gateway}}"
  isolation              = true
  disallow_mist_services = false
  name                   = "standalone_320_edge_vlan_1"
}
resource "mist_org_network" "network_AJ_HUB_QFX_lo019" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "4.4.4.4/32"
  disallow_mist_services = true
  name                   = "AJ-HUB-QFX-lo0"
}
resource "mist_org_network" "network_SC_WORKSTATIONS_VL10120" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 101
  subnet                 = "{{sc_workstations}}/{{sc_workstations_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_WORKSTATIONS_VL101"
}
resource "mist_org_network" "network_SC_PRINTERS_VL10221" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 102
  subnet                 = "{{sc_printers}}/{{sc_printers_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_PRINTERS_VL102"
}
resource "mist_org_network" "network_CTR_EFD_VOIP_VL3122" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 31
  subnet                 = "{{ctr_efd_voip}}/{{ctr_efd_voip_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_EFD-VOIP_VL31"
}
resource "mist_org_network" "network_PC_WIRED_CFD23" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 21
  subnet                 = "{{pc_wired_cfd}}/{{pc_wired_cfd_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_WIRED-CFD"
}
resource "mist_org_network" "network_PC_D_WIRED24" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 6
  subnet                 = "{{pc_d_wired}}/{{pc_d_wired_cidr}}"
  disallow_mist_services = true
  name                   = "PC_D-WIRED"
}
resource "mist_org_network" "network_CTR_M_PRINT25" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 32
  subnet                 = "{{ctr_m_print}}/{{ctr_m_print_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_M-PRINT"
}
resource "mist_org_network" "network_CTR_GWiFi26" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 202
  subnet                 = "{{ctr_gwifi}}/{{ctr_gwifi_cidr}}"
  disallow_mist_services = true
  name                   = "CTR_GWiFi"
}
resource "mist_org_network" "network_MWW_PRINTERS27" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 102
  subnet                 = "{{mww_printers}}/{{mww_printers_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "MWW_PRINTERS"
}
resource "mist_org_network" "network_SC_PRINTERS28" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 102
  subnet                 = "{{sc_printers}}/{{sc_printers_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_PRINTERS"
}
resource "mist_org_network" "network_PS_WORKSTATIONS29" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 101
  subnet                 = "{{ps_workstations}}/{{ps_workstations_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PS_WORKSTATIONS"
}
resource "mist_org_network" "network_SE_F_WIRED30" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 600
  subnet                 = "{{se_f_wired}}/{{se_f_wired_cidr}}"
  disallow_mist_services = true
  name                   = "SE_F-WIRED"
}
resource "mist_org_network" "network_MGMT_vlan131" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "{{Class_B_VLAN1}}{{Class_C_VLAN1}}0/{{CIDR_25}}"
  disallow_mist_services = false
  name                   = "MGMT_vlan1"
}
resource "mist_org_network" "network_SC_THIRD_PARTY_VL70032" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 700
  subnet                 = "{{sc_third_party}}/{{sc_third_party_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_THIRD-PARTY_VL700"
}
resource "mist_org_network" "network_SC_VOIP_VL50033" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 500
  subnet                 = "{{sc_voip}}/{{sc_voip_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_VOIP_VL500"
}
resource "mist_org_network" "network_aj_prod_vlan300_dc34" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 300
  subnet                 = "10.125.2.224/29"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "aj-prod-vlan300-dc"
}
resource "mist_org_network" "network_AT_TABLETS_VL20135" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 201
  subnet                 = "{{at_tablets}}/{{at_tablets_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "AT_TABLETS_VL201"
}
resource "mist_org_network" "network_PC_EFD_VOIP36" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 31
  subnet                 = "{{pc_efd_voip}}/{{pc_efd_voip_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_EFD-VOIP"
}
resource "mist_org_network" "network_PC_D_WIRELESS37" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 61
  subnet                 = "{{pc_d_wireless}}/{{pc_d_wireless_cidr}}"
  disallow_mist_services = true
  name                   = "PC_D-WIRELESS"
}
resource "mist_org_network" "network_CTR_FRONTIER38" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 51
  subnet                 = "{{ctr_frontier}}/{{ctr_frontier_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_FRONTIER"
}
resource "mist_org_network" "network_MWW_POS39" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 100
  subnet                 = "{{mww_pos}}/{{mww_pos_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "MWW_POS"
}
resource "mist_org_network" "network_VLAN30040" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 300
  subnet                 = "{{Class_B}}{{Class_C_Vlan_300}}0/{{CIDR}}"
  disallow_mist_services = true
  vpn_access = {
    aj-hub1 = {
      routed = true
    }
  }
  name = "VLAN300"
}
resource "mist_org_network" "network_aj_prod_vlan400_mgmt41" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 400
  subnet                 = "10.161.112.0/27"
  disallow_mist_services = false
  tenants = {
    HUB_VL400_MGMT_Subnet = {
      addresses = [
        "10.161.112.0/27"
      ]
    }
  }
  name = "aj-prod-vlan400-mgmt"
}
resource "mist_org_network" "network_SC_MGMT_VL142" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "{{sc_default}}/{{sc_default_cidr}}"
  disallow_mist_services = false
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_MGMT_VL1"
}
resource "mist_org_network" "network_SC_CORP_WIRELESS_VL20043" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 200
  subnet                 = "{{sc_corp_wireless}}/{{sc_corp_wireless_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_CORP-WIRELESS_VL200"
}
resource "mist_org_network" "network_AT_WORKSTATIONS_VL10144" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 101
  subnet                 = "{{at_workstations}}/{{at_workstations_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "AT_WORKSTATIONS_VL101"
}
resource "mist_org_network" "network_PC_FRONTIER45" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 51
  subnet                 = "{{pc_frontier}}/{{pc_frontier_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_FRONTIER"
}
resource "mist_org_network" "network_CTR_3P_DIGITAL46" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 7
  subnet                 = "{{ctr_3p_digital}}/{{ctr_3p_digital_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_3P-DIGITAL"
}
resource "mist_org_network" "network_AT_VOIP_VL50047" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 500
  subnet                 = "{{at_voip}}/{{at_voip_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "AT_VOIP_VL500"
}
resource "mist_org_network" "network_VLAN40048" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "192.168.140.0/24"
  disallow_mist_services = false
  vlan_id                = 400
  name                   = "VLAN400"
}
resource "mist_org_network" "network_CTR_MCAFEE_SERVERS_4249" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "42.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    CTR_MCAFEE_SERVER7 = {
      addresses = [
        "42.108.33.88/32"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_MCAFEE_SERVERS_42"
}
resource "mist_org_network" "network_Test_vlan250" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 2
  subnet                 = "2.2.2.0/30"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "Test_vlan2"
}
resource "mist_org_network" "network_SC_DVR_VL30051" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 300
  subnet                 = "{{sc_dvr}}/{{sc_dvr_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_DVR_VL300"
}
resource "mist_org_network" "network_PC_POS52" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 4
  subnet                 = "{{pc_pos}}/{{pc_pos_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_POS"
}
resource "mist_org_network" "network_PC_ESL53" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 52
  subnet                 = "{{pc_esl}}/{{pc_esl_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_ESL"
}
resource "mist_org_network" "network_CTR_D_WIRED54" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 6
  subnet                 = "{{ctr_d_wired}}/{{ctr_d_wired_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_D-WIRED"
}
resource "mist_org_network" "network_CTR_SQUID55" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 53
  subnet                 = "{{ctr_squid}}/{{ctr_squid_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_SQUID"
}
resource "mist_org_network" "network_MWW_DVR56" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 300
  subnet                 = "{{mww_dvr}}/{{mww_dvr_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "MWW_DVR"
}
resource "mist_org_network" "network_SC_THIRD_PARTY57" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 700
  subnet                 = "{{sc_third_party}}/{{sc_third_party_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_THIRD-PARTY"
}
resource "mist_org_network" "network_PHL_MGMT58" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "{{phl_default}}/{{phl_default_cidr}}"
  disallow_mist_services = false
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PHL_MGMT"
}
resource "mist_org_network" "network_vlan_210059" {
  org_id    = mist_org.terraform_test.id
  vlan_id   = 2100
  subnet    = "{{vlan_2100_subnet}}/29"
  gateway   = "{{vlan_2100_ip}}"
  isolation = true
  internet_access = {
    destination_nat = {
      "{{Edge_Public_IP}}:500" = {
        name        = "Mist_Edge_IKE"
        internal_ip = "{{Edge_Tunnel_IP}}"
        port        = "500"
      }
      "{{Edge_Public_IP}}:4500" = {
        name        = "Mist_Edge_IPSEC"
        internal_ip = "{{Edge_Tunnel_IP}}"
        port        = "4500"
      }
    }
  }
  disallow_mist_services = false
  name                   = "vlan_2100"
}
resource "mist_org_network" "network_10_235_223_0_2460" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "10.235.223.0/24"
  vpn_access = {
    OrgOverlay = {
      routed                    = true
      no_readvertise_to_overlay = false
      no_readvertise_to_lan_bgp = false
    }
  }
  disallow_mist_services = true
  name                   = "10_235_223_0_24"
}
resource "mist_org_network" "network_SC_DIGITAL_VL40061" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 400
  subnet                 = "{{sc_d_media}}/{{sc_d_media_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_DIGITAL_VL400"
}
resource "mist_org_network" "network_AT_POS_VL10062" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 100
  subnet                 = "{{at_pos}}/{{at_pos_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "AT_POS_VL100"
}
resource "mist_org_network" "network_SC_POS_VL10063" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 100
  subnet                 = "{{sc_pos}}/{{sc_pos_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_POS_VL100"
}
resource "mist_org_network" "network_wdc_prod_vlan300_dc64" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 300
  subnet                 = "10.16.2.224/29"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "wdc-prod-vlan300-dc"
}
resource "mist_org_network" "network_SC_GWiFi_VL20265" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 202
  subnet                 = "{{sc_gwifi}}/{{sc_gwifi_cidr}}"
  disallow_mist_services = true
  name                   = "SC_GWiFi_VL202"
}
resource "mist_org_network" "network_PC_CORP_WIRED66" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 2
  subnet                 = "{{pc_corp_wired}}/{{pc_corp_wired_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_CORP-WIRED"
}
resource "mist_org_network" "network_PC_3P_DIGITAL67" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 7
  subnet                 = "{{pc_3p_digital}}/{{pc_3p_digital_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_3P-DIGITAL"
}
resource "mist_org_network" "network_PC_SQUID68" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 53
  subnet                 = "{{pc_squid}}/{{pc_squid_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_SQUID"
}
resource "mist_org_network" "network_CTR_SRV_MGMT69" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 5
  subnet                 = "{{ctr_srv_mgmt}}/{{ctr_srv_mgmt_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_SRV-MGMT"
}
resource "mist_org_network" "network_CTR_CORP_WIRED70" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 2
  subnet                 = "{{ctr_corp_wired}}/{{ctr_corp_wired_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_CORP-WIRED"
}
resource "mist_org_network" "network_PC_GWiFi71" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 202
  subnet                 = "{{pc_gwifi}}/{{pc_gwifi_cidr}}"
  disallow_mist_services = false
  name                   = "PC_GWiFi"
}
resource "mist_org_network" "network_CTR_ESL72" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 52
  subnet                 = "{{ctr_esl}}/{{ctr_esl_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_ESL"
}
resource "mist_org_network" "network_MWW_TABLETS73" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 201
  subnet                 = "{{mww_tablets}}/{{mww_tablets_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "MWW_TABLETS"
}
resource "mist_org_network" "network_SC_POS74" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 100
  subnet                 = "{{sc_pos}}/{{sc_pos_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_POS"
}
resource "mist_org_network" "network_PHL_THIRD_PARTY75" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 700
  subnet                 = "{{phl_third_party}}/{{phl_third_party_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PHL_THIRD-PARTY"
}
resource "mist_org_network" "network_CTR_POS76" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 4
  subnet                 = "{{ctr_pos}}/{{ctr_pos_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_POS"
}
resource "mist_org_network" "network_MWW_WORKSTATIONS77" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 101
  subnet                 = "{{mww_workstations}}/{{mww_workstations_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "MWW_WORKSTATIONS"
}
resource "mist_org_network" "network_CTR_MGMT78" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "{{ctr_default}}/{{ctr_default_cidr}}"
  disallow_mist_services = false
  tenants = {
    CTR_Lab_Mgmt = {
      addresses = [
        "10.167.230.128/25"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_MGMT"
}
resource "mist_org_network" "network_PS_POS79" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 100
  subnet                 = "{{ps_pos}}/{{ps_pos_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PS_POS"
}
resource "mist_org_network" "network_SE_CORP_WIRELESS80" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 200
  subnet                 = "{{se_corp_wireless}}/{{se_corp_wireless_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SE_CORP-WIRELESS"
}
resource "mist_org_network" "network_AT_GWiFi_VL20281" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 202
  subnet                 = "{{at_gwifi}}/{{at_gwifi_cidr}}"
  disallow_mist_services = true
  name                   = "AT_GWiFi_VL202"
}
resource "mist_org_network" "network_TRIPWIRE_17282" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "172.16.0.0/20"
  disallow_mist_services = true
  tenants = {
    CTC_TRIPWIRE2 = {
      addresses = [
        "172.22.33.124/32",
        "172.22.33.125/32",
        "172.22.34.108/32"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "TRIPWIRE-172"
}
resource "mist_org_network" "network_MWW_TABLETS_VL20183" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 201
  subnet                 = "{{mww_tablets}}/{{mww_tablets_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "MWW_TABLETS_VL201"
}
resource "mist_org_network" "network_172_16_0_0_2084" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "172.16.0.0/20"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "172_16_0_0_20"
}
resource "mist_org_network" "network_MWW_GWiFi85" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 202
  subnet                 = "{{mww_gwifi}}/{{mww_gwifi_cidr}}"
  disallow_mist_services = true
  name                   = "MWW_GWiFi"
}
resource "mist_org_network" "network_SC_TABLETS86" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 201
  subnet                 = "{{sc_tablets}}/{{sc_tablets_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_TABLETS"
}
resource "mist_org_network" "network_PHL_POS87" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 100
  subnet                 = "{{phl_pos}}/{{phl_pos_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PHL_POS"
}
resource "mist_org_network" "network_wdc_prod_ctb_rtl88" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 124
  subnet                 = "10.161.254.8/29"
  disallow_mist_services = true
  name                   = "wdc-prod-ctb-rtl"
}
resource "mist_org_network" "network_CTP_PUMP89" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 41
  subnet                 = "{{ctp_pump}}/{{ctp_pump_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTP_PUMP"
}
resource "mist_org_network" "network_CTR_FRONTIER_VL5190" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 51
  subnet                 = "{{ctr_frontier}}/{{ctr_frontier_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_FRONTIER_VL51"
}
resource "mist_org_network" "network_PC_ESL_VL5291" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 52
  subnet                 = "{{pc_esl}}/{{pc_esl_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_ESL_VL52"
}
resource "mist_org_network" "network_MWW_THERMOSTATS92" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 701
  subnet                 = "{{mww_thermostats}}/{{mww_thermostats_cidr}}"
  disallow_mist_services = true
  name                   = "MWW_THERMOSTATS"
}
resource "mist_org_network" "network_SC_MGMT93" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "{{sc_default}}/{{sc_default_cidr}}"
  disallow_mist_services = false
  tenants = {
    SC_Lab_Mgmt = {
      addresses = [
        "10.164.255.0/26"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_MGMT"
}
resource "mist_org_network" "network_MWW_THIRD_PARTY94" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 700
  subnet                 = "{{mww_third_party}}/{{mww_third_party_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "MWW_THIRD-PARTY"
}
resource "mist_org_network" "network_SC_GWiFi95" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 202
  subnet                 = "{{sc_gwifi}}/{{sc_gwifi_cidr}}"
  disallow_mist_services = true
  name                   = "SC_GWiFi"
}
resource "mist_org_network" "network_PS_MGMT96" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "{{ps_default}}/{{ps_default_cidr}}"
  disallow_mist_services = false
  tenants = {
    PS_Lab_Mgmt = {
      addresses = [
        "10.168.194.192/26"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PS_MGMT"
}
resource "mist_org_network" "network_CTP_MGMT97" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "{{ctp_default}}/{{ctp_default_cidr}}"
  disallow_mist_services = false
  tenants = {
    CTP_Lab_Mgmt = {
      addresses = [
        "10.186.119.0/26"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTP_MGMT"
}
resource "mist_org_network" "network_PC_CORP_WIRED_VL298" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 2
  subnet                 = "{{pc_corp_wired}}/{{pc_corp_wired_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_CORP-WIRED_VL2"
}
resource "mist_org_network" "network_CTR_RF_GUNS_VL399" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3
  subnet                 = "{{ctr_rf_guns}}/{{ctr_rf_guns_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_RF-GUNS_VL3"
}
resource "mist_org_network" "network_MWW_DIGITAL_VL400100" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 400
  subnet                 = "{{mww_digital_media}}/{{mww_digital_media_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "MWW_DIGITAL_VL400"
}
resource "mist_org_network" "network_CTP_MGMT_VL1101" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "{{ctp_default}}/{{ctp_default_cidr}}"
  disallow_mist_services = false
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTP_MGMT_VL1"
}
resource "mist_org_network" "network_PS_MGMT_VL1102" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "{{ps_default}}/{{ps_default_cidr}}"
  disallow_mist_services = false
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PS_MGMT_VL1"
}
resource "mist_org_network" "network_MWW_CORP_WIRELESS103" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 200
  subnet                 = "{{mww_corp_wireless}}/{{mww_corp_wireless_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "MWW_CORP-WIRELESS"
}
resource "mist_org_network" "network_PHL_GWiFi104" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 202
  subnet                 = "{{phl_gwifi}}/{{phl_gwifi_cidr}}"
  disallow_mist_services = true
  name                   = "PHL_GWiFi"
}
resource "mist_org_network" "network_PHL_WORKSTATIONS105" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 101
  subnet                 = "{{phl_workstations}}/{{phl_workstations_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PHL_WORKSTATIONS"
}
resource "mist_org_network" "network_CTP_POS106" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 100
  subnet                 = "{{ctp_pos}}/{{ctp_pos_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTP_POS"
}
resource "mist_org_network" "network_CTR_M_PRINT_VL32107" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 32
  subnet                 = "{{ctr_m_print}}/{{ctr_m_print_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_M-PRINT_VL32"
}
resource "mist_org_network" "network_MWW_DIGITAL_MEDIA108" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 400
  subnet                 = "{{mww_digital_media}}/{{mww_digital_media_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "MWW_DIGITAL-MEDIA"
}
resource "mist_org_network" "network_PHL_PRINTERS109" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 102
  subnet                 = "{{phl_printers}}/{{phl_printers_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PHL_PRINTERS"
}
resource "mist_org_network" "network_PS_CORP_WIRELESS110" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 200
  subnet                 = "{{ps_corp_wireless}}/{{ps_corp_wireless_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PS_CORP-WIRELESS"
}
resource "mist_org_network" "network_CTP_NON_CTC111" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 700
  subnet                 = "{{ctp_non_ctc}}/{{ctp_non_ctc_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTP_NON-CTC"
}
resource "mist_org_network" "network_PC_3P_DIGITAL_VL7112" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 7
  subnet                 = "{{pc_3p_digital}}/{{pc_3p_digital_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_3P-DIGITAL_VL7"
}
resource "mist_org_network" "network_CTR_POS_VL4113" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 4
  subnet                 = "{{ctr_pos}}/{{ctr_pos_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_POS_VL4"
}
resource "mist_org_network" "network_PC_M_PRINT_VL32114" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 32
  subnet                 = "{{pc_m_print}}/{{pc_m_print_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_M-PRINT_VL32"
}
resource "mist_org_network" "network_SC_WORKSTATIONS115" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 101
  subnet                 = "{{sc_workstations}}/{{sc_workstations_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_WORKSTATIONS"
}
resource "mist_org_network" "network_PHL_CORP_WIRELESS116" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 200
  subnet                 = "{{phl_corp_wireless}}/{{phl_corp_wireless_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PHL_CORP-WIRELESS"
}
resource "mist_org_network" "network_SE_PRINTERS117" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 102
  subnet                 = "{{se_printers}}/{{se_printers_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SE_PRINTERS"
}
resource "mist_org_network" "network_AT_CORP_WIRELESS_VL20118" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 200
  subnet                 = "{{at_corp_wireless}}/{{at_corp_wireless_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "AT_CORP-WIRELESS_VL200"
}
resource "mist_org_network" "network_PC_SRV_MGMT_VL5119" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 5
  subnet                 = "{{pc_srv_mgmt}}/{{pc_srv_mgmt_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_SRV-MGMT_VL5"
}
resource "mist_org_network" "network_MWW_THIRD_PARTY_VL700120" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 700
  subnet                 = "{{mww_third_party}}/{{mww_third_party_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "MWW_THIRD-PARTY_VL700"
}
resource "mist_org_network" "network_192_168_0_0_16121" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "192.168.0.0/16"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "192_168_0_0_16"
}
resource "mist_org_network" "network_SC_VOIP122" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 500
  subnet                 = "{{sc_voip}}/{{sc_voip_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_VOIP"
}
resource "mist_org_network" "network_SE_SHOEWALL123" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 300
  subnet                 = "{{se_shoewall}}/{{se_shoewall_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SE_SHOEWALL"
}
resource "mist_org_network" "network_CTP_CORP_WIRED124" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 101
  subnet                 = "{{ctp_corp_wired}}/{{ctp_corp_wired_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTP_CORP-WIRED"
}
resource "mist_org_network" "network_CTR_SQUID_VL53125" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 53
  subnet                 = "{{ctr_squid}}/{{ctr_squid_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_SQUID_VL53"
}
resource "mist_org_network" "network_MWW_PRINTERS_VL102126" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 102
  subnet                 = "{{mww_printers}}/{{mww_printers_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "MWW_PRINTERS_VL102"
}
resource "mist_org_network" "network_PS_POS_VL100127" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 100
  subnet                 = "{{ps_pos}}/{{ps_pos_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PS_POS_VL100"
}
resource "mist_org_network" "network_PS_WORKSTATIONS_VL101128" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 101
  subnet                 = "{{ps_workstations}}/{{ps_workstations_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PS_WORKSTATIONS_VL101"
}
resource "mist_org_network" "network_CTR_SDM_SERVERS_10129" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "10.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    CTR_SDM_SERVER1 = {
      addresses = [
        "10.127.20.204/32"
      ]
    }
    CTR_SDM_SERVER2 = {
      addresses = [
        "10.127.20.44/32"
      ]
    }
    CTR_SDM_SERVER3 = {
      addresses = [
        "10.127.20.219/32"
      ]
    }
    CTR_SDM_SERVER4 = {
      addresses = [
        "10.120.56.44/32"
      ]
    }
    CTR_SDM_SERVER5 = {
      addresses = [
        "10.127.91.217/32"
      ]
    }
    CTR_SDM_SERVER6 = {
      addresses = [
        "10.127.91.190/32"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_SDM_SERVERS_10"
}
resource "mist_org_network" "network_SC_DIGITAL_CUSTOMER_I130" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 220
  subnet                 = "{{sc_d_customer_interactive}}/{{sc_d_customer_interactive_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_DIGITAL-CUSTOMER-INTERACTIVE"
}
resource "mist_org_network" "network_aj_prod_ctb_rtl131" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 125
  subnet                 = "10.161.255.8/29"
  disallow_mist_services = true
  name                   = "aj-prod-ctb-rtl"
}
resource "mist_org_network" "network_AT_PRINTERS_VL102132" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 102
  subnet                 = "{{at_printers}}/{{at_printers_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "AT_PRINTERS_VL102"
}
resource "mist_org_network" "network_CTR_ESL_VL52133" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 52
  subnet                 = "{{ctr_esl}}/{{ctr_esl_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_ESL_VL52"
}
resource "mist_org_network" "network_SC_DIGITAL_INFRASTRUC134" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 230
  subnet                 = "{{sc_d_infrastructure}}/{{sc_d_infrastructure_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_DIGITAL-INFRASTRUCTURE"
}
resource "mist_org_network" "network_PS_GWiFi135" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 202
  subnet                 = "{{ps_gwifi}}/{{ps_gwifi_cidr}}"
  disallow_mist_services = true
  name                   = "PS_GWiFi"
}
resource "mist_org_network" "network_SC_DIGITAL_MEDIA136" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 400
  subnet                 = "{{sc_d_media}}/{{sc_d_media_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_DIGITAL-MEDIA"
}
resource "mist_org_network" "network_SE_POS137" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 100
  subnet                 = "{{se_pos}}/{{se_pos_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SE_POS"
}
resource "mist_org_network" "network_MWW_MGMT138" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "{{mww_default}}/{{mww_default_cidr}}"
  disallow_mist_services = false
  tenants = {
    MWW_Lab_Mgmt = {
      addresses = [
        "10.157.255.0/26"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "MWW_MGMT"
}
resource "mist_org_network" "network_NEXPOSE_172139" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "172.16.0.0/20"
  disallow_mist_services = true
  tenants = {
    CTC_NEXPOSE2 = {
      addresses = [
        "172.31.98.50/32",
        "172.31.98.51/32",
        "172.22.3.20/32",
        "172.22.36.154/32",
        "172.25.36.65/32"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "NEXPOSE-172"
}
resource "mist_org_network" "network_42_0_0_0_8140" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "42.0.0.0/8"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "42_0_0_0_8"
}
resource "mist_org_network" "network_SC_CORP_WIRELESS141" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 200
  subnet                 = "{{sc_corp_wireless}}/{{sc_corp_wireless_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_CORP-WIRELESS"
}
resource "mist_org_network" "network_SC_SECURITY142" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 300
  subnet                 = "{{sc_security}}/{{sc_security_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SC_SECURITY"
}
resource "mist_org_network" "network_PS_VOIP143" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 500
  subnet                 = "{{ps_voip}}/{{ps_voip_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PS_VOIP"
}
resource "mist_org_network" "network_SE_F_WIRELESS144" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 601
  subnet                 = "{{se_f_wireless}}/{{se_f_wireless_cidr}}"
  disallow_mist_services = true
  name                   = "SE_F-WIRELESS"
}
resource "mist_org_network" "network_PS_THIRD_PARTY145" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 700
  subnet                 = "{{ps_third_party}}/{{ps_third_party_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PS_THIRD-PARTY"
}
resource "mist_org_network" "network_PHL_TABLETS146" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 201
  subnet                 = "{{phl_tablets}}/{{phl_tablets_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PHL_TABLETS"
}
resource "mist_org_network" "network_AT_MGMT_VL1147" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "{{at_default}}/{{at_default_cidr}}"
  disallow_mist_services = false
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "AT_MGMT_VL1"
}
resource "mist_org_network" "network_SE_TABLETS148" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 201
  subnet                 = "{{se_tablets}}/{{se_tablets_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SE_TABLETS"
}
resource "mist_org_network" "network_PS_PRINTERS149" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 102
  subnet                 = "{{ps_printers}}/{{ps_printers_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PS_PRINTERS"
}
resource "mist_org_network" "network_SE_WORKSTATIONS150" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 101
  subnet                 = "{{se_workstations}}/{{se_workstations_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SE_WORKSTATIONS"
}
resource "mist_org_network" "network_wdc_prod_wst_rtl151" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 120
  subnet                 = "10.160.254.0/28"
  disallow_mist_services = true
  tenants = {
    CTC_DC_10_Subnet = {
      addresses = [
        "10.0.0.0/8"
      ]
    }
    CTC_DC_192_Subnet = {
      addresses = [
        "192.168.0.0/16"
      ]
    }
    CTC_DC_172_Subnet = {
      addresses = [
        "172.16.0.0/12"
      ]
    }
    CTC_DC_42_Subnet = {
      addresses = [
        "42.0.0.0/8"
      ]
    }
  }
  name = "wdc-prod-wst-rtl"
}
resource "mist_org_network" "network_PC_FRONTIER_VL51152" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 51
  subnet                 = "{{pc_frontier}}/{{pc_frontier_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_FRONTIER_VL51"
}
resource "mist_org_network" "network_MWW_WORKSTATIONS_VL10153" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 101
  subnet                 = "{{mww_workstations}}/{{mww_workstations_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "MWW_WORKSTATIONS_VL101"
}
resource "mist_org_network" "network_PS_GWiFi_VL202154" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 202
  subnet                 = "{{ps_gwifi}}/{{ps_gwifi_cidr}}"
  disallow_mist_services = true
  name                   = "PS_GWiFi_VL202"
}
resource "mist_org_network" "network_AT_DIGITAL_VL400155" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 400
  subnet                 = "{{at_d_media}}/{{at_d_media_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "AT_DIGITAL_VL400"
}
resource "mist_org_network" "network_SE_THIRD_PARTY156" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 700
  subnet                 = "{{se_third_party}}/{{se_third_party_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SE_THIRD-PARTY"
}
resource "mist_org_network" "network_AT_DVR_VL300157" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 300
  subnet                 = "{{at_dvr}}/{{at_dvr_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "AT_DVR_VL300"
}
resource "mist_org_network" "network_CTR_GWiFi_VL202158" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 202
  subnet                 = "{{ctr_gwifi}}/{{ctr_gwifi_cidr}}"
  disallow_mist_services = true
  name                   = "CTR_GWiFi_VL202"
}
resource "mist_org_network" "network_PC_MGMT_VL1159" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "{{pc_default}}/{{pc_default_cidr}}"
  disallow_mist_services = false
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_MGMT_VL1"
}
resource "mist_org_network" "network_PS_CORP_WIRELESS_VL20160" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 200
  subnet                 = "{{ps_corp_wireless}}/{{ps_corp_wireless_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PS_CORP-WIRELESS_VL200"
}
resource "mist_org_network" "network_SE_MGMT161" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "{{se_default}}/{{se_default_cidr}}"
  disallow_mist_services = false
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SE_MGMT"
}
resource "mist_org_network" "network_CTP_CORP_WIRELESS162" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 200
  subnet                 = "{{ctp_corp_wireless}}/{{ctp_corp_wireless_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTP_CORP-WIRELESS"
}
resource "mist_org_network" "network_TRIPWIRE_10163" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "10.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    CTC_TRIPWIRE1 = {
      addresses = [
        "10.2.38.12/32",
        "10.2.38.13/32",
        "10.51.53.37/32",
        "10.17.58.54/32",
        "10.17.58.55/32",
        "10.2.36.106/32",
        "10.2.36.108/32",
        "10.17.36.106/32",
        "10.17.36.108/32",
        "10.18.80.150/32",
        "10.18.86.150/32",
        "10.127.80.150/32",
        "10.127.86.150/32"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "TRIPWIRE-10"
}
resource "mist_org_network" "network_PC_SQUID_VL53164" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 53
  subnet                 = "{{pc_squid}}/{{pc_squid_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_SQUID_VL53"
}
resource "mist_org_network" "network_CTP_POS_VL100165" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 300
  subnet                 = "{{ctp_pos}}/{{ctp_pos_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed                     = true
      no_readvertise_to_overlay  = false
      no_readvertise_to_lan_bgp  = false
      no_readvertise_to_lan_ospf = false
    }
  }
  name = "CTP_POS_VL100"
}
resource "mist_org_network" "network_SE_GWiFi166" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 202
  subnet                 = "{{se_gwifi}}/{{se_gwifi_cidr}}"
  disallow_mist_services = true
  name                   = "SE_GWiFi"
}
resource "mist_org_network" "network_PC_POS_VL4167" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 4
  subnet                 = "{{pc_pos}}/{{pc_pos_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_POS_VL4"
}
resource "mist_org_network" "network_CTR_CORP_WIRED_VL2168" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 2
  subnet                 = "{{ctr_corp_wired}}/{{ctr_corp_wired_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_CORP-WIRED_VL2"
}
resource "mist_org_network" "network_PC_GWiFi_VL202169" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 202
  subnet                 = "{{pc_gwifi}}/{{pc_gwifi_cidr}}"
  disallow_mist_services = true
  name                   = "PC_GWiFi_VL202"
}
resource "mist_org_network" "network_PS_PRINTERS_VL102170" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 102
  subnet                 = "{{ps_printers}}/{{ps_printers_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PS_PRINTERS_VL102"
}
resource "mist_org_network" "network_CTR_SCCM_BOPC_SERVERS171" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "42.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    CTR_SCCM_BOPC_SERVER14 = {
      addresses = [
        "42.201.3.40/32"
      ]
    }
    CTR_SCCM_BOPC_SERVER15 = {
      addresses = [
        "42.16.0.42/32"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_SCCM_BOPC_SERVERS_42"
}
resource "mist_org_network" "network_CTP_GWiFi172" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 202
  subnet                 = "{{ctp_gwifi}}/{{ctp_gwifi_cidr}}"
  disallow_mist_services = true
  name                   = "CTP_GWiFi"
}
resource "mist_org_network" "network_PC_RF_GUNS_VL3173" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3
  subnet                 = "{{pc_rf_guns}}/{{pc_rf_guns_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_RF-GUNS_VL3"
}
resource "mist_org_network" "network_CTR_MGMT_VL1174" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "{{ctr_default}}/{{ctr_default_cidr}}"
  disallow_mist_services = false
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_MGMT_VL1"
}
resource "mist_org_network" "network_MWW_VOIP_VL500175" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 500
  subnet                 = "{{mww_voip}}/{{mww_voip_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "MWW_VOIP_VL500"
}
resource "mist_org_network" "network_CTP_VAUGHN_OFFICE176" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "42.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    CTP_VAUGHN = {
      addresses = [
        "42.40.0.0/16"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTP_VAUGHN_OFFICE"
}
resource "mist_org_network" "network_wdc_prod_est_rtl177" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 122
  subnet                 = "10.160.254.32/28"
  disallow_mist_services = true
  tenants = {
    CTC_DC_10_Subnet = {
      addresses = [
        "10.0.0.0/8"
      ]
    }
    CTC_DC_192_Subnet = {
      addresses = [
        "192.168.0.0/16"
      ]
    }
    CTC_DC_172_Subnet = {
      addresses = [
        "172.16.0.0/12"
      ]
    }
    CTC_DC_42_Subnet = {
      addresses = [
        "42.0.0.0/8"
      ]
    }
  }
  name = "wdc-prod-est-rtl"
}
resource "mist_org_network" "network_PC_EFD_VOIP_VL31178" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 31
  subnet                 = "{{pc_efd_voip}}/{{pc_efd_voip_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_EFD-VOIP_VL31"
}
resource "mist_org_network" "network_10_0_0_0_8179" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "10.0.0.0/8"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "10_0_0_0_8"
}
resource "mist_org_network" "network_MWW_MGMT_VL1180" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "{{mww_default}}/{{mww_default_cidr}}"
  disallow_mist_services = false
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "MWW_MGMT_VL1"
}
resource "mist_org_network" "network_PC_WIRED_TABLETS_VL21181" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 21
  subnet                 = "{{pc_wired_tablets}}/{{pc_wired_tablets_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_WIRED-TABLETS_VL21"
}
resource "mist_org_network" "network_MWW_CORP_WIRELESS_VL2182" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 200
  subnet                 = "{{mww_corp_wireless}}/{{mww_corp_wireless_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "MWW_CORP-WIRELESS_VL200"
}
resource "mist_org_network" "network_CTP_PUMP_VL41183" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 41
  subnet                 = "{{ctp_pump}}/{{ctp_pump_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTP_PUMP_VL41"
}
resource "mist_org_network" "network_CTR_SRV_MGMT_VL5184" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 5
  subnet                 = "{{ctr_srv_mgmt}}/{{ctr_srv_mgmt_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_SRV-MGMT_VL5"
}
resource "mist_org_network" "network_CTP_BRAMPTON_SERVER_F185" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "42.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    CTP_BRAMPTON_SERVER_FARM1 = {
      addresses = [
        "42.106.0.0/16"
      ]
    }
    CTP_BRAMPTON_SERVER_FARM2 = {
      addresses = [
        "42.108.0.0/16"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTP_BRAMPTON_SERVER_FARM"
}
resource "mist_org_network" "network_CTR_MCAFEE_SERVERS_17186" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "172.16.0.0/12"
  disallow_mist_services = true
  tenants = {
    CTR_MCAFEE_SERVER8 = {
      addresses = [
        "172.21.131.43/32"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_MCAFEE_SERVERS_172"
}
resource "mist_org_network" "network_MWW_DVR_VL300187" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 300
  subnet                 = "{{mww_dvr}}/{{mww_dvr_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "MWW_DVR_VL300"
}
resource "mist_org_network" "network_CTR_D_VOIP188" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 63
  subnet                 = "{{ctr_d_voip}}/{{ctr_d_voip_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_D-VOIP"
}
resource "mist_org_network" "network_MWW_POS_VL100189" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 100
  subnet                 = "{{mww_pos}}/{{mww_pos_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "MWW_POS_VL100"
}
resource "mist_org_network" "network_CTP_GWiFi_VL202190" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 202
  subnet                 = "{{ctp_gwifi}}/{{ctp_gwifi_cidr}}"
  disallow_mist_services = true
  name                   = "CTP_GWiFi_VL202"
}
resource "mist_org_network" "network_CTR_3P_DIGITAL_VL7191" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 7
  subnet                 = "{{ctr_3p_digital}}/{{ctr_3p_digital_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_3P-DIGITAL_VL7"
}
resource "mist_org_network" "network_MWW_GWiFi_VL202192" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 202
  subnet                 = "{{mww_gwifi}}/{{mww_gwifi_cidr}}"
  disallow_mist_services = true
  name                   = "MWW_GWiFi_VL202"
}
resource "mist_org_network" "network_CTP_CORP_WIRED_VL101193" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 101
  subnet                 = "{{ctp_corp_wired}}/{{ctp_corp_wired_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTP_CORP-WIRED_VL101"
}
resource "mist_org_network" "network_CTR_D_WIRELESS194" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 61
  subnet                 = "{{ctr_d_wireless}}/{{ctr_d_wireless_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_D-WIRELESS"
}
resource "mist_org_network" "network_CTP_CORP_WIRELESS_VL2195" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 200
  subnet                 = "{{ctp_corp_wireless}}/{{ctp_corp_wireless_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTP_CORP-WIRELESS_VL200"
}
resource "mist_org_network" "network_SERVICENOW_10196" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "10.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    CTC_SERVICENOW = {
      addresses = [
        "10.126.57.239/32",
        "10.126.57.240/32",
        "10.126.57.157/32",
        "10.127.90.192/27"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "SERVICENOW-10"
}
resource "mist_org_network" "network_CTP_NON_CTC_VL700197" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 700
  subnet                 = "{{ctp_non_ctc}}/{{ctp_non_ctc_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTP_NON-CTC_VL700"
}
resource "mist_org_network" "network_OMNIBUS_10198" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "10.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    CTC_OMNIBUS = {
      addresses = [
        "10.126.18.51/32",
        "10.126.22.54/32",
        "10.126.26.51/32",
        "10.126.26.52/32",
        "10.18.218.62/32"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "OMNIBUS-10"
}
resource "mist_org_network" "network_CTP_2180_2190_SERVER_199" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "42.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    CTP_2180-2190_SERVER_FARM1 = {
      addresses = [
        "42.120.0.0/16"
      ]
    }
    CTP_2180-2190_SERVER_FARM2 = {
      addresses = [
        "42.124.0.0/16"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTP_2180-2190_SERVER_FARM"
}
resource "mist_org_network" "network_PS_VOIP_VL500200" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 500
  subnet                 = "{{ps_voip}}/{{ps_voip_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PS_VOIP_VL500"
}

resource "mist_org_network" "network_CTR_SCCM_BOPC_SERVERS201" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "10.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    CTR_SCCM_BOPC_SERVER1 = {
      addresses = [
        "10.127.20.44/32"
      ]
    }
    CTR_SCCM_BOPC_SERVER2 = {
      addresses = [
        "10.127.20.217/32"
      ]
    }
    CTR_SCCM_BOPC_SERVER3 = {
      addresses = [
        "10.127.20.218/32"
      ]
    }
    CTR_SCCM_BOPC_SERVER4 = {
      addresses = [
        "10.127.20.219/32"
      ]
    }
    CTR_SCCM_BOPC_SERVER5 = {
      addresses = [
        "10.127.20.204/32"
      ]
    }
    CTR_SCCM_BOPC_SERVER6 = {
      addresses = [
        "10.106.66.44/32"
      ]
    }
    CTR_SCCM_BOPC_SERVER7 = {
      addresses = [
        "10.127.21.43/32"
      ]
    }
    CTR_SCCM_BOPC_SERVER8 = {
      addresses = [
        "10.127.91.190/32"
      ]
    }
    CTR_SCCM_BOPC_SERVER9 = {
      addresses = [
        "10.127.91.217/32"
      ]
    }
    CTR_SCCM_BOPC_SERVER10 = {
      addresses = [
        "10.127.91.218/32"
      ]
    }
    CTR_SCCM_BOPC_SERVER11 = {
      addresses = [
        "10.127.91.219/32"
      ]
    }
    CTR_SCCM_BOPC_SERVER12 = {
      addresses = [
        "10.127.91.220/32"
      ]
    }
    CTR_SCCM_BOPC_SERVER13 = {
      addresses = [
        "10.120.56.44/32"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_SCCM_BOPC_SERVERS_10"
}
resource "mist_org_network" "network_CTP_WELAND202" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "172.21.0.0/16"
  disallow_mist_services = true
  tenants = {
    CTP_WELAND_PROD_POS = {
      addresses = [
        "172.21.0.80/28"
      ]
    }
    CTP_WELAND_TEST_POS = {
      addresses = [
        "172.21.0.208/28"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTP_WELAND"
}
resource "mist_org_network" "network_RETAIL_DMZ_HOST5203" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "10.100.116.125/32"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "RETAIL_DMZ_HOST5"
}
resource "mist_org_network" "network_test204" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "10.180.36.72/32"
  disallow_mist_services = false
  name                   = "test"
}
resource "mist_org_network" "network_PC_0834_Subnets_10205" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "10.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    PC_0834_VLAN-2-3-4-21 = {
      addresses = [
        "10.221.4.0/22"
      ]
    }
    PC_0834_VLAN-5 = {
      addresses = [
        "10.227.65.0/24"
      ]
    }
    PC_0834_VLAN-6-61-63 = {
      addresses = [
        "10.141.4.0/22"
      ]
    }
    PC_0834_VLAN-31-32 = {
      addresses = [
        "10.238.130.0/23"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_0834_Subnets_10"
}
resource "mist_org_network" "network_SE_POS_VL100206" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "{{se_pos}}/{{se_pos_cidr}}"
  vpn_access = {
    OrgOverlay = {
      routed                    = true
      no_readvertise_to_overlay = false
      no_readvertise_to_lan_bgp = false
    }
  }
  vlan_id                = 100
  disallow_mist_services = true
  name                   = "SE_POS_VL100"
}
resource "mist_org_network" "network_PHL_MGMT_VL1207" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "{{phl_default}}/{{phl_default_cidr}}"
  disallow_mist_services = false
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PHL_MGMT_VL1"
}
resource "mist_org_network" "network_BELL_SMI_NAT208" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "69.158.230.188/32"
  disallow_mist_services = true
  tenants = {
    BELL_SMI_NAT_Users_Subnet = {
      addresses = [
        "69.158.230.188/32"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed                    = true
      no_readvertise_to_overlay = false
      no_readvertise_to_lan_bgp = false
    }
  }
  name = "BELL_SMI_NAT"
}
resource "mist_org_network" "network_CTP_NEW_SSL_GW209" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "192.168.164.0/24"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTP_NEW_SSL_GW"
}
resource "mist_org_network" "network_RETAIL_DMZ_HOST6210" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "10.173.7.2/32"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "RETAIL_DMZ_HOST6"
}
resource "mist_org_network" "network_REMOTE_STORE_42_NETWO211" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "42.0.0.0/8"
  tenants = {
    REMOTE_STORES_42 = {
      addresses = [
        "{{Remote_Store1_VL_PRI_2}}",
        "{{Remote_Store1_VL_PRI_6}}",
        "{{Remote_Store2_VL_PRI_2}}",
        "{{Remote_Store2_VL_PRI_6}}",
        "{{Remote_Store3_VL_PRI_2}}",
        "{{Remote_Store3_VL_PRI_6}}",
        "{{Remote_Store4_VL_PRI_2}}",
        "{{Remote_Store4_VL_PRI_6}}"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed                    = true
      no_readvertise_to_overlay = false
      no_readvertise_to_lan_bgp = false
    }
  }
  vlan_id                = 3797
  disallow_mist_services = true
  name                   = "REMOTE_STORE_42_NETWORKS"
}
resource "mist_org_network" "network_PHL_POS_VL100212" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 100
  subnet                 = "{{phl_pos}}/{{phl_pos_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PHL_POS_VL100"
}
resource "mist_org_network" "network_iptest213" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "10.162.115.88/32"
  disallow_mist_services = false
  vlan_id                = 101
  name                   = "iptest"
}
resource "mist_org_network" "network_CTP_WDC_Subnets214" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "10.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    CTP_WDC_Subnet1 = {
      addresses = [
        "10.2.55.0/24"
      ]
    }
    CTP_WDC_Subnet2 = {
      addresses = [
        "10.2.56.0/23"
      ]
    }
    CTP_WDC_Subnet3 = {
      addresses = [
        "10.2.0.0/16"
      ]
    }
    CTP_WDC_Subnet4 = {
      addresses = [
        "10.11.0.0/16"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTP_WDC_Subnets"
}
resource "mist_org_network" "network_RETAIL_DMZ_HOST10215" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "10.100.34.97/32"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "RETAIL_DMZ_HOST10"
}
resource "mist_org_network" "network_REMOTE_STORE_GLOBAL_4216" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "42.0.0.0/8"
  tenants = {
    STORE_42_SUPERNETS = {
      addresses = [
        "42.60.0.0/14",
        "42.64.0.0/14",
        "42.240.0.0/13",
        "42.76.0.0/16"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed                    = true
      no_readvertise_to_overlay = false
      no_readvertise_to_lan_bgp = false
    }
  }
  vlan_id                = 3797
  disallow_mist_services = true
  name                   = "REMOTE_STORE_GLOBAL_42_BLOCK"
}
resource "mist_org_network" "network_MWW_LTE_MGMT_VL888217" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 888
  subnet                 = "192.168.98.0/30"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = false
      static_nat = {
        "{{mww_natted_ip}}" = {
          name        = "SMI-TO-LTE"
          internal_ip = "192.168.98.2"
        }
      }
      nat_pool = "{{mww_natted_ip}}/32"
    }
  }
  name = "MWW_LTE-MGMT_VL888"
}
resource "mist_org_network" "network_SE_TABLETS_VL201218" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "{{se_tablets}}/{{se_tablets_cidr}}"
  vpn_access = {
    OrgOverlay = {
      routed                    = true
      no_readvertise_to_overlay = false
      no_readvertise_to_lan_bgp = false
    }
  }
  vlan_id                = 201
  disallow_mist_services = true
  name                   = "SE_TABLETS_VL201"
}
resource "mist_org_network" "network_PHL_LTE_MGMT_VL888219" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 888
  subnet                 = "192.168.98.0/30"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = false
      static_nat = {
        "{{phl_natted_ip}}" = {
          name        = "SMI-TO-LTE"
          internal_ip = "192.168.98.2"
        }
      }
      no_readvertise_to_overlay = false
      no_readvertise_to_lan_bgp = false
      nat_pool                  = "{{phl_natted_ip}}/30"
    }
  }
  name = "PHL_LTE-MGMT_VL888"
}
resource "mist_org_network" "network_wdc_prod_dc220" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3796
  subnet                 = "10.16.2.224/29"
  disallow_mist_services = true
  tenants = {
    CTC_DC_10_Subnet = {
      addresses = [
        "10.0.0.0/8"
      ]
    }
    CTC_DC_192_Subnet = {
      addresses = [
        "192.168.0.0/16"
      ]
    }
    CTC_DC_172_Subnet = {
      addresses = [
        "172.16.0.0/12"
      ]
    }
    CTC_DC_42_Subnet = {
      addresses = [
        "42.0.0.0/8"
      ]
    }
    Azure-Supernet = {
      addresses = [
        "100.64.0.0/10"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed                    = true
      no_readvertise_to_overlay = false
      no_readvertise_to_lan_bgp = false
    }
  }
  name = "wdc-prod-dc"
}
resource "mist_org_network" "network_CTP_RPM_SERVERS_10221" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "10.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    CTP_RPM_SERVER5 = {
      addresses = [
        "10.127.91.110/32"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTP_RPM_SERVERS_10"
}
resource "mist_org_network" "network_CTR_LTE_MGMT_VL888222" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 888
  subnet                 = "192.168.98.0/30"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = false
      static_nat = {
        "{{ctr_natted_ip}}" = {
          name        = "SMI-TO-LTE"
          internal_ip = "192.168.98.2"
        }
      }
      nat_pool = "{{ctr_natted_ip}}/32"
    }
  }
  name = "CTR_LTE-MGMT_VL888"
}
resource "mist_org_network" "network_172_16_0_0_12223" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "172.16.0.0/12"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "172_16_0_0_12"
}
resource "mist_org_network" "network_CTP1407_PC_Test224" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "10.183.46.71/32"
  disallow_mist_services = false
  vlan_id                = 101
  name                   = "CTP1407-PC-Test"
}
resource "mist_org_network" "network_PHL_TABLETS_VL201225" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 201
  subnet                 = "{{phl_tablets}}/{{phl_tablets_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PHL_TABLETS_VL201"
}
resource "mist_org_network" "network_PC_D_VOIP_VL63226" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 63
  subnet                 = "{{pc_d_voip}}/{{pc_d_voip_cidr}}"
  disallow_mist_services = true
  tenants = {
    Any = {
      addresses = [
        "0.0.0.0/0"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed                    = true
      no_readvertise_to_overlay = false
      no_readvertise_to_lan_bgp = false
    }
  }
  name = "PC_D-VOIP_VL63"
}
resource "mist_org_network" "network_CTP_RPM_SERVERS_42227" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "42.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    CTP_RPM_SERVER1 = {
      addresses = [
        "42.108.133.133/32"
      ]
    }
    CTP_RPM_SERVER2 = {
      addresses = [
        "42.108.134.17/32"
      ]
    }
    CTP_RPM_SERVER3 = {
      addresses = [
        "42.108.34.77/32"
      ]
    }
    CTP_RPM_SERVER4 = {
      addresses = [
        "42.108.46.126/32"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTP_RPM_SERVERS_42"
}
resource "mist_org_network" "network_CTR_0424_Subnets_42228" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "42.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    CTR_0424_VLAN-2 = {
      addresses = [
        "42.63.85.0/24"
      ]
    }
    CTR_0424_VLAN-6 = {
      addresses = [
        "42.243.78.0/24"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_0424_Subnets_42"
}
resource "mist_org_network" "network_CTR_AD_SERVERS_10229" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "10.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    CTR_AD_SERVER1 = {
      addresses = [
        "10.125.120.0/23"
      ]
    }
    CTR_AD_SERVER3 = {
      addresses = [
        "10.2.55.0/24"
      ]
    }
    CTR_AD_SERVER4 = {
      addresses = [
        "10.2.56.0/23"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_AD_SERVERS_10"
}
resource "mist_org_network" "network_HykoKeys230" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "10.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    HykoKeys-Users = {
      addresses = [
        "10.168.194.97/27",
        "10.186.118.65/26"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "HykoKeys"
}
resource "mist_org_network" "network_CTR_0424_Subnets_10231" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "10.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    CTR_0424_VLAN-2-3-4-21 = {
      addresses = [
        "10.214.156.0/22"
      ]
    }
    CTR_0424_VLAN-5 = {
      addresses = [
        "10.225.167.0/24"
      ]
    }
    CTR_0424_VLAN-6-61-63 = {
      addresses = [
        "10.134.156.0/22"
      ]
    }
    CTR_0424_VLAN-31-32 = {
      addresses = [
        "10.235.78.0/23"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_0424_Subnets_10"
}
resource "mist_org_network" "network_PC_LTE_MGMT_VL888232" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 888
  subnet                 = "192.168.98.0/30"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = false
      static_nat = {
        "{{pc_natted_ip}}" = {
          name        = "SMI-TO-LTE"
          internal_ip = "192.168.98.2"
        }
      }
      nat_pool = "{{pc_natted_ip}}/32"
    }
  }
  name = "PC_LTE-MGMT_VL888"
}
resource "mist_org_network" "network_AT_LTE_MGMT_VL888233" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 888
  subnet                 = "192.168.98.0/30"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = false
      static_nat = {
        "{{at_natted_ip}}" = {
          name        = "SMI-to-LTE"
          internal_ip = "192.168.98.2"
        }
      }
      nat_pool = "{{at_natted_ip}}/32"
    }
  }
  name = "AT_LTE-MGMT_VL888"
}
resource "mist_org_network" "network_CTR_0043_Subnets_10234" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "10.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    CTR_0043_VLAN-2-3-4-21 = {
      addresses = [
        "10.208.168.0/22"
      ]
    }
    CTR_0043_VLAN-5 = {
      addresses = [
        "10.224.42.0/24"
      ]
    }
    CTR_0043_VLAN-6-61-63 = {
      addresses = [
        "10.128.168.0/22"
      ]
    }
    CTR_0043_VLAN-31-32 = {
      addresses = [
        "10.232.84.0/23"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_0043_Subnets_10"
}
resource "mist_org_network" "network_CTP_LTE_MGMT_VL888235" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 888
  subnet                 = "192.168.98.0/30"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = false
      static_nat = {
        "{{ctp_natted_ip}}" = {
          name        = "SMI-TO-LTE"
          internal_ip = "192.168.98.2"
        }
      }
      nat_pool = "{{ctp_natted_ip}}/32"
    }
  }
  name = "CTP_LTE-MGMT_VL888"
}
resource "mist_org_network" "network_REMOTE_STORE_GLOBAL_1236" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "10.0.0.0/8"
  tenants = {
    STORE_SUPERNETS = {
      addresses = [
        "10.128.0.0/12",
        "10.208.0.0/12",
        "10.224.0.0/14",
        "10.228.0.0/14",
        "10.232.0.0/13",
        "10.180.0.0/14",
        "10.184.0.0/14"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed                    = true
      no_readvertise_to_overlay = false
      no_readvertise_to_lan_bgp = false
    }
  }
  vlan_id                = 3797
  disallow_mist_services = true
  name                   = "REMOTE_STORE_GLOBAL_10_BLOCK"
}
resource "mist_org_network" "network_CTR_CORP_CAMPUS_10237" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "10.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    CTR_CORP_CAMPUS1 = {
      addresses = [
        "10.112.0.0/12"
      ]
    }
    CTR_CORP_CAMPUS9 = {
      addresses = [
        "10.16.0.0/14"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "CTR_CORP_CAMPUS_10"
}
resource "mist_org_network" "network_SE_WORKSTATIONS_VL101238" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "{{se_workstations}}/{{se_workstations_cidr}}"
  vpn_access = {
    OrgOverlay = {
      routed                    = true
      no_readvertise_to_overlay = false
      no_readvertise_to_lan_bgp = false
    }
  }
  vlan_id                = 101
  disallow_mist_services = true
  name                   = "SE_WORKSTATIONS_VL101"
}
resource "mist_org_network" "network_PHL_DIGITAL_VL400239" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 400
  subnet                 = "{{phl_d_media}}/{{phl_d_media_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PHL_DIGITAL_VL400"
}
resource "mist_org_network" "network_aj_prod_est_rtl240" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 122
  subnet                 = "10.160.255.32/28"
  disallow_mist_services = true
  tenants = {
    CTC_DC_10_Subnet = {
      addresses = [
        "10.0.0.0/8"
      ]
    }
    CTC_DC_192_Subnet = {
      addresses = [
        "192.168.0.0/16"
      ]
    }
    CTC_DC_172_Subnet = {
      addresses = [
        "172.16.0.0/12"
      ]
    }
    CTC_DC_42_Subnet = {
      addresses = [
        "42.0.0.0/8"
      ]
    }
  }
  name = "aj-prod-est-rtl"
}
resource "mist_org_network" "network_10_231_205_0_25241" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "10.231.205.0/25"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "10-231-205-0_25"
}
resource "mist_org_network" "network_RETAIL_DMZ_HOST8242" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "10.100.36.78/32"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "RETAIL_DMZ_HOST8"
}
resource "mist_org_network" "network_PS_LTE_MGMT_VL888243" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 888
  subnet                 = "192.168.98.0/30"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = false
      static_nat = {
        "{{ps_natted_ip}}" = {
          name        = "SMI-TO-LTE"
          internal_ip = "192.168.98.2"
        }
      }
      nat_pool = "{{ps_natted_ip}}/32"
    }
  }
  name = "PS_LTE-MGMT_VL888"
}
resource "mist_org_network" "network_SRX345_lan244" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "100.100.150.0/24"
  vpn_access = {
    OrgOverlay = {
      routed                     = true
      no_readvertise_to_overlay  = false
      no_readvertise_to_lan_bgp  = false
      no_readvertise_to_lan_ospf = false
    }
  }
  vlan_id                = 500
  disallow_mist_services = false
  name                   = "SRX345_lan"
}
resource "mist_org_network" "network_BELL_SMI_WIN_JUMP_SER245" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  subnet                 = "142.125.248.192/32"
  disallow_mist_services = true
  tenants = {
    BELL_SMI_JUMP_USERS_SUBNET = {
      addresses = [
        "142.125.248.192/32"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed                    = true
      no_readvertise_to_overlay = false
      no_readvertise_to_lan_bgp = false
    }
  }
  name = "BELL_SMI_WIN_JUMP_SERVER"
}
resource "mist_org_network" "network_SC_LTE_MGMT_VL888246" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 888
  subnet                 = "192.168.98.0/30"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = false
      static_nat = {
        "{{sc_natted_ip}}" = {
          name        = "SMI-TO-LTE"
          internal_ip = "192.168.98.2"
        }
      }
      nat_pool = "{{sc_natted_ip}}/32"
    }
  }
  name = "SC_LTE-MGMT_VL888"
}
resource "mist_org_network" "network_PC_0843_Subnets_10247" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3797
  subnet                 = "10.0.0.0/8"
  disallow_mist_services = true
  tenants = {
    PC_0843_VLAN-2-3-4-21 = {
      addresses = [
        "10.221.40.0/22"
      ]
    }
    PC_0843_VLAN-5 = {
      addresses = [
        "10.227.74.0/24"
      ]
    }
    PC_0843_VLAN-6-61-63 = {
      addresses = [
        "10.141.40.0/22"
      ]
    }
    PC_0843_VLAN-31-32 = {
      addresses = [
        "10.238.148.0/23"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PC_0843_Subnets_10"
}
resource "mist_org_network" "network_SE_MGMT_VL1248" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "{{se_default}}/{{se_default_cidr}}"
  vpn_access = {
    OrgOverlay = {
      routed                    = true
      no_readvertise_to_overlay = false
      no_readvertise_to_lan_bgp = false
    }
  }
  disallow_mist_services = false
  name                   = "SE_MGMT_VL1"
}
resource "mist_org_network" "network_PHL_VOIP_VL500249" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 500
  subnet                 = "{{phl_voip}}/{{phl_voip_cidr}}"
  disallow_mist_services = true
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  name = "PHL_VOIP_VL500"
}
resource "mist_org_network" "network_PC_D_WIRELESS_VL61250" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 61
  subnet                 = "{{pc_d_wireless}}/{{pc_d_wireless_cidr}}"
  disallow_mist_services = true
  tenants = {
    Any = {
      addresses = [
        "0.0.0.0/0"
      ]
    }
  }
  vpn_access = {
    OrgOverlay = {
      routed                    = true
      no_readvertise_to_overlay = false
      no_readvertise_to_lan_bgp = false
    }
  }
  name = "PC_D-WIRELESS_VL61"
}
