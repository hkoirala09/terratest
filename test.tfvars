vnet_resource_group        = "rg-eus-d-73808-data2-p"
rg_name                    = "rg-eus-d-73808-data2"
subscription_id            = "a98b5696-919e-4ac6-aba3-d3dbbf34f56a"

# Private Endpoint Configuration
private_endpoint_subnet_name = "snet-eus-d-73808-data2-pe-lz2"
private_endpoint_vnet_name   = "vnet-eus-d-73808-data2"

# Key Vault Information (Ensure correct format)
key_vault_id   = "/subscriptions/a98b5696-919e-4ac6-aba3-d3dbbf34f56a/resourceGroups/rg-eus-d-73808-data2/providers/Microsoft.KeyVault/vaults/kv-eus-d-73808-dev2-01"
key_vault_name = "kv-eus-d-73808-dev2-01"

# General Configuration
location      = "eastus"
env           = "dev"
ait           = "73808"
short_name    = "dev2"

# Security and Access
public_network_enabled = false
creator_id             = "dg.ec.p_cloud_foundation_eng_dev@bofa.com"

# Azure Data Factory Account Configuration
datafactory_account_configuration = {
  "01" = {
    managed_virtual_network_enabled = true
    managed_identity_type           = ["SystemAssigned", "UserAssigned"] # Converted to a valid list format
    purview_id                      = "/subscriptions/21f120bd-0dd9-4b52-a54c-32dd..." # Ensure this is a valid Azure resource path
  }
}

# Role Assignments for Key Vault Access
role_assignments = {
  "assignment1" = {
    role         = "Key Vault Crypto Officer"
    principal_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx" # Add actual user or service principal ID
  },
  "assignment2" = {
    role         = "Key Vault Secrets Officer"
    principal_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  },
  "assignment3" = {
    role         = "Key Vault Crypto Service Encryption User"
    principal_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  }
}
