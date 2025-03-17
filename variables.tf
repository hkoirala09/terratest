#-------------------------------- Tags --------------------------------#
variable "TFC_PROJECT_NAME" {
  default     = "Default Project"
  description = "Terraform Cloud project name"
}

variable "TFC_WORKSPACE_NAME" {
  type        = string
  description = "Terraform Cloud workspace name"
}

variable "TFC_RUN_ID" {
  type        = string
  description = "Terraform Cloud run ID"
}

variable "tags" {
  type        = map(string)
  description = "Tags to be applied to module resources"
  nullable    = true
  default = {
    environment            = "default-env"
    project                = "default_project"
    autoscale_max_throughput = "1000"
  }
}

#-------------------------------- Azure Context --------------------------------#
variable "subscription_id" {
  type        = string
  description = "Azure Subscription ID"
  nullable    = false
}

variable "resource_group_name" {
  type        = string
  description = "Resource Group Name"
  nullable    = false
  default     = "default-rg"
}

variable "location" {
  type        = string
  description = "Azure Location"
  nullable    = false
  default     = "eastus"
}

#-------------------------------- User Assigned Identity --------------------------------#
variable "uami_id" {
  type        = string
  description = "User Assigned Identity ID"
  nullable    = false
  default     = "default_uami_id"
}

#-------------------------------- CosmosDB Account --------------------------------#
variable "account_name" {
  type        = string
  description = "Name of CosmosDB account"
  nullable    = false
  default     = "default-account"
}

variable "offer_type" {
  type        = string
  description = "Offer Type for CosmosDB"
  nullable    = false
  default     = "Standard"
}

variable "total_throughput_limit" {
  type        = number
  description = "CosmosDB account level throughput limit"
  nullable    = false
  default     = 1000
}

variable "consistency_level" {
  type        = string
  description = "Consistency Level of CosmosDB account"
  default     = "Session"
}

variable "backup_type" {
  type        = string
  description = "Backup Type of CosmosDB account"
  default     = "Continuous"
}

variable "backup_tier" {
  type        = string
  description = "Backup Tier of CosmosDB account"
  default     = "Continuous7days"
}

variable "minimal_tls_version" {
  type        = string
  description = "TLS version for CosmosDB account"
  default     = "TLS1_2"
}

variable "mongo_server_version" {
  type        = string
  description = "Mongo Server Version"
  default     = "4.2"
}

variable "capabilities" {
  type        = list(string)
  description = "Capabilities of MongoDB"
  default     = ["EnableMongo"]
}

variable "max_interval_in_seconds" {
  type        = number
  description = "Max interval in seconds for CosmosDB account consistency policy"
  default     = 300
}

variable "max_staleness_prefix" {
  type        = number
  description = "Max stateless prefix for CosmosDB account consistency policy"
  default     = 100000
}

variable "failover_priority" {
  type        = number
  description = "Failover priority for CosmosDB account geolocation settings"
  default     = 0
}

#-------------------------------- Data Objects --------------------------------#
variable "mongo_databases" {
  type = map(object({
    name                     = string
    throughput               = number # Removed optional()
    autoscale_max_throughput = number # Removed optional()
  }))
  description = "CosmosDB database details"
  nullable    = true
  default = {
    database1 = {
      name                     = "default_db"
      throughput               = 400
      autoscale_max_throughput = 1000
    }
  }
}

variable "mongo_collections" {
  type = map(object({
    name                     = string
    database_name            = string
    default_ttl_seconds      = number
    shard_key                = string
    throughput               = number # Removed optional()
    autoscale_max_throughput = number # Removed optional()
    keys                     = list(string)
  }))
  description = "CosmosDB collection details"
  nullable    = true
  default = {
    collection1 = {
      name                     = "default_collection"
      database_name            = "default_db"
      default_ttl_seconds      = 3600
      shard_key                = "_id"
      throughput               = 400
      autoscale_max_throughput = 1000
      keys                     = ["field1", "field2"]
    }
  }
}

#-------------------------------- Private Endpoint --------------------------------#
variable "private_endpoint_name" {
  type        = string
  description = "Name of Private Endpoint"
  nullable    = false
  default     = "default-private-endpoint"
}

variable "pe_resource_group_name" {
  type        = string
  description = "Private Endpoint resource group name"
  nullable    = false
  default     = "default-pe-rg"
}

variable "virtual_network_resource_group_name" {
  type        = string
  description = "Name of Resource Group of Virtual Network"
  nullable    = false
  default     = "default-vnet-rg"
}

variable "virtual_network_name" {
  type        = string
  description = "Virtual Network Name"
  nullable    = false
  default     = "vnet-name"
}

variable "subnet_name" {
  type        = string
  description = "Subnet Name"
  nullable    = false
  default     = "default-vnet"
}

#-------------------------------- Encryption Key --------------------------------#
variable "key_vault_id" {
  type        = string
  description = "Key Vault ID"
  nullable    = false
}

variable "encryption_key_name" {
  type        = string
  description = "Encryption Key name"
  nullable    = false
  default     = "default-encryption-key"
}
