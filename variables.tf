terraform {
  experiments = ["module_variable_optional_attrs"]
}

variable "TFC_PROJECT_NAME" {
  default = "Default Project"
}

variable "TFC_WORKSPACE_NAME" {
  type = string
}

variable "TFC_RUN_ID" {
  type = string
}

variable "subscription_id" {
  type = string
  description = "Subscription Id"
}

variable "resource_group_name" {
  type = string
  description = "Resource group name"
}

variable "location" {
  type = string
  description = "Azure Location"
}

variable "uami_id" {
  type = string
  description = "User Assigned Identity ID"
}

variable "account_name" {
  type = string
  description = "Name of CosmosDB account"
}

variable "offer_type" {
  type = string
  description = "Offer Type for Cosmos DB"
  default = "Standard"
}

variable "total_throughput_limit" {
  type = number
  description = "CosmosDB account level throughput limit"
  default = 1000
}

variable "consistency_level" {
  type = string
  description = "Consistency Level of Cosmos DB account"
  default = "Session"
}

variable "backup_type" {
  type = string
  description = "Backup Type of Cosmos DB account"
  default = "Continuous"
}

variable "backup_tier" {
  type = string
  description = "Backup Tier of Cosmos DB account"
  default = "Continuous7days"
}

variable "minimal_tls_version" {
  type = string
  description = "TLS version Cosmos DB account"
  default = "Tls12"
}

variable "mongo_server_version" {
  type = string
  description = "Mongo Server Version"
  default = "4.0"
}

variable "capabilities" {
  type = list(string)
  description = "Capabilities of Mongo DB"
  default = ["EnableMongo"]
}

variable "max_interval_in_seconds" {
  type = number
  description = "Max interval in seconds for Cosmos DB account consistency policy"
  default = 300
}

variable "max_staleness_prefix" {
  type = number
  description = "Max stateless prefix for Cosmos DB account consistency policy"
  default = 100000
}

variable "failover_priority" {
  type = number
  description = "Failover priority for Cosmos DB geolocation settings"
  default = 0
}

variable "mongo_databases" {
  description = "MongoDB databases configuration"
  type = map(object({
    name  = string
    throughput = optional(number)
    autoscale_max_throughput = optional(number)
  }))
}

variable "mongo_collections" {
  description = "MongoDB collections configuration"
  type = map(object({
    name  = string
    database_name = string
    default_ttl_seconds = number
    shard_key = string
    throughput = optional(number)
    autoscale_max_throughput = optional(number)
    keys = list(string)
  }))
}

variable "private_endpoint_name" {
  type = string
  description = "Name of Private Endpoint"
}

variable "pe_resource_group_name" {
  type = string
  description = "Private Endpoint resource group name"
}

variable "virtual_network_resource_group_name" {
  type = string
  description = "Name of Resource Group of Virtual Network"
}

variable "virtual_network_name" {
  type = string
  description = "Virtual Network Name"
}

variable "subnet_name" {
  type = string
  description = "Subnet Name"
}

variable "key_vault_id" {
  type = string
  description = "Key Vault ID"
}

variable "encryption_key_name" {
  type = string
  description = "Encryption Key name"
}
