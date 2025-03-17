terraform {
  required_version = ">= 1.3"
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = ">= 3.0"
    }
  }
}

provider "azurerm" {
  features {}
  subscription_id = var.subscription_id
}

module "naming-convention" {
  source  = "localterraform.com/azurepmr/naming-convention/azurerm"
  version = "1.0.10"
}

module "cosmosdb-mongodb" {
  source  = "localterraform.com/azurepmr/cosmosdb-mongodb/azurerm"
  version = "1.0.2"

  TFC_PROJECT_NAME    = var.TFC_PROJECT_NAME
  TFC_WORKSPACE_NAME  = var.TFC_WORKSPACE_NAME
  TFC_RUN_ID          = var.TFC_RUN_ID
  subscription_id     = var.subscription_id
  resource_group_name = var.resource_group_name
  location            = var.location

  account_name              = var.account_name
  consistency_level         = var.consistency_level
  offer_type                = var.offer_type
  mongo_server_version      = var.mongo_server_version
  capabilities              = var.capabilities
  minimal_tls_version       = var.minimal_tls_version
  backup_type               = var.backup_type
  backup_tier               = var.backup_tier
  total_throughput_limit    = var.total_throughput_limit
  max_staleness_prefix      = var.max_staleness_prefix
  max_interval_in_seconds   = var.max_interval_in_seconds
  failover_priority         = var.failover_priority
  uami_id                   = var.uami_id
  encryption_key_name       = var.encryption_key_name
  key_vault_id              = var.key_vault_id

  mongo_databases           = var.mongo_databases
  mongo_collections         = var.mongo_collections

  private_endpoint_name               = var.private_endpoint_name
  virtual_network_resource_group_name = var.virtual_network_resource_group_name
  virtual_network_name                = var.virtual_network_name
  subnet_name                         = var.subnet_name
  pe_resource_group_name              = var.pe_resource_group_name
}
