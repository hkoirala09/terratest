module "cosmosdb_mongodb" {
  source  = "github.com/gruntwork-io/terratest/modules/azure"
  version = "1.0.2"

  # Required Variables
  subscription_id    = var.subscription_id
  resource_group_name = var.resource_group_name
  location           = var.location

  # Networking
  private_endpoint_name             = "cosmos-pe"
  virtual_network_resource_group_name = var.virtual_network_rg
  virtual_network_name              = var.virtual_network_name
  subnet_name                       = var.subnet_name
  pe_resource_group_name            = var.pe_resource_group_name

  # Identity
  uami_id = azurerm_user_assigned_identity.uami.id

  # Key Vault
  encryption_key_name = var.encryption_key_name
  key_vault_id        = data.azurerm_key_vault.key_vault.id

  # CosmosDB Account
  account_name           = "cosmos-mongo-${var.environment}"
  consistency_level      = "Session"
  offer_type            = "Standard"
  mongo_server_version   = "4.2"
  minimal_tls_version    = "TLS1_2"
  backup_type            = "Continuous"
  total_throughput_limit = 1000

  # Tags
  tags = local.tags
}
