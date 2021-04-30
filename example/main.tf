terraform {
    required_providers {
      vyos = {
          source = "ralphschuler/terraform-provider-vyos"
          version = "~>1.0.0"
      }
    }
}

provider "vyos" {
    host = var.vyos_host
    token = var.vyos_token
}