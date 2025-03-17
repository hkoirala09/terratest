terraform {
  required_version = ">= 1.3"
  experiments = [module_variable_optional_attrs]
}

variable "TFC_PROJECT_NAME" {
  type        = string
  description = "Terraform Cloud Project Name"
}

variable "TFC_WORKSPACE_NAME" {
  type        = string
  description = "Terraform Cloud Workspace Name"
}

variable "TFC_RUN_ID" {
  type        = string
  description = "Terraform Cloud Run ID"
}

variable "subscription_id" {
  type        = string
  description = "Azure Subscription ID"
}

variable "resource_group_name" {
  type        = string
  description = "Resource Group Name"
}

variable "location" {
  type        = string
  description = "Azure Region"
  default     = "eastus"
}

variable "account_name" {
  type        = string
  description = "Cosmos DB Account Name"
}

variable "consistency_level" {
  type        = string
  description = "Cosmos DB Consistency Level"
  default     = "Session"
}

variable "offer_type" {
  type        = string
  description = "Cosmos DB Offer Type"
  default     = "Standard"
}

variable "mongo_server_version" {
  type        = string
  description = "MongoDB Server Version"
  default     = "4.0"
}

variable "capabilities" {
  type        = list(string)
  description = "Capabilities for Cosmos DB"
  default     = ["EnableMongo"]
}

variable "minimal_tls_version" {
  type        = string
  description = "Minimum TLS Version"
  default     = "Tls12"
}

variable "backup_type" {
  type        = string
  description = "Backup Type for Cosmos DB"
  default     = "Continuous"
}

variable "backup_tier" {
  type        = string
  description = "Backup Tier for Cosmos DB"
  default     = "Continuous7days"
}

variable "total_throughput_limit" {
  type        = number
  description = "Total Throughput Limit for Cosmos DB"
  default     = 1000
}

variable "max_staleness_prefix" {
  type        = number
  description = "Max Staleness Prefix"
  default     = 100000
}

variable "max_interval_in_seconds" {
  type        = number
  description = "Max Interval in Seconds"
  default     = 300
}

variable "failover_priority" {
  type        = number
  description = "Failover Priority"
  default     = 0
}

variable "uami_id" {
  type        = string
  description = "User Assigned Managed Identity ID"
}

variable "encryption_key_name" {
  type        = string
  description = "Encryption Key Name"
}

variable "key_vault_id" {
  type        = string
  description = "Key Vault ID"
}

variable "mongo_databases" {
  type = map(object({
    name                        = string
    throughput                  = optional(number)
    autoscale_max_throughput    = optional(number)
  }))
  description = "Mongo Databases Configuration"
  default     = {}
}

variable "mongo_collections" {
  type = map(object({
    name                        = string
    database_name               = string
    default_ttl_seconds         = number
    shard_key                   = string
    throughput                  = optional(number)
    autoscale_max_throughput    = optional(number)
    keys                        = list(string)
  }))
  description = "Mongo Collections Configuration"
  default     = {}
}

variable "private_endpoint_name" {
  type        = string
  description = "Private Endpoint Name"
}

variable "virtual_network_resource_group_name" {
  type        = string
  description = "Virtual Network Resource Group Name"
}

variable "virtual_network_name" {
  type        = string
  description = "Virtual Network Name"
}

variable "subnet_name" {
  type        = string
  description = "Subnet Name"
}

variable "pe_resource_group_name" {
  type        = string
  description = "Private Endpoint Resource Group Name"
}