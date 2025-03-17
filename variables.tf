#--------------------------------- Tags --------------------------------#
variable "TFC_PROJECT_NAME" {
  default = "Default Project"
  description = "Terraform Cloud Project Name"
}

variable "TFC_WORKSPACE_NAME" {
  description = "Terraform Cloud Workspace Name"
}

variable "TFC_RUN_ID" {
  description = "Terraform Cloud Run ID"
}

variable "tags" {
  type        = map(string)
  description = "Tags to be applied to module resources"
  nullable    = false # Enforcing tags as required
  default = {
    environment               = "default-env"
    project                   = "default-project"
    autoscale_max_throughput   = "1000"
  }
}

#--------------------------- Azure Context ----------------------------#
variable "subscription_id" {
  type        = string
  description = "Azure Subscription ID"
  nullable    = false
}

variable "resource_group_name" {
  type        = string
  description = "Azure Resource Group Name"
  nullable    = false
  default     = "default-rg"
}

variable "location" {
  type        = string
  description = "Azure Region Location"
  default     = "eastus"
}

#----------------------- User Assigned Identity -----------------------#
variable "uami_id" {
  type        = string
  description = "User Assigned Managed Identity ID"
  nullable    = false
  default     = "default_uami_id"
}

#------------------------ CosmosDB Account ----------------------------#
variable "account_name" {
  type        = string
  description = "CosmosDB Account Name"
  nullable    = false
  default     = "default-account"
}

variable "offer_type" {
  type        = string
  description = "CosmosDB Offer Type"
  default     = "Standard"
}

variable "total_throughput_limit" {
  type        = number
  description = "Throughput limit for CosmosDB account"
  nullable    = false
  default     = 1000
}

variable "consistency_level" {
  type        = string
  description = "Consistency Level of Cosmos DB account"
  default     = "Session"

  validation {
    condition     = contains(["Strong", "BoundedStaleness", "Session", "Eventual", "ConsistentPrefix"], var.consistency_level)
    error_message = "Allowed values: Strong, BoundedStaleness, Session, Eventual, ConsistentPrefix."
  }
}

variable "backup_type" {
  type        = string
  description = "Backup Type for CosmosDB account"
  default     = "Continuous"

  validation {
    condition     = contains(["Periodic", "Continuous"], var.backup_type)
    error_message = "Allowed values: Periodic, Continuous."
  }
}

variable "backup_tier" {
  type        = string
  description = "Backup Tier for CosmosDB account"
  default     = "Continuous7days"

  validation {
    condition     = contains(["Continuous7days", "Continuous30days"], var.backup_tier)
    error_message = "Allowed values: Continuous7days, Continuous30days."
  }
}

variable "minimal_tls_version" {
  type        = string
  description = "Minimal TLS version for CosmosDB"
  default     = "TLS1_2"
}

variable "mongo_server_version" {
  type        = string
  description = "MongoDB Server Version"
  default     = "4.0"

  validation {
    condition     = contains(["3.6", "4.0", "4.2", "4.4", "5.0"], var.mongo_server_version)
    error_message = "Allowed values: 3.6, 4.0, 4.2, 4.4, 5.0."
  }
}

variable "capabilities" {
  type        = list(string)
  description = "MongoDB Capabilities"
  default     = ["EnableMongo"]
}

variable "max_interval_in_seconds" {
  type        = number
  description = "Max interval in seconds for CosmosDB consistency policy"
  default     = 300
}

variable "max_staleness_prefix" {
  type        = number
  description = "Max staleness prefix for CosmosDB consistency policy"
  default     = 100000
}

variable "failover_priority" {
  type        = number
  description = "Failover priority for CosmosDB"
  default     = 0
}

#------------------------- CosmosDB Data Objects ----------------------#
variable "mongo_databases" {
  type = map(object({
    name                  = string
    throughput            = optional(number, null)
    autoscale_max_throughput = optional(number, null)
  }))

  description = "CosmosDB Database details"
  nullable    = true
  default = {
    database1 = {
      name                  = "default_db"
      throughput            = 400
      autoscale_max_throughput = 1000
    }
  }
}

variable "mongo_collections" {
  type = map(object({
    name                  = string
    database_name         = string
    default_ttl_seconds   = number
    shard_key             = string
    throughput            = optional(number, null)
    autoscale_max_throughput = optional(number, null)
    keys                  = list(string)
  }))

  description = "CosmosDB Collection details"
  nullable    = true
  default = {
    collection1 = {
      name                  = "orders"
      database_name         = "ecommerce"
      default_ttl_seconds   = 3600
      shard_key             = "_id"
      throughput            = 400
      autoscale_max_throughput = 1000
      keys                  = ["customer_id", "order_date"]
    }
    collection2 = {
      name                  = "customers"
      database_name         = "ecommerce"
      default_ttl_seconds   = 7200
      shard_key             = "_id"
      throughput            = 500
      autoscale_max_throughput = 1200
      keys                  = ["email", "phone"]
    }
  }
}

#---------------------- Private Endpoint Configuration ----------------#
variable "private_endpoint_name" {
  type        = string
  description = "Private Endpoint Name"
  nullable    = false
  default     = "default-private-endpoint"
}

variable "pe_resource_group_name" {
  type        = string
  description = "Private Endpoint Resource Group Name"
  nullable    = false
  default     = "default-pe-rg"
}

variable "virtual_network_resource_group_name" {
  type        = string
  description = "Resource Group for Virtual Network"
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

#----------------------- Encryption Key Settings ----------------------#
variable "key_vault_id" {
  type        = string
  description = "Azure Key Vault ID"
  nullable    = false
}

variable "encryption_key_name" {
  type        = string
  description = "Encryption Key Name for CosmosDB"
  nullable    = false
  default     = "default-encryption-key"
}
