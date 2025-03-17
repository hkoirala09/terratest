// Naming convention module for standardizing resource names
module "naming-convention" {
  source  = "localterraform.com/azurepmr/naming-convention/azurerm"
  version = "1.0.10"
}

// CosmosDB MongoDB module
module "cosmosdb-mongodb" {
  source  = "localterraform.com/AzurePMR/cosmosdb-mongodb/azurerm"
  version = "1.0.2"

  # Required Variables
  subscription_id     = var.subscription_id
  resource_group_name = var.resource_group_name
  location           = var.location

  # Networking
  private_endpoint_name          = "${module.naming-convention.short_resource}-pe-${module.naming-convention.short_location}-${module.naming-convention.short_env}"
  virtual_network_resource_group_name = var.virtual_network_rg
  virtual_network_name           = var.virtual_network_name
  subnet_name                    = var.subnet_name
  pe_resource_group_name         = var.pe_resource_group_name

  # Identity
  uami_id = azurerm_user_assigned_identity.adfumi.id

  # Key Vault
  encryption_key_name = var.encryption_key_name
  key_vault_id       = data.azurerm_key_vault.key_vault.id

  # CosmosDB Account
  account_name           = "${module.naming-convention.short_resource}-cosmosdb-${module.naming-convention.short_location}-${module.naming-convention.short_env}"
  consistency_level      = var.consistency_level
  offer_type            = var.offer_type
  mongo_server_version  = var.mongo_server_version
  capabilities          = var.capabilities
  minimal_tls_version   = var.minimal_tls_version
  backup_type           = var.backup_type
  backup_tier           = var.backup_tier
  total_throughput_limit = var.total_throughput_limit
  max_staleness_prefix   = var.max_staleness_prefix
  max_interval_in_seconds = var.max_interval_in_seconds
  failover_priority      = var.failover_priority

  # CosmosDB Data Objects
  mongo_databases   = var.mongo_databases
  mongo_collections = var.mongo_collections

  # Tags
  tags = var.tags
}

#-----------------------------------------------------------#
#                     Output Values                         #
#-----------------------------------------------------------#
output "env_short" {
  value = module.naming-convention.short_env
}

output "location_short" {
  value = module.naming-convention.short_location
}

output "resource_short" {
  value = module.naming-convention.short_resource
}

output "cosmosdb_account_id" {
  value = module.cosmosdb-mongodb.account_name
}

output "cosmosdb_endpoint" {
  value = module.cosmosdb-mongodb.private_endpoint_name
}
