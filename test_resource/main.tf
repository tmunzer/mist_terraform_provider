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
  # username = local.envs["USERNAME"]
  # password = local.envs["PASSWORD"]
}

### ORG
resource "mist_org" "terraform_test" {
  name = "Terraform Testing"
}

resource "mist_org_servicepolicy" "test1" {
  org_id = mist_org.terraform_test.id
  tenants = [
    "PRD-Core"
  ]
  services = [
    "app_mqtt"
  ]
  action = "allow"
  idp = {
    enabled = false
  }
  name = "Policy-1"
}
resource "mist_org_servicepolicy" "test2" {
  org_id = mist_org.terraform_test.id
  tenants = [
    "cctv-e.PRD-IoT"
  ]
  services = [
    "app_proxy"
  ]
  action = "allow"
  name   = "Policy-3"
}
resource "mist_org_servicepolicy" "test3" {
  org_id = mist_org.terraform_test.id
  tenants = [
    "cctv-b.PRD-IoT"
  ]
  services = [
    "any"
  ]
  action = "allow"
  name   = "Policy-2"
}
resource "mist_org_servicepolicy" "test4" {
  org_id = mist_org.terraform_test.id
  tenants = [
    "saltlake-lan2"
  ]
  services = [
    "Internet"
  ]
  action = "allow"
  name   = "UDP-ZERO-DATA-DNS-GOOGLE"
}
resource "mist_org_servicepolicy" "test5" {
  org_id = mist_org.terraform_test.id
  tenants = [
    "corp_datacenter1",
    "corp_datacenter2",
    "core_dc2_network",
    "mgmt_datacenter1",
    "mgmt_datacenter2",
    "corp_network",
    "mgmt_network"
  ]
  services = [
    "rfc1918"
  ]
  action = "allow"
  idp = {
    enabled = false
  }
  name = "access-to-rfc1918"
}
resource "mist_org_servicepolicy" "test6" {
  org_id = mist_org.terraform_test.id
  tenants = [
    "guest"
  ]
  services = [
    "guest-internet"
  ]
  action = "allow"
  idp = {
    enabled    = true,
    profile    = "standard",
    alert_only = true
  }
  name = "Guest-IDP"
}
resource "mist_org_service" "service143" {
  org_id      = mist_org.terraform_test.id
  description = "App for detecting office 365"
  apps = [
    "office365"
  ]
  traffic_type = "default"
  name         = "XX-office365"
  type         = "apps"
}

resource "mist_org_idpprofile" "test1" {
  org_id       = mist_org.terraform_test.id
  base_profile = "standard"
  overwrites = [
    {
      name   = "r1"
      action = "drop"
      matching = {
        severity   = []
        dst_subnet = []
        attack_name = [
          "SSL:OVERFLOW:KEY-ARG-NO-ENTROPY"
        ]
      }
    },
    {
      name = "r2"
      matching = {
        severity   = []
        dst_subnet = []
        attack_name = [
          "UDP:ZERO-DATA"
        ]
      }
    }
  ]
  name = "custom-standard"
}
resource "mist_org_nactag" "nactag_VLAN_Crew451" {
  org_id = mist_org.terraform_test.id
  vlan   = "451"
  name   = "VLAN-Crew451"
  type   = "vlan"
}
resource "mist_org_idpprofile" "test2" {
  org_id       = mist_org.terraform_test.id
  base_profile = "standard"
  overwrites = [
    {
      name   = "UDP-ZERO-DATA-AllowGoogle1"
      action = "alert"
      matching = {
        severity = [],
        dst_subnet = [
          "8.8.8.8/32"
        ],
        attack_name = [
          "UDP:ZERO-DATA"
        ]
      }
    }
  ]
  name = "UDP-ZERO-DATA"
}
resource "mist_org_idpprofile" "test3" {
  org_id       = mist_org.terraform_test.id
  base_profile = "critical"
  overwrites = [
    {
      name = "server_bypass"
      matching = {
        severity   = []
        dst_subnet = []
        attack_name = [
          "SSL:OVERFLOW:KEY-ARG-NO-ENTROPY"
        ]
      }
    },
    {
      name = "aide-demo-guest-bypass"
      matching = {
        severity   = []
        dst_subnet = []
        attack_name = [
          "SSL:OVERFLOW:KEY-ARG-NO-ENTROPY"
        ]
      }
    }
  ]
  name = "aide-demo-dmz-bypass"
}
resource "mist_org_idpprofile" "test4" {
  org_id       = mist_org.terraform_test.id
  base_profile = "critical"
  name         = "nn-test"
}
resource "mist_org_idpprofile" "test5" {
  org_id       = mist_org.terraform_test.id
  base_profile = "standard"
  overwrites = [
    {
      name   = "zero-data"
      action = "close"
      matching = {
        severity = [],
        dst_subnet = [
          "9.9.9.9/32",
          "1.0.0.1/32"
        ]
        attack_name = [
          "UDP:ZERO-DATA"
        ]
      }
    }
  ]
  name = "zero-data"
}
resource "mist_org_idpprofile" "test6" {
  org_id       = mist_org.terraform_test.id
  base_profile = "standard"
  overwrites = [
    {
      name   = "CVE"
      action = "alert"
      matching = {
        severity = []
        dst_subnet = [
          "192.168.94.2/32",
          "93.184.215.14/32",
          "13.89.179.10/32",
          "1.0.0.1/32",
          "9.9.9.9/32",
          "1.1.1.1/32",
          "142.251.33.99/32",
          "23.102.110.248/32",
          "20.189.173.10/32",
          "13.89.179.14/32",
          "142.251.211.227/32"
        ]
        attack_name = [
          "HTTP:INFO-LEAK:NO-SP-AFTER-FLD",
          "HTTP:OVERFLOW:HEADER",
          "HTTP:INVALID:HDR-FIELD",
          "HTTP:INVALID:INV-CONT-ENC",
          "SSL:AUDIT:EXPORT-GRADE-CIPHER",
          "PORTMAPPER:ERR:SHORT-READ",
          "SSL:OVERFLOW:KEY-ARG-NO-ENTROPY"
        ]
      }
    }
  ]
  name = "CVE"
}



resource "mist_org_setting" "terraform_test" {
  org_id = mist_org.terraform_test.id
  password_policy = {
    enabled                  = true
    min_length               = 8
    requires_special_char    = false
    requires_two_factor_auth = false
  }
  disable_pcap    = false
  ui_idle_timeout = 0
  cacerts = [
    "-----BEGIN CERTIFICATE-----\nMIIFnTCCA4WgAwIBAgIHBfyZT9pPxzANBgkqhkiG9w0BAQsFADBmMQswCQYDVQQG\nEwJVUzEUMBIGA1UECgwLRm94cGFzcyBJbmMxGTAXBgNVBAsMEEN1c3RvbWVyIEVB\nUC1UTFMxJjAkBgNVBAMMHXN0YWdvbmUucmFkaXVzLmZveHBhc3MuY29tIENBMB4X\nDTIxMDQyNzA1MTUxNloXDTMxMDQyNTA1MTUxNlowZjELMAkGA1UEBhMCVVMxFDAS\nBgNVBAoMC0ZveHBhc3MgSW5jMRkwFwYDVQQLDBBDdXN0b21lciBFQVAtVExTMSYw\nJAYDVQQDDB1zdGFnb25lLnJhZGl1cy5mb3hwYXNzLmNvbSBDQTCCAiIwDQYJKoZI\nhvcNAQEBBQADggIPADCCAgoCggIBALdhq+AcZhsem5QGxGgXDMKUJxR+OPn3unvw\nyx+Hif/QVlZqHzpmw2biTz5qPkNS6+F0fu932vEnGUL5SfbagW+XKHtCM5NbfFeP\nlYlkZzifmljjaFG1BlnfdV4Zk79HBFWsI4CyLLer2bIJgjvWdLiaVnu/yOeiMCWU\nIPtQN5GbQf2TdlJv24AoFIZ2s6I7e30wrc1Yy4RH2i/LARfi0FWWxeuE9lD/Gk1n\nXCHiBZDRCoAmszmB466senLYPOdYg6WrlTFo7+C/ddrI4JDjHMJuK0SVM76uWrVd\nv2oOQVcCB8RqAz18424NURCbvJpd4rdX9n5wne/m33GNknNkHKz82SRDeUJPYQV0\n7D/eTIAtD9LRhizS/UD2DQkfYrWr/YVNU0KkS8FGeByjAAYiDKhuRHpqB6au8HPM\nj1+KLOPDlan5hl1SSu/lCtBSYKZgfA7S4fcWz27+NfAZR8fjy1kfWIPn8QGAc5uD\nOhiER73I5bPQBr6RLpQ1PxDHvO//JOt/ejd7aEi+DpK5GuwFxD0FmLWEQcyKkjUD\nhZijNtJRie08q/yhKySTlkr4fxeCR56p1DlB1S15myPILCUZMiHMm0ll9UAzYatT\nBTY/hyNkZUePtllOZXLvU1YUJIzlzpU+Kr45+hO5a0gLwoDHoidNxTK5TOx4KkYe\nrx9RtFprAgMBAAGjUDBOMB0GA1UdDgQWBBTaOaPuXmtLDTJVv++VYBiQr9gHCTAf\nBgNVHSMEGDAWgBTaOaPuXmtLDTJVv++VYBiQr9gHCTAMBgNVHRMEBTADAQH/MA0G\nCSqGSIb3DQEBCwUAA4ICAQAYt2Sg3fkXYo2/hke+gNA92nXHXZXh+xupOg4I5f0Y\nh07Y/I5uZKU5uUV33ls7BrLBQm9J47kncB5XiRLJR0fWAHEAA8FaoEw44mnyIn0N\nNxbfDayIFbUzShMD7GqIBNMd+R9WqhXrBwFscWHB0idHRnbY93sxuIRY+NnYNhDK\n20J0wU2cPkzCT/rDWbcE2gCS5OTasUCvndrSQBSKobKU6JIia0hxRu6gekbIr36u\nXIaxiRKv+N7xLgEXfImITcxF/YS7TWa98ZiaHWTvRArUxmBjitPerTXXMNvWxF0K\nw7+K3piKkD2tvEpTAYNcVAexE6ElTHYvKIvYaYlZ4qyoLSVxzkiXBrsPyiAGsELQ\nWazQlfdxk0vtulAcyv3bmXsGqjxTxBQ5fPZ7wYMJQfa1jW55Fz6osmphWbRuj6Ht\nCenlr8+3KqkwmR5fCHmcaY2CuJzx1/ClhH4dGplztI4MhCVnAplSwCxOjV0xcFj9\nWu1nheb8M9mTWX3/4vkopWBnxs8rYelZZ8XsVjW+L7vtSLSpiHjI8wMv+441g+Qp\nFJ+I/xoOXn3LT9+dfTZw+QnAHF4Yz3hhphZYS0XR9BgRZLsufuKU3pHmIbmnDNed\n+B814wVglGU6pVfjx7m+l3NcuLq0LszWHGcSRbpTduch7u24yQ8Y+VWg/RgNZHzM\nZw==\n-----END CERTIFICATE-----"
  ]
  mgmt = {
    use_wxtunnel = false
    use_mxtunnel = false
  }
  security = {
    limit_ssh_access = false
  }
  device_updown_threshold = 2
  installer = {
    grace_period      = 365
    allow_all_devices = false
    allow_all_sites   = false
  }
  mist_nac = {
    cacerts = [
      "-----BEGIN CERTIFICATE-----\nMIIIojCCBoqgAwIBAgITaQAAAA4ZVhqPIAs/ywAAAAAADjANBgkqhkiG9w0BAQ0F\nADBiMQswCQYDVQQGEwJVUzEdMBsGA1UEChMUSnVuaXBlciBOZXR3b3JrcyBJbmMx\nNDAyBgNVBAMTK0p1bmlwZXIgTmV0d29ya3MgUm9vdCBDZXJ0aWZpY2F0ZSBBdXRo\nb3JpdHkwHhcNMjEwOTAyMjA0NjU3WhcNMjYwOTAyMjA1NjU3WjBWMRMwEQYKCZIm\niZPyLGQBGRYDbmV0MRQwEgYKCZImiZPyLGQBGRYEam5wcjEpMCcGA1UEAxMgSnVu\naXBlciBOZXR3b3JrcyBJc3N1aW5nIEFXUzEgQ0EwggIiMA0GCSqGSIb3DQEBAQUA\nA4ICDwAwggIKAoICAQC8aPpLWDuPrNm0pPULwgzhCQYTReZNUZGIOK0fD6YAmwpy\nrY9AhFaGs0FJudBinT/LWg084OPQAWU8O2U8g6oMJvDyFRrE1O5nD9zJ0fs84A0u\n5q507j1i+Emvqj5aaPKZnggvPd3dNey9qn1WU3ElJK5jM6CfwSfu8cBHCRwpBUkb\nnvK31YZb8w6dQhQ9Gr292s6wzCjWryHO/EMa6wwxk7mtCKi0EfUascaaqLl35qun\nsSjQDYuNWPhVQ2Qhlse58+2Eb65q94V5S2wIfP/gmJ/kWCYeacWaS0MUy7FOCpDM\n4MHWJZPWMe72Om6XhMAebfJohejs4yHCbromjK9r1eeqyRnmUXkJmmF2ZfZ9RhLn\nBm86Q7uWpb3boupNNH4azsqeQBXfVUldZAATR9q1J2BBZd7Ox1DoroDZ6VVZUq9j\n4cR6DMnSxZc8+4TfDxQpfVnz56JEyDYOUEzjDTyRGIqx0CI5fhBMg11ey5LkO11C\n12sxgLLmCp+J5S6gziInAF+g45vBMWvn/N7oDgt1ubd6ZN/sIhjPeHT+AqErlfdQ\nurqpToQpM4WdKCFbi58boFVmVHVAN3pxDxoiqiI2qJO/Vm1ozyncEDDIvg6cTMLz\nhu3tQ3qg/tlee9ftjr3F49M9H5t6Cvona74GLPRwkwBvroUQYcpEh9ajzDll1QID\nAQABo4IDWzCCA1cwEAYJKwYBBAGCNxUBBAMCAQEwIwYJKwYBBAGCNxUCBBYEFBTa\ne+zk8uL7fhg+C4ubfYLIkeBnMB0GA1UdDgQWBBQfQbpi/2hX99sRJIirVNh7Lagi\nuDAZBgkrBgEEAYI3FAIEDB4KAFMAdQBiAEMAQTALBgNVHQ8EBAMCAYYwEgYDVR0T\nAQH/BAgwBgEB/wIBATAfBgNVHSMEGDAWgBT8FKc0zZCXqoGVmvg0uqkFY8Us4jCC\nAVEGA1UdHwSCAUgwggFEMIIBQKCCATygggE4hllodHRwOi8vY2RwYWlhLmpucHIu\nbmV0L0NlcnRFbnJvbGwvSnVuaXBlciUyME5ldHdvcmtzJTIwUm9vdCUyMENlcnRp\nZmljYXRlJTIwQXV0aG9yaXR5LmNybIaB2mxkYXA6Ly8vQ049SnVuaXBlciUyME5l\ndHdvcmtzJTIwUm9vdCUyMENlcnRpZmljYXRlJTIwQXV0aG9yaXR5LENOPUNBSk5Q\nUlJPT1QsQ049Q0RQLENOPVB1YmxpYyUyMEtleSUyMFNlcnZpY2VzLENOPVNlcnZp\nY2VzLENOPUNvbmZpZ3VyYXRpb24sREM9am5wcixEQz1uZXQ/Y2VydGlmaWNhdGVS\nZXZvY2F0aW9uTGlzdD9iYXNlP29iamVjdENsYXNzPWNSTERpc3RyaWJ1dGlvblBv\naW50MIIBSwYIKwYBBQUHAQEEggE9MIIBOTBlBggrBgEFBQcwAoZZaHR0cDovL2Nk\ncGFpYS5qbnByLm5ldC9DZXJ0RW5yb2xsL0p1bmlwZXIlMjBOZXR3b3JrcyUyMFJv\nb3QlMjBDZXJ0aWZpY2F0ZSUyMEF1dGhvcml0eS5jcnQwgc8GCCsGAQUFBzAChoHC\nbGRhcDovLy9DTj1KdW5pcGVyJTIwTmV0d29ya3MlMjBSb290JTIwQ2VydGlmaWNh\ndGUlMjBBdXRob3JpdHksQ049QUlBLENOPVB1YmxpYyUyMEtleSUyMFNlcnZpY2Vz\nLENOPVNlcnZpY2VzLENOPUNvbmZpZ3VyYXRpb24sREM9am5wcixEQz1uZXQ/Y0FD\nZXJ0aWZpY2F0ZT9iYXNlP29iamVjdENsYXNzPWNlcnRpZmljYXRpb25BdXRob3Jp\ndHkwDQYJKoZIhvcNAQENBQADggIBAGxmZ/rZicA8OVK+6qg1ikQGnMQp7iMEKVg5\n+tKiltwoAPrngFYYS/3bOtP0EjUfxUZkdkJncoSUUgjhs9E+SdUTNg7o0LtCL4lQ\n+CTcKPBJyKl5q40D2Zsmnegh4TADSy/t16yxJQtVHqjzuImr4f821UL0JaEwukXg\nO2MWQrsSbb481KSVE6hvv1aGqFZKH8XXEX0rxOc3tdVCclDZwmIbZysxQn2tERu7\nBg3fAnI1xUNFUhv1+7MHA5KIazBYxMVJWgtXJ5g81jBHN7meqvqYhT8scHyPsux6\nbiG/pctnEFwd3e1VotPxL4WzwYp2qfqX48+oeN9Qeeij+7vrtFATPnWsD5vDSCC0\nCKDqnvpVLmnbqKAK9UC6DpaPg7uf2C7rQl6SqzOjcyfORKqkZr5QtAbbX8eNc0iB\njqkX6m8+dOwEqcjuGV95PzEOk6jIqkO9O4lP2yx0EEHG08/itQ+UNc5RH6hCtaQT\nAwO0jTzJKdl1dHMNcvNFphHJqixO8pW1uwUExteExCx2gvKBa0n5pddXNBnmS5wv\nZMy2iYqjAR8r14hvI5TSjHlEMTITf4GS0DYWL+pkyVVi7Im5xjpNJ3rEfRSKQGhx\nKfdsZ5G6dAFisxgiZnqIkkCbB+ZJY3Fl6RaReqvKJHBFcJlnRtWrfsd1Up2A5RYd\nFZuetE3D\n-----END CERTIFICATE-----\n",
      "-----BEGIN CERTIFICATE-----\nMIIFojCCA4qgAwIBAgIQSkpzWeMzg55AjLh7Mx11ajANBgkqhkiG9w0BAQ0FADBi\nMQswCQYDVQQGEwJVUzEdMBsGA1UEChMUSnVuaXBlciBOZXR3b3JrcyBJbmMxNDAy\nBgNVBAMTK0p1bmlwZXIgTmV0d29ya3MgUm9vdCBDZXJ0aWZpY2F0ZSBBdXRob3Jp\ndHkwHhcNMTYxMDI4MDE0ODQ5WhcNMjYxMDI4MDE1ODQ4WjBiMQswCQYDVQQGEwJV\nUzEdMBsGA1UEChMUSnVuaXBlciBOZXR3b3JrcyBJbmMxNDAyBgNVBAMTK0p1bmlw\nZXIgTmV0d29ya3MgUm9vdCBDZXJ0aWZpY2F0ZSBBdXRob3JpdHkwggIiMA0GCSqG\nSIb3DQEBAQUAA4ICDwAwggIKAoICAQDTWXp1qN5eglLo58vjRw4jwz7OdESkxw5U\nNL2SBlE5TMPXarFVtfRnms9OYOwhNZvTd9Jh4lnZ1h7RMEFLw/9C1RysVKINxbXa\nhC1WSZeS8gnwcXnzhxRPdjZqnjl0DrYhMMuNhA+EzzIOArti+lJAvD29JnX3ItIC\nD+vwLptct0ZX8Z4gDv/INggLDUl1gXf0bfB3rAIZsDI+4ArWywt9xDZPRIL0refs\nowBL5pQ8OdMpRAwFSHhekSA+M3auFBdPsPdoU8kSQvbVvtZqPrRaSw601rRNIgdP\ncywVDA5OnlKQ4qxaAsREe0v1QN5NElEV97wXW2A6ujIq65dSCgFIODX6/wdxMXOd\n9vuKwuB5FiTPHFX1E/kb/p38KEYg672Ma6hjOWCOPDQKXsoMje8Ms3MDQPncllIZ\nkgirLBWfHnFpQbTNBZg2Vvnhzs0jrUz0WdSvmYHS1no2dK41erJmFzE8niL8xWE7\nrmFfX9ODzfg2wkQIcDmuQ2kaQUx+Kj7aGBLrY12Wy6QipOv1w6EI5I/7P6I1X5ku\ntCsVeGKNt/pNFXKEKqmQx0sjrewZqUOxYnMl3CgGfMT1m7M0hnw9uErO0J/DxcSP\nAb1rXvKVI65B0XUxIxiX6cuO2TKA0xgkoPfe/3QWP8vxkU/tFPsC3DFY+bJ2bV47\nSriO3YOEkQIDAQABo1QwUjALBgNVHQ8EBAMCAYYwEgYDVR0TAQH/BAgwBgEB/wIB\nAjAdBgNVHQ4EFgQU/BSnNM2Ql6qBlZr4NLqpBWPFLOIwEAYJKwYBBAGCNxUBBAMC\nAQAwDQYJKoZIhvcNAQENBQADggIBAHqQ+fMN2O2zZaj5Oj7AGQz8sOO/s2itxjdM\nNAJM3smFH1uXg6mre/qUaJKOLHcCJ6wa8Z+CYLz+3L1XLPPpcm78qVp6DtMBkQgu\nLdDOahXet4mfKIe96z1AouwSyJKhMStFiCzVeA229quuxoXDKELSpJLbhtaR3Kde\nLbM1RjMbFbVbg3xLm2k8oOGtmGFSbSZhy7sT4WLMHYALct1GgEoForkLSbveVfiX\nIgaIa6NXhOqYn7lcYxorVP5Z8AZzbdwMZ1UfoEV/sZkfJHUfsgh/ctP/EB0qXK6k\nOIWagncZpu8WHDYk8S/sT0ufbcISMaBjWuEuRz9GRcoN47EKHiTpDtBKbm5yZRMd\nz0pXLfQUvkityW80LCIJo+fRPduuKfI6o2OUie3ObZ8tbGuXIMZTzwRju9wOBk4H\nSoFV6QIebM808EEpUn1sUPKBbYyqvFMaBDxwzp1wIcPY2oOLf+K9MSZyaU48DJux\nFD4T/ZIjAQ9sR0uzJWM46sgjTc1Ctj+iAcI6DaFXNQGdwqKeA+KJJbEmoP70TIMp\nsqIlOKETn1IGunweyZ5dcl5Fj13ieiE0JpldD0R2uHcuwvjgqGrFBj29LhIDFkvI\nvIrpVyMmoOAhmQgNDgtB/L2LPNqnJPTZTxgFQeK+YnVwWOe2SMxuEzeijf/tH3TB\n6qsgP6NI\n-----END CERTIFICATE-----",
      "-----BEGIN CERTIFICATE-----\nMIIFWzCCA0OgAwIBAgIQP9LU9jzqnLlJDLkiViyp4zANBgkqhkiG9w0BAQsFADBA\nMRMwEQYKCZImiZPyLGQBGRYDb25lMRQwEgYKCZImiZPyLGQBGRYEc3RhZzETMBEG\nA1UEAxMKc3RhZy1EQy1DQTAeFw0xOTA5MjUxMjAxMjBaFw0yNDA5MjUxMjExMTla\nMEAxEzARBgoJkiaJk/IsZAEZFgNvbmUxFDASBgoJkiaJk/IsZAEZFgRzdGFnMRMw\nEQYDVQQDEwpzdGFnLURDLUNBMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKC\nAgEA1LBkXVV7A2rOhqm4X1L9ImC/0ofpF1RVHlGlDwfhIem7cqOLKpth/E34jyy+\n3azPwjv4uDWnYB6ZpUezvlHVt6W9h76C5zASP1xSnoAsNW2daJo8XmX9ajovz3R0\nUzG8Ds0ewDqkCRYIF/lyBo01ycbHV1Y5BNTzgqzX9jvQ7ZievURhwqjeikibXpA9\n5R8fyqsWz+Zj3oOuL02Ovcjs8k1+8jdzRPI9dI+jr/1LJtB0pXkQFAekuGLWo08z\nbecU4emmW0/sk1GMaLxGbUoNtnv7OYbsAefYOb6Evu4sqSI2UHADxYQiMWBII7wX\nGg5vzMfhWffp57ZxXkGGdNo9E7ozC7cKFVIDBbHrmqdAFWQ++TRTkv30ROU6w18w\nM6+3hGduqOdfaKdGXy8PBnZl/P54pUW9bPy57e3eEFWBKXh5+SMN3mQCGn03zojt\nZvS4VuowXBZIAG0zHQxByiARvF9RTcu+07h5Fz+rOmUxou1VsyINLY8w35Xr+MGl\nVbDiRVqD/VCaf0h/fyc2O5WVCQGjoZ9KgVWKC8JA58XwkodRlRcBtKcUV6cWrBKF\nu30ytBK45MTHO6jTS7jcHDtevJUK5JcecgcN3Sg/+vad4dw3eUdZC2mjQdk/Scv0\nb4LsrUhGS0EqTcr2NjKgdMjT3zxIgMwX4CbMxIG7MwuozvkCAwEAAaNRME8wCwYD\nVR0PBAQDAgGGMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFHUMeFVdvceaasAc\ngfeHRh/VvcMbMBAGCSsGAQQBgjcVAQQDAgEAMA0GCSqGSIb3DQEBCwUAA4ICAQAk\nF8gh7eH25xaDTeAoRUHTmVrWKfkdZcVm2mx+XG9NWNbw18fp0yn6mUfmYEUpcHez\ncXr9/ZyxKK12arnvamv53iX/ltYsOO335kDvvQFsT9hZc2oAvnUAuPFL8ncVsJw9\nIWcCduI1/71s7YQgUWfnIKPzNhhe0IW4HxQPyhyar5hoK0KDzdCf+RMGkyesskge\n2i6B/7kYK6F50B99bb05jRROLlhaBUOVaNARSKvSUOPjCFxbI7PIB5lSkKGR++RU\n9rSe9e/qTnFhNr+L55zCh/JEKUTMIjY4EKyc2Y0kHny9iVXRXVsoosvoeocz+WtR\nlXV96XHrNSkouobcVeWb7WDY4eXwb5sB2dPNeCVPhq95yBR8Ic0OqsQvJztI8F/1\nWhulfW9emwZk+bVBlsoFqhfLQtHfyXwhp4zKNYzYv09yzXQd8uXP6n4vrhXP6K4b\n1NNHWRntF9Cmuxz+HyqBsjHLEcmBy5njvLCYjYf4V2Fh48Xd9cieS5AxWdEJSW7k\noTN4bRa7qMtygGwMsZZw2FGbVaX+drNWw0njUlAPFpBckepWEmk6IwrwWCcbb+PB\nTo9iU9gjWwe2b0JGd2Meed9IcF80+oGMhxdEiFu6ZBo+E9MlegOQtluKypzlKSf5\ndSpW/jfrsM2iR+fp1ixqqZKBmcABUCryN2BGmDH93A==\n-----END CERTIFICATE-----",
      "-----BEGIN CERTIFICATE-----\nMIIF1TCCA72gAwIBAgIBATANBgkqhkiG9w0BAQsFADBbMQswCQYDVQQGEwJVUzEN\nMAsGA1UECgwETWlzdDEOMAwGA1UECwwFT3JnQ0ExLTArBgNVBAMMJDM5Y2UyMDg4\nLTFkYmUtNDM0Ni05ODdhLTFhNWE4OGJhYjVlZTAeFw0yMzA3MjAxODE5NTJaFw0z\nMzA3MTcxODE5NTJaMFsxCzAJBgNVBAYTAlVTMQ0wCwYDVQQKDARNaXN0MQ4wDAYD\nVQQLDAVPcmdDQTEtMCsGA1UEAwwkMzljZTIwODgtMWRiZS00MzQ2LTk4N2EtMWE1\nYTg4YmFiNWVlMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEArreBVg9a\nRjLCJIxuQqgTBvmVRoE2WOWIayfsZGJsJTnPhLeb1BuLXcwrsFZGY1nn/hgqeXvL\ni5ZJVyPCjH9YyXGodsQiVAhAxt4IBXKhItjxoTRok6g2WhJx6S2NC3J9iVNZsSAR\nC5gM5kxOnZ3QgWKAV6ecZkAtzzDMWiwLqhtRKDBi+pzqt+A5j9OLYiyKdj8aMAcI\nHYXdEYCdn0YibDMQNLnc9ADM2a3wAq4DIWqtXDH7DEGU5CMAo4RACPP9VZLLUVMm\nN6HyLVX1ff07Nc2AxXsYZo8L8oDeWnPCdY4FBCMtIyZ8TY2ZnpiDVx4vTDXhwp8Z\nRWj5V4ij7b0L7v3Ej8zIGZj+gsOYQEXkVArzNu20XERiPqX2H4O+ayCrzN874WfM\n5KVGpsyrZeUbeb3aBtBpm37JrArp6zZpfxTI+IJag/OYA9XK1xi0BRclDf+CENho\n462rWBB98t3yvXuog8z3iHIUtWcPon/+A+tWwGUGCa2/QabOjPfb7DREgIKdZpXz\nMAke70ILjEH+iFO5mh1xhCTbkhzIHaXH9QHPVd/caQd16OEK7GdMPjnYnVZPiav2\nPHs4zgAKareKdhVP+KrWJT5mmSTMvyA0i9a1cBtBvmX29DZyq+ZAQ611D6WzPlzV\n5YAPULcD8tA/J/o6V7ejSJ4ekTWTZOc/K/kCAwEAAaOBozCBoDAvBgNVHREEKDAm\ngiQzOWNlMjA4OC0xZGJlLTQzNDYtOTg3YS0xYTVhODhiYWI1ZWUwDwYDVR0TAQH/\nBAUwAwEB/zBcBgNVHR8EVTBTMFGgT6BNhktodHRwOi8vYXBpLmV1Lm1pc3QuY29t\nL2FwaS92MS9vcmdzLzM5Y2UyMDg4LTFkYmUtNDM0Ni05ODdhLTFhNWE4OGJhYjVl\nZS9jcmwwDQYJKoZIhvcNAQELBQADggIBAKXeTdOJ95kGImhkhokFstGNfTSL2mEr\nXTfqpri6yLmKOwn4U+ly1eZD58TxiyFmA/x7Q3r6UvpRL0tX2kntJ+7j8QLbxhnw\n4Mtwz5D6B2fBW9DDyq7VktpWgxIOaQrLnfJRvtwVdUCyEI6759CQOo3uNmBkqLaP\nALPos+Zq6VGNAbWdXEob9NSq+PrPtetQYuKH1u/fi/u36aZXSGEOHrreixY6RW/h\nxJPeTtn+wyYuK2JDXgANXJbsKUsKZtSGIqbS2SD2o+ePbwKkTtXVgEnqAXe0X0t4\nADcQ8f0FVDcRBs2RRNEMk1ty4hF0k6GrLxU3KRFvXvgeCS2a93wEWDMlUvaXm63q\nMNcR7KZZeAgiwlM5EDPnxftO0wFsgUZvFeqoPS+GpsR1p5g6s852WwQmkguW5Gqe\ntfObJOBx55XuwfmXR4wDwJKf9DvwCEn3a0fkKmslcRxBMaJydEDvUAxU025Ax5Gn\n1mRxjo8Y2CyBS9p/l+/chYeXr9ieMy+QBmEALCvQ2o+1OlABzXyn9QMsqYd+sWzs\nlnCilZL4DzabYqH3i/J7ryHQbhtDoeELK46EjS2sINohWRoJodcQCAjRkGBIwv0C\n279A891PzgljfdrGcuQ0WCtBJ1KZvUAqxvRR2wcSrS+5rmAyyc6zBUI/LvjtDExp\nbo5YxN39uE+y\n-----END CERTIFICATE-----\n",
      "-----BEGIN CERTIFICATE-----\nMIIF0jCCA7qgAwIBAgIBATANBgkqhkiG9w0BAQsFADBbMQswCQYDVQQGEwJVUzEN\nMAsGA1UECgwETWlzdDEOMAwGA1UECwwFT3JnQ0ExLTArBgNVBAMMJDIwM2QzZDAy\nLWRiYzAtNGMxYi05ZjQxLTc2ODk2YTMzMzBmNDAeFw0yMzExMjMwODU1MTRaFw0z\nMzExMjAwODU1MTRaMFsxCzAJBgNVBAYTAlVTMQ0wCwYDVQQKDARNaXN0MQ4wDAYD\nVQQLDAVPcmdDQTEtMCsGA1UEAwwkMjAzZDNkMDItZGJjMC00YzFiLTlmNDEtNzY4\nOTZhMzMzMGY0MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA0CLyOHkQ\nIeO9IkU0Mjb4XMkn9rHnOFM1orDQHss7nrH74FiQnPAb+oZhMOspGLd4r1Z4QFw1\nB1BMuRd7YzwuDl5evlEoNNFPaAslFHPyFPdqLHTNPyBSnmQM/7cCoqBT3JgMolQk\nM/XgNVHvWCW8dbtNmrWMdtcH9xZ3q4P/Z3+PHeMMZqgi0e79yRFUE/B77wCbubZ0\nEtApXK0HstK0a5dZVV2nph376UW8l5oCzHlUoMI1d3QogFtPsYUxCcqojabLU3h2\nVoRLpRjzPX7norDzSNFqYgM2355uih6PFsjy9Zl0U73E2Rzxu7hITosJbuuTEolZ\nAJ41vHB9vulx1fhNU0aA91BJ4t3e2BCHu1V8LwVtuBYrVeSRJwj19APmzEndzwv1\nWVyWdR9AzMNChz5jAzIJZShVMG3KiB0gBCoZuhbvopVerZCpMAdE8THhtWrRUpv4\niIPe4jjfJEjazWSYNaqhsnonMr9el5x4BPlpoQSO+LkdJW7hjlbanrDyu1EVsIwg\nTGPDttXjjU5IrMiPjAZy5bBnzwfjB0fJ7aGrAmxK7jDnuBjggQ6apNf68YkoVXNJ\nomV4vFawVQ/qykyXxoxom3MfBn/lcLod5+ALX3ukjdkfjYR8M7F1rERekklwC34z\nfhzaDtYSRyVn9+HOATjBKWMzGXGC6zicyvsCAwEAAaOBoDCBnTAvBgNVHREEKDAm\ngiQyMDNkM2QwMi1kYmMwLTRjMWItOWY0MS03Njg5NmEzMzMwZjQwDwYDVR0TAQH/\nBAUwAwEB/zBZBgNVHR8EUjBQME6gTKBKhkhodHRwOi8vYXBpLm1pc3QuY29tL2Fw\naS92MS9vcmdzLzIwM2QzZDAyLWRiYzAtNGMxYi05ZjQxLTc2ODk2YTMzMzBmNC9j\ncmwwDQYJKoZIhvcNAQELBQADggIBAKnfUNKE2U09i4RfEltOJc8cGK85yS4GhRTh\nPCtSKxY0iYiIs+m+2d6vkWbZXEeGjaZPdjQh8ixoZMAu1L6rtWqp2qzpmSxXn++w\nZkuMTHGn8vFXUSbb6CTFdB3SHzGsp17JJVfFhvshoZbpRP+d4cSACJvp0cy67JFN\n9QGvsPTyIZSClhUWYX7iYe4VW11s2gS982FMGwQdr5OR1DHmvNHGBuP/uHhW95D7\nTGqe43c/0AdVjQn1VYnavVScqKERCYh/uIZ+hSbpY7INR0t/nj4n+XBjpjIiQPVJ\nkeumCKVHaE0GHhV2Q1DbngU+uuk/K023xQima9kNn0TbgtFjVuH4Ip0Y8QSH5WVR\naV43oc+keqnirXcJqjwF1z5hVpQZKG8P9ZcEoJkA1JHo04J64lJ2auBUHFQlWJ1g\nEmoS+VT6/wR2Eiflc65ZONZ688ckbLRyioak3tDxUV/nJQfFMuaQASJGjaOod4iJ\nCN0b2g6tEKuA6nMLiKVpQpfo3bHHipZ2dtj9VFA48KPHftSpjJTsmWJvCdMYTr3f\nYaN6uRgV45iqUfO0fcxiGcS9W2EfIlR9Jsv2uoBYSqbXPFc8t1hOsezyOvjNqcR7\nV02f1S5kOfmTdbnp2e4aFpUFy73tZQ+MXMFe+HdabgVlRoVmGSKjE8Wu0/nXOhKA\n90aXHpbz\n-----END CERTIFICATE-----\n"
    ]
    server_cert = {
      cert = "-----BEGIN CERTIFICATE-----\nMIIGXDCCBESgAwIBAgITbgAAAPhdWS1GxAK1qQABAAAA+DANBgkqhkiG9w0BAQsF\nADBAMRMwEQYKCZImiZPyLGQBGRYDb25lMRQwEgYKCZImiZPyLGQBGRYEc3RhZzET\nMBEGA1UEAxMKc3RhZy1EQy1DQTAeFw0yMzAyMDYwODE3NTBaFw0yNTAyMDUwODE3\nNTBaMF0xCzAJBgNVBAYTAkZSMQwwCgYDVQQIEwNJREYxETAPBgNVBAcTCE5hbnRl\ncnJlMREwDwYDVQQKEwhzdGFnLm9uZTEaMBgGA1UEAxMRbWlzdC1uYWMuc3RhZy5v\nbmUwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCvi1GzzI0FjjqHQinL\n/cFXSVH8a2IQoGF7zdmPkDWlEil7Sbhz7DkPyZhNe6jSpr/fMPjstMs/eyt9CCFW\nqvtByR2iBTkN7I3LeURM213Q1HotFK5Cr6P0xMfdGOQXQwjQ+6X3twiJo3+e2uJ/\n+hv4zjvUX3rFrFQwjE6Os4SH8B1AgJaQ+jq65Q6u0bGyznZozjRE2O1V7lQM1K3q\nEbWBOQb66CPB/1xXK3NDbdHtPbP8pihvaU7aSmDsuOPVQggRWvP7Y87eUf6m9+Dn\nIwww9huZabHsy2+tHtQCpIfY6bMqlstpVfik0vVfh+qE8O6nfzvySrMrWIlozMSu\npWPlAgMBAAGjggIwMIICLDAOBgNVHQ8BAf8EBAMCBaAwHAYDVR0RBBUwE4IRbWlz\ndC1uYWMuc3RhZy5vbmUwHQYDVR0OBBYEFPR07EguUUKQx6Fzb6c53q7HUJ1wMB8G\nA1UdIwQYMBaAFHUMeFVdvceaasAcgfeHRh/VvcMbMIHHBgNVHR8Egb8wgbwwgbmg\ngbaggbOGgbBsZGFwOi8vL0NOPXN0YWctREMtQ0EsQ049V1MyMDIyLURDLENOPUNE\nUCxDTj1QdWJsaWMlMjBLZXklMjBTZXJ2aWNlcyxDTj1TZXJ2aWNlcyxDTj1Db25m\naWd1cmF0aW9uLERDPXN0YWcsREM9b25lP2NlcnRpZmljYXRlUmV2b2NhdGlvbkxp\nc3Q/YmFzZT9vYmplY3RDbGFzcz1jUkxEaXN0cmlidXRpb25Qb2ludDCBuQYIKwYB\nBQUHAQEEgawwgakwgaYGCCsGAQUFBzAChoGZbGRhcDovLy9DTj1zdGFnLURDLUNB\nLENOPUFJQSxDTj1QdWJsaWMlMjBLZXklMjBTZXJ2aWNlcyxDTj1TZXJ2aWNlcyxD\nTj1Db25maWd1cmF0aW9uLERDPXN0YWcsREM9b25lP2NBQ2VydGlmaWNhdGU/YmFz\nZT9vYmplY3RDbGFzcz1jZXJ0aWZpY2F0aW9uQXV0aG9yaXR5MCEGCSsGAQQBgjcU\nAgQUHhIAVwBlAGIAUwBlAHIAdgBlAHIwEwYDVR0lBAwwCgYIKwYBBQUHAwEwDQYJ\nKoZIhvcNAQELBQADggIBAEArEXJUjR7dXH53DAP3d42gc5fiHewn+XAHWFZJQNyy\nmOqGaDEleoK6indtgF17jl8rRZ+54GxaMoc1/snXiKsKfRNxdHohbEEMVje+M0OQ\nvQvDirP7gcHcUzvjMXLD6ElMGcAzOkzeJAntp0XZKFhk1rhEQMtxxEpqtx869ivK\nEVewgt7uLPskBJD+PUUpxPniA4AsD1q9R2ntBDbwcp/d/sI5rmmmJHVcogU3kPN/\n9OAvEgsr5xNaxn/Cler+4dblf6UbLVP1AbU3glr/v26H6sCwRL2ZAY+u+M94j9Y+\nxwewmo/tqxyldd0ot42BA9zEWycz694B6rbhhPMeuZ3+O1N4RiX98s8U+xV6NpDf\nTJ8yByf8C4V5kUnFFN6EUWvdnfNw/6k11nqGrp8ZQgBf5vzURGGNc5RM72zWp2qc\nkURgmvtCepGeDYB/tXzA1FgFMQw1wm7ZgICeZ7QJFJiywQMECO9c5AfEickvfGth\ntxBSGPZLJnZ9VBsoML6z53uxG3N/ysclvoNgfmRBlyGc6jbafsMhCP7mQcZx2u1x\nbx+D/ZPMFtc0VJVlrZN/ghla7GAtZ8WI2W+CuFh9q5DY3jqOZIoiFecq92FFNsVF\nRZZ3q+c22tC5m22j6QIDX3ZlJKkwt+YPHiN3hkSk/8yguPAUbpVEvk4+6F79nUb/\n-----END CERTIFICATE-----"
      key  = "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCvi1GzzI0FjjqH\nQinL/cFXSVH8a2IQoGF7zdmPkDWlEil7Sbhz7DkPyZhNe6jSpr/fMPjstMs/eyt9\nCCFWqvtByR2iBTkN7I3LeURM213Q1HotFK5Cr6P0xMfdGOQXQwjQ+6X3twiJo3+e\n2uJ/+hv4zjvUX3rFrFQwjE6Os4SH8B1AgJaQ+jq65Q6u0bGyznZozjRE2O1V7lQM\n1K3qEbWBOQb66CPB/1xXK3NDbdHtPbP8pihvaU7aSmDsuOPVQggRWvP7Y87eUf6m\n9+DnIwww9huZabHsy2+tHtQCpIfY6bMqlstpVfik0vVfh+qE8O6nfzvySrMrWIlo\nzMSupWPlAgMBAAECggEAXMY4q9mTPjho3kstY838fHEXj1wBO+BHJPMp5CHG+DKd\nEbdocHuLvLhCcdDE7G+MQBzVsePq9uXVmgLN+8vpdl8f9hzkhmNanQq9+hCFiHSD\nDYg8gKnq8NV++27IPTMEWc5lbIGjVXq/W+k6g4BfgCMeo6LCc+4boHzGl8xwYpa3\n7QTzsEvwJxm8A6SWEMjxkIXF0sTeYHmHcnGvPro848PTFwYI7nX+2jlaR1+U5Hxl\nH6BUVUkAyWz1jVIt+C3TG+jCtFex/3JORWWil44XmF2Zbbc+xnsDbpUzGlc7CJFO\nRT0JdGpPpuXRt/rtJu4Vj64IxjZauSWlNpB8vYFHAQKBgQDbFRbxphFeeI2c+ffI\n2uxlh9fvRCvVL0C7NQKDVewO0WYR+I46sYTTRxHso0/x7WRVDCmy7PiazFHDSkA5\njCbwGuR5J8rxsUDLNGOLPgvu3fH/UARAkTE7usjtsWcGHGJbI+U6vUgaNH5YDNeC\nxg1wE6xlb9xjR9vL2cTZYNahkQKBgQDNIA2p2H5dhdLaeSfjbD7hpzopgOhVVwRX\nF2+GuoLP4pTKiIlkOhWA3zLFNVhaEjqRQhpz3x2NZUrIi63T3ZSyuH8QU5+tFKyU\nc4a0lWy9bZX4EzqZnLC/IAIZ4aVqjyyPO+pACeTcvqLlp4FUfEFyePPMku809X1I\n5StucBZzFQKBgQDJhNWtA9+4MVW5zijBRbbrEZBr0x4qo6N8UB92pYNUggLMhwB1\nNrMCEL6JGBPCamW+4Ug+AEIRSd3vsA3xoVxzwZjlZNgFh9Ec2ZyzCSi62Mbv3S98\nINHTqlWGZodolJVYVmVVqaR4Dk8lOPLHpNvZew854zIq1H92PGMYvT3QMQKBgB6h\nHKkx9JyOcQ/Ca5xf+3vCdsLfxtGHbtR6YWLi0smr8b/UZ3+Z1MArk+8yqgC5HBm4\nny6SMF1+tI4BnXE3cw7E0eRkOEGmBqccIQ8yCMy4Yl5qsKNjbFj9tJKcyRCCVIsG\nzVHGqG/132kffAZbj4JCYVU405M63PxXjAavogqJAoGBANS7ojTfeJfCepAYAQNL\ndr6RaXoJ/u/re7ZIuuY14DAcFdYFsX9g/JEHJhbrOPpy2Ez6eo7hJ6ZU5OK6wl7f\naw1YujZMkzyJD8qpH7b7dozhNvki54b5TBhujBWFMkgRGh4H3h+OSkoCBibmJpI1\n9bYtllCav7gHDyGRk3k50pay\n-----END PRIVATE KEY-----"
    }
    eu_only = true
  }
  synthetic_test = {
    disabled = false
    vlans = [
      {
        vlan_ids = [
          "8",
          "999"
        ],
        disabled = true
      }
    ]

  }
  device_cert = null
  mxedge_mgmt = {
    fips_enabled = false
  }
  switch_mgmt = {
    use_mxedge_proxy = false
  }
  cradlepoint = {
    ecm_api_id            = "2e0c3b2e-d56b-4db1-8842-549e8146fb03"
    ecm_api_key           = "9b4824dc007b07f299f5a08708cd0d6d1c030337"
    cp_api_id             = "4bab0b43"
    cp_api_key            = "e5db30aa5bdd392153771edfdd54fc88"
    shared_secret         = "R99CCyvvD0ANxJufHuaR5h49ZPb93hWuCsNmeV8tZMxKPzpwhbWbpI39rQhVwxUh"
    destination_config_id = "d79a7abe-12ba-11ef-8c0a-fe633ad12f0e"
    alert_config_id       = "d8ec73c0-12ba-11ef-99ca-7678474402dd"
  }
  api_policy = {
    no_reveal = false
  }

}
resource "mist_org_wxtag" "test_1" {
  values = [
    "10.3.0.0/16"
  ]
  op     = "in"
  name   = "Internal Network"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "ip_range_subnet"
}
resource "mist_org_wxtag" "test_2" {
  values = [
    "53"
  ]
  op     = "in"
  name   = "DNS"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "port"
}
resource "mist_org_wxtag" "test_3" {
  values = [
    "443"
  ]
  op     = "in"
  name   = "HTTPS"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "port"
}
resource "mist_org_wxtag" "test_4" {
  values = [
    "80"
  ]
  op     = "in"
  name   = "HTTP"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "port"
}
resource "mist_org_wxtag" "test_5" {
  op = "in"
  values = [
    "www.myapp.com:443"
  ]
  name   = "my app"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "hostname"
}
resource "mist_org_wxtag" "test_6" {
  op = "in"
  values = [
    "test1"
  ]
  name   = "test1"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "radius_group"
}
resource "mist_org_wxtag" "test_7" {
  op = "in"
  values = [
    "10.3.20.0/24"
  ]
  name   = "Subnet SRV"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "ip_range_subnet"
}
resource "mist_org_wxtag" "test_8" {
  specs = [
    {
      subnets = [
        "10.3.20.105"
      ]
      protocol   = "tcp"
      port_range = "443"
    }
  ]
  name   = "IP traefik.stag.one:443"
  org_id = mist_org.terraform_test.id
  type   = "spec"
}
resource "mist_org_wxtag" "test_81" {
  specs = [
    {
    }
  ]
  name   = "test_81"
  org_id = mist_org.terraform_test.id
  type   = "spec"
}
resource "mist_org_wxtag" "test_9" {
  op = "in"
  values = [
    "10.3.20.105"
  ]
  name   = "IP traefik.stag.one"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "ip_range_subnet"
}
resource "mist_org_wxtag" "test_10" {
  name   = "ma station"
  org_id = mist_org.terraform_test.id
  type   = "client"
  mac    = "9a4fccaa8f82"
}
resource "mist_org_wxtag" "test_12" {
  op = "in"
  values = [
    "iot"
  ]
  name   = "iot_label"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "radius_group"
}
resource "mist_org_wxtag" "test_13" {
  op = "in"
  values = [
    "station"
  ]
  name   = "station"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "radius_group"
}
resource "mist_org_wxtag" "test_14" {
  op     = "not_in"
  name   = "AP_test"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "ap_id"
  values = [
    "0000000"
  ]
}
resource "mist_org_wxtag" "test_15" {
  op = "in"
  values = [
    "test"
  ]
  name   = "test"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "radius_group"
}
resource "mist_org_wxtag" "test_16" {
  values = [
    "test"
  ]
  name   = "VIP"
  op     = "in"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "psk_role"
}
resource "mist_org_wxtag" "test_16_1" {
  op = "in"
  values = [
    "10.3.20.164"
  ]
  name   = "vcsa.stag.one_ip"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "ip_range_subnet"
}
resource "mist_org_wxtag" "test_17" {
  op = "not_in"
  values = [
    "ringcentral",
    "gcp",
    "gsuite",
    "teamviewer",
    "ms-teams"
  ]
  name   = "Allowed Applications"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "app"
}
resource "mist_org_wxtag" "test_18" {
  op = "in"
  values = [
    "0.0.0.0/0"
  ]
  name   = "any"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "ip_range_subnet"
}
resource "mist_org_wxtag" "test_19" {
  op = "in"
  values = [
    "manage.mist.com",
    "api.mist.com"
  ]
  name   = "manage.mist.com"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "hostname"
}
resource "mist_org_wxtag" "test_19_1" {
  op = "in"
  values = [
    "demopsk"
  ]
  name   = "demopsk"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "radius_group"
}
resource "mist_org_wxtag" "test_20" {
  op = "in"
  values = [
    "manage.mist.com:443"
  ]
  name   = "resource_mist:443"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "hostname"
}
resource "mist_org_wxtag" "test_21" {
  specs = [
    {
      subnets = [
        "0.0.0.0/0"
      ]
      protocol   = "tcp"
      port_range = "443"
    }
  ]
  name   = "internet:443"
  org_id = mist_org.terraform_test.id
  type   = "spec"
}
resource "mist_org_wxtag" "test_22" {
  values = [
    "53"
  ]
  op     = "in"
  name   = "DNS_all"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "port"
}
resource "mist_org_wxtag" "test_23" {
  specs = [
    {
      subnets = [
        "10.3.20.201",
        "10.3.51.222"
      ]
      protocol   = "any"
      port_range = "53"
    }
  ]
  name   = "DNS_internal"
  org_id = mist_org.terraform_test.id
  type   = "spec"
}
resource "mist_org_wxtag" "test_24" {
  op = "in"
  values = [
    "netflix.com"
  ]
  name   = "netflix"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "hostname"
}
resource "mist_org_wxtag" "test_25" {
  op = "in"
  values = [
    "10.0.0.0/8",
    "172.16.0.0/22",
    "192.168.0.0/16"
  ]
  name   = "RFC1918"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "ip_range_subnet"
}
resource "mist_org_wxtag" "test_26" {
  op = "in"
  values = [
    "10.3.8.0/24"
  ]
  name   = "iot_subnet"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "ip_range_subnet"
}
resource "mist_org_wxtag" "test_27" {
  op = "in"
  values = [
    ".teamviewer.com:5938",
    ".teamviewer.com:443"
  ]
  name   = "teamviewer"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "hostname"
}
resource "mist_org_wxtag" "test_29" {
  op     = "in"
  name   = "my ssid"
  org_id = mist_org.terraform_test.id
  type   = "match"
  values = ["test"]
  match  = "wlan_id"
}
resource "mist_org_wxtag" "test_30" {
  op = "in"
  values = [
    "cdn.segment.com"
  ]
  name   = "cdn.segment.com"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "hostname"
}
resource "mist_org_wxtag" "test_32" {
  op     = "in"
  name   = "MlN.1X"
  org_id = mist_org.terraform_test.id
  type   = "match"
  match  = "wlan_id"
  values = ["test"]
}

resource "mist_org_vpn" "vpn_one" {
  org_id = mist_org.terraform_test.id
  name   = "VpnOrgOverlay"
  paths = {
    "AWS_Hub_Profile1-WAN1" : {
      bfd_profile = "broadband"
    },
    "AWS_Hub_Profile1-WAN2" : {},
  }
}

resource "mist_org_inventory" "inventory" {
  org_id = mist_org.terraform_test.id
  devices = [
    {
      claim_code = "CPKL2EXN8JY98AC"
      site_id    = mist_site.terraform_site.id
    },
    {
      claim_code = "G87JHBFXZJSFNMX"
      site_id    = mist_site.terraform_site.id
    },
    {
      claim_code = "CV4YAS8DQWYLL6M"
      site_id    = mist_site.terraform_site.id
    }
  ]
}

resource "mist_org_wxtag" "wxtag_one" {
  org_id  = mist_org.terraform_test.id
  name    = "wxtag_one"
  type    = "vlan"
  vlan_id = 10
}

resource "mist_device_gateway_cluster" "cluster_one" {
  site_id   = mist_site.terraform_site2.id
  device_id = "00000000-0000-0000-1000-4c96143de700"
  nodes = [
    { mac = "4c96143de700" },
    { mac = "4c9614aae100" }
  ]
}
resource "mist_device_gateway" "cluster_one" {
  device_id = mist_device_gateway_cluster.cluster_one.device_id
  site_id   = mist_device_gateway_cluster.cluster_one.site_id
  oob_ip_config = {
    type = "dhcp"
  }
  dns_servers = ["8.8.8.8"]
  name        = "cluster_one"
  additional_config_cmds = [
    "annotate system \" -- custom-main -- Template level --\"",
    "delete apply-groups custom-main",
    "delete groups custom-main",
    "set groups custom-main",
    "set groups custom-main system services ssh root-login allow",
    "set apply-groups custom-main",
  ]
}
### SITES
resource "mist_site" "terraform_site" {
  org_id       = mist_org.terraform_test.id
  name         = "terraform_site"
  country_code = "FR"
  timezone     = "Europe/Paris"
  address      = "77 Terrasse de l'Universit\u00e9, 92000 Nanterre, France"
  notes        = "Created with Terraform, Updated with Terraform"
  latlng = {
    lat = 48.899268
    lng = 2.214447
  }
  sitegroup_ids      = [mist_org_sitegroup.test_group.id, mist_org_sitegroup.test_group2.id]
  networktemplate_id = mist_org_networktemplate.switch_template.id
  rftemplate_id      = mist_org_rftemplate.test_rf.id
  gatewaytemplate_id = mist_org_gatewaytemplate.test-api.id
}

resource "mist_device_ap" "test_ap" {
  device_id = mist_org_inventory.inventory.devices[0].id
  site_id   = mist_org_inventory.inventory.devices[0].site_id
  name      = "test_ap"
}


resource "mist_site" "terraform_site2" {
  org_id       = mist_org.terraform_test.id
  name         = "terraform_site2"
  country_code = "FR"
  timezone     = "Europe/Paris"
  address      = "77 Terrasse de l'Universit\u00e9, 92000 Nanterre, France"
  notes        = "Created with Terraform, Updated with Terraform"
  latlng = {
    lat = 48.899268
    lng = 2.214447
  }
  sitegroup_ids = [mist_org_sitegroup.test_group.id, mist_org_sitegroup.test_group2.id]
}

resource "mist_org_sitegroup" "test_group" {
  org_id = mist_org.terraform_test.id
  name   = "test group"
}
resource "mist_org_sitegroup" "test_group2" {
  org_id = mist_org.terraform_test.id
  name   = "test group2b"
}
resource "mist_org_rftemplate" "rft1" {
  org_id        = mist_org.terraform_test.id
  band_24_usage = "24"
  country_code  = "US"
  band_24 = {
    power_min = 8
    power_max = 18
    preamble  = "short"
    bandwidth = 20
    power     = 8
    ant_gain  = 1
  }
  name = "rft_test"
}

# # ### ORG LEVEL
resource "mist_org_service" "lab" {
  org_id = mist_org.terraform_test.id
  addresses = [
    "10.3.0.0/24", "10.4.0.0/24"
  ]
  description = "the lab network"
  name        = "lab_network"
  type        = "custom"
  specs = [
    {
      protocol   = "tcp"
      port_range = "443"
    }
  ]
}
resource "mist_org_network" "networks14" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "0.0.0.0/0"
  internet_access = {
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  disallow_mist_services = true
  name                   = "any"
}
resource "mist_org_network" "networks17" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "{{VLAN494}}.0/24"
  internet_access = {
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  vlan_id                = "{{VLAN494}}"
  disallow_mist_services = true
  name                   = "VLAN494_Network"
}
resource "mist_org_network" "networks1" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "{{cnf_peer_subnet}}/{{cnf_peer_prefix}}"
  tenants = {
    ANY = {
      addresses = [
        "0.0.0.0/0"
      ]
    }
  }
  internet_access = {
  }
  name = "CNF-Peer"
}
resource "mist_org_network" "corp" {
  org_id                 = mist_org.terraform_test.id
  disallow_mist_services = false
  name                   = "prd_corp"
  subnet                 = "10.4.30.0/{{test}}"
}

resource "mist_org_network" "mgmt" {
  org_id  = mist_org.terraform_test.id
  vlan_id = "{{test}}"
  subnet  = "10.3.172.0/24"
  gateway = "10.3.172.9"
  vpn_access = {
    "OrgOverlay" : {
      routed                    = true
      no_readvertise_to_overlay = false
      no_readvertise_to_lan_bgp = false
    }
  }
  isolation              = true
  disallow_mist_services = false
  name                   = "SRX-Mgmt"
}
resource "mist_org_network" "reg" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = "12"
  subnet                 = "10.3.12.0/24"
  disallow_mist_services = false
  name                   = "SRX-REG"
}
resource "mist_org_network" "ssr" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 128
  subnet                 = "10.128.0.0/16"
  disallow_mist_services = false
  name                   = "SRX-Core-128T"
}
resource "mist_org_network" "mxe" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 1000
  subnet                 = "10.10.0.0/24"
  disallow_mist_services = false
  tenants = {
    "mxe-ext" = {
      addresses = [
        "10.10.0.10"
      ]
    }
  }
  internet_access = {
    destination_nat = {
      "192.168.1.9:500" : {
        name        = "ike"
        internal_ip = "10.10.0.10"
        port        = "500"
      }
      "192.168.1.9:4500" : {
        name        = "isakmp"
        internal_ip = "10.10.0.10"
        port        = "4500"
      }
    }
  }
  name = "SRX-MXE-tt-in"
}
resource "mist_org_network" "iot" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 8
  subnet                 = "10.3.8.0/24"
  disallow_mist_services = true
  tenants = {
    netatmo = {
      addresses = [
        "10.3.8.201/32"
      ]
    }
    "cctv-e" = {
      addresses = [
        "10.3.8.40/32"
      ]
    }
    "cctv-b" = {
      addresses = [
        "10.3.8.41/32"
      ]
    }
    "cctv-c" = {
      addresses = [
        "10.3.8.42/32"
      ]
    }
    "cctv-x" = {
      addresses = [
        "10.3.8.44/32"
      ]
    }
  }
  name = "SRX-IoT"
}
resource "mist_org_network" "dmz" {
  org_id                 = mist_org.terraform_test.id
  isolation              = true
  vlan_id                = 3
  subnet                 = "10.3.3.0/24"
  disallow_mist_services = true
  tenants = {
    aws = {
      addresses = [
        "10.3.3.3"
      ]
    }
  }
  internet_access = {
    destination_nat = {
      "192.168.1.9:5431" : {
        name        = "aws-wg"
        internal_ip = "10.3.3.3"
        port        = "51820"
      }
    }
  }
  name = "SRX-DMZ"
}


resource "mist_org_gatewaytemplate" "test-api" {

  type   = "spoke"
  name   = "test-api"
  org_id = mist_org.terraform_test.id
  tunnel_configs = {
    Prisma-Tunnel = {
      provider = "custom-ipsec"
      protocol = "ipsec"
      local_id = "{{sc_local_id}}"
      psk      = "psktest"
      primary = {
        hosts = [
          "{{service_ip_1}}"
        ]
        remote_ids = [
          "{{sc_remote_id}}"
        ]
        wan_names = [
          "WAN_0"
        ]
      }
      ike_proposals = [
        {
          auth_algo = "sha2"
          enc_algo  = "aes256"
          dh_group  = "19"
        }
      ]
      ike_lifetime = 28800
      ipsec_proposals = [
        {
          auth_algo = "sha2"
          enc_algo  = "aes256"
          dh_group  = "19"
        }
      ]
      ipsec_lifetime = 3600
      version        = "2"
    }
  }
  port_config = {
    "ge-0/0/3" = {
      name       = "FTTH"
      usage      = "wan"
      aggregated = false
      redundant  = false
      critical   = false
      wan_type   = "broadband"
      ip_config = {
        type    = "static"
        ip      = "192.168.1.8"
        netmask = "/24"
        gateway = "192.168.1.1"
      },
      disable_autoneg = false
      speed           = "auto"
      duplex          = "auto"
      wan_source_nat = {
        disabled = false
      },
      vpn_paths = {
        "SSR_HUB_DC-MPLS.OrgOverlay" = {
          key         = 0
          role        = "spoke"
          bfd_profile = "broadband"
        }
      }
    },
    "ge-0/0/4" = {
      name       = "LTE"
      usage      = "wan"
      aggregated = false
      redundant  = false
      critical   = false
      wan_type   = "broadband"
      ip_config = {
        type = "dhcp"
      },
      disable_autoneg = false
      speed           = "auto"
      duplex          = "auto"
      wan_source_nat = {
        disabled = false
      },
      vpn_paths = {
        "SSR_HUB_DC-MPLS.OrgOverlay" = {
          key         = 0
          role        = "spoke"
          bfd_profile = "broadband"
        }
      }
    },
    "ge-0/0/5" = {
      usage            = "lan"
      critical         = false
      aggregated       = true
      ae_disable_lacp  = false
      ae_lacp_force_up = true
      ae_idx           = 0
      redundant        = false
      networks = [
        "PRD-Core"
      ]
    },
    "ge-0/0/7" = {
      usage      = "lan"
      critical   = false
      aggregated = false
      redundant  = false
      networks = [
        "PRD-Mgmt"
      ]
    },
    "ge-0/0/6" = {
      usage      = "lan"
      critical   = false
      aggregated = false
      redundant  = false
      networks = [
        "PRD-Lab"
      ]
    }
  }
  ip_configs = {
    "PRD-Core" = {
      type    = "static"
      ip      = "10.3.100.9"
      netmask = "/24"
    },
    "PRD-Mgmt" = {
      type    = "static"
      ip      = "10.3.172.1"
      netmask = "/24"
    },
    "PRD-Lab" = {
      type    = "static"
      ip      = "10.3.171.1"
      netmask = "/24"
    }
  }
  dhcpd_config = {
    enable = true
  }
  path_preferences = {
    "HUB" = {
      strategy = "ordered"
      paths = [
        {
          name = "SSR_HUB_DC-MPLS.OrgOverlay"
          type = "vpn"
        }
      ]
    },
    "HUB-ORDERED" = {
      strategy = "ordered"
      paths = [
        {
          name     = "SSR_HUB_DC-MPLS.OrgOverlay"
          wan_name = "FTTH",
          type     = "vpn"
        },
        {
          name     = "SSR_HUB_DC-MPLS.OrgOverlay"
          wan_name = "LTE"
          type     = "vpn"
        }
      ]
    },
    "HUB-ECMP" = {
      strategy = "weighted"
      paths = [
        {
          name     = "SSR_HUB_DC-MPLS.OrgOverlay"
          wan_name = "LTE"
          cost     = 30
          type     = "vpn"
        },
        {
          name     = "SSR_HUB_DC-MPLS.OrgOverlay"
          wan_name = "FTTH"
          cost     = 30
          type     = "vpn"
        }
      ]
    }
  }
  service_policies = [
    {
      servicepolicy_id = mist_org_servicepolicy.test1.id
      path_preference  = "TO-INTERNET"
    },
    {
      name = "Policy-14"
      tenants = [
        "PRD-Core"
      ],
      services = [
        "any"
      ],
      action          = "allow"
      path_preference = "HUB"
      idp = {
        enabled    = true
        profile    = "critical"
        alert_only = false
      }
    },
    {
      name = "Policy-2"
      tenants = [
        "PRD-Mgmt"
      ],
      services = [
        "any"
      ],
      action          = "allow",
      path_preference = "HUB-ECMP"
      idp = {
        enabled    = true,
        profile    = "standard"
        alert_only = true
      }
    },
    {
      name = "Policy-3"
      tenants = [
        "PRD-Lab"
      ],
      services = [
        "any"
      ],
      action          = "allow"
      path_preference = "HUB-ORDERED"
      idp = {
        enabled = false
      }
    }
  ]
}

resource "mist_org_gatewaytemplate" "sdwan-westford" {
  ntp_override = true
  service_policies = [
    {
      name = "inbound-hub-to-corp-net"
      idp = {
        enabled = false
      },
      path_preference = "to-corp-lan"
      services = [
        "spoke-corp-network"
      ],
      action = "allow"
      tenants = [
        "corp_datacenter1",
        "corp_datacenter2",
        "core_dc2_network",
        "mgmt_datacenter1",
        "mgmt_datacenter2"
      ]
    },
    {
      name = "inbound-remote-spoke-to-corp-net"
      tenants = [
        "corp_network",
        "mgmt_network"
      ],
      services = [
        "spoke-corp-network"
      ],
      action          = "allow"
      path_preference = "to-corp-lan"
      idp = {
        enabled = false
      }
    },
    {
      name = "outbound-to-hub-workloads"
      tenants = [
        "corp_network"
      ],
      services = [
        "datacenter1-mgmt-network",
        "datacenter1-workloads",
        "datacenter2-lan",
        "datacenter2-mgmt-network",
        "datacenter2-workloads"
      ],
      action          = "allow"
      path_preference = "broadband-overlay"
      idp = {
        enabled = false
      }
    },
    {
      name = "outbound-corp-internet"
      tenants = [
        "corp_network"
      ],
      services = [
        "corp-internet"
      ],
      action          = "allow",
      path_preference = "internet-breakout"
      idp = {
        enabled = false
      }
    },
    {
      name = "application-test"
      tenants = [
        "corp_network"
      ],
      services = [
        "youtube"
      ],
      action = "allow"
      idp = {
        enabled = false
      },
      path_preference = "internet-breakout"
    },
    {
      name = "url-filtering"
      tenants = [
        "corp_network"
      ],
      services = [
        "blocked-categories"
      ],
      action        = "deny"
      local_routing = true
      idp = {
        enabled = false
      }
    }
  ]
  dns_servers = [
    "8.8.8.8"
  ]
  port_config = {
    "ge-0/0/1" : {
      name       = "wan1-broadband"
      usage      = "wan"
      aggregated = false
      redundant  = false
      critical   = false
      disabled   = false
      wan_type   = "broadband"
      ip_config = {
        type    = "static"
        ip      = "{{broadband_ip_var}}"
        netmask = "/{{broadband_prefix_var}}"
        gateway = "{{broadband_gw_var}}"
      },
      disable_autoneg = false
      wan_source_nat = {
        disabled = false
      },
      vpn_paths = {
        "sdwan_newyork_hub-external1.OrgOverlay" : {
          key         = 0
          role        = "spoke"
          bfd_profile = "broadband"
        }
      }
    },
    "ge-0/0/3" : {
      networks = [
        "corp_network"
      ],
      usage      = "lan"
      aggregated = false
      redundant  = false
      critical   = false
      disabled   = false
    }
  }
  path_preferences = {
    "broadband-overlay" : {
      strategy = "ordered"
      paths = [
        {
          name = "sdwan_newyork_hub-external1.OrgOverlay"
          type = "vpn"
        }
      ]
    },
    "internet-breakout" : {
      strategy = "ordered"
      paths = [
        {
          name = "wan1-broadband"
          type = "wan"
        }
      ]
    },
    "to-corp-lan" : {
      strategy = "ordered"
      paths = [
        {
          type = "local"
          networks = [
            "corp_network"
          ]
        }
      ]
    }
  }
  ntp_servers = [
    "pool.ntp.org"
  ]
  routing_policies = {
    "ibgp-hub-preference" : {
      terms = [
        {
          matching = {
            prefix = [
              "172.76.128.0/24",
              "172.36.128.0/24",
              "192.168.94.0/24",
              "192.168.93.0/24",
              "192.168.92.0/24",
              "192.168.95.0/24",
              "192.168.91.0/24",
              "172.46.128.0/24",
              "172.66.128.0/24",
              "172.56.128.0/24",
              "24.0.0.0/29",
              "192.168.90.0/24",
              "192.168.97.0/24",
              "192.168.98.0/24",
              "192.168.94.0/24"
            ],
            protocol = [
              "bgp"
            ],
            vpn_path = [
              "NewYork-newyork-broadband1.OrgOverlay",
              "NewYork-newyork-broadband1-lte.OrgOverlay",
              "SanFrancisco-sanfran-broadband1.OrgOverlay"
            ]
          },
          actions = {
            accept = true
          }
        }
      ]
    }
  }
  ip_configs = {
    corp_network = {
      type    = "static"
      ip      = "{{corp_ip}}"
      netmask = "/{{corp_CIDR}}"
    }
  }
  dns_override = true
  dhcpd_config = {
    enabled = true
    config = {
      corp_network = {
        type     = "local"
        ip_start = "{{corp_dhcp_start}}"
        ip_end   = "{{corp_dhcp_end}}"
        gateway  = "{{corp_ip}}"
        dns_servers = [
          "8.8.8.8",
          "1.1.1.1"
        ]
      }
    }
  }
  bgp_config = {
    "bgp-group1" : {
      via           = "vpn"
      vpn_name      = "OrgOverlay"
      import_policy = "ibgp-hub-preference"
    }
  }
  type   = "spoke"
  name   = "sdwan-westford"
  org_id = mist_org.terraform_test.id

}


resource "mist_org_rftemplate" "test_rf" {
  band_24_usage = "auto"
  band_5 = {
    ant_gain = 2
    power    = 8
    channels = [
      60,
      104,
      132
    ]
    bandwidth = 20
  }
  band_6 = {
    ant_gain = 2
    power    = 8
  }
  band_24 = {
    ant_gain          = 1
    allow_rrm_disable = true
    power_min         = 18
    power_max         = 18
    bandwidth         = 20
  }
  ant_gain_5   = 2
  ant_gain_6   = 2
  ant_gain_24  = 1
  country_code = "FR"
  name         = "tf"
  org_id       = mist_org.terraform_test.id
}

resource "mist_site_setting" "test" {
  site_id = mist_site.terraform_site.id
  # analytic = {
  #   enabled = true
  # }
  gateway_mgmt = {
    root_password = "pwd123"
    app_usage     = true
    auto_signature_update = {
      enable      = true
      time_of_day = "02:00"
    }
  }
  vars = {
  }
  ap_updown_threshold     = 5
  device_updown_threshold = 5
  auto_upgrade = {
    enabled     = true
    day_of_week = "tue"
    time_of_day = "02:00"
    version     = "beta"
  }
  config_auto_revert = true

  persist_config_on_device = true
  proxy = {
    url = "http://myproxy:3128"
  }
  rogue = {
    enabled          = true
    honeypot_enabled = true
    min_duration     = 5
  }
}

resource "mist_org_wlantemplate" "test101" {
  org_id = mist_org.terraform_test.id
  name   = "test101"
  applies = {
    site_ids = [
      mist_site.terraform_site.id
    ]
  }
}

resource "mist_org_wlan" "wlan_cwp" {
  ssid              = "MlN.test"
  bands             = ["5"]
  vlan_id           = "141"
  wlan_limit_up     = 10000
  wlan_limit_down   = 20000
  client_limit_up   = 512
  client_limit_down = 1000
  portal = {
    enabled                = true
    bypass_when_cloud_down = true
    auth                   = "sso"
    privacy                = false
    sso_issuer             = "https://sts.windows.net/f2532c2f-938c-4529-b6e4-aa26992b6b62/"
    sso_nameid_format      = "email"
    sso_idp_sign_algo      = "sha256"
    sso_idp_cert           = "-----BEGIN CERTIFICATE-----\nMIIC8DCCAdigAwIBAgIQE5pOI9W1DZFHbB9m2Q7ADzANBgkqhkiG9w0BAQsFADA0MTIwMAYDVQQD\nEylNaWNyb3NvZnQgQXp1cmUgRmVkZXJhdGVkIFNTTyBDZXJ0aWZpY2F0ZTAeFw0yMjAyMDIxNDEz\nMTNaFw0yNTAyMDIxNDEzMTNaMDQxMjAwBgNVBAMTKU1pY3Jvc29mdCBBenVyZSBGZWRlcmF0ZWQg\nU1NPIENlcnRpZmljYXRlMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5gQTCccB3oE7\nelNYH2+11Q69Iq/2f3qf5KUZEQKwL++HyoBCOAM3wL3uLWwvRaih4+qpAeZvNsuShNIyB08SDcWN\nYsqVxaUsLYfzDD0c9VG9mwAx0Kh01S2JvtaLCaFveac7UXVfn/E/QbPXibS1EQvHUj0hwNXMrdS4\nh4TOk4D1Q70+OnCWyy7ykG1/RuO8UerIfqkQEy4C3QFb3Cyo4E7bEaYQo0NiCqD9IoM3B0wZib8Y\n3yRGJKdzXyDxuVJFb5rF7XMAHTWWAbxaN4KOLhZnjaJla7Pu/sFAj2Npm8Hm5pYEYBaUz4fc/8kg\nIwakFb3mnbnYw0xQwf+aJss1vQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQCF+oKuLmnooDzALwaE\nbFVI7PVGhU7/UZzAnq6HHI9ngF0Af2+uIrvAz6rdUM1bsGhRbj3SV2oaj26pe/1TDrGescXWhTPw\nKcXOwBnVmFr8FlMkozwpHRNzCQyFYGiTAztgQcmtwF7pilVndOmEc+p3LvCdI5JZB+LtMM/o9V+1\n+Yhm4MEWO6wTSY+j7goc/vi5f76TDZPN6PkRv17+EkybEudJuTOuIoNiqsAbNB52bVNHtxFHGIwb\nH9iS45QJ4/RG1WUr91xe3Vzh/fp1BkiHZVL4iOywOIF0TYcW7h958JEf+q0HD5LUMO47NPEbc/Cd\n+fVCTXWzABXXy4D+S8gA\n-----END CERTIFICATE-----"
    sso_idp_sso_url        = "https://login.microsoftonline.com/f2532c2f-938c-4529-b6e4-aa26992b6b62/saml2"
    email_enabled          = true
  }
  portal_allowed_hostnames = [
    "login.microsoftonline.com",
    "portal.mist.com",
    "login.live.com",
    "aadcdn.msauth.net",
    "logincdn.msauth.net"
  ]
  auth = {
    type = "psk"
    psk  = "Juniper123"
  }
  apply_to    = "site"
  org_id      = mist_org.terraform_test.id
  template_id = mist_org_wlantemplate.test101.id
  interface   = "all"
}




# # ################### SITE LEVEL


resource "mist_org_networktemplate" "switch_template" {
  name                   = "test_switch"
  org_id                 = mist_org.terraform_test.id
  dns_servers            = ["10.3.51.222"]
  dns_suffix             = ["stag.one"]
  ntp_servers            = ["10.3.51.222"]
  additional_config_cmds = ["set system hostnam test", "set system services ssh root-login allow"]

  extra_routes = {
    "1.2.0.0/24" = {
      via = "1.2.3.4"
    }
  }
  networks = {
    test = {
      subnet  = "1.2.3.0/24"
      vlan_id = 10
    }
    test2 = {
      subnet  = "1.2.3.0/24"
      vlan_id = 11
    }
  }
  radius_config = {
    acct_interim_interval = 60
    coa_enabled           = true
    network               = "test"
    acct_servers = [
      {
        host   = "1.2.3.4"
        secret = "secret"
      }
    ]
    auth_servers = [
      {
        host   = "1.2.3.4"
        secret = "secret"
      }
    ]
  }
  port_usages = {
    disabled_port = {
      mode         = "access"
      disable      = true
      port_network = "default"
    }
  }
  remote_syslog = {
    archive = {
      files = 2
      size  = "5m"
    }
    console = {
      contents = [
        {
          facility = "kernel"
          severity = "alert"
        }
      ]
    }
    enabled             = true
    network             = "test"
    send_to_all_servers = false
    servers = [
      {
        contents = [
          {
            facility = "any"
            severity = "any"
          }
        ]
        host = "1.2.3.4"
      }
    ]
    structured_data = true
  }
  switch_mgmt = {
    config_revert = 5
    protect_re = {
      enabled = true
    }
    root_password = "Juniper123!"
  }
  switch_matching = {
    enable = true
    rules = [
      {
        match_type  = "match_name[0:3]"
        match_value = "abc"
        additional_config_cmds = [
          "set system name-server 8.8.8.8"
        ]
        match_role = "access"
        name       = "access"
        port_config = {
          "ge-0/0/0-10" = {
            usage = "trunk"
          }
        }
      },
      {
        additional_config_cmds = [
          "set system name-server 8.8.8.8"
        ]
        match_role = "core"
        name       = "core"
        port_config = {
          "ge-0/0/0-10" = {
            usage = "trunk"
          }
        }
      }
    ]
  }
}

# ################### SITE LEVEL


resource "mist_site_networktemplate" "site_switch_template" {
  site_id                = mist_site.terraform_site.id
  dns_servers            = ["10.3.51.222"]
  dns_suffix             = ["stag.one"]
  ntp_servers            = ["10.3.51.222"]
  additional_config_cmds = ["set system hostnam test", "set system services ssh root-login allow"]
  networks = {
    test = {
      subnet  = "1.2.3.0/24"
      vlan_id = 10
    }
    test2 = {
      subnet  = "1.2.3.0/24"
      vlan_id = 11

    }
  }
  radius_config = {
    acct_interim_interval = 60
    coa_enabled           = true
    network               = "test"
    acct_servers = [
      {
        host   = "1.2.3.4"
        secret = "secret"
      }
    ]
    auth_servers = [
      {
        host   = "1.2.3.4"
        secret = "secret"
      }
    ]
  }
  port_usages = {
    trunk = {
      all_networks = true
      description  = "profile for trunk ports"
      enable_qos   = true
      mode         = "trunk"
      port_network = "test2"
    }
    disabled_port = {
      mode         = "access"
      disabled     = true
      port_network = "default"
    }
  }
  remote_syslog = {
    archive = {
      files = 2
      size  = "5m"
    }
    console = {
      contents = [
        {
          facility = "kernel"
          severity = "alert"
        }
      ]
    }
    enabled             = true
    network             = "test"
    send_to_all_servers = false
    servers = [
      {
        contents = [
          {
            facility = "any"
            severity = "any"
          }
        ]
        host = "1.2.3.4"
      }
    ]
    structured_data = true
  }
  switch_mgmt = {
    config_revert = 5
    protect_re = {
      enabled = true
    }
    root_password = "Juniper123"
  }
  switch_matching = {
    enable = true
    rules = [
      {
        match_type  = "match_name[0:3]"
        match_value = "abc"
        additional_config_cmds = [
          "set system name-server 8.8.8.8"
        ]
        match_role = "access"
        name       = "access"
        port_config = {
          "ge-0/0/0-10" = {
            usage = "trunk"
          }
        }
      },
      {
        additional_config_cmds = [
          "set system name-server 8.8.8.8"
        ]
        match_role = "core"
        name       = "core"
        port_config = {
          "ge-0/0/0-10" = {
            usage = "trunk"
          }
        }
      }
    ]
  }
}
resource "mist_org_wlan" "wlan_one" {
  ssid              = "wlan_one"
  org_id            = mist_org.terraform_test.id
  template_id       = mist_org_wlantemplate.test101.id
  bands             = ["5", "6"]
  vlan_id           = 143
  wlan_limit_up     = 10000
  wlan_limit_down   = 20000
  client_limit_up   = 512
  client_limit_down = 1000
  auth = {
    type = "psk"
    psk  = "secretpsk"
  }
  interface = "all"
  dynamic_vlan = {
    enabled = true
    type    = "standard"
    vlans = {
      460 = ""
      492 = ""
      494 = ""
    }
    default_vlan_id = "494"
  }
}
resource "mist_site_wlan" "wlan_cwp2" {
  ssid              = "MlN.test"
  bands             = ["5"]
  vlan_id           = 143
  wlan_limit_up     = 10000
  wlan_limit_down   = 20000
  client_limit_up   = 512
  client_limit_down = 1000
  portal = {
    enabled                = true
    bypass_when_cloud_down = true
    auth                   = "sso"
    privacy                = false
    sso_issuer             = "https://sts.windows.net/f2532c2f-938c-4529-b6e4-aa26992b6b62/"
    sso_nameid_format      = "email"
    sso_idp_sign_algo      = "sha256"
    sso_idp_cert           = "-----BEGIN CERTIFICATE-----\nMIIC8DCCAdigAwIBAgIQE5pOI9W1DZFHbB9m2Q7ADzANBgkqhkiG9w0BAQsFADA0MTIwMAYDVQQD\nEylNaWNyb3NvZnQgQXp1cmUgRmVkZXJhdGVkIFNTTyBDZXJ0aWZpY2F0ZTAeFw0yMjAyMDIxNDEz\nMTNaFw0yNTAyMDIxNDEzMTNaMDQxMjAwBgNVBAMTKU1pY3Jvc29mdCBBenVyZSBGZWRlcmF0ZWQg\nU1NPIENlcnRpZmljYXRlMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5gQTCccB3oE7\nelNYH2+11Q69Iq/2f3qf5KUZEQKwL++HyoBCOAM3wL3uLWwvRaih4+qpAeZvNsuShNIyB08SDcWN\nYsqVxaUsLYfzDD0c9VG9mwAx0Kh01S2JvtaLCaFveac7UXVfn/E/QbPXibS1EQvHUj0hwNXMrdS4\nh4TOk4D1Q70+OnCWyy7ykG1/RuO8UerIfqkQEy4C3QFb3Cyo4E7bEaYQo0NiCqD9IoM3B0wZib8Y\n3yRGJKdzXyDxuVJFb5rF7XMAHTWWAbxaN4KOLhZnjaJla7Pu/sFAj2Npm8Hm5pYEYBaUz4fc/8kg\nIwakFb3mnbnYw0xQwf+aJss1vQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQCF+oKuLmnooDzALwaE\nbFVI7PVGhU7/UZzAnq6HHI9ngF0Af2+uIrvAz6rdUM1bsGhRbj3SV2oaj26pe/1TDrGescXWhTPw\nKcXOwBnVmFr8FlMkozwpHRNzCQyFYGiTAztgQcmtwF7pilVndOmEc+p3LvCdI5JZB+LtMM/o9V+1\n+Yhm4MEWO6wTSY+j7goc/vi5f76TDZPN6PkRv17+EkybEudJuTOuIoNiqsAbNB52bVNHtxFHGIwb\nH9iS45QJ4/RG1WUr91xe3Vzh/fp1BkiHZVL4iOywOIF0TYcW7h958JEf+q0HD5LUMO47NPEbc/Cd\n+fVCTXWzABXXy4D+S8gA\n-----END CERTIFICATE-----"
    sso_idp_sso_url        = "https://login.microsoftonline.com/f2532c2f-938c-4529-b6e4-aa26992b6b62/saml2"
    email_enabled          = true
  }
  portal_allowed_hostnames = [
    "login.microsoftonline.com",
    "portal.mist.com",
    "login.live.com",
    "aadcdn.msauth.net",
    "logincdn.msauth.net"
  ]
  auth = {
    type = "psk"
    psk  = "Juniper123"
  }
  apply_to  = "site"
  interface = "all"
  site_id   = mist_site.terraform_site2.id
}

resource "mist_device_switch" "test_switch" {
  managed = true
  role    = "test"
  networks = {
    "prx" = {
      vlan_id = "18"
    }
  }
  port_usages = {
    "prx" = {
      mode         = "trunk"
      disabled     = false
      port_network = "default"
      stp_edge     = false
      all_networks = false
      networks = [
        "default",
        "prx"
      ],
      speed         = "auto"
      duplex        = "auto"
      mac_limit     = 0
      poe_disabled  = false
      enable_qos    = false
      storm_control = {}
      description   = ""
    },
    "autogenerated_ge-0_0_10" = {
      mode         = "trunk"
      all_networks = false
      networks = [
        "default",
        "prx"
      ],
      port_network = "prx"
      autoneg      = true
      disabled     = false
      poe_disabled = false
      duplex       = "auto"
      speed        = "auto"
    },
    "autogenerated_ge-0_0_11" = {
      mode         = "trunk"
      all_networks = false
      networks = [
        "default",
        "prx"
      ],
      port_network = "prx"
      autoneg      = true
      disabled     = false
      poe_disabled = false
      duplex       = "auto"
      speed        = "auto"
    }
  }
  additional_config_cmds = [
    "set groups top system proxy password \"$9$VpsgajHmFnCq.pBIEeK4aZUDin6Ctu136eW\"",
    "",
    "annotate system \" -- custom-auth -- Template level --\"",
    "delete apply-groups custom-auth",
    "delete groups custom-auth",
    "set groups custom-auth",
    "set groups custom-auth  system login user tmunzer class super-user",
    "set groups custom-auth  system login user tmunzer authentication ssh-rsa \"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCxk1CXmrc/wNlzuelQoBgEdnA3Mq8r05UhXgBUwqFdcbwkEmgYnciQWXYeFdA+UwR3SWX4VtyeenQi/S9pSOjvoZQ5OYBWNqak+AWjXXzX7sO8cf59TSDDN2lybdH6kWbvRKYzMbfpZvwluYwQX7ftN/D4/iS7288i5knaMVklhDJGdcuh6Xlve6PueZ5ov2twPzpplPVdWOCqgaCMhvVplG3oNTvBpeJ6BCXzhgJdkFNinFehPnvUR6MRwdjD/MvS1r+yrggHGQzJ57VEDsw2qH8wRaUUn0+Nd5hMYO4vamfintlm5hEAW9+my6HKNzBmd/TMeLctKUOg/q4Bc9TfYB51HEghHZ1i3ClJ6nE/1M6Ev9B+YwBJpeRaieWHbOHWrmMKM5rnYpX+uShPobvHKC8Z5HU5UkAtWvRyXVqbeeVv7xlGDTKnm0UQtzv9UDR/BN3VSnh8JKGLeiPDRZ0oDdgqK1ijxXBCFKb6FXAUiORUQLtI/7dmV4YxK4oS+ZRU/uxCCNDx6R90o08RFDFpWNBPY1yj332T/6JnA7nzlAYdQ1s1hzgzamZoEZuO/a+NObV9gQbJC13ZtZEWKXSLHPvEsORc5nb7hN76RNI7yKnAWNy5Zj2aaPlTGofbdr9M8+Tez4wldmYB/k86M4lENv5fZRx9A6VmfXFDoSI33w== tmunzer@stag.one\"",
    "set apply-groups custom-auth"
  ]
  ip_config = {
    type    = "static"
    ip      = "10.3.18.99"
    netmask = "255.255.255.0"
    network = "prx"
    gateway = "10.3.18.11"
  }
  # routing_policies= {
  #     "tyes"= {
  #         terms= [
  #             {
  #                 matching= {
  #                     prefix= [
  #                         "0.0.0.0/0"
  #                     ],
  #                     protocol= [
  #                         "ospf"
  #                     ]
  #                 },
  #                 actions= {
  #                     accept= true
  #                 },
  #                 name= "gfds"
  #             }
  #         ]
  #     }
  # }
  port_config = {
    "ge-0/0/0" = {
      usage                = "prx"
      critical             = false
      description          = ""
      "no_local_overwrite" = true
    },
    "ge-0/0/10" = {
      usage              = "autogenerated_ge-0_0_10"
      no_local_overwrite = false
      speed              = "100m"
    },
    "ge-0/0/11" = {
      usage                = "default"
      port_network         = "prx"
      critical             = false
      description          = ""
      "no_local_overwrite" = false
    }
  }
  port_mirroring = {
    "test" = {
      output_port_id = "ge-0/0/10"
      input_port_ids_ingress = [
        "ge-0/0/2"
      ],
      input_port_ids_egress = [
        "ge-0/0/2"
      ],
      input_networks_ingress = [
        "default"
      ]
    },
    "test2" = {
      output_network = "prx"
      input_networks_ingress = [
        "default"
      ]
    }
  }
  mist_nac = {
    enabled = true
  }
  vrf_instances = {
    "fds" = {
      networks = [
        "prx"
      ],
      extra_routes = {
        "1.2.0.0/24" = {
          via = "1.2.3.4"
        }
      }
    }
  }
  vrf_config = {
    enabled = true
  }
  dhcpd_config = {
    enabled = true
    "prx" = {
      type = "relay"
      servers = [
        "1.2.3.4"
      ]
    }
  }
  # ospf_areas= {
  #     "0"= {
  #         type= "default"
  #         networks= {
  #             "prx"= {
  #                 passive= false
  #                 interface_type= "p2p"
  #                 hello_interval= 10
  #                 dead_interval= 40
  #                 auth_type= "password"
  #                 auth_password= "hpihpi"
  #                 metric= 1
  #                 bfd_minimum_interval= 1
  #             }
  #         },
  #         include_loopback= false
  #     }
  # }
  ospf_config = {
    enabled             = true
    reference_bandwidth = "10M"
    areas = {
      "0" = {
        no_summary = false
      }
    }
  }
  router_id = "1.2.3.4"
  oob_ip_config = {
    type                      = "static"
    ip                        = "2.2.2.2"
    netmask                   = "/24"
    use_mgmt_vrf              = true
    use_mgmt_vrf_for_host_out = true
    gateway                   = "2.2.2.1"
  }
  device_id = mist_org_inventory.inventory.devices[1].id
  name      = "demo-ex"
  site_id   = mist_org_inventory.inventory.devices[1].site_id
}


resource "mist_device_gateway" "srx" {

  name      = "srx"
  device_id = mist_org_inventory.inventory.devices[2].id
  site_id   = mist_org_inventory.inventory.devices[2].site_id

  port_config = {
    "ge-0/0/3" = {
      name       = "FTTH"
      usage      = "wan"
      aggregated = false
      redundant  = false
      critical   = false
      wan_type   = "broadband"
      ip_config = {
        type    = "static"
        ip      = "192.168.1.8"
        netmask = "/24"
        gateway = "192.168.1.1"
      },
      disable_autoneg = false
      speed           = "auto"
      duplex          = "auto"
      wan_source_nat = {
        disabled = false
      },
      vpn_paths = {
        "SSR_HUB_DC-MPLS.OrgOverlay" = {
          key         = 0
          role        = "spoke"
          bfd_profile = "broadband"
        }
      }
    },
    "ge-0/0/4" = {
      name       = "LTE"
      usage      = "wan"
      aggregated = false
      redundant  = false
      critical   = false
      wan_type   = "broadband"
      ip_config = {
        type = "dhcp"
      },
      disable_autoneg = false
      speed           = "auto"
      duplex          = "auto"
      wan_source_nat = {
        disabled = false
      },
      vpn_paths = {
        "SSR_HUB_DC-MPLS.OrgOverlay" = {
          key         = 0
          role        = "spoke"
          bfd_profile = "broadband"
        }
      }
    },
    "ge-0/0/5" = {
      usage            = "lan"
      critical         = false
      aggregated       = true
      ae_disable_lacp  = false
      ae_lacp_force_up = true
      ae_idx           = 0
      redundant        = false
      networks = [
        "PRD-Core"
      ]
    },
    "ge-0/0/7" = {
      usage      = "lan"
      critical   = false
      aggregated = false
      redundant  = false
      networks = [
        "PRD-Mgmt"
      ]
    },
    "ge-0/0/6" = {
      usage      = "lan"
      critical   = false
      aggregated = false
      redundant  = false
      networks = [
        "PRD-Lab"
      ]
    }
  }
  ip_configs = {
    "PRD-Core" = {
      type    = "static"
      ip      = "10.3.100.9"
      netmask = "/24"
    },
    "PRD-Mgmt" = {
      type    = "static"
      ip      = "10.3.172.1"
      netmask = "/24"
    },
    "PRD-Lab" = {
      type    = "static"
      ip      = "10.3.171.1"
      netmask = "/24"
    }
  }
  dhcpd_config = {
    enable = true
  }
  path_preferences = {
    "HUB" = {
      strategy = "ordered"
      paths = [
        {
          name = "SSR_HUB_DC-MPLS.OrgOverlay"
          type = "vpn"
        }
      ]
    },
    "HUB-ORDERED" = {
      strategy = "ordered"
      paths = [
        {
          name     = "SSR_HUB_DC-MPLS.OrgOverlay"
          wan_name = "FTTH"
          type     = "vpn"
        },
        {
          name     = "SSR_HUB_DC-MPLS.OrgOverlay"
          wan_name = "LTE"
          type     = "vpn"
        }
      ]
    },
    "HUB-ECMP" = {
      strategy = "weighted"
      paths = [
        {
          name     = "SSR_HUB_DC-MPLS.OrgOverlay"
          wan_name = "LTE"
          cost     = 30
          type     = "vpn"
        },
        {
          name     = "SSR_HUB_DC-MPLS.OrgOverlay"
          wan_name = "FTTH"
          cost     = 30
          type     = "vpn"
        }
      ]
    }
  }
  service_policies = [
    {
      name = "Policy-14"
      tenants = [
        "PRD-Core"
      ],
      services = [
        "any"
      ],
      action          = "allow"
      path_preference = "HUB"
      idp = {
        enabled    = true
        profile    = "critical"
        alert_only = false
      }
    },
    {
      name = "Policy-2"
      tenants = [
        "PRD-Mgmt"
      ],
      services = [
        "any"
      ],
      action          = "allow",
      path_preference = "HUB-ECMP"
      idp = {
        enabled    = true,
        profile    = "standard"
        alert_only = true
      }
    },
    {
      name = "Policy-3"
      tenants = [
        "PRD-Lab"
      ],
      services = [
        "any"
      ],
      action          = "allow"
      path_preference = "HUB-ORDERED"
      idp = {
        enabled = false
      }
    }
  ]
}

resource "mist_org_deviceprofile_ap" "app_one" {
  name         = "app_one"
  org_id       = mist_org.terraform_test.id
  disable_eth1 = true
}

resource "mist_org_deviceprofile_gateway" "hubp_one" {
  name   = "hubp_one2"
  org_id = mist_org.terraform_test.id

  port_config = {
    "ge-0/0/3" = {
      name       = "FTTH"
      usage      = "wan"
      aggregated = false
      redundant  = false
      critical   = false
      wan_type   = "broadband"
      ip_config = {
        type    = "static"
        ip      = "192.168.1.8"
        netmask = "/24"
        gateway = "192.168.1.1"
      },
      disable_autoneg = false
      speed           = "auto"
      duplex          = "auto"
      wan_source_nat = {
        disabled = false
      },
      vpn_paths = {
        "SSR_HUB_DC-MPLS.OrgOverlay" = {
          key         = 0
          role        = "spoke"
          bfd_profile = "broadband"
        }
      }
    },
    "ge-0/0/4" = {
      name       = "LTE"
      usage      = "wan"
      aggregated = false
      redundant  = false
      critical   = false
      wan_type   = "broadband"
      ip_config = {
        type = "dhcp"
      },
      disable_autoneg = false
      speed           = "auto"
      duplex          = "auto"
      wan_source_nat = {
        disabled = false
      },
      vpn_paths = {
        "SSR_HUB_DC-MPLS.OrgOverlay" = {
          key         = 0
          role        = "spoke"
          bfd_profile = "broadband"
        }
      }
    },
    "ge-0/0/5" = {
      usage            = "lan"
      critical         = false
      aggregated       = true
      ae_disable_lacp  = false
      ae_lacp_force_up = true
      ae_idx           = 0
      redundant        = false
      networks = [
        "PRD-Core"
      ]
    },
    "ge-0/0/7" = {
      usage      = "lan"
      critical   = false
      aggregated = false
      redundant  = false
      networks = [
        "PRD-Mgmt"
      ]
    },
    "ge-0/0/6" = {
      usage      = "lan"
      critical   = false
      aggregated = false
      redundant  = false
      networks = [
        "PRD-Lab"
      ]
    }
  }
  ip_configs = {
    "PRD-Core" = {
      type    = "static"
      ip      = "10.3.100.9"
      netmask = "/24"
    },
    "PRD-Mgmt" = {
      type    = "static"
      ip      = "10.3.172.1"
      netmask = "/24"
    },
    "PRD-Lab" = {
      type    = "static"
      ip      = "10.3.171.1"
      netmask = "/24"
    }
  }
  dhcpd_config = {
    enable = true
  }
  path_preferences = {
    "HUB" = {
      strategy = "ordered"
      paths = [
        {
          name = "SSR_HUB_DC-MPLS.OrgOverlay"
          type = "vpn"
        }
      ]
    },
    "HUB-ORDERED" = {
      strategy = "ordered"
      paths = [
        {
          name     = "SSR_HUB_DC-MPLS.OrgOverlay"
          wan_name = "FTTH",
          type     = "vpn"
        },
        {
          name     = "SSR_HUB_DC-MPLS.OrgOverlay"
          wan_name = "LTE"
          type     = "vpn"
        }
      ]
    },
    "HUB-ECMP" = {
      strategy = "weighted"
      paths = [
        {
          name     = "SSR_HUB_DC-MPLS.OrgOverlay"
          wan_name = "LTE"
          cost     = 30
          type     = "vpn"
        },
        {
          name     = "SSR_HUB_DC-MPLS.OrgOverlay"
          wan_name = "FTTH"
          cost     = 30
          type     = "vpn"
        }
      ]
    }
  }
  service_policies = [
    {
      name = "Policy-14"
      tenants = [
        "PRD-Core"
      ],
      services = [
        "any"
      ],
      action          = "allow"
      path_preference = "HUB"
      idp = {
        enabled    = true
        profile    = "critical"
        alert_only = false
      }
    },
    {
      name = "Policy-2"
      tenants = [
        "PRD-Mgmt"
      ],
      services = [
        "any"
      ],
      action          = "allow",
      path_preference = "HUB-ECMP"
      idp = {
        enabled    = true,
        profile    = "standard"
        alert_only = true
      }
    },
    {
      name = "Policy-3"
      tenants = [
        "PRD-Lab"
      ],
      services = [
        "any"
      ],
      action          = "allow"
      path_preference = "HUB-ORDERED"
      idp = {
        enabled = false
      }
    }
  ]
}

resource "mist_org_deviceprofile_assign" "hub" {
  org_id           = mist_org.terraform_test.id
  deviceprofile_id = mist_org_deviceprofile_gateway.hubp_one.id
  macs = [
    "4c96143de700",
    "e8a24550e732"
  ]
}

resource "mist_org_network" "CORE_VLAN410" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "0.0.0.0/0"
  internet_access = {
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  disallow_mist_services = true
  name                   = "CORE_VLAN410"
}
resource "mist_org_network" "RESTRICTED_VLAN420" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "0.0.0.0/0"
  internet_access = {
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  disallow_mist_services = true
  name                   = "RESTRICTED_VLAN420"
}
resource "mist_org_network" "AUTOMATION_VLAN460" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "0.0.0.0/0"
  internet_access = {
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  disallow_mist_services = true
  name                   = "AUTOMATION_VLAN460"
}
resource "mist_org_network" "UNTRUSTED_VLAN490" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "0.0.0.0/0"
  internet_access = {
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  disallow_mist_services = true
  name                   = "UNTRUSTED_VLAN490"
}
resource "mist_org_network" "MANAGEMENT_VLAN430" {
  org_id    = mist_org.terraform_test.id
  isolation = true
  subnet    = "0.0.0.0/0"
  internet_access = {
  }
  vpn_access = {
    OrgOverlay = {
      routed = true
    }
  }
  disallow_mist_services = true
  name                   = "MANAGEMENT_VLAN430"
}

resource "mist_org_gatewaytemplate" "gatewaytemplate_one" {
  type   = "spoke"
  name   = "gatewaytemplate_one"
  org_id = mist_org.terraform_test.id
  vrf_instances = {
    VRF_ONE = {
      networks = [
        "PRD-CORE_VLAN410",
        "RESTRICTED_VLAN420"
      ]
    }
  }
  vrf_config = {
    enabled = true
  }
  port_config = {
    "ge-0/0/7,ge-5/0/7" = {
      usage     = "lan"
      redundant = true
      networks = [
        "CORE_VLAN410",
        "RESTRICTED_VLAN420",
        "AUTOMATION_VLAN460",
        "UNTRUSTED_VLAN490",
        "MANAGEMENT_VLAN430",
      ]
      reth_idx  = "3"
      reth_node = "node0"
    }
    "ge-0/0/3" = {
      name       = "FTTH"
      usage      = "wan"
      aggregated = false
      redundant  = false
      critical   = false
      wan_type   = "broadband"
      ip_config = {
        type    = "static"
        ip      = "192.168.1.8"
        netmask = "/24"
        gateway = "192.168.1.1"
      }
      disable_autoneg = false
      speed           = "auto"
      duplex          = "auto"
      wan_source_nat = {
        disabled = false
      }
    }
  }
  ip_configs = {
    "CORE_VLAN410" = {
      type    = "static"
      ip      = "10.4.171.1"
      netmask = "/24"
    }
    "RESTRICTED_VLAN420" = {
      type    = "static"
      ip      = "10.5.171.1"
      netmask = "/24"
    }
    "AUTOMATION_VLAN460" = {
      type    = "static"
      ip      = "10.6.171.1"
      netmask = "/24"
    }
    "UNTRUSTED_VLAN490" = {
      type    = "static"
      ip      = "10.7.171.1"
      netmask = "/24"
    }
    "MANAGEMENT_VLAN430" = {
      type    = "static"
      ip      = "10.8.171.1"
      netmask = "/24"
    }
  }
  path_preferences = {
    HUB = {
      paths = [
        {
          type = "wan"
          name = "FTTH"
        }
      ]
    }
  }
  service_policies = [
    {
      name            = "Policy-14"
      tenants         = ["MANAGEMENT_VLAN430"]
      services        = ["any"]
      action          = "allow"
      path_preference = "HUB"
      idp = {
        enabled    = true
        profile    = "critical"
        alert_only = false
      }
    },
    {
      name             = "lan"
      servicepolicy_id = mist_org_servicepolicy.test5.id
      path_preference  = "HUB"
    }
  ]
}
